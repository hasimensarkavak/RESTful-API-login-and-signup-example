package modules

import (
	"Login_And_SignUp_Example/database"
	"Login_And_SignUp_Example/helpers"
	"log"
)

type User struct {
	ID       int    `json:"ID" binding:"required,alphanum"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func UserIsExist(email string) bool {
	var ID int
	ID = -1
	db := database.MyDB()
	err := db.QueryRow("SELECT ID FROM users WHERE email = ?", email).Scan(&ID)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	if ID == -1 {
		return false
	} else {
		log.Printf("ID si %d olan user bulundu.\n", ID)
		return true
	}
}

func CallUser(email string) (User, error) {
	var user User
	db := database.MyDB()
	err := db.QueryRow("SELECT * FROM users WHERE Email = ?", email).Scan(&user.ID, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		log.Print(err)
	}
	log.Printf("userID: %d  userName: %s  Email: %s  password: secret:) .\n", user.ID, user.FullName, user.Email)
	defer db.Close()
	return user, err
}

func CreateUser(user User) bool {
	db := database.MyDB()
	res, err := db.Exec("INSERT INTO users (FullName,Email,HashedPassword) VALUES (?,?,?)", user.FullName, user.Email, user.Password)
	defer db.Close()
	if err != nil {
		//log.Print(err)
		return false
	}
	rowCount, err := res.RowsAffected()
	helpers.ErrorCheck(err)
	log.Printf("Created User ID: %d  FullName: %s  Email: %s  Password: secret:) .\n", user.ID, user.FullName, user.Email)
	log.Printf("Inserted %d rows", rowCount)
	return true
}
