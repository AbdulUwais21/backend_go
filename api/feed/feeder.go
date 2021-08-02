package feed

import (
	"backend_go/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "User Test 1",
		Email:    "user1@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "User Test 2",
		Email:    "user2@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {
	// err := db.Debug().DropTable(&models.User{}).Error
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	// fmt.Println("Return value: %v", err)
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Attaching foreign key error: %v", err)
		}
		
		
	}
}
