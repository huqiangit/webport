package main

import (
	"fmt"
	"net/http"

	session "github.com/huqiangit/negroni_session"

	"encoding/json"
	"github.com/urfave/negroni"
	"os"
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
func main() {
	fmt.Println(getUserInfos())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	})

	n := negroni.Classic()
	n.Use(session.DefaultSession)
	n.UseHandler(mux)

	http.ListenAndServe(":3003", n)
}
