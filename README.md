## Mercedes SDK

- Receive push messages from server.
- Control your car by sending command to server. (developing)


```go
c := New(
    WithAccessToken("access_token"),
    WithRefreshToken("refresh_token"),
)
_ = c.Connect(context.TODO())

_ = c.OnListen(func(event MercedesEvent) {
    fmt.Printf("%s %T %+v\n", time.Now(), event, event)
})
```