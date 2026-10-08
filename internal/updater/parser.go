package updater

import (
	"strings"

	"github.com/Quak1/chuy-gbf/internal/store"
)

type DataItem struct {
	store.Item
	Rarity string
}

type CategoryLists map[string]map[string]any

func parseLookup(id, input string, lists *CategoryLists) DataItem {
	item := DataItem{
		Item: store.Item{ID: id, Enabled: false},
	}

	if strings.HasPrefix(input, "/") {
		input = " " + input
	}

	parts := strings.SplitSeq(input, " /")
	for part := range parts {
		if part == "" {
			continue
		}

		key := part[0]
		value := strings.TrimSpace(part[1:])

		switch key {
		case 'e':
			item.Element = value
		case 't':
			item.Type = value
		case 'a':
			item.Rarity = value
		case 'n':
			item.Name = value
		case 'c':
			item.Series = value
		}
	}

	for k, v := range *lists {
		if _, ok := v[item.ID]; ok {
			item.Category, _ = strings.CutSuffix(k, "s")
			break
		}
	}

	return item
}

func parseData(data *Data) []DataItem {
	items := []DataItem{}
	lists := &CategoryLists{
		"weapons":    data.Weapons,
		"summons":    data.Summons,
		"characters": data.Characters,
	}

	for id, l := range data.Lookup {
		items = append(items, parseLookup(id, l, lists))
	}

	return items
}
