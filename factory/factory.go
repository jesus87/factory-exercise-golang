package factory

import (
	"fmt"
	"sync"

	".main.go/assemblyspot"
	".main.go/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) []*assemblyspot.AssemblySpot {
	vehicleList := f.generateVehicleLots(amountOfVehicles)

	var wg sync.WaitGroup
	aSpots := []*assemblyspot.AssemblySpot{}

	for i, vehicle := range vehicleList {
		fmt.Println("Assembling vehicle...", i)
		if len(f.AssemblingSpots) == 0 {
			wg.Wait()
		}

		wg.Add(1)
		idleSpot := <-f.AssemblingSpots
		aSpots = append(aSpots, idleSpot)
		go f.process(vehicle, idleSpot, &wg)
	}

	return aSpots
}

func (f *Factory) process(v vehicle.Car, idleSpot *assemblyspot.AssemblySpot, wg *sync.WaitGroup) {
	defer wg.Done()

	idleSpot.SetVehicle(&v)
	vehicle, err := idleSpot.AssembleVehicle()

	if err != nil {
		return
	}

	vehicle.TestingLog = f.testCar(vehicle)
	vehicle.AssembleLog = idleSpot.GetAssembledLogs()

	idleSpot.SetVehicle(nil)
	f.AssemblingSpots <- idleSpot
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	var vehicles = []vehicle.Car{}
	var index = 0

	for {
		vehicles = append(vehicles, vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
