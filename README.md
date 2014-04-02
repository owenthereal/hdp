HTTP Debug Proxy (HDP)
======================

HTTP Debug Proxy (HDP) logs HTTP requests for inspection.
It makes testing an HTTP library easier.

Installation
------------

### Binary

[![Gobuild Download](http://gobuild.io/badge/github.com/jingweno/hdp/download.png)](http://gobuild.io/github.com/jingweno/hdp)

### Source

```
go install github.com/jingweno/hdp
```

Usage
-----

```
$ hdp
HDP listening on localhost:12345
# `curl http://localhost:12345` in another terminal session
> GET / HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:12345
> Accept: */*
```

See a [demo](http://showterm.io/d4b406229385b2d7cc8c0).

License
-------

`hdp` is released under the MIT license. See [LICENSE.md](https://github.com/jingweno/hdp/blob/master/LICENSE.md).
