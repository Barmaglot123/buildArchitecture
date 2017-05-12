package datasource

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DataSource *gorm.DB

func СonnectToDatabase(){
	var err error
	DataSource, err = gorm.Open("postgres", "host=localhost dbname=my_first_databse sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}