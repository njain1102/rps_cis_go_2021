package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}
func main() {

	db, err := sql.Open("mysql",
		"root:vignesh@tcp(127.0.0.1:3306)/cisdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)

	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	}

	sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"
	resv1, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := resv1.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)


	res_v2, err := db.Query("SELECT * FROM cities WHERE id = ?", 2)
	defer res_v2.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res_v2.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	} else {

		fmt.Println("No city found")
	}
}