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


