// 1 column
// 10 floors
// 0 basements
// 2 elevators
//
// 18 call buttons
// 20 inside buttons



function user_floor() {

    if (floor1.callBtn === true) {
        userFloor = 1;
        
    } else if (floor2.callBtn === true) {
        userFloor = 2;

    } else if (floor3.callBtn === true) {
        userFloor = 3;
        
    } else if (floor4.callBtn === true) {
        userFloor = 4;
        
    } else if (floor5.callBtn === true) {
        userFloor = 5;
        
    } else if (floor6.callBtn === true) {
        userFloor = 6;
        
    } else if (floor7.callBtn === true) {
        userFloor = 7;
        
    } else if (floor8.callBtn === true) {
        userFloor = 8;
        
    } else if (floor9.callBtn === true) {
        userFloor = 9;
        
    } else if (floor10.callBtn === true) {
        userFloor = 10;
        
    };

};


class floor {

    constructor(id, callBtn) {
        this.id = id;
        this.callBtn = callBtn;
    };

};


class elevator {

    constructor(id, floor, status, minFloor, maxFloor, door) {
        this.id = id;
        this.floor = floor;
        this.status = status;
        this.minFloor = minFloor;
        this.maxFloor = maxFloor;
        this.door = door;
    };

};


class column {

    constructor(id, status, minFloor, maxFloor, elevNum) {
        this.id = id;
        this.status = status;
        this.minFloor = minFloor;
        this.maxFloor = maxFloor;
        this.elevNum = elevNum;
    };

};




function requestElev(requestedFloor, direction) {

    if (direction === "up") {

        if (userFloor === elev1.floor && elev1.status === "goingUp") {

            console.log('pick elevator 1');
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor === elev2.floor && elev2.status === "goingUp") { 

            console.log('pick elevator 2');
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor > elev1.floor && elev1.status === "goingUp" && userFloor > elev2.floor && elev2.status === "goingUp") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x < y) {
                
                console.log('pick elevator 1');

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };
    
                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x > y) {
                
                console.log('pick elevator 2');

                while (elev2.floor < userFloor) {
                    elev2.floor++;
                };

                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor < requestedFloor) {
                    elev2.floor++;
                };
    
                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            } else {
                
                console.log('pick elevator 1');

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };
    
                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            };


        } else if (userFloor > elev1.floor && elev1.status === "goingUp") { 

            console.log('pick elevator 1');

            while (elev1.floor < userFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor > elev2.floor && elev2.status === "goingUp") { 

            console.log('pick elevator 2');

            while (elev2.floor < userFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor === elev1.floor && elev1.status === "idle") { 
                
            console.log('pick elevator 1');
            elev1.status = "goingUp"
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor === elev2.floor && elev2.status === "idle") { 
                
            console.log('pick elevator 2');
            elev2.status = "goingUp"
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor > elev1.floor && elev1.status === "idle" && userFloor > elev2.floor && elev2.status === "idle") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x < y) {
                
                console.log('pick elevator 1');
                elev1.status = "goingUp";

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };
    
                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x > y) {
                
                console.log('pick elevator 2');
                elev2.status = "goingUp";

                while (elev2.floor < userFloor) {
                    elev2.floor++;
                };

                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor < requestedFloor) {
                    elev2.floor++;
                };
    
                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            } else {
                
                console.log('pick elevator 1');
                elev1.status = "goingUp";

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };
    
                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            };


        } else if (userFloor > elev1.floor && elev1.status === "idle") { 

            console.log('pick elevator 1');
            elev1.status = "goingUp";

            while (elev1.floor < userFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor > elev2.floor && elev2.status === "idle") { 

            console.log('pick elevator 2');
            elev2.status = "goingUp";

            while (elev2.floor < userFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor < elev1.floor && elev1.status === "idle" && userFloor < elev2.floor && elev2.status === "idle") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x > y) {
                
                console.log('pick elevator 1');
                elev1.status = "goingUp";

                while (elev1.floor > userFloor) {
                    elev1.floor--;
                };

                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x < y) {
                
                console.log('pick elevator 2');
                elev2.status = "goingUp";

                while (elev2.floor > userFloor) {
                    elev2.floor--;
                };

                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor < requestedFloor) {
                    elev2.floor++;
                };

                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            } else {
                
                console.log('pick elevator 1');
                elev1.status = "goingUp";

                while (elev1.floor > userFloor) {
                    elev1.floor--;
                };

                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor < requestedFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            };


        } else if (userFloor < elev1.floor && elev1.status === "idle") { 

            console.log('pick elevator 1');
            elev1.status = "goingUp";

            while (elev1.floor > userFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor < elev2.floor && elev2.status === "idle") { 

            console.log('pick elevator 2');
            elev2.status = "goingUp";

            while (elev2.floor > userFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor === 10 || userFloor === requestedFloor || requestedFloor <= 1 || requestedFloor > 10) {

            console.log(`please select a valid floor`);


        } else {

            console.log(`all elevators are busy, try again in a few moments`)


        };

    } else if (direction === "down") {

        if (userFloor === elev1.floor && elev1.status === "goingDown") {

            console.log('pick elevator 1');
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor === elev2.floor && elev2.status === "goingDown") { 

            console.log('pick elevator 2');
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor < elev1.floor && elev1.status === "goingDown" && userFloor < elev2.floor && elev2.status === "goingDown") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x > y) {
                
                console.log('pick elevator 1');

                while (elev1.floor > userFloor) {
                    elev1.floor--;
                };

                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor > requestedFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x < y) {
                
                console.log('pick elevator 2');

                while (elev2.floor > userFloor) {
                    elev2.floor--;
                };

                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor > requestedFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";

            } else {
                
                console.log('pick elevator 1');

                while (elev1.floor > userFloor) {
                    elev1.floor--;
                };

                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor > requestedFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            };


        } else if (userFloor < elev1.floor && elev1.status === "goingDown") { 

            console.log('pick elevator 1');

            while (elev1.floor > userFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor < elev2.floor && elev2.status === "goingDown") { 

            console.log('pick elevator 2');

            while (elev2.floor > userFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor === elev1.floor && elev1.status === "idle") { 
                
            console.log('pick elevator 1');
            elev1.status = "goingDown";
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor === elev2.floor && elev2.status === "idle") { 
                
            console.log('pick elevator 2');
            elev2.status = "goingDown";
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor < elev1.floor && elev1.status === "idle" && userFloor < elev2.floor && elev2.status === "idle") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x > y) {
                
                console.log('pick elevator 1');
                elev1.status = "goingDown";

                while (elev1.floor > userFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor > requestedFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x < y) {
                
                console.log('pick elevator 2');
                elev2.status = "goingDown";

                while (elev2.floor > userFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor > requestedFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            } else {
                
                console.log('pick elevator 2');
                elev2.status = "goingDown";

                while (elev2.floor > userFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor > requestedFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            };


        } else if (userFloor < elev1.floor && elev1.status === "idle") { 

            console.log('pick elevator 1');
            elev1.status = "goingDown";

            while (elev1.floor > userFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor < elev2.floor && elev2.status === "idle") { 

            console.log('pick elevator 2');
            elev2.status = "goingDown";

            while (elev2.floor > userFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor > elev1.floor && elev1.status === "idle" && userFloor > elev2.floor && elev2.status === "idle") { 

            x = (userFloor - elev1.floor);
            y = (userFloor - elev2.floor);
            // console.log(x);
            // console.log(y);

            if (x < y) {
                
                console.log('pick elevator 1');
                elev1.status = "goingDown";

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor > requestedFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            } else if (x > y) {
                
                console.log('pick elevator 2');
                elev2.status = "goingDown";

                while (elev2.floor < userFloor) {
                    elev2.floor++;
                };

                console.log(`move up elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
    
                while (elev2.floor > requestedFloor) {
                    elev2.floor--;
                };
    
                console.log(`move down elevator 2 to floor ` + elev2.floor);
                elev2.door = "opened";
                console.log(`door ` + elev2.door);
                console.log(`a few seconds later`);
                elev2.door = "closed";
                console.log(`door ` + elev2.door);
                elev2.status = "idle";


            } else {
                
                console.log('pick elevator 1');
                elev1.status = "goingDown";

                while (elev1.floor < userFloor) {
                    elev1.floor++;
                };

                console.log(`move up elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
    
                while (elev1.floor > requestedFloor) {
                    elev1.floor--;
                };
    
                console.log(`move down elevator 1 to floor ` + elev1.floor);
                elev1.door = "opened";
                console.log(`door ` + elev1.door);
                console.log(`a few seconds later`);
                elev1.door = "closed";
                console.log(`door ` + elev1.door);
                elev1.status = "idle";


            };


        } else if (userFloor > elev1.floor && elev1.status === "idle") { 

            console.log('pick elevator 1');
            elev1.status = "goingDown";

            while (elev1.floor < userFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor > elev2.floor && elev2.status === "idle") { 

            console.log('pick elevator 2');
            elev2.status = "goingDown";

            while (elev2.floor < userFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor === 1 || userFloor === requestedFloor || requestedFloor < 1 || requestedFloor >= 10) {

            console.log(`please select a valid floor`);


        } else {

            console.log(`all elevators are busy, try again in a few moments`)


        };

    };

};



let col = new column(1, "idle", 1, 10, 2);
let elev1 = new elevator("1", 1, "idle", 1, 10, "closed");
let elev2 = new elevator("2", 1, "idle", 1, 10, "closed");
let floor1 = new floor('1', false);
let floor2 = new floor('2', false);
let floor3 = new floor('3', false);
let floor4 = new floor('4', false);
let floor5 = new floor('5', false);
let floor6 = new floor('6', false);
let floor7 = new floor('7', false);
let floor8 = new floor('8', false);
let floor9 = new floor('9', false);
let floor10 = new floor('10', false);



/*
    //    scenario #1
    elev1.floor = 2;
    elev2.floor = 6;

    console.log(`\\\\\\\\ scenario #1 \\\\\\\\`);
    floor3.callBtn = true;
    user_floor();
    floor3.callBtn = false;
    requestElev(7, "up");
*/



/*
    //    scenario #2
    elev1.floor = 10;
    elev2.floor = 3;

    console.log(`\\\\\\\\ scenario #2 \\\\\\\\`);
    console.log(`----- first person -----`);
    floor1.callBtn = true;
    user_floor();
    floor1.callBtn = false;
    requestElev(6, "up");

    console.log(`----- second person -----`);
    floor3.callBtn = true;
    user_floor();
    floor3.callBtn = false;
    requestElev(5, "up");

    console.log(`----- third person -----`);
    floor9.callBtn = true;
    user_floor();
    floor9.callBtn = false;
    requestElev(2, "down");
*/



/*
    //    scenario #3
    elev1.floor = 1;
    elev2.floor = 3;
    elev2.status = "goingUp";

    console.log(`\\\\\\\\ scenario #3 \\\\\\\\`);
    console.log(`----- first person -----`);
    floor3.callBtn = true;
    user_floor();
    floor3.callBtn = false;
    requestElev(2, "down");


    elev2.floor = 6;
    elev2.status = "idle";

    console.log(`----- second person -----`);
    floor10.callBtn = true;
    user_floor();
    floor10.callBtn = false;
    requestElev(3, "down");
*/



function requestFloor(elev, requestedFloor) {

    if (elev === "1") {

        if (userFloor < requestedFloor) {

            console.log('elevator 1');
            elev1.status = "goingUp";
            
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor < requestedFloor) {
                elev1.floor++;
            };

            console.log(`move up elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor > requestedFloor) {

            console.log('elevator 1');
            elev1.status = "goingDown";
            
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);

            while (elev1.floor > requestedFloor) {
                elev1.floor--;
            };

            console.log(`move down elevator 1 to floor ` + elev1.floor);
            elev1.door = "opened";
            console.log(`door ` + elev1.door);
            console.log(`a few seconds later`);
            elev1.door = "closed";
            console.log(`door ` + elev1.door);
            elev1.status = "idle";


        } else if (userFloor === requestedFloor || requestedFloor < 1 || requestedFloor > 10) {

            console.log(`please select a valid floor`);


        };

    } else if (elev === "2") {

        if (userFloor < requestedFloor) {

            console.log('elevator 2');
            elev2.status = "goingUp";
            
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor < requestedFloor) {
                elev2.floor++;
            };

            console.log(`move up elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor > requestedFloor) {

            console.log('elevator 2');
            elev2.status = "goingDown";
            
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);

            while (elev2.floor > requestedFloor) {
                elev2.floor--;
            };

            console.log(`move down elevator 2 to floor ` + elev2.floor);
            elev2.door = "opened";
            console.log(`door ` + elev2.door);
            console.log(`a few seconds later`);
            elev2.door = "closed";
            console.log(`door ` + elev2.door);
            elev2.status = "idle";


        } else if (userFloor === requestedFloor || requestedFloor < 1 || requestedFloor > 10) {

            console.log(`please select a valid floor`);


        };

    };

};



/*
    //    test #1
    elev1.floor = 7;
    elev2.floor = 1;

    console.log('\\\\\\\\ test #1 \\\\\\\\');
    userFloor = elev2.floor;
    requestFloor("2", 5);
*/


/*
    //    test #2
    elev1.floor = 5;
    elev2.floor = 3;

    console.log('\\\\\\\\ test #2 \\\\\\\\');
    userFloor = elev1.floor;
    requestFloor("1", 8);
*/


/*
    //    test #3
    elev1.floor = 7;
    elev2.floor = 4;

    console.log('\\\\\\\\ test #3 \\\\\\\\');
    userFloor = elev2.floor;
    requestFloor("2", -2);
*/





/*
    //    test requestElev();
    // replace x1 with the first elevator's initial floor (1 to 10 inclusive)
    // replace x2 with the second elevator's initial floor (1 to 10 inclusive)
    // replace both X3's with the user's current floor (1 to 10 inclusive)
    // replace x4 with the user's requested floor (1 to 10 inclusive)
    // replace x5 with the user's requested direction (up or down)


    elev1.floor = x1;
    elev2.floor = x2;

    console.log(`\\\\\\\\ test \\\\\\\\`);
    floorX3.callBtn = true;
    user_floor();
    floorX3.callBtn = false;
    requestElev(x4, "x5");
*/



/*
    //    test requestFloor();
    // replace x1 with the first elevator's initial floor (1 to 10 inclusive)
    // replace x2 with the second elevator's initial floor (1 to 10 inclusive)
    // replace both X3's with the elevator's id (1 or 2)
    // replace x4 with the user's requested floor (1 to 10 inclusive)


    elev1.floor = x1;
    elev2.floor = x2;

    console.log('\\\\\\\\ test \\\\\\\\');
    userFloor = elevX3.floor;
    requestFloor("X3", x4);
*/