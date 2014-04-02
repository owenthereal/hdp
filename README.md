HTTP Debug Proxy (HDP)
======================

HTTP Debug Proxy (HDP) logs HTTP requests for inspection.
It makes testing an HTTP library easier.

Usage
-----

```
$ hdp
HDP listening on localhost:12345
$ curl http://localhost:12345
> GET / HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:12345
> Accept: */*
```

License
-------

`hdp` is released under the MIT license. See[LICENSE.md](https://github.com/jingweno/hdp/blob/master/LICENSE.md).
