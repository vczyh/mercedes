## Mercedes SDK

- Receive push messages from server.
- Control your car by sending command to server. (developing)


## Usage

```go
var f mercedes.EventListenFun = func(event mercedes.MercedesEvent, err error) {
    if err != nil {
        return
    }
    fmt.Printf("%s %T %+v\n", time.Now(), event, event)
}

c := mercedes.New(
    mercedes.WithAccessToken("access_token"),
    mercedes.WithRefreshToken("refresh_token"),
    mercedes.WithEventListen(f),
)
ctx := context.TODO()

_ = c.Connect(ctx)
defer c.Close()

_ = c.UnLock(ctx, "vin", "pin")
```

## Get Token

```shell
go install github.com/vczyh/mercedes/cmd/mercedes
```

Then login and get token:

```shell
mercedes --debug login <email>
```

## Thanks 

- [ReneNulschDE/mbapi2020](https://github.com/ReneNulschDE/mbapi2020)
- [mercedes-benz/MBSDK-CarKit-iOS](https://github.com/mercedes-benz/MBSDK-CarKit-iOS)