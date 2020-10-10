package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type Battery struct {
	ColumnList       []Column
	FloorButtonsList []FloorButtons
	status           string
	AmountOfColumns  int
	LowestFloor      int
	AmountOfFloors   int
}

func battery(Status string, AmountOfColumns, LowestFloor, AmountOfFloors int) Battery {
	battery := new(Battery)
	battery.AmountOfColumns = AmountOfColumns
	battery.LowestFloor = LowestFloor
	battery.AmountOfFloors = LowestFloor

	//creating column list
	for i := 1; i <= AmountOfColumns; i++ {
		battery.ColumnList = append(battery.ColumnList, column(i, 5))
	}
	return *battery

}

type Column struct {
	ID                 int
	ElevatorsPerColumn int
	ElevatorList       []Elevator
}

func column(ID, ElevatorsPerColumn int) Column {
	column := new(Column)
	column.ElevatorsPerColumn = ElevatorsPerColumn
	column.ID = ID
	fmt.Println(column.ElevatorList)

	for i := 1; i <= ElevatorsPerColumn; i++ {
		column.ElevatorList = append(column.ElevatorList, elevator(i))
	}
	return *column

}

//chooses the best elevator
func (c *Column) BestElevator(UserFloor int) Elevator {
	var bestElevator Elevator

	for _, elevator := range c.ElevatorList {

		if UserFloor == 1 {
			if elevator.CurrentFloor == 1 {
				bestElevator = elevator
				return bestElevator
			} else if elevator.RequestedFloor == 1 {
				bestElevator = elevator
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				return bestElevator
			}
		} else if UserFloor > 1 {
			if elevator.Direction == "Down" && UserFloor < elevator.CurrentFloor {
				bestElevator = elevator
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor == 1 {
				bestElevator = elevator
				return bestElevator
			} else {
				bestElevator = elevator
				return bestElevator
			}
		} else {
			if elevator.Direction == "Up" && UserFloor > elevator.CurrentFloor {
				bestElevator = elevator
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor == 1 {
				bestElevator = elevator
				return bestElevator
			}

		}
	}
}

type Elevator struct {
	ID             int
	Direction      string
	CurrentFloor   int
	RequestedFloor int
}

func elevator(ID int) Elevator {
	elevator := new(Elevator)
	elevator.ID = ID

	return *elevator
}

type FloorButtons struct {
	Direction    string
	CurrentFloor int
}
