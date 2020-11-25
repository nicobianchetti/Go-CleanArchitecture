package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		NewPostgresDB()
	}
}

//NewPostgresDB created new connection in postgres
func NewPostgresDB() {
	//Singleton
	once.Do(func() {
		var err error

		host := "localhost"
		port := "5432"
		user := "postgres"
		pass := "postgres"
		dbName := "api_seguridad"

		db, err = gorm.Open("postgres", "postgres://"+user+":"+pass+"@"+host+":"+port+"/"+dbName+"?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

func newMySQLDB() {
	// once.Do(func() {
	// 	var err error
	// 	db, err = gorm.Open("mysql", "edteam:edteam@tcp(localhost:7531)/godb?parseTime=true")
	// 	if err != nil {
	// 		log.Fatalf("can't open db: %v", err)
	// 	}

	// 	fmt.Println("conectado a mySQL")
	// })
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
