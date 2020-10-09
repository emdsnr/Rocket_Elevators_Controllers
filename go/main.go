package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"time"
)

type battery struct {
	id     string
	status string
	colNum int
}

type column struct {
	id       string
	status   string
	minFloor int
	maxFloor int
	elevNum  int
}

type elevator struct {
	id       string
	floor    int
	status   string
	minFloor int
	maxFloor int
	door     string
}

func status() {

	A := battery{"1", "idle", 4}

	a := column{"a", "idle", -5, 0, 5}
	b := column{"b", "idle", 1, 20, 5}
	c := column{"c", "idle", 21, 40, 5}
	d := column{"d", "idle", 41, 60, 5}

	a1 := elevator{"a1", 1, "idle", -5, 0, "closed"}
	a2 := elevator{"a2", 1, "idle", -5, 0, "closed"}
	a3 := elevator{"a3", 1, "idle", -5, 0, "closed"}
	a4 := elevator{"a4", 1, "idle", -5, 0, "closed"}
	a5 := elevator{"a5", 1, "idle", -5, 0, "closed"}

	b1 := elevator{"b1", 1, "idle", 1, 20, "closed"}
	b2 := elevator{"b2", 1, "idle", 1, 20, "closed"}
	b3 := elevator{"b3", 1, "idle", 1, 20, "closed"}
	b4 := elevator{"b4", 1, "idle", 1, 20, "closed"}
	b5 := elevator{"b5", 1, "idle", 1, 20, "closed"}

	c1 := elevator{"c1", 1, "idle", 21, 40, "closed"}
	c2 := elevator{"c2", 1, "idle", 21, 40, "closed"}
	c3 := elevator{"c3", 1, "idle", 21, 40, "closed"}
	c4 := elevator{"c4", 1, "idle", 21, 40, "closed"}
	c5 := elevator{"c5", 1, "idle", 21, 40, "closed"}

	d1 := elevator{"d1", 1, "idle", 41, 60, "closed"}
	d2 := elevator{"d2", 1, "idle", 41, 60, "closed"}
	d3 := elevator{"d3", 1, "idle", 41, 60, "closed"}
	d4 := elevator{"d4", 1, "idle", 41, 60, "closed"}
	d5 := elevator{"d5", 1, "idle", 41, 60, "closed"}

	fmt.Println(A.status)
	fmt.Println(a.status)
	fmt.Println(b.status)
	fmt.Println(c.status)
	fmt.Println(d.status)

	b1.floor = 20
	b1.status = "goingDown"
	b2.floor = 3
	b2.status = "goingUp"
	b3.floor = 13
	b3.status = "goingDown"
	b4.floor = 15
	b4.status = "goingDown"
	b5.floor = 6
	b5.status = "goingDown"
	// requestElevB(1, "up");

	c1.floor = 1
	c1.status = "goingUp"
	c2.floor = 23
	c2.status = "goingUp"
	c3.floor = 33
	c3.status = "goingDown"
	c4.floor = 40
	c4.status = "goingDown"
	c5.floor = 39
	c5.status = "goingDown"
	// requestElevC(1, "up");

	d1.floor = 58
	d1.status = "goingDown"
	d2.floor = 50
	d2.status = "goingUp"
	d3.floor = 46
	d3.status = "goingUp"
	d4.floor = 1
	d4.status = "goingUp"
	d5.floor = 60
	d5.status = "goingDown"
	// requestElevD(54, "down");

	a1.floor = -3
	a1.status = "idle"
	a2.floor = 1
	a2.status = "idle"
	a3.floor = -2
	a3.status = "goingDown"
	a4.floor = -5
	a4.status = "goingUp"
	a5.floor = 0
	a5.status = "goingDown"
	// requestElevA(-2, "up");

	if a1.status != "idle" || a2.status != "idle" || a3.status != "idle" || a4.status != "idle" || a5.status != "idle" {
		a.status = "active"

	} else {
		a.status = "idle"

	}

	if b1.status != "idle" || b2.status != "idle" || b3.status != "idle" || b4.status != "idle" || b5.status != "idle" {
		b.status = "active"

	} else {
		b.status = "idle"

	}

	if c1.status != "idle" || c2.status != "idle" || c3.status != "idle" || c4.status != "idle" || c5.status != "idle" {
		c.status = "active"

	} else {
		c.status = "idle"

	}

	if d1.status != "idle" || d2.status != "idle" || d3.status != "idle" || d4.status != "idle" || d5.status != "idle" {
		d.status = "active"

	} else {
		d.status = "idle"

	}

}


func elevA1() {
	fmt.Printf("elevator a1");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a1.floor);

	for (a1.floor < userFloor) 
	{
	a1.floor++;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a1.floor);
	};

	a1.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a1.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a1.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a1.door);

			for (; a1.floor < requestedFloor ; a1.floor++) {
					
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a1.floor);
	
				};

			a1.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a1.door);
			a1.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a1.door);
			a1.status = "idle";
			status();
		
		};
	};
}

func elevA2() {
	fmt.Printf("elevator a2");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a2.floor);

	for (a2.floor < userFloor) 
	{
	a2.floor++;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a2.floor);
	};

	a2.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a2.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a2.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a2.door);

			for (a2.floor < requestedFloor) {
					
				a2.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a2.floor);

			};

			a2.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a2.door);
			a2.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a2.door);
			a2.status = "idle";
			status();
		
		};
	};
}

func elevA3() {
	fmt.Printf("elevator a3");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a3.floor);

	for (a3.floor < userFloor) 
	{
	a3.floor++;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a3.floor);
	};

	a3.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a3.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a3.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a3.door);

			for (a3.floor < requestedFloor) {
					
				a3.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a3.floor);

			};

			a3.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a3.door);
			a3.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a3.door);
			a3.status = "idle";
			status();
		
		};
	};
}

func elevA4() {
	fmt.Printf("elevator a4");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a4.floor);

	for (a4.floor < userFloor) 
	{
	a4.floor++;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a4.floor);
	};

	a4.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a4.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a4.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a4.door);

			for (a4.floor < requestedFloor) {
					
				a4.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a4.floor);

			};

			a4.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a4.door);
			a4.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a4.door);
			a4.status = "idle";
			status();
		
		};
	};
}

func elevA5() {
	fmt.Printf("elevator a5");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: {a5.floor}");

	for (a5.floor < userFloor) 
	{
	a5.floor++;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: {a5.floor}");
	};

	a5.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a5.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a5.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a5.door);

			for (a5.floor < requestedFloor) {
					
				a5.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a5.floor);

			};

			a5.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a5.door);
			a5.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a5.door);
			a5.status = "idle";
			status();
		
		};
	};
}


func elevA1v2() {
	fmt.Printf("elevator a1");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a1.floor);

	for (a1.floor > userFloor) 
	{
	a1.floor--;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a1.floor);
	};

	a1.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a1.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a1.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a1.door);

			for (a1.floor < requestedFloor) {
					
				a1.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a1.floor);

			};

			a1.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a1.door);
			a1.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a1.door);
			a1.status = "idle";
			status();
		
		};
	};
}

func elevA2v2() {
	fmt.Printf("elevator a2");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a2.floor);

	for (a2.floor > userFloor) 
	{
	a2.floor--;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a2.floor);
	};

	a2.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a2.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a2.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a2.door);

			for (a2.floor < requestedFloor) {
					
				a2.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a2.floor);

			};

			a2.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a2.door);
			a2.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a2.door);
			a2.status = "idle";
			status();
		
		};
	};
}

func elevA3v2() {
	fmt.Printf("elevator a3");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a3.floor);

	for (a3.floor > userFloor) 
	{
	a3.floor--;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a3.floor);
	};

	a3.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a3.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a3.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a3.door);

			for (a3.floor < requestedFloor) {
					
				a3.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a3.floor);

			};

			a3.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a3.door);
			a3.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a3.door);
			a3.status = "idle";
			status();
		
		};
	};
}

func elevA4v2() {
	fmt.Printf("elevator a4");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: %d\n", a4.floor);

	for (a4.floor > userFloor) 
	{
	a4.floor--;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: %d\n", a4.floor);
	};

	a4.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a4.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a4.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a4.door);

			for (a4.floor < requestedFloor) {
					
				a4.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a4.floor);

			};

			a4.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a4.door);
			a4.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a4.door);
			a4.status = "idle";
			status();
		
		};
	};
}

func elevA5v2() {
	fmt.Printf("elevator a5");
	time.Sleep((1/2) * time.Second);
	fmt.Printf("elevator's floor: {a5.floor}");

	for (a5.floor > userFloor) 
	{
	a5.floor--;
	time.Sleep(time.Second);
	fmt.Printf("elevator's floor: {a5.floor}");
	};

	a5.door = "opened";
	time.Sleep((1/2) * time.Second);
	fmt.Printf("door: " + a5.door);
	time.Sleep((1/2) * time.Second);
	fmt.Printf("which floor would u like to go to?");
	
	
	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if (requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor) {

			time.Sleep((1/2) * time.Second);
			fmt.Printf("please select a valid floor");

		} else {
			
			input = false;

			a5.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a5.door);

			for (a5.floor < requestedFloor) {
					
				a5.floor++;
				time.Sleep(time.Second);
				fmt.Printf("floor display: %d\n", a5.floor);

			};

			a5.door = "opened";
			time.Sleep((1/2) * time.Second);
			fmt.Printf("door: " + a5.door);
			a5.door = "closed";
			time.Sleep(time.Second);
			fmt.Printf("door: " + a5.door);
			a5.status = "idle";
			status();
		
		};
	};
}



func requestElevA(userFloor int, direction string) {

	if direction == "up" && userFloor >= -5 && userFloor < 1 {
		fmt.Println("hello xd")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("which floor would u like to go to? ")

	for input := true; input == true; {

		scanner.Scan()
		requestedFloor, _ := strconv.Atoi(scanner.Text())

		if requestedFloor <= -5 || requestedFloor <= 5 {
			time.Sleep((1 / 2) * time.Second)
			fmt.Printf("please select a valid floor \n")

		} else {

			input = false

			fmt.Printf("door: ")

		}
	}

}
