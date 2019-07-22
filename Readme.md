# golang-plugin-error
This is an error-reporting repository, which aims at helping to re-create the error.

### Dockerfile

* **For Current Repository:**
	```dockerfile
	FROM shivam010/golang:latest
	RUN go env
	
	RUN mkdir /app
	COPY . /app
	WORKDIR /app
	
	RUN go run main.go
	```
* **For shivam010/golang**: [link](https://hub.docker.com/r/shivam010/golang) <br>
	```dockerfile
	FROM golang:alpine
	RUN apk add git build-base
	```

