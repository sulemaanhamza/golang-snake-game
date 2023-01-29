package main

import "math/rand"

type Apple struct {
	point  *Point
	symbol rune
}

func AppleIsInsideSnake() bool {
	for _, p := range snake.parts {
		if p.row == apple.point.row && p.col == apple.point.col {
			return true
		}
	}

	return false
}

func UpdateApple() {
	for AppleIsInsideSnake() {
		apple.point.row, apple.point.col = rand.Intn(GameFrameHeight), rand.Intn(GameFrameWidth)
	}
}

func PrintApple() {
	PrintFilledRectInGameFrame(apple.point.row, apple.point.col, 1, 1, apple.symbol)
}
