# Assignment

The goal of this assignment is to make the main.go work.

Run the integration via either 
```shell
go run cmd/main.go
```
or 
```shell
make all
```
## Stepwise implementation

1. Familiarize yourself with Chargeamps REST API. Documentation is available here https://eapi.charge.space/swagger/index.html
2. Using username and password provided separately, complete TokenSource with appropriate code that retrieves access 
   token from charge amps and passes it forward.
3. Using the access tokens implement DeviceListProvider interface that retrieves device list belonging to given 
   charge amps user.
4. Using device ID retrieved and access tokens, retrieve charger status.

## Additional stuff

In order for code to look neat, use golang linter. Can be installed here https://golangci-lint.run/welcome/install/
if you don't already have it. After that the following make should work.
```shell
make lint
```

