package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var users = map[string]string{
	"syahraazly@gmail.com": "123",
	"jennie@gmail.com":     "123",
	"rose@gmail.com":       "123",
}

var biodata = map[string]map[string]string{
	"syahraazly@gmail.com": {
		"Name":    "Syahra Zulya",
		"Age":     "20",
		"Address": "Jakarta, Indonesia",
	},
	"jennie@gmail.com": {
		"Name":    "Jennie Kim",
		"Age":     "27",
		"Address": "Surabaya, Indonesia",
	},
	"rose@gmail.com": {
		"Name":    "Rose Kim",
		"Age":     "26",
		"Address": "Bandung, Indonesia",
	},
}

func main() {

	http.HandleFunc("/", loginHandler)

	http.HandleFunc("/biodata", biodataHandler)

	http.HandleFunc("/404", notFoundHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		_ = r.FormValue("password") // diisi kosong karna tidak ada validasi password

		if _, ok := users[email]; ok  {
			http.Redirect(w, r, "/biodata?email="+email, http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

	}

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, biodata) // return biodata untuk ditampilkankan ke template

}

func biodataHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	data := biodata[email]

	tmpl, err := template.ParseFiles("templates/biodata.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
