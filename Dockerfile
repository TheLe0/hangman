FROM golang:1.17.6-bullseye

WORKDIR /hangman

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY data data
COPY ui ui

RUN go mod download
RUN go install
RUN go build

CMD  ["./hangman"]