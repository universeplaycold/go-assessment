# ASSESSMENT - GO

This project is a Go-based web server that accepts and validates JSON input, logs the input to STDOUT, and simulates an API request.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need to have Go installed on your machine. If you do not have it, you can download it from the [official Go website](https://golang.org/dl/).

### Installing

To get the project running on your machine, follow these steps:

1. Clone this repository to your local machine using `git clone <repository_url>`.
2. Navigate to the project directory: `cd <project_directory>`.
3. Initialize the Go module: `go mod init <module_name>`.
4. Build the project: `go build`.

## Running and Testing

To run the application and perform a cURL request, follow these steps:

1. Open a terminal or command prompt.
2. Run the application by executing the following command:

$ go run ./cmd/frameplaygo/main.go

This will start the server and it will be listening on port 8081.

3. Open another terminal or command prompt.
4. Perform a cURL request to the server using the following command:

$ curl -X POST -d "{"username":"giljohn","email":"giljohn@gmail.com"}" http://localhost:8081/users

This will send a POST request with the specified JSON payload to the `/users` endpoint of the running server.

5. Observe the output in the terminal where the application is running. It will display the received request and the validated user data.

Please note that the application needs to be running in one terminal while you perform the CURL request in another terminal. This allows you to see the output and simulate the API request.

## SAMPLE OUTPUT 
C:\go\go-assessment>go run ./cmd/frameplaygo/main.go
Server listening on port 8081
Press Enter to exit
Received a request:
Username: example
Email: example@example.com
Mocking an actual API request call with user: {example example@example.com}
Successfully called API with user: {example example@example.com}
Received a request:
Username: example
Email: example@example.com
Mocking an actual API request call with user: {example example@example.com}
Successfully called API with user: {example example@example.com}

## Thought Process and Architectural Considerations

During the architecture design phase of this solution, the primary considerations were:

1. **Separation of Concerns**: Breaking down tasks into separate, manageable functions was a priority. These are setting up the HTTP server (main.go), handling HTTP requests (handler.go), and validating user data (validation.go). Each file has its own responsibilities, making the code easier to understand and maintain.

2. **Simplicity**: The codebase needed to be easy to understand, navigate, and test. Therefore, a conventional Go project structure was adopted: /cmd/frameplaygo/main.go as the application's entry point, and /pkg for the application's packages. This structure is straightforward, scales well, and is familiar to most Go developers.

## Descriptive Overview of the Solution

The solution is a Go app that:

1. Runs a web server listening on port 8081.
2. Accepts incoming HTTP POST requests at the /users endpoint.
3. Validates the user data contained in the request body, which should be a JSON object with a username and an email.
4. Prints the received request and the validated user data to the console.

The app consists of three main files:

1. `main.go`: Initializes the HTTP server and routes.
2. `handler.go`: Contains the HTTP handler function, which decodes the JSON input into a User struct and simulates the API call.
3. `validation.go`: Contains the validation function, which ensures that the username and email fields of the User struct are populated and that the email is in a valid format.

## Challenges, Considerations, and Requirements

One of the main challenges was ensuring that the app properly validates the input JSON. A custom validation function was required to check the username and email fields in the User struct, including verifying that the email is in a valid format.

Error handling was carefully considered. The app handles all possible error scenarios, including invalid JSON input and validation errors. Proper error messages are returned to the user in each case.

The requirement to log the user input to STDOUT was fulfilled by printing the user data in the HTTP handler function, once the data has been validated.

Finally, while the project description did not require the app to call an actual API, the structure of the `handler.go` file is such that a real API call could be easily added in the `HandleRequest` function.

---

## Go Playground Link

Please note that due to limitations in Go Playground, the code in the Go Playground link does not include the full functionality of making an API request. However, you can still see the structure and basic functionality of the application.

[Go Playground Link](https://go.dev/play/p/eLbSBjej2ZS)

You can copy the code from the Go Playground and run it in your local environment to test the full functionality.

---


