package main

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	//"time"
)

type Tags struct {
	note string
	tag  string
}

var tagPrefix = "##"
var wikiDir = "/home/sh/vimwiki"

func getTagAll() ([]string, error) {
	result := []string{}
	err := filepath.Walk(wikiDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".md" {
			return nil
		}
		filename := path
		val, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer val.Close()

		scanner := bufio.NewScanner(val)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, tagPrefix) {
				result = append(result, line)
			}
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	log.Println("Get Tag All Done")
	return result, nil
}

func getTag(tagline string) string {
	tag := strings.Split(tagline, "\n")
	if len(tag) == 0 {
		return "error"
	}
	log.Println("Get Tag in Tagline", tag[0])

	return tag[0]
}

func getTaglineAll() ([]string, error) {
	result := []string{}
	err := filepath.Walk(wikiDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".md" {
			return nil
		}
		filename := path
		val, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		ss := strings.Split(string(val), "\n\n")
		for _, s := range ss {
			if strings.HasPrefix(s, tagPrefix) {
				result = append(result, s)
			}
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	log.Println("Get Tagline All Done")
	return result, nil
}

func getRandom(arg interface{}) []int {
	var value int
	switch arg.(type) {
	case string:
		value = len(arg.(string))
		log.Println("check")
	case int:
		value = arg.(int)
	default:
		value = 10
	}

	var numbers []int
	for i := 0; i < 5; i++ {
		numbers = append(numbers, rand.Intn(value+i))
	}
	log.Println(numbers)
	return numbers
}

func makeCSVForm(tags []string) ([][]string, error) {
	length := len(tags)
	result := make([][]string, length)
	//current := time.Now().String()
	// if i saw the tag.
	// change current
	for i := range result {
		tag := strings.ReplaceAll(tags[i], "\n", " ")
		tag = "\"" + tag + "\""
		result[i] = []string{strconv.Itoa(i), tag}
	}
	return result, nil
}

func toCSV(tags []string) error {
	file, err := os.OpenFile("../tags.csv", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	// TagName Latest-Update Latest-Read Weight
	// How to check update date?
	var contents [][]string
	contents, err = makeCSVForm(tags)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = writer.Write([]string{"id", "description"})
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = writer.WriteAll(contents)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func textToString(filename string) ([]string, error) {
	val, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Fatal Read File for text To string job")
	}

	ss := strings.Split(string(val), "\n")
	var result []string
	for _, s := range ss {
		result = append(result, s)
	}
	return result, nil
}

func toHTML(contents []string) error {
	file, err := os.OpenFile("index.html", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	htmlForm := `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width">
        <title>Ta-da</title>
    </head>
    <body>`
	file.WriteString(htmlForm)
	for _, content := range contents {
		content = "<p>" + content + "</p>"
		file.WriteString(content)
	}

	htmlTail := `
    </body>
</html>
	`
	file.WriteString(htmlTail)

	return nil
}
