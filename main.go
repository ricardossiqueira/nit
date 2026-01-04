/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"nit/cmd"
	"nit/internal/db"
)

func main() {
	sqliteDB, err := db.Open("nit.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.InitSchema(sqliteDB); err != nil {
		log.Fatal(err)
	}

	store, err := db.NewStore(sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	cmd.SetRunStore(store)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
