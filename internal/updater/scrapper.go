package updater

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WeaponSkill struct {
	Name string
	url  string
	Img  string
}

func ParseWeaponSkills(client *http.Client, dir, wikiUrl string) ([]*WeaponSkill, error) {
	req, err := http.NewRequest("GET", wikiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Chuy GBF")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status code error: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	var skills []*WeaponSkill

	sel := doc.Find(".skill:not(.skill-upgrade-next)")
	for i := range sel.Nodes {
		node := sel.Eq(i)

		url, _ := node.Find(".skill-icon img").First().Attr("src")
		name := strings.TrimSpace(node.Find(".skill-name").First().Text())
		w := &WeaponSkill{Name: name, url: url}

		imgPath, err := downloadSkillImg(dir, w)
		if err != nil {
			return nil, err
		}

		w.Img = filepath.Base(imgPath)
		skills = append(skills, w)
	}

	return skills, nil
}

func downloadSkillImg(dir string, skill *WeaponSkill) (string, error) {
	fpath := filepath.Join(dir, filepath.Base(skill.url))

	err := downloadFileExcl(skill.url, fpath)
	if err != nil {
		return "", nil
	}

	return fpath, nil
}
