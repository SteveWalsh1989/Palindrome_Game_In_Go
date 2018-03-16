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

		// listOfNumbers = [6]int{2, 3, 5, 7, 11, 13}		//  array of numbers for the game

		listOfNumbers2 = [6]int{2, 3, 5, 5, 3, 2}		// TEST :   array of numbers for the game


		positionCursor = 2								// position of the cursor

		numberOfGoes   = 0								// number of valid commands the user enters

		numberDigits   = 6								// total number of digits in the list

	)

	 //-------------------------------------
	 //    Calling Game Methods
	 //-------------------------------------

	 displayState(listOfNumbers2, positionCursor,numberDigits, numberOfGoes)		// display game state

	 is_palindrome(listOfNumbers2, numberDigits)

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

	fmt.Printf("\t\t\t")                                      // separator of data

	fmt.Printf("< numberAtCursor is %d >", listOfNumbers[cursorValue])// print cursor value

	fmt.Printf("\n")                                          // new line

	fmt.Printf("   ")             							 //  Space before cursor

	//fmt.Printf("^")                                           	 // print cursor

	fmt.Printf("\n -----------------------------\n")          // line break to split content

	fmt.Printf("\n Enter your move: ")                        // ask  user to enter move

}

/**
 * moveCursorRight
 *
 * moves the cursor one value to the right
 *
 *@param  positionCursor : location of cursor
 *@param  numberDigits   : number of digits in the list

 */
 func moveCursorRight( positionCursor *int, numberDigits int, listOfNumbers [6]int) {

	if *positionCursor == numberDigits { 		// Scenario 1: position of cursor at the end of list

		*positionCursor=0						// set it to the beginning

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

*/
func moveCursorLeft(positionCursor *int, numberDigits int,listOfNumbers [6]int) {

	if *positionCursor == 0 { 						  // Scenario 1: position of cursor at the beginning of list

		*positionCursor=numberDigits				  // set it to the end

		positionCursor = &listOfNumbers[numberDigits] // sets address of position to be that of the correct element

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
func is_palindrome(listOfNumbers [6]int,numberDigits int ) {

	for i := 0 ; i < numberDigits ; i++ {                                     // loop through array size

		if  (listOfNumbers[i]) == ((listOfNumbers[numberDigits - 1]) - i) {  // copy first element of array with its last each time
			fmt.Printf("\n Text: Is a Palindrome ")                        // ask  user to enter move

		}
	}
}

