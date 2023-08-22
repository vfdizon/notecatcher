# notecatcher

Notecatcher is a Go-Based command line tool for linux that helps you create notes for school, work, etc. 

Prerequisites: 
    - Go

Intructions: 
    1. Build the source code using "go build main.go"
    2. Give the binary run permissions using "chmod +x notecatcher"
    3. Move the binary to /bin 

Usage: 
    ``` notecatcher create-subject -name exampleSubject ```
    ``` notecatcher create-unit -subject exampleSubject -name exampleUnit ```
    ``` notecatcher create-note -subject exampleSubject -unit exampleUnit -name exampleNote ``` 
*create-unit will automatically make the subject if it doesn't exist already, but create-note will only make the note if the subject and the unit exist. 

All notes will be a .md (markdown) file to allow for formatting such as underscores, bold, headers, etc. All notes will default to ~/Documents/notes/ if you want another location, change the config file in ~/.config/notecatcher
