package main

import (
	"fmt"
	"net/http"
	"sort"
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

	writer.Write([]byte("\n\n----- Headers Received -----\n"))
	var keys []string
	for header := range request.Header {
		keys = append(keys, header)
	}
	sort.Strings(keys)
	for _, key := range keys {
		writer.Write([]byte(fmt.Sprintf("%s : %v \n", key, request.Header[key])))
	}
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