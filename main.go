package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=golangdemo sslmode=disable password=123")
	defer db.Close()

	if err != nil {
		log.Print(err)
	}

	db.AutoMigrate(&Product{})

	db.FirstOrCreate(&Product{Code: "aaa123", Price: 2000})
	db.FirstOrCreate(&Product{Code: "bbb456", Price: 3999})
	db.FirstOrCreate(&Product{Code: "ccc789", Price: 1998})

	var product Product
	db.First(&product, "price > ?", 2000)

	fmt.Println(product)

	var products []Product
	db.Find(&products, "price >= ?", 2000)
	fmt.Println(products)

	db.Create(&Product{Code: "0001111", Price: 5003})

	var p Product
	db.First(&p, "code = ?", "0001111")
	fmt.Println("Initial p: ", p)
	db.Model(&p).Update("Price", 9999)

	fmt.Println("Updated p: ", p)

	var pFromDb Product
	db.First(&pFromDb, "code = ?", "0001111")
	fmt.Println("P from DB: ", pFromDb)

	db.Delete(&p)

	var missing Product
	dbResult := db.First(&missing, "code = ?", "-1")

	if dbResult.RecordNotFound() {
		fmt.Println("Record not found")
	}

	fmt.Println("Should have nothing: ", missing)
}
