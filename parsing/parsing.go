package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Tags struct {
	note string
	tag  string
}

const tagPrefix = "#"

var wikiDir = os.Getenv("VIMWIKI")

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
		taglines, err := getTaglines(path)
		result = append(result, taglines...)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	log.Println("Get All Tagline is Done")
	return result, nil
}

func getTaglines(filename string) ([]string, error) {
	var result []string
	val, err := ioutil.ReadFile(filename)
	if err != nil {
		return result, err
	}

	ss := strings.Split(string(val), "\n\n"+tagPrefix)
	for _, s := range ss {
		if strings.HasPrefix(s, tagPrefix) {
			result = append(result, s)
		}
	}
	log.Println("Get Taglines from filename", filename)
	return result, nil
}

func getNoteDir() error {
	if wikiDir == "" {
		wikiDir = "/home/sh/vimwiki"
	}
	return nil
}

func makeCSVForm(tags []string) ([][]string, error) {
	length := len(tags)
	result := make([][]string, length)
	//current := time.Now().String()
	// if i saw the tag.
	// change current
	for i := range result {
		tag := strings.ReplaceAll(tags[i], "\n", "==")
		//tag = "\"" + tag + "\""
		result[i] = []string{strconv.Itoa(i), tag}
	}
	log.Println("Successfully make csv form")
	return result, nil
}

func toCSV(tags []string) error {
	path := os.Getenv("CSV_PATH")
	if path == "" {
		path = "../tags.csv"
	}
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
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
	log.Println("Successfully convert to csv")
	return nil
}
