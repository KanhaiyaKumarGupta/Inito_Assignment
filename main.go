package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
)

func main() {
	fs := commands.NewFileSystem()

	mvCmd := flag.NewFlagSet("mv", flag.ExitOnError)
	cpCmd := flag.NewFlagSet("cp", flag.ExitOnError)
	rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)

	mvSrc := mvCmd.String("src", "", "Source path")
	mvDest := mvCmd.String("dest", "", "Destination path")

	cpSrc := cpCmd.String("src", "", "Source path")
	cpDest := cpCmd.String("dest", "", "Destination path")

	rmPath := rmCmd.String("path", "", "Path to remove")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'mv', 'cp', or 'rm' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "mv":
		mvCmd.Parse(os.Args[2:])
		if *mvSrc == "" || *mvDest == "" {
			fmt.Println("Please provide source and destination paths for mv command")
			os.Exit(1)
		}
		err := fs.Mv(*mvSrc, *mvDest)
		if err != nil {
			fmt.Printf("Error moving file: %s\n", err)
		}

	case "cp":
		cpCmd.Parse(os.Args[2:])
		if *cpSrc == "" || *cpDest == "" {
			fmt.Println("Please provide source and destination paths for cp command")
			os.Exit(1)
		}
		err := fs.Cp(*cpSrc, *cpDest)
		if err != nil {
			fmt.Printf("Error copying file: %s\n", err)
		}

	case "rm":
		rmCmd.Parse(os.Args[2:])
		if *rmPath == "" {
			fmt.Println("Please provide a path for rm command")
			os.Exit(1)
		}
		err := fs.Rm(*rmPath)
		if err != nil {
			fmt.Printf("Error removing file or directory: %s\n", err)
		}

	default:
		fmt.Println("Expected 'mv', 'cp', or 'rm' subcommands")
		os.Exit(1)
	}
}
