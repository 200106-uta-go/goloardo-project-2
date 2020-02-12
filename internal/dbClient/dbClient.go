package dbClient

import (
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

//UsrLogin stores the name, email, and password for user authentication
type UsrLogin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Sunsign  string
	Login    int
	Blocked  bool
}

//AddUser adds a user (name, email, and password) to database for user authertication
func AddUser(userName string, usrEmail string, usrPassword string, usrSunsign string, loginAttempts int, isBlocked bool) {
	//returns DB object, specify the type of DB and the DB file
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	myHash := HashPswd(usrPassword)
	//Migrate the schema
	db.AutoMigrate(&UsrLogin{})
	db.Create(&UsrLogin{Name: userName, Email: usrEmail, Password: myHash, Sunsign: usrSunsign, Login: loginAttempts, Blocked: isBlocked})

}

//FindEmail checks if given email is on out database
func FindEmail(userEmail string) string {
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin
	db.Find(&currUsr, "email = ?", userEmail)
	return currUsr.Email
}

//Auth checks if given password matches hashed password on database
func Auth(userEmail string, pw string) bool {
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin
	db.Find(&currUsr, "email = ?", userEmail)
	hash := currUsr.Password
	answer := IsPassword([]byte(hash), pw)
	return answer
}

//HashPswd hashes user password
func HashPswd(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

//IsPassword determines if user password matches hashed password in database
func IsPassword(hashedPassword []byte, password string) bool {
	pw := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, pw)
	if err != nil {
		return false
	}
	return true
}

//IsBlocked determines if user password matches hashed password in database
func IsBlocked(usrEmail string) bool {
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin
	db.Find(&currUsr, "email = ?", usrEmail)
	blocked := currUsr.Blocked

	return blocked
}

//FailedAttempt checks if given email is on out database
func FailedAttempt(userEmail string) {
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin
	db.Find(&currUsr, "email = ?", userEmail)
	count := currUsr.Login + 1
	db.Model(&currUsr).Update("login", count)
	if count == 3 {
		db.Model(&currUsr).Update("blocked", true)
	}
}

func FindUsr(userEmail string) [5]string {

	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin

	db.Find(&currUsr, "email = ?", userEmail)
	var answer = [5]string{currUsr.Name, currUsr.Password, currUsr.Sunsign, string(currUsr.Login), strconv.FormatBool(currUsr.Blocked)}
	return answer

}
