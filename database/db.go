package database

import (
	"DebTour/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MainDB *gorm.DB

func InitDB() {
	// Connect to the postgres db
	var err error

	MainDB, err = gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=DebTour port=5432 TimeZone=Asia/Bangkok"), &gorm.Config{})

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
		&models.Activity{},
		&models.ActivityLocation{},
		&models.Admin{},
		&models.Agency{},
		&models.CommentReview{},
		&models.Issue{},
		&models.IssueImage{},
		&models.Joining{},
		&models.Location{},
		&models.Notification{},
		&models.Review{},
		&models.Suggestion{},
		&models.SuggestionLocation{},
		&models.Tour{},
		&models.Tourist{},
		&models.Transaction{},
		&models.TransactionPayment{},
		&models.User{},
	}

	MainDB.AutoMigrate(Models...)
}
