package main

import (
  "time"
)

func greedyAiMoves() {
  for !clock.gameover {
    for i:=0; i<9; i++ {
      currentMino.moveLeft()
    }
    maxscore := 0
    var optx, opttheta int
    for i:=0; i<9; i++ {
      // Movements
      for j:=0; j<3; j++ {
        // Rotations
        dstMino := *currentMino
        score := dstMino.putBottom()
        if score>maxscore {
          maxscore = score
          optx = i
          opttheta = j
        }

        currentMino.rotateLeft()
      }
      currentMino.moveRight()
    }

    // Fix position to original
    currentMino.rotateLeft()
    for i:=0; i<9; i++ {
      currentMino.moveLeft()
    }

    for i:=0; i<optx; i++ {
      currentMino.moveRight()
    }

    for i:=0; i<opttheta; i++ {
      currentMino.rotateLeft()
    }

    time.Sleep(1*time.Millisecond)

    currentMino.drop()

    refreshScreen()
  }
}
