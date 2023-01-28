# Errors

Go does not have exceptions. In Go, you deal with errors that must be handled or actively ignored.

Later, you will see panic and recover. This an equivalent to exceptions, but very rarely used in go.

Go has a concept of a stack trace, it is hidden inside the standard error, and you normally don't need it.

## The definition of error

See bultin.go 

An error in Go is nothing, but an interface with the function Error() which returns a string.

```golang
type error interface {
  Error() string
}
```

## Creating errors

Creating errors is easy and can be done with the errors.New(string) function. 

```golang
func DoBusinessLogic() error {
	....
	return errors.New("Things happen")	
	
}
```


## Creating own errors

Since error is an interface, you can define your own type that implements (duck type) the interface.

```golang

type myError struct {}

func (m myError) Error() string {
  return "Bad Day"
}

func DoBusinessLogic() error {
  ...
  return myError{}	
	
}
```


## Handling errors

An Error needs to be handled

```golang

func Test() {
  v, err := DoBusinessLogic()
  if err != nil {
	 // do error stuff	
  }
}

func DoBusinessLogic() *Value, error {
  ...
}
```


Or an Error can be ignored! 
You should know what you are doing here!

```golang

func Test() {
  v,_ := DoBusinessLogic()
}

func DoBusinessLogic() *Value, error {
  ...
}
```


## Pattern for aggregating errors
```golang
func HigherLogic() error {

  v, err := DoBusinessLogic()
  if err != nil {
    return err
  }

  x, err := DoMoreBusinessLogic(v)
  if err != nil {
    return err
  }
  
  err := DoEvenMoreBusinessLogic(x)
  if err != nil {
    return err
  }
  return nil
}
```

# Panic

panic is an exception like element and stops the executing of the actual function and keeps
returning until a recover is handling the panic.

Normally the main function, started by the go runtime will handle the panic and stops your programm.

So, conclusion is, panic without recover is only for non-server programs acceptable.

```golang

func main() {
	
	...
	DoBusinessLogic()
	...
}

func DoBusinessLogic() {
	
	...
	panic("Not Expecting this")
}
```

# Recover

A panic can be caught within a deferred function and with the function
recover(), the panic can be resolved or re-panic'd.


```golang
func main() {
	Recovering()
}


func Recovering() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	Panicking()
}

func Panicking() {
	panic(errors.New("One of these days..."))
}

```