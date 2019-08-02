Hello Athens
===

Try Athens project.
https://docs.gomods.io/

How to try
---

```
$ docker-compose up -d
```

Then, running Athens, MongoDB and 2 sample Go apps: `app` and `app-with-athens`.
`app` and `app-with-athens` are same application but differ in downloading sources.
`app` downloads go modules from GitHub directly. `app-with-athens` downloads them from Athens. And Athens downloads from GitHub if doesn't have cache.

```
$ docker-compose ps
             Name                           Command               State           Ports         
------------------------------------------------------------------------------------------------
hello-athens_app-with-athens_1   go run main.go                   Up      0.0.0.0:8081->8080/tcp
hello-athens_app_1               go run main.go                   Up      0.0.0.0:8080->8080/tcp
hello-athens_athens_1            /sbin/tini -- athens-proxy ...   Up      0.0.0.0:3000->3000/tcp
hello-athens_mongo_1             docker-entrypoint.sh mongod      Up      27017/tcp   
```

Sample Go app uses `github.com/julienschmidt/httprouter`.
So app downloads it before running.

You can see it as log.

```
$ docker-compose logs app
Attaching to hello-athens_app_1
app_1              | go: finding github.com/julienschmidt/httprouter v1.2.0
app_1              | go: downloading github.com/julienschmidt/httprouter v1.2.0
app_1              | go: extracting github.com/julienschmidt/httprouter v1.2.0
app_1              | 2019/08/02 23:04:29 starting HTTP server...

$ docker-compose logs app-with-athens
Attaching to hello-athens_app-with-athens_1
app-with-athens_1  | go: finding github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | go: downloading github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | go: extracting github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | 2019/08/02 23:04:29 starting HTTP server...
```

Now 2 apps are same behavior.

Next, down 2 apps.

```
$ docker-compose rm -s -f app app-with-athens
```

And disconnect internet.

Then, up 2 apps.

```
$ docker-compose up -d app app-with-athens
```

`app` fails to start because cannot download modules. 

```
$ docker-compose logs app
Attaching to hello-athens_app_1
app_1              | go: finding github.com/julienschmidt/httprouter v1.2.0
app_1              | go: github.com/julienschmidt/httprouter@v1.2.0: unknown revision v1.2.0
app_1              | go: error loading module requirements
```

On the other hand, `app-with-athens` serve successfully because Athens has the modules cache.

```
$ docker-compose logs app-with-athens
Attaching to hello-athens_app-with-athens_1
app-with-athens_1  | go: finding github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | go: downloading github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | go: extracting github.com/julienschmidt/httprouter v1.2.0
app-with-athens_1  | 2019/08/02 23:18:06 starting HTTP server...
```