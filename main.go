package main

import (
	"fmt"
	"os"
	"service-tools/service"
)

func printUsageAndQuit() {
	fmt.Println("Usage: st add <service name> <service path>")
	fmt.Println("       st remove <service name>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsageAndQuit()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: st add <service path>")
			os.Exit(1)
			return
		}
		path := os.Args[2]
		name := os.Args[3]
		err := service.Add(path, name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: st remove <service name>")
			os.Exit(1)
			return
		}
		name := os.Args[2]
		err := service.Remove(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		printUsageAndQuit()
	}
}
