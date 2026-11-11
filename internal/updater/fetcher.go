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
	Lookup     map[string]string     `json:"lookup"`
	Weapons    map[string][][]string `json:"weapons"`
	Summons    map[string][][]string `json:"summons"`
	Characters map[string][][]string `json:"characters"`
	Jobs       map[string][][]string `json:"job"`
}

func downloadFileOptionalOverwrite(url, filepath string, overwrite bool) error {
	flags := os.O_CREATE | os.O_WRONLY
	if overwrite {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_EXCL
	}

	out, err := os.Create(filepath)
	if errors.Is(err, os.ErrExist) {
		return nil
	}
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

func downloadFile(url, filepath string) error {
	return downloadFileOptionalOverwrite(url, filepath, true)
}

func downloadFileExcl(url, filepath string) error {
	return downloadFileOptionalOverwrite(url, filepath, false)
}

func decodeData(r io.Reader) (*Data, error) {
	var data Data
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func GetChangelog() (*Changelog, error) {
	res, err := http.Get(changelogURL)
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
