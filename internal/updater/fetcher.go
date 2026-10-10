package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	dataURL      = "https://raw.githubusercontent.com/MizaGBF/GBFAL/main/json/data.json"
	changelogURL = "https://raw.githubusercontent.com/MizaGBF/GBFAL/main/json/changelog.json"
	dataPath     = "data.json"
)

type Changelog struct {
	Timestamp int64              `json:"timestamp"`
	New       map[string][][]any `json:"new"`
	Stat      string             `json:"stat"`
	Issues    []any              `json:"issues"`
	Help      bool               `json:"help"`
}

type Data struct {
	Lookup     map[string]string `json:"lookup"`
	Weapons    map[string]any    `json:"weapons"`
	Summons    map[string]any    `json:"summons"`
	Characters map[string]any    `json:"characters"`
}

func downloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func decodeData(r io.Reader) (*Data, error) {
	var data Data
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func GetChangelog() (*Changelog, error) {
	res, err := http.Get(dataURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var c Changelog
	if err := json.NewDecoder(res.Body).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func fetchData() (*Data, error) {
	res, err := http.Get(dataURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return decodeData(res.Body)
}

func loadFile() (*Data, error) {
	if _, err := os.Stat(dataPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Downloading data...")
		if err = downloadFile(dataURL, dataPath); err != nil {
			return nil, err
		}
		fmt.Printf("%s downloaded\n", dataPath)
	}

	f, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return decodeData(f)
}
