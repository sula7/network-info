# network-info

### Program run arguments
This app could be started using args. Just set IP addresses in args like

`go run ./main.go 192.168.0.1 192.168.0.2`

or run executable with the same args

### Run
This app sends ICMP packets (ping) to get some info and need higher privileges.

For Windows just run executable as administrator (right click -> Run as administrator)

For Linux run with `sudo`

### ENV
````
BIND_PORT=      Listen to port. Default is :8080
````

### Endpoints
````
/ping
Responses by string "pong" and code 200
Use for server availability test
````
````
/api/v1/addresses
Responses by string "OK" and code 200
Use to set list of IP address

Request should be:
Header
Content-Type: Application/JSON

Body
[
    {
        "ip":""
    }
    {
        "ip":""
    }
    ...
]
````
