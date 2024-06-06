package main

import "fmt"

type studentFormatString struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	data := studentFormatString{
		name:        "wick",
		height:      175.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"reading", "gaming", "coding"},
	}

	// layout format %b -> to show biner
	fmt.Printf("%b\n", data.age) // 11010

	// layout format %c -> to show unicode
	fmt.Printf("%c\n", 1400) //n
	fmt.Printf("%c\n", 1235) //ӓ

	// layout format %d -> to show numerik data
	fmt.Printf("%d\n", data.age)

	// layout format %e -> decimal to standart numerik
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)

	// layout format %f -> to show decimal with specific comma lenght
	fmt.Printf("%f\n", data.height)   //175.500000
	fmt.Printf("%.9f\n", data.height) //175.500000000
	fmt.Printf("%.2f\n", data.height) //175.50
	fmt.Printf("%.f\n", data.height)  //176

	// layout format %g -> to show decimal with same lenght, cannot be custom like %f
	fmt.Printf("%g\n", 0.12)   //0.12
	fmt.Printf("%.5g\n", 0.12) //0.12

	// layout format %o -> convert numerik to string (oktal basis)
	fmt.Printf("%o\n", data.age)

	// layout format %p -> to show reference from data pointer
	fmt.Printf("%p\n", &data.name)

	// layout format %q -> to escape string using \
	fmt.Printf("%q\n", `" name \ height "`)

	// layout format %t -> to show boolean data
	fmt.Printf("%t\n", data.isGraduated) //false

	// layout format %T -> to get variable type
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)

	// layout format %+v to format struct and return struct value
	fmt.Printf("%+v\n", data)

	// layout format %#v to format struct and returning the name and value of each property according to the structure of the struct and also how the object is declared.
	fmt.Printf("%#v\n", data)

	var dataAnonymouseStruct = struct {
		name   string
		height float64
	}{
		name:   "wick",
		height: 184.5,
	}
	fmt.Printf("%#v\n", dataAnonymouseStruct)

	// layout format %x or %X to format numerik data to string hexadecimal
	fmt.Printf("%x\n", data.age)

	d := data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])

	fmt.Printf("%x\n", d)

	// layout format %% to type % on string format
	fmt.Printf("%%\n")
}
