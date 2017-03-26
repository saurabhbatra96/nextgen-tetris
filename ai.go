package main

import (
  "time"
)

func aiMoves() {
  // Change the scope of these to global once genetic programming starts.
  heightMult, lineMult := -1,8

  for !clock.gameover {

    maxCost := -1000
    var optx, opttheta int

    for j:=0; j<4; j++ {
      for i:=0; i<10; i++ {
        currentMino.moveLeft()
      }

      // Check for all columns.
      for i:=0; i<11; i++ {
        dstMino := *currentMino
        dstMino.putBottom()
        h,maxHeight,numlines := 0,0,0

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

            // Check if this row is a line.
            for _, cell := range dstMino.cells() {
              hasLine := true
              for l:=0; l<10; l++ {
                if cell.y>=0 && board.colors[l][cell.y]==blankColor && cell.x!=l {
                  foundInMino := false
                  for _, nestedCell := range dstMino.cells() {
                    if nestedCell.y==cell.y && nestedCell.x==l {
                      foundInMino = true
                      break
                    }
                  }
                  if !foundInMino {
                    hasLine = false
                  }
                }
              }

              if hasLine {
                numlines++
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

        cost := heightMult*maxHeight + lineMult*numlines

        if cost>maxCost {
          maxCost = cost
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

    time.Sleep(1*time.Millisecond)

    refreshScreen()
  }
}
