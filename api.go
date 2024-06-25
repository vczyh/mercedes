package mercedes

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/google/uuid"
	"net/url"
)

type API struct {
	region Region
}

func NewAPI(opts ...APIOption) *API {
	api := &API{
		region: RegionChina,
	}
	for _, opt := range opts {
		opt.apply(api)
	}
	return api
}

func (a *API) Config(ctx context.Context) (*ConfigResponse, error) {
	var res ConfigResponse
	err := requests.
		URL(RegionProviders[a.region].ConfigURL).
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("ris-application-version", "1.41.2 (2167)").
		Header("X-SessionId", uuid.New().String()).
		Header("X-ApplicationName", "mycar-store-cn").
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) Login(ctx context.Context, emailOrPhoneNumber, nonce string) (*LoginResponse, error) {
	var res LoginResponse
	err := requests.
		URL(RegionProviders[a.region].LoginURL).
		Post().
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("X-SessionId", uuid.New().String()).
		BodyJSON(map[string]any{
			"emailOrPhoneNumber": emailOrPhoneNumber,
			"countryCode":        RegionProviders[a.region].CountryCode,
			"nonce":              nonce,
		}).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) OAuth2(ctx context.Context, email, nonce, password string) (*OAuth2Response, error) {
	var res OAuth2Response
	err := requests.
		URL(RegionProviders[a.region].OAuth2URL).
		Post().
		BodyForm(url.Values{
			"client_id":    {RegionProviders[a.region].ClientId},
			"grant_type":   {"password"},
			"password":     {fmt.Sprintf("%s:%s", nonce, password)},
			"scope":        {"openid email phone profile offline_access ciam-uid"},
			"username":     {email},
			"Stage":        {"prod"},
			"X-Device-Id":  {uuid.New().String()},
			"X-Request-Id": {uuid.New().String()},
		}).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

type ConfigResponse struct {
	ForceUpdate struct {
		Status   string `json:"status"`
		StoreUrl string `json:"storeUrl"`
	} `json:"forceUpdate"`
	VehicleImagesBaseUrl string `json:"vehicleImagesBaseUrl"`
}

type LoginResponse struct {
	IsEmail  bool   `json:"isEmail"`
	Username string `json:"username"`
}

type OAuth2Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func WithAPIRegion(region Region) APIOption {
	return APIOptionFun(func(api *API) {
		api.region = region
	})
}

type APIOption interface {
	apply(*API)
}

type APIOptionFun func(*API)

func (f APIOptionFun) apply(api *API) {
	f(api)
}
