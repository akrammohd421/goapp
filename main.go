// package main

// import "fmt"

// func main(){
// 	fmt.Printf("hello, Go");
// }

// add 
// package main 
// import "fmt"
// func main(){
// 	var x = 5

// 	x +=3
// 	fmt.Println(x);
// }

// subtract 

// package main 
// import "fmt"
// func main(){
// 	var x = 5

// 	x -=3
// 	fmt.Println(x);
// }

// multiply

// package main 
// import "fmt"
// func main(){
// 	var x = 5

// 	x *=3
// 	fmt.Println(x);
// }

// divide
// package main 
// import "fmt"
// func main(){
// 	var x = 5

// 	x /=3
// 	fmt.Println(x);
// }

// modulus
// package main 
// import "fmt"
// func main(){
// 	var x = 5

// 	x %=3
// 	fmt.Println(x);
// }

// x is a value: 10

// p is a pointer to x (it stores the memory address of x)

// & &&&&&&&&&&&&&&&

// and operator

// package main

// import "fmt"

// func main(){
// 	a:= 5
// 	b:= &a 
// 	fmt.Println("a=", a)
// 	fmt.Println("b address=", b)
// 	fmt.Println("b= actual value using pointer", *b)
// }


// output:-

// a = 5
// b = 0xc000014090   <-- memory address
// *b = 5

// x := 10
// p := &x   // p points to x

// *p = 20   // change value using pointer

// fmt.Println(x)  // 20

// -------------------------------------


// Go Comparison Operators


// Equal to=======================
// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x==y)
// }


// not equal operator	

// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x!=y)
// }

//Greater than


// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x>y)
// }

//lesser than


// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x<y)
// }


// greater than or equal to

// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x>=y)
// }

// lesser than or equal to

// package main
// import("fmt")
// func main(){

// 	var x = 5
// 	var y = 3

// 	fmt.Println(x<=y)
// }


// comparision operator end 



// logical operators
// Logical operators are used to determine the logic between variables or values:


// and operator

// package main
// import ("fmt")

// func main(){
// 	var x = 5

// 	fmt.Println(x < 5 && x < 10)
// }

//or operator


// package main
// import ("fmt")

// func main(){
// 	var x = 5

// 	fmt.Println(x < 5 || x < 10)
// }

// not operator

// package main
// import ("fmt")

// func main(){
// 	var x = 5

// 	fmt.Println(!(x < 5 && x < 10))
// }

// logical operator ennd-------------------------------


// Go Bitwise Operators

// package main
// import ("fmt")
// func main() {

//   var x = 9
//   var y = 8

//   fmt.Printf("x = %b\n",x)
//   fmt.Printf("y = %b\n",y)
    
//   fmt.Printf("x & y is %b\n",x & y)
// }




// 0r-------------------
// package main
// import ("fmt")
// func main() {

//   var x = 9
//   var y = 8

//   fmt.Printf("x = %b\n",x)
//   fmt.Printf("y = %b\n",y)
    
//   fmt.Printf("x | y is %b\n",x | y)
// }



// xor-----------------------

// package main
// import ("fmt")
// func main() {

//   var x = 9
//   var y = 8

//   fmt.Printf("x = %b\n",x)
//   fmt.Printf("y = %b\n",y)
    
//   fmt.Printf("x ^ y is %b\n",x ^ y)
// }


// Conditions 

// Use if to specify a block of code to be executed, if a specified condition is true
// Use else to specify a block of code to be executed, if the same condition is false
// Use else if to specify a new condition to test, if the first condition is false
//switch to specify many alternative blocks of code to be executed


//!if-------------------------------------
//? Use the if statement to specify a block of Go code to be executed if a condition is true.

// package main

// import ("fmt")

// func main(){
// 	if 20 > 18{
// 		fmt.Printf("hello this statement is correct")
// 	}
// }


// package main
// import ("fmt")

// func main() {
// 	x:= 20
// 	y:= 18

// 	if x > y {
//         fmt.Println("this works")
// 	}
// }


// The else Statement start--------------------
// Use the else statement to specify a block of code to be executed if the condition is false.

// Syntax
// if condition {
//   ? code to be executed if condition is true
// } else {
//   ? code to be executed if condition is false
// }


// package main
// import ("fmt")

// func main(){
// 	time := 20

// 	if time < 18 {
// 		fmt.Printf("good morning")
// 	}else{
// 		fmt.Printf("good evening")
// 	}
// }


// package main
// import ("fmt")

// func main(){
// 	temperature := 14

// 	if (temperature > 15){
// 		fmt.Printf("temperature is less and hot")
// 	}else{
// 		fmt.Printf("temperature is high and cold ")
// 	}
// }

// The else Statement end--------------------


// Go else if Statement start----------

// Use the else if statement to specify a new condition if the first condition is false.

// Syntax
// if condition1 {
//    // code to be executed if condition1 is true
// } else if condition2 {
//    // code to be executed if condition1 is false and condition2 is true
// } else {
//    // code to be executed if condition1 and condition2 are both false
// }


// package main
// import ("fmt")

// func main(){
// time := 20

// if time < 10{
// 	fmt.Println("good morning")
// }else if time < 20{
// fmt.Println("good afternon")
// }else{
// 	fmt.Println("Good evening")
// }
// }


// package main
// import ("fmt")

// func main(){
// 	a := 12
// 	b := 12

// 	if(a>b){
// 		fmt.Println("a is bigger")
// 	}else if(a<b){
// 		fmt.Println("b is bigger")
// 	}else{
// 		fmt.Println("a and b are eaual")
// 	}
// }



// Note: If condition1 and condition2 are BOTH true, only the code for condition1 are executed:

// package main

// import ("fmt")

// func main(){
// 	x:= 30

// 	if x > 20 {
// 		fmt.Println("30 is bigger than 20")
// 	}else if x>= 10{
//      fmt.Println("20 is bigger than 30")
// 	}else{
//        fmt.Println("i dont know")
// 	}
// }



// --------------------------------------

// The Nested if Statement
// You can have if statements inside if statements, this is called a nested if.

// Syntax
// if condition1 {
//    // code to be executed if condition1 is true
//   if condition2 {
//      // code to be executed if both condition1 and condition2 are true
//   }
// }




// This example shows how to use nested if statements:

// package main
// import ("fmt")

// func main() {
//   num := 20
//   if num >= 10 {
//     fmt.Println("Num is more than 10.")
//     if num > 15 {
//       fmt.Println("Num is also more than 15.")
//      }
//   } else {
//     fmt.Println("Num is less than 10.")
//   }
// }



// The switch Statement--------------------



// Single-Case switch Syntax
// Syntax
// switch expression {
// case x:
//    // code block
// case y:
//    // code block
// case z:
// ...
// default:
//    // code block
// }



// package main
// import ("fmt")

// func main() {
//   day := 4

//   switch day {
//   case 1:
//     fmt.Println("Monday")
//   case 2:
//     fmt.Println("Tuesday")
//   case 3:
//     fmt.Println("Wednesday")
//   case 4:
//     fmt.Println("Thursday")
//   case 5:
//     fmt.Println("Friday")
//   case 6:
//     fmt.Println("Saturday")
//   case 7:
//     fmt.Println("Sunday")
//   }
// }


// -------------------practice for nested if and single case switch

// package main
// import ("fmt")

// func main(){
// 	num := 20

// 	if num >= 10{
// 		fmt.Println("10 is lesser than 20")
// 		if num > 15{
// 			fmt.Println("15 is also greaater than 20")
// 		}else{
// 			fmt.Println("ITS WRONG OUTPUT")
// 		}
// 	}
// }



// package main
// import ("fmt")

// func main(){
// 	day := 4

// 	switch day{
// 	case 1:
// 		fmt.Println("sunday")
//     case 2:
// 		fmt.Println("monday")
// 	case 3:
// 		fmt.Println("tuesday")
// 	case 4:
// 		fmt.Println("wednesday")
// 	case 5:
// 		fmt.Println("thursday")
// 	case 6: 
// 	    fmt.Println("friday")
// 	case 7:
// 		fmt.Println("saturday")
// 	}


// }



// ------------------------------multi case switch 

// It is possible to have multiple values for each case in the switch statement:

// switch expression {
// case x,y:
//    // code block if expression is evaluated to x or y
// case v,w:
//    // code block if expression is evaluated to v or w
// case z:
// ...
// default:
//    // code block if expression is not found in any cases
// }	


// package main
// import ("fmt")

// func main(){

// 	day := 5

// 	switch day{
// 	case 2,4, 6:
// 		fmt.Println("this comes in even days ")
// 	case 1, 3, 5, 7:
// 	    fmt.Println("this are the odd days")
	
//     default:
// 	    fmt.Println("no days are as per u bro just enjoy u are not employed hahahah")
// 	}
// }



// go loops -------------------------------

// statement1 Initializes the loop counter value.

// statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

// statement3 Increases the loop counter value.

// package main
// import ("fmt")

// func main(){
	
// 	for i:=0; i < 5; i++{
// 		fmt.Println(i)
// 	}
// }


// explained
// i:=0; - Initialize the loop counter (i), and set the start value to 0
// i < 5; - Continue the loop as long as i is less than 5
// i++ - Increase the loop counter value by 1 for each iteration


// counts to 100 by tens

// package main
// import ("fmt")

// func main(){
// 	for i:=0; i <= 100; i+=10{
// 		fmt.Println(i)
// 	}
// }


// i:=0; - Initialize the loop counter (i), and set the start value to 0
// i <= 100; - Continue the loop as long as i is less than or equal to 100
// i+=10 - Increase the loop counter value by 10 for each iteration


// The continue Statement------------

// The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

// This example skips the value of 3:

// package main
// import ("fmt")

// func main(){
// 	for i:=0; i < 5; i++{
// 		if i ==3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }


// The break Statement

// The break statement is used to break/terminate the loop execution.

//  example breaks out of the loop when i is equal to 3:

// package main
// import("fmt")
// func main(){
// 	for i:=0; i < 5; i++{
// 		if i ==3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }


// package main
// import ("fmt")

// func main(){
// 	adj := [2]string{"apple", "banana"}
// 	fruits := [3]string{"facebook", "instagram", "twitter"}

// 	for i:= 0; i < len(adj);i++{
//     for j:=0; j < len(fruits); j++{
// 		fmt.Println(adj[i], fruits[j])
// 	}


// 	}
// }

// This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):


// package main
// import ("fmt")

// func main(){
// fruits := [3]string{"apple", "mango", "banana"}

// for idx, val := range fruits {
// fmt.Printf("%v\t%v\n", idx, val)
// }


// }


// package main
// import ("fmt")

// func main(){
// 	fruits := [3]string{"ahmednawaz", "hooriya", "akram"}

// 	for _, val := range fruits{
// 		fmt.Printf("%v\n", val)
// 	} 
      
// }



// here i will remove the values and prnt the index 

// package main
// import ("fmt")

// func main(){
// 	sports := [3]string{"calisthenics", "swimming", "gym"}

// 	for idx, _ := range sports{
// 		fmt.Printf("%v\n", idx)
// 	}
// }


// package main
// import ("fmt")

// func myMessage() {
//   fmt.Println("I just got executed!")
// }

// func main() {
//   myMessage() // call the function
// }



// package main
// import ("fmt")

// func myMessage() {
//   fmt.Println("I just got executed!")
// }

// func main() {
//   myMessage()
//   myMessage()
//   myMessage()
// }




// Function With Parameter 

// package main
// import ("fmt")

// func familyName(fname string) {
//   fmt.Println("Hello", fname, "Refsnes")
// }

// func main() {
//   familyName("Liam")
//   familyName("Jenny")
//   familyName("Anja")
// }


// package main
// import ("fmt")

// func familyName(fname string, age int) {
//   fmt.Println("Hello", age, "year old", fname, "Refsnes")
// }

// func main() {
//   familyName("Liam", 3)
//   familyName("Jenny", 14)
//   familyName("Anja", 30)
// }

// function returns

// package main
// import ("fmt")

// func myFunction(x int, y int) int {
//   return x + y
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }


// Named Return Values

// package main
// import ("fmt")

// func myFunction(x int, y int) (result int) {
//   result = x + y
//   return
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }



// package main
// import ("fmt")

// func myMessage(){
// 	fmt.Println("i just got executed")
// }

// func main() {
// 	myMessage()
// }

// function can be called multiple times

// package main
// import ("fmt")

// func myMessage(){
// 	fmt.Println("i just got executed")
// }

// func main() {
// 	myMessage()
// 	myMessage()
// 	myMessage()
// }


// Function With Parameter Example




// package main
// import ("fmt")

// func myFamily(name string){
// fmt.Println("Idrisi",  name )
// }

// func main() {
// 	myFamily("akram")
// 	myFamily("nawaz")
// 	myFamily("hooriya")
// }

// package main
// import("fmt")

// func myFamily(fname string, age int){
// fmt.Println("My name is idrisi",fname, "and my age is", age  )
// }

// func main(){
// 	myFamily("Akram", 22)
// 	myFamily("Hooriya", 18)
// 	myFamily("Ahmed Nawaz", 25)
// }

// Function Return Example

// package main
// import ("fmt")

// func myAddition(age1 int, age2 int)int{

// return age1 + age2

// }

// func main(){
// 	fmt.Println(myAddition(1, 2))
// }


// Named Return Values

// package main
// import("fmt")

// func myFunction(age1 int, age2 int)( result int ){
// result = age1 + age2
// return
// }

// func main(){
// 	fmt.Println((myFunction(3, 4)))
// }


// Store the Return Value in a Variable


// package main
// import("fmt")

// func myMessage(age1 int, age2 int)(result int){
// 	result = age1 + age2 
// return

// }

// func main(){
// total := myMessage(2, 4)
// fmt.Println(total)
// }

//multiple return values

// package main
// import ("fmt")

// func total(x int, y string)(result1 int, result2 string){
// 	result1 = x + x;
// 	result2 = y + "bhai hu"
//     return
// }

// func main(){
// 	fmt.Println(total(3, "Akram"))
// }




// package main
// import ("fmt")

// func myMsg(rollno int, name string) (result int, total string){
// 	result = rollno + rollno
// 	total = name + "world"
// return
// }

// func main(){
// 	a, b := myMsg(3, "Hello")
// 	fmt.Println(a, b)
// }

// package main
// import ("fmt")

// func final_total(roll int, name string)( fully int, half string){
// 	fully = roll + roll
// 	half = name + "World"
// 	return
// }

// func main(){
// 	_, s := final_total(6, "Hola Migo")
// 	fmt.Println(s)
// }


// package main
// import ("fmt")

// func final_total(roll int, name string)( fully int, half string){
// 	fully = roll + roll
// 	half = name + "World"
// 	return
// }

// func main(){
// 	s, _ := final_total(6, "Hola Migo")
// 	fmt.Println(s)
// }

