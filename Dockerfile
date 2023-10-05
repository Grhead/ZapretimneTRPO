FROM golang:latest

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build main.go
EXPOSE 8085
CMD ["./main"]
