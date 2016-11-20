package main

import "fmt"

var (
	//Version ...
	Version = "EMPTY"
	//BuildTime ...
	BuildTime = "EMPTY"
	//CommitHash ...
	CommitHash = "EMPTY"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(fmt.Sprintf("Version %s", Version))
	fmt.Println(fmt.Sprintf("Build Time %s", BuildTime))
	fmt.Println(fmt.Sprintf("Commit Hash %s", CommitHash))
}
