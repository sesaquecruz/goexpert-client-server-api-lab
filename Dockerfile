FROM golang:1.19-bullseye

WORKDIR /src

COPY . .

RUN apt-get update && apt-get install sqlite3

RUN sqlite3 server.db < ./sql/quotes_table.sql

RUN go mod download

RUN go build -o build/server cmd/server/main.go

RUN go build -o build/client cmd/client/main.go

EXPOSE 8080

CMD ["sleep", "infinity"]
