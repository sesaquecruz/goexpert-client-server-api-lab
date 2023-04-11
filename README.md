# Client-Server API-Lab - Go Expert

This project contains a simple client-server service to get the USD-BRL exchange rate. The client gets the the exchange rate from the server, then displays and saves it to a text file. The server fetches the exchange rate from an external API, saves it to a database, then returns it to the client. 

## Requirements

To use this application, you will need:

- Docker
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

3. Run the docker compose:

```
docker compose up --build
```

## Usage

### Server

1. In the project directory, enter the *lab-app* container.

```
docker compose exec lab-app bash
```

2. Inside the container, run the server:

```
./build/server
```
- The server runs on port 8080 and exposes a single endpoint at /cotacao.
- The request timeout to the server calls the external API is 200ms.
- The timeout to server saves the result in a database is 10ms.
- If a timeout is reached, an error status is returned to the client.

### Client

1. In the project directory, enter the *lab-app* container.

```
docker compose exec lab-app bash
```

2. Inside the container, run the client:

```
./build/client
```

- The client makes a request to the server at the /cotacao endpoint within a request timeout of 300ms. 
- If the request times out, the client cancels the request and displays an error.
- Upon receiving the quotation value, the client displays it on the screen, and saves it as a new line in the file "cotacao.txt" in the following format: Dólar: {value}.

## Troubleshooting

- The timeouts used can be very small depending on the internet connection. Consider changing them if necessary.

- If the client receives the status 503, consider increasing the timeout on the server to call the external API.

- If the client receives status 500 see [here](https://github.com/mattn/go-sqlite3/issues/855) for more information.

## License

This project is licensed under the MIT License. See LICENSE file for more information.
