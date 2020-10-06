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
        }

        static void Main(string[] args)
        {
            Battery A = new Battery("1", "idle");

            Column a = new Column("a", "idle", -5, 0, 5);
            Column b = new Column("b", "idle", 1, 20, 5);
            Column c = new Column("c", "idle", 21, 40, 5);
            Column d = new Column("d", "idle", 41, 60, 5);

            Elevator a1 = new Elevator("a1", -5, "goingUp", -5, 1, "closed");
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

            void floorDisplayA1()
            {
                if (a1.floor == 1)
                {
                    a1.floorDisplay = 1;


                } else if (a1.floor == 0)
                {
                    a1.floorDisplay = -1;


                } else if (a1.floor == -1)
                {
                    a1.floorDisplay = -2;


                } else if (a1.floor == -2)
                {
                    a1.floorDisplay = -3;

                    
                } else if (a1.floor == -3)
                {
                    a1.floorDisplay = -4;

                    
                } else if (a1.floor == -4)
                {
                    a1.floorDisplay = -5;

                
                } else if (a1.floor == -5)
                {
                    a1.floorDisplay = -6;

                    
                };
            }

            void floorDisplayA2()
            {
                if (a2.floor == 0)
                {
                    a2.floorDisplay = -1;


                } else if (a2.floor == -1)
                {
                    a2.floorDisplay = -2;


                } else if (a2.floor == -2)
                {
                    a2.floorDisplay = -3;

                    
                } else if (a2.floor == -3)
                {
                    a2.floorDisplay = -4;

                    
                } else if (a2.floor == -4)
                {
                    a2.floorDisplay = -5;

                
                } else if (a2.floor == -5)
                {
                    a2.floorDisplay = -6;

                    
                };
            }

            void floorDisplayA3()
            {
                if (a3.floor == 0)
                {
                    a3.floorDisplay = -1;


                } else if (a3.floor == -1)
                {
                    a3.floorDisplay = -2;


                } else if (a3.floor == -2)
                {
                    a3.floorDisplay = -3;

                    
                } else if (a3.floor == -3)
                {
                    a3.floorDisplay = -4;

                    
                } else if (a3.floor == -4)
                {
                    a3.floorDisplay = -5;

                
                } else if (a3.floor == -5)
                {
                    a3.floorDisplay = -6;

                    
                };
            }

            void floorDisplayA4()
            {
                if (a4.floor == 0)
                {
                    a4.floorDisplay = -1;


                } else if (a4.floor == -1)
                {
                    a4.floorDisplay = -2;


                } else if (a4.floor == -2)
                {
                    a4.floorDisplay = -3;

                    
                } else if (a4.floor == -3)
                {
                    a4.floorDisplay = -4;

                    
                } else if (a4.floor == -4)
                {
                    a4.floorDisplay = -5;

                
                } else if (a4.floor == -5)
                {
                    a4.floorDisplay = -6;

                    
                };
            }

            void floorDisplayA5()
            {
                if (a5.floor == 0)
                {
                    a5.floorDisplay = -1;


                } else if (a5.floor == -1)
                {
                    a5.floorDisplay = -2;


                } else if (a5.floor == -2)
                {
                    a5.floorDisplay = -3;

                    
                } else if (a5.floor == -3)
                {
                    a5.floorDisplay = -4;

                    
                } else if (a5.floor == -4)
                {
                    a5.floorDisplay = -5;

                
                } else if (a5.floor == -5)
                {
                    a5.floorDisplay = -6;

                    
                };
            }


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

            status();


            void requestElev(int userFloor, string direction)
            {
                if (direction == "up") {

                    if (userFloor == a1.floor && a1.status == "goingUp") {

                        Console.WriteLine("elevator a1");
                        a1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a1.door);
                        a1.door = "closed";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine("door: " + a1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        floorDisplayA1();
                        Console.WriteLine($"floor display: {a1.floorDisplay}");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");

                        bool invalid = true;
                        while (invalid) {
                            var request = Console.ReadLine();
                            var requestedFloor = Convert.ToInt32(request);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor == userFloor) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                invalid = false;

                                while (a1.floor < requestedFloor) 
                                {
                                a1.floor++;
                                a1.floorDisplay++;
                                floorDisplayA1();
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a1.floorDisplay}");
                                };

                                a1.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a1.door);
                                a1.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a1.door);
                                a1.status = "idle";
                                status();
                                Console.WriteLine(a.status);
                            
                            }
                        }

                    }

                }
            }

            requestElev(-5, "up");
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