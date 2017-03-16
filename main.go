package toy_robot

import (
	"fmt"
	"flag"
	"os"
)

var (
	fileName string
	readFromConsole bool
)

func main() {
	flag.StringVar(&fileName, "cmdfile", "commands.txt", "Files to get commands")
	flag.BoolVar(&readFromConsole, "console", false, "Read commands from console")
    
    flag.Parse()

    var inputFile *os.File

    if readFromConsole {
    	inputFile = os.Stdin
    } else {
    	f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Command file not found : %s\n", fileName)
			return
		}
		inputFile = f
    }

    run(inputFile)
}
