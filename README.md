# network-info

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
