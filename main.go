package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Problem struct {
	Number     string
	Title      string
	Difficulty string
	Status     string
	Notes      string
	File       string
}

func formatTitle(title string) string {
	words := strings.Split(title, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, " ")
}

func parseFileAnnotations(path string) (difficulty, status, notes string) {
	file, err := os.Open(path)
	if err != nil {
		return "-", "-", "-"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < 10 && scanner.Scan(); i++ { // 只扫描文件头10行
		line := scanner.Text()
		if strings.HasPrefix(line, "// Difficulty:") {
			difficulty = strings.TrimSpace(strings.TrimPrefix(line, "// Difficulty:"))
		} else if strings.HasPrefix(line, "// Status:") {
			status = strings.TrimSpace(strings.TrimPrefix(line, "// Status:"))
		} else if strings.HasPrefix(line, "// Notes:") {
			notes = strings.TrimSpace(strings.TrimPrefix(line, "// Notes:"))
		}
	}

	if difficulty == "" {
		difficulty = "-"
	}
	if status == "" {
		status = "-"
	}
	if notes == "" {
		notes = "-"
	}

	return
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

			filePath := filepath.Join(dir, file.Name())
			difficulty, status, notes := parseFileAnnotations(filePath)

			problems = append(problems, Problem{
				Number:     number,
				Title:      title,
				Difficulty: difficulty,
				Status:     status,
				Notes:      notes,
				File:       file.Name(),
			})
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	table := "| # | Title | Difficulty | Language | Status | Notes |\n"
	table += "|---|-------|------------|----------|--------|-------|\n"
	for _, p := range problems {
		table += fmt.Sprintf("| %s | %s | %s | Go | %s | %s |\n",
			p.Number, p.Title, p.Difficulty, p.Status, p.Notes)
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
