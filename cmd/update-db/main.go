package main

import (
	"fmt"

	"github.com/Quak1/chuy-gbf/internal/updater"
)

func main() {
	err := updater.LoadData(false)
	if err != nil {
		fmt.Println(err)
	}
}
