# Interfaces
In Go, interfaces define a set of methods that a struct must implement in order to be considered "implementing" the interface. Interfaces allow for polymorphism in Go, as a struct that implements an interface can be used in place of any other struct that also implements the same interface.

## Interface Syntax
```golang
type interfaceName interface {
    methodName1(parameterList1) returnType1
    methodName2(parameterList2) returnType2
    ...
}
```
Here is an example of an interface Shape with a method Area():

```golang
type Shape interface {
    Area() float64
}
```
Now we can create structs that implement this interface by defining the Area() method on them:

```golang
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```
We can now use these structs in a function that takes in a Shape:

```golang
func getArea(s Shape) float64 {
    return s.Area()
}

rect := Rectangle{width: 3, height: 4}
circ := Circle{radius: 5}

fmt.Println(getArea(rect)) // output: 12
fmt.Println(getArea(circ)) // output: 78.53981633974483
```
## Implicit Interface
In Go, if a struct implements all the methods of an interface, then it is considered to implement that interface. There is no need to explicitly declare that a struct implements an interface.

## Empty Interface
The empty interface interface{} has no methods, which means that any type implements it. This can be useful when you need a function to accept any type of argument.

Interfaces in Go provide a way to define a contract for the behavior of different types, which allows for more flexibility and modularity in your code.

## DuckTyping
Duck typing is a concept that is closely related to interfaces in Go. It refers to the ability of a struct to be considered as "implementing" an interface, not based on its type, but based on its behavior (i.e. the methods it implements). This is similar to the saying "if it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck".

In Go, since interfaces are satisfied implicitly, and there is no need to explicitly declare that a struct implements an interface, duck typing is a natural behavior of the language.

For example, the `Shape` interface from the previous example, can be implemented by any struct that has an `Area() float64` method, regardless of its type. So, any struct that has this method, it can be passed as argument to function that takes Shape as parameter, this is the main concept of duck typing.

This allows for greater flexibility and modularity in your code, as you can define interfaces with specific behavior, and then use any struct that implements that behavior, regardless of its type.

In summary, duck typing in Go refers to the ability of structs to be considered as "implementing" an interface based on the methods they implement, rather than their type. And this is a natural behavior of Go language due to the implicit interface implementation and no need of explicit declaration.