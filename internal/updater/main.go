package updater

import (
	"context"

	"github.com/Quak1/chuy-gbf/internal/database"
	"github.com/Quak1/chuy-gbf/internal/store"
)

func LoadData(download bool) error {
	data, err := loadFile()
	if err != nil {
		return err
	}

	db, err := database.InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	items := parseData(data)
	q := store.New(db)

	for _, item := range items {
		if item.Rarity == "ssr" {
			q.CreateItem(context.Background(), store.CreateItemParams(item.Item))
		}
	}

	return nil
}
