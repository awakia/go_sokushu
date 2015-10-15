package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var instance *gorm.DB

// Get gets opened db instance
func Get() *gorm.DB {
	if instance != nil {
		return instance
	}
	db, err := gorm.Open("postgres", "user=awakia dbname=go_sokushu sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	instance = &db
	return instance
}
