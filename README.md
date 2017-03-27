# NextGen-Tetris
![Tetris](output.gif)

The aim of this project is to understand how different AI strategies perform while playing a game of Tetris on a 10x18 board.

The end gold standard of the AI being a genetic programming algorithm which should ideally be able to play on for ages!

## Analysis

Randomized bot (makes random choices):
* Average score: 166
* Number of trials: 6685
* Sum: 1111315

Greedy bot (makes greedy choices, i.e. selects the move which maximizes the score):
* Average: 210
* Number: 10000
* Sum: 2105309

Bot that minimizes the maximum height at every step:
* Average: 726
* Number: 10049
* Sum: 7300649

Prev bot + awards line clears:
* Average: 1382
* Number: 9281
* Sum: 12831920

## Tetris

Shoutout to https://github.com/k0kubun/ for the Tetris shell!
