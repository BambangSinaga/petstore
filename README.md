# petstore
Build REST pet service using golang

### Execute Migration
```bash
# install goose
go get bitbucket.org/liamstask/goose/cmd/goose

# up table pet
goose -path="$GOPATH/src/github.com/BambangSinaga/petstore/db" up
```

### How To Run This Project


```bash
# GET WITH GO GET
go get github.com/BambangSinaga/petstore

# Go to directory
cd $GOPATH/src/github.com/BambangSinaga/petstore

# Install Dependencies
glide install -v

# Run Project
go run main.go
```
