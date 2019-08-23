Hello Athens
===

Try Athens project.
https://docs.gomods.io/

How to try
---

```
$ docker-compose up -d
```

Then, running Athens and a sample a Go app.
A Go app downloads them from Athens. And Athens downloads from GitHub if doesn't have cache.

```
$ docker-compose ps
        Name                       Command               State           Ports         
---------------------------------------------------------------------------------------
hello-athens_app_1      go run main.go                   Up      0.0.0.0:8081->8080/tcp
hello-athens_athens_1   /sbin/tini -- athens-proxy ...   Up      0.0.0.0:3000->3000/tcp
```

Sample Go app uses [Gin](https://github.com/gin-gonic/gin).
So app downloads it before running.

You can see it as log.

```
$ docker-compose logs app
Attaching to hello-athens_app_1
app_1     | go: finding github.com/ugorji/go v1.1.7
app_1     | go: finding github.com/json-iterator/go v1.1.7
app_1     | go: finding github.com/golang/protobuf v1.3.2
app_1     | go: finding github.com/stretchr/testify v1.4.0
...
app_1     | [GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
app_1     | [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
app_1     | [GIN-debug] Listening and serving HTTP on :8080
```

Next, down the app.

```
$ docker-compose rm -s -f app
```

And disconnect internet.

Then, up the app.

```
$ docker-compose up -d app
```

`app` download modules and serve successfully because Athens has the modules cache.

```
$ docker-compose logs app
Attaching to hello-athens_app_1
app_1     | go: finding github.com/kr/pretty v0.1.0
app_1     | go: finding github.com/ugorji/go v1.1.7
app_1     | go: finding github.com/gin-gonic/gin v1.4.0
app_1     | go: finding github.com/gin-contrib/sse v0.1.0
app_1     | go: finding github.com/mattn/go-isatty v0.0.9
...
app_1     | [GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
app_1     | [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
app_1     | [GIN-debug] Listening and serving HTTP on :8080
```