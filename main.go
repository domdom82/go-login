package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/test", test)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}


func login(writer http.ResponseWriter, request *http.Request) {

	JSESSIONID := &http.Cookie{
		Name:       "JSESSIONID",
		Value:      "0815",
		Path:       "/",
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteNoneMode,
	}
	http.SetCookie(writer, JSESSIONID)
}

func test(writer http.ResponseWriter, request *http.Request) {

	JSESSIONID, err := request.Cookie("JSESSIONID")

	if err != nil {
		writer.WriteHeader(401)
		return
	}

	if JSESSIONID.Value != "0815" {
		writer.WriteHeader(401)
		return
	}

	writer.Write([]byte("Cool! You're logged in now. JSESSIONID=" + JSESSIONID.Value))
}