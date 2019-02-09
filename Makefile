
deps:
	dep ensure 
	
clean: 
	rm bin/delete
	rm bin/get
	rm bin/list-by-year
	rm bin/post 
	rm bin/put
	
build: 
	GOOS=linux GOARCH=amd64 go build -o  bin/delete ./cmd/delete/delete.go
	GOOS=linux GOARCH=amd64 go build -o  bin/get ./cmd/get/get.go 
	GOOS=linux GOARCH=amd64 go build -o  bin/list-by-year ./cmd/list/list-by-year.go
	GOOS=linux GOARCH=amd64 go build -o  bin/post ./cmd/post/post.go
	GOOS=linux GOARCH=amd64 go build -o  bin/put ./cmd/put/put.go 
	 
