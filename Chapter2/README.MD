# Chapter 2

### Structural Hierarchy

The basic structural hierarchy of go programming language is as follows:
Expressions -> Statements -> Functions -> Packages

### Naming Conventions
In Go it's usually a best practice to start names with a letter or an underscore followed by multiple letters or numbers.  
There are a set of 25 reserved keywords in Go that cannot be used for naming a variable:

1. `break`		*-> used to break out of a loop or case statement*
2. `case`		*-> used as part of switch statement*
3. `chan`		*-> used to create*
4. `const`		*-> used to declare a constant*
5. `continue`	*-> used to take control back to for statement when executed*
6. `defualt`		*-> default case used as part of switch statement*
7. `defer`		*-> used to defer the execution of a function call toward the end of execution context*
8. `else`	*-> used to execute statements when if is evaluated to false*
9. `fallthrough` *-> used to fallthrough the next case statement in switch*
10. `for` *-> used to loop through a condition*
11. `func` *-> a keyword to create a function*
12. `go` *-> used to launch a go routine to run code in parallel threads*
13. `goto` *-> used to goto a specific statement in execution context*
14. `if` *-> executes statements wrapped in if block if the condition evaluates to true*
15. `import` *-> used to import specific packages*
16. `interface` *->used to create interfaces*
17. `map`	*-> used to fallthrough the next case statement in switch*
18. `package` *-> used to declare the name of the package*
19. `range` *-> used to create key, value pair iterator*
20. `return` *-> used to return from function execution context*
21. `select` *-> lets a goroutine wait on multiple communication operations.*
22. `struct` *-> used to create a struct*
23. `switch` *-> is a branching technique*
24. `type` *-> a keyword used when creating complex types in golang*
25. `var` *->  a keyword used when declaring a variable*

### Declaration
Usually when we declare something in a programming language we are essentially
we are creating named entities and declare some or all of its its properties

<pre>
func add(a int, b int) {
    return a+b
}
</pre>

In the above example we are declaring a function and its behavior. Similarly we also create
a variable like shown below:

<pre>
var a int = 1
</pre>

In the example above we are creating a variable by name 'a' of type int that takes a value of 1.

### Variables

#### Declaration
'var' is the keyword that creates a variable in Go programming language.

<pre>
var name type = value
</pre>

Variable declaration follows the above pattern starting with the keyword var and then the name of the variable
and then its data type and the value it would hold. But Go is powerful enough to interpret the datatype at compilation.

<pre>
var a = 1
fmt.Printf("%v of type %T \n", a, a)    // prints ->  1 of type int

<a href="https://play.golang.org/p/tl7mNi-zTEv" target="_blank">Example Here</a>
</pre>

Here compiler is interpreting the variable data type to be int. But what if want the compiler
to interpret it as something else? Then we have to explicitly mention the data type as below:

<pre>
var a int64 = 1
fmt.Printf("%v of type %T \n", a, a)    // prints ->  1 of type int64

var b float32 = 1
fmt.Printf("%v of type %T \n", b, b)    // prints ->  1 of type float32

<a href="https://play.golang.org/p/YRt53csyhvV" target="_blank">Example Here</a>
</pre>

Or we can make use of type casting to solve this like below:

<pre>
var c = float32(1)
fmt.Printf("%v of type %T \n", c, c)    // prints ->  1 of type float32

<a href="https://play.golang.org/p/YRt53csyhvV" target="_blank">Example Here</a>
</pre>

But that's not the only way of declaring a variable in Go programming language

#### Short Variable Declaration:
We can omit using the keyword var and data type when if we want to if we follow the below method:
<pre>
a := 1
</pre>

Two things are happening here: a variable is getting declared and its value is being assigned.
To understand why its important to understand why `:=` is different from a normal `=` operator look at the below example:

<pre>
a := 1 // a is created and takes the type int and value 1
a := 2 // this results in the error: <span color="red">no new variables on left side of :=</span>
a = 2 // existing variable 'a' takes the value 2
</pre>


### Scope of a Variable:
When a variable is declared inside a function it is visible only with in that function. This visibility is also called scope.
A variable declared outside the function is accessible throughout the package. Instead, if we want to make a variable visible outside
the package then the variable name should start with a capital letter.

### Pointers
Pointers are a special kind of variable that can hold the address of a value.  Since a value can be of any type, it is essential to denote the data type of the value in a statically typed language like golang. Hence, the syntax for declaring a pointer goes something like this:

`var variableName * int = & OtherVariabelOfTypeInt`

The `*` symbol in the above expression helps the compiler understand that we are trying to create a variable of name `variableName` of the type int pointer.
Below is an example on working with pointers.

<pre>
func main() {
	var a int = 1
	var b * int = &a
	
	fmt.Println(b)
}
<a href="https://play.golang.org/p/qrNBGTlk7Nl" > Example </a>
</pre>

An important thing to remember is that pointers are variable that hold address location of a value and it gets that address location
as shared in the example above with `&` operator. Now knowing that fact look at the example below and estimate what would be the result.

<pre>
func main() {
	var a * int = & int(1)
	
	fmt.Println(a)
}
</pre>

If you have guessed that the above piece of code errors out then you are correct. When the compiler intercepts the statement `var a * int = & int(1)`, 
it tries to evaluate from right to left since it's an assignment operator. Since the left hand side `int(1)` never got assigned to a variable so it is never
stored and hence using `&` operator to get address location errors out. Look at this behavior we might be tempted to alter the definition of a pointer to:
A pointer is a special variable that takes address of a variable as value. This would be true if a pointer always takes a variable's address.
But there is `new` operator in golang which creates a pointer without creating a variable.

<pre>
func main(){

    var a * int = new(int)
    
    fmt.Println(a)
}
</pre>

The de-referencing operator `*` is used to get the value at a given address.

<pre>
    var a int = 8
    var c * int = &a
    
    fmt.Println(*c) // outputs value: 8
</pre> 

Consider the below example:

<pre>
    var a int = 8
    var b * int = &a
    var c int = *&a
    
    fmt.Println(b, c) // output: address location 8
</pre>

### Packages
Packages are what libraries are in languages like c, javascript. They provide modular code development. A package name is declared at the top 
of the go files. Usually when we write main function we declare that in the package called `main`. In a code base the main package is initialized first before other
packages. 

<pre>
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Hello World!")
}
</pre>

All the files that belong to a package will have the same package name in the declaration section of a package.
Usually the package name should take the same name as the folder in which it resides. This is a good practice to follow
as it makes it easy to know which package source files reside in which folder.

When a package is initialized its dependencies are resolved first. This holds true for variables as well. Look at the below example:

<pre>
package somepackage

var a = c + b

var c = b + 1

var b = 2

func main(){
    fmt.Println(a, b, c) // output: 5 3 2 
}
 
</pre>

The above code is completely valid and `b` is resolved first before resolving `a` and `c` as they both are dependent on b even though `a` and `c`
appear before `b`. 

There is special function `init` which initializes a package. This function cannot be called or referenced manually.
But it is automatically called whenever the package is initialized. If every file in a package has a init function the all those functions will be called 
in the ascending order of the file names in which they are written. This is because the compiler processes the files in ascending order or there name by default.
 