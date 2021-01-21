package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserModel struct{
	Id int `gorm:"primary_key";"AUTO_INCREMENT"`
	Name string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func main(){
	db, err := gorm.Open("mysql", "root:vignesh@tcp(127.0.0.1:3306)/cisdb?charset=utf8&parseTime=True")

	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})

	user:=&UserModel{Name:"Parameswari",Address:"Chennai"}
	newUser:=&UserModel{Name:"Bala",Address:"Bangalore"}
	tx := db.Begin()
	//To insert or create the record.
	//NOTE: we can insert multiple records too
	db.Debug().Create(user)
	//Also we can use save that will return primary key
	resultSave:=db.Debug().Save(newUser)
	tx.Commit()
	//fmt.Println(resultSave)
    log.Println(resultSave)

	//Update Record
	db.Debug().Find(&user).Update("address", "Pune")
	//It will update user address to Pune

	// Select, edit, and save
	db.Debug().Find(&user)
	user.Address = "New Delhi"

	db.Debug().Save(&user)

	// Update with column names, not attribute names
	db.Debug().Model(&user).Update("Name", "Parameswari Bala")

	db.Debug().Model(&user).Updates(
		map[string]interface{}{
			"Name": "Vignesh",
			"Address": "Hyderabad",
		})

	// UpdateColumn()
	db.Debug().Model(&user).UpdateColumn("Address", "Pune")
	db.Debug().Model(&user).UpdateColumns(
		map[string]interface{}{
			"Name": "Jayanth",
			"Address": "Hyderabad",
		})
	// Using Find()
	db.Debug().Find(&user).Update("Address", "Shimla")

	// Batch Update
	db.Debug().Table("user_models").Where("address = ?", "calicut").Update("name", "Vijay")

	// Select records and delete it
	//db.Debug().Table("user_models").Where("address= ?", "San Diego").Delete(&UserModel{})

	db.Debug().Where("address = ?", "Bangalore").First(&user)
	log.Println(user)
	db.Debug().Where("address = ?", "Pune").Find(&user)
	log.Println(user)
	db.Debug().Where("address <> ?", "New Delhi").Find(&user)
	log.Println(user)
	// IN
	db.Debug().Where("name in (?)", []string{"Jayanth", "Manohar"}).Find(&user)
	log.Println(user)
	// LIKE
	db.Debug().Where("name LIKE ?", "%ti%").Find(&user)
	log.Println(user)
	// AND
	db.Debug().Where("name = ? AND address >= ?", "Mohan", "Pune").Find(&user)
	log.Println(user)

	//Find the record and delete it
	//db.Where("address=?", "Los Angeles").Delete(&UserModel{})

	// Select all records from a model and delete all
//	db.Debug().Model(&UserModel{}).Delete(&UserModel{})

	}
