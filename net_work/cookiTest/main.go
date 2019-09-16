package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
	// read cookie
	cookie, err := req.Cookie("testcookiename")
	if err == nil {
		cookievalue := cookie.Value
		w.Write([]byte("<b>cookie的值是：" + cookievalue + "</b>\n"))
	} else {
		w.Write([]byte("<b>读取出现错误：" + err.Error() + "</b>\n"))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>设置cookie成功。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>删除cookie成功。</b>\n"))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/readcookie", ReadCookieServer)
	http.HandleFunc("/writecookie", WriteCookieServer)
	http.HandleFunc("/deletecookie", DeleteCookieServer)

	http.ListenAndServe(":8083", nil)
}