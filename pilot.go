package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){
	phpDir := "/var/www/html"
	goDir := "Desktop/go_projects"
	laravelDir := "Desktop/laravel_projects"
	fmt.Println("Enter Language Or Frmawork: \n1.Php\n2.Go\n3.Laravel")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ",err)
		return
	}
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		err := os.Chdir(phpDir)
		if err != nil {
			fmt.Println("Error changing directory: ",err)
			return
		}
	case "2":
		err := os.Chdir(goDir)
		if err != nil {
			fmt.Println("Error changing directory: ",err)
			return
		}
	case "3":
		err := os.Chdir(laravelDir)
		if err != nil {
			fmt.Println("Error cahnging directory: ",err)
			return
		}
	default: 	
	fmt.Println("Invalid input, Exiting...")
	os.Exit(1)	
	}
	fmt.Println("Choose action: \n1.Create\n2.Open\n3.Delete")
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ",err)
		return
	}
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		fmt.Println("Enter project name:")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ",err)
			return
		}
		input = strings.TrimSpace(input)
		cmd := exec.Command("sudo", "mkdir",input)
	    fmt.Printf("Running command: %s\n", cmd.String())
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error while creating your peoject: ",err)
			return
		}
		fmt.Println("Your project was created successfully")
	case "2":
		cmd := exec.Command("ls")
	    fmt.Printf("Running command: %s\n", cmd.String())	
	    out, err := cmd.Output()
	    if err != nil {
			println("Could not run the command: ",err)
			return
		}
		fmt.Println(string(out))
		fmt.Println("Enter project name:")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ",err)
			return
		}
		input = strings.TrimSpace(input)
		cmd = exec.Command("code", input)
	    fmt.Printf("Running command: %s\n", cmd.String())
		fmt.Println("Opneing your project...")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error while opening your peoject: ",err)
			return
		}
		fmt.Println("exiting ...")
	case "3":
		cmd := exec.Command("ls")
	    fmt.Printf("Running command: %s\n", cmd.String())	
	    out, err := cmd.Output()
	    if err != nil {
			println("Could not run the command: ",err)
			return
		}
		fmt.Println(string(out))
		fmt.Println("Enter project name:")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ",err)
			return
		}
		input = strings.TrimSpace(input)
		cmd = exec.Command("sudo","rm","-r", input)
	    fmt.Printf("Running command: %s\n", cmd.String())
		fmt.Println("Deleting your project...")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error while opening your peoject: ",err)
			return
		}
		fmt.Println("exiting ...")
	default: 
	fmt.Println("Invalid input, Exiting...")
	os.Exit(1)	
}
	
}