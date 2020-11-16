package getch

import (
	"errors"
	"fmt"
	"regexp"
)

func Get(element string, doc string) (map[int]string, error) {
	pattern := fmt.Sprintf(`<%s.*?>(.*)</%s>`, element, element)
	re := regexp.MustCompile(pattern)
	submatchAll := re.FindAllStringSubmatch(doc, -1)
	matches := make(map[int]string)
	for i, val := range submatchAll {
		matches[i] = val[1]
	}
	if len(matches) > 0 {
		return matches, nil
	} else {
		return nil, errors.New("No Matches Find!")
	}
}

func GetLinks(doc string) (map[int]string, error) {
	pattern := `href[ ]{0,1}=[ ]{0,1}"([^\"]{0,})"`
	re := regexp.MustCompile(pattern)
	submatchAll := re.FindAllStringSubmatch(doc, -1)
	matches := make(map[int]string)
	for i, val := range submatchAll {
		matches[i] = val[1]
	}
	if matches != nil {
		return matches, nil
	} else {
		return nil, errors.New("No Links Find!")
	}
}
