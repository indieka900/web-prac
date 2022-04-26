package main

import "fmt"

//moving the variable outside the function
var x string = "My variable"

func f() {
	fmt.Println(x)
}
func third() {
	fmt.Println("Printing the third variable", x)
}
func add(m int, y int) {
	total := 0
	total = m + y
	fmt.Println("The sum is", total)
}

//function with int as a return type
func addr(m int, y int) int {
	total := 0
	total = m + y
	return total
}
func rectangle(l int, w int) (area int) {
	var parameter int = 2 * (l + w)
	fmt.Println("Parameter: ", parameter)

	area = l * w
	return
}

//returning multiple lines
func rectangle1(l int, w int) (area1 int, parameter1 int) {
	parameter1 = 2 * (l + w)
	area1 = l * w
	return
}

//passing address to a function
func update(a *int, t *string) {
	*a = *a + 5          //defrencing pointer address
	*t = *t + " Indieka" // defrencing pointer address
	return
}

// Anonymous functions
var (
	area2 = func(l int, w int) int {
		return l * w
	}
)

func main() {
	/*
		fmt.Println("My name is Joseph")
		fmt.Println(len("My name is Joseph"))
		fmt.Println("Hello World"[1])
		fmt.Println("Hello " + "Joseph")
		fmt.Println(2*2*2*2*2*2*2*2)
		var x string = "Hello Joseph"
		fmt.Println(x)
		OR
		var xy string
		xy = x + " Hello Indieka"
		fmt.Println(xy)
		var s string = "joseph"
		var ss string = "indieka"
		var j string = "joseph"
		fmt.Println(s == ss)
		fmt.Println(s == j)
		var g string = "Hello"
		g := 5
		fmt.Println(g)
		name := "Joseph"
		fmt.Println("My name is ", name)
		calling the function in the main function
		fmt.Println(x)
		f()
		third()
		passing arguments
		add(23, 45)
		sum := addr(20, 30)
		accepting return value in variable
		fmt.Println(sum)
		fmt.Println("Area: ", rectangle(32, 21))
	*/

	//returning multiple lines
	/*
		var a, p int = rectangle1(30, 15)
		fmt.Println("Area:", a)
		fmt.Println("Perimeter:", p)
	*/

	//passing address
	/*
		 var age = 20
		var text = "Joseph"
		fmt.Println("Before:", text, age)

		update(&age, &text)
		fmt.Println("After :", text, age) */

	//anonymous function
	/*
		 fmt.Println(area2(20, 3))
		fmt.Printf(
			"100 (`f) = %.2f (`C)\n",
			func(f float64) float64{
				return (f - 32.0) * (5.0 / 9.0)
			}(100),
		)
	*/

	//closures functions
	/*
		l := 23
		w := 17
		func() {
			var area3 int = l * w
			fmt.Println(area3)
		}()
		for i := 10.0; i < 100; i += 10.0 {
			rad := func() float64 {
				return i * 39.370
			}()
			fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
		}
	*/

	//program thet accepts number from the user
	/*
		fmt.Print("Enter a number: ")
		var input float64
		fmt.Scanf("%f", &input)

		output := input * 2
		fmt.Println(output)
		y := 5
		y += 1
		fmt.Print(y)*/

	//constants
	/*const k string = "Hello World"
	k = "Any String" // will result to an error
	fmt.Println(k)
	*/

	//control structures
	//printing 1 to 10 using for loop
	/*
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
		//or
		for j := 1; j <= 10; j++ {
			fmt.Println(j)
		}
	*/

	//if statements to determine if the number is odd or even
	/*
		i := 1
		for i <= 10 {
			if i%2 == 0 {
				fmt.Println("Even ", i)
			} else {
				fmt.Println("Odd ", i)
			}
			i = i + 1
		}
	*/

	//Switch statements
	/*
		i := 1
		for i <= 10 {
			switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			case 6: fmt.Println("Six")
			default: fmt.Println("Unknown Number")
			}
			i = i + 1
		}
	*/

	//Arrays, Slices and Maps
	//Arrays
	/*
		var y [5]float64
		y[4] = 10 //the fifth element
		y[0] = 2  //first element
		y[2] = 3
		y[1] = 41
		y[3] = 16
		var total float64 = 0
		// for i := 0; i < len(y); i++ {
		// 	total += y[i]
		// }
		for _, value := range y {
			total += value
		}
		fmt.Println(total)
		fmt.Println(total / float64(len(y)))

		m := [5]float64{98, 93, 77, 82, 83, //90,
		}
		fmt.Print(m[2])
	*/

	//slices
	/*
	// we use buit-in make function
	//k := make([]float64, 5, 10)
	arr := [6]float64{1, 3, 7, 9, 8, 6}
	k := arr[0:5]
	fmt.Println(k)
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
	*/
	s := make(map[string]int)
	s["key"] = 12
	fmt.Println(s["key"])
}
