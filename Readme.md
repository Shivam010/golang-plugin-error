# golang-plugin-error
This is an error-reporting repository, which aims at helping to re-create the error.

### Error: changing the same plugin file and building it, does not change .so file
On changing the plugin code and re-compiling it, the new changes do not add up on opening the newly generated `.so file`.

Follow the code in [main.go](./main.go) or run the docker [image](./Dockerfile)

### Output of the code
```
Iteration: 0
Iteration: 0
Iteration: 0
Iteration: 0
Iteration: 0
```

### Expected Output
```
Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
```

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

