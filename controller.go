package main


import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func run(fp *os.File) {

	if fp == nil {
		fmt.Printf("Command file not defined\n")
		return
	}

	myRobot := newRobot()

	bufScanner := bufio.NewScanner(fp)
    for bufScanner.Scan() {
        line := bufScanner.Text()
        if len(line) != 0 {
            line = strings.TrimSpace(line)
        	
        	executeCommad(line, &myRobot)
        }
    }
    if err := bufScanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error in reading stdin :", err)
    }
}

func executeCommad(cmdLine string, r *robot) {
	
	// tokenize commnad, make all uppercase and put into a vector
	tokens := strings.Split(cmdLine, " ")
	if len(tokens) == 0 && len(tokens) > 2 {
		fmt.Printf("Invalid command : [%s]\n", cmdLine)
		return
	}

	command := strings.ToUpper(tokens[0])
	
	switch(command) {
		case "PLACE":
			params := strings.Split(tokens[1], ",")
			if len(params) != 3 {
				fmt.Printf("Invalid command : [%s]\n", cmdLine)
				return
			}

			x, err := strconv.Atoi(params[0])
			if err != nil {
				fmt.Printf("Invalid command : [%s] \n", cmdLine)
				return
			}

			y, err := strconv.Atoi(params[1])
			if err != nil {
				fmt.Printf("Invalid command : [%s] \n", cmdLine)
				return
			}

			r.place(x, y, params[2])

		case "MOVE":
			r.move()
		case "LEFT":
			r.left()
		case "RIGHT":
			r.right()
		case "REPORT":
			r.report()
		default:
			fmt.Printf("Invalid command : [%s]\n", cmdLine)
	}
}