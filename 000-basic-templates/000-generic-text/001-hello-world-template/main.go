package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
)

func ordinal(num int) string {
	var dict = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}
	posNum := int(math.Abs(float64(num)))
	if ((posNum % 100) >= 11) && ((posNum % 100) <= 13) {
		return fmt.Sprintf("%dth", num)
	}
	return fmt.Sprintf("%d%s", num, dict[posNum])
}

func main() {

	bigData := struct { // an
		MaxInt int // unnecessarily
	}{ // complicated
		MaxInt: func() int { // way of making
			return int(^uint(0) >> 1) // a big number
		}(), // for
	} // the template

	templ := `Hello, world for the {{ .MaxInt | ordinal }} time.` // ordinal because erudition
	exhaustion, err := template.New("exhaustion").Funcs(template.FuncMap{"ordinal": ordinal}).Parse(templ)
	if err != nil {
		panic(fmt.Sprintf("as if this wasn't bad enough already: %v", err))
	}
	if err := exhaustion.Execute(os.Stdout, bigData); err != nil {
		log.Fatal(err)
	}
}
