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
  f, err := os.OpenFile("/home/saurabh/projects/tetris-data/data", os.O_RDWR|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  for !clock.gameover {
    for i:=0; i<9; i++ {
      currentMino.moveLeft()
    }

    minAggrHeight := 19
    var optx, opttheta int
    for i:=0; i<9; i++ {

      // Movements
      for j:=0; j<3; j++ {

        // Rotations

        dstMino := *currentMino
        dstMino.putBottom()
        totHeight,h,aggrHeight := 0,0,0

        for k:=0;k<10;k++ {
          if _, err = f.WriteString(strconv.Itoa(h)+"\t"); err != nil {
            panic(err)
          }
          for m:=0; m<18; m++ {
            if board.colors[k][m]!=blankColor {
              h = m
            }
          }

          for _, cell := range dstMino.cells() {
            if cell.x == k && cell.y == h {
              h++
            }
          }
          if _, err = f.WriteString(strconv.Itoa(h)+"\n"); err != nil {
            panic(err)
          }

          totHeight += h
          h = 0
        }

        if totHeight != 0 {
          aggrHeight = totHeight/10
        }



        if aggrHeight<minAggrHeight {
          minAggrHeight = aggrHeight
          optx = i
          opttheta = j
          if _, err = f.WriteString(strconv.Itoa(totHeight)+","+strconv.Itoa(optx)+","+strconv.Itoa(opttheta)+"\n"); err != nil {
            panic(err)
          }
        }

        currentMino.rotateLeft()
        if _, err = f.WriteString("rotleft\n"); err != nil {
          panic(err)
        }
      }
      if _, err = f.WriteString("movright\n"); err != nil {
        panic(err)
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


    currentMino.drop()

    time.Sleep(1000*time.Millisecond)

    refreshScreen()
  }
}
