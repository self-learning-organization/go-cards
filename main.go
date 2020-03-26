package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Breaking down of the syntax, check point 1 below
	// card := "Ace of Spades" // This is the 100% equivalent of the code right above. Here we rely upon the compiler to infer that card is supposed to contain a string
	// card = "Five of Diamonds" // Important note on point 3 below
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // append function does not modify the existing slice. Instead it returns a new slice that we then assign back to a variable of cards

	for i, card := range cards { // Breaking down of the syntax, check point 6 below
		fmt.Println(i, card)
	}
}

func newCard() string { // Breaking down of the syntax, check point 4 below
	return "Five of Diamonds"
}

// Notes:
// 1. var card string = "Ace of Spades"
//    var tells go we're about to create a new variable. var is simply the short of variable
//    card is the name of the variable
//    string tells go that only a "string" will ever be assigned to this variable. Go is a statically typed language. So whenever we define a variable we assign it a type
//    "Ace of Spades" is basically the value assigned to this variable "card" using an equal sign
//
// 2. Basic Go Types (this is not an exhaustive list of basic types by the way)
//    Type    | Example
//    bool    | true, false
//    string  | "Hi!", "How's it going?"
//    int     | 0, -10000, 99999
//    float64 | 10.000001, 0.00009, -100.003
//
// 3. Only use := syntax when you're defining a new variable. To resign a new value to in this case, our variable "card", simply use the following:
//    card = "Five of Diamonds"
//
// 4. func newCard() string {
//	     return "Five of Diamonds"
//    }
//    newCard there helps define a function called 'newCard'
//    string there tells go that when executed, this function will return a value of type 'string'
// 5. Basic data structures in Go for handling lists of records
//    - Array: Fixed length list of things
//    - Slice: An array that can grow or shrink
//    Slices and arrays both must be defined with a data type. Every single record inside of it must be of an identical type
// 6. for i, card := range cards {
//       fmt.Println(i, card)
//    }
//    i is our index of this element in the array. i is short of index
//    card is the current card we're iterating over
//    range cards means to take the slice of 'cards' and loop over it
//    fmt.Println(i, card) within the for loop body basically means to run this line of code one time for each card in the slice
//    We use our colon equals syntax to assign the index and the value for every element inside the slice to the variables of i and card
//    Well with for loops every single time that we step through this list of cards we're really throwing away the previous index and the previous card that had been declared.
//    And so that's why we are re-declaring or re-initializing the variables in card here by using the colon equals syntax

// Quiz 2: Test Your Knowledge: Variable Assignment
// 1. Is the following a valid way of initializing and assigning a value to a variable? Yes
// var bookTitle string = "Harry Potter"
// 2. Is the following a valid way of initializing an assigning a value to a variable? Yes
// fruitCount := 5
// 3. After running the following code, Go will assume that the variable quizQuestionCount is of what type? Integer
// quizQuestionCount := 10
// 4. Will the following code compile?  Why or why not? No, because a variable can only be initialized one time. In this case, the ':=' operator is being used to intialize 'paperColor' twice
// paperColor := "Green"
// paperColor := "Blue"
// 5. Are the two lines following ways of initializing the variable 'pi' equivalent? Yes
// pi := 3.14
// var pi float = 3.14
// 6. This might require a bit of experimentation on your end :)
// Remember that you can use the Go Playground at https://play.golang.org/ to quickly run a snippet of code
// Is the following code valid? Due to a syntax error: non-declaration statement outside function body that'd occur, it's a no
/* 	package main

	import "fmt"
	deckSize := 20

	func main() {
  		fmt.Println(deckSize)
	} */
// 7. We might be able to initialize a variable and then later assign a variable to it. Is the following code valid? Yes
/*	package main

	import "fmt"

	func main() {
  		var deckSize int
  		deckSize = 52
  	fmt.Println(deckSize)
	} */
// 8. Here's another one that might need some testing on your end!
// Remember that you can use the Go Playground at https://play.golang.org/ to quickly run a snippet of code.
// Is the following code valid? Yes, we can initialize a variable outside of a function, we just can't assign a value to it
/* 	package main

	import "fmt"

	var deckSize int

	func main() {
  		deckSize = 50
  		fmt.Println(deckSize)
	} */
// 9. Is the following code valid? Why or why not? No. This is because, variables must first be initialized with the ':=' operator or the 'var variableName type' syntax
/* 	package main

	import "fmt"

	func main() {
 		deckSize = 52
  		fmt.Println(deckSize)
	} */

// Quiz 3: Test Your Knowledge: Functions
// 1. If a function returns a value, do we have to annotate, or mark, the function with the type of value that is being returned? Yes, every function that returns a value must indicate what type of value it is returning
// 2. The Go compiler is presenting an error message about the following function.  What should we do to fix the error? Add a return type of 'string' to the function, like "func getSize() string {"
/* 	func getSize() {
    	return "30 meters"
	} */
// 3. Is the following function valid? Yes
/* 	func estimatePi() float64 {
    	return 3.14
	} */
// 4. Is the following code valid? Why or why not? Yes, it is valid
/* 	package main

	import "fmt"

	func main() {
    	fmt.Println(getTitle())
	}

	func getTitle() string {
    	return "Harry Potter"
	} */
// 5. Here's a tough one!
// Imagine we have two files in the same folder with the same package name.
// This might take a little experimentation on your side.
// If you do try this out yourself, run your code with the command go run main.go state.go
// Will the following code successfully compile? Yes, because both files are part of the same package. Files in the same package can freely call functions defined in other files
/* 	In main.go:
	package main

	func main() {
    	printState()
	}

	In a separate file called state.go:
	package main

	import "fmt"

	func printState() {
    	fmt.Println("California")
	} */

// Quiz 4: Test Your Knowledge: Slices and For Loops
// 1. Which of the following represents a slice where each element in it is of type int? []int{}
// 2. Is the following code valid? No, because "strings" is not a valid type
// colors := []strings{"Red", "Yellow", "Blue"}
// 3. How do we iterate through each element in a slice and print out its value? (Will answer in next line)
// colors := []string{"Red", "Yellow", "Blue"}
// for index, color := range colors {
//    fmt.Println(index, color)
// }
// 4. Would the following code compile successfully? No, because every variable we declare must be used in our code. In this case, 'index' is not being used
/* 	for index, card := range cards {
    	fmt.Println(card)
	} */
// 5. Can a slice have both values of type 'int' and of type 'string' in it? No, because a slice can only have one type of value in it
