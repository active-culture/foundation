FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/gorilla/mux

EXPOSE 8081

CMD ["go", "run", "main.go"]
