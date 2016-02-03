package cli

import (
	"bufio"
	"fmt"
	"os"
)

var (
	yesNoDisabled = false
)

func Run(fn func()) {
	ParseArgs()
	fn()
}

func RunFromGitRoot(fn func()) {
	err := os.Chdir(GetGitRoot())
	if err != nil {
		panic(err)
	}
	Run(fn)
}

func DisableYesNo() {
	yesNoDisabled = true
}

func YesNo(argv ...interface{}) bool {
	fmt.Println(fmt.Sprint(argv) + "? (y/n/q)")
	if yesNoDisabled {
		fmt.Println("Yes/No disabled - continue")
		return true
	}
	for {
		line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		switch string(line) {
		case "y": // proceed
			return true
		case "n": // exit
			return false
		case "q": // quit
			os.Exit(0)
		default: // repeat
			fmt.Println("(y/n)")
		}
	}
	return false
}
