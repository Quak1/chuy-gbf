package updater

import (
	"fmt"
)

func LoadData(download bool) error {
	data, err := loadFile()
	if err != nil {
		return err
	}

	items := parseData(data)
	for i := range 10 {
		fmt.Printf("%+v\n", items[i])
	}

	return nil
}
