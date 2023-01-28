# Structs in Go

Structs in Go are a way to store data in an organized manner. They allow you to group together multiple fields with different data types.

## Creating a Struct

A struct is created by using the `struct` data type and specifying the fields within curly braces. Here is an example of a struct that represents a person:

```golang
type Person struct {
    Name string
    Age int
    Address string
}
```

In this example, we have created a struct named Person that has fields for the name, age, and address.

## Instantiating a Struct
To create an instance of a struct, you can use the syntax variable := Type{values}.

```golang
p := Person{"John Smith", 30, "123 Main St"}
```

## Accessing Fields
To access the fields of an instance of a struct, you use the dot operator (.). Here is an example:

```golang
fmt.Println(p.Name) // prints "John Smith"
```
## Structs and Methods
Structs can also have methods that can access their fields. Here is an example of a method that increases the age of a person by one:

```golang
func (p *Person) IncreaseAge() {
    p.Age++
}
```
This method can then be applied to an instance of Person:

```golang
p.IncreaseAge()
fmt.Println(p.Age) // prints 31
```
Structs are a very useful concept in Go as they allow you to store data in an organized manner and access that data in a consistent way.