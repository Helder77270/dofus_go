FROM golang:latest

COPY . .
RUN go get -u github.com/eiannone/keyboard

RUN  go run Dofus3.go