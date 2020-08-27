## Basic Data Types
Data types in go fall in one of four categories:
1. Basic Types:
    * Numbers
    * Strings
    * Booleans
2. Aggregate Types:
    * Arrays
    * Structs
    * Combinations of these two
3. Reference Types:
<br> They all refer to program state or variables indirectly so that the effect if an operation applied to one reference
 is observed in all copies.
    * Slices
    * Maps
    * Functions
    * Channels
4. Interfaces

### Numbers
There are several types of numbers in golang. integers, folating point numbers, and complex numbers (like imaginary numbers)

#### Integers:
Go provides signed and unsigned integers. They come in four sizes 8, 16, 32, 64 bits.
And referenced as int8, int16, int32, int64 for signed integers and uint8, uint16, uint32, uint64 for unsigned integers.
Signed integers can takes values on positive and negative side of the number line but, on the other hand, unsigned integers 
can only take positive integers. Also, there are synonyms for int32 and int8 which are rune and byte respectively.

In signed integers out of all bits available one bit is allocated to denote the sign of the integer hence the number of integer an
integer can take would be 2<sup>(n-1)</sup> -1 and when working with usigned integers, all the bits can be used to store numbers so
they would accommodate 2<sup>n</sup>-1 numbers. 

There are two additional types in go namely int and uint without any numbers at the end.
These act as proxy for either 32bit or 64bit integers. The underlying type cannot be assumed as it can 
change based on the compiler architecture so when converting to 32bit or 64bit integer explicit conversion is required.

When working with integers we have to be extra cautious about the operations we might perform on them in future. Otherwise funky things can happen.
Look at the example below:

<pre>
var a int8 = 127
fmt.Printf("value of a + 1 is %v and is of type %T \n", a+1, a+1) // output: value of a + 1 is -128 and is of type int8

<a href="https://play.golang.org/p/9LSTmg5VtuG">Example</a>
</pre>

Now let's look at what happens when we work with unsigned integers

<pre>
var b uint8 = 127
fmt.Printf("value of b + 1 is %v and is of type %T \n", b+1, b+1) // output: value of b + 1 is 0 and is of type uint8

<a href="https://play.golang.org/p/9LSTmg5VtuG">Example</a>
</pre>

Integers have a limited domain of numbers with which it can operate from the above example its clear when the result of an arithmetic operation
makes the value out of bound of the integer, the compiler tries to start counting from the smallest number of the domain. In case of int8 the smallest integer it can hold is -128
so adding 1 to 127 makes the value to be -128 since +128 is out of bound in the domain int8 operates. in case of uint8 the smallest number is 0 so adding 1 to 255 makes the value to be 0 and not 256 as 256 is out of bound.

#### Floating-Point Numbers:
floats come in two formats 32bit and 64bit. float32 can store numbers with a precision of 6 decimal points and float64 can 
hold 15 digits.
Its best to use float64 as float32 can quickly accumulate error as below:

<pre>
    var f float32 = 167772161
    var g float32 = 167772160
    fmt.Println(f == g)    // "true"!
</pre>


#### Complex Numbers:
Complex numbers come in two sizes complex64 and complex128 with float32 and float64 as its components.
Creating a complex number:

<pre>
var x complex128 = complex(1, 2) // 1 + 2i
var y complex128 = complex(3, 4) // 3 + 4i
var z := 1 + 3i // 1 + 3i
</pre>

There are two special functions `real` and `imag` to extract real and imaginary parts of a complex number
<pre>

var z := 1 + 3i

fmt.Printf("real: %v and imaginary: %v", real(x), imag(z))

</pre> 