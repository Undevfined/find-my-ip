# Find My IP

Find My IP is a simple HTTP server written in Go that returns the public IP address of the client making the request. It's useful for checking your public IP address programmatically.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software:

- [Go (Golang)](https://golang.org/dl/) (version 1.x or later)

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

1. **Clone the Repository**

   ```sh
   git clone https://github.com/Undevfined/find-my-ip.git
   cd find-my-ip
   ```

2. **Build the Project**

Navigate to the project directory and run the following command:

```sh
go build -o ./bin/find-my-ip
```

This command compiles the code and generates an executable in the ./bin directory.

3. **Run the Project**

To start the server, run:

```
./bin/find-my-ip
```

By default, the server will listen on port 80.

## Usage

Once the server is running, you can find your public IP by visiting `http://localhost:80` in your web browser or using a tool like `curl`:

```
curl http://localhost
```

This will return your public IP address.
