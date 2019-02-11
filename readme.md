# How to run

1. Execute `go build`
2. Make sure `main.go` is executable by running `ls -halt`. If it is not exectuable then run `chmod +x main.go`. 
3. Execute `go run main.go`


# Assumptions
* HTTP Status code from main server is sent as it is i.e., `401` from main server would return `401` from my go rest api
* No extensive error handling is done
* I assumed that the input coming from the client is valid and hence no validation is done 