import time
#defining the column class
class column():
    "this is the column"
    def __init__(self, _flooramount, _minimumFloor, _maximumFloor, _amountOfElevator):
        self.flooramount = _flooramount
        self.minimumFloor = _minimumFloor
        self.maximumFloor = _maximumFloor
        self.amountOfElevator = _amountOfElevator
        self.ElevatorList = []
        self.FloorButtonList = []
        
        #elevator list
        for i in range (_amountOfElevator):
           elevatorlist = elevator(i + 1, _flooramount, 9, 999)
           self.ElevatorList.append(elevatorlist)
        
        #floor buttons 
        for i in range(1, 11):
            self.FloorButtonList.append(i)

    def listUpdate(self, elevatorList, _currentFloor):
        elevatorList.append(_currentFloor)
        elevatorList.sort()
        

    def elevPriority(self, _floorRequest, _direction):
            score = 0
            bestElevator = None
            comparaisonDistance = 999
            distance = 0
        
            if _direction == 'up':
                
                for elevator in self.ElevatorList:
                
                    print('elevator number', elevator.Id, 'on floor', elevator._currentFloor, 'elevator direction', elevator.direction)
                    if (elevator._currentFloor < _floorRequest and elevator.direction == 'up'):
                        score = 3
                        if (score <= 3):
                            bestElevator = elevator
                    
                            return elevator
                    elif (elevator._currentFloor == _floorRequest):
                        score = 2
                        if (score <= 2):
                            bestElevator = elevator
                            
                            return elevator
                    elif (elevator.direction == 'idle'):
                        distance = abs(int(elevator._currentFloor)-int(_floorRequest))
                        score = 1
                        if (distance < comparaisonDistance):
                            if(score <=1):
                                bestElevator = elevator
                                comparaisonDistance = distance
                        
               
                
            elif _direction == 'down':
                
                for elevator in self.ElevatorList:
                    print('elevator number', elevator.Id, 'on floor', elevator._currentFloor, 'elevator direction ', elevator.direction)
                    if(elevator._currentFloor > _floorRequest and elevator.direction == 'down'):
                        score = 5
                        if (score <= 5):
                            elevator = bestElevator
                            
                            return elevator
                    elif (elevator._currentFloor == _floorRequest):
                        score = 4
                        if (score <= 4):
                            elevator = bestElevator
                        
                            return elevator
                    elif (elevator.direction == 'idle'):
                        distance = abs(int(elevator._currentFloor)-int(_floorRequest))
                        score = 1
                        if (distance < comparaisonDistance):
                            if(score <=1):
                                elevator = bestElevator
                                comparaisonDistance = distance
                        
                    
                           


                   


    def move_up(self, _floorRequest, elevator, currentFloor):
        while currentFloor != int(_floorRequest):
            currentFloor += 1
            print(currentFloor)
        elevator.open_door()
        time.sleep(.5)
    def move_down(self, _floorRequest, elevator, currentFloor):
        while currentFloor != int(_floorRequest):
            currentFloor -= 1
            print(currentFloor)
        elevator.open_door()
        time.sleep(.5)

    def move_upIn(self, _insideRequest, elevator, currentFloor):
        while currentFloor != int(_insideRequest):
            currentFloor += 1
            print(currentFloor)
        elevator.open_door(self)
        time.sleep(.5)
    def move_downIn(self, _insideRequest, elevator, currentFloor):
        while currentFloor != int(_insideRequest):
            currentFloor -= 1
            print(currentFloor)
        elevator.open_door(self)
        time.sleep(.5)

    def RequestFloor(self,elevator1, _direction, currentFloor):
        
        
        if (_direction == 'down'):
            print('//////////////////// elevator goes to floor', _floorRequest, ' ///////////////////////')
            self.move_down(_floorRequest, elevator, currentFloor)
            
        elif (_direction == 'up'):
            print('//////////////////// elevator goes to floor', _floorRequest, ' ///////////////////////')
            self.move_up(_floorRequest, elevator, currentFloor)
        return
    
    def RequestFloorInside(self, elevator2, _direction, currentFloor):
       
        if (_direction =='down'):
            print('///////////////////// elevator goes to floor', _insideRequest, 'with customer ///////////////')
            self.move_downIn(_insideRequest, elevator, currentFloor)

        elif (_direction =='up'):
            print('///////////////////// elevator goes to floor', _insideRequest, 'with customer ///////////////')
            self.move_upIn(_insideRequest, elevator, currentFloor)
        return


class elevator():
    "creates the elevator class"

    def __init__(self, Id, _flooramount, distance, comparaisonDistance):
        self.distance = distance
        self.Id = Id
        self.requestList = []
        self._currentFloor = 1
        self._direction = 'idle'
        self._status = 'idle'
        self._door = 'closed'


    def open_door(self):
        print('open doors')
        time.sleep(1)
        print('close doors')

    
        

class floorRequestButton():
    "buttons on each floor"
    def __init__(self, _floor, _direction, _status = 'off'):
        self.floor = _floor
        self.direction = _direction
        self.status = _status
    

 #scnenario 1
#column1 = column(10, 1, 10, 2)
#column1.ElevatorList[0]._currentFloor = 2
#column1.ElevatorList[0].direction = 'idle'
#column1.ElevatorList[1]._currentFloor = 6
#column1.ElevatorList[1].direction = 'idle'
#_floorRequest = 3
#_direction = 'up'
#_insideRequest = 7
#_directioninside = 'up'

#print('///////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest,' /////////////////////')
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator1, _floorRequest, 2)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 3)
#column1.RequestFloor( elevator2, _insideRequest, 3)

#scenario 2
#column1 = column(10, 1, 10, 2)
#column1.ElevatorList[0]._currentFloor = 10
#column1.ElevatorList[0].direction = 'idle'
#column1.ElevatorList[1]._currentFloor = 3
#column1.ElevatorList[1].direction = 'idle'
#_floorRequest = 1
#_direction = 'down'
#_insideRequest = 6
#_directioninside = 'up'

#print('////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest, '///////////////////////////')
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator1, _floorRequest, 3)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 1)
#column1.RequestFloor( elevator2, _insideRequest, 1)


#column1.ElevatorList[1]._currentFloor = 6

#floorRequest = 3
#_direction = 'down'
#_insideRequest = 5
#_directioninside = 'up'
#print('////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest, '////////////////////////////' )
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator1, _floorRequest, 6)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 3)
#column1.RequestFloor( elevator2, _insideRequest, 3)

#column1.ElevatorList[0]._currentFloor = 10
#column1.ElevatorList[1]._currentFloor = 5
#_floorRequest = 9
#_direction = 'down'
#_insideRequest = 2
#_directioninside = 'down'
#print('////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest, '////////////////////////////' )
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator1, _floorRequest, 10)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 9)
#column1.RequestFloorInside( elevator2, _insideRequest, 9)

#scenario 3

#column1 = column(10, 1, 10, 2)
#column1.ElevatorList[0]._currentFloor = 10
#column1.ElevatorList[0].direction = 'idle'
#column1.ElevatorList[1]._currentFloor = 3
#column1.ElevatorList[1].direction = 'up'
#_floorRequest = 3
#_direction = 'down'
#_insideRequest = '2'
#_directioninside = 'down'

#print('///////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest,' /////////////////////')
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator, _floorRequest, 10)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 3)
#column1.RequestFloor( elevator2, _insideRequest, 3)

#column1.ElevatorList[1]._currentFloor = 6

#floorRequest = 10
#direction = 'up'
#_insideRequest = 3
#_directioninside = 'down'
#print('////////////////// user at floor', _floorRequest, 'going', _directioninside, 'to floor', _insideRequest, '////////////////////////////' )
#elevator1 = column1.elevPriority(_floorRequest, _direction)
#column1.RequestFloor(elevator1, _floorRequest, 6)
#elevator2 = column1.RequestFloorInside(_insideRequest, _directioninside, 10)
#column1.RequestFloor( elevator2, _insideRequest, 10)

