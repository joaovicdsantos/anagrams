# Anagrams !["golang version"](https://img.shields.io/badge/golang-v1.15.6-green)

A program that calculates anagrams made in Go using the resources of **concurrency**.

### When they act?

Goroutines are implemented when calculating anagram variations. In summary, for each letter in the word, the program will generate a process that will calculate the variants with that initial letter.

<img src="https://i.ibb.co/SrbvqXd/anagrams-func.png" alt="scheme of the program's operation" width="500"/>

## How to use

For a simple run, make sure you have "golang" installed on your machine and run the following command:
```console
budinha@budinha-pc:/anagrams/$ go run main.go golang
```
In addition, you can choose to compile the program and then run it. There is probably no performance improvement, as the `go run` command compiles to a temporary folder and executes the project. The advantage is that you will have an executable file instead of source code.
```console
budinha@budinha-pc:/anagrams/$ go build
budinha@budinha-pc:/anagrams/$ ./anagrams golang
```
As Golang is beautiful, you can still install it on your pc (if you have GOPATH configured).
```console
budinha@budinha-pc:/anagrams/$ go install
// restart the console and run anywhere
budinha@budinha-pc:/anagrams/$ anagrams golang
```

## Example
!["example"](https://media.giphy.com/media/mV8yyjfJ7viSrCPwg1/giphy.gif)
