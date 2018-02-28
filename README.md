# ceflog

> Common Event Format Logger

## about

The ArcSight Common Event Format can be used to emit audit logs in a format
which threat analysis and monitor tools are able to ingest.

## usage

```go
logger := ceflog.New(w, "vendor", "product", "version")

logger.Event(
    "User login",
    "auth.new",
    ceflog.Sev(0),
    ceflog.Ext("dst", "127.0.0.1"),
)
```

More documentation can be found in the [GoDoc][godoc].

[godoc]: https://godoc.org/github.com/xoebus/ceflog
