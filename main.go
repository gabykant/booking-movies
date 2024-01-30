package main

import (
	"fmt"
	"math/rand"
)

var arrayOfChoices [] int

func main() {

	var day uint  

	welcomeMessage()

	day = getDayOfWeek()  

	getPrograms(day)
}

func welcomeMessage() {
	fmt.Println("Hi! Welcome to the movie show")
	fmt.Println(" ")
	
}

func getDayOfWeek () uint {

	var day uint

	fmt.Println("Pick a day in the week and the entire program will be displayed")
 
	fmt.Println("1. Monday")
	
	fmt.Println("2. Tuesday")
	
	fmt.Println("3. Wednesday")
	
	fmt.Println("4. Thursday")
	
	fmt.Println("5. Friday")
	
	fmt.Println("6. Saturday")
	
	fmt.Println("7. Sunday")

	fmt.Scan(&day)
	
	return day
}

func getPrograms(choiceDay uint) {

	var c uint 
	mondayMovies := rand.New(rand.NewSource(99))
	arrayOfChoices = nil
	arrayOfChoices = append(arrayOfChoices, 0) 
	switch choiceDay {
	case 1:
		fmt.Println("Movies of Monday") 
		fmt.Println("	1 One Flew Over the Cuckoo's Nest") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r) 
		fmt.Printf("	This movie has %v tickets avalaible\n", r) 
		fmt.Println("	2 The Lord of the Rings: The Return of the King") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)  
		fmt.Println("Select a movie") 
	case 2:
		fmt.Println("Movies of Tuesday")  
		fmt.Println("	1 The Blues Brothers") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r )
		fmt.Printf("	This movie has %v tickets avalaible\n", r ) 
		fmt.Println("	2 Titanic") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r) 
		fmt.Println("Select a movie") 
	case 3:
		fmt.Println("Movies of Wednesday") 
		fmt.Println("	1 Fight Club") 
		r2 := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r2)
		fmt.Printf("	This movie has %v tickets avalaible\n", r2)
		fmt.Println("Select a movie") 
	case 4:
		fmt.Println("Movies of Thursday") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Println("	1 Once Upon a Time in the West") 
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("Select a movie") 
	case 5:
		fmt.Println("Movies of Friday") 
		fmt.Println("	1 Star Wars: Episode IV - A New Hope") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("	2 The Godfather") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("	3 Apocalypse Now") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("Select a movie") 
	case 6:
		fmt.Println("Movies of Saturday") 
		fmt.Println("	1 The Matrix") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("	2 American History X") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("Select a movie") 
	case 7:
		fmt.Println("Movies of Sunday") 
		fmt.Println("	1 The Deer Hunter") 
		r := mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("	2 Forrest Gump") 
		r = mondayMovies.Intn(100)
		arrayOfChoices = append(arrayOfChoices, r)
		fmt.Printf("	This movie has %v tickets avalaible\n", r)
		fmt.Println("Select a movie") 
	default:
		choiceDay = 0
		fmt.Println("Incorrect choice. Type 0 to return") 
		setChoice(choiceDay, 0)
		break;
	} 
	
	fmt.Scan(&c)
	setChoice(choiceDay, c)
}

func setChoice(myDay uint, myChoice uint) {
	
	if myChoice == 0 {
		getDayOfWeek()
	}

	numberTickets := 0 

	mySelection := arrayOfChoices[myChoice]

	fmt.Println("How many tickets do you want to purchase ?")
	fmt.Scan(&numberTickets)
	if numberTickets > mySelection {
		fmt.Printf("Sorry. There is only %v tickets remaining\n", mySelection)
		getPrograms(myDay)
	} else {
		var fullname string
		fmt.Printf("Please enter your full name:\n")
		// fmt.Scan(&numberTickets)
		fmt.Scan(&fullname)

		arrayOfChoices[myChoice] = mySelection - numberTickets

		fmt.Printf("Congratulation %v. You successfuly purchase %v tickets for the movie \n", fullname, (numberTickets))
	}
}