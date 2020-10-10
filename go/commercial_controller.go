// 60 floors (including G)
// 6 basements
// 4 columns
// 20 elevators

// 90 call buttons
// 345 elevator buttons

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

func main() {

	A := battery{"1", "idle", 4}
	A.colNum = 4

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

	status := func() {

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

	requestElevA := func(userFloor int, direction string) {
		scanner := bufio.NewScanner(os.Stdin)

		if direction == "up" && userFloor >= -5 && userFloor < 1 {

			elevA1 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a1.floor)

				for a1.floor < userFloor {
					a1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)
				}

				a1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)

						for a1.floor < requestedFloor {

							a1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a1.floor)

						}

						a1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.status = "idle"
						status()

					}
				}
			}

			elevA2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a2.floor)

				for a2.floor < userFloor {
					a2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)
				}

				a2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)

						for a2.floor < requestedFloor {

							a2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a2.floor)

						}

						a2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.status = "idle"
						status()

					}
				}
			}

			elevA3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a3.floor)

				for a3.floor < userFloor {
					a3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)
				}

				a3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)

						for a3.floor < requestedFloor {

							a3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a3.floor)

						}

						a3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.status = "idle"
						status()

					}
				}
			}

			elevA4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a4.floor)

				for a4.floor < userFloor {
					a4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)
				}

				a4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)

						for a4.floor < requestedFloor {

							a4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a4.floor)

						}

						a4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.status = "idle"
						status()

					}
				}
			}

			elevA5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a5.floor)

				for a5.floor < userFloor {
					a5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)
				}

				a5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)

						for a5.floor < requestedFloor {

							a5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a5.floor)

						}

						a5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.status = "idle"
						status()

					}
				}
			}

			elevA1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a1.floor)

				for a1.floor > userFloor {
					a1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)
				}

				a1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)

						for a1.floor < requestedFloor {

							a1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a1.floor)

						}

						a1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.status = "idle"
						status()

					}
				}
			}

			elevA2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a2.floor)

				for a2.floor > userFloor {
					a2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)
				}

				a2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)

						for a2.floor < requestedFloor {

							a2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a2.floor)

						}

						a2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.status = "idle"
						status()

					}
				}
			}

			elevA3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a3.floor)

				for a3.floor > userFloor {
					a3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)
				}

				a3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)

						for a3.floor < requestedFloor {

							a3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a3.floor)

						}

						a3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.status = "idle"
						status()

					}
				}
			}

			elevA4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a4.floor)

				for a4.floor > userFloor {
					a4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)
				}

				a4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)

						for a4.floor < requestedFloor {

							a4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a4.floor)

						}

						a4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.status = "idle"
						status()

					}
				}
			}

			elevA5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a5.floor)

				for a5.floor > userFloor {
					a5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)
				}

				a5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)

						for a5.floor < requestedFloor {

							a5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a5.floor)

						}

						a5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingUp
			if userFloor == a1.floor && a1.status == "goingUp" {
				elevA1()

			} else if userFloor == a2.floor && a2.status == "goingUp" {
				elevA2()

			} else if userFloor == a3.floor && a3.status == "goingUp" {
				elevA3()

			} else if userFloor == a4.floor && a4.status == "goingUp" {
				elevA4()

			} else if userFloor == a5.floor && a5.status == "goingUp" {
				elevA5()

				// userFloor > elevFloor && status == goingUp | 5 OPTIONS
			} else if userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				if v <= w && v <= x && v <= y && v <= z {
					elevA1()

				} else if w <= v && w <= x && w <= y && w <= z {
					elevA2()

				} else if x <= v && x <= w && x <= y && x <= z {
					elevA3()

				} else if y <= v && y <= w && y <= x && y <= z {
					elevA4()

				} else if z <= v && z <= w && z <= x && z <= y {
					elevA5()

				} else {
					elevA1()

				}

				// userFloor > elevFloor && status == goingUp | 4 OPTIONS
			} else if userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						elevA1()

					} else if w <= v && w <= x && w <= y {
						elevA2()

					} else if x <= v && x <= w && x <= y {
						elevA3()

					} else if y <= v && y <= w && y <= x {
						elevA4()

					} else {
						elevA1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						elevA1()

					} else if w <= v && w <= x && w <= z {
						elevA2()

					} else if x <= v && x <= w && x <= z {
						elevA3()

					} else if z <= v && z <= w && z <= x {
						elevA5()

					} else {
						elevA1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						elevA1()

					} else if w <= v && w <= y && w <= z {
						elevA2()

					} else if y <= v && y <= w && y <= z {
						elevA4()

					} else if z <= v && z <= w && z <= y {
						elevA5()

					} else {
						elevA1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						elevA1()

					} else if x <= v && x <= y && x <= z {
						elevA3()

					} else if y <= v && y <= x && y <= z {
						elevA4()

					} else if z <= v && z <= x && z <= y {
						elevA5()

					} else {
						elevA1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						elevA2()

					} else if x <= w && x <= y && x <= z {
						elevA3()

					} else if y <= w && y <= x && y <= z {
						elevA4()

					} else if z <= w && z <= x && z <= y {
						elevA5()

					} else {
						elevA2()

					}
				}

				// userFloor > elevFloor && status == goingUp | 3 OPTIONS
			} else if userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						elevA1()

					} else if w <= v && w <= x {
						elevA2()

					} else if x <= v && x <= w {
						elevA3()

					} else {
						elevA1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						elevA1()

					} else if w <= v && w <= y {
						elevA2()

					} else if y <= v && y <= w {
						elevA4()

					} else {
						elevA1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						elevA1()

					} else if w <= v && w <= z {
						elevA2()

					} else if z <= v && z <= w {
						elevA5()

					} else {
						elevA1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						elevA1()

					} else if x <= v && x <= y {
						elevA3()

					} else if y <= v && y <= x {
						elevA4()

					} else {
						elevA1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						elevA1()

					} else if x <= v && x <= z {
						elevA3()

					} else if z <= v && z <= x {
						elevA5()

					} else {
						elevA1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						elevA1()

					} else if y <= v && y <= z {
						elevA4()

					} else if z <= v && z <= y {
						elevA5()

					} else {
						elevA1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						elevA2()

					} else if x <= w && x <= y {
						elevA3()

					} else if y <= w && y <= x {
						elevA4()

					} else {
						elevA2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						elevA2()

					} else if x <= w && x <= z {
						elevA3()

					} else if z <= w && z <= x {
						elevA4()

					} else {
						elevA2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						elevA2()

					} else if y <= w && y <= z {
						elevA4()

					} else if z <= w && z <= y {
						elevA5()

					} else {
						elevA2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						elevA3()

					} else if y <= x && y <= z {
						elevA4()

					} else if z <= x && z <= y {
						elevA5()

					} else {
						elevA3()

					}
				}

				// userFloor > elevFloor && status == goingUp | 2 OPTIONS
			} else if userFloor > a1.floor && a1.status == "goingUp" && userFloor > a2.floor && a2.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a1.floor && a1.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a3.floor && a3.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a2.floor && a2.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a4.floor && a4.status == "goingUp" || userFloor > a3.floor && a3.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" || userFloor > a4.floor && a4.status == "goingUp" && userFloor > a5.floor && a5.status == "goingUp" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						elevA1()

					} else if w <= v {
						elevA2()

					} else {
						elevA1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						elevA1()

					} else if x <= v {
						elevA3()

					} else {
						elevA1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						elevA1()

					} else if y <= v {
						elevA4()

					} else {
						elevA1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						elevA1()

					} else if z <= v {
						elevA5()

					} else {
						elevA1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						elevA2()

					} else if x <= w {
						elevA3()

					} else {
						elevA2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						elevA2()

					} else if y <= w {
						elevA4()

					} else {
						elevA2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						elevA2()

					} else if z <= w {
						elevA5()

					} else {
						elevA2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						elevA3()

					} else if y <= x {
						elevA4()

					} else {
						elevA3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						elevA3()

					} else if z <= x {
						elevA5()

					} else {
						elevA3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						elevA4()

					} else if z <= y {
						elevA5()

					} else {
						elevA4()

					}
				}

				// userFloor > elevFloor && status == goingUp | 1 OPTION
			} else if userFloor > a1.floor && a1.status == "goingUp" {
				elevA1()

			} else if userFloor > a2.floor && a2.status == "goingUp" {
				elevA2()

			} else if userFloor > a3.floor && a3.status == "goingUp" {
				elevA3()

			} else if userFloor > a4.floor && a4.status == "goingUp" {
				elevA4()

			} else if userFloor > a5.floor && a5.status == "goingUp" {
				elevA5()

				// userFloor == elevFloor && status == idle | a1
			} else if userFloor == a1.floor && a1.status == "idle" {
				a1.status = "goingUp"
				status()
				elevA1()

				// userFloor == elevFloor && status == idle | a2
			} else if userFloor == a2.floor && a2.status == "idle" {
				a2.status = "goingUp"
				status()
				elevA2()

				// userFloor == elevFloor && status == idle | a3
			} else if userFloor == a3.floor && a3.status == "idle" {
				a3.status = "goingUp"
				status()
				elevA3()

				// userFloor == elevFloor && status == idle | a4
			} else if userFloor == a4.floor && a4.status == "idle" {
				a4.status = "goingUp"
				status()
				elevA4()

				// userFloor == elevFloor && status == idle | a5
			} else if userFloor == a5.floor && a5.status == "idle" {
				a5.status = "goingUp"
				status()
				elevA5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					a1.status = "goingUp"
					status()
					elevA1()

				} else if w <= v && w <= x && w <= y && w <= z {
					a2.status = "goingUp"
					status()
					elevA2()

				} else if x <= v && x <= w && x <= y && x <= z {
					a3.status = "goingUp"
					status()
					elevA3()

				} else if y <= v && y <= w && y <= x && y <= z {
					a4.status = "goingUp"
					status()
					elevA4()

				} else if z <= v && z <= w && z <= x && z <= y {
					a5.status = "goingUp"
					status()
					elevA5()

				} else {
					a1.status = "goingUp"
					status()
					elevA1()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= x && w <= y {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= v && x <= w && x <= y {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= v && y <= w && y <= x {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= x && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= v && x <= w && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if z <= v && z <= w && z <= x {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= y && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if y <= v && y <= w && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= v && z <= w && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if x <= v && x <= y && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= v && y <= x && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= v && z <= x && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= w && x <= y && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= w && y <= x && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= w && z <= x && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= x {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= v && x <= w {
						a3.status = "goingUp"
						status()
						elevA3()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= y {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if y <= v && y <= w {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if z <= v && z <= w {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if x <= v && x <= y {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= v && y <= x {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if x <= v && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if z <= v && z <= x {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if y <= v && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= v && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= w && x <= y {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= w && y <= x {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= w && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if z <= w && z <= x {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if y <= w && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= w && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= x && y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= x && z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a3.status = "goingUp"
						status()
						elevA3()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if w <= v {
						a2.status = "goingUp"
						status()
						elevA2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if x <= v {
						a3.status = "goingUp"
						status()
						elevA3()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if y <= v {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						a1.status = "goingUp"
						status()
						elevA1()

					} else if z <= v {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a1.status = "goingUp"
						status()
						elevA1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if x <= w {
						a3.status = "goingUp"
						status()
						elevA3()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if y <= w {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						a2.status = "goingUp"
						status()
						elevA2()

					} else if z <= w {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a2.status = "goingUp"
						status()
						elevA2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if y <= x {
						a4.status = "goingUp"
						status()
						elevA4()

					} else {
						a3.status = "goingUp"
						status()
						elevA3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						a3.status = "goingUp"
						status()
						elevA3()

					} else if z <= x {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a3.status = "goingUp"
						status()
						elevA3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						a4.status = "goingUp"
						status()
						elevA4()

					} else if z <= y {
						a5.status = "goingUp"
						status()
						elevA5()

					} else {
						a4.status = "goingUp"
						status()
						elevA4()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > a1.floor && a1.status == "idle" {
				a1.status = "goingUp"
				status()
				elevA1()

			} else if userFloor > a2.floor && a2.status == "idle" {
				a2.status = "goingUp"
				status()
				elevA2()

			} else if userFloor > a3.floor && a3.status == "idle" {
				a3.status = "goingUp"
				status()
				elevA3()

			} else if userFloor > a4.floor && a4.status == "idle" {
				a4.status = "goingUp"
				status()
				elevA4()

			} else if userFloor > a5.floor && a5.status == "idle" {
				a5.status = "goingUp"
				status()
				elevA5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					a1.status = "goingUp"
					status()
					elevA1v2()

				} else if w >= v && w >= x && w >= y && w >= z {
					a2.status = "goingUp"
					status()
					elevA2v2()

				} else if x >= v && x >= w && x >= y && x >= z {
					a3.status = "goingUp"
					status()
					elevA3v2()

				} else if y >= v && y >= w && y >= x && y >= z {
					a4.status = "goingUp"
					status()
					elevA4v2()

				} else if z >= v && z >= w && z >= x && z >= y {
					a5.status = "goingUp"
					status()
					elevA5v2()

				} else {
					a1.status = "goingUp"
					status()
					elevA1v2()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= x && w >= y {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= v && x >= w && x >= y {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= v && y >= w && y >= x {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= x && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= v && x >= w && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if z >= v && z >= w && z >= x {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= y && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if y >= v && y >= w && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= v && z >= w && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if x >= v && x >= y && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= v && y >= x && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= v && z >= x && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= w && x >= y && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= w && y >= x && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= w && z >= x && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= x {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= v && x >= w {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= y {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if y >= v && y >= w {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if z >= v && z >= w {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if x >= v && x >= y {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= v && y >= x {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if x >= v && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if z >= v && z >= x {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if y >= v && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= v && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= w && x >= y {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= w && y >= x {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= w && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if z >= w && z >= x {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if y >= w && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= w && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= x && y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= x && z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a3.status = "goingUp"
						status()
						elevA3v2()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if w >= v {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if x >= v {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if y >= v {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						a1.status = "goingUp"
						status()
						elevA1v2()

					} else if z >= v {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a1.status = "goingUp"
						status()
						elevA1v2()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if x >= w {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if y >= w {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						a2.status = "goingUp"
						status()
						elevA2v2()

					} else if z >= w {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a2.status = "goingUp"
						status()
						elevA2v2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if y >= x {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else {
						a3.status = "goingUp"
						status()
						elevA3v2()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						a3.status = "goingUp"
						status()
						elevA3v2()

					} else if z >= x {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a3.status = "goingUp"
						status()
						elevA3v2()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						a4.status = "goingUp"
						status()
						elevA4v2()

					} else if z >= y {
						a5.status = "goingUp"
						status()
						elevA5v2()

					} else {
						a4.status = "goingUp"
						status()
						elevA4v2()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < a1.floor && a1.status == "idle" {
				a1.status = "goingUp"
				status()
				elevA1v2()

			} else if userFloor < a2.floor && a2.status == "idle" {
				a2.status = "goingUp"
				status()
				elevA2v2()

			} else if userFloor < a3.floor && a3.status == "idle" {
				a3.status = "goingUp"
				status()
				elevA3v2()

			} else if userFloor < a4.floor && a4.status == "idle" {
				a4.status = "goingUp"
				status()
				elevA4v2()

			} else if userFloor < a5.floor && a5.status == "idle" {
				a5.status = "goingUp"
				status()
				elevA5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else if direction == "down" && userFloor >= -5 && userFloor <= 1 {

			elevA1 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator a1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a1.floor)

				for a1.floor > userFloor {
					a1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)
				}

				a1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)

						for a1.floor > requestedFloor {

							a1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a1.floor)

						}

						a1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.status = "idle"
						status()

					}
				}
			}

			elevA2 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator a2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a2.floor)

				for a2.floor > userFloor {
					a2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)
				}

				a2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)

						for a2.floor > requestedFloor {

							a2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a2.floor)

						}

						a2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.status = "idle"
						status()

					}
				}
			}

			elevA3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a3.floor)

				for a3.floor > userFloor {
					a3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)
				}

				a3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)

						for a3.floor > requestedFloor {

							a3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a3.floor)

						}

						a3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.status = "idle"
						status()

					}
				}
			}

			elevA4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a4.floor)

				for a4.floor > userFloor {
					a4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)
				}

				a4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)

						for a4.floor > requestedFloor {

							a4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a4.floor)

						}

						a4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.status = "idle"
						status()

					}
				}
			}

			elevA5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a5.floor)

				for a5.floor > userFloor {
					a5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)
				}

				a5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)

						for a5.floor > requestedFloor {

							a5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a5.floor)

						}

						a5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.status = "idle"
						status()

					}
				}
			}

			elevA1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a1.floor)

				for a1.floor < userFloor {
					a1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)
				}

				a1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)

						for a1.floor > requestedFloor {

							a1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a1.floor)

						}

						a1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a1.door)
						a1.status = "idle"
						status()

					}
				}
			}

			elevA2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a2.floor)

				for a2.floor < userFloor {
					a2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)
				}

				a2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)

						for a2.floor > requestedFloor {

							a2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a2.floor)

						}

						a2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a2.door)
						a2.status = "idle"
						status()

					}
				}
			}

			elevA3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a3.floor)

				for a3.floor < userFloor {
					a3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)
				}

				a3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)

						for a3.floor > requestedFloor {

							a3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a3.floor)

						}

						a3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a3.door)
						a3.status = "idle"
						status()

					}
				}
			}

			elevA4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a4.floor)

				for a4.floor < userFloor {
					a4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)
				}

				a4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)

						for a4.floor > requestedFloor {

							a4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a4.floor)

						}

						a4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a4.door)
						a4.status = "idle"
						status()

					}
				}
			}

			elevA5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator a5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", a5.floor)

				for a5.floor < userFloor {
					a5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)
				}

				a5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", a5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)

						for a5.floor > requestedFloor {

							a5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", a5.floor)

						}

						a5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", a5.door)
						a5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingDown
			if userFloor == a1.floor && a1.status == "goingDown" {
				elevA1()

			} else if userFloor == a2.floor && a2.status == "goingDown" {
				elevA2()

			} else if userFloor == a3.floor && a3.status == "goingDown" {
				elevA3()

			} else if userFloor == a4.floor && a4.status == "goingDown" {
				elevA4()

			} else if userFloor == a5.floor && a5.status == "goingDown" {
				elevA5()

				// userFloor < elevFloor && status == goingDown | 5 OPTIONS
			} else if userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					elevA1()

				} else if w >= v && w >= x && w >= y && w >= z {
					elevA2()

				} else if x >= v && x >= w && x >= y && x >= z {
					elevA3()

				} else if y >= v && y >= w && y >= x && y >= z {
					elevA4()

				} else if z >= v && z >= w && z >= x && z >= y {
					elevA5()

				} else {
					elevA1()

				}

				// userFloor < elevFloor && status == goingDown | 4 OPTIONS
			} else if userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						elevA1()

					} else if w >= v && w >= x && w >= y {
						elevA2()

					} else if x >= v && x >= w && x >= y {
						elevA3()

					} else if y >= v && y >= w && y >= x {
						elevA4()

					} else {
						elevA1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						elevA1()

					} else if w >= v && w >= x && w >= z {
						elevA2()

					} else if x >= v && x >= w && x >= z {
						elevA3()

					} else if z >= v && z >= w && z >= x {
						elevA5()

					} else {
						elevA1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						elevA1()

					} else if w >= v && w >= y && w >= z {
						elevA2()

					} else if y >= v && y >= w && y >= z {
						elevA4()

					} else if z >= v && z >= w && z >= y {
						elevA5()

					} else {
						elevA1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						elevA1()

					} else if x >= v && x >= y && x >= z {
						elevA3()

					} else if y >= v && y >= x && y >= z {
						elevA4()

					} else if z >= v && z >= x && z >= y {
						elevA5()

					} else {
						elevA1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						elevA2()

					} else if x >= w && x >= y && x >= z {
						elevA3()

					} else if y >= w && y >= x && y >= z {
						elevA4()

					} else if z >= w && z >= x && z >= y {
						elevA5()

					} else {
						elevA2()

					}
				}

				// userFloor < elevFloor && status == goingDown | 3 OPTIONS
			} else if userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						elevA1()

					} else if w >= v && w >= x {
						elevA2()

					} else if x >= v && x >= w {
						elevA3()

					} else {
						elevA1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						elevA1()

					} else if w >= v && w >= y {
						elevA2()

					} else if y >= v && y >= w {
						elevA4()

					} else {
						elevA1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						elevA1()

					} else if w >= v && w >= z {
						elevA2()

					} else if z >= v && z >= w {
						elevA5()

					} else {
						elevA1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						elevA1()

					} else if x >= v && x >= y {
						elevA3()

					} else if y >= v && y >= x {
						elevA4()

					} else {
						elevA1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						elevA1()

					} else if x >= v && x >= z {
						elevA3()

					} else if z >= v && z >= x {
						elevA5()

					} else {
						elevA1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						elevA1()

					} else if y >= v && y >= z {
						elevA4()

					} else if z >= v && z >= y {
						elevA5()

					} else {
						elevA1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						elevA2()

					} else if x >= w && x >= y {
						elevA3()

					} else if y >= w && y >= x {
						elevA4()

					} else {
						elevA2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						elevA2()

					} else if x >= w && x >= z {
						elevA3()

					} else if z >= w && z >= x {
						elevA4()

					} else {
						elevA2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						elevA2()

					} else if y >= w && y >= z {
						elevA4()

					} else if z >= w && z >= y {
						elevA5()

					} else {
						elevA2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						elevA3()

					} else if y >= x && y >= z {
						elevA4()

					} else if z >= x && z >= y {
						elevA5()

					} else {
						elevA3()

					}
				}

				// userFloor < elevFloor && status == goingDown | 2 OPTIONS
			} else if userFloor < a1.floor && a1.status == "goingDown" && userFloor < a2.floor && a2.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a1.floor && a1.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a3.floor && a3.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a2.floor && a2.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a4.floor && a4.status == "goingDown" || userFloor < a3.floor && a3.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" || userFloor < a4.floor && a4.status == "goingDown" && userFloor < a5.floor && a5.status == "goingDown" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						elevA1()

					} else if w >= v {
						elevA2()

					} else {
						elevA1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						elevA1()

					} else if x >= v {
						elevA3()

					} else {
						elevA1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						elevA1()

					} else if y >= v {
						elevA4()

					} else {
						elevA1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						elevA1()

					} else if z >= v {
						elevA5()

					} else {
						elevA1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						elevA2()

					} else if x >= w {
						elevA3()

					} else {
						elevA2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						elevA2()

					} else if y >= w {
						elevA4()

					} else {
						elevA2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						elevA2()

					} else if z >= w {
						elevA5()

					} else {
						elevA2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						elevA3()

					} else if y >= x {
						elevA4()

					} else {
						elevA3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						elevA3()

					} else if z >= x {
						elevA5()

					} else {
						elevA3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						elevA4()

					} else if z >= y {
						elevA5()

					} else {
						elevA4()

					}
				}

				// userFloor < elevFloor && status == goingDown | 1 OPTION
			} else if userFloor < a1.floor && a1.status == "goingDown" {
				elevA1()

			} else if userFloor < a2.floor && a2.status == "goingDown" {
				elevA2()

			} else if userFloor < a3.floor && a3.status == "goingDown" {
				elevA3()

			} else if userFloor < a4.floor && a4.status == "goingDown" {
				elevA4()

			} else if userFloor < a5.floor && a5.status == "goingDown" {
				elevA5()

				// userFloor == elevFloor && status == idle | a1
			} else if userFloor == a1.floor && a1.status == "idle" {
				a1.status = "goingDown"
				status()
				elevA1()

				// userFloor == elevFloor && status == idle | a2
			} else if userFloor == a2.floor && a2.status == "idle" {
				a2.status = "goingDown"
				status()
				elevA2()

				// userFloor == elevFloor && status == idle | a3
			} else if userFloor == a3.floor && a3.status == "idle" {
				a3.status = "goingDown"
				status()
				elevA3()

				// userFloor == elevFloor && status == idle | a4
			} else if userFloor == a4.floor && a4.status == "idle" {
				a4.status = "goingDown"
				status()
				elevA4()

				// userFloor == elevFloor && status == idle | a5
			} else if userFloor == a5.floor && a5.status == "idle" {
				a5.status = "goingDown"
				status()
				elevA5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					a1.status = "goingDown"
					status()
					elevA1()

				} else if w >= v && w >= x && w >= y && w >= z {
					a2.status = "goingDown"
					status()
					elevA2()

				} else if x >= v && x >= w && x >= y && x >= z {
					a3.status = "goingDown"
					status()
					elevA3()

				} else if y >= v && y >= w && y >= x && y >= z {
					a4.status = "goingDown"
					status()
					elevA4()

				} else if z >= v && z >= w && z >= x && z >= y {
					a5.status = "goingDown"
					status()
					elevA5()

				} else {
					a1.status = "goingDown"
					status()
					elevA1()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= x && w >= y {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= v && x >= w && x >= y {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= v && y >= w && y >= x {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= x && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= v && x >= w && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if z >= v && z >= w && z >= x {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= y && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if y >= v && y >= w && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= v && z >= w && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if x >= v && x >= y && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= v && y >= x && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= v && z >= x && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= w && x >= y && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= w && y >= x && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= w && z >= x && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= x {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= v && x >= w {
						a3.status = "goingDown"
						status()
						elevA3()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= y {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if y >= v && y >= w {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if z >= v && z >= w {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if x >= v && x >= y {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= v && y >= x {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if x >= v && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if z >= v && z >= x {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if y >= v && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= v && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= w && x >= y {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= w && y >= x {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= w && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if z >= w && z >= x {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if y >= w && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= w && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= x && y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= x && z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a3.status = "goingDown"
						status()
						elevA3()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < a1.floor && a1.status == "idle" && userFloor < a2.floor && a2.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a1.floor && a1.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a3.floor && a3.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a2.floor && a2.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a4.floor && a4.status == "idle" || userFloor < a3.floor && a3.status == "idle" && userFloor < a5.floor && a5.status == "idle" || userFloor < a4.floor && a4.status == "idle" && userFloor < a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if w >= v {
						a2.status = "goingDown"
						status()
						elevA2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if x >= v {
						a3.status = "goingDown"
						status()
						elevA3()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if y >= v {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						a1.status = "goingDown"
						status()
						elevA1()

					} else if z >= v {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a1.status = "goingDown"
						status()
						elevA1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if x >= w {
						a3.status = "goingDown"
						status()
						elevA3()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if y >= w {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						a2.status = "goingDown"
						status()
						elevA2()

					} else if z >= w {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a2.status = "goingDown"
						status()
						elevA2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if y >= x {
						a4.status = "goingDown"
						status()
						elevA4()

					} else {
						a3.status = "goingDown"
						status()
						elevA3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						a3.status = "goingDown"
						status()
						elevA3()

					} else if z >= x {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a3.status = "goingDown"
						status()
						elevA3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						a4.status = "goingDown"
						status()
						elevA4()

					} else if z >= y {
						a5.status = "goingDown"
						status()
						elevA5()

					} else {
						a4.status = "goingDown"
						status()
						elevA4()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < a1.floor && a1.status == "idle" {
				a1.status = "goingDown"
				status()
				elevA1()

			} else if userFloor < a2.floor && a2.status == "idle" {
				a2.status = "goingDown"
				status()
				elevA2()

			} else if userFloor < a3.floor && a3.status == "idle" {
				a3.status = "goingDown"
				status()
				elevA3()

			} else if userFloor < a4.floor && a4.status == "idle" {
				a4.status = "goingDown"
				status()
				elevA4()

			} else if userFloor < a5.floor && a5.status == "idle" {
				a5.status = "goingDown"
				status()
				elevA5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					a1.status = "goingDown"
					status()
					elevA1v2()

				} else if w <= v && w <= x && w <= y && w <= z {
					a2.status = "goingDown"
					status()
					elevA2v2()

				} else if x <= v && x <= w && x <= y && x <= z {
					a3.status = "goingDown"
					status()
					elevA3v2()

				} else if y <= v && y <= w && y <= x && y <= z {
					a4.status = "goingDown"
					status()
					elevA4v2()

				} else if z <= v && z <= w && z <= x && z <= y {
					a5.status = "goingDown"
					status()
					elevA5v2()

				} else {
					a1.status = "goingDown"
					status()
					elevA1v2()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= x && w <= y {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= v && x <= w && x <= y {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= v && y <= w && y <= x {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= x && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= v && x <= w && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if z <= v && z <= w && z <= x {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= y && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if y <= v && y <= w && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= v && z <= w && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if x <= v && x <= y && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= v && y <= x && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= v && z <= x && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= w && x <= y && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= w && y <= x && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= w && z <= x && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= x {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= v && x <= w {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= y {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if y <= v && y <= w {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if z <= v && z <= w {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if x <= v && x <= y {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= v && y <= x {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if x <= v && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if z <= v && z <= x {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if y <= v && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= v && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= w && x <= y {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= w && y <= x {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= w && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if z <= w && z <= x {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if y <= w && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= w && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= x && y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= x && z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a3.status = "goingDown"
						status()
						elevA3v2()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > a1.floor && a1.status == "idle" && userFloor > a2.floor && a2.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a1.floor && a1.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a3.floor && a3.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a2.floor && a2.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a4.floor && a4.status == "idle" || userFloor > a3.floor && a3.status == "idle" && userFloor > a5.floor && a5.status == "idle" || userFloor > a4.floor && a4.status == "idle" && userFloor > a5.floor && a5.status == "idle" {
				var v int = (userFloor - a1.floor)
				var w int = (userFloor - a2.floor)
				var x int = (userFloor - a3.floor)
				var y int = (userFloor - a4.floor)
				var z int = (userFloor - a5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if w <= v {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if x <= v {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if y <= v {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						a1.status = "goingDown"
						status()
						elevA1v2()

					} else if z <= v {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a1.status = "goingDown"
						status()
						elevA1v2()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if x <= w {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if y <= w {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						a2.status = "goingDown"
						status()
						elevA2v2()

					} else if z <= w {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a2.status = "goingDown"
						status()
						elevA2v2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if y <= x {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else {
						a3.status = "goingDown"
						status()
						elevA3v2()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						a3.status = "goingDown"
						status()
						elevA3v2()

					} else if z <= x {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a3.status = "goingDown"
						status()
						elevA3v2()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						a4.status = "goingDown"
						status()
						elevA4v2()

					} else if z <= y {
						a5.status = "goingDown"
						status()
						elevA5v2()

					} else {
						a4.status = "goingDown"
						status()
						elevA4v2()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > a1.floor && a1.status == "idle" {
				a1.status = "goingDown"
				status()
				elevA1v2()

			} else if userFloor > a2.floor && a2.status == "idle" {
				a2.status = "goingDown"
				status()
				elevA2v2()

			} else if userFloor > a3.floor && a3.status == "idle" {
				a3.status = "goingDown"
				status()
				elevA3v2()

			} else if userFloor > a4.floor && a4.status == "idle" {
				a4.status = "goingDown"
				status()
				elevA4v2()

			} else if userFloor > a5.floor && a5.status == "idle" {
				a5.status = "goingDown"
				status()
				elevA5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else {

			time.Sleep(time.Second)
			fmt.Printf("please enter valid information\n")

		}
	}

	requestElevB := func(userFloor int, direction string) {
		scanner := bufio.NewScanner(os.Stdin)

		if direction == "up" && userFloor >= 1 && userFloor < 20 {

			elevB1 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b1.floor)

				for b1.floor < userFloor {
					b1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)
				}

				b1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)

						for b1.floor < requestedFloor {

							b1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b1.floor)

						}

						b1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.status = "idle"
						status()

					}
				}
			}

			elevB2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b2.floor)

				for b2.floor < userFloor {
					b2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)
				}

				b2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)

						for b2.floor < requestedFloor {

							b2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b2.floor)

						}

						b2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.status = "idle"
						status()

					}
				}
			}

			elevB3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b3.floor)

				for b3.floor < userFloor {
					b3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)
				}

				b3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)

						for b3.floor < requestedFloor {

							b3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b3.floor)

						}

						b3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.status = "idle"
						status()

					}
				}
			}

			elevB4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b4.floor)

				for b4.floor < userFloor {
					b4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)
				}

				b4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)

						for b4.floor < requestedFloor {

							b4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b4.floor)

						}

						b4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.status = "idle"
						status()

					}
				}
			}

			elevB5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b5.floor)

				for b5.floor < userFloor {
					b5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)
				}

				b5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)

						for b5.floor < requestedFloor {

							b5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b5.floor)

						}

						b5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.status = "idle"
						status()

					}
				}
			}

			elevB1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b1.floor)

				for b1.floor > userFloor {
					b1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)
				}

				b1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)

						for b1.floor < requestedFloor {

							b1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b1.floor)

						}

						b1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.status = "idle"
						status()

					}
				}
			}

			elevB2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b2.floor)

				for b2.floor > userFloor {
					b2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)
				}

				b2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)

						for b2.floor < requestedFloor {

							b2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b2.floor)

						}

						b2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.status = "idle"
						status()

					}
				}
			}

			elevB3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b3.floor)

				for b3.floor > userFloor {
					b3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)
				}

				b3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)

						for b3.floor < requestedFloor {

							b3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b3.floor)

						}

						b3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.status = "idle"
						status()

					}
				}
			}

			elevB4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b4.floor)

				for b4.floor > userFloor {
					b4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)
				}

				b4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)

						for b4.floor < requestedFloor {

							b4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b4.floor)

						}

						b4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.status = "idle"
						status()

					}
				}
			}

			elevB5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b5.floor)

				for b5.floor > userFloor {
					b5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)
				}

				b5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)

						for b5.floor < requestedFloor {

							b5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b5.floor)

						}

						b5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingUp
			if userFloor == b1.floor && b1.status == "goingUp" {
				elevB1()

			} else if userFloor == b2.floor && b2.status == "goingUp" {
				elevB2()

			} else if userFloor == b3.floor && b3.status == "goingUp" {
				elevB3()

			} else if userFloor == b4.floor && b4.status == "goingUp" {
				elevB4()

			} else if userFloor == b5.floor && b5.status == "goingUp" {
				elevB5()

				// userFloor > elevFloor && status == goingUp | 5 OPTIONS
			} else if userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				if v <= w && v <= x && v <= y && v <= z {
					elevB1()

				} else if w <= v && w <= x && w <= y && w <= z {
					elevB2()

				} else if x <= v && x <= w && x <= y && x <= z {
					elevB3()

				} else if y <= v && y <= w && y <= x && y <= z {
					elevB4()

				} else if z <= v && z <= w && z <= x && z <= y {
					elevB5()

				} else {
					elevB1()

				}

				// userFloor > elevFloor && status == goingUp | 4 OPTIONS
			} else if userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						elevB1()

					} else if w <= v && w <= x && w <= y {
						elevB2()

					} else if x <= v && x <= w && x <= y {
						elevB3()

					} else if y <= v && y <= w && y <= x {
						elevB4()

					} else {
						elevB1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						elevB1()

					} else if w <= v && w <= x && w <= z {
						elevB2()

					} else if x <= v && x <= w && x <= z {
						elevB3()

					} else if z <= v && z <= w && z <= x {
						elevB5()

					} else {
						elevB1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						elevB1()

					} else if w <= v && w <= y && w <= z {
						elevB2()

					} else if y <= v && y <= w && y <= z {
						elevB4()

					} else if z <= v && z <= w && z <= y {
						elevB5()

					} else {
						elevB1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						elevB1()

					} else if x <= v && x <= y && x <= z {
						elevB3()

					} else if y <= v && y <= x && y <= z {
						elevB4()

					} else if z <= v && z <= x && z <= y {
						elevB5()

					} else {
						elevB1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						elevB2()

					} else if x <= w && x <= y && x <= z {
						elevB3()

					} else if y <= w && y <= x && y <= z {
						elevB4()

					} else if z <= w && z <= x && z <= y {
						elevB5()

					} else {
						elevB2()

					}
				}

				// userFloor > elevFloor && status == goingUp | 3 OPTIONS
			} else if userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						elevB1()

					} else if w <= v && w <= x {
						elevB2()

					} else if x <= v && x <= w {
						elevB3()

					} else {
						elevB1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						elevB1()

					} else if w <= v && w <= y {
						elevB2()

					} else if y <= v && y <= w {
						elevB4()

					} else {
						elevB1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						elevB1()

					} else if w <= v && w <= z {
						elevB2()

					} else if z <= v && z <= w {
						elevB5()

					} else {
						elevB1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						elevB1()

					} else if x <= v && x <= y {
						elevB3()

					} else if y <= v && y <= x {
						elevB4()

					} else {
						elevB1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						elevB1()

					} else if x <= v && x <= z {
						elevB3()

					} else if z <= v && z <= x {
						elevB5()

					} else {
						elevB1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						elevB1()

					} else if y <= v && y <= z {
						elevB4()

					} else if z <= v && z <= y {
						elevB5()

					} else {
						elevB1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						elevB2()

					} else if x <= w && x <= y {
						elevB3()

					} else if y <= w && y <= x {
						elevB4()

					} else {
						elevB2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						elevB2()

					} else if x <= w && x <= z {
						elevB3()

					} else if z <= w && z <= x {
						elevB4()

					} else {
						elevB2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						elevB2()

					} else if y <= w && y <= z {
						elevB4()

					} else if z <= w && z <= y {
						elevB5()

					} else {
						elevB2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						elevB3()

					} else if y <= x && y <= z {
						elevB4()

					} else if z <= x && z <= y {
						elevB5()

					} else {
						elevB3()

					}
				}

				// userFloor > elevFloor && status == goingUp | 2 OPTIONS
			} else if userFloor > b1.floor && b1.status == "goingUp" && userFloor > b2.floor && b2.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b1.floor && b1.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b3.floor && b3.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b2.floor && b2.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b3.floor && b3.status == "goingUp" && userFloor > b4.floor && b4.status == "goingUp" || userFloor > b3.floor && b3.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" || userFloor > b4.floor && b4.status == "goingUp" && userFloor > b5.floor && b5.status == "goingUp" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						elevB1()

					} else if w <= v {
						elevB2()

					} else {
						elevB1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						elevB1()

					} else if x <= v {
						elevB3()

					} else {
						elevB1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						elevB1()

					} else if y <= v {
						elevB4()

					} else {
						elevB1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						elevB1()

					} else if z <= v {
						elevB5()

					} else {
						elevB1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						elevB2()

					} else if x <= w {
						elevB3()

					} else {
						elevB2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						elevB2()

					} else if y <= w {
						elevB4()

					} else {
						elevB2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						elevB2()

					} else if z <= w {
						elevB5()

					} else {
						elevB2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						elevB3()

					} else if y <= x {
						elevB4()

					} else {
						elevB3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						elevB3()

					} else if z <= x {
						elevB5()

					} else {
						elevB3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						elevB4()

					} else if z <= y {
						elevB5()

					} else {
						elevB4()

					}
				}

				// userFloor > elevFloor && status == goingUp | 1 OPTION
			} else if userFloor > b1.floor && b1.status == "goingUp" {
				elevB1()

			} else if userFloor > b2.floor && b2.status == "goingUp" {
				elevB2()

			} else if userFloor > b3.floor && b3.status == "goingUp" {
				elevB3()

			} else if userFloor > b4.floor && b4.status == "goingUp" {
				elevB4()

			} else if userFloor > b5.floor && b5.status == "goingUp" {
				elevB5()

				// userFloor == elevFloor && status == idle | b1
			} else if userFloor == b1.floor && b1.status == "idle" {
				b1.status = "goingUp"
				status()
				elevB1()

				// userFloor == elevFloor && status == idle | b2
			} else if userFloor == b2.floor && b2.status == "idle" {
				b2.status = "goingUp"
				status()
				elevB2()

				// userFloor == elevFloor && status == idle | b3
			} else if userFloor == b3.floor && b3.status == "idle" {
				b3.status = "goingUp"
				status()
				elevB3()

				// userFloor == elevFloor && status == idle | b4
			} else if userFloor == b4.floor && b4.status == "idle" {
				b4.status = "goingUp"
				status()
				elevB4()

				// userFloor == elevFloor && status == idle | b5
			} else if userFloor == b5.floor && b5.status == "idle" {
				b5.status = "goingUp"
				status()
				elevB5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					b1.status = "goingUp"
					status()
					elevB1()

				} else if w <= v && w <= x && w <= y && w <= z {
					b2.status = "goingUp"
					status()
					elevB2()

				} else if x <= v && x <= w && x <= y && x <= z {
					b3.status = "goingUp"
					status()
					elevB3()

				} else if y <= v && y <= w && y <= x && y <= z {
					b4.status = "goingUp"
					status()
					elevB4()

				} else if z <= v && z <= w && z <= x && z <= y {
					b5.status = "goingUp"
					status()
					elevB5()

				} else {
					b1.status = "goingUp"
					status()
					elevB1()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= x && w <= y {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= v && x <= w && x <= y {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= v && y <= w && y <= x {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= x && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= v && x <= w && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if z <= v && z <= w && z <= x {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= y && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if y <= v && y <= w && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= v && z <= w && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if x <= v && x <= y && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= v && y <= x && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= v && z <= x && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= w && x <= y && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= w && y <= x && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= w && z <= x && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= x {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= v && x <= w {
						b3.status = "goingUp"
						status()
						elevB3()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= y {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if y <= v && y <= w {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if z <= v && z <= w {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if x <= v && x <= y {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= v && y <= x {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if x <= v && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if z <= v && z <= x {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if y <= v && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= v && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= w && x <= y {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= w && y <= x {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= w && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if z <= w && z <= x {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if y <= w && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= w && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= x && y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= x && z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b3.status = "goingUp"
						status()
						elevB3()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if w <= v {
						b2.status = "goingUp"
						status()
						elevB2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if x <= v {
						b3.status = "goingUp"
						status()
						elevB3()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if y <= v {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						b1.status = "goingUp"
						status()
						elevB1()

					} else if z <= v {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b1.status = "goingUp"
						status()
						elevB1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if x <= w {
						b3.status = "goingUp"
						status()
						elevB3()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if y <= w {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						b2.status = "goingUp"
						status()
						elevB2()

					} else if z <= w {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b2.status = "goingUp"
						status()
						elevB2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if y <= x {
						b4.status = "goingUp"
						status()
						elevB4()

					} else {
						b3.status = "goingUp"
						status()
						elevB3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						b3.status = "goingUp"
						status()
						elevB3()

					} else if z <= x {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b3.status = "goingUp"
						status()
						elevB3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						b4.status = "goingUp"
						status()
						elevB4()

					} else if z <= y {
						b5.status = "goingUp"
						status()
						elevB5()

					} else {
						b4.status = "goingUp"
						status()
						elevB4()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > b1.floor && b1.status == "idle" {
				b1.status = "goingUp"
				status()
				elevB1()

			} else if userFloor > b2.floor && b2.status == "idle" {
				b2.status = "goingUp"
				status()
				elevB2()

			} else if userFloor > b3.floor && b3.status == "idle" {
				b3.status = "goingUp"
				status()
				elevB3()

			} else if userFloor > b4.floor && b4.status == "idle" {
				b4.status = "goingUp"
				status()
				elevB4()

			} else if userFloor > b5.floor && b5.status == "idle" {
				b5.status = "goingUp"
				status()
				elevB5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					b1.status = "goingUp"
					status()
					elevB1v2()

				} else if w >= v && w >= x && w >= y && w >= z {
					b2.status = "goingUp"
					status()
					elevB2v2()

				} else if x >= v && x >= w && x >= y && x >= z {
					b3.status = "goingUp"
					status()
					elevB3v2()

				} else if y >= v && y >= w && y >= x && y >= z {
					b4.status = "goingUp"
					status()
					elevB4v2()

				} else if z >= v && z >= w && z >= x && z >= y {
					b5.status = "goingUp"
					status()
					elevB5v2()

				} else {
					b1.status = "goingUp"
					status()
					elevB1v2()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= x && w >= y {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= v && x >= w && x >= y {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= v && y >= w && y >= x {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= x && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= v && x >= w && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if z >= v && z >= w && z >= x {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= y && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if y >= v && y >= w && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= v && z >= w && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if x >= v && x >= y && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= v && y >= x && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= v && z >= x && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= w && x >= y && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= w && y >= x && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= w && z >= x && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= x {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= v && x >= w {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= y {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if y >= v && y >= w {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if z >= v && z >= w {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if x >= v && x >= y {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= v && y >= x {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if x >= v && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if z >= v && z >= x {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if y >= v && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= v && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= w && x >= y {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= w && y >= x {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= w && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if z >= w && z >= x {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if y >= w && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= w && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= x && y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= x && z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b3.status = "goingUp"
						status()
						elevB3v2()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if w >= v {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if x >= v {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if y >= v {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						b1.status = "goingUp"
						status()
						elevB1v2()

					} else if z >= v {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b1.status = "goingUp"
						status()
						elevB1v2()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if x >= w {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if y >= w {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						b2.status = "goingUp"
						status()
						elevB2v2()

					} else if z >= w {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b2.status = "goingUp"
						status()
						elevB2v2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if y >= x {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else {
						b3.status = "goingUp"
						status()
						elevB3v2()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						b3.status = "goingUp"
						status()
						elevB3v2()

					} else if z >= x {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b3.status = "goingUp"
						status()
						elevB3v2()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						b4.status = "goingUp"
						status()
						elevB4v2()

					} else if z >= y {
						b5.status = "goingUp"
						status()
						elevB5v2()

					} else {
						b4.status = "goingUp"
						status()
						elevB4v2()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < b1.floor && b1.status == "idle" {
				b1.status = "goingUp"
				status()
				elevB1v2()

			} else if userFloor < b2.floor && b2.status == "idle" {
				b2.status = "goingUp"
				status()
				elevB2v2()

			} else if userFloor < b3.floor && b3.status == "idle" {
				b3.status = "goingUp"
				status()
				elevB3v2()

			} else if userFloor < b4.floor && b4.status == "idle" {
				b4.status = "goingUp"
				status()
				elevB4v2()

			} else if userFloor < b5.floor && b5.status == "idle" {
				b5.status = "goingUp"
				status()
				elevB5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else if direction == "down" && userFloor > 1 && userFloor <= 20 {

			elevB1 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator b1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b1.floor)

				for b1.floor > userFloor {
					b1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)
				}

				b1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)

						for b1.floor > requestedFloor {

							b1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b1.floor)

						}

						b1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.status = "idle"
						status()

					}
				}
			}

			elevB2 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator b2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b2.floor)

				for b2.floor > userFloor {
					b2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)
				}

				b2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)

						for b2.floor > requestedFloor {

							b2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b2.floor)

						}

						b2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.status = "idle"
						status()

					}
				}
			}

			elevB3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b3.floor)

				for b3.floor > userFloor {
					b3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)
				}

				b3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)

						for b3.floor > requestedFloor {

							b3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b3.floor)

						}

						b3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.status = "idle"
						status()

					}
				}
			}

			elevB4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b4.floor)

				for b4.floor > userFloor {
					b4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)
				}

				b4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)

						for b4.floor > requestedFloor {

							b4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b4.floor)

						}

						b4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.status = "idle"
						status()

					}
				}
			}

			elevB5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b5.floor)

				for b5.floor > userFloor {
					b5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)
				}

				b5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)

						for b5.floor > requestedFloor {

							b5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b5.floor)

						}

						b5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.status = "idle"
						status()

					}
				}
			}

			elevB1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b1.floor)

				for b1.floor < userFloor {
					b1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)
				}

				b1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)

						for b1.floor > requestedFloor {

							b1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b1.floor)

						}

						b1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b1.door)
						b1.status = "idle"
						status()

					}
				}
			}

			elevB2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b2.floor)

				for b2.floor < userFloor {
					b2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)
				}

				b2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)

						for b2.floor > requestedFloor {

							b2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b2.floor)

						}

						b2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b2.door)
						b2.status = "idle"
						status()

					}
				}
			}

			elevB3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b3.floor)

				for b3.floor < userFloor {
					b3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)
				}

				b3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)

						for b3.floor > requestedFloor {

							b3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b3.floor)

						}

						b3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b3.door)
						b3.status = "idle"
						status()

					}
				}
			}

			elevB4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b4.floor)

				for b4.floor < userFloor {
					b4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)
				}

				b4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)

						for b4.floor > requestedFloor {

							b4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b4.floor)

						}

						b4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b4.door)
						b4.status = "idle"
						status()

					}
				}
			}

			elevB5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator b5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", b5.floor)

				for b5.floor < userFloor {
					b5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)
				}

				b5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", b5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)

						for b5.floor > requestedFloor {

							b5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", b5.floor)

						}

						b5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", b5.door)
						b5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingDown
			if userFloor == b1.floor && b1.status == "goingDown" {
				elevB1()

			} else if userFloor == b2.floor && b2.status == "goingDown" {
				elevB2()

			} else if userFloor == b3.floor && b3.status == "goingDown" {
				elevB3()

			} else if userFloor == b4.floor && b4.status == "goingDown" {
				elevB4()

			} else if userFloor == b5.floor && b5.status == "goingDown" {
				elevB5()

				// userFloor < elevFloor && status == goingDown | 5 OPTIONS
			} else if userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					elevB1()

				} else if w >= v && w >= x && w >= y && w >= z {
					elevB2()

				} else if x >= v && x >= w && x >= y && x >= z {
					elevB3()

				} else if y >= v && y >= w && y >= x && y >= z {
					elevB4()

				} else if z >= v && z >= w && z >= x && z >= y {
					elevB5()

				} else {
					elevB1()

				}

				// userFloor < elevFloor && status == goingDown | 4 OPTIONS
			} else if userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						elevB1()

					} else if w >= v && w >= x && w >= y {
						elevB2()

					} else if x >= v && x >= w && x >= y {
						elevB3()

					} else if y >= v && y >= w && y >= x {
						elevB4()

					} else {
						elevB1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						elevB1()

					} else if w >= v && w >= x && w >= z {
						elevB2()

					} else if x >= v && x >= w && x >= z {
						elevB3()

					} else if z >= v && z >= w && z >= x {
						elevB5()

					} else {
						elevB1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						elevB1()

					} else if w >= v && w >= y && w >= z {
						elevB2()

					} else if y >= v && y >= w && y >= z {
						elevB4()

					} else if z >= v && z >= w && z >= y {
						elevB5()

					} else {
						elevB1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						elevB1()

					} else if x >= v && x >= y && x >= z {
						elevB3()

					} else if y >= v && y >= x && y >= z {
						elevB4()

					} else if z >= v && z >= x && z >= y {
						elevB5()

					} else {
						elevB1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						elevB2()

					} else if x >= w && x >= y && x >= z {
						elevB3()

					} else if y >= w && y >= x && y >= z {
						elevB4()

					} else if z >= w && z >= x && z >= y {
						elevB5()

					} else {
						elevB2()

					}
				}

				// userFloor < elevFloor && status == goingDown | 3 OPTIONS
			} else if userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						elevB1()

					} else if w >= v && w >= x {
						elevB2()

					} else if x >= v && x >= w {
						elevB3()

					} else {
						elevB1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						elevB1()

					} else if w >= v && w >= y {
						elevB2()

					} else if y >= v && y >= w {
						elevB4()

					} else {
						elevB1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						elevB1()

					} else if w >= v && w >= z {
						elevB2()

					} else if z >= v && z >= w {
						elevB5()

					} else {
						elevB1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						elevB1()

					} else if x >= v && x >= y {
						elevB3()

					} else if y >= v && y >= x {
						elevB4()

					} else {
						elevB1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						elevB1()

					} else if x >= v && x >= z {
						elevB3()

					} else if z >= v && z >= x {
						elevB5()

					} else {
						elevB1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						elevB1()

					} else if y >= v && y >= z {
						elevB4()

					} else if z >= v && z >= y {
						elevB5()

					} else {
						elevB1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						elevB2()

					} else if x >= w && x >= y {
						elevB3()

					} else if y >= w && y >= x {
						elevB4()

					} else {
						elevB2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						elevB2()

					} else if x >= w && x >= z {
						elevB3()

					} else if z >= w && z >= x {
						elevB4()

					} else {
						elevB2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						elevB2()

					} else if y >= w && y >= z {
						elevB4()

					} else if z >= w && z >= y {
						elevB5()

					} else {
						elevB2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						elevB3()

					} else if y >= x && y >= z {
						elevB4()

					} else if z >= x && z >= y {
						elevB5()

					} else {
						elevB3()

					}
				}

				// userFloor < elevFloor && status == goingDown | 2 OPTIONS
			} else if userFloor < b1.floor && b1.status == "goingDown" && userFloor < b2.floor && b2.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b1.floor && b1.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b3.floor && b3.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b2.floor && b2.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b3.floor && b3.status == "goingDown" && userFloor < b4.floor && b4.status == "goingDown" || userFloor < b3.floor && b3.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" || userFloor < b4.floor && b4.status == "goingDown" && userFloor < b5.floor && b5.status == "goingDown" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						elevB1()

					} else if w >= v {
						elevB2()

					} else {
						elevB1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						elevB1()

					} else if x >= v {
						elevB3()

					} else {
						elevB1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						elevB1()

					} else if y >= v {
						elevB4()

					} else {
						elevB1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						elevB1()

					} else if z >= v {
						elevB5()

					} else {
						elevB1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						elevB2()

					} else if x >= w {
						elevB3()

					} else {
						elevB2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						elevB2()

					} else if y >= w {
						elevB4()

					} else {
						elevB2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						elevB2()

					} else if z >= w {
						elevB5()

					} else {
						elevB2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						elevB3()

					} else if y >= x {
						elevB4()

					} else {
						elevB3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						elevB3()

					} else if z >= x {
						elevB5()

					} else {
						elevB3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						elevB4()

					} else if z >= y {
						elevB5()

					} else {
						elevB4()

					}
				}

				// userFloor < elevFloor && status == goingDown | 1 OPTION
			} else if userFloor < b1.floor && b1.status == "goingDown" {
				elevB1()

			} else if userFloor < b2.floor && b2.status == "goingDown" {
				elevB2()

			} else if userFloor < b3.floor && b3.status == "goingDown" {
				elevB3()

			} else if userFloor < b4.floor && b4.status == "goingDown" {
				elevB4()

			} else if userFloor < b5.floor && b5.status == "goingDown" {
				elevB5()

				// userFloor == elevFloor && status == idle | b1
			} else if userFloor == b1.floor && b1.status == "idle" {
				b1.status = "goingDown"
				status()
				elevB1()

				// userFloor == elevFloor && status == idle | b2
			} else if userFloor == b2.floor && b2.status == "idle" {
				b2.status = "goingDown"
				status()
				elevB2()

				// userFloor == elevFloor && status == idle | b3
			} else if userFloor == b3.floor && b3.status == "idle" {
				b3.status = "goingDown"
				status()
				elevB3()

				// userFloor == elevFloor && status == idle | b4
			} else if userFloor == b4.floor && b4.status == "idle" {
				b4.status = "goingDown"
				status()
				elevB4()

				// userFloor == elevFloor && status == idle | b5
			} else if userFloor == b5.floor && b5.status == "idle" {
				b5.status = "goingDown"
				status()
				elevB5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					b1.status = "goingDown"
					status()
					elevB1()

				} else if w >= v && w >= x && w >= y && w >= z {
					b2.status = "goingDown"
					status()
					elevB2()

				} else if x >= v && x >= w && x >= y && x >= z {
					b3.status = "goingDown"
					status()
					elevB3()

				} else if y >= v && y >= w && y >= x && y >= z {
					b4.status = "goingDown"
					status()
					elevB4()

				} else if z >= v && z >= w && z >= x && z >= y {
					b5.status = "goingDown"
					status()
					elevB5()

				} else {
					b1.status = "goingDown"
					status()
					elevB1()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= x && w >= y {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= v && x >= w && x >= y {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= v && y >= w && y >= x {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= x && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= v && x >= w && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if z >= v && z >= w && z >= x {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= y && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if y >= v && y >= w && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= v && z >= w && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if x >= v && x >= y && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= v && y >= x && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= v && z >= x && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= w && x >= y && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= w && y >= x && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= w && z >= x && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= x {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= v && x >= w {
						b3.status = "goingDown"
						status()
						elevB3()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= y {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if y >= v && y >= w {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if z >= v && z >= w {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if x >= v && x >= y {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= v && y >= x {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if x >= v && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if z >= v && z >= x {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if y >= v && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= v && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= w && x >= y {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= w && y >= x {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= w && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if z >= w && z >= x {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if y >= w && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= w && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= x && y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= x && z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b3.status = "goingDown"
						status()
						elevB3()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < b1.floor && b1.status == "idle" && userFloor < b2.floor && b2.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b1.floor && b1.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b3.floor && b3.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b2.floor && b2.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b4.floor && b4.status == "idle" || userFloor < b3.floor && b3.status == "idle" && userFloor < b5.floor && b5.status == "idle" || userFloor < b4.floor && b4.status == "idle" && userFloor < b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if w >= v {
						b2.status = "goingDown"
						status()
						elevB2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if x >= v {
						b3.status = "goingDown"
						status()
						elevB3()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if y >= v {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						b1.status = "goingDown"
						status()
						elevB1()

					} else if z >= v {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b1.status = "goingDown"
						status()
						elevB1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if x >= w {
						b3.status = "goingDown"
						status()
						elevB3()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if y >= w {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						b2.status = "goingDown"
						status()
						elevB2()

					} else if z >= w {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b2.status = "goingDown"
						status()
						elevB2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if y >= x {
						b4.status = "goingDown"
						status()
						elevB4()

					} else {
						b3.status = "goingDown"
						status()
						elevB3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						b3.status = "goingDown"
						status()
						elevB3()

					} else if z >= x {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b3.status = "goingDown"
						status()
						elevB3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						b4.status = "goingDown"
						status()
						elevB4()

					} else if z >= y {
						b5.status = "goingDown"
						status()
						elevB5()

					} else {
						b4.status = "goingDown"
						status()
						elevB4()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < b1.floor && b1.status == "idle" {
				b1.status = "goingDown"
				status()
				elevB1()

			} else if userFloor < b2.floor && b2.status == "idle" {
				b2.status = "goingDown"
				status()
				elevB2()

			} else if userFloor < b3.floor && b3.status == "idle" {
				b3.status = "goingDown"
				status()
				elevB3()

			} else if userFloor < b4.floor && b4.status == "idle" {
				b4.status = "goingDown"
				status()
				elevB4()

			} else if userFloor < b5.floor && b5.status == "idle" {
				b5.status = "goingDown"
				status()
				elevB5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					b1.status = "goingDown"
					status()
					elevB1v2()

				} else if w <= v && w <= x && w <= y && w <= z {
					b2.status = "goingDown"
					status()
					elevB2v2()

				} else if x <= v && x <= w && x <= y && x <= z {
					b3.status = "goingDown"
					status()
					elevB3v2()

				} else if y <= v && y <= w && y <= x && y <= z {
					b4.status = "goingDown"
					status()
					elevB4v2()

				} else if z <= v && z <= w && z <= x && z <= y {
					b5.status = "goingDown"
					status()
					elevB5v2()

				} else {
					b1.status = "goingDown"
					status()
					elevB1v2()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= x && w <= y {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= v && x <= w && x <= y {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= v && y <= w && y <= x {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= x && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= v && x <= w && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if z <= v && z <= w && z <= x {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= y && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if y <= v && y <= w && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= v && z <= w && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if x <= v && x <= y && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= v && y <= x && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= v && z <= x && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= w && x <= y && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= w && y <= x && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= w && z <= x && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= x {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= v && x <= w {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= y {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if y <= v && y <= w {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if z <= v && z <= w {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if x <= v && x <= y {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= v && y <= x {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if x <= v && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if z <= v && z <= x {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if y <= v && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= v && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= w && x <= y {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= w && y <= x {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= w && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if z <= w && z <= x {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if y <= w && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= w && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= x && y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= x && z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b3.status = "goingDown"
						status()
						elevB3v2()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > b1.floor && b1.status == "idle" && userFloor > b2.floor && b2.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b1.floor && b1.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b3.floor && b3.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b2.floor && b2.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b4.floor && b4.status == "idle" || userFloor > b3.floor && b3.status == "idle" && userFloor > b5.floor && b5.status == "idle" || userFloor > b4.floor && b4.status == "idle" && userFloor > b5.floor && b5.status == "idle" {
				var v int = (userFloor - b1.floor)
				var w int = (userFloor - b2.floor)
				var x int = (userFloor - b3.floor)
				var y int = (userFloor - b4.floor)
				var z int = (userFloor - b5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if w <= v {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if x <= v {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if y <= v {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						b1.status = "goingDown"
						status()
						elevB1v2()

					} else if z <= v {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b1.status = "goingDown"
						status()
						elevB1v2()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if x <= w {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if y <= w {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						b2.status = "goingDown"
						status()
						elevB2v2()

					} else if z <= w {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b2.status = "goingDown"
						status()
						elevB2v2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if y <= x {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else {
						b3.status = "goingDown"
						status()
						elevB3v2()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						b3.status = "goingDown"
						status()
						elevB3v2()

					} else if z <= x {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b3.status = "goingDown"
						status()
						elevB3v2()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						b4.status = "goingDown"
						status()
						elevB4v2()

					} else if z <= y {
						b5.status = "goingDown"
						status()
						elevB5v2()

					} else {
						b4.status = "goingDown"
						status()
						elevB4v2()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > b1.floor && b1.status == "idle" {
				b1.status = "goingDown"
				status()
				elevB1v2()

			} else if userFloor > b2.floor && b2.status == "idle" {
				b2.status = "goingDown"
				status()
				elevB2v2()

			} else if userFloor > b3.floor && b3.status == "idle" {
				b3.status = "goingDown"
				status()
				elevB3v2()

			} else if userFloor > b4.floor && b4.status == "idle" {
				b4.status = "goingDown"
				status()
				elevB4v2()

			} else if userFloor > b5.floor && b5.status == "idle" {
				b5.status = "goingDown"
				status()
				elevB5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else {

			time.Sleep(time.Second)
			fmt.Printf("please enter valid information\n")

		}
	}

	requestElevC := func(userFloor int, direction string) {
		scanner := bufio.NewScanner(os.Stdin)

		if direction == "up" && userFloor >= 1 && userFloor < 40 || direction == "up" && userFloor == 1 {

			elevC1 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c1.floor)

				for c1.floor < userFloor {
					c1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)
				}

				c1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)

						for c1.floor < requestedFloor {

							c1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c1.floor)

						}

						c1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.status = "idle"
						status()

					}
				}
			}

			elevC2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c2.floor)

				for c2.floor < userFloor {
					c2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)
				}

				c2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)

						for c2.floor < requestedFloor {

							c2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c2.floor)

						}

						c2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.status = "idle"
						status()

					}
				}
			}

			elevC3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c3.floor)

				for c3.floor < userFloor {
					c3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)
				}

				c3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)

						for c3.floor < requestedFloor {

							c3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c3.floor)

						}

						c3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.status = "idle"
						status()

					}
				}
			}

			elevC4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c4.floor)

				for c4.floor < userFloor {
					c4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)
				}

				c4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)

						for c4.floor < requestedFloor {

							c4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c4.floor)

						}

						c4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.status = "idle"
						status()

					}
				}
			}

			elevC5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c5.floor)

				for c5.floor < userFloor {
					c5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)
				}

				c5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)

						for c5.floor < requestedFloor {

							c5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c5.floor)

						}

						c5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.status = "idle"
						status()

					}
				}
			}

			elevC1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c1.floor)

				for c1.floor > userFloor {
					c1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)
				}

				c1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)

						for c1.floor < requestedFloor {

							c1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c1.floor)

						}

						c1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.status = "idle"
						status()

					}
				}
			}

			elevC2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c2.floor)

				for c2.floor > userFloor {
					c2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)
				}

				c2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)

						for c2.floor < requestedFloor {

							c2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c2.floor)

						}

						c2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.status = "idle"
						status()

					}
				}
			}

			elevC3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c3.floor)

				for c3.floor > userFloor {
					c3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)
				}

				c3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)

						for c3.floor < requestedFloor {

							c3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c3.floor)

						}

						c3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.status = "idle"
						status()

					}
				}
			}

			elevC4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c4.floor)

				for c4.floor > userFloor {
					c4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)
				}

				c4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)

						for c4.floor < requestedFloor {

							c4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c4.floor)

						}

						c4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.status = "idle"
						status()

					}
				}
			}

			elevC5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c5.floor)

				for c5.floor > userFloor {
					c5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)
				}

				c5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)

						for c5.floor < requestedFloor {

							c5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c5.floor)

						}

						c5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingUp
			if userFloor == c1.floor && c1.status == "goingUp" {
				elevC1()

			} else if userFloor == c2.floor && c2.status == "goingUp" {
				elevC2()

			} else if userFloor == c3.floor && c3.status == "goingUp" {
				elevC3()

			} else if userFloor == c4.floor && c4.status == "goingUp" {
				elevC4()

			} else if userFloor == c5.floor && c5.status == "goingUp" {
				elevC5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					c1.status = "goingUp"
					status()
					elevC1()

				} else if w <= v && w <= x && w <= y && w <= z {
					c2.status = "goingUp"
					status()
					elevC2()

				} else if x <= v && x <= w && x <= y && x <= z {
					c3.status = "goingUp"
					status()
					elevC3()

				} else if y <= v && y <= w && y <= x && y <= z {
					c4.status = "goingUp"
					status()
					elevC4()

				} else if z <= v && z <= w && z <= x && z <= y {
					c5.status = "goingUp"
					status()
					elevC5()

				} else {
					c1.status = "goingUp"
					status()
					elevC1()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= x && w <= y {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= v && x <= w && x <= y {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= v && y <= w && y <= x {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= x && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= v && x <= w && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if z <= v && z <= w && z <= x {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= y && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if y <= v && y <= w && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= v && z <= w && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if x <= v && x <= y && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= v && y <= x && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= v && z <= x && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= w && x <= y && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= w && y <= x && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= w && z <= x && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= x {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= v && x <= w {
						c3.status = "goingUp"
						status()
						elevC3()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= y {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if y <= v && y <= w {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if z <= v && z <= w {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if x <= v && x <= y {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= v && y <= x {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if x <= v && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if z <= v && z <= x {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if y <= v && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= v && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= w && x <= y {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= w && y <= x {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= w && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if z <= w && z <= x {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if y <= w && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= w && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= x && y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= x && z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c3.status = "goingUp"
						status()
						elevC3()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if w <= v {
						c2.status = "goingUp"
						status()
						elevC2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if x <= v {
						c3.status = "goingUp"
						status()
						elevC3()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if y <= v {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						c1.status = "goingUp"
						status()
						elevC1()

					} else if z <= v {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c1.status = "goingUp"
						status()
						elevC1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if x <= w {
						c3.status = "goingUp"
						status()
						elevC3()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if y <= w {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						c2.status = "goingUp"
						status()
						elevC2()

					} else if z <= w {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c2.status = "goingUp"
						status()
						elevC2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if y <= x {
						c4.status = "goingUp"
						status()
						elevC4()

					} else {
						c3.status = "goingUp"
						status()
						elevC3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						c3.status = "goingUp"
						status()
						elevC3()

					} else if z <= x {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c3.status = "goingUp"
						status()
						elevC3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						c4.status = "goingUp"
						status()
						elevC4()

					} else if z <= y {
						c5.status = "goingUp"
						status()
						elevC5()

					} else {
						c4.status = "goingUp"
						status()
						elevC4()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > c1.floor && c1.status == "idle" {
				c1.status = "goingUp"
				status()
				elevC1()

			} else if userFloor > c2.floor && c2.status == "idle" {
				c2.status = "goingUp"
				status()
				elevC2()

			} else if userFloor > c3.floor && c3.status == "idle" {
				c3.status = "goingUp"
				status()
				elevC3()

			} else if userFloor > c4.floor && c4.status == "idle" {
				c4.status = "goingUp"
				status()
				elevC4()

			} else if userFloor > c5.floor && c5.status == "idle" {
				c5.status = "goingUp"
				status()
				elevC5()

				// userFloor > elevFloor && status == goingUp | 5 OPTIONS
			} else if userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				if v <= w && v <= x && v <= y && v <= z {
					elevC1()

				} else if w <= v && w <= x && w <= y && w <= z {
					elevC2()

				} else if x <= v && x <= w && x <= y && x <= z {
					elevC3()

				} else if y <= v && y <= w && y <= x && y <= z {
					elevC4()

				} else if z <= v && z <= w && z <= x && z <= y {
					elevC5()

				} else {
					elevC1()

				}

				// userFloor > elevFloor && status == goingUp | 4 OPTIONS
			} else if userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						elevC1()

					} else if w <= v && w <= x && w <= y {
						elevC2()

					} else if x <= v && x <= w && x <= y {
						elevC3()

					} else if y <= v && y <= w && y <= x {
						elevC4()

					} else {
						elevC1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						elevC1()

					} else if w <= v && w <= x && w <= z {
						elevC2()

					} else if x <= v && x <= w && x <= z {
						elevC3()

					} else if z <= v && z <= w && z <= x {
						elevC5()

					} else {
						elevC1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						elevC1()

					} else if w <= v && w <= y && w <= z {
						elevC2()

					} else if y <= v && y <= w && y <= z {
						elevC4()

					} else if z <= v && z <= w && z <= y {
						elevC5()

					} else {
						elevC1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						elevC1()

					} else if x <= v && x <= y && x <= z {
						elevC3()

					} else if y <= v && y <= x && y <= z {
						elevC4()

					} else if z <= v && z <= x && z <= y {
						elevC5()

					} else {
						elevC1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						elevC2()

					} else if x <= w && x <= y && x <= z {
						elevC3()

					} else if y <= w && y <= x && y <= z {
						elevC4()

					} else if z <= w && z <= x && z <= y {
						elevC5()

					} else {
						elevC2()

					}
				}

				// userFloor > elevFloor && status == goingUp | 3 OPTIONS
			} else if userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						elevC1()

					} else if w <= v && w <= x {
						elevC2()

					} else if x <= v && x <= w {
						elevC3()

					} else {
						elevC1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						elevC1()

					} else if w <= v && w <= y {
						elevC2()

					} else if y <= v && y <= w {
						elevC4()

					} else {
						elevC1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						elevC1()

					} else if w <= v && w <= z {
						elevC2()

					} else if z <= v && z <= w {
						elevC5()

					} else {
						elevC1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						elevC1()

					} else if x <= v && x <= y {
						elevC3()

					} else if y <= v && y <= x {
						elevC4()

					} else {
						elevC1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						elevC1()

					} else if x <= v && x <= z {
						elevC3()

					} else if z <= v && z <= x {
						elevC5()

					} else {
						elevC1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						elevC1()

					} else if y <= v && y <= z {
						elevC4()

					} else if z <= v && z <= y {
						elevC5()

					} else {
						elevC1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						elevC2()

					} else if x <= w && x <= y {
						elevC3()

					} else if y <= w && y <= x {
						elevC4()

					} else {
						elevC2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						elevC2()

					} else if x <= w && x <= z {
						elevC3()

					} else if z <= w && z <= x {
						elevC4()

					} else {
						elevC2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						elevC2()

					} else if y <= w && y <= z {
						elevC4()

					} else if z <= w && z <= y {
						elevC5()

					} else {
						elevC2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						elevC3()

					} else if y <= x && y <= z {
						elevC4()

					} else if z <= x && z <= y {
						elevC5()

					} else {
						elevC3()

					}
				}

				// userFloor > elevFloor && status == goingUp | 2 OPTIONS
			} else if userFloor > c1.floor && c1.status == "goingUp" && userFloor > c2.floor && c2.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c1.floor && c1.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c3.floor && c3.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c2.floor && c2.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c3.floor && c3.status == "goingUp" && userFloor > c4.floor && c4.status == "goingUp" || userFloor > c3.floor && c3.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" || userFloor > c4.floor && c4.status == "goingUp" && userFloor > c5.floor && c5.status == "goingUp" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						elevC1()

					} else if w <= v {
						elevC2()

					} else {
						elevC1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						elevC1()

					} else if x <= v {
						elevC3()

					} else {
						elevC1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						elevC1()

					} else if y <= v {
						elevC4()

					} else {
						elevC1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						elevC1()

					} else if z <= v {
						elevC5()

					} else {
						elevC1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						elevC2()

					} else if x <= w {
						elevC3()

					} else {
						elevC2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						elevC2()

					} else if y <= w {
						elevC4()

					} else {
						elevC2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						elevC2()

					} else if z <= w {
						elevC5()

					} else {
						elevC2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						elevC3()

					} else if y <= x {
						elevC4()

					} else {
						elevC3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						elevC3()

					} else if z <= x {
						elevC5()

					} else {
						elevC3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						elevC4()

					} else if z <= y {
						elevC5()

					} else {
						elevC4()

					}
				}

				// userFloor > elevFloor && status == goingUp | 1 OPTION
			} else if userFloor > c1.floor && c1.status == "goingUp" {
				elevC1()

			} else if userFloor > c2.floor && c2.status == "goingUp" {
				elevC2()

			} else if userFloor > c3.floor && c3.status == "goingUp" {
				elevC3()

			} else if userFloor > c4.floor && c4.status == "goingUp" {
				elevC4()

			} else if userFloor > c5.floor && c5.status == "goingUp" {
				elevC5()

				// userFloor == elevFloor && status == idle | c1
			} else if userFloor == c1.floor && c1.status == "idle" {
				c1.status = "goingUp"
				status()
				elevC1()

				// userFloor == elevFloor && status == idle | c2
			} else if userFloor == c2.floor && c2.status == "idle" {
				c2.status = "goingUp"
				status()
				elevC2()

				// userFloor == elevFloor && status == idle | c3
			} else if userFloor == c3.floor && c3.status == "idle" {
				c3.status = "goingUp"
				status()
				elevC3()

				// userFloor == elevFloor && status == idle | c4
			} else if userFloor == c4.floor && c4.status == "idle" {
				c4.status = "goingUp"
				status()
				elevC4()

				// userFloor == elevFloor && status == idle | c5
			} else if userFloor == c5.floor && c5.status == "idle" {
				c5.status = "goingUp"
				status()
				elevC5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					c1.status = "goingUp"
					status()
					elevC1v2()

				} else if w >= v && w >= x && w >= y && w >= z {
					c2.status = "goingUp"
					status()
					elevC2v2()

				} else if x >= v && x >= w && x >= y && x >= z {
					c3.status = "goingUp"
					status()
					elevC3v2()

				} else if y >= v && y >= w && y >= x && y >= z {
					c4.status = "goingUp"
					status()
					elevC4v2()

				} else if z >= v && z >= w && z >= x && z >= y {
					c5.status = "goingUp"
					status()
					elevC5v2()

				} else {
					c1.status = "goingUp"
					status()
					elevC1v2()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= x && w >= y {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= v && x >= w && x >= y {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= v && y >= w && y >= x {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= x && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= v && x >= w && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if z >= v && z >= w && z >= x {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= y && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if y >= v && y >= w && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= v && z >= w && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if x >= v && x >= y && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= v && y >= x && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= v && z >= x && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= w && x >= y && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= w && y >= x && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= w && z >= x && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= x {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= v && x >= w {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= y {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if y >= v && y >= w {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if z >= v && z >= w {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if x >= v && x >= y {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= v && y >= x {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if x >= v && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if z >= v && z >= x {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if y >= v && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= v && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= w && x >= y {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= w && y >= x {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= w && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if z >= w && z >= x {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if y >= w && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= w && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= x && y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= x && z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c3.status = "goingUp"
						status()
						elevC3v2()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if w >= v {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if x >= v {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if y >= v {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						c1.status = "goingUp"
						status()
						elevC1v2()

					} else if z >= v {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c1.status = "goingUp"
						status()
						elevC1v2()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if x >= w {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if y >= w {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						c2.status = "goingUp"
						status()
						elevC2v2()

					} else if z >= w {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c2.status = "goingUp"
						status()
						elevC2v2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if y >= x {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else {
						c3.status = "goingUp"
						status()
						elevC3v2()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						c3.status = "goingUp"
						status()
						elevC3v2()

					} else if z >= x {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c3.status = "goingUp"
						status()
						elevC3v2()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						c4.status = "goingUp"
						status()
						elevC4v2()

					} else if z >= y {
						c5.status = "goingUp"
						status()
						elevC5v2()

					} else {
						c4.status = "goingUp"
						status()
						elevC4v2()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < c1.floor && c1.status == "idle" {
				c1.status = "goingUp"
				status()
				elevC1v2()

			} else if userFloor < c2.floor && c2.status == "idle" {
				c2.status = "goingUp"
				status()
				elevC2v2()

			} else if userFloor < c3.floor && c3.status == "idle" {
				c3.status = "goingUp"
				status()
				elevC3v2()

			} else if userFloor < c4.floor && c4.status == "idle" {
				c4.status = "goingUp"
				status()
				elevC4v2()

			} else if userFloor < c5.floor && c5.status == "idle" {
				c5.status = "goingUp"
				status()
				elevC5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else if direction == "down" && userFloor >= 21 && userFloor <= 40 {

			elevC1 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator c1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c1.floor)

				for c1.floor > userFloor {
					c1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)
				}

				c1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)

						for c1.floor > requestedFloor {

							c1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c1.floor)

						}

						c1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.status = "idle"
						status()

					}
				}
			}

			elevC2 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator c2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c2.floor)

				for c2.floor > userFloor {
					c2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)
				}

				c2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)

						for c2.floor > requestedFloor {

							c2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c2.floor)

						}

						c2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.status = "idle"
						status()

					}
				}
			}

			elevC3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c3.floor)

				for c3.floor > userFloor {
					c3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)
				}

				c3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)

						for c3.floor > requestedFloor {

							c3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c3.floor)

						}

						c3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.status = "idle"
						status()

					}
				}
			}

			elevC4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c4.floor)

				for c4.floor > userFloor {
					c4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)
				}

				c4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)

						for c4.floor > requestedFloor {

							c4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c4.floor)

						}

						c4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.status = "idle"
						status()

					}
				}
			}

			elevC5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c5.floor)

				for c5.floor > userFloor {
					c5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)
				}

				c5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)

						for c5.floor > requestedFloor {

							c5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c5.floor)

						}

						c5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.status = "idle"
						status()

					}
				}
			}

			elevC1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c1.floor)

				for c1.floor < userFloor {
					c1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)
				}

				c1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)

						for c1.floor > requestedFloor {

							c1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c1.floor)

						}

						c1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c1.door)
						c1.status = "idle"
						status()

					}
				}
			}

			elevC2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c2.floor)

				for c2.floor < userFloor {
					c2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)
				}

				c2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)

						for c2.floor > requestedFloor {

							c2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c2.floor)

						}

						c2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c2.door)
						c2.status = "idle"
						status()

					}
				}
			}

			elevC3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c3.floor)

				for c3.floor < userFloor {
					c3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)
				}

				c3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)

						for c3.floor > requestedFloor {

							c3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c3.floor)

						}

						c3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c3.door)
						c3.status = "idle"
						status()

					}
				}
			}

			elevC4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c4.floor)

				for c4.floor < userFloor {
					c4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)
				}

				c4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)

						for c4.floor > requestedFloor {

							c4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c4.floor)

						}

						c4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c4.door)
						c4.status = "idle"
						status()

					}
				}
			}

			elevC5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator c5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", c5.floor)

				for c5.floor < userFloor {
					c5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)
				}

				c5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", c5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if requestedFloor < 1 || requestedFloor >= 20 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)

						for c5.floor > requestedFloor {

							c5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", c5.floor)

						}

						c5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", c5.door)
						c5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingDown
			if userFloor == c1.floor && c1.status == "goingDown" {
				elevC1()

			} else if userFloor == c2.floor && c2.status == "goingDown" {
				elevC2()

			} else if userFloor == c3.floor && c3.status == "goingDown" {
				elevC3()

			} else if userFloor == c4.floor && c4.status == "goingDown" {
				elevC4()

			} else if userFloor == c5.floor && c5.status == "goingDown" {
				elevC5()

				// userFloor < elevFloor && status == goingDown | 5 OPTIONS
			} else if userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					elevC1()

				} else if w >= v && w >= x && w >= y && w >= z {
					elevC2()

				} else if x >= v && x >= w && x >= y && x >= z {
					elevC3()

				} else if y >= v && y >= w && y >= x && y >= z {
					elevC4()

				} else if z >= v && z >= w && z >= x && z >= y {
					elevC5()

				} else {
					elevC1()

				}

				// userFloor < elevFloor && status == goingDown | 4 OPTIONS
			} else if userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						elevC1()

					} else if w >= v && w >= x && w >= y {
						elevC2()

					} else if x >= v && x >= w && x >= y {
						elevC3()

					} else if y >= v && y >= w && y >= x {
						elevC4()

					} else {
						elevC1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						elevC1()

					} else if w >= v && w >= x && w >= z {
						elevC2()

					} else if x >= v && x >= w && x >= z {
						elevC3()

					} else if z >= v && z >= w && z >= x {
						elevC5()

					} else {
						elevC1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						elevC1()

					} else if w >= v && w >= y && w >= z {
						elevC2()

					} else if y >= v && y >= w && y >= z {
						elevC4()

					} else if z >= v && z >= w && z >= y {
						elevC5()

					} else {
						elevC1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						elevC1()

					} else if x >= v && x >= y && x >= z {
						elevC3()

					} else if y >= v && y >= x && y >= z {
						elevC4()

					} else if z >= v && z >= x && z >= y {
						elevC5()

					} else {
						elevC1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						elevC2()

					} else if x >= w && x >= y && x >= z {
						elevC3()

					} else if y >= w && y >= x && y >= z {
						elevC4()

					} else if z >= w && z >= x && z >= y {
						elevC5()

					} else {
						elevC2()

					}
				}

				// userFloor < elevFloor && status == goingDown | 3 OPTIONS
			} else if userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						elevC1()

					} else if w >= v && w >= x {
						elevC2()

					} else if x >= v && x >= w {
						elevC3()

					} else {
						elevC1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						elevC1()

					} else if w >= v && w >= y {
						elevC2()

					} else if y >= v && y >= w {
						elevC4()

					} else {
						elevC1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						elevC1()

					} else if w >= v && w >= z {
						elevC2()

					} else if z >= v && z >= w {
						elevC5()

					} else {
						elevC1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						elevC1()

					} else if x >= v && x >= y {
						elevC3()

					} else if y >= v && y >= x {
						elevC4()

					} else {
						elevC1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						elevC1()

					} else if x >= v && x >= z {
						elevC3()

					} else if z >= v && z >= x {
						elevC5()

					} else {
						elevC1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						elevC1()

					} else if y >= v && y >= z {
						elevC4()

					} else if z >= v && z >= y {
						elevC5()

					} else {
						elevC1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						elevC2()

					} else if x >= w && x >= y {
						elevC3()

					} else if y >= w && y >= x {
						elevC4()

					} else {
						elevC2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						elevC2()

					} else if x >= w && x >= z {
						elevC3()

					} else if z >= w && z >= x {
						elevC4()

					} else {
						elevC2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						elevC2()

					} else if y >= w && y >= z {
						elevC4()

					} else if z >= w && z >= y {
						elevC5()

					} else {
						elevC2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						elevC3()

					} else if y >= x && y >= z {
						elevC4()

					} else if z >= x && z >= y {
						elevC5()

					} else {
						elevC3()

					}
				}

				// userFloor < elevFloor && status == goingDown | 2 OPTIONS
			} else if userFloor < c1.floor && c1.status == "goingDown" && userFloor < c2.floor && c2.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c1.floor && c1.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c3.floor && c3.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c2.floor && c2.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c3.floor && c3.status == "goingDown" && userFloor < c4.floor && c4.status == "goingDown" || userFloor < c3.floor && c3.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" || userFloor < c4.floor && c4.status == "goingDown" && userFloor < c5.floor && c5.status == "goingDown" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						elevC1()

					} else if w >= v {
						elevC2()

					} else {
						elevC1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						elevC1()

					} else if x >= v {
						elevC3()

					} else {
						elevC1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						elevC1()

					} else if y >= v {
						elevC4()

					} else {
						elevC1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						elevC1()

					} else if z >= v {
						elevC5()

					} else {
						elevC1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						elevC2()

					} else if x >= w {
						elevC3()

					} else {
						elevC2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						elevC2()

					} else if y >= w {
						elevC4()

					} else {
						elevC2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						elevC2()

					} else if z >= w {
						elevC5()

					} else {
						elevC2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						elevC3()

					} else if y >= x {
						elevC4()

					} else {
						elevC3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						elevC3()

					} else if z >= x {
						elevC5()

					} else {
						elevC3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						elevC4()

					} else if z >= y {
						elevC5()

					} else {
						elevC4()

					}
				}

				// userFloor < elevFloor && status == goingDown | 1 OPTION
			} else if userFloor < c1.floor && c1.status == "goingDown" {
				elevC1()

			} else if userFloor < c2.floor && c2.status == "goingDown" {
				elevC2()

			} else if userFloor < c3.floor && c3.status == "goingDown" {
				elevC3()

			} else if userFloor < c4.floor && c4.status == "goingDown" {
				elevC4()

			} else if userFloor < c5.floor && c5.status == "goingDown" {
				elevC5()

				// userFloor == elevFloor && status == idle | c1
			} else if userFloor == c1.floor && c1.status == "idle" {
				c1.status = "goingDown"
				status()
				elevC1()

				// userFloor == elevFloor && status == idle | c2
			} else if userFloor == c2.floor && c2.status == "idle" {
				c2.status = "goingDown"
				status()
				elevC2()

				// userFloor == elevFloor && status == idle | c3
			} else if userFloor == c3.floor && c3.status == "idle" {
				c3.status = "goingDown"
				status()
				elevC3()

				// userFloor == elevFloor && status == idle | c4
			} else if userFloor == c4.floor && c4.status == "idle" {
				c4.status = "goingDown"
				status()
				elevC4()

				// userFloor == elevFloor && status == idle | c5
			} else if userFloor == c5.floor && c5.status == "idle" {
				c5.status = "goingDown"
				status()
				elevC5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					c1.status = "goingDown"
					status()
					elevC1()

				} else if w >= v && w >= x && w >= y && w >= z {
					c2.status = "goingDown"
					status()
					elevC2()

				} else if x >= v && x >= w && x >= y && x >= z {
					c3.status = "goingDown"
					status()
					elevC3()

				} else if y >= v && y >= w && y >= x && y >= z {
					c4.status = "goingDown"
					status()
					elevC4()

				} else if z >= v && z >= w && z >= x && z >= y {
					c5.status = "goingDown"
					status()
					elevC5()

				} else {
					c1.status = "goingDown"
					status()
					elevC1()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= x && w >= y {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= v && x >= w && x >= y {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= v && y >= w && y >= x {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= x && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= v && x >= w && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if z >= v && z >= w && z >= x {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= y && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if y >= v && y >= w && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= v && z >= w && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if x >= v && x >= y && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= v && y >= x && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= v && z >= x && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= w && x >= y && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= w && y >= x && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= w && z >= x && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= x {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= v && x >= w {
						c3.status = "goingDown"
						status()
						elevC3()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= y {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if y >= v && y >= w {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if z >= v && z >= w {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if x >= v && x >= y {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= v && y >= x {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if x >= v && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if z >= v && z >= x {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if y >= v && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= v && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= w && x >= y {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= w && y >= x {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= w && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if z >= w && z >= x {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if y >= w && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= w && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= x && y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= x && z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c3.status = "goingDown"
						status()
						elevC3()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < c1.floor && c1.status == "idle" && userFloor < c2.floor && c2.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c1.floor && c1.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c3.floor && c3.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c2.floor && c2.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c4.floor && c4.status == "idle" || userFloor < c3.floor && c3.status == "idle" && userFloor < c5.floor && c5.status == "idle" || userFloor < c4.floor && c4.status == "idle" && userFloor < c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if w >= v {
						c2.status = "goingDown"
						status()
						elevC2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if x >= v {
						c3.status = "goingDown"
						status()
						elevC3()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if y >= v {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						c1.status = "goingDown"
						status()
						elevC1()

					} else if z >= v {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c1.status = "goingDown"
						status()
						elevC1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if x >= w {
						c3.status = "goingDown"
						status()
						elevC3()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if y >= w {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						c2.status = "goingDown"
						status()
						elevC2()

					} else if z >= w {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c2.status = "goingDown"
						status()
						elevC2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if y >= x {
						c4.status = "goingDown"
						status()
						elevC4()

					} else {
						c3.status = "goingDown"
						status()
						elevC3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						c3.status = "goingDown"
						status()
						elevC3()

					} else if z >= x {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c3.status = "goingDown"
						status()
						elevC3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						c4.status = "goingDown"
						status()
						elevC4()

					} else if z >= y {
						c5.status = "goingDown"
						status()
						elevC5()

					} else {
						c4.status = "goingDown"
						status()
						elevC4()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < c1.floor && c1.status == "idle" {
				c1.status = "goingDown"
				status()
				elevC1()

			} else if userFloor < c2.floor && c2.status == "idle" {
				c2.status = "goingDown"
				status()
				elevC2()

			} else if userFloor < c3.floor && c3.status == "idle" {
				c3.status = "goingDown"
				status()
				elevC3()

			} else if userFloor < c4.floor && c4.status == "idle" {
				c4.status = "goingDown"
				status()
				elevC4()

			} else if userFloor < c5.floor && c5.status == "idle" {
				c5.status = "goingDown"
				status()
				elevC5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					c1.status = "goingDown"
					status()
					elevC1v2()

				} else if w <= v && w <= x && w <= y && w <= z {
					c2.status = "goingDown"
					status()
					elevC2v2()

				} else if x <= v && x <= w && x <= y && x <= z {
					c3.status = "goingDown"
					status()
					elevC3v2()

				} else if y <= v && y <= w && y <= x && y <= z {
					c4.status = "goingDown"
					status()
					elevC4v2()

				} else if z <= v && z <= w && z <= x && z <= y {
					c5.status = "goingDown"
					status()
					elevC5v2()

				} else {
					c1.status = "goingDown"
					status()
					elevC1v2()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= x && w <= y {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= v && x <= w && x <= y {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= v && y <= w && y <= x {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= x && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= v && x <= w && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if z <= v && z <= w && z <= x {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= y && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if y <= v && y <= w && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= v && z <= w && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if x <= v && x <= y && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= v && y <= x && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= v && z <= x && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= w && x <= y && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= w && y <= x && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= w && z <= x && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= x {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= v && x <= w {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= y {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if y <= v && y <= w {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if z <= v && z <= w {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if x <= v && x <= y {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= v && y <= x {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if x <= v && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if z <= v && z <= x {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if y <= v && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= v && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= w && x <= y {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= w && y <= x {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= w && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if z <= w && z <= x {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if y <= w && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= w && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= x && y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= x && z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c3.status = "goingDown"
						status()
						elevC3v2()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > c1.floor && c1.status == "idle" && userFloor > c2.floor && c2.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c1.floor && c1.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c3.floor && c3.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c2.floor && c2.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c4.floor && c4.status == "idle" || userFloor > c3.floor && c3.status == "idle" && userFloor > c5.floor && c5.status == "idle" || userFloor > c4.floor && c4.status == "idle" && userFloor > c5.floor && c5.status == "idle" {
				var v int = (userFloor - c1.floor)
				var w int = (userFloor - c2.floor)
				var x int = (userFloor - c3.floor)
				var y int = (userFloor - c4.floor)
				var z int = (userFloor - c5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if w <= v {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if x <= v {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if y <= v {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						c1.status = "goingDown"
						status()
						elevC1v2()

					} else if z <= v {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c1.status = "goingDown"
						status()
						elevC1v2()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if x <= w {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if y <= w {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						c2.status = "goingDown"
						status()
						elevC2v2()

					} else if z <= w {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c2.status = "goingDown"
						status()
						elevC2v2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if y <= x {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else {
						c3.status = "goingDown"
						status()
						elevC3v2()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						c3.status = "goingDown"
						status()
						elevC3v2()

					} else if z <= x {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c3.status = "goingDown"
						status()
						elevC3v2()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						c4.status = "goingDown"
						status()
						elevC4v2()

					} else if z <= y {
						c5.status = "goingDown"
						status()
						elevC5v2()

					} else {
						c4.status = "goingDown"
						status()
						elevC4v2()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > c1.floor && c1.status == "idle" {
				c1.status = "goingDown"
				status()
				elevC1v2()

			} else if userFloor > c2.floor && c2.status == "idle" {
				c2.status = "goingDown"
				status()
				elevC2v2()

			} else if userFloor > c3.floor && c3.status == "idle" {
				c3.status = "goingDown"
				status()
				elevC3v2()

			} else if userFloor > c4.floor && c4.status == "idle" {
				c4.status = "goingDown"
				status()
				elevC4v2()

			} else if userFloor > c5.floor && c5.status == "idle" {
				c5.status = "goingDown"
				status()
				elevC5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else {

			time.Sleep(time.Second)
			fmt.Printf("please enter valid information\n")

		}
	}

	requestElevD := func(userFloor int, direction string) {
		scanner := bufio.NewScanner(os.Stdin)

		if direction == "up" && userFloor >= 41 && userFloor < 60 || direction == "up" && userFloor == 1 {

			elevD1 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d1.floor)

				for d1.floor < userFloor {
					d1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)
				}

				d1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)

						for d1.floor < requestedFloor {

							d1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d1.floor)

						}

						d1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.status = "idle"
						status()

					}
				}
			}

			elevD2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d2.floor)

				for d2.floor < userFloor {
					d2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)
				}

				d2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)

						for d2.floor < requestedFloor {

							d2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d2.floor)

						}

						d2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.status = "idle"
						status()

					}
				}
			}

			elevD3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d3.floor)

				for d3.floor < userFloor {
					d3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)
				}

				d3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)

						for d3.floor < requestedFloor {

							d3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d3.floor)

						}

						d3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.status = "idle"
						status()

					}
				}
			}

			elevD4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d4.floor)

				for d4.floor < userFloor {
					d4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)
				}

				d4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)

						for d4.floor < requestedFloor {

							d4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d4.floor)

						}

						d4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.status = "idle"
						status()

					}
				}
			}

			elevD5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d5.floor)

				for d5.floor < userFloor {
					d5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)
				}

				d5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)

						for d5.floor < requestedFloor {

							d5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d5.floor)

						}

						d5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.status = "idle"
						status()

					}
				}
			}

			elevD1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d1.floor)

				for d1.floor > userFloor {
					d1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)
				}

				d1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)

						for d1.floor < requestedFloor {

							d1.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d1.floor)

						}

						d1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.status = "idle"
						status()

					}
				}
			}

			elevD2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d2.floor)

				for d2.floor > userFloor {
					d2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)
				}

				d2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)

						for d2.floor < requestedFloor {

							d2.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d2.floor)

						}

						d2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.status = "idle"
						status()

					}
				}
			}

			elevD3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d3.floor)

				for d3.floor > userFloor {
					d3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)
				}

				d3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)

						for d3.floor < requestedFloor {

							d3.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d3.floor)

						}

						d3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.status = "idle"
						status()

					}
				}
			}

			elevD4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d4.floor)

				for d4.floor > userFloor {
					d4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)
				}

				d4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)

						for d4.floor < requestedFloor {

							d4.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d4.floor)

						}

						d4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.status = "idle"
						status()

					}
				}
			}

			elevD5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d5.floor)

				for d5.floor > userFloor {
					d5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)
				}

				d5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)

						for d5.floor < requestedFloor {

							d5.floor++
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d5.floor)

						}

						d5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingUp
			if userFloor == d1.floor && d1.status == "goingUp" {
				elevD1()

			} else if userFloor == d2.floor && d2.status == "goingUp" {
				elevD2()

			} else if userFloor == d3.floor && d3.status == "goingUp" {
				elevD3()

			} else if userFloor == d4.floor && d4.status == "goingUp" {
				elevD4()

			} else if userFloor == d5.floor && d5.status == "goingUp" {
				elevD5()

				// userFloor > elevFloor && status == goingUp | 5 OPTIONS
			} else if userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				if v <= w && v <= x && v <= y && v <= z {
					elevD1()

				} else if w <= v && w <= x && w <= y && w <= z {
					elevD2()

				} else if x <= v && x <= w && x <= y && x <= z {
					elevD3()

				} else if y <= v && y <= w && y <= x && y <= z {
					elevD4()

				} else if z <= v && z <= w && z <= x && z <= y {
					elevD5()

				} else {
					elevD1()

				}

				// userFloor > elevFloor && status == goingUp | 4 OPTIONS
			} else if userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						elevD1()

					} else if w <= v && w <= x && w <= y {
						elevD2()

					} else if x <= v && x <= w && x <= y {
						elevD3()

					} else if y <= v && y <= w && y <= x {
						elevD4()

					} else {
						elevD1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						elevD1()

					} else if w <= v && w <= x && w <= z {
						elevD2()

					} else if x <= v && x <= w && x <= z {
						elevD3()

					} else if z <= v && z <= w && z <= x {
						elevD5()

					} else {
						elevD1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						elevD1()

					} else if w <= v && w <= y && w <= z {
						elevD2()

					} else if y <= v && y <= w && y <= z {
						elevD4()

					} else if z <= v && z <= w && z <= y {
						elevD5()

					} else {
						elevD1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						elevD1()

					} else if x <= v && x <= y && x <= z {
						elevD3()

					} else if y <= v && y <= x && y <= z {
						elevD4()

					} else if z <= v && z <= x && z <= y {
						elevD5()

					} else {
						elevD1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						elevD2()

					} else if x <= w && x <= y && x <= z {
						elevD3()

					} else if y <= w && y <= x && y <= z {
						elevD4()

					} else if z <= w && z <= x && z <= y {
						elevD5()

					} else {
						elevD2()

					}
				}

				// userFloor > elevFloor && status == goingUp | 3 OPTIONS
			} else if userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						elevD1()

					} else if w <= v && w <= x {
						elevD2()

					} else if x <= v && x <= w {
						elevD3()

					} else {
						elevD1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						elevD1()

					} else if w <= v && w <= y {
						elevD2()

					} else if y <= v && y <= w {
						elevD4()

					} else {
						elevD1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						elevD1()

					} else if w <= v && w <= z {
						elevD2()

					} else if z <= v && z <= w {
						elevD5()

					} else {
						elevD1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						elevD1()

					} else if x <= v && x <= y {
						elevD3()

					} else if y <= v && y <= x {
						elevD4()

					} else {
						elevD1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						elevD1()

					} else if x <= v && x <= z {
						elevD3()

					} else if z <= v && z <= x {
						elevD5()

					} else {
						elevD1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						elevD1()

					} else if y <= v && y <= z {
						elevD4()

					} else if z <= v && z <= y {
						elevD5()

					} else {
						elevD1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						elevD2()

					} else if x <= w && x <= y {
						elevD3()

					} else if y <= w && y <= x {
						elevD4()

					} else {
						elevD2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						elevD2()

					} else if x <= w && x <= z {
						elevD3()

					} else if z <= w && z <= x {
						elevD4()

					} else {
						elevD2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						elevD2()

					} else if y <= w && y <= z {
						elevD4()

					} else if z <= w && z <= y {
						elevD5()

					} else {
						elevD2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						elevD3()

					} else if y <= x && y <= z {
						elevD4()

					} else if z <= x && z <= y {
						elevD5()

					} else {
						elevD3()

					}
				}

				// userFloor > elevFloor && status == goingUp | 2 OPTIONS
			} else if userFloor > d1.floor && d1.status == "goingUp" && userFloor > d2.floor && d2.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d1.floor && d1.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d3.floor && d3.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d2.floor && d2.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d3.floor && d3.status == "goingUp" && userFloor > d4.floor && d4.status == "goingUp" || userFloor > d3.floor && d3.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" || userFloor > d4.floor && d4.status == "goingUp" && userFloor > d5.floor && d5.status == "goingUp" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v\n);
				// fmt.Printf(w\n);
				// fmt.Printf(x\n);
				// fmt.Printf(y\n);
				// fmt.Printf(z\n);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						elevD1()

					} else if w <= v {
						elevD2()

					} else {
						elevD1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						elevD1()

					} else if x <= v {
						elevD3()

					} else {
						elevD1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						elevD1()

					} else if y <= v {
						elevD4()

					} else {
						elevD1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						elevD1()

					} else if z <= v {
						elevD5()

					} else {
						elevD1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						elevD2()

					} else if x <= w {
						elevD3()

					} else {
						elevD2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						elevD2()

					} else if y <= w {
						elevD4()

					} else {
						elevD2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						elevD2()

					} else if z <= w {
						elevD5()

					} else {
						elevD2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						elevD3()

					} else if y <= x {
						elevD4()

					} else {
						elevD3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						elevD3()

					} else if z <= x {
						elevD5()

					} else {
						elevD3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						elevD4()

					} else if z <= y {
						elevD5()

					} else {
						elevD4()

					}
				}

				// userFloor > elevFloor && status == goingUp | 1 OPTION
			} else if userFloor > d1.floor && d1.status == "goingUp" {
				elevD1()

			} else if userFloor > d2.floor && d2.status == "goingUp" {
				elevD2()

			} else if userFloor > d3.floor && d3.status == "goingUp" {
				elevD3()

			} else if userFloor > d4.floor && d4.status == "goingUp" {
				elevD4()

			} else if userFloor > d5.floor && d5.status == "goingUp" {
				elevD5()

				// userFloor == elevFloor && status == idle | d1
			} else if userFloor == d1.floor && d1.status == "idle" {
				d1.status = "goingUp"
				status()
				elevD1()

				// userFloor == elevFloor && status == idle | d2
			} else if userFloor == d2.floor && d2.status == "idle" {
				d2.status = "goingUp"
				status()
				elevD2()

				// userFloor == elevFloor && status == idle | d3
			} else if userFloor == d3.floor && d3.status == "idle" {
				d3.status = "goingUp"
				status()
				elevD3()

				// userFloor == elevFloor && status == idle | d4
			} else if userFloor == d4.floor && d4.status == "idle" {
				d4.status = "goingUp"
				status()
				elevD4()

				// userFloor == elevFloor && status == idle | d5
			} else if userFloor == d5.floor && d5.status == "idle" {
				d5.status = "goingUp"
				status()
				elevD5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					d1.status = "goingUp"
					status()
					elevD1()

				} else if w <= v && w <= x && w <= y && w <= z {
					d2.status = "goingUp"
					status()
					elevD2()

				} else if x <= v && x <= w && x <= y && x <= z {
					d3.status = "goingUp"
					status()
					elevD3()

				} else if y <= v && y <= w && y <= x && y <= z {
					d4.status = "goingUp"
					status()
					elevD4()

				} else if z <= v && z <= w && z <= x && z <= y {
					d5.status = "goingUp"
					status()
					elevD5()

				} else {
					d1.status = "goingUp"
					status()
					elevD1()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= x && w <= y {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= v && x <= w && x <= y {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= v && y <= w && y <= x {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= x && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= v && x <= w && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if z <= v && z <= w && z <= x {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= y && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if y <= v && y <= w && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= v && z <= w && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if x <= v && x <= y && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= v && y <= x && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= v && z <= x && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= w && x <= y && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= w && y <= x && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= w && z <= x && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= x {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= v && x <= w {
						d3.status = "goingUp"
						status()
						elevD3()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= y {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if y <= v && y <= w {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if z <= v && z <= w {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if x <= v && x <= y {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= v && y <= x {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if x <= v && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if z <= v && z <= x {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if y <= v && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= v && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= w && x <= y {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= w && y <= x {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= w && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if z <= w && z <= x {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if y <= w && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= w && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= x && y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= x && z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d3.status = "goingUp"
						status()
						elevD3()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if w <= v {
						d2.status = "goingUp"
						status()
						elevD2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if x <= v {
						d3.status = "goingUp"
						status()
						elevD3()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if y <= v {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						d1.status = "goingUp"
						status()
						elevD1()

					} else if z <= v {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d1.status = "goingUp"
						status()
						elevD1()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if x <= w {
						d3.status = "goingUp"
						status()
						elevD3()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if y <= w {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						d2.status = "goingUp"
						status()
						elevD2()

					} else if z <= w {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d2.status = "goingUp"
						status()
						elevD2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if y <= x {
						d4.status = "goingUp"
						status()
						elevD4()

					} else {
						d3.status = "goingUp"
						status()
						elevD3()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						d3.status = "goingUp"
						status()
						elevD3()

					} else if z <= x {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d3.status = "goingUp"
						status()
						elevD3()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						d4.status = "goingUp"
						status()
						elevD4()

					} else if z <= y {
						d5.status = "goingUp"
						status()
						elevD5()

					} else {
						d4.status = "goingUp"
						status()
						elevD4()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > d1.floor && d1.status == "idle" {
				d1.status = "goingUp"
				status()
				elevD1()

			} else if userFloor > d2.floor && d2.status == "idle" {
				d2.status = "goingUp"
				status()
				elevD2()

			} else if userFloor > d3.floor && d3.status == "idle" {
				d3.status = "goingUp"
				status()
				elevD3()

			} else if userFloor > d4.floor && d4.status == "idle" {
				d4.status = "goingUp"
				status()
				elevD4()

			} else if userFloor > d5.floor && d5.status == "idle" {
				d5.status = "goingUp"
				status()
				elevD5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					d1.status = "goingUp"
					status()
					elevD1v2()

				} else if w >= v && w >= x && w >= y && w >= z {
					d2.status = "goingUp"
					status()
					elevD2v2()

				} else if x >= v && x >= w && x >= y && x >= z {
					d3.status = "goingUp"
					status()
					elevD3v2()

				} else if y >= v && y >= w && y >= x && y >= z {
					d4.status = "goingUp"
					status()
					elevD4v2()

				} else if z >= v && z >= w && z >= x && z >= y {
					d5.status = "goingUp"
					status()
					elevD5v2()

				} else {
					d1.status = "goingUp"
					status()
					elevD1v2()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= x && w >= y {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= v && x >= w && x >= y {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= v && y >= w && y >= x {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= x && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= v && x >= w && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if z >= v && z >= w && z >= x {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= y && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if y >= v && y >= w && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= v && z >= w && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if x >= v && x >= y && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= v && y >= x && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= v && z >= x && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= w && x >= y && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= w && y >= x && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= w && z >= x && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= x {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= v && x >= w {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= y {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if y >= v && y >= w {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if z >= v && z >= w {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if x >= v && x >= y {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= v && y >= x {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if x >= v && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if z >= v && z >= x {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if y >= v && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= v && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= w && x >= y {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= w && y >= x {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= w && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if z >= w && z >= x {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if y >= w && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= w && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= x && y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= x && z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d3.status = "goingUp"
						status()
						elevD3v2()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if w >= v {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if x >= v {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if y >= v {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						d1.status = "goingUp"
						status()
						elevD1v2()

					} else if z >= v {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d1.status = "goingUp"
						status()
						elevD1v2()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if x >= w {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if y >= w {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						d2.status = "goingUp"
						status()
						elevD2v2()

					} else if z >= w {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d2.status = "goingUp"
						status()
						elevD2v2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if y >= x {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else {
						d3.status = "goingUp"
						status()
						elevD3v2()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						d3.status = "goingUp"
						status()
						elevD3v2()

					} else if z >= x {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d3.status = "goingUp"
						status()
						elevD3v2()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						d4.status = "goingUp"
						status()
						elevD4v2()

					} else if z >= y {
						d5.status = "goingUp"
						status()
						elevD5v2()

					} else {
						d4.status = "goingUp"
						status()
						elevD4v2()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < d1.floor && d1.status == "idle" {
				d1.status = "goingUp"
				status()
				elevD1v2()

			} else if userFloor < d2.floor && d2.status == "idle" {
				d2.status = "goingUp"
				status()
				elevD2v2()

			} else if userFloor < d3.floor && d3.status == "idle" {
				d3.status = "goingUp"
				status()
				elevD3v2()

			} else if userFloor < d4.floor && d4.status == "idle" {
				d4.status = "goingUp"
				status()
				elevD4v2()

			} else if userFloor < d5.floor && d5.status == "idle" {
				d5.status = "goingUp"
				status()
				elevD5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else if direction == "down" && userFloor >= 41 && userFloor <= 60 {

			elevD1 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator d1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d1.floor)

				for d1.floor > userFloor {
					d1.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)
				}

				d1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)

						for d1.floor > requestedFloor {

							d1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d1.floor)

						}

						d1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.status = "idle"
						status()

					}
				}
			}

			elevD2 := func() {

				time.Sleep(time.Second)
				fmt.Printf("elevator d2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d2.floor)

				for d2.floor > userFloor {
					d2.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)
				}

				d2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)

						for d2.floor > requestedFloor {

							d2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d2.floor)

						}

						d2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.status = "idle"
						status()

					}
				}
			}

			elevD3 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d3.floor)

				for d3.floor > userFloor {
					d3.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)
				}

				d3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)

						for d3.floor > requestedFloor {

							d3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d3.floor)

						}

						d3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.status = "idle"
						status()

					}
				}
			}

			elevD4 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d4.floor)

				for d4.floor > userFloor {
					d4.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)
				}

				d4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)

						for d4.floor > requestedFloor {

							d4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d4.floor)

						}

						d4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.status = "idle"
						status()

					}
				}
			}

			elevD5 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d5.floor)

				for d5.floor > userFloor {
					d5.floor--
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)
				}

				d5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)

						for d5.floor > requestedFloor {

							d5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d5.floor)

						}

						d5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.status = "idle"
						status()

					}
				}
			}

			elevD1v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d1\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d1.floor)

				for d1.floor < userFloor {
					d1.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)
				}

				d1.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d1.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)

						for d1.floor > requestedFloor {

							d1.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d1.floor)

						}

						d1.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d1.door)
						d1.status = "idle"
						status()

					}
				}
			}

			elevD2v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d2\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d2.floor)

				for d2.floor < userFloor {
					d2.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)
				}

				d2.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d2.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)

						for d2.floor > requestedFloor {

							d2.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d2.floor)

						}

						d2.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d2.door)
						d2.status = "idle"
						status()

					}
				}
			}

			elevD3v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d3\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d3.floor)

				for d3.floor < userFloor {
					d3.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)
				}

				d3.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d3.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)

						for d3.floor > requestedFloor {

							d3.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d3.floor)

						}

						d3.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d3.door)
						d3.status = "idle"
						status()

					}
				}
			}

			elevD4v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d4\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d4.floor)

				for d4.floor < userFloor {
					d4.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)
				}

				d4.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d4.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)

						for d4.floor > requestedFloor {

							d4.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d4.floor)

						}

						d4.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d4.door)
						d4.status = "idle"
						status()

					}
				}
			}

			elevD5v2 := func() {
				time.Sleep(time.Second)
				fmt.Printf("elevator d5\n")
				time.Sleep(time.Second)
				fmt.Printf("elevator's floor: %d\n", d5.floor)

				for d5.floor < userFloor {
					d5.floor++
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)
				}

				d5.door = "opened"
				time.Sleep(time.Second)
				fmt.Printf("door: %v\n", d5.door)
				time.Sleep(time.Second)
				fmt.Printf("which floor would u like to go to?\n")

				for input := true; input == true; {

					scanner.Scan()
					requestedFloor, _ := strconv.Atoi(scanner.Text())

					if 1 < requestedFloor && requestedFloor < 41 || requestedFloor < 1 || requestedFloor >= 60 || requestedFloor >= userFloor {

						time.Sleep(time.Second)
						fmt.Printf("please select a valid floor\n")

					} else {

						input = false

						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)

						for d5.floor > requestedFloor {

							d5.floor--
							time.Sleep(time.Second)
							fmt.Printf("floor display: %d\n", d5.floor)

						}

						d5.door = "opened"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.door = "closed"
						time.Sleep(time.Second)
						fmt.Printf("door: %v\n", d5.door)
						d5.status = "idle"
						status()

					}
				}
			}

			// userfloor == elevFloor && status == goingDown
			if userFloor == d1.floor && d1.status == "goingDown" {
				elevD1()

			} else if userFloor == d2.floor && d2.status == "goingDown" {
				elevD2()

			} else if userFloor == d3.floor && d3.status == "goingDown" {
				elevD3()

			} else if userFloor == d4.floor && d4.status == "goingDown" {
				elevD4()

			} else if userFloor == d5.floor && d5.status == "goingDown" {
				elevD5()

				// userFloor < elevFloor && status == goingDown | 5 OPTIONS
			} else if userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					elevD1()

				} else if w >= v && w >= x && w >= y && w >= z {
					elevD2()

				} else if x >= v && x >= w && x >= y && x >= z {
					elevD3()

				} else if y >= v && y >= w && y >= x && y >= z {
					elevD4()

				} else if z >= v && z >= w && z >= x && z >= y {
					elevD5()

				} else {
					elevD1()

				}

				// userFloor < elevFloor && status == goingDown | 4 OPTIONS
			} else if userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						elevD1()

					} else if w >= v && w >= x && w >= y {
						elevD2()

					} else if x >= v && x >= w && x >= y {
						elevD3()

					} else if y >= v && y >= w && y >= x {
						elevD4()

					} else {
						elevD1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						elevD1()

					} else if w >= v && w >= x && w >= z {
						elevD2()

					} else if x >= v && x >= w && x >= z {
						elevD3()

					} else if z >= v && z >= w && z >= x {
						elevD5()

					} else {
						elevD1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						elevD1()

					} else if w >= v && w >= y && w >= z {
						elevD2()

					} else if y >= v && y >= w && y >= z {
						elevD4()

					} else if z >= v && z >= w && z >= y {
						elevD5()

					} else {
						elevD1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						elevD1()

					} else if x >= v && x >= y && x >= z {
						elevD3()

					} else if y >= v && y >= x && y >= z {
						elevD4()

					} else if z >= v && z >= x && z >= y {
						elevD5()

					} else {
						elevD1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						elevD2()

					} else if x >= w && x >= y && x >= z {
						elevD3()

					} else if y >= w && y >= x && y >= z {
						elevD4()

					} else if z >= w && z >= x && z >= y {
						elevD5()

					} else {
						elevD2()

					}
				}

				// userFloor < elevFloor && status == goingDown | 3 OPTIONS
			} else if userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						elevD1()

					} else if w >= v && w >= x {
						elevD2()

					} else if x >= v && x >= w {
						elevD3()

					} else {
						elevD1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						elevD1()

					} else if w >= v && w >= y {
						elevD2()

					} else if y >= v && y >= w {
						elevD4()

					} else {
						elevD1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						elevD1()

					} else if w >= v && w >= z {
						elevD2()

					} else if z >= v && z >= w {
						elevD5()

					} else {
						elevD1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						elevD1()

					} else if x >= v && x >= y {
						elevD3()

					} else if y >= v && y >= x {
						elevD4()

					} else {
						elevD1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						elevD1()

					} else if x >= v && x >= z {
						elevD3()

					} else if z >= v && z >= x {
						elevD5()

					} else {
						elevD1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						elevD1()

					} else if y >= v && y >= z {
						elevD4()

					} else if z >= v && z >= y {
						elevD5()

					} else {
						elevD1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						elevD2()

					} else if x >= w && x >= y {
						elevD3()

					} else if y >= w && y >= x {
						elevD4()

					} else {
						elevD2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						elevD2()

					} else if x >= w && x >= z {
						elevD3()

					} else if z >= w && z >= x {
						elevD4()

					} else {
						elevD2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						elevD2()

					} else if y >= w && y >= z {
						elevD4()

					} else if z >= w && z >= y {
						elevD5()

					} else {
						elevD2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						elevD3()

					} else if y >= x && y >= z {
						elevD4()

					} else if z >= x && z >= y {
						elevD5()

					} else {
						elevD3()

					}
				}

				// userFloor < elevFloor && status == goingDown | 2 OPTIONS
			} else if userFloor < d1.floor && d1.status == "goingDown" && userFloor < d2.floor && d2.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d1.floor && d1.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d3.floor && d3.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d2.floor && d2.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d3.floor && d3.status == "goingDown" && userFloor < d4.floor && d4.status == "goingDown" || userFloor < d3.floor && d3.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" || userFloor < d4.floor && d4.status == "goingDown" && userFloor < d5.floor && d5.status == "goingDown" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						elevD1()

					} else if w >= v {
						elevD2()

					} else {
						elevD1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						elevD1()

					} else if x >= v {
						elevD3()

					} else {
						elevD1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						elevD1()

					} else if y >= v {
						elevD4()

					} else {
						elevD1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						elevD1()

					} else if z >= v {
						elevD5()

					} else {
						elevD1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						elevD2()

					} else if x >= w {
						elevD3()

					} else {
						elevD2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						elevD2()

					} else if y >= w {
						elevD4()

					} else {
						elevD2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						elevD2()

					} else if z >= w {
						elevD5()

					} else {
						elevD2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						elevD3()

					} else if y >= x {
						elevD4()

					} else {
						elevD3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						elevD3()

					} else if z >= x {
						elevD5()

					} else {
						elevD3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						elevD4()

					} else if z >= y {
						elevD5()

					} else {
						elevD4()

					}
				}

				// userFloor < elevFloor && status == goingDown | 1 OPTION
			} else if userFloor < d1.floor && d1.status == "goingDown" {
				elevD1()

			} else if userFloor < d2.floor && d2.status == "goingDown" {
				elevD2()

			} else if userFloor < d3.floor && d3.status == "goingDown" {
				elevD3()

			} else if userFloor < d4.floor && d4.status == "goingDown" {
				elevD4()

			} else if userFloor < d5.floor && d5.status == "goingDown" {
				elevD5()

				// userFloor == elevFloor && status == idle | d1
			} else if userFloor == d1.floor && d1.status == "idle" {
				d1.status = "goingDown"
				status()
				elevD1()

				// userFloor == elevFloor && status == idle | d2
			} else if userFloor == d2.floor && d2.status == "idle" {
				d2.status = "goingDown"
				status()
				elevD2()

				// userFloor == elevFloor && status == idle | d3
			} else if userFloor == d3.floor && d3.status == "idle" {
				d3.status = "goingDown"
				status()
				elevD3()

				// userFloor == elevFloor && status == idle | d4
			} else if userFloor == d4.floor && d4.status == "idle" {
				d4.status = "goingDown"
				status()
				elevD4()

				// userFloor == elevFloor && status == idle | d5
			} else if userFloor == d5.floor && d5.status == "idle" {
				d5.status = "goingDown"
				status()
				elevD5()

				// userFloor < elevFloor && status == idle | 5 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v >= w && v >= x && v >= y && v >= z {
					d1.status = "goingDown"
					status()
					elevD1()

				} else if w >= v && w >= x && w >= y && w >= z {
					d2.status = "goingDown"
					status()
					elevD2()

				} else if x >= v && x >= w && x >= y && x >= z {
					d3.status = "goingDown"
					status()
					elevD3()

				} else if y >= v && y >= w && y >= x && y >= z {
					d4.status = "goingDown"
					status()
					elevD4()

				} else if z >= v && z >= w && z >= x && z >= y {
					d5.status = "goingDown"
					status()
					elevD5()

				} else {
					d1.status = "goingDown"
					status()
					elevD1()

				}

				// userFloor < elevFloor && status == idle | 4 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v < 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v >= w && v >= x && v >= y {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= x && w >= y {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= v && x >= w && x >= y {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= v && y >= w && y >= x {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v >= w && v >= x && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= x && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= v && x >= w && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if z >= v && z >= w && z >= x {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v >= w && v >= y && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= y && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if y >= v && y >= w && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= v && z >= w && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v >= x && v >= y && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if x >= v && x >= y && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= v && y >= x && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= v && z >= x && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z < 0 {
					if w >= x && w >= y && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= w && x >= y && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= w && y >= x && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= w && z >= x && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}
				}

				// userFloor < elevFloor && status == idle | 3 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v >= w && v >= x {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= x {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= v && x >= w {
						d3.status = "goingDown"
						status()
						elevD3()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [2] 3 elevs — v, w, y
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v >= w && v >= y {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= y {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if y >= v && y >= w {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [3] 3 elevs — v, w, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v >= w && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if z >= v && z >= w {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [4] 3 elevs — v, x, y
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v >= x && v >= y {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if x >= v && x >= y {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= v && y >= x {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [5] 3 elevs — v, x, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v >= x && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if x >= v && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if z >= v && z >= x {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [6] 3 elevs — v, y, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v >= y && v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if y >= v && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= v && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [7] 3 elevs — w, x, y
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if w >= x && w >= y {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= w && x >= y {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= w && y >= x {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [8] 3 elevs — w, x, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if w >= x && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= w && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if z >= w && z >= x {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [9] 3 elevs — w, y, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if w >= y && w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if y >= w && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= w && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [10] 3 elevs — x, y, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if x >= y && x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= x && y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= x && z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d3.status = "goingDown"
						status()
						elevD3()

					}
				}

				// userFloor < elevFloor && status == idle | 2 OPTIONS
			} else if userFloor < d1.floor && d1.status == "idle" && userFloor < d2.floor && d2.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d1.floor && d1.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d3.floor && d3.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d2.floor && d2.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d4.floor && d4.status == "idle" || userFloor < d3.floor && d3.status == "idle" && userFloor < d5.floor && d5.status == "idle" || userFloor < d4.floor && d4.status == "idle" && userFloor < d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v >= w {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if w >= v {
						d2.status = "goingDown"
						status()
						elevD2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [2] 2 elevs — v, x
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v >= x {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if x >= v {
						d3.status = "goingDown"
						status()
						elevD3()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [3] 2 elevs — v, y
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v >= y {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if y >= v {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [4] 2 elevs — v, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v >= z {
						d1.status = "goingDown"
						status()
						elevD1()

					} else if z >= v {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d1.status = "goingDown"
						status()
						elevD1()

					}

					// [5] 2 elevs — w, x
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if w >= x {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if x >= w {
						d3.status = "goingDown"
						status()
						elevD3()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [6] 2 elevs — w, y
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if w >= y {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if y >= w {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [7] 2 elevs — w, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if w >= z {
						d2.status = "goingDown"
						status()
						elevD2()

					} else if z >= w {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d2.status = "goingDown"
						status()
						elevD2()

					}

					// [8] 2 elevs — x, y
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if x >= y {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if y >= x {
						d4.status = "goingDown"
						status()
						elevD4()

					} else {
						d3.status = "goingDown"
						status()
						elevD3()

					}

					// [9] 2 elevs — x, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if x >= z {
						d3.status = "goingDown"
						status()
						elevD3()

					} else if z >= x {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d3.status = "goingDown"
						status()
						elevD3()

					}

					// [10] 2 elevs — y, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if y >= z {
						d4.status = "goingDown"
						status()
						elevD4()

					} else if z >= y {
						d5.status = "goingDown"
						status()
						elevD5()

					} else {
						d4.status = "goingDown"
						status()
						elevD4()

					}
				}

				// userFloor < elevFloor && status == idle | 1 OPTION
			} else if userFloor < d1.floor && d1.status == "idle" {
				d1.status = "goingDown"
				status()
				elevD1()

			} else if userFloor < d2.floor && d2.status == "idle" {
				d2.status = "goingDown"
				status()
				elevD2()

			} else if userFloor < d3.floor && d3.status == "idle" {
				d3.status = "goingDown"
				status()
				elevD3()

			} else if userFloor < d4.floor && d4.status == "idle" {
				d4.status = "goingDown"
				status()
				elevD4()

			} else if userFloor < d5.floor && d5.status == "idle" {
				d5.status = "goingDown"
				status()
				elevD5()

				// userFloor > elevFloor && status == idle | 5 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				if v <= w && v <= x && v <= y && v <= z {
					d1.status = "goingDown"
					status()
					elevD1v2()

				} else if w <= v && w <= x && w <= y && w <= z {
					d2.status = "goingDown"
					status()
					elevD2v2()

				} else if x <= v && x <= w && x <= y && x <= z {
					d3.status = "goingDown"
					status()
					elevD3v2()

				} else if y <= v && y <= w && y <= x && y <= z {
					d4.status = "goingDown"
					status()
					elevD4v2()

				} else if z <= v && z <= w && z <= x && z <= y {
					d5.status = "goingDown"
					status()
					elevD5v2()

				} else {
					d1.status = "goingDown"
					status()
					elevD1v2()

				}

				// userFloor > elevFloor && status == idle | 4 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 4 elevs — v, w, x, y
				if v > 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if v <= w && v <= x && v <= y {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= x && w <= y {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= v && x <= w && x <= y {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= v && y <= w && y <= x {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [2] 4 elevs — v, w, x, z
				} else if v > 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if v <= w && v <= x && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= x && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= v && x <= w && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if z <= v && z <= w && z <= x {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [3] 4 elevs — v, w, y, z
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if v <= w && v <= y && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= y && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if y <= v && y <= w && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= v && z <= w && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [4] 4 elevs — v, x, y, z
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if v <= x && v <= y && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if x <= v && x <= y && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= v && y <= x && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= v && z <= x && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [5] 4 elevs — w, x, y, z
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z > 0 {
					if w <= x && w <= y && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= w && x <= y && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= w && y <= x && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= w && z <= x && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}
				}

				// userFloor > elevFloor && status == idle | 3 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 3 elevs — v, w, x
				if v > 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if v <= w && v <= x {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= x {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= v && x <= w {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [2] 3 elevs — v, w, y
				} else if v > 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if v <= w && v <= y {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= y {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if y <= v && y <= w {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [3] 3 elevs — v, w, z
				} else if v > 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if v <= w && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if z <= v && z <= w {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [4] 3 elevs — v, x, y
				} else if v > 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if v <= x && v <= y {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if x <= v && x <= y {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= v && y <= x {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [5] 3 elevs — v, x, z
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if v <= x && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if x <= v && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if z <= v && z <= x {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [6] 3 elevs — v, y, z
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if v <= y && v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if y <= v && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= v && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [7] 3 elevs — w, x, y
				} else if v < 0 && w > 0 && x > 0 && y > 0 && z < 0 {
					if w <= x && w <= y {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= w && x <= y {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= w && y <= x {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [8] 3 elevs — w, x, z
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z > 0 {
					if w <= x && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= w && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if z <= w && z <= x {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [9] 3 elevs — w, y, z
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z > 0 {
					if w <= y && w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if y <= w && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= w && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [10] 3 elevs — x, y, z
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z > 0 {
					if x <= y && x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= x && y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= x && z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d3.status = "goingDown"
						status()
						elevD3v2()

					}
				}

				// userFloor > elevFloor && status == idle | 2 OPTIONS
			} else if userFloor > d1.floor && d1.status == "idle" && userFloor > d2.floor && d2.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d1.floor && d1.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d3.floor && d3.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d2.floor && d2.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d4.floor && d4.status == "idle" || userFloor > d3.floor && d3.status == "idle" && userFloor > d5.floor && d5.status == "idle" || userFloor > d4.floor && d4.status == "idle" && userFloor > d5.floor && d5.status == "idle" {
				var v int = (userFloor - d1.floor)
				var w int = (userFloor - d2.floor)
				var x int = (userFloor - d3.floor)
				var y int = (userFloor - d4.floor)
				var z int = (userFloor - d5.floor)
				// fmt.Printf(v);
				// fmt.Printf(w);
				// fmt.Printf(x);
				// fmt.Printf(y);
				// fmt.Printf(z);

				// [1] 2 elevs — v, w
				if v > 0 && w > 0 && x < 0 && y < 0 && z < 0 {
					if v <= w {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if w <= v {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [2] 2 elevs — v, x
				} else if v > 0 && w < 0 && x > 0 && y < 0 && z < 0 {
					if v <= x {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if x <= v {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [3] 2 elevs — v, y
				} else if v > 0 && w < 0 && x < 0 && y > 0 && z < 0 {
					if v <= y {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if y <= v {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [4] 2 elevs — v, z
				} else if v > 0 && w < 0 && x < 0 && y < 0 && z > 0 {
					if v <= z {
						d1.status = "goingDown"
						status()
						elevD1v2()

					} else if z <= v {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d1.status = "goingDown"
						status()
						elevD1v2()

					}

					// [5] 2 elevs — w, x
				} else if v < 0 && w > 0 && x > 0 && y < 0 && z < 0 {
					if w <= x {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if x <= w {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [6] 2 elevs — w, y
				} else if v < 0 && w > 0 && x < 0 && y > 0 && z < 0 {
					if w <= y {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if y <= w {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [7] 2 elevs — w, z
				} else if v < 0 && w > 0 && x < 0 && y < 0 && z > 0 {
					if w <= z {
						d2.status = "goingDown"
						status()
						elevD2v2()

					} else if z <= w {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d2.status = "goingDown"
						status()
						elevD2v2()

					}

					// [8] 2 elevs — x, y
				} else if v < 0 && w < 0 && x > 0 && y > 0 && z < 0 {
					if x <= y {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if y <= x {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else {
						d3.status = "goingDown"
						status()
						elevD3v2()

					}

					// [9] 2 elevs — x, z
				} else if v < 0 && w < 0 && x > 0 && y < 0 && z > 0 {
					if x <= z {
						d3.status = "goingDown"
						status()
						elevD3v2()

					} else if z <= x {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d3.status = "goingDown"
						status()
						elevD3v2()

					}

					// [10] 2 elevs — y, z
				} else if v < 0 && w < 0 && x < 0 && y > 0 && z > 0 {
					if y <= z {
						d4.status = "goingDown"
						status()
						elevD4v2()

					} else if z <= y {
						d5.status = "goingDown"
						status()
						elevD5v2()

					} else {
						d4.status = "goingDown"
						status()
						elevD4v2()

					}
				}

				// userFloor > elevFloor && status == idle | 1 OPTION
			} else if userFloor > d1.floor && d1.status == "idle" {
				d1.status = "goingDown"
				status()
				elevD1v2()

			} else if userFloor > d2.floor && d2.status == "idle" {
				d2.status = "goingDown"
				status()
				elevD2v2()

			} else if userFloor > d3.floor && d3.status == "idle" {
				d3.status = "goingDown"
				status()
				elevD3v2()

			} else if userFloor > d4.floor && d4.status == "idle" {
				d4.status = "goingDown"
				status()
				elevD4v2()

			} else if userFloor > d5.floor && d5.status == "idle" {
				d5.status = "goingDown"
				status()
				elevD5v2()

			} else {

				time.Sleep(time.Second)
				fmt.Printf("all elevators are busy, please try again in a few moments\n")

			}

		} else {

			time.Sleep(time.Second)
			fmt.Printf("please enter valid information\n")

		}
	}

	var userFloor int

	requestFloorA := func(elev int, requestedFloor int) {

		if elev == 1 {

			if userFloor < requestedFloor {

				if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)

					a1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					time.Sleep(time.Second)
					a1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)

					for a1.floor < requestedFloor {

						a1.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a1.floor)

					}

					a1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					a1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					a1.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a1.floor)

					a1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					time.Sleep(time.Second)
					a1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)

					for a1.floor > requestedFloor {

						a1.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a1.floor)

					}

					a1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					a1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a1.door)
					a1.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 2 {

			if userFloor < requestedFloor {

				if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)

					a2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					time.Sleep(time.Second)
					a2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)

					for a2.floor < requestedFloor {

						a2.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a2.floor)

					}

					a2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					a2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					a2.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a2.floor)

					a2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					time.Sleep(time.Second)
					a2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)

					for a2.floor > requestedFloor {

						a2.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a2.floor)

					}

					a2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					a2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a2.door)
					a2.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 3 {

			if userFloor < requestedFloor {

				if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)

					a3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					time.Sleep(time.Second)
					a3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)

					for a3.floor < requestedFloor {

						a3.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a3.floor)

					}

					a3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					a3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					a3.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a3.floor)

					a3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					time.Sleep(time.Second)
					a3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)

					for a3.floor > requestedFloor {

						a3.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a3.floor)

					}

					a3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					a3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a3.door)
					a3.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 4 {

			if userFloor < requestedFloor {

				if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)

					a4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					time.Sleep(time.Second)
					a4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)

					for a4.floor < requestedFloor {

						a4.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a4.floor)

					}

					a4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					a4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					a4.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a4.floor)

					a4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					time.Sleep(time.Second)
					a4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)

					for a4.floor > requestedFloor {

						a4.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a4.floor)

					}

					a4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					a4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a4.door)
					a4.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 5 {

			if userFloor < requestedFloor {

				if requestedFloor <= -5 || requestedFloor > 1 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)

					a5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					time.Sleep(time.Second)
					a5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)

					for a5.floor < requestedFloor {

						a5.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a5.floor)

					}

					a5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					a5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					a5.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator a5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", a5.floor)

					a5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					time.Sleep(time.Second)
					a5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)

					for a5.floor > requestedFloor {

						a5.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", a5.floor)

					}

					a5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					a5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", a5.door)
					a5.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		}
	}

	requestFloorB := func(elev int, requestedFloor int) {

		if elev == 1 {

			if userFloor < requestedFloor {

				if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)

					b1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					time.Sleep(time.Second)
					b1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)

					for b1.floor < requestedFloor {

						b1.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b1.floor)

					}

					b1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					b1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					b1.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b1.floor)

					b1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					time.Sleep(time.Second)
					b1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)

					for b1.floor > requestedFloor {

						b1.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b1.floor)

					}

					b1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					b1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b1.door)
					b1.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 2 {

			if userFloor < requestedFloor {

				if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)

					b2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					time.Sleep(time.Second)
					b2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)

					for b2.floor < requestedFloor {

						b2.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b2.floor)

					}

					b2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					b2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					b2.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b2.floor)

					b2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					time.Sleep(time.Second)
					b2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)

					for b2.floor > requestedFloor {

						b2.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b2.floor)

					}

					b2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					b2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b2.door)
					b2.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 3 {

			if userFloor < requestedFloor {

				if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)

					b3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					time.Sleep(time.Second)
					b3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)

					for b3.floor < requestedFloor {

						b3.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b3.floor)

					}

					b3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					b3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					b3.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b3.floor)

					b3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					time.Sleep(time.Second)
					b3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)

					for b3.floor > requestedFloor {

						b3.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b3.floor)

					}

					b3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					b3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b3.door)
					b3.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 4 {

			if userFloor < requestedFloor {

				if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)

					b4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					time.Sleep(time.Second)
					b4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)

					for b4.floor < requestedFloor {

						b4.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b4.floor)

					}

					b4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					b4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					b4.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b4.floor)

					b4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					time.Sleep(time.Second)
					b4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)

					for b4.floor > requestedFloor {

						b4.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b4.floor)

					}

					b4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					b4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b4.door)
					b4.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 5 {

			if userFloor < requestedFloor {

				if requestedFloor <= 1 || requestedFloor > 20 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)

					b5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					time.Sleep(time.Second)
					b5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)

					for b5.floor < requestedFloor {

						b5.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b5.floor)

					}

					b5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					b5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					b5.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator b5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", b5.floor)

					b5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					time.Sleep(time.Second)
					b5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)

					for b5.floor > requestedFloor {

						b5.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", b5.floor)

					}

					b5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					b5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", b5.door)
					b5.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		}
	}

	requestFloorC := func(elev int, requestedFloor int) {

		if elev == 1 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)

					c1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					time.Sleep(time.Second)
					c1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)

					for c1.floor < requestedFloor {

						c1.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c1.floor)

					}

					c1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					c1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					c1.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c1.floor)

					c1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					time.Sleep(time.Second)
					c1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)

					for c1.floor > requestedFloor {

						c1.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c1.floor)

					}

					c1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					c1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c1.door)
					c1.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 2 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)

					c2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					time.Sleep(time.Second)
					c2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)

					for c2.floor < requestedFloor {

						c2.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c2.floor)

					}

					c2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					c2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					c2.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c2.floor)

					c2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					time.Sleep(time.Second)
					c2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)

					for c2.floor > requestedFloor {

						c2.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c2.floor)

					}

					c2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					c2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c2.door)
					c2.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 3 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)

					c3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					time.Sleep(time.Second)
					c3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)

					for c3.floor < requestedFloor {

						c3.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c3.floor)

					}

					c3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					c3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					c3.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c3.floor)

					c3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					time.Sleep(time.Second)
					c3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)

					for c3.floor > requestedFloor {

						c3.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c3.floor)

					}

					c3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					c3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c3.door)
					c3.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 4 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)

					c4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					time.Sleep(time.Second)
					c4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)

					for c4.floor < requestedFloor {

						c4.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c4.floor)

					}

					c4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					c4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					c4.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c4.floor)

					c4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					time.Sleep(time.Second)
					c4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)

					for c4.floor > requestedFloor {

						c4.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c4.floor)

					}

					c4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					c4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c4.door)
					c4.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 5 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 21 || requestedFloor > 40 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)

					c5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					time.Sleep(time.Second)
					c5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)

					for c5.floor < requestedFloor {

						c5.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c5.floor)

					}

					c5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					c5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					c5.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator c5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", c5.floor)

					c5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					time.Sleep(time.Second)
					c5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)

					for c5.floor > requestedFloor {

						c5.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", c5.floor)

					}

					c5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					c5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", c5.door)
					c5.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		}
	}

	requestFloorD := func(elev int, requestedFloor int) {

		if elev == 1 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)

					d1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					time.Sleep(time.Second)
					d1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)

					for d1.floor < requestedFloor {

						d1.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d1.floor)

					}

					d1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					d1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					d1.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d1\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d1.floor)

					d1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					time.Sleep(time.Second)
					d1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)

					for d1.floor > requestedFloor {

						d1.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d1.floor)

					}

					d1.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					d1.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d1.door)
					d1.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 2 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)

					d2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					time.Sleep(time.Second)
					d2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)

					for d2.floor < requestedFloor {

						d2.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d2.floor)

					}

					d2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					d2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					d2.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d2\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d2.floor)

					d2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					time.Sleep(time.Second)
					d2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)

					for d2.floor > requestedFloor {

						d2.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d2.floor)

					}

					d2.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					d2.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d2.door)
					d2.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 3 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)

					d3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					time.Sleep(time.Second)
					d3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)

					for d3.floor < requestedFloor {

						d3.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d3.floor)

					}

					d3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					d3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					d3.status = "idle"
					status()

				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d3\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d3.floor)

					d3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					time.Sleep(time.Second)
					d3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)

					for d3.floor > requestedFloor {

						d3.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d3.floor)

					}

					d3.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					d3.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d3.door)
					d3.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 4 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)

					d4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					time.Sleep(time.Second)
					d4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)

					for d4.floor < requestedFloor {

						d4.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d4.floor)

					}

					d4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					d4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					d4.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d4\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d4.floor)

					d4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					time.Sleep(time.Second)
					d4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)

					for d4.floor > requestedFloor {

						d4.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d4.floor)

					}

					d4.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					d4.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d4.door)
					d4.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		} else if elev == 5 {

			if userFloor < requestedFloor {

				if 1 <= requestedFloor && requestedFloor < 41 || requestedFloor > 60 || requestedFloor <= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)

					d5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					time.Sleep(time.Second)
					d5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)

					for d5.floor < requestedFloor {

						d5.floor++
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d5.floor)

					}

					d5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					d5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					d5.status = "idle"
					status()
				}

			} else if userFloor > requestedFloor {

				if requestedFloor < -5 || requestedFloor >= 1 || requestedFloor >= userFloor {

					time.Sleep(time.Second)
					fmt.Printf("please enter valid information\n")

				} else {

					time.Sleep(time.Second)
					fmt.Printf("elevator d5\n")
					time.Sleep(time.Second)
					fmt.Printf("elevator's floor: %d\n", d5.floor)

					d5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					time.Sleep(time.Second)
					d5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)

					for d5.floor > requestedFloor {

						d5.floor--
						time.Sleep(time.Second)
						fmt.Printf("floor display: %d\n", d5.floor)

					}

					d5.door = "opened"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					d5.door = "closed"
					time.Sleep(time.Second)
					fmt.Printf("door: %v\n", d5.door)
					d5.status = "idle"
					status()

				}

			} else {

				time.Sleep(time.Second)
				fmt.Printf("please enter valid information\n")

			}

		}
	}

	//    scenario #1
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
	fmt.Printf("-------- scenario #1 --------\n")
	requestElevB(16, "down")

	//    scenario #2
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
	fmt.Printf("-------- scenario #2 --------\n")
	requestElevC(1, "up")

	//    scenario #3
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
	fmt.Printf("-------- scenario #3 --------\n")
	requestElevD(54, "down")

	//    scenario #4
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
	fmt.Printf("-------- scenario #4 --------\n")
	requestElevA(-2, "up")

	//    test #1
	b1.floor = 7
	b2.floor = 19
	b3.floor = 12
	b4.floor = 3
	b5.floor = 9
	fmt.Printf("-------- test #1 --------\n")
	userFloor = b3.floor
	requestFloorB(3, 5)

	//    test #2
	d1.floor = 43
	d2.floor = 58
	d3.floor = 47
	d4.floor = 1
	d5.floor = 51
	fmt.Printf("-------- test #2 --------\n")
	userFloor = d4.floor
	requestFloorD(4, 41)

	//    test #3
	a1.floor = -5
	a2.floor = 0
	a3.floor = -3
	a4.floor = 1
	a5.floor = -2
	fmt.Printf("-------- test #3 --------\n")
	userFloor = a1.floor
	requestFloorA(1, -3)

	//    test #4
	c1.floor = 21
	c2.floor = 36
	c3.floor = 29
	c4.floor = 1
	c5.floor = 39
	fmt.Printf("-------- test #4 --------\n")
	userFloor = c2.floor
	requestFloorC(2, 24)

}
