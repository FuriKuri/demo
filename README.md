# demo

## Usage

### Start server

In docker

```
$ docker run -p 8080:8080 furikuri/demo

$ http localhost:8080/host
```

or in a kubernetes cluster

```
$ kubectl apply -f https://raw.githubusercontent.com/FuriKuri/demo/master/deploy.yaml

$ http 127.0.0.1:8001/api/v1/namespaces/default/services/http:client:80/proxy/host
```

or in docker swarm

```
$ docker stack deploy --compose-file compose.yaml stackdemo

$ http 127.0.0.1:8080/host
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

**/echo** returns the request data

```
$ http localhost:8080/echo my-header:header-value my-body=body-value
HTTP/1.1 200 OK
Content-Length: 256
Content-Type: text/plain; charset=utf-8
Date: Mon, 02 Apr 2018 12:57:57 GMT

POST /echo HTTP/1.1
Host: localhost:8080
user-agent: HTTPie/0.9.9
accept: application/json, */*
content-type: application/json
content-length: 25
accept-encoding: gzip, deflate
connection: keep-alive
my-header: header-value

Body: {"my-body": "body-value"}
```

**/raw** returns predefined content over environment variable `RAW_CONTENT`

```
$ docker run -p 8080:8080 -e RAW_CONTENT='{"key":"value"}' furikuri/demo

$ http localhost:8080/raw
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: text/plain; charset=utf-8
Date: Mon, 02 Apr 2018 13:00:11 GMT

{
    "key": "value"
}
```

**/random** returns a uuid

```
$ http localhost:8080/random
HTTP/1.1 200 OK
Content-Length: 36
Content-Type: text/plain; charset=utf-8
Date: Mon, 02 Apr 2018 13:01:04 GMT

BB113495-AAF1-3C35-F40F-C8DEF014641C
```

**/html** returns a simple html page with predefined content over `HTML_TITLE` and `HTML_BODY`

```
$ docker run -p 8080:8080 -e HTML_TITLE=Hello -e HTML_BODY=World furikuri/demo

$ http localhost:8080/html
HTTP/1.1 200 OK
Content-Length: 84
Content-Type: text/html; charset=utf-8
Date: Mon, 02 Apr 2018 13:02:19 GMT

<html>
  <body>
    <h1>Hello</h1>
    <p>
        World
    </p>
  </body>
</html>
```

**/delay/{seconds}** delay the response

```
$ time http localhost:8080/delay/4
HTTP/1.1 200 OK
Content-Length: 8
Content-Type: text/plain; charset=utf-8
Date: Mon, 02 Apr 2018 13:04:24 GMT

delay: 4

http localhost:8080/delay/4  0,36s user 0,06s system 9% cpu 4,130 total
```

**/http/{url}** do a http request to another url

```
$ http localhost:8080/http/other-server:8080/path
```