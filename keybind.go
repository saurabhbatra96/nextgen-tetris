package main

import (
	"os"
	"bufio"
	"time"
)

func waitKeyInput() {
	file, err := os.Open("/home/saurabh/go-work/src/github.com/saurabhbatra96/nextgen-tetris/output.txt")
	if (err != nil) {
		panic(err)
	}
  defer file.Close()

	// logFile, logErr := os.Open("log.txt")


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "rotleft" {
			currentMino.rotateLeft()
		} else if scanner.Text() == "rotright" {
			currentMino.rotateRight()
		} else if scanner.Text() == "movleft" {
			currentMino.moveLeft()
		} else if scanner.Text() == "movright" {
			currentMino.moveRight()
		} else if scanner.Text() == "drop" {
			currentMino.drop()
		}

		refreshScreen()
	}
}
