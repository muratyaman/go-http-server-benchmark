# go-http-server-benchmark

Comparison of performance - requests per second - using Go HTTP server, text, JSON, PostgreSQL query results.

## Requirements

* [Go](https://go.dev/)

## Build

```sh
go build main.go
```

## Execution

Terminal 1:

```sh
./main
```

Terminal 2:

```sh
autocannon http://localhost:8000/
autocannon http://localhost:8000/json
autocannon http://localhost:8000/sql1
autocannon http://localhost:8000/sql2
```

## Summary

Average Req/Sec:

| Text | JSON | PgSQL 1 | PgSQL 2 |
| - | - | - | - |
| 109k | 107k | 46k | ?k |

## Results for Hello world - Text output

Send fixed/same text all the time 'Hello world'.

```
Running 10s test @ http://localhost:8080
10 connections

┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬──────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max  │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼──────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.02 ms │ 5 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴──────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬───────────┬─────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg       │ Stdev   │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼───────────┼─────────┼─────────┤
│ Req/Sec   │ 102847  │ 102847  │ 110527  │ 111359  │ 109262.55 │ 2722.29 │ 102821  │
├───────────┼─────────┼─────────┼─────────┼─────────┼───────────┼─────────┼─────────┤
│ Bytes/Sec │ 13.5 MB │ 13.5 MB │ 14.5 MB │ 14.6 MB │ 14.3 MB   │ 354 kB  │ 13.5 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴───────────┴─────────┴─────────┘

Req/Bytes counts sampled once per second.
# of samples: 11

1202k requests in 11.01s, 157 MB read
```

## Results for simple JSON message

```go
tsObj := map[string]interface{}{
	"message": "Hello world",
	"ts":      time.Now().String(),
}
```

```
Running 10s test @ http://localhost:8080/json
10 connections

┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬──────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max  │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼──────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.02 ms │ 5 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴──────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev   │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
│ Req/Sec   │ 102079  │ 102079  │ 108287  │ 109887  │ 107488  │ 2468.63 │ 102018  │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
│ Bytes/Sec │ 20.6 MB │ 20.6 MB │ 21.9 MB │ 22.2 MB │ 21.7 MB │ 500 kB  │ 20.6 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┘

Req/Bytes counts sampled once per second.
# of samples: 11

1182k requests in 11.01s, 239 MB read
```

## Results for time from PostgreSQL v1

Send result of query:

```sql
SELECT CAST(now() AS VARCHAR)
```

```
Running 10s test @ http://localhost:8080/sql1
10 connections

┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬──────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max  │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼──────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.03 ms │ 5 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴──────┘
┌───────────┬─────────┬─────────┬─────────┬────────┬──────────┬────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%  │ Avg      │ Stdev  │ Min     │
├───────────┼─────────┼─────────┼─────────┼────────┼──────────┼────────┼─────────┤
│ Req/Sec   │ 44319   │ 44319   │ 46719   │ 46975  │ 46398.55 │ 732.44 │ 44318   │
├───────────┼─────────┼─────────┼─────────┼────────┼──────────┼────────┼─────────┤
│ Bytes/Sec │ 7.93 MB │ 7.93 MB │ 8.36 MB │ 8.4 MB │ 8.3 MB   │ 131 kB │ 7.93 MB │
└───────────┴─────────┴─────────┴─────────┴────────┴──────────┴────────┴─────────┘

Req/Bytes counts sampled once per second.
# of samples: 11

510k requests in 11.01s, 91.3 MB read
```

## Results for random list of records from PostgreSQL v2

TODO
