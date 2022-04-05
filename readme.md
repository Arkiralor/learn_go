# Documentation for Learning Go (Language by Alphabet. Inc)

## Summary

Go (_or Golang as it is colloquially called_) is a fast, functional programming language developed by Google (_now Alphabet. Inc_).
Development began in 2007 and the first stable release was on November 10th, 2009.

## Features

1. Statically Typed
2. Structurally Typed
3. Functional Language (_non-OOP_)
4. Memory Safe
5. Runtime Efficient
6. Garbage Collected

## Aim

Go was developed by google to:

_Improve programming productivity in an era of multicore, networked machines and large codebases. The designers wanted to address criticism of other languages in use at Google, but keep their useful characteristics_:

- _**Static Typing** & **Runtime Efficiancy** like C_
- _**Readability** & **Usability** like Python and/or Javascript_
- _High-Performance **Networking** & **Multiprocessing**_

## Basics

**Go** is a compiled language, as such it is translated directly to machine code in one go. This makes programs
tremendously faster than interpreted languages like Python or Javascript or than VM-based languages like Java.

Let us go over some basics of a Go program.

### Basic Structure

```go
package main

import (
    "fmt"
    )

func main() {
	fmt.Print("Hello world.\n")
}
```

Let us go over each line one-by-one:

1. `package main`: This is the package declaration. It is required for all Go programs.

2. `import ("fmt")`: This imports the required packages for the current module. 
In this case, we are importing the `fmt` package, which stands for `formatting` and 
contains basic IO functions.

3. `func main()`: This is the entry point of the Go program. 
Every Go program requires an entry point function. It basically tells the compiler where the program starts.

4. `fmt.Print()`: The `Print()` function outputs basic symbol lists to the console. Here, we are using it 
to print the string _Hello world_ to the console.

### Variables

Go has a very simple syntax for declaring variables. But the simple syntax can have any of the three forms mentioned below.
The most straight-forward of these is:
```go
var x int = a
```
Here, we are decalaring a variable named `x` of type `int` and assigning it the value `a`. 
But what if we just want to declare a variable without assigning it a value?

```go
var x int
```
In the example directly above, we have only declared a variable `x` of type `int` and have not assigned it any value. In these cases, the compiler assigns a default value to the variable. The default value depends on the type of the variable. For example, if we declare a variable of type `int`, the default value is `0`.

It should be noted that there is also a short-hand method for decalaring variables in Go, which goes something like the following:
    
```go
x := a
```

Here, the `:=` operator is used to assign a value to a variable at the time of declaration and the data-type is automatically infered from the assigned value.

#### Types of Data in Go

Keyword|Description|Value Range|Default Value
---|---|---|---|
**nil**|_Describes a NULL datatype_|_nil_|_nil_
**int8**|_8-bit signed integer_|_-128 to 127_|_0_
**int16**|_16-bit signed integer_|_-32768 to 32767_|_0_
**int32**|_32-bit signed integer_|_-2147483648 to 2147483647_|_0_
**int64**|_64-bit signed integer_|_-9223372036854775808 to 9223372036854775807_|_0_
**uint8**|_8-bit unsigned integer_|_0 to 255_|_0_
**uint16**|_16-bit unsigned integer_|_0 to 65535_|_0_
**uint32**|_32-bit unsigned integer_|_0 to 4294967295_|_0_
**uint64**|_64-bit unsigned integer_|_0 to 18446744073709551615_|_0_
**int**|_Both int and uint contain same size, either 32 or 64 bit._|_-2147483648 to 2147483647_|_0_
**uint**|_Both int and uint contain same size, either 32 or 64 bit._|_0 to 4294967295_|_0_
**bool**|_Boolean datatype_|_true or false_|_false_
**string**|_String datatype_|_any sequence of characters_|_empty string_
**rune**|_It is a synonym of int32 and also represent Unicode code points_|_0 to 0x10FFFF_|_0_
**byte**|_It is a synonym of uint8._|_0 to 255_|_0_
**uintptr**|_It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value._|_0 to 18446744073709551615_|_0_
**float32**|_32-bit IEEE 754 floating-point number_|_-3.402823e+38 to 3.402823e+38_|_0.0_
**float64**|_64-bit IEEE 754 floating-point number_|_-1.7976931348623157e+308 to 1.7976931348623157e+308_|_0.0_
**complex64**|_Complex numbers which contain float32 as a real and imaginary component._|_-3.402823e+38 to 3.402823e+38_|_0.0+0.0i_
**complex128**|_Complex numbers which contain float64 as a real and imaginary component._|_-1.7976931348623157e+308 to 1.7976931348623157e+308_|_0.0+0.0i_

#### Constants

Besides variables, Go also has another type of volatile datastore called _constants_. While the value of a variable
may change during runtime, the value of a constant is fixed at compile time and cannot be changed once assigned/declared.

Like variables, constants can be declared in almost the exact same way:

```go
const x = a
```

Here, the `const` keyword denotes that the volatile data mentioned next is assigned to a constant similiarly to how
`var` denotes that the volatile data declared next is assigned to a variable.