// 60 floors (including G)
// 6 basements
// 4 columns
// 20 elevators

// 90 call buttons
// 345 elevator buttons

using System;

namespace Controller
{
    class Program
    {
        static void ok(string[] args)
        {
            Console.WriteLine("\nwhat's your name?");
            var name = Console.ReadLine();
            var date = DateTime.Now;
            Console.WriteLine($"\nhello, {name}! \n({date:d} at {date:t})");
            Console.Write("\nPress any key to exit...");
            Console.ReadKey(true);
        }
    }

    class Battery
    {
        public string id;
        public string status;

        public Battery(string battId, string battStatus)
        {
            id = battId;
            status = battStatus;
            
        }

    class Column
    {
        public string id;
        public string status;
        public int minFloor;
        public int maxFloor;
        public int elevNum;

        public Column(string colId, string colStatus, int colMinFloor, int colMaxFloor, int colElevNum)
        {
            id = colId;
            status = colStatus;
            minFloor = colMinFloor;
            maxFloor = colMaxFloor;
            elevNum = colElevNum;
        }

    class Elevator 
    {
        public string id;
        public int floor;
        public string status;
        public int minFloor;
        public int maxFloor;
        public string door;
        public int floorDisplay;

        public Elevator(string elevId, int elevFloor, string elevStatus, int elevMinFloor, int elevMaxFloor, string elevDoor)
        {
            id = elevId;
            floor = elevFloor;
            status = elevStatus;
            minFloor = elevMinFloor;    // excluding ground floor
            maxFloor = elevMaxFloor;    // excluding ground floor
            door = elevDoor;
            floorDisplay = elevFloor;

            if (floor == 0)
            {
                floorDisplay = -1;


            } else if (floor == -1)
            {
                floorDisplay = -2;


            } else if (floor == -2)
            {
                floorDisplay = -3;

                
            } else if (floor == -3)
            {
                floorDisplay = -4;

                
            } else if (floor == -4)
            {
                floorDisplay = -5;

            
            } else if (floor == -5)
            {
                floorDisplay = -6;

                
            };
        }

        static void Main(string[] args)
        {
            Battery A = new Battery("1", "idle");

            Column a = new Column("a", "idle", -5, 0, 5);
            Column b = new Column("b", "idle", 1, 20, 5);
            Column c = new Column("c", "idle", 21, 40, 5);
            Column d = new Column("d", "idle", 41, 60, 5);

            Elevator a1 = new Elevator("a1", -5, "idle", -5, 1, "closed");
            Elevator a2 = new Elevator("a2", 1, "idle", -5, 1, "closed");
            Elevator a3 = new Elevator("a3", 1, "idle", -5, 1, "closed");
            Elevator a4 = new Elevator("a4", 1, "idle", -5, 1, "closed");
            Elevator a5 = new Elevator("a5", 1, "idle", -5, 1, "closed");

            Elevator b1 = new Elevator("b1", 1, "goingUp", 1, 20, "closed");
            Elevator b2 = new Elevator("b2", 1, "idle", 1, 20, "closed");
            Elevator b3 = new Elevator("b3", 1, "idle", 1, 20, "closed");
            Elevator b4 = new Elevator("b4", 1, "idle", 1, 20, "closed");
            Elevator b5 = new Elevator("b5", 1, "idle", 1, 20, "closed");

            Elevator c1 = new Elevator("c1", 1, "idle", 21, 40, "closed");
            Elevator c2 = new Elevator("c2", 1, "idle", 21, 40, "closed");
            Elevator c3 = new Elevator("c3", 1, "idle", 21, 40, "closed");
            Elevator c4 = new Elevator("c4", 1, "idle", 21, 40, "closed");
            Elevator c5 = new Elevator("c5", 1, "idle", 21, 40, "closed");
            
            Elevator d1 = new Elevator("d1", 1, "idle", 41, 60, "closed");
            Elevator d2 = new Elevator("d2", 1, "idle", 41, 60, "closed");
            Elevator d3 = new Elevator("d3", 1, "idle", 41, 60, "closed");
            Elevator d4 = new Elevator("d4", 1, "idle", 41, 60, "closed");
            Elevator d5 = new Elevator("d5", 1, "idle", 41, 60, "closed");
            
            void status()
            {

                if (a1.status != "idle" | a2.status != "idle" | a3.status != "idle" | a4.status != "idle" | a5.status != "idle")
                {
                    a.status = "active";
                }

                if (b1.status != "idle" | b2.status != "idle" | b3.status != "idle" | b4.status != "idle" | b5.status != "idle")
                {
                    b.status = "active";
                };

                if (c1.status != "idle" | c2.status != "idle" | c3.status != "idle" | c4.status != "idle" | c5.status != "idle")
                {
                    c.status = "active";
                };

                if (d1.status != "idle" | d2.status != "idle" | d3.status != "idle" | d4.status != "idle" | d5.status != "idle")
                {
                    d.status = "active";
                };


                if (a1.status != "idle" | a2.status != "idle" | a3.status != "idle" | a4.status != "idle" | a5.status != "idle")
                {
                    a.status = "active";
                };
                

                if (a.status != "idle" | b.status != "idle" | c.status != "idle" | d.status != "idle")
                {
                    A.status = "active";
                };

            }


            void requestElev(int userFloor, string direction)
            {
                if (direction == "up") {

                    if (userFloor == b1.floor && b1.status == "goingUp") {

                        Console.WriteLine("elevator b1");
                        b1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + b1.door);
                        b1.door = "closed";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(3));
                        Console.WriteLine("door: " + b1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"floor display: {b1.floorDisplay}");

                        while (b1.floor < 5) {
                            b1.floor++;
                            b1.floorDisplay++;
                            System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                            Console.WriteLine($"floor display: {b1.floorDisplay}");
                        };

                        b1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine("door: " + b1.door);
                        b1.door = "closed";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(3));
                        Console.WriteLine("door: " + b1.door);
                        b1.status = "idle";
                        status();

                    }

                }
            }

            requestElev(1, "up");
            status();
            // Console.WriteLine(c.minFloor);
            // Console.WriteLine(a1.floorDisplay);
            // Console.WriteLine(a.status);
            // Console.WriteLine(A.status);
        }
    }
    }
    }
}