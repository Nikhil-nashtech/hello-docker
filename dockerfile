FROM golang:1.20 as build
WORKDIR  /assignment
COPY . .
RUN go mod download 
RUN go build -o hello-docker 
CMD ["/assignment/main.go"]


FROM alpine:latest
COPY --from=build /assignment/hello-docker /assignment/hello-docker
EXPOSE 8081


