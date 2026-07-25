FROM golang:1.25.12-alpine

WORKDIR /hangman

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY data data
COPY ui ui
COPY model model

RUN go mod download
RUN go install
RUN go build

CMD  ["./hangman"]