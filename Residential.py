#defining the column class
class column:
    "this is the column"

    def __init__(self, _flooramount, _minimumFloor, _maximumFloor, _amountOfElevator, _status = 'online'):
        self.flooramount = _flooramount
        self.minimumFloor = _minimumFloor
        self.maximumFloor = _maximumFloor
        self.amountOfElevator = _amountOfElevator
        self.status = _status
   

    #elevator list
    ElevatorList = []
    for i in range (1, 3):
        ElevatorList.append(i)
    print(ElevatorList)

    #floor buttons
    FloorButtonList = []
    for i in range(1, 11):
        FloorButtonList.append(i)
    print(FloorButtonList)

    
    def get_data(self):
        print(f'{self.flooramount},{self.minimumFloor},{self.maximumFloor}, {self.status}')



class elevator():
    "creates the elevator class"

    def __init__(self, _id, _curentFloor, _direction = None, _requestList = None,  _status = 'idle', _door = 'closed'):
        self.id = _id
        self.status = _status
        self.currentFloor = _curentFloor
        self.direction = _direction
        self.requestList = _requestList
        self.door = _door

    def get_data1(self):
            print(f'{self.currentFloor},{self.id},{self.status},{self.direction},{self.requestList},{self.door}')
    class inside_buttons():
        def __init__(self, _flooramount, _direction, _status = 'off', _minimumFloor = 1):
            self.flooramount = _flooramount
            self.direction = _direction
            self.status = _status
            self.minimumFloor = _minimumFloor
    
        def get_data2(self):
            print(f'{self.flooramount},{self.direction},{self.status},{self.minimumFloor}')

class floorRequestButton():
    "buttons on each floor"
    def __init__(self, _floor, _direction, _status = 'off'):
        self.floor = _floor
        self.direction = _direction
        self.status = _status

elevator1 = elevator(1, 1)

elevator1.get_data1()
#priotities for the elevators
def elevPriority():
    print("hi")