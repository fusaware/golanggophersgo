/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

func main() {

	// get the template from CLI
	var goTemplate string
	flag.StringVar(&goTemplate, "go-template", "", "The go template.")
	flag.Parse()

	// default if no flag found
	if goTemplate == "" {
		goTemplate = `You said: {{.Input}}`
	}

	d := struct {
		Duration  time.Duration
		EndTime   time.Time
		StartTime time.Time
		Input     string
	}{
		StartTime: time.Now(),
	}

	t, err := template.New("say-something").Parse(goTemplate)
	if err != nil {
		log.Fatalln("doh:", err)
	}

	// prompt the user
	fmt.Println("Say something:")

	// read input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// complete meta data
	d.EndTime = time.Now()
	d.Duration = d.EndTime.Sub(d.StartTime)
	d.Input = input

	if err := t.Execute(os.Stdout, d); err != nil {
		log.Fatalln("template:", err)
	}
}
