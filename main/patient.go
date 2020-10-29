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

var db *sql.DB
var err error



// func sign up is used to register new patient 
func PatientSignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "form/patientSignup.html")
		return
	}
	db := sqlconnection.DbConn()

	username := req.FormValue("username")
	password := req.FormValue("password")
	patientname := req.FormValue("patientname")

	var patient string

	err := db.QueryRow("SELECT username FROM patients WHERE username=?", username).Scan(&patient)


	switch {
	case err == sql.ErrNoRows:
	   
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO patients(username, password, patientname) VALUES(?, ?, ?)", username, hashedPassword,  patientname)
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




// log in so patient will be able to view his profile and book or cancel appointements, view all doctors
func PatientsLoginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "form/patientsLogin.html")
		return
	}
	
	db := sqlconnection.DbConn()
	
	username := req.FormValue("username")
	password := req.FormValue("password")

	var DatabaseUsername string
	var databasePassword string
	

	err := db.QueryRow("SELECT username, password FROM patients WHERE username=?", username).Scan(&DatabaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/PatientsLogin", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/PatientsLogin", 301)
		return
	}
	log.Println(username)
	sessionhandler.SetSession(username, res)
	http.Redirect(res, req, "/PatientProfile", 302)

}


// functions to get patient info along with his appointments 
func PatientProfile(w http.ResponseWriter, r *http.Request) {
    db := sqlconnection.DbConn()
	npatientName := sessionhandler.GetUserName(r)

	log.Println(npatientName)
	patientInfo := domain.Patient{}
	db.QueryRow("SELECT username , patientname FROM patients WHERE username=?", npatientName).Scan(
	&patientInfo.DatabaseUsername, &patientInfo.DatabasePatientname) 
	
    selDB, err := db.Query("SELECT * FROM Appointment WHERE patient=? ORDER BY id DESC", npatientName)
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
	
	doctors:= getDoctors()
	
	res:= domain.PatientAccount{
   PatientInfo: "Hello " + patientInfo.DatabaseUsername + " ,Patientname: " + patientInfo.DatabasePatientname,
   Appointments : appointments,
   Doctors : doctors,
   }
	
    tmpl.ExecuteTemplate(w, "PatientProfile", res)
    defer db.Close()
}




// book and cancel appointement functions
func Book(w http.ResponseWriter, r *http.Request) {
    log.Println("start booking")
    db := sqlconnection.DbConn()
    appoi := r.URL.Query().Get("id")
	npatientName := sessionhandler.GetUserName(r)
	insForm, err := db.Prepare("UPDATE appointment SET patient=?, available=? WHERE id=?")
    if err != nil {
            panic(err.Error())
    }
    insForm.Exec(npatientName, false, appoi)
    log.Println("done book: appoi: " + appoi + " | npatientName: " + npatientName)
    
    defer db.Close()
    http.Redirect(w, r, "/PatientProfile", 301)
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
    log.Println("start cancel booking")
    db := sqlconnection.DbConn()
    appoi := r.URL.Query().Get("id")
	npatientName := sessionhandler.GetUserName(r)
	insForm, err := db.Prepare("UPDATE appointment SET patient=?, available=? WHERE id=?")
    if err != nil {
            panic(err.Error())
    }
    insForm.Exec(" ", true, appoi)
    log.Println("done cancel booking: appoi: " + appoi + " | npatientName: " + npatientName)
    
    defer db.Close()
    http.Redirect(w, r, "/PatientProfile", 301)
}





//show to view certain doctor info along with his available appointments
func Show(w http.ResponseWriter, r *http.Request) {
    db := sqlconnection.DbConn()
    ndocName := r.URL.Query().Get("docname")
     
	doctorInfo := domain.Doctor{}
	db.QueryRow("SELECT username, specialization, address, doctorname FROM users WHERE username=?", ndocName).Scan(
	&doctorInfo.DatabaseUsername, &doctorInfo.DatabaseSpecialization, &doctorInfo.DatabaseAddress, &doctorInfo.DatabaseDoctorname) 
	
    selDB, err := db.Query("SELECT * FROM Appointment WHERE docName=? AND available=? ORDER BY id DESC", ndocName ,true)
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
   DoctorInfo : "This is " + doctorInfo.DatabaseUsername + " ,Specialization: " + 
   doctorInfo.DatabaseSpecialization + " ,Address: " + doctorInfo.DatabaseAddress+" ,Doctorname: " + doctorInfo.DatabaseDoctorname,
   Appointments : appointments,
   }
	//log.Println(res)
    tmpl.ExecuteTemplate(w, "DoctorProfile", res)
    defer db.Close()
}




// to get all doctors info from database
func getDoctors()[]domain.Doctor{
    db := sqlconnection.DbConn()
    selDB, err := db.Query("SELECT * FROM users ORDER BY specialization DESC")
    if err != nil {
        panic(err.Error())
    }
    doc := domain.Doctor{}
	res := []domain.Doctor{}
    for selDB.Next() {
        var id int
        var userName, password, specialization ,address ,doctorname string
        err = selDB.Scan(&id, &userName, &password, &specialization , &address , &doctorname)
        if err != nil {
            panic(err.Error())
        }
        doc.DatabaseUsername = userName
        doc.DatabaseAddress = address
        doc.DatabaseSpecialization =specialization
		doc.DatabaseDoctorname=doctorname
		res = append(res, doc)
    }
    
    defer db.Close()
	return res
}


