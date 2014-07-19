# Example Vedis API server (KvStore/KvFetch)

Just a simple interface to [KvStore](http://vedis.symisc.net/c_api/vedis_kv_store.html)
and [KvFetch](http://vedis.symisc.net/c_api/vedis_kv_fetch.html) over HTTP.

## Benchmarks

Starting the server

``` bash
$ GOMAXPROCS=2 go run main.go -db="data.vdb"
```

Storing some data

``` bash
curl -X POST -d "xyz123" http://localhost:6600/abc
```

Running a latency benchmark using [wrk](https://github.com/wg/wrk)

``` bash
$ wrk --latency -c 100 -d 20s -t 2 http://localhost:6600/abc
Running 20s test @ http://localhost:6600/abc
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.75ms  332.78us  10.19ms   80.66%
    Req/Sec    29.65k     2.51k   38.67k    69.98%
  Latency Distribution
     50%    1.75ms
     75%    1.87ms
     90%    2.05ms
     99%    2.77ms
  1121444 requests in 20.00s, 130.48MB read
Requests/sec:  56072.75
Transfer/sec:      6.52MB
```
