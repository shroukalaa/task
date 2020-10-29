package domain

import (
	"time"
)

// to represent the Appointment
//each appointment is created by the doctor with  a default available status and the choosed date
//then when a patient book it, it is not available any more and the patient username is added

type Appointment struct {
    Id    int
    DocName  string
	Available bool
    Patient string
	DateTime time.Time

}
