# duplicate-request
Receive POST and duplicate this POST to others environments

# Start app-test
    $ cd app-test
    $ go run main.go

# Start main app
    $ cd ..
    $ export ENVIRONMENTS="http://localhost:8181/testepostA;http://localhost:8181/testepostB"
    $ go run main.go
    2024/07/09 17:59:42 Duplicator started

# Send test
    $ curl -L -d "name=foo&content=bar" -X POST http://localhost:8080/middleware
