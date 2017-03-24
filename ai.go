package main

import (
  "time"
  "os"
  "strconv"
)

// func greedyAiMoves() {
//   for !clock.gameover {
//     for i:=0; i<9; i++ {
//       currentMino.moveLeft()
//     }
//     maxscore := 0
//     var optx, opttheta int
//     for i:=0; i<9; i++ {
//       // Movements
//       for j:=0; j<3; j++ {
//         // Rotations
//         dstMino := *currentMino
//         score := dstMino.putBottom()
//         if score>maxscore {
//           maxscore = score
//           optx = i
//           opttheta = j
//         }
//
//         currentMino.rotateLeft()
//       }
//       currentMino.moveRight()
//     }
//
//     // Fix position to original
//     currentMino.rotateLeft()
//     for i:=0; i<9; i++ {
//       currentMino.moveLeft()
//     }
//
//     for i:=0; i<optx; i++ {
//       currentMino.moveRight()
//     }
//
//     for i:=0; i<opttheta; i++ {
//       currentMino.rotateLeft()
//     }
//
//     time.Sleep(1*time.Millisecond)
//
//     currentMino.drop()
//
//     refreshScreen()
//   }
// }

func aiMoves() {
  for !clock.gameover {

    minMaxHeight := 18
    var optx, opttheta int

    for j:=0; j<4; j++ {
      for i:=0; i<10; i++ {
        currentMino.moveLeft()
      }

      // Check for all columns.
      for i:=0; i<10; i++ {
        dstMino := *currentMino
        dstMino.putBottom()
        h,maxHeight := 0,0

        for k:=0;k<10;k++ {
          for m:=17; m>=0; m-- {
            if board.colors[k][m]!=blankColor {
              h = 18-m
            }
          }

          maxCellHeight := 0

          for _, cell := range dstMino.cells() {
            if cell.x == k {
              if 18-cell.y > maxCellHeight {
                maxCellHeight = 18-cell.y
              }
            }
          }
          if maxCellHeight > 0 {
            h = maxCellHeight
          }

          if h>maxHeight {
            maxHeight = h
          }
          h = 0
        }

        if maxHeight<minMaxHeight {
          minMaxHeight = maxHeight
          optx = i
          opttheta = j
        }

        currentMino.moveRight()
      }

      // After we're done checking, place mino somewhere in the middle.
      for i:=0; i<4; i++ {
        currentMino.moveLeft()
      }

      // Now it's safe to rotate the mino.
      currentMino.rotateLeft()
    }

    // Fix position to original

    for i:=0; i<opttheta; i++ {
      currentMino.rotateLeft()
    }


    for i:=0; i<9; i++ {
      currentMino.moveLeft()
    }

    for i:=0; i<optx; i++ {
      currentMino.moveRight()
    }

    currentMino.drop()

    time.Sleep(1000*time.Millisecond)

    refreshScreen()
  }
}
