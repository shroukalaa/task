package domain


// to represent the patient user
//patient have to sign up with thier info username, full name
// patients are able to view all doctors and thier info when logged in
//patients can book view and cancel appointements
type Patient struct {
  	 DatabaseUsername string 
	 DatabasePatientname string
}


// to represent the patient profile
// / this struct combine the patient info along with thier available and booked appointements 
// it also show a table of all doctors 
type PatientAccount struct {
   PatientInfo string
   Appointments []Appointment
   Doctors []Doctor

}