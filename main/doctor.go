package main
import(
 "database/sql"
 "golang.org/x/crypto/bcrypt"
 "time"
 "net/http"
 "log"	
 "../sqlconnection"
 "../sessionhandler"
 "../domain"
)



// functions to get doctor info along with his appointments to be shown when he log in
func Profile(w http.ResponseWriter, r *http.Request) {
    db := sqlconnection.DbConn()
	ndocName :=sessionhandler.GetUserName(r)

	log.Println(ndocName)
	doctorInfo := domain.Doctor{}
	db.QueryRow("SELECT username, specialization, address, doctorname FROM users WHERE username=?", ndocName).Scan(
	&doctorInfo.DatabaseUsername, &doctorInfo.DatabaseSpecialization, &doctorInfo.DatabaseAddress, &doctorInfo.DatabaseDoctorname) 
	
    selDB, err := db.Query("SELECT * FROM Appointment WHERE docName=? ORDER BY id DESC", ndocName)
    if err != nil {
        panic(err.Error())
    }
    appoi := domain.Appointment{}
    appointments := []domain.Appointment{}
    for selDB.Next() {
        var id int
		var available bool
		var dateTime time.Time
        var docName, patient string
        err = selDB.Scan(&id, &docName, &available, &dateTime, &patient)
        if err != nil {
            panic(err.Error())
        }
        appoi.Id = id
        appoi.DocName = docName
        appoi.Patient = patient
		appoi.Available = available
        appoi.DateTime = dateTime
        appointments = append(appointments, appoi)
    }
	res:= domain.DocAccount{
   DoctorInfo : "Hello " + doctorInfo.DatabaseUsername + " ,Specialization: " + 
   doctorInfo.DatabaseSpecialization + " ,Address: " + doctorInfo.DatabaseAddress+" ,Doctorname: " + doctorInfo.DatabaseDoctorname,
   Appointments : appointments,
   }
	//log.Println(res)
    tmpl.ExecuteTemplate(w, "Profile", res)
    defer db.Close()
}




// sign up function to register doctor info and create an account
func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "form/signup.html")
		return
	}
	db := sqlconnection.DbConn()

	username := req.FormValue("username")
	password := req.FormValue("password")
	specialization := req.FormValue("specialization")
	address := req.FormValue("address")
	doctorname := req.FormValue("doctorname")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)


	switch {
	case err == sql.ErrNoRows:
	   
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password, specialization, address, doctorname) VALUES(?, ?, ?, ?, ?)",
		username, hashedPassword,  specialization, address, doctorname)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}
          
		res.Write([]byte("User created!"))
		return
	case err != nil:
	   
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
	  
		http.Redirect(res, req, "/", 301)
	}
}





// log in so doctor will be able to view his profile and add or delete appointements
func loginPage(res http.ResponseWriter, req *http.Request) {
   
	if req.Method != "POST" {
		http.ServeFile(res, req, "form/login.html")
		return
	}
	
	db := sqlconnection.DbConn()
	
	username := req.FormValue("username")
	password := req.FormValue("password")

	var DatabaseUsername string
	var databasePassword string
	

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&DatabaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}
	
    sessionhandler.SetSession(username, res)
	http.Redirect(res, req, "/Profile", 302)    
}




// new and insert funcs are used to create new appointement with the choosed date
func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
    db := sqlconnection.DbConn()
    if r.Method == "POST" {
        DocName := sessionhandler.GetUserName(r)
		DateTime := r.FormValue("Date")
		
        insForm, err := db.Prepare("INSERT INTO appointment(docName, available,dateTime,patient) VALUES(?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(DocName, true, DateTime ," ")
        log.Println("INSERT: DocName: " + DocName + " | DateTime: " + DateTime)
    }
    defer db.Close()
    http.Redirect(w, r, "/Profile", 302)
}



// delete is used to delete the choosed appointement
func Delete(w http.ResponseWriter, r *http.Request) {
    db := sqlconnection.DbConn()
    appoi := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM appointment WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(appoi)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/Profile", 301)
}