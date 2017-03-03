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
)

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
		Index       int
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
			Index       int
			Public_port string
			Local_port  string
			Local_ip    string
			Proto       string
		}{
			Index:       i,
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
func (ip *iptables) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Println(r.PostForm)
	}
	ip.show(w, r)
}
func main() {
	//fmt.Println(getUserInfos())

	mux := http.NewServeMux()
	mux.Handle("/", &iptables{})

	n := negroni.Classic()
	n.Use(session.DefaultSession)
	n.UseHandler(mux)

	http.ListenAndServe(":3003", n)
}
