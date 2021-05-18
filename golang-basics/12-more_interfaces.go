package main

import (
	"fmt"
)

type RailRoadWideChecker interface {
	CheckRailsWidth() int
}

type RailRoad struct {
	Width int
}

func (r *RailRoad) IsCorrectSizeTrain(checker RailRoadWideChecker) bool {
	return checker.CheckRailsWidth() != r.Width
}

type Train struct {
	TrainWidth int
}

func (t *Train) CheckRailsWidth() int {
	return t.TrainWidth
}

func main() {
	railRoad := RailRoad{Width: 10}
	passengerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}

	canPassengerTrainPass := railRoad.IsCorrectSizeTrain(&passengerTrain)
	canCargoTrainPass := railRoad.IsCorrectSizeTrain(&cargoTrain)

	fmt.Printf("Can passenger train pass? %t\n", canPassengerTrainPass)
	fmt.Printf("Can cargo train pass? %t\n", canCargoTrainPass)

}
