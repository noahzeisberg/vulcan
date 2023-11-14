package main

import (
	"github.com/NoahOnFyre/gengine/color"
	"github.com/NoahOnFyre/gengine/logging"
	"github.com/google/go-github/github"
	"os"
	"strings"
)

var (
	githubClient = github.NewClient(nil)
)

func VulcanCommand(args []string) {
}

func CdCommand(args []string) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		logging.Error(err)
		return
	}
}

func LsCommand(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		logging.Error(err)
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		logging.Error(err)
	}
	for _, file := range files {
		if file.IsDir() {
			logging.Log(color.Green + "/" + file.Name())
		} else if strings.HasPrefix(file.Name(), ".") {
			logging.Log(color.Gray + file.Name())
		} else {
			logging.Log(file.Name())
		}
	}
}

func HelpCommand(args []string) {
	for _, command := range commands {
		s := ""
		for _, argument := range command.Args.Get {
			s += "<" + argument + "> "
		}
		logging.Log(color.Green + strings.ToUpper(command.Name) + color.Gray + ": " + color.Reset + command.Description + color.Gray + " - " + color.Green + s)
	}
}

func ClearCommand(args []string) {
	logging.Clear()
	Menu()
}

func ExitCommand(args []string) {
	logging.Log("Shutting down FyUTILS...")
	os.Exit(0)
}
