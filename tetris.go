package main

import (
	"github.com/k0kubun/termbox-go"
	"os"
	"strconv"
	"time"
)

const (
	levelMax         = 20
	scoreMax         = 999999
	gameoverDuration = 10
)

var (
	board       *Board
	clock       *Clock
	currentMino *Mino
	nextMino    *Mino
	score       int
	level       int
	initLevel   int
	deleteLines int
	heightMult  float64
	lineMult		float64
	holeMult		float64
)

func initGame() {
	board = NewBoard()
	initMino()
	score = 0
	level = initLevel
	deleteLines = 0

	refreshScreen()
}

func initMino() {
	currentMino, nextMino = nil, nil
	pushMino()
	pushMino()
}

func deleteCheck() {
	if !board.hasFullLine() {
		return
	}
	clock.pause()

	lines := board.fullLines()
	board.showDeleteAnimation(lines)
	for _, line := range lines {
		board.deleteLine(line)
	}
	deleteLines += len(lines)
	switch len(lines) {
	case 1:
		addScore(40 * (level + 1))
	case 2:
		addScore(100 * (level + 1))
	case 3:
		addScore(300 * (level + 1))
	case 4:
		addScore(1200 * (level + 1))
	}
	levelUpdate()

	clock.start()
}

func levelUpdate() {
	if level == levelMax {
		return
	}

	targetLevel := deleteLines / 10
	if level < targetLevel {
		level = targetLevel
		clock.updateInterval()
	}
}

func addScore(add int) {
	score += add
	if score > scoreMax {
		score = scoreMax
	}
}

func pushMino() {
	deleteCheck()

	currentMino = nextMino
	if currentMino != nil {
		currentMino.x, currentMino.y = defaultMinoX, defaultMinoY
		if currentMino.conflicts() {
			ranking := NewRanking()
			ranking.insertScore(score)
			ranking.save()
			gameOver()
			return
		}
	}
	nextMino = NewMino()
}

func gameOver() {
	clock.over()

	clock.lock = true
	for j := 0; j < boardHeight; j++ {
		rewriteScreen(func() {
			for y := boardHeight - 1; y > boardHeight-1-j; y -= 1 {
				board.colorizeLine(y, termbox.ColorBlack)
			}
		})
		timer := time.NewTimer(gameoverDuration * time.Millisecond)
		<-timer.C
	}
	clock.lock = false
}

func setMultipliers(args []string) {
	if s, err := strconv.ParseFloat(args[0], 64); err == nil {
		heightMult = s
	}

	if s, err := strconv.ParseFloat(args[1], 64); err == nil {
		lineMult = s
	}

	if s, err := strconv.ParseFloat(args[2], 64); err == nil {
		holeMult = s
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	clock = NewClock(func() {
		currentMino.applyGravity()
		refreshScreen()
	})

	initLevel = 1

	initGame()
	clock.start()

	setMultipliers(os.Args[1:])
	aiMoves()
	// storeScore()
}
