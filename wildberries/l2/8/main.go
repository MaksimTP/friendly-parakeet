package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type OS struct {
	currentDir string
}

func (o *OS) jobs(...[]string) {

}
func (o *OS) StartEventLoop() {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)
	var output string
	for text != "\\quit" {
		fmt.Printf("> ")
		scanner.Scan()
		text = scanner.Text()
		cmd := strings.Split(text, " ")
		log.Printf("Scanned text <%s>\n", text)
		log.Println("Splitted text:", cmd)
		switch cmd[0] {
		case "ls":

		case "cd":
			o.ChangeDirectory(cmd[1:]...)
		case "pwd":
			output, _ = o.Pwd()
		case "echo":
			o.Echo(cmd[1:]...)
		case "kill":
			o.Kill(cmd[1:]...)
		case "ps":
			o.Ps(cmd[1:]...)
		default:
		}
		fmt.Println(output)
	}
}

func (o *OS) ChangeDirectory(args ...string) (string, error) {
	dirs, err := os.ReadDir(args[0])
	if err != nil {
		return "", err
	}
	fmt.Println(dirs)
	o.currentDir = args[0]
	return "", nil
}
func (o *OS) Pwd() (string, error) {
	return o.currentDir, nil
}

func (o *OS) Echo(args ...string) (string, error) {
	return "", nil

}
func (o *OS) Kill(args ...string) (string, error) {
	// os.
	return "", nil

}
func (o *OS) Ps(args ...string) (string, error) {
	return "", nil

}

func (o *OS) Ls(args ...string) (string, error) {
	return "", nil

}

func main() {
	dir, _ := os.Getwd()
	OS := &OS{currentDir: dir}
	log.Println("Starting event loop")
	OS.StartEventLoop()
}
