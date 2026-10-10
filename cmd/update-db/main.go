package main

import (
	"fmt"
	"log"

	"github.com/Quak1/chuy-gbf/internal/database"
	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/Quak1/chuy-gbf/internal/updater"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	defer db.Close()

	q := store.New(db)

	err, msgs := updater.LoadData(false, q)
	for msg := range msgs {
		fmt.Println(msg)
	}
	if err != nil {
		fmt.Println(err)
	}
}
