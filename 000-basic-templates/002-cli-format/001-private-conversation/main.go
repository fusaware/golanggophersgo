package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

func main() {
	type data struct {
		Duration  time.Duration
		EndTime   time.Time
		StartTime time.Time
		Input     string
	}
	d0 := data{
		StartTime: time.Now(),
	}

	t, err := template.ParseFiles("meta-data.gohtml")
	if err != nil {
		panic(fmt.Sprintf("doh: %v", err))
	}

	t = t.Funcs(template.FuncMap{
		"trim": func(v string) string {
			return strings.Trim(v, "\n")
		},
	})

	// prompt the user
	fmt.Println("Say something:")

	// read input
	reader := bufio.NewReader(os.Stdin)
	input0, _ := reader.ReadString('\n')

	// complete meta data
	d0.EndTime = time.Now()
	d0.Duration = d0.EndTime.Sub(d0.StartTime)
	d0.Input = input0

	d1 := data{
		StartTime: time.Now(),
	}

	// prompt the user
	fmt.Println("Now respond:")
	// read input
	input1, _ := reader.ReadString('\n')
	// complete meta data
	d1.EndTime = time.Now()
	d1.Duration = d0.EndTime.Sub(d0.StartTime)
	d1.Input = input1

	ds := []data{d0, d1}

	if err := t.Execute(os.Stdout, ds); err != nil {
		log.Fatalln("template:", err)
	}

}
