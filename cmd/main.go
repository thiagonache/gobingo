package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gookit/color"
)

const rows, cols, totalNumbers, totalBalls = 8, 10, 75, 24

type card map[int]struct{}
type board [rows][cols]int

func cardGenerator() card {
	playerCard := card{}
	for totalGenerated := 0; totalGenerated < totalBalls; totalGenerated++ {
		rand.Seed(time.Now().UnixNano())
		// We have to add one since rand.Intn returns [0,max-1] as range for the
		// random value
		n := rand.Intn(totalNumbers) + 1
		_, found := playerCard[n]
		tries := 0
		for found {
			if tries > totalNumbers {
				log.Fatalf("Max tries of %d has exceeded", totalNumbers)
			}
			rand.Seed(time.Now().UnixNano())
			n = rand.Intn(totalNumbers) + 1
			_, found = playerCard[n]
			tries++
		}
		playerCard[n] = struct{}{}
	}

	return playerCard
}

func ballBoardMarker(card card) board {
	board := board{}
	count := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			key := count + 1
			board[x][y] = key
			count++
		}
	}
	return board
}

func boardPrinter(board board, card card, drewBalls card) {
	for _, row := range board {
		for _, col := range row {
			_, exist := card[col]
			_, exist2 := drewBalls[col]
			if exist && exist2 {
				color.Green.Printf("%4d", col)
			} else if exist {
				color.Blue.Printf("%4d", col)
			} else {
				fmt.Printf("%4d", col)
			}
		}
		fmt.Println()

	}
}

func drawBall(drewBalls card) card {
	rand.Seed(time.Now().UnixNano())
	// We have to add one since rand.Intn returns [0,max-1] as range for the
	// random value
	n := rand.Intn(totalNumbers) + 1
	_, found := drewBalls[n]
	tries := 0
	for found {
		if tries > totalNumbers {
			log.Fatalf("Max tries of %d has exceeded", totalNumbers)
		}
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(totalNumbers) + 1
		_, found = drewBalls[n]
		tries++
	}
	drewBalls[n] = struct{}{}

	return drewBalls
}

func main() {
	playerCard := cardGenerator()
	drewBalls := card{}
	var bingo bool
	for !bingo {
		var waitForEnter string
		board := ballBoardMarker(playerCard)
		boardPrinter(board, playerCard, drewBalls)
		fmt.Println()
		fmt.Print("Press enter to draw a ball...")
		fmt.Scanln(&waitForEnter)
		drewBalls = drawBall(drewBalls)
		fmt.Println(drewBalls)
	}

}
