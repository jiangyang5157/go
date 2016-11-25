package main

import "fmt"

/*
int: either 32 or 64 bits depending on architecture
uint: unsigned (positive only), either 32 or 64 bits depending on architecture
uintptr: integer pointer


float32
float64: generally used, unless there are memory concerns
complex64: two float32
complex128: two float64


Strings
Created with double quotes allow the use of escape characters like \n and \t, while strings created with backticks allow newlines.

"Go is awesome!\n Hello, World"

`Go is awesome!
Hello, World`


Arrays
var arr [10]string
var arr [5]int{ 1, 2, 3, 4, 5 }


Slices:
var x []int
var x = make([]int, 5)

var arr = [5]int{ 1, 2, 3, 4, 5 }
x := arr[2:5] // x == [3, 4, 5] elements 2 through 4


Maps
var x map[string]float64
x := make(map[string]float64)
x["apple"] = 0.99

Variables
var x int = 7

x := 7


Constants
const message string = "Go is awesome!"
message = "Hello world!" // Throws compile error


switch
Go switch statements do not "fall through". If there are several else if statements in your if block, it may be more efficient to use the switch statement.

if x == 5 {
	fmt.Println("x is 5")
} else if x == 10 {
	fmt.Println("x is 10")
} else {
	fmt.Println("x is not 5 or 10")
}

switch x {
case 5: fmt.Println("x is 5")
case 10: fmt.Println("x is 10")
default: fmt.Println("x is not 5 or 10")
}

for
for is the only looping structure in Go.

for {
	// ever and ever
	break // just kidding
}

x := 10
for x > 0 {
	// do something
	x--
}

for x := 0; x < 10; x++ {
	// do something
}


Functions
func sum(x []int) int {...}


Returning multiple values
x, y := multVal()
func multVal() (int, int) { return 5, 9 }


Structs
Go does not have classes and does not have a class keyword, but structs are roughly analogous.
Methods are associated to structs rather than being composed within the struct.

type Person struct {
	firstName string
	lastName string
	age integer
}


var me Person
me := new(Person)
me:= Person{firstName: "Bill", lastName: "Broughton", age: 29}
me:= Person{"Bill", "Broughton", 29}


Fields
function arguments are always copied in Go, so functions will not change field values unless the instance is passed as a pointer.

func incrementAge(p Person) integer {
	p.age++
	return p.age
}

me:= Person{"Bill", "Broughton", 29}
fmt.Println(incrementAge(me)) // 30
fmt.Println(me.age) // 29


Methods
Methods are associated with structs rather than composed in the struct.

func (p *Person) incrementAge(p Person) integer {
	p.age++
	return p.age
}

me:= Person{"Bill", "Broughton", 29}
fmt.Println(incrementAge(me)) // 30
fmt.Println(me.age) // 30


Type embedding
Rather than classical inheritance, Go uses embedded types.

type Author struct {
	Person
	publishedBooks []string
}

me := Author{
	Person: Person{
		firstName true: "Bill",
		lastName: "Broughton",
		age: 29
	},
	publishedBooks: make([]string,1)
}

me.incrementAge() // Now properties and methods from Person can be used with Author
fmt.Println(me.age) // Output: 30
 */
func main() {
	var two int = 2
	var three int = 3
	fmt.Println("2 + 3 =", two + three)
	fmt.Println("3 / 2 =", three / two)

	var seven float64 = 7.0
	var twoandahalf float64 = 2.5
	fmt.Println("7.0 / 2.5 =", seven / twoandahalf)

	var strAwesome string = "Go is awesome!\nHello, World"
	fmt.Println(strAwesome)
	strAwesome = `Go is awesome!
Hello, World`
	fmt.Println(strAwesome)

	var x int = 10
	if x == 5 {
		fmt.Println("x is 5")
	} else if x == 10 {
		fmt.Println("x is 10")
	} else {
		fmt.Println("x is not 5 or 10")
	}
	x = 11
	switch x {
	case 5: fmt.Println("x is 5")
	case 10: fmt.Println("x is 10")
	default: fmt.Println("x is not 5 or 10")
	}

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
