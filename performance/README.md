Performance Comparing
===

Compare downloading modules time with downloading modules from VCS directly, using https://proxy.golang.org as `GOPROXY` and using Athens local container as `GOPROXY`.

How To Compare
---

At first, build base image.

```
$ docker-compose build
```

Download modules from VCS directly
---

```
$ GOPROXY=direct docker-compose run --rm go-get
...
143.61user 111.61system 3:06.13elapsed 137%CPU (0avgtext+0avgdata 644164maxresident)k
27096inputs+2432880outputs (179major+4427752minor)pagefaults 0swaps
```

Use https://proxy.golang.org as `GOPROXY`
---

```
$ GOPROXY=https://proxy.golang.org docker-compose run --rm go-get
...
67.06user 68.40system 0:47.32elapsed 286%CPU (0avgtext+0avgdata 644576maxresident)k
0inputs+1546448outputs (2major+2200778minor)pagefaults 0swaps
```


Use Athens local container as `GOPROXY`
---

```
$ GOPROXY=http://athens:3000 docker-compose run --rm go-get
...
64.69user 64.91system 8:16.90elapsed 26%CPU (0avgtext+0avgdata 664428maxresident)k
8inputs+1546456outputs (2major+2208477minor)pagefaults 0swaps
```

It took much time because Athens didn't have any Modules caches. Athens returned modules after downloaded from upstream. So Athens will returns more quickly from next time.

Try again.

```
$ GOPROXY=http://athens:3000 docker-compose run --rm go-get
...
64.48user 61.88system 0:26.67elapsed 473%CPU (0avgtext+0avgdata 635104maxresident)k
232inputs+1546440outputs (3major+2207396minor)pagefaults 0swaps
```
