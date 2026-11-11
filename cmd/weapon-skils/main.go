package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/Quak1/chuy-gbf/internal/database"
	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/Quak1/chuy-gbf/internal/updater"
)

type WeaponSkill struct {
	Name        string
	Description string
	Img         string
}

func main() {
	concurrentLimit := flag.Int("c", 12, "Concurrent weapon parsing limit")
	dir := flag.String("dir", "weapon-skills", "Directory to save weapon skill images")

	db, err := database.InitDB()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	defer db.Close()

	q := store.New(db)

	items, err := q.GetAllItems(context.Background())
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	var wg sync.WaitGroup
	var count atomic.Int32
	var failed atomic.Int32
	sem := make(chan struct{}, *concurrentLimit)

	client := &http.Client{}

	log.Printf("Parsing %d items, for weapon skills\n", len(items))

	for _, item := range items {
		if item.Category != "weapon" {
			continue
		}

		sem <- struct{}{}
		wg.Go(func() {
			defer func() { <-sem }()

			wikiUrl := "https://gbf.wiki/" + item.WikiPath
			_, err := updater.ParseWeaponSkills(client, *dir, wikiUrl)
			if err != nil {
				failed.Add(1)
				log.Println(err)
			}

			count.Add(1)
		})
	}

	wg.Wait()
	log.Printf("Done with %d weapons, and %d error(s)\n", count.Load(), failed.Load())
}
