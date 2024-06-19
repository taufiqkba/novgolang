package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Exec on Go

	// Using Exec -> to execute program on CLI
	output1, _ := exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(output1))

	output2, _ := exec.Command("pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(output2))

	output3, _ := exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))

}
