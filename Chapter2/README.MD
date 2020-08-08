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

### Scope of a Variable:
When a variable is declared inside a function its scope exists within that function.
Similarly when a function is declared in a block { } its scope exists withing that block { }.
A variable declared outside the function is accessible throughout the package and if a
package level variable is starts with a capital letter it is visible outside the package and
if the variable starts with a small letter its visibility is withing the package and cannot be exported out.
