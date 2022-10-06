Up db:
```shell

```

Up service:
```shell
go run cmd/main.go
```

Make test request:
```shell
curl localhost:8080/add -XPOST -d '{"name": "asdf", "country": "Argentina", "phone": "1231223", "website": "asdafafad"}' | jq
curl localhost:8080/b
curl localhost:8080/b -XPUT -d '{"name": "a", "country": "Argentina", "phone": "1", "website": "a"}'
curl curl localhost:8080/b -XDELETE
```