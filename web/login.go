package web

import (
	"fmt"
	m "mysql"
	"net/http"
	"text/template"
)

var (
	RegisterHtmlRoutie = "/Users/terry_hsiesh/go/src/web/register.html"
	port               = "8800"
)

func Server() {
	http.HandleFunc("/register", RegisterEntrance)
	http.HandleFunc("/register/registersucess", Register)

	fmt.Println("造訪地址 localhost:8800")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("server open failed")
	}
}

func RegisterEntrance(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(RegisterHtmlRoutie)
	if err != nil {
		fmt.Println("html open fail", err)
	}
	t.Execute(w, nil)

}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	r.ParseForm()
	if r.Method == "GET" {

	} else if r.Method == "POST" {
		fmt.Println(r.Method, r.Form)

		///傳入user 的資訊
		var userinfo = map[string]string{}
		var checkempty bool
		for _, temp := range r.Form {
			for _, tempinslice := range temp {
				if tempinslice == "" {
					checkempty = true
					// t, err := template.ParseFiles(RegisterHtmlRoutie)
					// if err != nil {
					// 	fmt.Println("html open fail", err)
					// }
					// t.Execute(w, nil)

				}

			}
		}

		if checkempty == false {
			userinfo["username"] = r.Form["username"][0]
			userinfo["password"] = r.Form["password"][0]
			m.InitDB()
			m.GetUserInfo(userinfo)

			///判斷有無重複帳號
			if m.UserAccoutDouble == false {

				fmt.Fprintln(w, "註冊成功")
			} else {
				LauchRegisterHtml(w)
				fmt.Fprintln(w, "您的帳號已有人使用")
			}

		} else {
			LauchRegisterHtml(w)

			fmt.Fprintln(w, "欄位內不能為空")
		}

	}

}

func LauchRegisterHtml(w http.ResponseWriter) {
	t, err := template.ParseFiles(RegisterHtmlRoutie)
	if err != nil {
		fmt.Println("html open fail", err)
	}
	t.Execute(w, nil)
}
