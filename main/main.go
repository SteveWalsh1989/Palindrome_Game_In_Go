package main

import(
	_ "fmt"
	  "fmt"
	  "math/rand"
	  "time"
)
/************************************************************************
 *
 *  Golang Project
 *
 *  Steve Walsh
 *
 *  17/11/2018
 *
 *  Palindrome game
 *
 *  Goal:
 *  Present list of numbers to user and goal is to make it a palindrome
 *  User can use keyboard to move between numbers using a and d keys
 *  user can increment or decrement values in list using w and x keys
 *  Program will count number of goes it takes user to complete game
 *
 *
 *
 **************************************************************************/

/**
 * main
 *
 * main program function to run
 */
func main () {

	startGame()  // start the game

}
/**
 * startGame
 *
 * starts the game
 */
 func startGame(){

	 //-------------------------------------
	 //    Declaring Variables
	 //-------------------------------------
	var (

		listOfNumbers [6]int  						    //  array of numbers for the game

		positionCursor = 0								// position of the cursor

		numberOfGoes   = 0								// number of valid commands the user enters

		numberDigits   = 6								// total number of digits in the list

		isWinner       = false                          // stores if user has won game
	)
	 //-------------------------------------
	 //    Calling Game Methods
	 //-------------------------------------

	 displayGameHeader() 																			// display header for game

	 listOfNumbers = createRandomList(listOfNumbers, numberDigits)								    // creates random list of numbers

	 randomizeCursorPosition(&positionCursor, numberDigits)											// randomize cursor starting position

	 for isWinner != true {																			// keep looping until user is winner

		displayState(listOfNumbers, positionCursor,numberDigits, numberOfGoes)						// display game state

		userCommand := getCommand()																	// store user command

		processCommand(userCommand, &listOfNumbers, &positionCursor, &numberOfGoes, numberDigits) 	// process user command

		isWinner = isPalindrome(listOfNumbers, numberDigits )										// checks if list is a palindrome
	 }
		userWon(&numberOfGoes)                                                          			// prints congratulations to user if winner
 }
/**
 * displayGameHeader
 *
 * displays header for game with : name
 *                               : goal and for controls
 *                               : valid commands
 *
 */
func displayGameHeader(){

	fmt.Printf("\n     ---------------------        ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n      THE PALINDROME GAME         ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n     ---------------------        ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n     Convert the List into        ")
	fmt.Printf("\n      a palindrome to win         ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n     ---------------------        ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n    'd' = move cursor left        ")
	fmt.Printf("\n    'a' = move cursor right       ")
	fmt.Printf("\n    'w' = increase digit value    ")
	fmt.Printf("\n    'x' = decrease digit value    ")
	fmt.Printf("\n    enter key = submit command    ")
	fmt.Printf("\n                                  ")
	fmt.Printf("\n     ---------------------        ")
}
/**
 * displayState
 *
 * shows the state of the 'game'
 *
 * show the list of numbers in the array, the value of the cursor and value of element at cursor.
 *
 *@param  listOfNumbers[]  : address of listOfNumbers
 *@param  cursorValue      : location of cursor
 *@param  numberDigits     : size of list
 *@param  numberOfGoes     : number of goes the user takes
 */
func displayState ( listOfNumbers [6]int ,  cursorValue int,numberDigits int,   numberOfGoes int) {

	fmt.Printf("\n\n\t\t Game State ")                         // print to console

	fmt.Printf("\n -----------------------------\n\n")         // line break to split content

	fmt.Printf("< ")                                           // open array for numbers

	for i := 0; i < numberDigits; i++ {
		fmt.Printf("%d ",listOfNumbers[i])      		      	  // print list
	}
	fmt.Printf(">")                                            // close array for numbers

	fmt.Printf("\t\t")                                         // separator of data

	fmt.Printf("< Num Goes: %d >", numberOfGoes )              // print number of goes

	// fmt.Printf("\t\t\t")                                       		 //
	// fmt.Printf("< Cursor at position: %d >", cursorValue+1 )   		 // TEST: cursor value
	// fmt.Printf("\t\t\t")                                      		 //
	// fmt.Printf("< Value of cursor: %d >", listOfNumbers[cursorValue]) // TEST: element in list at cursor value

	fmt.Printf("\n")                                            // new line

	fmt.Printf("  ")             							   //  Space before cursor

	for i:= 0 ; i < (cursorValue); i++  {              			  	   // loop until number of cursor

		fmt.Printf("  ")                                        // print space
	}
	fmt.Printf("^")                                             // print cursor

	fmt.Printf("\n -----------------------------\n")            // line break to split content

	fmt.Printf("\n Enter your move: ")                          // ask  user to enter move
}
/**
 * moveCursorRight
 *
 * moves the cursor one value to the right
 *
 *@param  positionCursor : location of cursor
 *@param  numberDigits   : number of digits in the list
 *@param  listOfNumbers  : list of numbers

 */
 func moveCursorRight( positionCursor *int, numberDigits int, listOfNumbers [6]int)  {

	if *positionCursor == numberDigits-1 { 		// Scenario 1: position of cursor at the end of list

		*positionCursor = 0						// set it to the beginning

		positionCursor = &listOfNumbers[0]	    // sets address of position to be that of the correct element

	} else {									// Scenario 2: position of cursor is not at the end of list

		*positionCursor++						// increase the value of position of the cursor

		var temp  = *positionCursor				// temp variable for position of cursor

		positionCursor = &listOfNumbers[temp]	// sets address of position to be that of the correct element
	}
 }
/**
 * moveCursorLeft
 *
 * moves the cursor one value to the right
 *
 *@param  positionCursor : location of cursor
 *@param  numberDigits   : number of digits in the list
 *@param  listOfNumbers  : list of numbers
 */
func moveCursorLeft(positionCursor *int, numberDigits int,listOfNumbers [6]int) {

	if *positionCursor == 0 { 						  // Scenario 1: position of cursor at the beginning of list

		*positionCursor=numberDigits-1				  // set it to the end

		positionCursor = &listOfNumbers[numberDigits-1] // sets address of position to be that of the correct element

	} else {										  // Scenario 2: position of cursor is not at the beginning of list

		*positionCursor--							  // decrease the value of position of the cursor

		var temp  = *positionCursor					  // temp variable for position of cursor

		positionCursor = &listOfNumbers[temp]   	  // sets address of position to be that of the correct element
	}
}
/**
 * incrementDigitInListAtPos
 *
 * increments value at cursor position
 *
 *@param  numberAtCursor  : value in list stored at cursor position
 */
func incrementDigitInListAtPos( numberAtCursor *int ) {

	if  *numberAtCursor == 9  {    				// Scenario 1: Value is 9

		*numberAtCursor = 0        				// resets value at position to 0

	} else {                       				// Scenario 2: value not 9

		*numberAtCursor++          				// increments value by one
	}
}
/**
 * decrementDigitInListAtPos
 *
 * decrements value at cursor position
 *
 *@param  numberAtCursor  : value in list stored at cursor position
 */
func decrementDigitInListAtPos( numberAtCursor *int ) {

	if  *numberAtCursor == 0  {    				// Scenario 1: Value is 0

		*numberAtCursor = 9        				// resets value at position to 9

	} else {                       				// Scenario 2: value not 0

		*numberAtCursor--         				// decrements value by one
	}
}
/**
 * getCommand
 *
 * Stores command entered by user
 *
 */
func getCommand( )*string {

	var inputCommand string 		//define variable to store inputCommand

	fmt.Scanln(&inputCommand) 		// stores character from user

	return &inputCommand 			// returns users character
}
/**
 * processCommand
 *
 * processes user command by calling relative method
 *
 *@param command    	: address of command the user enters
 *@param listOfNumbers  : list of numbers
 *@param positionCursor : position of the cursor
 *@param numberOfGoes   : address of number of goes value
 *@param numberDigits   : number of digits in the list
 */
func processCommand(command *string, listOfNumbers *[6]int, positionCursor *int, numberOfGoes *int, numberDigits int) {

	if *command == "w" { 											  // Scenario 1: User enters w

		incrementDigitInListAtPos(&listOfNumbers[*positionCursor])     // increases number

		*numberOfGoes += 1 											  // increase number of goes

	} else if *command == "x" { 								      // Scenario 2: User enters x

		decrementDigitInListAtPos(&listOfNumbers[*positionCursor])     // decreases number

		*numberOfGoes += 1 											  // increase number of goes

	} else if *command == "a" { 									  // Scenario 3: User enters a

		moveCursorLeft(positionCursor, numberDigits, *listOfNumbers)  // move cursor left

		*numberOfGoes += 1 										  	  // increase number of goes

	} else if *command == "d" { 									  // Scenario 4: User enters d

		 moveCursorRight(positionCursor, numberDigits, *listOfNumbers)// move cursor right

		*numberOfGoes += 1 											  // increase number of goes

	} else {														  // Scenario 5 : invalid command entered

		fmt.Printf("\n*------------------------------------------*")
		fmt.Printf("\n*      Error: Enter a valid command        *")
		fmt.Printf("\n*------------------------------------------*")



	}
}
/**
 * isPalindrome
 *
 * checks if the array is a palindrome
 *
 * Returns true if a palindrome
 * Returns false if not a palindrome
 *
 *@param pListOfNumbers : address of first element in list
 *@param size           : size of the list
 */
func isPalindrome(listOfNumbers [6]int,numberDigits int ) bool  {

	isValid := true                                                 		  // initiate return value

	for i := 0 ; i < numberDigits ; i++ {                                     // loop through array size

		if  listOfNumbers[i] != listOfNumbers[(numberDigits - 1)-i]  {  	  // compare first element of array with its mirrored index value

			isValid = false                                                	  // if not the same then it is not a palindrome
		}
	}
	return isValid															  // return if is palindrome
}
/**
 *
 * userWon
 *
 * Tells user they won and how many goes it look
 *
 * @param numberOfGoes  : address of number of goes it took user to win
 */
func userWon( numberOfGoes *int){

	fmt.Printf("\n")
	fmt.Printf("\n*--------------------------------*")
	fmt.Printf("\n*    Congratulations You Won!    *")
	fmt.Printf("\n*--------------------------------*")
	fmt.Printf("\n\n You took a total number of %d goes to create the palindrome!\n\n", *numberOfGoes)
}


/**
 * randomizeCursorPosition
 *
 * randomizes the cursor position at the start of the game
 *
 *@param pPosOfCursor   : address of first cursor
 *@param size           : size of the list
 */
func randomizeCursorPosition( positionCursor *int, numberDigits int) {

	rand.Seed(time.Now().UTC().UnixNano())     // seeds rand by time to varey each time program runs

	*positionCursor = rand.Intn(numberDigits)  // set position of the cursor to random number between 0 - 6

}

/**
 * createRandomList
 *
 * creates new list of 6 random numbers between 0 - 9
 *
 *@param  pPosOfCursor   : address of first cursor
 *@param  size           : size of the list
 *@return listOfNumbers  : returns list
 */
func createRandomList( listOfNumbers [6]int, numberDigits int) [6]int {

	rand.Seed(time.Now().UTC().UnixNano())     // seeds rand by time to varey each time program runs

	for i := 0 ; i < numberDigits ; i++ {      // loop through list

		listOfNumbers[i] = rand.Intn(9)     // set element in list to random number between  0 - 9
	}
  return listOfNumbers  					   // return list of random numbers
}
