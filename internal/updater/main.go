package updater

import (
	"context"
	"fmt"

	"github.com/Quak1/chuy-gbf/internal/store"
)

func LoadData(download bool, q *store.Queries) (error, []string) {
	var data *Data
	var err error
	if download {
		data, err = fetchData()
	} else {
		data, err = loadFile()
	}
	if err != nil {
		return err, nil
	}

	items := parseData(data)

	var errorMsg []string
	for _, item := range items {
		if item.Rarity != "ssr" {
			continue
		}

		err := q.CreateItem(context.Background(), store.CreateItemParams(item.Item))
		if err != nil {
			errorMsg = append(errorMsg, fmt.Sprintf("%+v\n%v\n", item, err))
		}
	}

	return nil, errorMsg
}
