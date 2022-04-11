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

### Constants

Besides variables, Go also has another type of volatile datastore called _constants_. While the value of a variable
may change during runtime, the value of a constant is fixed at compile time and cannot be changed once assigned/declared.

Like variables, constants can be declared in almost the exact same way:

```go
const x = a
```

Here, the `const` keyword denotes that the volatile data mentioned next is assigned to a constant, similiarly to how
`var` denotes that the volatile data declared next is assigned to a variable.

### Operators

As with any programming language, Go has a set of operators that one can utilize to perform mathematical, relational
and logical operations (_as the entire subject of **Computer Science** can basically be summarised as **Applied Mathematics**_);
given below are some of the most common operators:

#### Arithmetic Operators

These are used to perform arithmetic/mathematical operations on operands in Go language.

Operator|Name|Type|Operation|Usage
---|---|---|---|---|
+|Addition Operator|Binary|Adds two values| a+b
-|Subtraction Operator|Binary|Subtracts the second value from the first value| a-b
\*|Multiplication Operator|Binary|Multiplies two values| a\*b
/|Division Operator|Binary|Divides the first value by the second value| a/b
%|Modulo Operator|Binary|Returns the remainder when the first value is divided by the second value| a%b

#### Relational Operators

Relational operators are used for comparison of two values.

Operator|Name|Type|Operation|Usage
---|---|---|---|---|
==|Equal To|Binary|Return boolean value of whether the first statement and the second statement are equivalent| a==b
!=|Not Equal To|Binary|Return boolean value of whether the first statement and the second statement are unequal| a!=b
\>|Greater Than|Binary|Return boolean value of whether the first operand is greater than the second operand.| a>b
<|Less Than|Binary|Return boolean value of whether the first operand is less than the second operand.| a\<b
\>=|Greater Than Or Equal To|Binary|Return boolean value of whether the first operand is greater than or equal to the second operand.| a>=b
<=|Less Than Or Equal To|Binary|Return boolean value of whether the first operand is lesser than or equal to the second operand.| a<=b

#### Logical Operators

They are used to combine two or more conditions/constraints or to complement the evaluation of the original condition in consideration.  

Operator|Name|Type|Operation|Usage
---|---|---|---|---|
&&|Logical And|Binary|Returns the boolean value of whether both statements are true| a&&b
\|\||Logical Or|Binary|Returns the boolean value of whether either statement is true| a||b
!|Logical Not|Unary|Returns the boolean value of whether the statement is false| !a

#### Bitwise Operators

In Go, there are six bitwise operators which work at bit level i.e, thry are used to perform bit by bit operations.

Operator|Name|Type|Operation|Usage
---|---|---|---|---|
&|Bitwise AND|Binary|Takes two numbers as operands and does AND on every bit of two numbers.| a&b
\||Bitwise OR|Binary|Takes two numbers as operands and does OR on every bit of two numbers.| a\|b
^|Bitwise XOR|Binary|Takes two numbers as operands and does XOR on every bit of two numbers.| a^b
<<|Left Shift|Binary|Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.| a\<\<b
\>\>|Right Shift|Binary|Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.| a>>b
&^|AND NOT|Binary|This is a bit clear operator.| a&^b

#### Assignment Operators

Assignment operators are used to assigning a value to a variable. The left side operand of the assignment operator is a variable and right side operand of the assignment operator is a value. The value on the right side must be of the same data-type of the variable on the left side otherwise the compiler will raise an error.

Operator|Name|Type|Operation|Usage
---|---|---|---|---|
=|Simple Assignment|Unary|This operator is used to assign the value on the right to the variable on the left.|x = 43
+=|Add assignment|Binary|This operator is combination of ‘+’ and ‘=’ operators. This operator first adds the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.| a+=b
-=|Subtract assignment|Binary|This operator is combination of ‘-‘ and ‘=’ operators. This operator first subtracts the current value of the variable on left from the value on the right and then assigns the result to the variable on the left.| a-=b
\*=|Multiply assignment|Binary|This operator is combination of ‘\*’ and ‘=’ operators. This operator first multiplies the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.| a\*=b
/=|Division assignment|Binary|This operator is combination of ‘/’ and ‘=’ operators. This operator first divides the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.| a/=b
%=|Modulo assignment|Binary|This operator is combination of ‘%’ and ‘=’ operators. This operator first modulo the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.| a%=b
&=|Bitwise AND assignment|Binary|This operator is combination of ‘&’ and ‘=’ operators. This operator first bitwise AND the current value of the variable on left with the value on the right and then assigns the result to the variable on the left.| a&=b
\|=|Bitwise OR assignment|Binary|This operator is combination of ‘\|’ and ‘=’ operators. This operator first bitwise OR the current value of the variable on left with the value on the right and then assigns the result to the variable on the left.| a\|=b
^=|Bitwise XOR assignment|Binary|This operator is combination of ‘^’ and ‘=’ operators. This operator first bitwise XOR the current value of the variable on left with the value on the right and then assigns the result to the variable on the left.| a^=b
<<=|Left Shift assignment|Binary|This operator is combination of ‘<<’ and ‘=’ operators. This operator first left shifts the current value of the variable on left with the value on the right and then assigns the result to the variable on the left.| a<<=b
\>\>=|Right Shift assignment|Binary|This operator is combination of ‘>>’ and ‘=’ operators. This operator first right shifts the current value of the variable on left with the value on the right and then assigns the result to the variable on the left.| a>>=b

### Operator Precedence in Go

Operator precedence determines the grouping of terms in an expression. This affects how an expression is evaluated. Certain operators have higher precedence than others; for example, the multiplication operator has higher precedence than the addition operator.

For example x = 7 + 3 \* 2; here, x is assigned 13, not 20 because operator \* has higher precedence than +, so it first gets multiplied with 3\*2 and then adds into 7.

Here, operators with the highest precedence appear at the top of the table, those with the lowest appear at the bottom. Within an expression, higher precedence operators will be evaluated first.

Category|Operator|Associativity
---|---|---|
Postfix|() [] -> . ++ - -|Left to right
Unary|+ - ! ~ ++ -- (type)* & sizeof|Left to right
Multiplicative|* / %|Left to right
Additive|+ -|Left to right
Shift|<< >>|Left to right
Relational|< <= > >=|Left to right
Equality|== !=|Left to right
Bitwise AND|&|Left to right
Bitwise XOR|^|Left to right
Bitwise OR|\||Left to right
Logical AND|&&|Left to right
Logical OR|\|\||Left to right
Assignment|= += -= *= /= %= >>= <<= &= ^= |=|Right to left
Comma|,|Left to right

Source: [Tutorials Point](https://www.tutorialspoint.com/go/go_operators_precedence.htm)
