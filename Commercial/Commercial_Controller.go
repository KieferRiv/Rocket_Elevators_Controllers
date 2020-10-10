package main

import "fmt"

func main() {
	/*Battery1 := battery(4, -6, 66)
	Battery1.ColumnList[1].ElevatorList[0].CurrentFloor = 20
	Battery1.ColumnList[1].ElevatorList[0].Direction = "Down"
	Battery1.ColumnList[1].ElevatorList[0].RequestedFloor = 5
	Battery1.ColumnList[1].ElevatorList[1].CurrentFloor = 3
	Battery1.ColumnList[1].ElevatorList[1].Direction = "Up"
	Battery1.ColumnList[1].ElevatorList[1].RequestedFloor = 15
	Battery1.ColumnList[1].ElevatorList[2].CurrentFloor = 13
	Battery1.ColumnList[1].ElevatorList[2].Direction = "Down"
	Battery1.ColumnList[1].ElevatorList[2].RequestedFloor = 1
	Battery1.ColumnList[1].ElevatorList[3].CurrentFloor = 15
	Battery1.ColumnList[1].ElevatorList[3].Direction = "Down"
	Battery1.ColumnList[1].ElevatorList[3].RequestedFloor = 2
	Battery1.ColumnList[1].ElevatorList[4].CurrentFloor = 6
	Battery1.ColumnList[1].ElevatorList[4].Direction = "Down"
	Battery1.ColumnList[1].ElevatorList[4].RequestedFloor = 1
	fmt.Println("User at floor 1 going up going to floor 20")
	Battery1.ColumnList[1].BestElevator(1, 20)*/

	//scenario 2
	/*Battery1 := battery(4, -6, 66)
	Battery1.ColumnList[2].ElevatorList[0].CurrentFloor = 1
	Battery1.ColumnList[2].ElevatorList[0].Direction = "Up"
	Battery1.ColumnList[2].ElevatorList[0].RequestedFloor = 21
	Battery1.ColumnList[2].ElevatorList[1].CurrentFloor = 23
	Battery1.ColumnList[2].ElevatorList[1].Direction = "Up"
	Battery1.ColumnList[2].ElevatorList[1].RequestedFloor = 28
	Battery1.ColumnList[2].ElevatorList[2].CurrentFloor = 33
	Battery1.ColumnList[2].ElevatorList[2].Direction = "Down"
	Battery1.ColumnList[2].ElevatorList[2].RequestedFloor = 1
	Battery1.ColumnList[2].ElevatorList[3].CurrentFloor = 40
	Battery1.ColumnList[2].ElevatorList[3].Direction = "Down"
	Battery1.ColumnList[2].ElevatorList[3].RequestedFloor = 24
	Battery1.ColumnList[2].ElevatorList[4].CurrentFloor = 39
	Battery1.ColumnList[2].ElevatorList[4].Direction = "Down"
	Battery1.ColumnList[2].ElevatorList[4].RequestedFloor = 1
	fmt.Println("User at floor 1 going to floor 36")
	Battery1.ColumnList[2].BestElevator(1, 36)*/

	//scenario 3
	/*Battery1 := battery(4, -6, 66)
	Battery1.ColumnList[3].ElevatorList[0].CurrentFloor = 58
	Battery1.ColumnList[3].ElevatorList[0].Direction = "Down"
	Battery1.ColumnList[3].ElevatorList[0].RequestedFloor = 1
	Battery1.ColumnList[3].ElevatorList[1].CurrentFloor = 50
	Battery1.ColumnList[3].ElevatorList[1].Direction = "Up"
	Battery1.ColumnList[3].ElevatorList[1].RequestedFloor = 60
	Battery1.ColumnList[3].ElevatorList[2].CurrentFloor = 46
	Battery1.ColumnList[3].ElevatorList[2].Direction = "Up"
	Battery1.ColumnList[3].ElevatorList[2].RequestedFloor = 58
	Battery1.ColumnList[3].ElevatorList[3].CurrentFloor = 1
	Battery1.ColumnList[3].ElevatorList[3].Direction = "Up"
	Battery1.ColumnList[3].ElevatorList[3].RequestedFloor = 54
	Battery1.ColumnList[3].ElevatorList[4].CurrentFloor = 60
	Battery1.ColumnList[3].ElevatorList[4].Direction = "Down"
	Battery1.ColumnList[3].ElevatorList[4].RequestedFloor = 1
	fmt.Println("User at floor 54 going to floor 1")
	Battery1.ColumnList[3].BestElevator(54, 1)*/

	//scenario 4
	/*Battery1 := battery(4, -6, 66)
	Battery1.ColumnList[0].ElevatorList[0].CurrentFloor = -4
	Battery1.ColumnList[0].ElevatorList[0].Direction = "Idle"
	Battery1.ColumnList[0].ElevatorList[1].CurrentFloor = 1
	Battery1.ColumnList[0].ElevatorList[1].Direction = "Idle"
	Battery1.ColumnList[0].ElevatorList[2].CurrentFloor = -3
	Battery1.ColumnList[0].ElevatorList[2].Direction = "Down"
	Battery1.ColumnList[0].ElevatorList[2].RequestedFloor = -5
	Battery1.ColumnList[0].ElevatorList[3].CurrentFloor = -6
	Battery1.ColumnList[0].ElevatorList[3].Direction = "Up"
	Battery1.ColumnList[0].ElevatorList[3].RequestedFloor = 1
	Battery1.ColumnList[0].ElevatorList[4].CurrentFloor = -1
	Battery1.ColumnList[0].ElevatorList[4].Direction = "Down"
	Battery1.ColumnList[0].ElevatorList[4].RequestedFloor = -6
	fmt.Println("User at basement 3 going to floor 1")
	Battery1.ColumnList[0].BestElevator(-3, 1)*/

}

type Battery struct {
	ColumnList       []Column
	FloorButtonsList []FloorButtons
	AmountOfColumns  int
	LowestFloor      int
	AmountOfFloors   int
}

func battery(AmountOfColumns, LowestFloor, AmountOfFloors int) Battery {
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

	for i := 1; i <= ElevatorsPerColumn; i++ {
		column.ElevatorList = append(column.ElevatorList, elevator(i))
	}
	return *column

}

//moves the elevator
func (el *Elevator) Move(elevator Elevator, UserFloor int) {

	fmt.Println("moving")
	if UserFloor < elevator.CurrentFloor {
		i := elevator.CurrentFloor
		for UserFloor < i {
			i--
			fmt.Println(i)
		}
		fmt.Println("opening Doors")
		fmt.Println("closing doors")
	} else if UserFloor > elevator.CurrentFloor {
		i := elevator.CurrentFloor
		for UserFloor > i {
			i++
			fmt.Println(i)
		}
		fmt.Println("opening Doors")
		fmt.Println("closing doors")
	} else {
		fmt.Println("opening doors")
		fmt.Println("closing doors")
	}

}

//moving to requested floor from user
func (el *Elevator) MoveUser(elevator Elevator, RequestedFloor, UserFloor int) {
	if RequestedFloor < elevator.CurrentFloor {
		fmt.Println("moving Down")
		i := UserFloor
		for RequestedFloor < i {
			i--
			fmt.Println(i)
		}
		fmt.Println("opening Doors")
		fmt.Println("closing doors")
	} else {
		fmt.Println("moving up")
		i := UserFloor
		for RequestedFloor > i {
			i++
			fmt.Println(i)
		}
		fmt.Println("opening Doors")
		fmt.Println("closing doors")
	}
}

//chooses the best elevator
func (c *Column) BestElevator(UserFloor, RequestedFloor int) Elevator {
	var bestElevator Elevator

	for _, elevator := range c.ElevatorList {
		if UserFloor == 1 {
			if elevator.CurrentFloor == 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.RequestedFloor == 1 {
				fmt.Println("requested floor 1")
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			}
		} else if UserFloor > 1 {
			if elevator.Direction == "Down" && UserFloor < elevator.CurrentFloor {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor == 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			}
		} else {
			if elevator.Direction == "Up" && UserFloor > elevator.CurrentFloor {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor != 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			} else if elevator.Direction == "Idle" && elevator.CurrentFloor == 1 {
				bestElevator = elevator
				bestElevator.Move(elevator, UserFloor)
				bestElevator.MoveUser(elevator, RequestedFloor, UserFloor)
				return bestElevator
			}

		}
	}
	return bestElevator

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
