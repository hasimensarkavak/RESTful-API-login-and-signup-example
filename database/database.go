package database

import (
	"Login_And_SignUp_Example/helpers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func MyDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/databaseName") // should write your mysql root, password and database name
	helpers.ErrorCheck(err)
	return db
}

func DatabaseBegin() {
	db := MyDB()
	CreateUserTableDB(db)
	defer db.Close()
}

func CreateUserTableDB(db *sql.DB) {
	createStatement :=
		`CREATE TABLE IF NOT EXISTS kutuphane_projesi.users (
		  ID INT NOT NULL AUTO_INCREMENT,
		  FullName VARCHAR(45) NOT NULL,
		  Email VARCHAR(45) NOT NULL,
		  HashedPassword VARCHAR(100) NOT NULL,
		  PRIMARY KEY (ID),
		  UNIQUE INDEX ID_UNIQUE (ID ASC) VISIBLE,
		  UNIQUE INDEX Email_UNIQUE (Email ASC) VISIBLE
		)
			ENGINE = InnoDB
			DEFAULT CHARACTER SET = utf8;`
	_, err := db.Exec(createStatement)
	helpers.ErrorCheck(err)
}
