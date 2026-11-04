package updater

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Quak1/chuy-gbf/internal/store"
)

type DataItem struct {
	store.Item
	Rarity string
}

type CategoryLists map[string]map[string][][]string

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
		case 'w':
			item.WikiPath = value
		}
	}

	var img string
	for k, v := range *lists {
		if i, ok := v[item.ID]; ok {
			item.Category, _ = strings.CutSuffix(k, "s")

			switch k {
			case "jobs":
				item.Element = "any"
				item.Rarity = "ssr"
				item.Category = "character"
				img = last(i[2])
			case "weapons":
				img = last(i[0])
			case "summons":
				img = last(i[0])
			case "characters":
				img = last(i[6])
				if n := findJumpIndex(i[6]); n != -1 {
					img = i[6][n-1]
				}
			}

			if img == "" {
				img = item.ID
				if k == "character" {
					img = fmt.Sprintf("%s_01", item.ID)
				}
			}
			item.ImageName = img

			break
		}
	}

	return item
}

func last(a []string) string {
	if len(a) == 0 {
		return ""
	}
	return a[len(a)-1]
}

func findJumpIndex(items []string) int {
	if len(items) < 2 {
		return -1
	}

	prevNum, err := parseSuffix(items[0])
	if err != nil {
		return -1
	}

	for i := 1; i < len(items); i++ {
		currNum, err := parseSuffix(items[i])
		if err != nil {
			return -1
		}

		if currNum-prevNum != 1 {
			return i
		}
		prevNum = currNum
	}

	return -1
}

func parseSuffix(s string) (int, error) {
	parts := strings.Split(s, "_")
	if len(parts) < 2 {
		return 0, fmt.Errorf("no underscore delimiter found in %s", s)
	}
	return strconv.Atoi(parts[1])
}

func parseData(data *Data) []DataItem {
	items := []DataItem{}
	lists := &CategoryLists{
		"jobs":       data.Jobs,
		"weapons":    data.Weapons,
		"summons":    data.Summons,
		"characters": data.Characters,
	}

	for id, l := range data.Lookup {
		items = append(items, parseLookup(id, l, lists))
	}

	return items
}
