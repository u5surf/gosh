package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
)

func clearHistory() string {
	f, _ := os.OpenFile("/.gosh/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Truncate(0)
	fmt.Printf("%s ✔\n", color.FgYellow.Render("History has been cleared"))
	return "History has been cleared ✔\n"
}
func history() {
	file, _ := os.Open("/.gosh/history.txt")
	scanner := bufio.NewScanner(file)
	var num = 1
	fmt.Printf("   %s      %s\n", color.FgGreen.Render("#"), color.FgGreen.Render("command"))
	fmt.Println(" ╭━━━━━━━━━━━━━━━━━━━╮")
	for scanner.Scan() {
		z := 14 - len(scanner.Text())
		spaces := strings.Repeat(" ", z)
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		if num < 10 {
			fmt.Printf(" │ %s  │ %s%s│\n", color.FgGreen.Render(num), scanner.Text(), spaces)
		} else {
			fmt.Printf(" │ %s │ %s%s│\n", color.FgGreen.Render(num), scanner.Text(), spaces)
		}
		num++
	}
	fmt.Println(" ╰━━━━━━━━━━━━━━━━━━━╯")
}

func updateHistory(command string) {
	f, err := os.OpenFile("/.gosh/history.txt"
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + command); err != nil {
		log.Println(err)
	}
}
