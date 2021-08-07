package main

import (
	"fmt"

	".main.go/factory"
)

const carsAmount = 100

func main() {
	factory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	aSpots := factory.StartAssemblingProcess(carsAmount)
	for _, aSpot := range aSpots {
		if aSpot == nil {
			continue
		}
		v := aSpot.GetAssembledVehicle()
		if v == nil {
			continue
		}

		fmt.Println(fmt.Sprintf("id: %v, testingLog: %v, assemblyLog: %v",
			v.Id,
			v.TestingLog,
			aSpot.GetAssembledLogs()))
	}
}
