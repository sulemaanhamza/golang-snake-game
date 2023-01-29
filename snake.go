package main

type Snake struct {
	parts          []*Point
	velRow, velCol int
	symbol         rune
}

func GetSnakeHead() *Point {
	return snake.parts[len(snake.parts)-1]
}

func UpdateSnake() {
	head := GetSnakeHead()
	snake.parts = append(snake.parts, &Point{row: head.row + snake.velRow, col: head.col + snake.velCol})

	if !AppleIsInsideSnake() {
		snake.parts = snake.parts[1:]
	} else {
		score++
	}

	//
	if IsSnakeHittingTheWall() || IsSnakeEatingItSelf() {
		isGameOver = true
	}
}

func IsSnakeHittingTheWall() bool {
	head := GetSnakeHead()
	return head.row < 0 || head.row >= GameFrameHeight || head.col < 0 || head.col >= GameFrameWidth
}

func IsSnakeEatingItSelf() bool {
	head := GetSnakeHead()
	for _, p := range snake.parts[:len(snake.parts)-1] {
		if p.row == head.row && p.col == head.col {
			return true
		}
	}

	return false
}

func PrintSnake() {
	for _, p := range snake.parts {
		PrintFilledRectInGameFrame(p.row, p.col, 1, 1, snake.symbol)
	}
}
