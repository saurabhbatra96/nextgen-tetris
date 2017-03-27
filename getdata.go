package main

import (
  "os"
  "strconv"
)

func storeScore() {
  f, err := os.OpenFile("/home/saurabh/projects/tetris-data/aionlylinesdata", os.O_RDWR|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scoreString := strconv.Itoa(score)

  if _, err = f.WriteString(scoreString+"\n"); err != nil {
    panic(err)
  }
}
