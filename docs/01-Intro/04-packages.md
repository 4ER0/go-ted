# Go Packages

Go uses packages as a way to organize and reuse code. A package is a collection of Go source files in a single directory.

## Importing Packages

To use a package in your code, you need to import it. This is done using the `import` keyword. For example:
```golang
import "fmt"
```
You can also import multiple packages at once:
```golang
import (
    "fmt"
    "math"
)
```
You can also give a package an alias name while importing it:
```golang
import m "math"
```
## Creating a Package
To create a package, simply create a new directory and put your Go source files in it. The package name is the same as the name of the directory.

## Package Visibility
Go follows a simple rule when it comes to package visibility: if a name starts with a lowercase letter, it is private to the package and can only be accessed within the package. If a name starts with an uppercase letter, it is exported and can be accessed from other packages.

## Standard Library
Go has a large standard library that includes a wide variety of packages for everything from web development to working with data. Some of the most commonly used standard library packages include:

- `fmt` for formatted I/O
- `math` for basic mathematical operations
- `net/http` for building HTTP servers and clients
- `encoding/json` for working with JSON data
- `os` for interacting with the operating system
