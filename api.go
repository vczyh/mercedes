package mercedes

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/google/uuid"
	"net/url"
)

const (
	CommandNameAuxHeatConfigure                      = "AUXHEAT_CONFIGURE"
	CommandNameAuxHeatStart                          = "AUXHEAT_START"
	CommandNameAuxHeatStop                           = "AUXHEAT_STOP"
	CommandNameBatteryChargeProgramConfigure         = "BATTERY_CHARGE_PROGRAM_CONFIGURE"
	CommandNameBatteryMaxSocConfigure                = "BATTERY_MAX_SOC_CONFIGURE"
	CommandNameChargeOptConfigure                    = "CHARGE_OPT_CONFIGURE"
	CommandNameChargeOptStart                        = "CHARGE_OPT_START"
	CommandNameChargeOptStop                         = "CHARGE_OPT_STOP"
	CommandNameChargeProgramConfigure                = "CHARGE_PROGRAM_CONFIGURE"
	CommandNameChildPresenceDetectionDeactivateAlarm = "CHILDPRESENCEDETECTION_DEACTIVATEALARM"
	CommandNameRemoteUpdateStart                     = "REMOTE_UPDATE_START"
	CommandNameDoorsLock                             = "DOORS_LOCK"
	CommandNameDoorsUnlock                           = "DOORS_UNLOCK"
	CommandNameEngineStart                           = "ENGINE_START"
	CommandNameEngineStop                            = "ENGINE_STOP"
	CommandNameTheftAlarmDeselectInterior            = "THEFTALARM_DESELECT_INTERIOR"
	CommandNameTheftAlarmDeselectTow                 = "THEFTALARM_DESELECT_TOW"
	CommandNameTheftAlarmSelectInterior              = "THEFTALARM_SELECT_INTERIOR"
	CommandNameTheftAlarmSelectTow                   = "THEFTALARM_SELECT_TOW"
	CommandNameTheftAlarmStart                       = "THEFTALARM_START"
	CommandNameTheftAlarmStop                        = "THEFTALARM_STOP"
	CommandNameTheftAlarmConfirmDamagedetection      = "THEFTALARM_CONFIRM_DAMAGEDETECTION"
	CommandNameTheftAlarmDeselectDamagedetection     = "THEFTALARM_DESELECT_DAMAGEDETECTION"
	CommandNameTheftAlarmSelectDamagedetection       = "THEFTALARM_SELECT_DAMAGEDETECTION"
	CommandNameTheftAlarmSelectPictureTaking         = "THEFTALARM_SELECT_PICTURE_TAKING"
	CommandNameTheftAlarmDeselectPictureTaking       = "THEFTALARM_DESELECT_PICTURE_TAKING"
	CommandNameSunroofOpen                           = "SUNROOF_OPEN"
	CommandNameSunroofLift                           = "SUNROOF_LIFT"
	CommandNameSunroofClose                          = "SUNROOF_CLOSE"
	CommandNameSpeedAlertStart                       = "SPEEDALERT_START"
	CommandNameSpeedAlertStop                        = "SPEEDALERT_STOP"
	CommandNameWeekProfileConfigure                  = "WEEK_PROFILE_CONFIGURE"
	CommandNameWeekProfileConfigureV2                = "WEEK_PROFILE_CONFIGURE_V2"
	CommandNameChargingClockTimerConfigure           = "CHARGING_CLOCK_TIMER_CONFIGURE"
	CommandNameChargingConfigure                     = "CHARGING_CONFIGURE"
	CommandNameWindowsOpen                           = "WINDOWS_OPEN"
	CommandNameWindowsClose                          = "WINDOWS_CLOSE"
	CommandNameWindowsVentilate                      = "WINDOWS_VENTILATE"
	CommandNameWiperHealthReset                      = "WIPER_HEALTH_RESET"
	CommandNameZevPreconditionConfigure              = "ZEV_PRECONDITION_CONFIGURE"
	CommandNameZevPreconditionConfigureSeats         = "ZEV_PRECONDITION_CONFIGURE_SEATS"
	CommandNameZevPreconditioningStart               = "ZEV_PRECONDITIONING_START"
	CommandNameZevPreconditioningStop                = "ZEV_PRECONDITIONING_STOP"
	CommandNameSigposStart                           = "SIGPOS_START"
	CommandNameTemperatureConfigure                  = "TEMPERATURE_CONFIGURE"
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

func (a *API) RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error) {
	var res RefreshTokenResponse
	err := requests.
		URL(RegionProviders[a.region].OAuth2URL).
		Post().
		ContentType("application/x-www-form-urlencoded").
		BodyForm(url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {refreshToken},
			"X-Device-Id":   {uuid.New().String()},
			"X-Request-Id":  {uuid.New().String()},
		}).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) GetVehicles(ctx context.Context, token string) (*GetVehiclesResponse, error) {
	var res GetVehiclesResponse
	err := requests.
		URL(RegionProviders[a.region].ResetURL).
		Path("/v2/vehicles").
		Header("Authorization", "Bearer "+token).
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("X-SessionId", uuid.New().String()).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) GetUserDetail(ctx context.Context, token string) (*GetUserDetailResponse, error) {
	var res GetUserDetailResponse
	err := requests.
		URL(RegionProviders[a.region].ResetURL).
		Path("/v1/user").
		Header("Authorization", "Bearer "+token).
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("X-SessionId", uuid.New().String()).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) GetCapabilities(ctx context.Context, token, vin string) (*GetCapabilitiesResponse, error) {
	var res GetCapabilitiesResponse
	err := requests.
		URL(RegionProviders[a.region].ResetURL).
		Pathf("/v1/vehicle/%s/capabilities", vin).
		Header("Authorization", "Bearer "+token).
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("X-SessionId", uuid.New().String()).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

func (a *API) GetCommandCapabilities(ctx context.Context, token, vin string) (*GetCommandCapabilitiesResponse, error) {
	var res GetCommandCapabilitiesResponse
	err := requests.
		URL(RegionProviders[a.region].ResetURL).
		Pathf("/v1/vehicle/%s/capabilities/commands", vin).
		Header("Authorization", "Bearer "+token).
		Header("X-TrackingId", uuid.New().String()).
		Header("ris-os-name", "ios").
		Header("ris-os-version", "17.5.1").
		Header("ris-sdk-version", "2.113.0").
		Header("X-Locale", RegionProviders[a.region].XLocal).
		Header("X-SessionId", uuid.New().String()).
		ToJSON(&res).
		Fetch(ctx)
	return &res, err
}

type APIOption interface {
	apply(*API)
}

type APIOptionFun func(*API)

func (f APIOptionFun) apply(api *API) {
	f(api)
}
