package mercedes

type Region int

const (
	RegionChina Region = iota
)

var (
	RegionProviderChina = RegionProvider{
		ConfigURL:    "https://bff.cn-prod.mobilesdk.mercedes-benz.com/v1/config",
		LoginURL:     "https://bff.cn-prod.mobilesdk.mercedes-benz.com/v1/login",
		WebSocketURL: "wss://websocket.cn-prod.mobilesdk.mercedes-benz.com/v2/ws",
		OAuth2URL:    "https://ciam-1.mercedes-benz.com.cn/as/token.oauth2",
		XLocal:       " zh-CN",
		CountryCode:  "CN",
		ClientId:     "3f36efb1-f84b-4402-b5a2-68a118fec33e",
	}

	RegionProviders = map[Region]RegionProvider{
		RegionChina: RegionProviderChina,
	}
)

type RegionProvider struct {
	ConfigURL    string
	LoginURL     string
	WebSocketURL string
	OAuth2URL    string
	XLocal       string
	CountryCode  string
	ClientId     string
}
