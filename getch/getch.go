package getch

import (
	"errors"
	"fmt"
	"regexp"
)

func Get(element string, doc string) (map[int]string, error) {
	expression := fmt.Sprintf(`<%s.*?>(.*)</%s>`, element, element)
	re := regexp.MustCompile(expression)
	submatchAll := re.FindAllStringSubmatch(doc, -1)
	matches := make(map[int]string)
	for i, val := range submatchAll {
		matches[i] = val[1]
	}
	if matches != nil {
		return matches, nil
	} else {
		return nil, errors.New("No Matches Find!")
	}
}