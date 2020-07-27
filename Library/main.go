package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	port, _ := os.LookupEnv("PORT")
	host, _ := os.LookupEnv("HOST")
	user, _ := os.LookupEnv("USER")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")
	var db *gorm.DB
	var err error
	db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = db.AutoMigrate(&Books{}).Error
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = db.AutoMigrate(&Member{}).Error
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(port, host, user, password, dbname)
}
