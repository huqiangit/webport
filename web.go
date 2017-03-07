package main

import (
	"fmt"
	"net/http"

	session "github.com/huqiangit/negroni_session"

	"encoding/json"
	"github.com/urfave/negroni"
	"os"
	"os/exec"

	"html/template"
	"regexp"
	"strconv"
)

type Config struct {
	Public_ip string
}

func getConfig() Config {
	f, err := os.Open("config.json")
	if err != nil {
		return Config{}
	}
	defer f.Close()

	var config Config
	d := json.NewDecoder(f)
	err = d.Decode(&config)

	return config
}

type UserInfo struct {
	Username string
	Password string
}

func getUserInfos() []UserInfo {
	f, err := os.Open("user.json")
	if err != nil {
		return nil
	}
	defer f.Close()

	var userInfos []UserInfo
	d := json.NewDecoder(f)
	err = d.Decode(&userInfos)

	return userInfos
}
func getPREROUTING() string {
	cmd := exec.Command("/bin/sh", "-c", `iptables-save > tmp && grep "A PREROUTING" tmp && rm tmp`)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	prerouting := string(stdout)
	return prerouting
}
func getPOSTROUTING() string {
	cmd := exec.Command("/bin/sh", "-c", `iptables-save > tmp && grep "A POSTROUTING" tmp && rm tmp`)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	postrouting := string(stdout)
	return postrouting
}

type iptables struct {
	tpl template.Template
}

func (ip *iptables) showall(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`-A PREROUTING -d.*\n`)
	entrys_string := re.FindAllString(getPREROUTING(), -1)

	var Entrys []*struct {
		Index       string
		Public_port string
		Local_port  string
		Local_ip    string
		Proto       string
	}
	for i, v := range entrys_string {
		re = regexp.MustCompile(`--to-destination .*\n`)
		local_ip_port_string := re.FindString(v)
		re = regexp.MustCompile(`[\d\.]+`)
		local_ip_string := re.FindAllString(local_ip_port_string, -1)[0]
		local_port_string := re.FindAllString(local_ip_port_string, -1)[1]

		re = regexp.MustCompile(`-p .* `)
		proto_string := re.FindString(v)
		re = regexp.MustCompile(`(tcp)|(udp)`)
		proto_string = re.FindString(proto_string)

		re = regexp.MustCompile(`--dport .* `)
		public_port_string := re.FindString(v)
		re = regexp.MustCompile(`[\d\.]+`)
		public_port_string = re.FindString(public_port_string)

		entry := struct {
			Index       string
			Public_port string
			Local_port  string
			Local_ip    string
			Proto       string
		}{
			Index:       strconv.Itoa(i),
			Public_port: public_port_string,
			Local_port:  local_port_string,
			Local_ip:    local_ip_string,
			Proto:       proto_string,
		}
		Entrys = append(Entrys, &entry)
	}
	tHeader := template.Must(template.ParseFiles("./templ/header.tpl"))
	tHeader.Execute(w, nil)

	tIptable := template.Must(template.ParseFiles("./templ/iptable.tpl"))
	tIptable.Execute(w, Entrys)
	tNewEntry := template.Must(template.ParseFiles("./templ/newentry.tpl"))
	tNewEntry.Execute(w, nil)

	tMonitor := template.Must(template.ParseFiles("./templ/monitor.tpl"))
	tMonitor.Execute(w, Entrys)
	tTail := template.Must(template.ParseFiles("./templ/tail.tpl"))
	tTail.Execute(w, nil)

}

func validOperate(v string) (valid bool, op string) {
	if (v == "add") || (v == "del") {
		return true, v
	}
	return false, v
}
func validPublicPort(v string) (bool, string) {
	if s, err := strconv.Atoi(v); err == nil {
		if s >= 1 && s < 65535 {
			return true, v
		}
	}
	return false, v

}
func validLocalIP(v string) (bool, string) {
	re := regexp.MustCompile(`^((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))$`)
	if re.MatchString(v) {
		return true, v
	}
	return false, v

}
func validLocalPort(v string) (bool, string) {
	if s, err := strconv.Atoi(v); err == nil {
		if s >= 1 && s < 65535 {
			return true, v
		}
	}
	return false, v
}
func validProto(v string) (bool, string) {
	if (v == "tcp") || (v == "udp") {
		return true, v
	}
	return false, v
}
func validRecord(v string) (bool, string) {
	if (v == "") || (v == "on") {
		return true, v
	}
	return false, v
}
func checkNewEntryValid(r *http.Request) (bool, struct {
	sessionkey  string
	operate     string
	public_port string
	local_port  string
	local_ip    string
	proto       string
	record      string
}) {
	newEntry := struct {
		sessionkey  string
		operate     string
		public_port string
		local_port  string
		local_ip    string
		proto       string
		record      string
	}{
		operate:     r.FormValue("operate"),
		public_port: r.FormValue("new_public_port"),
		local_port:  r.FormValue("new_local_port"),
		local_ip:    r.FormValue("new_local_ip"),
		proto:       r.FormValue("new_proto"),
		record:      r.FormValue("new_record"),
	}

	if ok, _ := validOperate(newEntry.operate); !ok {
		fmt.Println("check operate fail")
		return false, newEntry
	}
	if ok, _ := validPublicPort(newEntry.public_port); !ok {
		fmt.Println(newEntry.public_port)
		fmt.Println("check public port fail")
		return false, newEntry
	}
	if ok, _ := validLocalPort(newEntry.local_port); !ok {
		fmt.Println("check local port fail")
		return false, newEntry
	}
	if ok, _ := validLocalIP(newEntry.local_ip); !ok {
		fmt.Println("check local ip fail")
		return false, newEntry
	}
	if ok, _ := validProto(newEntry.proto); !ok {
		fmt.Println("check proto fail")
		return false, newEntry
	}
	if ok, _ := validRecord(newEntry.record); !ok {
		fmt.Println("check record fail")
		return false, newEntry
	}
	fmt.Println("newEntry:", newEntry)
	return true, newEntry

}
func (ip *iptables) wrapNAT_PREROUTING(action string, public_ip string, entry struct {
	sessionkey  string
	operate     string
	public_port string
	local_port  string
	local_ip    string
	proto       string
	record      string
}) string {
	return `iptables -t nat ` + action + ` PREROUTING -d ` + public_ip + `/32 -p ` + entry.proto + ` -m ` + entry.proto + ` --dport ` + entry.public_port + ` -j DNAT --to-destination ` + entry.local_ip + `:` + entry.local_port
}
func (ip *iptables) wrapNAT_POSTROUTING(action string, public_ip string, entry struct {
	sessionkey  string
	operate     string
	public_port string
	local_port  string
	local_ip    string
	proto       string
	record      string
}) string {
	return `iptables -t nat ` + action + ` POSTROUTING -d ` + entry.local_ip + `/32 -p ` + entry.proto + ` -m ` + entry.proto + ` --dport ` + entry.local_port + ` -j SNAT --to-source ` + public_ip
}
func (ip *iptables) wrapNAT_FORWARDING(action string, entry struct {
	sessionkey  string
	operate     string
	public_port string
	local_port  string
	local_ip    string
	proto       string
	record      string
}) string {
	return `iptables -t filter ` + action + ` FORWARD -d ` + entry.local_ip + `/32 -p ` + entry.proto + ` -m ` + entry.proto + ` --dport ` + entry.local_port + ` -j ACCEPT`
}
func (ip *iptables) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if session.IsInclude(w, r) == false {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {

		fmt.Println("entry")
		if r.Method == "POST" {
			fmt.Println("this is a post")
			flag, newEntry := checkNewEntryValid(r) //this function just check format,you should check iptable

			if newEntry.operate == "add" {
				if flag {
					fmt.Println("check ok")
					fmt.Println("newEntry:", newEntry)

					cmd := exec.Command("/bin/sh", "-c", ip.wrapNAT_FORWARDING("-A", newEntry))
					stdout, err := cmd.CombinedOutput()
					fmt.Println(string(stdout))
					cmd = exec.Command("/bin/sh", "-c", ip.wrapNAT_POSTROUTING("-A", getConfig().Public_ip, newEntry))
					stdout, err = cmd.CombinedOutput()
					fmt.Println(string(stdout))
					cmd = exec.Command("/bin/sh", "-c", ip.wrapNAT_PREROUTING("-A", getConfig().Public_ip, newEntry))
					stdout, err = cmd.CombinedOutput()
					fmt.Println(stdout, err)

				} else {
					fmt.Println("check fail")
				}
			}
			if newEntry.operate == "del" {
				fmt.Println("this is a del", flag)
				if flag {

					fmt.Println("1")
					cmd := exec.Command("/bin/sh", "-c", ip.wrapNAT_FORWARDING("-D", newEntry))
					stdout, err := cmd.CombinedOutput()
					fmt.Println(string(stdout), err)
					fmt.Println("2")
					cmd = exec.Command("/bin/sh", "-c", ip.wrapNAT_POSTROUTING("-D", getConfig().Public_ip, newEntry))
					stdout, err = cmd.CombinedOutput()
					fmt.Println(string(stdout), err)
					fmt.Println("3")
					cmd = exec.Command("/bin/sh", "-c", ip.wrapNAT_PREROUTING("-D", getConfig().Public_ip, newEntry))
					stdout, err = cmd.CombinedOutput()
					fmt.Println(string(stdout), err)
					fmt.Println("4")
				}
			}
		} else {

		}
		ip.showall(w, r)
	}
}

type login struct {
}

func (l *login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !session.IsInclude(w, r) {

			username := r.FormValue("username")
			password := r.FormValue("password")

			users := getUserInfos()
			for _, v := range users {
				if v.Username == username && v.Password == password {
					session.Add(w, r, 200)
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)

	}

	if r.Method == "GET" || r.Method == "" {
		if session.IsInclude(w, r) == true {
			w.Write([]byte("you has login"))
		} else {
			t := template.Must(template.ParseFiles("./templ/login.tpl"))
			t.Execute(w, nil)
		}
	}
}

func main() {
	//fmt.Println(getUserInfos())

	mux := http.NewServeMux()
	mux.Handle("/", &iptables{})
	//mux.Handle("/script/", http.FileServer(http.Dir("./script")))

	mux.Handle("/script/", http.StripPrefix("/script", http.FileServer(http.Dir("./script"))))

	mux.Handle("/login", &login{})

	n := negroni.Classic()
	//	n.Use(session.DefaultSession)
	n.UseHandler(mux)

	http.ListenAndServe(":3001", n)
}
