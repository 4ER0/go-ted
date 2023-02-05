# Go-TED

This programm is used to demonstrate go capabilities to create an API.

To run the programm you need to have go installed on your machine. Then you can run the programm with the following command:

    go run ./...
or

    go run ./main/main.go

## Build the binary
You can build the binary to the dist folder with the following command:

    go build -o=./dist/estatestics ./main
An executable file named "estatestics" will be created in the dist folder.
If you want to build the binary for a different operating system you can use the following command:
    
    GOOS=darwin GOARCH=amd64 go build -o=./dist/estatestics ./main
This will create an executable file for MacOS. You can build for other operating systems as well. For more information about the GOOS and GOARCH variables see [here](https://golang.org/doc/install/source#environment).

You can install the binary with the following command:

    go build -o=$GOPATH/bin/estatestics ./main

## Test the programm

To test the programm you can run the following command:

    go test ./...

To run the tests with coverage you can run the following command:

    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

To run benchmarks you can run the following command:

    go test -bench=. -run=^# -count=5 ./...
