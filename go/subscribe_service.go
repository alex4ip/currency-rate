package exchangerate

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//type subscriber struct {
//	ID    int    `json:"id"`
//	Email string `json:"email" unique:"true"`
//	// Other user fields...
//}

//type rate struct {
//	ID   int     `json:"id"`
//	Date string  `json:"date"`
//	Rate float64 `json:"rate"`
//	// Other rate fields...
//}

const DB_PATH = "./db/"
const DB_SOURCE_NAME = DB_PATH + "database.sqlite"
const DB_DRIVER_NAME = "sqlite3"

func subscribe(email string) (statusCode uint16) {

	db, err := sql.Open(DB_DRIVER_NAME, DB_SOURCE_NAME)
	if err != nil {
		panic(err)
		// Create tables if they don't exist
		dbInit()
	}
	defer db.Close() // Ensure database connection is closed at the end

	// Use the 'db' object for database operations

	// search for user in the 'users' table
	user, err := db.Exec("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		panic(err)
	}
	if user != nil {
		return 409
	}

	// Insert a new user into the 'users' table
	_, err = db.Exec("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		panic(err)
		return 510
	}

	// Insert a new rate into the 'rates' table
	_, err = db.Exec("INSERT INTO rates (date, rate) VALUES (?, ?)", "2022-01-01", 39.4272)
	if err != nil {
		panic(err)
	}
	return
}
func BbInit1() { dbInit() }
func dbInit() {
	// Check if the database file exists
	if _, err := os.Stat(DB_SOURCE_NAME); err != nil {
		// Create the database file if it doesn't exist
		dbFile, err := os.Create(DB_SOURCE_NAME)
		if err != nil {
			panic(err)
		}
		defer dbFile.Close()
		log.Printf("File for Database created: %v\n", dbFile.Name())

		//log.Printf("Database file already exists: %v\n", DB_SOURCE_NAME)

		//return
	}

	log.Printf("Database file already exists: %v\n", DB_SOURCE_NAME)

	db, err := sql.Open(DB_DRIVER_NAME, DB_SOURCE_NAME)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Ensure database connection is closed at the end

	//log.Printf("Database file created: %v\n", dbFile.Name())

	//err = createUsersTable(db)
	//if err != nil {
	//	panic(err)
	//}

	_, err = db.Exec(`
		CREATE TABLE users (
# 			id INTEGER  AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL PRIMARY KEY
		);
	`)
	if err != nil {
		panic(err)
	}

	//err = createRatesTable(db)
	//if err != nil {
	//	panic(err)
	//}

	// Proceed with your application logic
	//return err
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE users (
# 			id INTEGER  AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL PRIMARY KEY
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func createRatesTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE rates (
# 		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			date DATE NOT NULL PRIMARY KEY,
			rate FLOAT NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
