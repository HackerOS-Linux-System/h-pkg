package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: h-pkg-cli <command> [args...]")
		fmt.Println("Commands: install, remove, update, upgrade, search")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	var aptCmd string
	var aptArgs []string

	switch command {
		case "install":
			aptCmd = "apt"
			aptArgs = append([]string{"install", "-y"}, args...)
		case "remove":
			aptCmd = "apt"
			aptArgs = append([]string{"remove", "-y"}, args...)
		case "update":
			aptCmd = "apt"
			aptArgs = []string{"update"}
		case "upgrade":
			aptCmd = "apt"
			aptArgs = []string{"upgrade", "-y"}
		case "search":
			aptCmd = "apt"
			aptArgs = append([]string{"search"}, args...)
		default:
			fmt.Printf("Unknown command: %s\n", command)
			os.Exit(1)
	}

	cmd := exec.Command("sudo", append([]string{aptCmd}, aptArgs...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
