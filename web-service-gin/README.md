# go-web-service-gin

Simple REST api in Go.

Followed [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

Resolve and download dependencies

```sh
go get .
```

Run code in current directory

```sh
go run .
```

Get all albums

```sh
curl http://localhost:8080/albums
```

Add a new album

```sh
 curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
