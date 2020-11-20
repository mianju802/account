# account
### run service
```shell script
cd svr && go run main.go
```

### run cgi
```shell script
cd cgi && go run main.go
```

### shell access

```shell script
curl localhost:1234/account/register -X POST -d 'username=victor&password=victor'
```