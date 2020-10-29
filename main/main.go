package main

import(
 "database/sql"
 "text/template"
 "net/http"
 "../sessionhandler"
)


// templates
var tmpl = template.Must(template.ParseGlob("form/*"))

func homePage(res http.ResponseWriter, req *http.Request) {
    sessionhandler.ClearSession(res)
	http.ServeFile(res, req, "form/index.html")
}

func main() {
    db, err :=sql.Open("mysql", "root:@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/PatientSignup", PatientSignupPage)
	http.HandleFunc("/PatientsLogin", PatientsLoginPage)
	http.HandleFunc("/Profile", Profile)
	http.HandleFunc("/New", New)
	http.HandleFunc("/Delete", Delete)
	http.HandleFunc("/Insert", Insert)
	http.HandleFunc("/PatientProfile", PatientProfile)
	http.HandleFunc("/Show", Show)
	http.HandleFunc("/book", Book)
	http.HandleFunc("/cancelBooking", CancelBooking)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
	
}
