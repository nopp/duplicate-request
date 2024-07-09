# duplicatePost
Receive POST and duplicate this POST to others environments

# Start app-test
    $ cd app-test
    $ go run main.go

# Start main app
    $ cd ..
    $ go run main.go
    $ go run main.go
    [GIN-debug] POST   /middleware               --> main.main.func1 (1 handlers)
    2024/07/09 17:59:42 Duplicator started
    [GIN-debug] Listening and serving HTTP on :8080
    2024/07/09 18:01:09 Send to environment A
    2024/07/09 18:01:09 "name=foo\u0026content=bar"
    2024/07/09 18:01:09 Send to environment B
    2024/07/09 18:01:09 "name=foo\u0026content=bar"

# Send test
    $ curl -L -d "name=foo&content=bar" -X POST http://localhost:8080/middleware
