package domain

// to represent the doctor user
// doctor have to sign up with thier info name,Specialization,Address,full name
//doctors will be able to add new appointmment slots and delete appointmments after being loged in
type Doctor struct {
  	 DatabaseUsername string
     DatabaseSpecialization string  
	 DatabaseAddress string  
	 DatabaseDoctorname string
}



// to represent the doctor profile
// this struct combine the doctor info along with thier available and booked appointements 
type DocAccount struct {
   DoctorInfo string
   Appointments []Appointment

}