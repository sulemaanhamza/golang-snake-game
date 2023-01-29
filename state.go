package main

func InitGameState() {
	snake = &Snake{
		parts: []*Point{
			{row: 9, col: 3},
			{row: 8, col: 3},
			{row: 7, col: 3},
			{row: 6, col: 3},
			{row: 5, col: 3},
		},
		velRow: -1,
		velCol: 0,
		symbol: SnakeSymbol,
	}
	apple = &Apple{
		point:  &Point{row: 10, col: 10},
		symbol: AppleSymbol,
	}
}

func DrawState() {
	if isGamePaused {
		return
	}

	screen.Clear()

	PrintString(0, 0, debugLog)
	PrintGameFrame()
	PrintSnake()
	PrintApple()

	screen.Show()
}

func UpdateState() {
	if isGamePaused {
		return
	}

	UpdateSnake()
	UpdateApple()
}
