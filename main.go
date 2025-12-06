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