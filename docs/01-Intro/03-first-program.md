# How to write your first Go program


## Exercise

* Create a new folder

* Run  ```go mod init "github.com/myorg/repo/first"```
* Open an editor and create a file named main.go
* Add package line with package "main"
* Write a main function
* Compile it with ```go build main.go```
* Run ```./main```

or

* Run it directly with ```go run main.go```



## Example


  ```golang
package main

import (
	"fmt"
)

func main() {
    fmt.Println("Hello, Go TED!")
}
 ```

## Run In Playground

[Run code in Go Playground](https://play.golang.org/p/2Pw61eHoF02)