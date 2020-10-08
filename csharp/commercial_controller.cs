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

            Elevator a1 = new Elevator("a1", -3, "idle", -5, 0, "closed");
            Elevator a2 = new Elevator("a2", 1, "idle", -5, 0, "closed");
            Elevator a3 = new Elevator("a3", -2, "goingDown", -5, 0, "closed");
            Elevator a4 = new Elevator("a4", -5, "goingUp", -5, 0, "closed");
            Elevator a5 = new Elevator("a5", 0, "goingDown", -5, 0, "closed");

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

                } else 
                {
                    a.status = "idle";
                };

                if (b1.status != "idle" | b2.status != "idle" | b3.status != "idle" | b4.status != "idle" | b5.status != "idle")
                {
                    b.status = "active";
                } else 
                {
                    b.status = "idle";
                };

                if (c1.status != "idle" | c2.status != "idle" | c3.status != "idle" | c4.status != "idle" | c5.status != "idle")
                {
                    c.status = "active";
                } else 
                {
                    c.status = "idle";
                };

                if (d1.status != "idle" | d2.status != "idle" | d3.status != "idle" | d4.status != "idle" | d5.status != "idle")
                {
                    d.status = "active";
                } else 
                {
                    d.status = "idle";
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


            void requestElevA(int userFloor, string direction)
            {     
                if (direction == "up") 
                {
                    void elevA1()
                    {
                        Console.WriteLine("elevator a1");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                
                        while (a1.floor < userFloor) 
                        {
                        a1.floor++;
                        a1.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                        };

                        a1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a1.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a1.door);

                                while (a1.floor < requestedFloor) 
                                {
                                a1.floor++;
                                a1.floorDisplay++;
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
                            
                            };
                        };
                    }

                    void elevA2()
                    {
                        Console.WriteLine("elevator a2");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                
                        while (a2.floor < userFloor) 
                        {
                        a2.floor++;
                        a2.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                        };

                        a2.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a2.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);

                                while (a2.floor < requestedFloor) 
                                {
                                a2.floor++;
                                a2.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a2.floorDisplay}");
                                };

                                a2.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a2.door);
                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);
                                a2.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA3()
                    {
                        Console.WriteLine("elevator a3");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                
                        while (a3.floor < userFloor) 
                        {
                        a3.floor++;
                        a3.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                        };

                        a3.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a3.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);

                                while (a3.floor < requestedFloor) 
                                {
                                a3.floor++;
                                a3.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a3.floorDisplay}");
                                };

                                a3.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a3.door);
                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);
                                a3.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA4()
                    {
                        Console.WriteLine("elevator a4");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                
                        while (a4.floor < userFloor) 
                        {
                        a4.floor++;
                        a4.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                        };

                        a4.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a4.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);

                                while (a4.floor < requestedFloor) 
                                {
                                a4.floor++;
                                a4.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a4.floorDisplay}");
                                };

                                a4.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a4.door);
                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);
                                a4.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA5()
                    {
                        Console.WriteLine("elevator a5");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                
                        while (a5.floor < userFloor) 
                        {
                        a5.floor++;
                        a5.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                        };

                        a5.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a5.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);

                                while (a5.floor < requestedFloor) 
                                {
                                a5.floor++;
                                a5.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a5.floorDisplay}");
                                };

                                a5.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a5.door);
                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);
                                a5.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA1v2()
                    {
                        Console.WriteLine("elevator a1");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                
                        while (a1.floor > userFloor) 
                        {
                        a1.floor--;
                        a1.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                        };

                        a1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a1.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a1.door);

                                while (a1.floor < requestedFloor) 
                                {
                                a1.floor++;
                                a1.floorDisplay++;
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
                            
                            };
                        };
                    }

                    void elevA2v2()
                    {
                        Console.WriteLine("elevator a2");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                
                        while (a2.floor > userFloor) 
                        {
                        a2.floor--;
                        a2.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                        };

                        a2.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a2.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);

                                while (a2.floor < requestedFloor) 
                                {
                                a2.floor++;
                                a2.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a2.floorDisplay}");
                                };

                                a2.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a2.door);
                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);
                                a2.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA3v2()
                    {
                        Console.WriteLine("elevator a3");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                
                        while (a3.floor > userFloor) 
                        {
                        a3.floor--;
                        a3.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                        };

                        a3.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a3.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);

                                while (a3.floor < requestedFloor) 
                                {
                                a3.floor++;
                                a3.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a3.floorDisplay}");
                                };

                                a3.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a3.door);
                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);
                                a3.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA4v2()
                    {
                        Console.WriteLine("elevator a4");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                
                        while (a4.floor > userFloor) 
                        {
                        a4.floor--;
                        a4.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                        };

                        a4.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a4.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);

                                while (a4.floor < requestedFloor) 
                                {
                                a4.floor++;
                                a4.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a4.floorDisplay}");
                                };

                                a4.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a4.door);
                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);
                                a4.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA5v2()
                    {
                        Console.WriteLine("elevator a5");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                
                        while (a5.floor > userFloor) 
                        {
                        a5.floor--;
                        a5.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                        };

                        a5.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a5.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);

                                while (a5.floor < requestedFloor) 
                                {
                                a5.floor++;
                                a5.floorDisplay++;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a5.floorDisplay}");
                                };

                                a5.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a5.door);
                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);
                                a5.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    // userfloor == elevFloor && status == goingUp
                    if (userFloor == a1.floor && a1.status == "goingUp") 
                    {
                        elevA1();

                    } else if (userFloor == a2.floor && a2.status == "goingUp") 
                    {
                        elevA2();

                    } else if (userFloor == a3.floor && a3.status == "goingUp") 
                    {
                        elevA3();

                    } else if (userFloor == a4.floor && a4.status == "goingUp") 
                    {
                        elevA4();

                    } else if (userFloor == a5.floor && a5.status == "goingUp") 
                    {
                        elevA5();
                        
                        
                    // userFloor > elevFloor && status == goingUp | 5 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v <= w && v <= x && v <= y && v <= z)
                        {
                            elevA1();

                        } else if (w <= v && w <= x && w <= y && w <= z)
                        {
                            elevA2();

                        } else if (x <= v && x <= w && x <= y && x <= z)
                        {
                            elevA3();

                        } else if (y <= v && y <= w && y <= x && y <= z)
                        {
                            elevA4();

                        } else if (z <= v && z <= w && z <= x && z <= y)
                        {
                            elevA5();

                        } else
                        {
                            elevA1();

                        };
                        

                    // userFloor > elevFloor && status == goingUp | 4 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v > 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= x && v <= y) 
                            {
                                elevA1();

                            } else if (w <= v && w <= x && w <= y) 
                            {
                                elevA2();

                            } else if (x <= v && x <= w && x <= y)
                            {
                                elevA3();

                            } else if (y <= v && y <= w && y <= x)
                            {
                                elevA4();

                            } else 
                            {
                                elevA1();
                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= x && v <= z) 
                            {
                                elevA1();

                            } else if (w <= v && w <= x && w <= z) 
                            {
                                elevA2();

                            } else if (x <= v && x <= w && x <= z)
                            {
                                elevA3();

                            } else if (z <= v && z <= w && z <= x)
                            {
                                elevA5();

                            } else 
                            {
                                elevA1();

                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= w && v <= y && v <= z) 
                            {
                                elevA1();

                            } else if (w <= v && w <= y && w <= z) 
                            {
                                elevA2();

                            } else if (y <= v && y <= w && y <= z)
                            {
                                elevA4();

                            } else if (z <= v && z <= w && z <= y)
                            {
                                elevA5();

                            } else 
                            {
                                elevA1();

                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v <= x && v <= y && v <= z) 
                            {
                                elevA1();

                            } else if (x <= v && x <= y && x <= z) 
                            {
                                elevA3();

                            } else if (y <= v && y <= x && y <= z)
                            {
                                elevA4();

                            } else if (z <= v && z <= x && z <= y)
                            {
                                elevA5();

                            } else 
                            {
                                elevA1();

                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (w <= x && w <= y && w <= z) 
                            {
                                elevA2();

                            } else if (x <= w && x <= y && x <= z) 
                            {
                                elevA3();

                            } else if (y <= w && y <= x && y <= z)
                            {
                                elevA4();

                            } else if (z <= w && z <= x && z <= y)
                            {
                                elevA5();

                            } else 
                            {
                                elevA2();

                            };
                        };
                        

                    // userFloor > elevFloor && status == goingUp | 3 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= w && v <= x)
                            {
                                elevA1();

                            } else if (w <= v && w <= x)
                            {
                                elevA2();

                            } else if (x <= v && x <= w)
                            {
                                elevA3();

                            } else
                            {
                                elevA1();

                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= y)
                            {
                                elevA1();

                            } else if (w <= v && w <= y)
                            {
                                elevA2();

                            } else if (y <= v && y <= w)
                            {
                                elevA4();

                            } else
                            {
                                elevA1();

                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= z)
                            {
                                elevA1();

                            } else if (w <= v && w <= z)
                            {
                                elevA2();

                            } else if (z <= v && z <= w)
                            {
                                elevA5();

                            } else
                            {
                                elevA1();

                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= x && v <= y)
                            {
                                elevA1();

                            } else if (x <= v && x <= y)
                            {
                                elevA3();

                            } else if (y <= v && y <= x)
                            {
                                elevA4();

                            } else
                            {
                                elevA1();

                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= x && v <= z)
                            {
                                elevA1();

                            } else if (x <= v && x <= z)
                            {
                                elevA3();

                            } else if (z <= v && z <= x)
                            {
                                elevA5();

                            } else
                            {
                                elevA1();

                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= y && v <= z)
                            {
                                elevA1();

                            } else if (y <= v && y <= z)
                            {
                                elevA4();

                            } else if (z <= v && z <= y)
                            {
                                elevA5();

                            } else
                            {
                                elevA1();

                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w <= x && w <= y)
                            {
                                elevA2();

                            } else if (x <= w && x <= y)
                            {
                                elevA3();

                            } else if (y <= w && y <= x)
                            {
                                elevA4();

                            } else
                            {
                                elevA2();

                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w <= x && w <= z)
                            {
                                elevA2();

                            } else if (x <= w && x <= z)
                            {
                                elevA3();

                            } else if (z <= w && z <= x)
                            {
                                elevA4();

                            } else
                            {
                                elevA2();

                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w <= y && w <= z)
                            {
                                elevA2();

                            } else if (y <= w && y <= z)
                            {
                                elevA4();

                            } else if (z <= w && z <= y)
                            {
                                elevA5();

                            } else
                            {
                                elevA2();

                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (x <= y && x <= z)
                            {
                                elevA3();

                            } else if (y <= x && y <= z)
                            {
                                elevA4();

                            } else if (z <= x && z <= y)
                            {
                                elevA5();

                            } else
                            {
                                elevA3();

                            };
                        };
                    

                    // userFloor > elevFloor && status == goingUp | 2 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v <= w)
                            {
                                elevA1();

                            } else if (w <= v)
                            {
                                elevA2();

                            } else
                            {
                                elevA1();

                            };

                        // [2] 2 elevs — v, x
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= x)
                            {
                                elevA1();

                            } else if (x <= v)
                            {
                                elevA3();

                            } else
                            {
                                elevA1();

                            };

                        // [3] 2 elevs — v, y
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= y)
                            {
                                elevA1();

                            } else if (y <= v)
                            {
                                elevA4();

                            } else
                            {
                                elevA1();

                            };

                        // [4] 2 elevs — v, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= z)
                            {
                                elevA1();

                            } else if (z <= v)
                            {
                                elevA5();

                            } else
                            {
                                elevA1();

                            };

                        // [5] 2 elevs — w, x
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w <= x)
                            {
                                elevA2();

                            } else if (x <= w)
                            {
                                elevA3();

                            } else
                            {
                                elevA2();

                            };

                        // [6] 2 elevs — w, y
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w <= y)
                            {
                                elevA2();

                            } else if (y <= w)
                            {
                                elevA4();

                            } else
                            {
                                elevA2();

                            };

                        // [7] 2 elevs — w, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w <= z)
                            {
                                elevA2();

                            } else if (z <= w)
                            {
                                elevA5();

                            } else
                            {
                                elevA2();

                            };

                        // [8] 2 elevs — x, y
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (x <= y)
                            {
                                elevA3();

                            } else if (y <= x)
                            {
                                elevA4();

                            } else
                            {
                                elevA3();

                            };

                        // [9] 2 elevs — x, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (x <= z)
                            {
                                elevA3();

                            } else if (z <= x)
                            {
                                elevA5();

                            } else
                            {
                                elevA3();

                            };

                        // [10] 2 elevs — y, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (y <= z)
                            {
                                elevA4();

                            } else if (z <= y)
                            {
                                elevA5();

                            } else
                            {
                                elevA4();

                            };
                        };
                        
                    // userFloor > elevFloor && status == goingUp | 1 OPTION
                    } else if (userFloor > a1.floor && a1.status == "goingUp") 
                    {
                        elevA1();
                        
                    } else if (userFloor > a2.floor && a2.status == "goingUp") 
                    {
                        elevA2();
                        
                    } else if (userFloor > a3.floor && a3.status == "goingUp") 
                    {
                        elevA3();
                        
                    } else if (userFloor > a4.floor && a4.status == "goingUp") 
                    {
                        elevA4();
                        
                    } else if (userFloor > a5.floor && a5.status == "goingUp") 
                    {
                        elevA5();
                    

                    // userFloor == elevFloor && status == idle | a1
                    } else if (userFloor == a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingUp";
                        status();
                        elevA1();
                    
                    // userFloor == elevFloor && status == idle | a2
                    } else if (userFloor == a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingUp";
                        status();
                        elevA2();
                    
                    // userFloor == elevFloor && status == idle | a3
                    } else if (userFloor == a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingUp";
                        status();
                        elevA3();
                    
                    // userFloor == elevFloor && status == idle | a4
                    } else if (userFloor == a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingUp";
                        status();
                        elevA4();
                    
                    // userFloor == elevFloor && status == idle | a5
                    } else if (userFloor == a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingUp";
                        status();
                        elevA5();
                    

                    // userFloor > elevFloor && status == idle | 5 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v <= w && v <= x && v <= y && v <= z)
                        {
                            a1.status = "goingUp";
                            status();
                            elevA1();

                        } else if (w <= v && w <= x && w <= y && w <= z)
                        {
                            a2.status = "goingUp";
                            status();
                            elevA2();

                        } else if (x <= v && x <= w && x <= y && x <= z)
                        {
                            a3.status = "goingUp";
                            status();
                            elevA3();

                        } else if (y <= v && y <= w && y <= x && y <= z)
                        {
                            a4.status = "goingUp";
                            status();
                            elevA4();

                        } else if (z <= v && z <= w && z <= x && z <= y)
                        {
                            a5.status = "goingUp";
                            status();
                            elevA5();

                        } else
                        {
                            a1.status = "goingUp";
                            status();
                            elevA1();

                        };

                        
                    // userFloor > elevFloor && status == idle | 4 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v > 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= x && v <= y) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= x && w <= y) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= v && x <= w && x <= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= v && y <= w && y <= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= x && v <= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= x && w <= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= v && x <= w && x <= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (z <= v && z <= w && z <= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= w && v <= y && v <= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= y && w <= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (y <= v && y <= w && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= v && z <= w && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v <= x && v <= y && v <= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (x <= v && x <= y && x <= z) 
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= v && y <= x && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= v && z <= x && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (w <= x && w <= y && w <= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= w && x <= y && x <= z) 
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= w && y <= x && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= w && z <= x && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };
                        };
                        

                    // userFloor > elevFloor && status == idle | 3 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= w && v <= x)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= x)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= v && x <= w)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (y <= v && y <= w)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v && w <= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (z <= v && z <= w)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= x && v <= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (x <= v && x <= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= v && y <= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= x && v <= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (x <= v && x <= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (z <= v && z <= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= y && v <= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (y <= v && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= v && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w <= x && w <= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= w && x <= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= w && y <= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w <= x && w <= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= w && x <= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (z <= w && z <= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w <= y && w <= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (y <= w && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= w && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (x <= y && x <= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= x && y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= x && z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            };
                        };
                    

                    // userFloor > elevFloor && status == idle | 2 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v <= w)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (w <= v)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [2] 2 elevs — v, x
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= x)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (x <= v)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [3] 2 elevs — v, y
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (y <= v)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [4] 2 elevs — v, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            } else if (z <= v)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1();
                                
                            };

                        // [5] 2 elevs — w, x
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w <= x)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (x <= w)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [6] 2 elevs — w, y
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w <= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (y <= w)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [7] 2 elevs — w, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w <= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            } else if (z <= w)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2();
                                
                            };

                        // [8] 2 elevs — x, y
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (x <= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (y <= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            };

                        // [9] 2 elevs — x, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (x <= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            } else if (z <= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3();
                                
                            };

                        // [10] 2 elevs — y, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (y <= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            } else if (z <= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4();
                                
                            };
                        };
                        
                    // userFloor > elevFloor && status == idle | 1 OPTION
                    } else if (userFloor > a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingUp";
                        status();
                        elevA1();
                        
                    } else if (userFloor > a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingUp";
                        status();
                        elevA2();
                        
                    } else if (userFloor > a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingUp";
                        status();
                        elevA3();
                        
                    } else if (userFloor > a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingUp";
                        status();
                        elevA4();
                        
                    } else if (userFloor > a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingUp";
                        status();
                        elevA5();


                    // userFloor < elevFloor && status == idle | 5 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v >= w && v >= x && v >= y && v >= z)
                        {
                            a1.status = "goingUp";
                            status();
                            elevA1v2();

                        } else if (w >= v && w >= x && w >= y && w >= z)
                        {
                            a2.status = "goingUp";
                            status();
                            elevA2v2();

                        } else if (x >= v && x >= w && x >= y && x >= z)
                        {
                            a3.status = "goingUp";
                            status();
                            elevA3v2();

                        } else if (y >= v && y >= w && y >= x && y >= z)
                        {
                            a4.status = "goingUp";
                            status();
                            elevA4v2();

                        } else if (z >= v && z >= w && z >= x && z >= y)
                        {
                            a5.status = "goingUp";
                            status();
                            elevA5v2();

                        } else
                        {
                            a1.status = "goingUp";
                            status();
                            elevA1v2();

                        };

                        
                    // userFloor < elevFloor && status == idle | 4 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v < 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= x && v >= y) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= x && w >= y) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= v && x >= w && x >= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= v && y >= w && y >= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= x && v >= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= x && w >= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= v && x >= w && x >= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (z >= v && z >= w && z >= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= w && v >= y && v >= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= y && w >= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (y >= v && y >= w && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= v && z >= w && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v >= x && v >= y && v >= z) 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (x >= v && x >= y && x >= z) 
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= v && y >= x && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= v && z >= x && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (w >= x && w >= y && w >= z) 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= w && x >= y && x >= z) 
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= w && y >= x && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= w && z >= x && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };
                        };
                        

                    // userFloor < elevFloor && status == idle | 3 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= w && v >= x)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= x)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= v && x >= w)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (y >= v && y >= w)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v && w >= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (z >= v && z >= w)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= x && v >= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (x >= v && x >= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= v && y >= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= x && v >= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (x >= v && x >= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (z >= v && z >= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= y && v >= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (y >= v && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= v && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w >= x && w >= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= w && x >= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= w && y >= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w >= x && w >= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= w && x >= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (z >= w && z >= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w >= y && w >= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (y >= w && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= w && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (x >= y && x >= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= x && y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= x && z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            };
                        };
                    

                    // userFloor < elevFloor && status == idle | 2 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v >= w)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (w >= v)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [2] 2 elevs — v, x
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= x)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (x >= v)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 2 elevs — v, y
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= y)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (y >= v)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 2 elevs — v, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= z)
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            } else if (z >= v)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingUp";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 2 elevs — w, x
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w >= x)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (x >= w)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [6] 2 elevs — w, y
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w >= y)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (y >= w)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [7] 2 elevs — w, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w >= z)
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            } else if (z >= w)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a2.status = "goingUp";
                                status();
                                elevA2v2();
                                
                            };

                        // [8] 2 elevs — x, y
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (x >= y)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (y >= x)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            };

                        // [9] 2 elevs — x, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (x >= z)
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            } else if (z >= x)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a3.status = "goingUp";
                                status();
                                elevA3v2();
                                
                            };

                        // [10] 2 elevs — y, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (y >= z)
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            } else if (z >= y)
                            {
                                a5.status = "goingUp";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a4.status = "goingUp";
                                status();
                                elevA4v2();
                                
                            };
                        };
                        
                    // userFloor < elevFloor && status == idle | 1 OPTION
                    } else if (userFloor < a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingUp";
                        status();
                        elevA1v2();
                        
                    } else if (userFloor < a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingUp";
                        status();
                        elevA2v2();
                        
                    } else if (userFloor < a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingUp";
                        status();
                        elevA3v2();
                        
                    } else if (userFloor < a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingUp";
                        status();
                        elevA4v2();
                        
                    } else if (userFloor < a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingUp";
                        status();
                        elevA5v2();


                    };

                } else if (direction == "down") 
                {
                    void elevA1()
                    {
                        Console.WriteLine("elevator a1");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                
                        while (a1.floor > userFloor) 
                        {
                        a1.floor--;
                        a1.floorDisplay--;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                        };

                        a1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a1.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a1.door);

                                while (a1.floor > requestedFloor) 
                                {
                                a1.floor--;
                                a1.floorDisplay--;
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
                            
                            };
                        };
                    }

                    void elevA2()
                    {
                        Console.WriteLine("elevator a2");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                
                        while (a2.floor < userFloor) 
                        {
                        a2.floor++;
                        a2.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                        };

                        a2.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a2.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);

                                while (a2.floor > requestedFloor) 
                                {
                                a2.floor--;
                                a2.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a2.floorDisplay}");
                                };

                                a2.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a2.door);
                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);
                                a2.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA3()
                    {
                        Console.WriteLine("elevator a3");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                
                        while (a3.floor < userFloor) 
                        {
                        a3.floor++;
                        a3.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                        };

                        a3.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a3.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);

                                while (a3.floor > requestedFloor) 
                                {
                                a3.floor--;
                                a3.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a3.floorDisplay}");
                                };

                                a3.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a3.door);
                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);
                                a3.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA4()
                    {
                        Console.WriteLine("elevator a4");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                
                        while (a4.floor < userFloor) 
                        {
                        a4.floor++;
                        a4.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                        };

                        a4.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a4.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);

                                while (a4.floor > requestedFloor) 
                                {
                                a4.floor--;
                                a4.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a4.floorDisplay}");
                                };

                                a4.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a4.door);
                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);
                                a4.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA5()
                    {
                        Console.WriteLine("elevator a5");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                
                        while (a5.floor < userFloor) 
                        {
                        a5.floor++;
                        a5.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                        };

                        a5.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a5.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);

                                while (a5.floor > requestedFloor) 
                                {
                                a5.floor--;
                                a5.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a5.floorDisplay}");
                                };

                                a5.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a5.door);
                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);
                                a5.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA1v2()
                    {
                        Console.WriteLine("elevator a1");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                
                        while (a1.floor < userFloor) 
                        {
                        a1.floor++;
                        a1.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a1.floorDisplay}");
                        };

                        a1.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a1.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a1.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a1.door);

                                while (a1.floor > requestedFloor) 
                                {
                                a1.floor--;
                                a1.floorDisplay--;
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
                            
                            };
                        };
                    }

                    void elevA2v2()
                    {
                        Console.WriteLine("elevator a2");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                
                        while (a2.floor < userFloor) 
                        {
                        a2.floor++;
                        a2.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a2.floorDisplay}");
                        };

                        a2.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a2.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);

                                while (a2.floor > requestedFloor) 
                                {
                                a2.floor--;
                                a2.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a2.floorDisplay}");
                                };

                                a2.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a2.door);
                                a2.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a2.door);
                                a2.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA3v2()
                    {
                        Console.WriteLine("elevator a3");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                
                        while (a3.floor < userFloor) 
                        {
                        a3.floor++;
                        a3.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a3.floorDisplay}");
                        };

                        a3.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a3.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);

                                while (a3.floor > requestedFloor) 
                                {
                                a3.floor--;
                                a3.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a3.floorDisplay}");
                                };

                                a3.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a3.door);
                                a3.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a3.door);
                                a3.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA4v2()
                    {
                        Console.WriteLine("elevator a4");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                
                        while (a4.floor < userFloor) 
                        {
                        a4.floor++;
                        a4.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a4.floorDisplay}");
                        };

                        a4.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a4.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);

                                while (a4.floor > requestedFloor) 
                                {
                                a4.floor--;
                                a4.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a4.floorDisplay}");
                                };

                                a4.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a4.door);
                                a4.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a4.door);
                                a4.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    void elevA5v2()
                    {
                        Console.WriteLine("elevator a5");
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                
                        while (a5.floor < userFloor) 
                        {
                        a5.floor++;
                        a5.floorDisplay++;
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                        Console.WriteLine($"elevator's floor: {a5.floorDisplay}");
                        };

                        a5.door = "opened";
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("door: " + a5.door);
                        System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                        Console.WriteLine("which floor would u like to go to?");
                        
                        bool input = true;
                        while (input) 
                        {
                            int requestedFloor;
                            bool valid = int.TryParse(Console.ReadLine(), out requestedFloor);

                            if (requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor || valid == false) 
                            {
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("please select a valid floor");

                            } else {
                                
                                input = false;

                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);

                                while (a5.floor > requestedFloor) 
                                {
                                a5.floor--;
                                a5.floorDisplay--;
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine($"floor display: {a5.floorDisplay}");
                                };

                                a5.door = "opened";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(.5));
                                Console.WriteLine("door: " + a5.door);
                                a5.door = "closed";
                                System.Threading.Thread.Sleep(TimeSpan.FromSeconds(1));
                                Console.WriteLine("door: " + a5.door);
                                a5.status = "idle";
                                status();
                            
                            };
                        };
                    }

                    // userfloor == elevFloor && status == goingDown
                    if (userFloor == a1.floor && a1.status == "goingDown") 
                    {
                        elevA1();

                    } else if (userFloor == a2.floor && a2.status == "goingDown") 
                    {
                        elevA2();

                    } else if (userFloor == a3.floor && a3.status == "goingDown") 
                    {
                        elevA3();

                    } else if (userFloor == a4.floor && a4.status == "goingDown") 
                    {
                        elevA4();

                    } else if (userFloor == a5.floor && a5.status == "goingDown") 
                    {
                        elevA5();
                        

                    // userFloor < elevFloor && status == goingDown | 5 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v >= w && v >= x && v >= y && v >= z)
                        {
                            elevA1();

                        } else if (w >= v && w >= x && w >= y && w >= z)
                        {
                            elevA2();

                        } else if (x >= v && x >= w && x >= y && x >= z)
                        {
                            elevA3();

                        } else if (y >= v && y >= w && y >= x && y >= z)
                        {
                            elevA4();

                        } else if (z >= v && z >= w && z >= x && z >= y)
                        {
                            elevA5();

                        } else
                        {
                            elevA1();

                        };

                        
                    // userFloor < elevFloor && status == goingDown | 4 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v < 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= x && v >= y) 
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= x && w >= y) 
                            {
                                elevA2();
                                
                            } else if (x >= v && x >= w && x >= y)
                            {
                                elevA3();
                                
                            } else if (y >= v && y >= w && y >= x)
                            {
                                elevA4();
                                
                            } else 
                            {
                                elevA1();
                                                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= x && v >= z) 
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= x && w >= z) 
                            {
                                elevA2();
                                
                            } else if (x >= v && x >= w && x >= z)
                            {
                                elevA3();
                                
                            } else if (z >= v && z >= w && z >= x)
                            {
                                elevA5();
                                
                            } else 
                            {
                                elevA1();
                                
                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= w && v >= y && v >= z) 
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= y && w >= z) 
                            {
                                elevA2();
                                
                            } else if (y >= v && y >= w && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= v && z >= w && z >= y)
                            {
                                elevA5();
                                
                            } else 
                            {
                                elevA1();
                                
                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v >= x && v >= y && v >= z) 
                            {
                                elevA1();
                                
                            } else if (x >= v && x >= y && x >= z) 
                            {
                                elevA3();
                                
                            } else if (y >= v && y >= x && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= v && z >= x && z >= y)
                            {
                                elevA5();
                                
                            } else 
                            {
                                elevA1();
                                
                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (w >= x && w >= y && w >= z) 
                            {
                                elevA2();
                                
                            } else if (x >= w && x >= y && x >= z) 
                            {
                                elevA3();
                                
                            } else if (y >= w && y >= x && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= w && z >= x && z >= y)
                            {
                                elevA5();
                                
                            } else 
                            {
                                elevA2();
                                
                            };
                        };
                        

                    // userFloor < elevFloor && status == goingDown | 3 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= w && v >= x)
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= x)
                            {
                                elevA2();
                                
                            } else if (x >= v && x >= w)
                            {
                                elevA3();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= y)
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= y)
                            {
                                elevA2();
                                
                            } else if (y >= v && y >= w)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= z)
                            {
                                elevA1();
                                
                            } else if (w >= v && w >= z)
                            {
                                elevA2();
                                
                            } else if (z >= v && z >= w)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= x && v >= y)
                            {
                                elevA1();
                                
                            } else if (x >= v && x >= y)
                            {
                                elevA3();
                                
                            } else if (y >= v && y >= x)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= x && v >= z)
                            {
                                elevA1();
                                
                            } else if (x >= v && x >= z)
                            {
                                elevA3();
                                
                            } else if (z >= v && z >= x)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= y && v >= z)
                            {
                                elevA1();
                                
                            } else if (y >= v && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= v && z >= y)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w >= x && w >= y)
                            {
                                elevA2();
                                
                            } else if (x >= w && x >= y)
                            {
                                elevA3();
                                
                            } else if (y >= w && y >= x)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w >= x && w >= z)
                            {
                                elevA2();
                                
                            } else if (x >= w && x >= z)
                            {
                                elevA3();
                                
                            } else if (z >= w && z >= x)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w >= y && w >= z)
                            {
                                elevA2();
                                
                            } else if (y >= w && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= w && z >= y)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (x >= y && x >= z)
                            {
                                elevA3();
                                
                            } else if (y >= x && y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= x && z >= y)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA3();
                                
                            };
                        };
                    

                    // userFloor < elevFloor && status == goingDown | 2 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v >= w)
                            {
                                elevA1();
                                
                            } else if (w >= v)
                            {
                                elevA2();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [2] 2 elevs — v, x
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= x)
                            {
                                elevA1();
                                
                            } else if (x >= v)
                            {
                                elevA3();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [3] 2 elevs — v, y
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= y)
                            {
                                elevA1();
                                
                            } else if (y >= v)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [4] 2 elevs — v, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= z)
                            {
                                elevA1();
                                
                            } else if (z >= v)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA1();
                                
                            };

                        // [5] 2 elevs — w, x
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w >= x)
                            {
                                elevA2();
                                
                            } else if (x >= w)
                            {
                                elevA3();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [6] 2 elevs — w, y
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w >= y)
                            {
                                elevA2();
                                
                            } else if (y >= w)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [7] 2 elevs — w, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w >= z)
                            {
                                elevA2();
                                
                            } else if (z >= w)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA2();
                                
                            };

                        // [8] 2 elevs — x, y
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (x >= y)
                            {
                                elevA3();
                                
                            } else if (y >= x)
                            {
                                elevA4();
                                
                            } else
                            {
                                elevA3();
                                
                            };

                        // [9] 2 elevs — x, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (x >= z)
                            {
                                elevA3();
                                
                            } else if (z >= x)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA3();
                                
                            };

                        // [10] 2 elevs — y, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (y >= z)
                            {
                                elevA4();
                                
                            } else if (z >= y)
                            {
                                elevA5();
                                
                            } else
                            {
                                elevA4();
                                
                            };
                        };

                        
                    // userFloor < elevFloor && status == goingDown | 1 OPTION
                    } else if (userFloor < a1.floor && a1.status == "goingDown") 
                    {
                        elevA1();
                        
                    } else if (userFloor < a2.floor && a2.status == "goingDown") 
                    {
                        elevA2();
                        
                    } else if (userFloor < a3.floor && a3.status == "goingDown") 
                    {
                        elevA3();
                        
                    } else if (userFloor < a4.floor && a4.status == "goingDown") 
                    {
                        elevA4();
                        
                    } else if (userFloor < a5.floor && a5.status == "goingDown") 
                    {
                        elevA5();
                    

                    // userFloor == elevFloor && status == idle | a1
                    } else if (userFloor == a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingDown";
                        status();
                        elevA1();
                    
                    // userFloor == elevFloor && status == idle | a2
                    } else if (userFloor == a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingDown";
                        status();
                        elevA2();
                    
                    // userFloor == elevFloor && status == idle | a3
                    } else if (userFloor == a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingDown";
                        status();
                        elevA3();
                    
                    // userFloor == elevFloor && status == idle | a4
                    } else if (userFloor == a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingDown";
                        status();
                        elevA4();
                    
                    // userFloor == elevFloor && status == idle | a5
                    } else if (userFloor == a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingDown";
                        status();
                        elevA5();
                    

                    // userFloor < elevFloor && status == idle | 5 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v >= w && v >= x && v >= y && v >= z)
                        {
                            a1.status = "goingDown";
                            status();
                            elevA1();

                        } else if (w >= v && w >= x && w >= y && w >= z)
                        {
                            a2.status = "goingDown";
                            status();
                            elevA2();

                        } else if (x >= v && x >= w && x >= y && x >= z)
                        {
                            a3.status = "goingDown";
                            status();
                            elevA3();

                        } else if (y >= v && y >= w && y >= x && y >= z)
                        {
                            a4.status = "goingDown";
                            status();
                            elevA4();

                        } else if (z >= v && z >= w && z >= x && z >= y)
                        {
                            a5.status = "goingDown";
                            status();
                            elevA5();

                        } else
                        {
                            a1.status = "goingDown";
                            status();
                            elevA1();

                        };

                        
                    // userFloor < elevFloor && status == idle | 4 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v < 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= x && v >= y) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= x && w >= y) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= v && x >= w && x >= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= v && y >= w && y >= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= x && v >= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= x && w >= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= v && x >= w && x >= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (z >= v && z >= w && z >= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= w && v >= y && v >= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= y && w >= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (y >= v && y >= w && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= v && z >= w && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v >= x && v >= y && v >= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (x >= v && x >= y && x >= z) 
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= v && y >= x && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= v && z >= x && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (w >= x && w >= y && w >= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= w && x >= y && x >= z) 
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= w && y >= x && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= w && z >= x && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };
                        };
                        

                    // userFloor < elevFloor && status == idle | 3 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= w && v >= x)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= x)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= v && x >= w)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= w && v >= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (y >= v && y >= w)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= w && v >= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v && w >= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (z >= v && z >= w)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v >= x && v >= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (x >= v && x >= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= v && y >= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v >= x && v >= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (x >= v && x >= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (z >= v && z >= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v >= y && v >= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (y >= v && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= v && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w >= x && w >= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= w && x >= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= w && y >= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w >= x && w >= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= w && x >= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (z >= w && z >= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w >= y && w >= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (y >= w && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= w && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (x >= y && x >= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= x && y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= x && z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            };
                        };
                    

                    // userFloor < elevFloor && status == idle | 2 OPTIONS
                    } else if (userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v >= w)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (w >= v)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [2] 2 elevs — v, x
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v >= x)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (x >= v)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [3] 2 elevs — v, y
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v >= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (y >= v)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [4] 2 elevs — v, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v >= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            } else if (z >= v)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1();
                                
                            };

                        // [5] 2 elevs — w, x
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w >= x)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (x >= w)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [6] 2 elevs — w, y
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w >= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (y >= w)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [7] 2 elevs — w, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w >= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            } else if (z >= w)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2();
                                
                            };

                        // [8] 2 elevs — x, y
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (x >= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (y >= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            };

                        // [9] 2 elevs — x, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (x >= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            } else if (z >= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3();
                                
                            };

                        // [10] 2 elevs — y, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (y >= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            } else if (z >= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5();
                                
                            } else
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4();
                                
                            };
                        };

                        
                    // userFloor < elevFloor && status == idle | 1 OPTION
                    } else if (userFloor < a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingDown";
                        status();
                        elevA1();
                        
                    } else if (userFloor < a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingDown";
                        status();
                        elevA2();
                        
                    } else if (userFloor < a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingDown";
                        status();
                        elevA3();
                        
                    } else if (userFloor < a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingDown";
                        status();
                        elevA4();
                        
                    } else if (userFloor < a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingDown";
                        status();
                        elevA5();


                    // userFloor > elevFloor && status == idle | 5 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        if (v <= w && v <= x && v <= y && v <= z)
                        {
                            a1.status = "goingDown";
                            status();
                            elevA1v2();

                        } else if (w <= v && w <= x && w <= y && w <= z)
                        {
                            a2.status = "goingDown";
                            status();
                            elevA2v2();

                        } else if (x <= v && x <= w && x <= y && x <= z)
                        {
                            a3.status = "goingDown";
                            status();
                            elevA3v2();

                        } else if (y <= v && y <= w && y <= x && y <= z)
                        {
                            a4.status = "goingDown";
                            status();
                            elevA4v2();

                        } else if (z <= v && z <= w && z <= x && z <= y)
                        {
                            a5.status = "goingDown";
                            status();
                            elevA5v2();

                        } else
                        {
                            a1.status = "goingDown";
                            status();
                            elevA1v2();

                        };

                        
                    // userFloor > elevFloor && status == idle | 4 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 4 elevs — v, w, x, y
                        if (v > 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= x && v <= y) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= x && w <= y) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= v && x <= w && x <= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= v && y <= w && y <= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                                                
                            };

                        // [2] 4 elevs — v, w, x, z
                        } else if (v > 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= x && v <= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= x && w <= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= v && x <= w && x <= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (z <= v && z <= w && z <= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 4 elevs — v, w, y, z
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= w && v <= y && v <= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= y && w <= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (y <= v && y <= w && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= v && z <= w && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 4 elevs — v, x, y, z
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (v <= x && v <= y && v <= z) 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (x <= v && x <= y && x <= z) 
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= v && y <= x && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= v && z <= x && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 4 elevs — w, x, y, z
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (w <= x && w <= y && w <= z) 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= w && x <= y && x <= z) 
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= w && y <= x && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= w && z <= x && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else 
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };
                        };
                        

                    // userFloor > elevFloor && status == idle | 3 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 3 elevs — v, w, x
                        if (v > 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= w && v <= x)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= x)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= v && x <= w)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [2] 3 elevs — v, w, y
                        } else if (v > 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= w && v <= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (y <= v && y <= w)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 3 elevs — v, w, z
                        } else if (v > 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= w && v <= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v && w <= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (z <= v && z <= w)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 3 elevs — v, x, y
                        } else if (v > 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (v <= x && v <= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (x <= v && x <= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= v && y <= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 3 elevs — v, x, z
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (v <= x && v <= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (x <= v && x <= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (z <= v && z <= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [6] 3 elevs — v, y, z
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (v <= y && v <= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (y <= v && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= v && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [7] 3 elevs — w, x, y
                        } else if (v < 0 && w > 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (w <= x && w <= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= w && x <= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= w && y <= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [8] 3 elevs — w, x, z
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (w <= x && w <= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= w && x <= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (z <= w && z <= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [9] 3 elevs — w, y, z
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (w <= y && w <= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (y <= w && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= w && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [10] 3 elevs — x, y, z
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z > 0)
                        {
                            if (x <= y && x <= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= x && y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= x && z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            };
                        };
                    

                    // userFloor > elevFloor && status == idle | 2 OPTIONS
                    } else if (userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle") 
                    {
                        int v = (userFloor - a1.floor);
                        int w = (userFloor - a2.floor);
                        int x = (userFloor - a3.floor);
                        int y = (userFloor - a4.floor);
                        int z = (userFloor - a5.floor);
                        // Console.WriteLine(v);
                        // Console.WriteLine(w);
                        // Console.WriteLine(x);
                        // Console.WriteLine(y);
                        // Console.WriteLine(z);

                        // [1] 2 elevs — v, w
                        if (v > 0 && w > 0 && x < 0 && y < 0 && z < 0)
                        {
                            if (v <= w)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (w <= v)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [2] 2 elevs — v, x
                        } else if (v > 0 && w < 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (v <= x)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (x <= v)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [3] 2 elevs — v, y
                        } else if (v > 0 && w < 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (v <= y)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (y <= v)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [4] 2 elevs — v, z
                        } else if (v > 0 && w < 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (v <= z)
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            } else if (z <= v)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a1.status = "goingDown";
                                status();
                                elevA1v2();
                                
                            };

                        // [5] 2 elevs — w, x
                        } else if (v < 0 && w > 0 && x > 0 && y < 0 && z < 0)
                        {
                            if (w <= x)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (x <= w)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [6] 2 elevs — w, y
                        } else if (v < 0 && w > 0 && x < 0 && y > 0 && z < 0)
                        {
                            if (w <= y)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (y <= w)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [7] 2 elevs — w, z
                        } else if (v < 0 && w > 0 && x < 0 && y < 0 && z > 0)
                        {
                            if (w <= z)
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            } else if (z <= w)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a2.status = "goingDown";
                                status();
                                elevA2v2();
                                
                            };

                        // [8] 2 elevs — x, y
                        } else if (v < 0 && w < 0 && x > 0 && y > 0 && z < 0)
                        {
                            if (x <= y)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (y <= x)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            };

                        // [9] 2 elevs — x, z
                        } else if (v < 0 && w < 0 && x > 0 && y < 0 && z > 0)
                        {
                            if (x <= z)
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            } else if (z <= x)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a3.status = "goingDown";
                                status();
                                elevA3v2();
                                
                            };

                        // [10] 2 elevs — y, z
                        } else if (v < 0 && w < 0 && x < 0 && y > 0 && z > 0)
                        {
                            if (y <= z)
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            } else if (z <= y)
                            {
                                a5.status = "goingDown";
                                status();
                                elevA5v2();
                                
                            } else
                            {
                                a4.status = "goingDown";
                                status();
                                elevA4v2();
                                
                            };
                        };
                        
                        
                    // userFloor > elevFloor && status == idle | 1 OPTION
                    } else if (userFloor > a1.floor && a1.status == "idle") 
                    {
                        a1.status = "goingDown";
                        status();
                        elevA1v2();
                        
                    } else if (userFloor > a2.floor && a2.status == "idle") 
                    {
                        a2.status = "goingDown";
                        status();
                        elevA2v2();
                        
                    } else if (userFloor > a3.floor && a3.status == "idle") 
                    {
                        a3.status = "goingDown";
                        status();
                        elevA3v2();
                        
                    } else if (userFloor > a4.floor && a4.status == "idle") 
                    {
                        a4.status = "goingDown";
                        status();
                        elevA4v2();
                        
                    } else if (userFloor > a5.floor && a5.status == "idle") 
                    {
                        a5.status = "goingDown";
                        status();
                        elevA5v2();


                    };

                };
            }

            requestElevA(-2, "up");
            // Console.WriteLine(c.minFloor);
            // Console.WriteLine(a1.floorDisplay);
            // Console.WriteLine(a.status);
            // Console.WriteLine(A.status);
        }
    }
    }
    }
}