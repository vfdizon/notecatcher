package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	NotesDirectory string
}

func main() {
	createSubjectCmd := flag.NewFlagSet("create-subject", flag.ExitOnError)
	subjectName := createSubjectCmd.String("name", "", "Subject name")

	createUnitCmd := flag.NewFlagSet("create-unit", flag.ExitOnError)
	subjectNameUnit := createUnitCmd.String("subject", "", "Subject name")
	unitName := createUnitCmd.String("name", "", "Unit name")

	createNoteCmd := flag.NewFlagSet("create-note", flag.ExitOnError)
	subjectNameNote := createNoteCmd.String("subject", "", "Subject name")
	unitNameNote := createNoteCmd.String("unit", "", "Unit name")
	noteName := createNoteCmd.String("name", "", "Note name")

	switch os.Args[1] {
	case "create-subject":
		createSubjectCmd.Parse(os.Args[2:])
		CreateSubject(*subjectName, readConfig())
	case "create-unit":
		createUnitCmd.Parse(os.Args[2:])
		CreateUnit(*subjectNameUnit, *unitName, readConfig())
	case "create-note":
		createNoteCmd.Parse(os.Args[2:])
		CreateNote(*noteName, *unitNameNote, *subjectNameNote, readConfig())

	default:
		fmt.Println("Invalid command use: create-subject, create-unit, create-note")
		os.Exit(1)
	}
}

func readConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(homeDir, ".config", "notecatcher")

	os.Mkdir(configPath, 0755)

	files, err := os.ReadDir(configPath)
	if err != nil {
		panic(err)
	}

	var containsConfig bool
	for _, file := range files {
		if file.Name() == "config" {
			containsConfig = true
			break
		}
	}

	if !containsConfig {
		configFile, err := os.Create(filepath.Join(configPath, "config"))
		if err != nil {
			panic(err)
		}

		configWriter := bufio.NewWriter(configFile)
		configWriter.WriteString("notes_directory = " + filepath.Join(homeDir, "/Documents/notes"))
		configWriter.Flush()
		fmt.Println("Created Config file at " + filepath.Join(configPath, "config"))
		return Config{
			NotesDirectory: filepath.Join(homeDir, "/Documents/notes"),
		}
	}

	configFile, err := os.Open(filepath.Join(configPath, "config"))
	if err != nil {
		panic(err)
	}

	configReader := bufio.NewScanner(configFile)
	config := Config{}

	for configReader.Scan() {
		line := configReader.Text()
		if strings.Contains(line, "notes_directory") {
			config.NotesDirectory = strings.TrimSpace(strings.Split(line, "=")[1])
		}
	}

	return config
}

func CreateSubject(subjectName string, config Config) {
	subjectPath := filepath.Join(config.NotesDirectory, subjectName)
	os.Mkdir(subjectPath, 0755)
	fmt.Println(subjectName + " created at " + subjectPath)
}

func CreateUnit(subject, unit string, config Config) {
	noteDirectory := config.NotesDirectory

	files, err := os.ReadDir(noteDirectory)
	if err != nil {
		panic(err)
	}

	var containsSubject bool = false
	for _, file := range files {
		if file.Name() == subject && file.IsDir() {
			containsSubject = true
			break
		}
	}

	if !containsSubject {
		fmt.Println("Subject doesn't exist, creating subject...")
		CreateSubject(subject, config)
	}

	subjectPath := filepath.Join(noteDirectory, subject)
	unitPath := filepath.Join(subjectPath, unit)
	os.Mkdir(unitPath, 0755)
	fmt.Println(unit + " created at " + unitPath)
}

func CreateNote(note, unit, subject string, config Config) {
	notePath := filepath.Join(config.NotesDirectory, subject, unit, note)
	os.Create(notePath + ".md")
	fmt.Println(note + " created at " + notePath)
}
