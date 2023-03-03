package main

import "go-rock-paper-scissor/game"

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	newGame := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go newGame.Rounds()
	newGame.ClearScreen()
	newGame.PrintIntro()

	for {
		newGame.RoundChan <- 1
		<-newGame.RoundChan

		if newGame.Round.RoundNumber > 3 {
			break
		}

		if !newGame.PlayRound() {
			newGame.RoundChan <- -1
			<-newGame.RoundChan
		}
	}

	newGame.PrintSummary()
}
