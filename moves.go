package main

import (
  "time"
  "math/rand"
)

func movesList() {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for !clock.gameover {
    leftRots := r1.Intn(3) + 1
    lOrR := r1.Intn(2)
    moves := r1.Intn(10) + 1

    for i:=0; i<leftRots; i++ {
      currentMino.rotateLeft()
    }

    for i:=0; i<moves; i++ {
      if lOrR == 1 {
        currentMino.moveLeft()
      } else {
        currentMino.moveRight()
      }
    }
    currentMino.drop()

    refreshScreen()
  }
}
