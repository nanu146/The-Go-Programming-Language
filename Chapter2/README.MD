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

### Variable

##### Declaration
'var' is the keyword that creates a variable in Go programming language.

<pre>
var name type = value
</pre>

Variable declaration follow the above pattern starting with the keyword var and then the name of the variable
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

But that's not the only way of declaring a variable in Go
programming language
##### Short Variable Declaration:
We can omit using the keyword var  and data type when if we want to if we follow the below method:  
<pre>
a := 1
</pre>

Here in the type of the variable 'a' is interpreted by the compiler. But what if the value that you are trying to
hold inside the variable is different from what it may see to the compiler?
<pre>
a := 1 // here the compiler interprets a to be of type int
b float32 := 1 // here the compiler interprets
</pre>


### Scope of a Variable:
When a variable is declared inside a function its scope exists within that function.
Similarly when a function is declared in a block { } its scope exists withing that block { }.
A variable declared outside the function is accessible throughout the package and if a
package level variable is starts with a capital letter it is visible outside the package and
if the variable starts with a small letter its visibility is withing the package and cannot be exported out.
