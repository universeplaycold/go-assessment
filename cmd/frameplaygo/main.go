package main

import (
	"fmt"
	"net/http"
	"api-request/pkg"
)

func main() {
	http.HandleFunc("/users", pkg.HandleRequest)
	fmt.Println("Server listening on port 8081")
	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic(err)
		}
	}()

	// Wait for user input to keep the terminal open
	// curl -X POST -d "{\"username\":\"giljohn\",\"email\":\"giljohn@gmail.com\"}" http://localhost:8081/users

	fmt.Println("Don't press until you send a curl request")
	fmt.Println("<Press Enter to exit")
	fmt.Scanln()
}
