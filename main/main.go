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

		listOfNumbers = [6]int{2, 3, 5, 7, 11, 13}

		positionCursor = 0

		numberOfGoes   = 0

		numberDigits   = 6

	)


	 displayState(listOfNumbers, positionCursor, numberDigits, numberOfGoes)		// display game state


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
*@param  max              : number of values
*@param numberOfGoes      : number of goes the user takes
*/
func displayState ( plistOfNumbers [6]int ,  cursorValue int,  max int,  numberOfGoes int) {

//var i = 0;                                             			  // initialise counter index

	fmt.Printf("\n\n\t\t Game State ")                        // print to console

	fmt.Printf("\n -----------------------------\n\n")        // line break to split content

	fmt.Printf("< ")                                          // open array for numbers

	fmt.Printf(" %d ",plistOfNumbers)      				      // print list

	fmt.Printf(" >")                                          // close array for numbers

	fmt.Printf("\t\t")                                        // separator of data

	fmt.Printf("< Num Goes: %d >", numberOfGoes )             // print number of goes

	fmt.Printf("\t\t\t")                                        // separator of data

	fmt.Printf("< Cursor at %d >", cursorValue )              // print cursor value

	fmt.Printf("\n")                                          // new line

	fmt.Printf("\t")             						//  Space before cursor


	fmt.Printf("^")                                           // print cursor

	fmt.Printf("\n -----------------------------\n")          // line break to split content

	fmt.Printf(  "\n Enter your move: ")                      // ask  user to enter move

}