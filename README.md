# demo

## Usage

### Start server

```
$ docker run -p 8080:8080 furikuri/demo
```

### Endpoints

**/host** returns the hostname

```
$ http localhost:8080/host
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain; charset=utf-8
Date: Mon, 02 Apr 2018 12:50:30 GMT

432530f24f72
```