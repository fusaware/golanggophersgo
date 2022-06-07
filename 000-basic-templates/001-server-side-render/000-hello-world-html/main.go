package main

import (
	"fmt"
	"html/template"
	"os"
)

const helloWorld = `
<!DOCTYPE html>
<html>
<body>

<h1>Hello World</h1>
<p>Again.</p>

</body>
</html>
`

func main() {
	// use template.Must because to not waste our lives wondering what might have been if we made better choices
	t := template.Must(template.New("hello-world.gohtml").Parse(helloWorld))

	if err := t.Execute(os.Stdout, struct{}{}); err != nil {
		panic(fmt.Sprintf("as if this wasn't bad enough already: %v", err))
	}
}
