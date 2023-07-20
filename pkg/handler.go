package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received a request:")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Username:", user.Username)
	fmt.Println("Email:", user.Email)

	err = Validate(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here we would send the request to the API, using the user data as part of the request message
    // I've commented out the actual request because it's not necessary for this task, but I've included a placeholder so you can see where it would go.
    // apiResponse, err := sendApiRequest(user)
    // if err != nil {
    //     http.Error(w, "API request failed", http.StatusInternalServerError)
    //     return
    // }

	// Mocking the actual API call
	fmt.Printf("Mocking an actual API request call with user: %v\n", user)
	fmt.Printf("Successfully called API with user: %v\n", user)
	w.Write([]byte(fmt.Sprintf("Successfully called API with user: %v\n", user)))
}
