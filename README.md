## Mercedes SDK

- Receive push messages from server.
- Control your car by sending command to server. (developing)


## Usage

```go
var f EventListenFun = func(event MercedesEvent, err error) {
    if err != nil {
        t.Error(err)
    }
    t.Logf("%s %T %+v\n", time.Now(), event, event)
}

c := mercedes.New(
    mercedes.WithAccessToken("access_token"),
    mercedes.WithRefreshToken("refresh_token"),
    WithEventListen(f),
)
_ = c.Connect(context.TODO())

_ = c.UnLock("vin", "pin")

<-make(chan struct{})
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