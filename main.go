package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"math/rand"
	"time"
)

const SnakeSymbol = 0x2588
const AppleSymbol = 0x25CF
const GameFrameWidth = 30
const GameFrameHeight = 15
const GameFrameSymbol = '‖'

type Point struct {
	row, col int
}

var screen tcell.Screen
var snake *Snake
var apple *Apple
var isGamePaused bool
var isGameOver bool
var score int
var debugLog string

func main() {

	rand.Seed(time.Now().UnixNano())

	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	for !isGameOver {
		HandleUserInput(ReadInput(inputChan))
		UpdateState()
		DrawState()

		time.Sleep(75 * time.Millisecond)
	}

	screenWidth, screenHeight := screen.Size()
	PrintStringCentered(screenHeight/2, screenWidth/2, "Game Over!")
	PrintStringCentered(screenHeight/2+1, screenWidth/2, fmt.Sprintf("Your score is %d", score))
	screen.Show()
	time.Sleep(3 * time.Second)
	screen.Fini()
}
