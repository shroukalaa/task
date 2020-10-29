### Golang MySQL 
*This website is designed to allow patients to book appointments with doctors.  
*doctor have to sign up with thier info name,Specialization,Address,full name.  
*doctors will be able to add new appointmment slots and delete appointments after being loged in.  
*patient have to sign up with thier info username, full name.  
*patients are able to view all doctors and thier info when logged in.  
*each appointment is created by the doctor with  a default available status and the choosed date   
*then when a patient can book it and it will not available any more and the patient username is added.  

#### Requires: 

* ![golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

* ![github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

* ![database/sql]

* ![github.com/gorilla/securecookie]

### How To Run 

Create a new database with  users,patients,appointment tables 
or run the 3 go files in the mysql folder
```sql
CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
	specialization VARCHAR(120),
	address VARCHAR(120),
	doctorname VARCHAR(120)
   );
   
CREATE TABLE patients(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
	patientname VARCHAR(120)
   );
   
CREATE TABLE appointment (
		id int(6) unsigned NOT NULL AUTO_INCREMENT,
		docName varchar(30) NOT NULL,
		available bool DEFAULT true,
		dateTime TIMESTAMP ,
		patient varchar(30),
		PRIMARY KEY (id)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
```

Go get both required packages listed below 

```bash
go get golang.org/x/crypto/bcrypt

go get github.com/go-sql-driver/mysql

go get github.com/gorilla/securecookie

```

Inside of **main/main.go** line **20** replace <example> with your own credentials
Inside of **sqlconnection/connection.go** line **17** replace <example> with your own credentials
```go
db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
// Replace with 
db, err = sql.Open("mysql", "myUsername:myPassword@/myDatabase")
```

* Cd to main folder
* Run the following command:

go build
go run .







