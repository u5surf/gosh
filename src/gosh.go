package main
import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "os/exec"
    "log"
)
func executeCommand(theCommand string) {
	cmd := exec.Command(theCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}
func getDir(commandString string) string {
	s:=strings.Split(commandString," ")
	if len(s) > 1 {
		return s[1]
	}
	return "error"
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to gosh the Go Shell!")
    fmt.Println("-----------------------------")
    for {
        executeCommand("/workspace/gosh/src/commands/bin/prompt")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        if strings.Compare("help",command) == 0 {
			executeCommand("/workspace/gosh/src/commands/bin/help")
        } else if strings.Compare("exit",command) == 0 {
            break;
		} else if(strings.Compare("ls", command)==0) {
			executeCommand("/workspace/gosh/src/commands/list.py")
		} else if(strings.Compare("",command) == 0) {
			continue;
		} else if strings.HasPrefix(command,"cd") {
			dir := getDir(command)
			if dir == "error" {
				println("gosh: cd: file not specified")
			}
			os.Chdir(dir)
		} else {
			cmd := exec.Command(command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                fmt.Println("gosh: " + command + ": command not found")
            }
		}
	}
}
