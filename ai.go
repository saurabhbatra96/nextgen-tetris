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

    minMaxHeight := 18
    var optx, opttheta int
    for i:=0; i<9; i++ {

      var j int
      // Movements
      for j=0; j<3; j++ {

        // Rotations

        dstMino := *currentMino
        dstMino.putBottom()
        h,maxHeight := 0,0

        for k:=0;k<10;k++ {
          for m:=17; m>=0; m-- {
            if board.colors[k][m]!=blankColor {
              if _, err = f.WriteString(strconv.Itoa(k)+","+strconv.Itoa(17-m)+" is colored.\t"); err != nil {
                panic(err)
              }
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
          if _, err = f.WriteString(strconv.Itoa(h)+"\n"); err != nil {
            panic(err)
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
          if _, err = f.WriteString(strconv.Itoa(minMaxHeight)+","+strconv.Itoa(optx)+","+strconv.Itoa(opttheta)+"\n"); err != nil {
            panic(err)
          }
        }

        testMino := *currentMino
        testMino.forceRotateLeft()
        if testMino.conflicts() {
          break
        }

        currentMino.rotateLeft()
        if _, err = f.WriteString("rotleft: "+strconv.Itoa(j)+"\n"); err != nil {
          panic(err)
        }
      }

      testMino := *currentMino
      testMino.forceRotateLeft()
      if !testMino.conflicts() {
        currentMino.rotateLeft()
      } else {
        for z:=j;z>=0;z-- {
          currentMino.rotateRight()
        }
      }
      if _, err = f.WriteString("movright\n"); err != nil {
        panic(err)
      }
      currentMino.moveRight()
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
