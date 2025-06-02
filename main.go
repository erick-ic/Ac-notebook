package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Problem struct {
	Number string
	Title  string
	File   string
}

func formatTitle(title string) string {
	words := strings.Split(title, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, " ")
}

func main() {
	dir := "Go"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var problems []Problem
	r := regexp.MustCompile(`^(\d+)_([\w_]+)\.go$`)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		match := r.FindStringSubmatch(file.Name())
		if len(match) == 3 {
			number := match[1]
			rawTitle := match[2]
			title := formatTitle(rawTitle)
			problems = append(problems, Problem{
				Number: number,
				Title:  title,
				File:   file.Name(),
			})
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	table := "| # | Title | Difficulty | Language | Status | Notes |\n"
	table += "|---|-------|------------|----------|--------|-------|\n"

	for _, p := range problems {
		table += fmt.Sprintf("| %s | %s | - | Go |  |  |\n", p.Number, p.Title)
	}

	updateReadme("README.md", table)
}

func updateReadme(readmePath, table string) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		panic(err)
	}

	startMarker := "<!-- start -->"
	endMarker := "<!-- end -->"

	old := string(content)
	startIdx := strings.Index(old, startMarker)
	endIdx := strings.Index(old, endMarker)

	if startIdx == -1 || endIdx == -1 || startIdx > endIdx {
		fmt.Println("Markers not found or invalid.")
		return
	}

	newContent := old[:startIdx+len(startMarker)] + "\n\n" + table + "\n" + old[endIdx:]
	err = os.WriteFile(readmePath, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("README.md updated!")
}
