package main

import (
	"bytes"
	"context"
	"github.com/NoahOnFyre/gengine/color"
	"github.com/NoahOnFyre/gengine/convert"
	"github.com/NoahOnFyre/gengine/logging"
	"github.com/NoahOnFyre/gengine/utils"
	"github.com/google/go-github/github"
	"golang.org/x/mod/semver"
	"io"
	"os"
	"os/exec"
	"strings"
)

var (
	username, _ = strings.CutPrefix(convert.ValueOf(utils.Catch(os.UserHomeDir())), "C:\\Users\\")
	device, _   = os.Hostname()
	version     = "v1.0.0"
	homeDir, _  = os.UserHomeDir()
	vulcanDir   = homeDir + "\\.vulcan\\"
	moduleDir   = vulcanDir + "module\\"
	tempDir     = vulcanDir + "temp\\"
	configPath  = vulcanDir + "config.json"

	commands []Command
)

func Menu() {
	logging.Print()
	logging.Print(color.Green + "____   ____ __   __  __     _____    ___     ___  __ ")
	logging.Print(color.Green + "\\   \\ /   /|  | |  ||  |   /  __/   /   \\   |   \\|  |")
	logging.Print(color.Green + " \\   V   / |  | |  ||  |   | /     /  ∆  \\  |  \\ |  |")
	logging.Print(color.Green + "  \\     /  |  |_|  ||  |__ | \\__  /  ___  \\ |  |\\   |")
	logging.Print(color.Green + "   \\___/    \\_____/ |_____|\\____\\/__/   \\__\\|__| \\__|")
	logging.Print()
}

func main() {
	logging.SetMainColor(color.GreenBg)

	var newestRelease *github.RepositoryRelease

	go func() {
		release, _, err := githubClient.Repositories.GetLatestRelease(context.Background(), "NoahOnFyre", "vulcan")
		if err != nil {
			return
		}

		if semver.Compare(release.GetTagName(), version) == 1 {
			newestRelease = release
		}
	}()

	CheckPaths([]string{
		homeDir,
		vulcanDir,
		moduleDir,
		tempDir,
	})

	CommandRegistration()
	Menu()

	for {
		SetState("Idle")
		logging.Print()
		// currentDir, _ := os.Getwd()
		input := logging.Input(color.GreenBg + color.Black + " \ue62a " + username + "@" + device + " " + color.Reset + color.Green + "\ue0b8" + " Noah " + color.Reset)
		if input == "" {
			continue
		}
		logging.Print()
		command, args := ParseCommand(input)
		RunCommand(command, args)
		if newestRelease != nil {
			logging.Print()
			logging.Print(color.Gray + "┌" + MultiString("─", 119))
			logging.Print(color.Gray + "│ " + color.Reset + "A new version of Vulcan is available!")
			logging.Print(color.Gray + "│ " + color.Reset + "Version Diff: " + color.Red + version + color.Gray + " -> " + color.Green + newestRelease.GetTagName())
			logging.Print(color.Gray + "└" + MultiString("─", 119))
		}
	}
}

func ParseCommand(input string) (string, []string) {
	split := strings.Split(input, " ")
	command := split[0]
	args := utils.RemoveElement(split, 0)
	return command, args
}

func RunCommand(command string, args []string) {
	commandFound := false
	for _, cmd := range commands {
		if cmd.Name == command {
			if len(args) == cmd.Args.Count {
				SetState("Running: " + cmd.Name)
				cmd.Run(args)
				commandFound = true
			} else {
				s := ""
				for _, argument := range cmd.Args.Get {
					s = s + "<" + argument + "> "
				}
				logging.Error("Invalid arguments!" + color.Gray + " - " + color.Red + "Usage: " + command + " " + s)
				commandFound = true
			}
		}
	}
	if !commandFound {
		_, err := exec.LookPath(command)
		if err != nil {
			logging.Error("Command not found! - Run \"help\" to see all commands.")
			return
		}
		var cmdArgs []string
		cmdArgs = append(cmdArgs, "/c")
		cmdArgs = append(cmdArgs, command)
		cmdArgs = append(cmdArgs, args...)
		SetState("Running: " + command)
		runnable := exec.Command("cmd.exe", cmdArgs...)

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		runnable.Stdout = mw
		runnable.Stderr = mw

		err = runnable.Run()
		if err != nil {
			return
		}
	}
}
