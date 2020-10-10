using System;
using System.Collections.Generic;

namespace odyssey
{
    class Program
    {
        static void Main(string[] args)
        {
            //scenario 1 test
            //Battery battery1 = new Battery(4, -6, 66);
            //battery1.status = "online";
            //battery1.ColumnList[1].ElevatorList[0].CurrentFloor = 20;
            //battery1.ColumnList[1].ElevatorList[0].Direction = "Down";
            //battery1.ColumnList[1].ElevatorList[0].RequestedFloor = 5;
            //battery1.ColumnList[1].ElevatorList[1].CurrentFloor = 3;
            //battery1.ColumnList[1].ElevatorList[1].Direction = "Up";
            //battery1.ColumnList[1].ElevatorList[1].RequestedFloor = 15;
            //battery1.ColumnList[1].ElevatorList[2].CurrentFloor = 13;
            //battery1.ColumnList[1].ElevatorList[2].Direction = "Down";
           // battery1.ColumnList[1].ElevatorList[2].RequestedFloor = 1;
            //battery1.ColumnList[1].ElevatorList[3].CurrentFloor = 15;
            //battery1.ColumnList[1].ElevatorList[3].Direction = "Down";
            //battery1.ColumnList[1].ElevatorList[3].RequestedFloor = 2;
            //battery1.ColumnList[1].ElevatorList[4].CurrentFloor = 6;
            //battery1.ColumnList[1].ElevatorList[4].Direction = "Down";
            //battery1.ColumnList[1].ElevatorList[4].RequestedFloor = 1;
            //Console.WriteLine("User is at floor 1 and going to floor 20");
            //battery1.ColumnList[1].BestElevator(1);

            // scenario 2
             //Battery battery1 = new Battery(4, -6, 66);
            //battery1.status = "online";
            //battery1.ColumnList[2].ElevatorList[0].CurrentFloor = 1;
            //battery1.ColumnList[2].ElevatorList[0].Direction = "Up";
            //battery1.ColumnList[2].ElevatorList[0].RequestedFloor = 21;
            //battery1.ColumnList[2].ElevatorList[1].CurrentFloor = 23;
            //battery1.ColumnList[2].ElevatorList[1].Direction = "Up";
            //battery1.ColumnList[2].ElevatorList[1].RequestedFloor = 28;
            //battery1.ColumnList[2].ElevatorList[2].CurrentFloor = 33;
            //battery1.ColumnList[2].ElevatorList[2].Direction = "Down";
            //battery1.ColumnList[2].ElevatorList[2].RequestedFloor = 1;
            //battery1.ColumnList[2].ElevatorList[3].CurrentFloor = 40;
            //battery1.ColumnList[2].ElevatorList[3].Direction = "Down";
           // battery1.ColumnList[2].ElevatorList[3].RequestedFloor = 24;
            //battery1.ColumnList[2].ElevatorList[4].CurrentFloor = 39;
           // battery1.ColumnList[2].ElevatorList[4].Direction = "Down";
            //battery1.ColumnList[2].ElevatorList[4].RequestedFloor = 1;
            //Console.WriteLine("User is at floor 1 and going to floor 36");
            //battery1.ColumnList[2].BestElevator(1);

            //scenario 3
            // Battery battery1 = new Battery(4, -6, 66);
            //battery1.status = "online";
          //  battery1.ColumnList[3].ElevatorList[0].CurrentFloor = 58;
            //battery1.ColumnList[3].ElevatorList[0].Direction = "Down";
            //battery1.ColumnList[3].ElevatorList[0].RequestedFloor = 1;
            //battery1.ColumnList[3].ElevatorList[1].CurrentFloor = 50;
            //battery1.ColumnList[3].ElevatorList[1].Direction = "Up";
            //battery1.ColumnList[3].ElevatorList[1].RequestedFloor = 60;
            //battery1.ColumnList[3].ElevatorList[2].CurrentFloor = 46;
            //battery1.ColumnList[3].ElevatorList[2].Direction = "Up";
            //battery1.ColumnList[3].ElevatorList[2].RequestedFloor = 58;
            //battery1.ColumnList[3].ElevatorList[3].CurrentFloor = 1;
            //battery1.ColumnList[3].ElevatorList[3].Direction = "Up";
            //battery1.ColumnList[3].ElevatorList[3].RequestedFloor = 54;
            //battery1.ColumnList[3].ElevatorList[4].CurrentFloor = 60;
            //battery1.ColumnList[3].ElevatorList[4].Direction = "Down";
            //battery1.ColumnList[3].ElevatorList[4].RequestedFloor = 1;
            //Console.WriteLine("User is at floor 54 and going to floor 1");
            //battery1.ColumnList[3].BestElevator(54);

            //scenario 4
             //Battery battery1 = new Battery(4, -6, 66);
            //battery1.status = "online";
            //battery1.ColumnList[0].ElevatorList[0].CurrentFloor = -4;
            //battery1.ColumnList[0].ElevatorList[0].Direction = "Idle";
            //battery1.ColumnList[0].ElevatorList[1].CurrentFloor = 1;
            //battery1.ColumnList[0].ElevatorList[1].Direction = "Idle";
            //battery1.ColumnList[0].ElevatorList[2].CurrentFloor = -3;
            //battery1.ColumnList[0].ElevatorList[2].Direction = "Down";
            //battery1.ColumnList[0].ElevatorList[2].RequestedFloor = -5;
            //battery1.ColumnList[0].ElevatorList[3].CurrentFloor = -6;
            //battery1.ColumnList[0].ElevatorList[3].Direction = "Up";
            //battery1.ColumnList[0].ElevatorList[3].RequestedFloor = 1;
            //battery1.ColumnList[0].ElevatorList[4].CurrentFloor = -1;
            //battery1.ColumnList[0].ElevatorList[4].Direction = "Down";
            //battery1.ColumnList[0].ElevatorList[4].RequestedFloor = -6;
            //Console.WriteLine("User is at Basement 3 and going to floor 1");
            //battery1.ColumnList[0].BestElevator(-3);
            
           
           
           
           
            
        }
    }
    class Battery
    {
        public List<Column> ColumnList;
        public List<FloorButtons> FloorButtonsList;
        public string status;
        public int AmountOfColumn;
        public int LowestFloor;
        public int AmountOfFloors;

        public Battery(int _AmountOfColumn, int _LowestFloor, int _AmountOfFloors)
        {
            ColumnList = new List<Column>();
            FloorButtonsList = new List<FloorButtons>();
            AmountOfColumn = _AmountOfColumn;
            LowestFloor = _LowestFloor;
            AmountOfFloors = _AmountOfFloors;
            Console.WriteLine(AmountOfColumn);

          
             // creating the column list
             // int = 1 so first column is column 1
            for (int i = 1; i<AmountOfColumn + 1; i++)
            {
                ColumnList.Add( new Column(id:i + 1, 5));
                Console.WriteLine(ColumnList);
            }  
            // creating the list of buttons for the elevator
            // if it s a basement = down
            //if it s a upper floor = up
            //start floors at 1 so doesnt count 0
            for (int i = LowestFloor; i < AmountOfFloors - 5; i++ )
            {
                if (i != 0)
                {
                    if (i < 6)
                    {
                        FloorButtonsList.Add( new FloorButtons("Down", i));
                    }
                    else if (i > 6)
                    {
                        FloorButtonsList.Add( new FloorButtons("Up", i));
                    }
                    else 
                    {
                        FloorButtonsList.Add( new FloorButtons("Down", i));
                        FloorButtonsList.Add( new FloorButtons("Up", i));
                    }
                }}
                //Console.WriteLine(FloorButtonsList.Count);
                Console.WriteLine(LowestFloor);
                
            }
            // which column you are getting 
            public Column FindBestColumn(string _Direction, int _RequestedFloor)
            {
                for (int i = LowestFloor; i < AmountOfFloors -5; i++ )
            {// if you are on ground floor 
                if (i == 6)
                {
                    if (_Direction == "Down")
                    {
                        var BestCol = (ColumnList[0]);
                        return BestCol;
                    }
                    else if (_Direction == "Up")
                    {
                        if (_RequestedFloor > 1 && _RequestedFloor <= 20)
                        {
                            var BestCol = (ColumnList[1]);
                            return BestCol;
                        }
                        else if (_RequestedFloor > 20 && _RequestedFloor <= 40)
                        {
                            var BestCol = (ColumnList[2]);
                            return BestCol;
                        }
                        else if (_RequestedFloor > 40 && _RequestedFloor <= 60)
                        {
                            var BestCol = (ColumnList[3]);
                            return BestCol;
                        }                    
                        
                    }
                   
                }// finding the column on other floors
                else if (i != 6)
                {
                    if (i > 1 && i <= 20)
                    {
                        var BestCol = (ColumnList[1]);
                        return BestCol;
                    }
                    else if (i > 20 && i <= 40)
                    {
                        var BestCol = (ColumnList[2]);
                        return BestCol;
                    }
                    else if (i > 40 && i <= 60)
                    {
                        var BestCol = (ColumnList[3]);
                        return BestCol;
                    }       
                }
            }return null;//try for the missing statement
            
        }
       
    }
    class Column
    {
        public int Id;
        public int ElevatorsPerColumn;
        public List<Elevator> ElevatorList;

        public Column(int id, int _ElevatorsPerColumn)
        {
            ElevatorList = new List<Elevator>();
            Id = id;
            ElevatorsPerColumn = _ElevatorsPerColumn;

            //creating the elevator list 
            //using the amount of elevators per column
             for (int i = 1; i<=ElevatorsPerColumn; i++)
            {
                ElevatorList.Add( new Elevator(id:i + 1));
                Console.WriteLine(ElevatorList);
            }  
        }
        //Choosing the best elevator on the ground floor
        public Elevator BestElevator(int UserFloor)
        {   Elevator BestElevator = null;

            foreach ( Elevator elevator in ElevatorList)
            {Console.WriteLine(elevator.Id);
                if (UserFloor == 1)
                {
                    if ( elevator.CurrentFloor == 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else if (elevator.RequestedFloor == 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                       
                    }
                    else if (elevator.Direction == "idle" && elevator.CurrentFloor != 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;

                    }
                }
                else if (UserFloor > 1)
                {
                    if(elevator.Direction == "Down" && UserFloor < elevator.CurrentFloor)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else if (elevator.Direction == "idle" && elevator.CurrentFloor != 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else if(elevator.Direction == "idle" && elevator.CurrentFloor == 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else 
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                }
                else if (UserFloor < 1)
                {
                    if(elevator.Direction == "Up" && UserFloor > elevator.CurrentFloor)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else if (elevator.Direction == "idle" && elevator.CurrentFloor != 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else if(elevator.Direction == "idle" && elevator.CurrentFloor == 1)
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                    else 
                    {
                        BestElevator = elevator;
                        Move(elevator, UserFloor);
                        return BestElevator;
                        
                    }
                
                
                }   }return null;
                
        }
        public void Move(Elevator elevator, int UserFloor)
        {   if (UserFloor < elevator.CurrentFloor){
                while (elevator.CurrentFloor > UserFloor )
                {
                    elevator.CurrentFloor -= 1;
                    Console.WriteLine(elevator.CurrentFloor);
                }
                OpenDoor();
            }
            else {
                 while (elevator.CurrentFloor < UserFloor)
            {
                elevator.CurrentFloor ++;
                Console.WriteLine(elevator.CurrentFloor);
            }
            OpenDoor();
            }
        }
        
        public void OpenDoor()
        {
            Console.WriteLine("Opening Doors");
            Console.WriteLine("Closing Doors");
        }
    }
    
    class Elevator
    {
        public int Id;
        public string Direction;
        public int CurrentFloor;
        public int RequestedFloor;

        public Elevator(int id)
        {
            Id = id;
        }

    }
    class FloorButtons
    {
        public string Direction;
        public int CurrentFloor;


        public FloorButtons(string _Direction, int _CurrentFloor)
        {
            Direction = _Direction;
            CurrentFloor = _CurrentFloor;
        }
    }
    class FloorButtonsGf
    {
        public string Direction;
        public int RequestedFloor;

        public FloorButtonsGf(string _Direction, int _RequestedFloor)
        {
            Direction = _Direction;
            RequestedFloor = _RequestedFloor;
        }
    }
}
