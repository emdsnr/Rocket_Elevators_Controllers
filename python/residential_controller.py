# 1 column
# 10 floors
# 0 basements
# 2 elevators

# 18 call buttons
# 20 inside buttons


def user_floor():

    if floor1.callBtn == True:
        userFloor = 1
        return userFloor

    elif floor2.callBtn == True:
        userFloor = 2
        return userFloor

    elif floor3.callBtn == True:
        userFloor = 3
        return userFloor
        
    elif floor4.callBtn == True:
        userFloor = 4
        return userFloor
        
    elif floor5.callBtn == True:
        userFloor = 5
        return userFloor
        
    elif floor6.callBtn == True:
        userFloor = 6
        return userFloor
        
    elif floor7.callBtn == True:
        userFloor = 7
        return userFloor
        
    elif floor8.callBtn == True:
        userFloor = 8
        return userFloor
        
    elif floor9.callBtn == True:
        userFloor = 9
        return userFloor
        
    elif floor10.callBtn == True:
        userFloor = 10   
        return userFloor

userFloor = 0


class floor:

    def __init__(self, id, callBtn):
        self.id = id
        self.callBtn = callBtn


class elevator:

    def __init__(self, id, floor, status, minFloor, maxFloor, door):
        self.id = id
        self.floor = floor
        self.status = status
        self.minFloor = minFloor
        self.maxFloor = maxFloor
        self.door = door


class column:

    def __init__(self, id, status, minFloor, maxFloor, elevNum):
        self.id = id 
        self.status = status
        self.minFloor = minFloor
        self.maxFloor = maxFloor
        self.elevNum = elevNum




def requestElev(requestedFloor, direction):

    if direction == 'up':

        if userFloor == elev1.floor and elev1.status == 'goingUp':

            print('pick elevator 1')

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor < requestedFloor:
                elev1.floor += 1

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'
            

        elif userFloor == elev2.floor and elev2.status == 'goingUp':

            print('pick elevator 2')
            
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor < requestedFloor:
                elev2.floor += 1

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'            


        elif userFloor > elev1.floor and elev1.status == 'goingUp' and userFloor > elev2.floor and elev2.status == 'goingUp':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x < y:
                
                print('pick elevator 1')

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x > y:
                
                print('pick elevator 2')

                while (elev2.floor < userFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor < requestedFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor > elev1.floor and elev1.status == 'goingUp':

            print('pick elevator 1')

            while (elev1.floor < userFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor < requestedFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor > elev2.floor and elev2.status == 'goingUp':

            print('pick elevator 2')

            while (elev2.floor < userFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor < requestedFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor == elev1.floor and elev1.status == 'idle':

            print('pick elevator 1')
            elev1.status = 'goingUp'

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor < requestedFloor:
                elev1.floor += 1

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor == elev2.floor and elev2.status == 'idle':

            print('pick elevator 2')
            elev2.status = 'goingUp'

            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor < requestedFloor:
                elev2.floor += 1

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor > elev1.floor and elev1.status == 'idle' and userFloor > elev2.floor and elev2.status == 'idle':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x < y:
                
                print('pick elevator 1')
                elev1.status = 'goingUp'

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x > y:
                
                print('pick elevator 2')
                elev2.status = 'goingUp'

                while (elev2.floor < userFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor < requestedFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')
                elev1.status = 'goingUp'

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor > elev1.floor and elev1.status == 'idle':

            print('pick elevator 1')
            elev1.status = 'goingUp'

            while (elev1.floor < userFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor < requestedFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor > elev2.floor and elev2.status == 'idle':

            print('pick elevator 2')
            elev2.status = 'goingUp'

            while (elev2.floor < userFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor < requestedFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor < elev1.floor and elev1.status == 'idle' and userFloor < elev2.floor and elev2.status == 'idle':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x > y:
                
                print('pick elevator 1')
                elev1.status = 'goingUp'

                while (elev1.floor > userFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x < y:
                
                print('pick elevator 2')
                elev2.status = 'goingUp'

                while (elev2.floor > userFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor < requestedFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')
                elev1.status = 'goingUp'

                while (elev1.floor > userFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor < requestedFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor < elev1.floor and elev1.status == 'idle':
                
            print('pick elevator 1')
            elev1.status = 'goingUp'

            while (elev1.floor > userFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor < requestedFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor < elev2.floor and elev2.status == 'idle':
                
            print('pick elevator 2')
            elev2.status = 'goingUp'

            while (elev2.floor > userFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor < requestedFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor == 10 or userFloor == requestedFloor or requestedFloor <= 1 or requestedFloor > 10:

            print('please select a valid floor')


        else:

            print('all elevators are busy, try again in a few moments')



    elif direction == 'down':

        if userFloor == elev1.floor and elev1.status == 'goingDown':

            print('pick elevator 1')

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor > requestedFloor:
                elev1.floor -= 1

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'
            

        elif userFloor == elev2.floor and elev2.status == 'goingDown':

            print('pick elevator 2')

            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor > requestedFloor:
                elev2.floor -= 1

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'            


        elif userFloor < elev1.floor and elev1.status == 'goingDown' and userFloor < elev2.floor and elev2.status == 'goingDown':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x > y:
                
                print('pick elevator 1')

                while (elev1.floor > userFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x < y:
                
                print('pick elevator 2')

                while (elev2.floor > userFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor > requestedFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')

                while (elev1.floor > userFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor < elev1.floor and elev1.status == 'goingDown':

            print('pick elevator 1')

            while (elev1.floor > userFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor > requestedFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor < elev2.floor and elev2.status == 'goingDown':

            print('pick elevator 2')

            while (elev2.floor > userFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor > requestedFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor == elev1.floor and elev1.status == 'idle':

            print('pick elevator 1')
            elev1.status = 'goingDown'

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor > requestedFloor:
                elev1.floor -= 1

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor == elev2.floor and elev2.status == 'idle':

            print('pick elevator 2')
            elev2.status = 'goingDown'

            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor > requestedFloor:
                elev2.floor -= 1

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor < elev1.floor and elev1.status == 'idle' and userFloor < elev2.floor and elev2.status == 'idle':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x > y:
                
                print('pick elevator 1')
                elev1.status = 'goingDown'

                while (elev1.floor > userFloor):
                    elev1.floor += 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor += 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x < y:
                
                print('pick elevator 2')
                elev2.status = 'goingDown'

                while (elev2.floor > userFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor > requestedFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')
                elev1.status = 'goingDown'

                while (elev1.floor > userFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor < elev1.floor and elev1.status == 'idle':

            print('pick elevator 1')
            elev1.status = 'goingDown'

            while (elev1.floor > userFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor > requestedFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor < elev2.floor and elev2.status == 'idle':

            print('pick elevator 2')
            elev2.status = 'goingDown'

            while (elev2.floor > userFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor > requestedFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor > elev1.floor and elev1.status == 'idle' and userFloor > elev2.floor and elev2.status == 'idle':

            x = (userFloor - elev1.floor)
            y = (userFloor - elev2.floor)
            # print(x)
            # print(y)

            if x < y:
                
                print('pick elevator 1')
                elev1.status = 'goingDown'

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


            elif x > y:
                
                print('pick elevator 2')
                elev2.status = 'goingDown'

                while (elev2.floor < userFloor):
                    elev2.floor += 1
                

                print('move up elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
    
                while (elev2.floor > requestedFloor):
                    elev2.floor -= 1
                

                print('move down elevator 2 to floor ', elev2.floor)
                elev2.door = 'opened'
                print('door ', elev2.door)
                print('a few seconds later')
                elev2.door = 'closed'
                print('door ', elev2.door)
                elev2.status = 'idle'


            else:
                
                print('pick elevator 1')
                elev1.status = 'goingDown'

                while (elev1.floor < userFloor):
                    elev1.floor += 1
                

                print('move up elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
    
                while (elev1.floor > requestedFloor):
                    elev1.floor -= 1
                

                print('move down elevator 1 to floor ', elev1.floor)
                elev1.door = 'opened'
                print('door ', elev1.door)
                print('a few seconds later')
                elev1.door = 'closed'
                print('door ', elev1.door)
                elev1.status = 'idle'


        elif userFloor > elev1.floor and elev1.status == 'idle':
                
            print('pick elevator 1')
            elev1.status = 'goingDown'

            while (elev1.floor < userFloor):
                elev1.floor += 1
            

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while (elev1.floor > requestedFloor):
                elev1.floor -= 1
            

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor > elev2.floor and elev2.status == 'idle':
                
            print('pick elevator 2')
            elev2.status = 'goingDown'

            while (elev2.floor < userFloor):
                elev2.floor += 1
            

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while (elev2.floor > requestedFloor):
                elev2.floor -= 1
            

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor == 1 or userFloor == requestedFloor or requestedFloor < 1 or requestedFloor >= 10:

            print('please select a valid floor')


        else:
            print('all elevators are busy, try again in a few moments')



col = column('1', 'idle', 1, 10, 2)
elev1 = elevator('1', 1, 'idle', 1, 10, 'closed')
elev2 = elevator('2', 1, 'idle', 1, 10, 'closed')
floor1 = floor('1', False)
floor2 = floor('2', False)
floor3 = floor('3', False)
floor4 = floor('4', False)
floor5 = floor('5', False)
floor6 = floor('6', False)
floor7 = floor('7', False)
floor8 = floor('8', False)
floor9 = floor('9', False)
floor10 = floor('10', False)



'''
    #    scenario #1
    elev1.floor = 2
    elev2.floor = 6

    print('\\\\\\\\ scenario #1 \\\\\\\\')
    print('----- first person -----')
    floor3.callBtn = True
    userFloor = user_floor()
    floor3.callBtn = False
    requestElev(7, 'up')
'''



'''
    #    scenario #2
    elev1.floor = 10
    elev2.floor = 3


    print('\\\\\\\\ scenario #2 \\\\\\\\')
    print('----- first person -----')
    floor1.callBtn = True
    userFloor = user_floor()
    floor1.callBtn = False
    requestElev(6, 'up')

    print('----- second person -----')
    floor3.callBtn = True
    userFloor = user_floor()
    floor3.callBtn = False
    requestElev(5, 'up')

    print('----- third person -----')
    floor9.callBtn = True
    userFloor = user_floor()
    floor9.callBtn = False
    requestElev(2, 'down')
'''



'''
    #    scenario #3
    elev1.floor = 10
    elev2.floor = 3
    elev2.status = 'goingUp'


    print('\\\\\\\\ scenario #3 \\\\\\\\')
    print('----- first person -----')
    floor3.callBtn = True
    userFloor = user_floor()
    floor3.callBtn = False
    requestElev(2, 'down')

    print('----- second person -----')
    elev2.status = 'idle'
    floor10.callBtn = True
    userFloor = user_floor()
    floor10.callBtn = False
    requestElev(3, 'down')
'''



def requestFloor(elev, requestedFloor):

    if elev == '1':

        if userFloor < requestedFloor:

            print('pick elevator 1')
            elev1.status = 'goingUp'

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor < requestedFloor:
                elev1.floor += 1

            print('move up elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor > requestedFloor:

            print('elevator 1')
            elev1.status = 'goingDown'

            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)

            while elev1.floor > requestedFloor:
                elev1.floor -= 1

            print('move down elevator 1 to floor ', elev1.floor)
            elev1.door = 'opened'
            print('door ', elev1.door)
            print('a few seconds later')
            elev1.door = 'closed'
            print('door ', elev1.door)
            elev1.status = 'idle'


        elif userFloor == requestedFloor or requestedFloor < 1 or requestedFloor > 10:

            print('please select a valid floor')



    elif elev == '2':

        if userFloor < requestedFloor:

            print('elevator 2')
            elev2.status = 'goingUp'

            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor < requestedFloor:
                elev2.floor += 1

            print('move up elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor > requestedFloor:

            print('elevator 2')
            elev2.status = 'goingDown'

            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)

            while elev2.floor > requestedFloor:
                elev2.floor -= 1

            print('move down elevator 2 to floor ', elev2.floor)
            elev2.door = 'opened'
            print('door ', elev2.door)
            print('a few seconds later')
            elev2.door = 'closed'
            print('door ', elev2.door)
            elev2.status = 'idle'


        elif userFloor == requestedFloor or requestedFloor < 1 or requestedFloor > 10:

            print('please select a valid floor')



'''
    #    scenario #1
    elev1.floor = 7
    elev2.floor = 1

    print('\\\\\\\\ scenario #1 \\\\\\\\')
    print('----- first person -----')
    userFloor = elev2.floor
    requestFloor('2', 5)
'''


'''
    #    scenario #2
    elev1.floor = 5
    elev2.floor = 3

    print('\\\\\\\\ scenario #2 \\\\\\\\')
    print('----- first person -----')
    userFloor = elev1.floor
    requestFloor('1', 8)
'''


'''
    #    scenario #3
    elev1.floor = 7
    elev2.floor = 4

    print('\\\\\\\\ scenario #3 \\\\\\\\')
    print('----- first person -----')
    userFloor = elev2.floor
    requestFloor('2', -2)
'''





'''
    #    test requestElev()
    # replace x1 with the first elevator's initial floor (1 to 10 inclusive)
    # replace x2 with the second elevator's initial floor (1 to 10 inclusive)
    # replace both X3's with the user's current floor (1 to 10 inclusive)
    # replace x4 with the user's requested floor (1 to 10 inclusive)
    # replace x5 with the user's requested direction (up or down)


    elev1.floor = x1
    elev2.floor = x2

    print('\\\\\\\\ test \\\\\\\\')
    floorX3.callBtn = True
    userFloor = user_floor()
    floorX3.callBtn = False
    requestElev(x4, 'x5')
'''



'''
    #    test requestFloor()
    # replace x1 with the first elevator's initial floor (1 to 10 inclusive)
    # replace x2 with the second elevator's initial floor (1 to 10 inclusive)
    # replace both X3's with the elevator's id (1 or 2)
    # replace x4 with the user's requested floor (1 to 10 inclusive)
    elev1.floor = x1
    elev2.floor = x2

    print('\\\\\\\\ test \\\\\\\\')
    userFloor = elevX3.floor
    requestFloor('X3', x4)
'''