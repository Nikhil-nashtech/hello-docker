FROM golang:1.20 as build
WORKDIR  /assignment
COPY . .
RUN go mod download 
RUN go build -o main.go
RUN chmod 777 main.go
CMD ["./main.go"]






