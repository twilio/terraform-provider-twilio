FROM golang:1.18

RUN mkdir /twilio
WORKDIR /twilio

COPY client ./client
COPY core ./core
COPY twilio ./twilio
COPY main.go .

# Fetch dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download
