package updater

import (
	"strings"
)

type Item struct {
	ID      string
	Element string
	Type    string
	Rarity  string
	Name    string
	Series  string
}

func parseLookup(id, input string) Item {
	item := Item{
		ID: id,
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

	return item
}

func parseData(data *Data) []Item {
	items := []Item{}

	for id, l := range data.Lookup {
		items = append(items, parseLookup(id, l))
	}

	return items
}
