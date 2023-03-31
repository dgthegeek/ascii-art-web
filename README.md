# ASCII-Art Web Application

A simple web application that generates ASCII art banners based on user input.

## Usage

To run this application, execute the following command:

```go run main.go```

Then open your web browser and navigate to http://localhost:8080/.

note : make sure there no vpn running in the computer !

## Implementation Details

### Algorithm
The algorithm for generating ASCII art banners is implemented in the `runprog` package. The `RunProgram` function takes in two arguments: the input text and the banner type. It then applies the appropriate ASCII art algorithm based on the banner type and returns the generated banner.

### Package main
The `main` package is responsible for setting up the HTTP server and handling incoming requests. It defines two HTTP endpoints: `/` for the home page and `/ascii-art` for the result page. The `home` function serves the `home.html` template and the `result` function generates the ASCII art banner based on the user input and serves the `home.html` template with the generated banner.

## Authors

This application was created by :

#### Mouhamed Soumare (mnom)
#### Pape Saliou Ka (pka)
#### Dame Gaye (dgaye)

