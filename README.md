# AppRestInventory

RESTful API of the inventory control system GoStock.

Access to the functional [demo](https://app-rest-inventory.herokuapp.com/swagger).

## Installing

### Installing Beego

Follow the next steps to install the framework Beego in Cloud9. See the official
guide here https://beego.me/docs/install/bee.md

1. Run the following command.

```
$ go get github.com/beego/bee
```

2. bee is installed in GOBIN by default, so we need to add GOBIN to the PATH, 
otherwise the bee command won’t work. To do that we need modify the source 
~/.profile and add the following lines at the end.

```
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

3. Run the bee command to verify.

```
$ bee
```

## Deployment

### Development server

**NOTE:** The `app.conf` file will not be pushed to the public repository for security. It will be added to the resources folder in the final package.

Run `bee run` for a development server. Navigate to `https://{HOST}:{PORT}/swagger` to see the API definition.

## Testing

### Running Tests

```
go test -run {TEST_NAME}
```