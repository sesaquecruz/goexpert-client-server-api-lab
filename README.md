# Client-Server API-Lab - Go Expert

This project contains a simple client-server service to get the USD-BRL exchange rate. The client gets the the exchange rate from the server, then displays and saves it to a text file. The server fetches the exchange rate from an external API, saves it to a database, then returns it to the client. 

## Requirements

To use this application, you will need:

- Go version 1.18 or higher installed on your computer
- SQLite3
- A stable internet connection

## Installation

To install this application: 

1. Clone this repository:

```
git clone https://github.com/sesaquecruz/goexpert-client-server-api-lab
```

2. Enter the project directory:

```
cd goexpert-client-server-api-lab
```

3. Install the required packages:

```
go mod download
```

## Usage

See Makefile.

1. Create the DB file by executing:

```
make up-db
```

2. Run the server by executing:

```
make up-server
```

- The server runs on port 8080 and exposes a single endpoint at /cotacao.
- The request timeout to the server calls the external API is 200ms.
- The timeout to server saves the result in a database is 10ms.
- If a timeout is reached, an error status is returned to the client.

3. Run the client by executing:


```
make up-client
```

- The client makes a request to the server at the /cotacao endpoint within a request timeout of 300ms. 
- If the request times out, the client cancels the request and displays an error.
- Upon receiving the quotation value, the client displays it on the screen, and saves it as a new line in the file "cotacao.txt" in the following format: Dólar: {value}.

## Troubleshooting

The timeouts used can be very small depending on the internet connection. Consider changing them if necessary.

## License

This project is licensed under the MIT License. See LICENSE file for more information.
