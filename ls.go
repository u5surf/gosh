package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"strings"
)

func ls(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	blue := color.FgCyan.Render
	yellow := color.FgYellow.Render
	fmt.Println("┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━┯")
	for i, file := range files {
		x := 3 - len(string(i))
		z := 33 - len(file.Name())
		spaces := strings.Repeat(" ", z)
		spaces2 := strings.Repeat(" ", x)
		if file.IsDir() {
			if i >= 10 {
				fmt.Printf("│%d%s│ %s%s│ Directory       │\n", i, spaces2, blue(file.Name()), spaces)
			} else {
				fmt.Printf("│%d%s │ %s%s│ Directory       │\n", i, spaces2, blue(file.Name()), spaces)
			}
		} else if file.Mode().String() == "-rwxr-xr-x" {
			if i >= 10 {
				fmt.Printf("│%d%s│ %s%s│ Executable File │\n", i, spaces2, yellow(file.Name()), spaces)
			} else {
				fmt.Printf("│%d%s │ %s%s│ Executable File │\n", i, spaces2, yellow(file.Name()), spaces)
			}
		} else {
			if i >= 10 {
				fmt.Printf("│%d%s│ %s%s│ File            │\n", i, spaces2, file.Name(), spaces)
			} else {
				fmt.Printf("│%d%s │ %s%s│ File            │\n", i, spaces2, file.Name(), spaces)
			}
		}
	}
	fmt.Println("┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷")
}
