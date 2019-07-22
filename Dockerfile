FROM shivam010/golang:latest
RUN go env

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go run main.go
