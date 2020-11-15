package main

import (
	"fmt"

	getch "github.com/blackestwhite/crawler/getch"
)

func main() {
	theDoc := `<html><body>
	<form name="query" action="http://www.example.net/action.php" method="post">
		<textarea type="text" name="nameiknow">The text I want</textarea>
		<div id="button">
			<input type="submit" value="Submit" />
			<a href="https://google.com">google</a>
		</div>
	</form>
	</body></html>`
	res, err := getch.Get("textarea", theDoc)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range res {
		fmt.Println(val)
	}

	links, err := getch.GetLinks(theDoc)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range links {
		fmt.Println(val)
	}
}
