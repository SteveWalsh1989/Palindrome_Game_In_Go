package main


import(
	_ "fmt"

	"fmt"
)

/******************************************************************************
 *
 * Golang Project
 *
 * Steve Walsh
 *
 * 16/11/2018
 *
 * Palindrome game
 *
 * Goal:
 * Present list of numbers to user and goal is to make it a palindrome
 * User can use keyboard to move between numbers using a and d keys
 * user can increment or decrement values in list using w and x keys
 * Program will count number of goes it takes user to complete game
 *
 *
 *
 *
 *******************************************************************************/
func main () {

	start_game()  // start the game

}


/**
 * start_game
 *
 * starts the game
 *
 * creates an integer array to use based on an integer
 * passed into method as a parameter
 * *
 */
 func start_game(){

	 //-------------------------------------
	 //    Declaring Variables
	 //-------------------------------------

	var (

		listOfNumbers = [6]int{2, 3, 5, 7, 11, 13}		//  array of numbers for the game

		positionCursor = 0								// position of the cursor

		numberOfGoes   = 0								// number of valid commands the user enters

		numberDigits   = 6

	)


	 //-------------------------------------
	 //    Calling Game Methods
	 //-------------------------------------

	 displayState(listOfNumbers, positionCursor,numberDigits, numberOfGoes)		// display game state

	 moveCursorRight(&positionCursor,numberDigits)				  				// move the cursor right

	 displayState(listOfNumbers, positionCursor,numberDigits, numberOfGoes)		// display game state

	 moveCursorLeft(&positionCursor,numberDigits)				  				// move the cursor left

	 displayState(listOfNumbers, positionCursor,numberDigits, numberOfGoes)		// display game state


 }

/**
* displayState
*
* shows the state of the 'game'
*
* show the list of numbers in the array and the value of the cursor.
*
*@param  listOfNumbers[]  :  address of listOfNumbers
*@param  positionOfCursor : location of cursor
*@param numberOfGoes      : number of goes the user takes
*/
func displayState ( listOfNumbers [6]int ,  cursorValue int,numberDigits int,   numberOfGoes int) {

//var i = 0;                                             			 // initialise counter index

	fmt.Printf("\n\n\t\t Game State ")                        // print to console

	fmt.Printf("\n -----------------------------\n\n")        // line break to split content

	fmt.Printf("< ")                                          // open array for numbers

	for i := 0; i < numberDigits; i++ {
		fmt.Printf(" %d ",listOfNumbers[i])      		     // print list
	}

	fmt.Printf(" >")                                          // close array for numbers

	fmt.Printf("\t\t")                                        // separator of data

	fmt.Printf("< Num Goes: %d >", numberOfGoes )             // print number of goes

	fmt.Printf("\t\t\t")                                      // separator of data

	fmt.Printf("< Cursor at %d >", cursorValue )              // print cursor value

	fmt.Printf("\n")                                          // new line

	fmt.Printf("   ")             							 //  Space before cursor

	fmt.Printf("^")                                           // print cursor

	fmt.Printf("\n -----------------------------\n")          // line break to split content

	fmt.Printf(  "\n Enter your move: ")                      // ask  user to enter move

}

/**
 * moveCursorRight
 *
 * moves the cursor one value to the right
 *
 *@param  positionCursor : location of cursor
 *@param  numberDigits   : number of digits in the list

 */
 func moveCursorRight(positionCursor *int, numberDigits int) {

	if *positionCursor == numberDigits { 		// Scenario 1: position of cursor at the end of list

		*positionCursor=0						// set it to the beginning

	} else {									// Scenario 2: position of cursor is not at the end of list

		*positionCursor++						// increase the value of position of the cursor
	}
 }

/**
* moveCursorLeft
*
* moves the cursor one value to the right
*
*@param  positionCursor : location of cursor
*@param  numberDigits   : number of digits in the list

*/
func moveCursorLeft(positionCursor *int, numberDigits int) {

	if *positionCursor == 0 { 					// Scenario 1: position of cursor at the beginning of list

		*positionCursor=numberDigits			// set it to the end

	} else {									// Scenario 2: position of cursor is not at the beginning of list

		*positionCursor--						// decrease the value of position of the cursor
	}
}