package mercedes

type Region int

const (
	RegionChina Region = iota
)

var (
	RegionProviderChina = RegionProvider{
		WebSocketURL: "wss://websocket.cn-prod.mobilesdk.mercedes-benz.com/v2/ws",
		OAuth2URL:    "https://ciam-1.mercedes-benz.com.cn/as/token.oauth2",
	}

	RegionProviders = map[Region]RegionProvider{
		RegionChina: RegionProviderChina,
	}
)

type RegionProvider struct {
	WebSocketURL string
	OAuth2URL    string
}
