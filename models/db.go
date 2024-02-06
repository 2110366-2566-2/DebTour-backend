package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	// Connect to the postgres db
	var err error

	db, err = gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=DebTour port=5432 TimeZone=Asia/Bangkok"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Unable to connect to database: ", err)
	} else {
		log.Println("Connected to database")
	}

	MigrateDB()
}

func MigrateDB() {
	// Migrate the schema
	Models := []interface{}{
		&Activity{},
		&ActivityLocation{},
		&Admin{},
		&Agency{},
		&CommentReview{},
		&Issue{},
		&IssueImage{},
		&Joining{},
		&Location{},
		&Notification{},
		&Review{},
		&Suggestion{},
		&SuggestionLocation{},
		&Tour{},
		&Tourist{},
		&Transaction{},
		&TransactionPayment{},
		&User{},
	}

	db.AutoMigrate(Models...)
}

