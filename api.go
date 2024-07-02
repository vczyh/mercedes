package mercedes

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/google/uuid"
	"net/url"
	"time"
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

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type GetVehiclesResponse struct {
	AssignedVehicles []struct {
		AuthorizationType    string `json:"authorizationType"`
		Carline              string `json:"carline"`
		DataCollectorVersion string `json:"dataCollectorVersion"`
		Dealers              struct {
			Items []struct {
				DealerData struct {
					Address struct {
						City           string `json:"city"`
						CountryIsoCode string `json:"countryIsoCode"`
						Street         string `json:"street"`
						ZipCode        string `json:"zipCode"`
					} `json:"address"`
					Communication struct {
						Fax     string `json:"fax"`
						Phone   string `json:"phone"`
						Website string `json:"website"`
					} `json:"communication"`
					GeoCoordinates struct {
						Latitude  int `json:"latitude"`
						Longitude int `json:"longitude"`
					} `json:"geoCoordinates"`
					ID           string `json:"id"`
					LegalName    string `json:"legalName"`
					OpeningHours struct {
						Friday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"FRIDAY"`
						Monday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"MONDAY"`
						Saturday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"SATURDAY"`
						Sunday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"SUNDAY"`
						Thursday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"THURSDAY"`
						Tuesday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"TUESDAY"`
						Wednesday struct {
							Periods []struct {
								From  string `json:"from"`
								Until string `json:"until"`
							} `json:"periods"`
							Status string `json:"status"`
						} `json:"WEDNESDAY"`
					} `json:"openingHours"`
					Region struct {
						Region string `json:"region"`
					} `json:"region"`
					SpokenLanguages interface{} `json:"spokenLanguages"`
				} `json:"dealerData"`
				DealerID  string    `json:"dealerId"`
				Role      string    `json:"role"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"items"`
		} `json:"dealers"`
		Fin                             string `json:"fin"`
		IsOwner                         bool   `json:"isOwner"`
		IsTemporarilyAccessible         bool   `json:"isTemporarilyAccessible"`
		LicensePlate                    string `json:"licensePlate"`
		Mopf                            bool   `json:"mopf"`
		NormalizedProfileControlSupport string `json:"normalizedProfileControlSupport"`
		ProfileSyncSupport              string `json:"profileSyncSupport"`
		SalesRelatedInformation         struct {
			Baumuster struct {
				Baumuster            string `json:"baumuster"`
				BaumusterDescription string `json:"baumusterDescription"`
			} `json:"baumuster"`
			Line struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"line"`
			Paint struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"paint"`
			Upholstery struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"upholstery"`
		} `json:"salesRelatedInformation"`
		TirePressureMonitoringType string `json:"tirePressureMonitoringType"`
		TrustLevel                 int    `json:"trustLevel"`
		VehicleConnectivity        string `json:"vehicleConnectivity"`
		VehicleSegment             string `json:"vehicleSegment"`
		Vin                        string `json:"vin"`
		WindowsLiftCount           string `json:"windowsLiftCount"`
	} `json:"assignedVehicles"`
	Fleets []interface{} `json:"fleets"`
}

type GetUserDetailResponse struct {
	AccountCountryCode string `json:"accountCountryCode"`
	AccountIdentifier  string `json:"accountIdentifier"`
	AccountVerified    bool   `json:"accountVerified"`
	AdaptionValues     struct {
		Alias         string `json:"alias"`
		BodyHeight    int    `json:"bodyHeight"`
		PreAdjustment bool   `json:"preAdjustment"`
	} `json:"adaptionValues"`
	Address struct {
		AddressLine1  string `json:"addressLine1"`
		AddressLine2  string `json:"addressLine2"`
		AddressLine3  string `json:"addressLine3"`
		City          string `json:"city"`
		CountryCode   string `json:"countryCode"`
		DoorNo        string `json:"doorNo"`
		FloorNo       string `json:"floorNo"`
		HouseName     string `json:"houseName"`
		HouseNo       string `json:"houseNo"`
		PostOfficeBox string `json:"postOfficeBox"`
		Province      string `json:"province"`
		State         string `json:"state"`
		Street        string `json:"street"`
		StreetType    string `json:"streetType"`
		ZipCode       string `json:"zipCode"`
	} `json:"address"`
	AddressLines            []string `json:"addressLines"`
	Birthday                string   `json:"birthday"`
	CiamID                  string   `json:"ciamId"`
	CommunicationPreference struct {
		ContactedByEmail  bool `json:"contactedByEmail"`
		ContactedByLetter bool `json:"contactedByLetter"`
		ContactedByPhone  bool `json:"contactedByPhone"`
		ContactedBySms    bool `json:"contactedBySms"`
	} `json:"communicationPreference"`
	CreatedAt             time.Time `json:"createdAt"`
	CreatedBy             string    `json:"createdBy"`
	Email                 string    `json:"email"`
	FirstName             string    `json:"firstName"`
	IsEmailVerified       bool      `json:"isEmailVerified"`
	IsMobileVerified      bool      `json:"isMobileVerified"`
	LandlinePhone         string    `json:"landlinePhone"`
	LastName1             string    `json:"lastName1"`
	LastName2             string    `json:"lastName2"`
	MiddleInitial         string    `json:"middleInitial"`
	MobilePhoneNumber     string    `json:"mobilePhoneNumber"`
	NamePrefix            string    `json:"namePrefix"`
	Nickname              string    `json:"nickname"`
	PreferredLanguageCode string    `json:"preferredLanguageCode"`
	SalutationCode        string    `json:"salutationCode"`
	TaxNumber             string    `json:"taxNumber"`
	Title                 string    `json:"title"`
	UnitPreferences       struct {
		ClockHours                     string   `json:"clockHours"`
		ConsumptionCo                  string   `json:"consumptionCo"`
		ConsumptionEv                  string   `json:"consumptionEv"`
		ConsumptionGas                 string   `json:"consumptionGas"`
		RelevantConsumptionUnitOptions []string `json:"relevantConsumptionUnitOptions"`
		SpeedDistance                  string   `json:"speedDistance"`
		Temperature                    string   `json:"temperature"`
		TirePressure                   string   `json:"tirePressure"`
		UserLengthUnit                 string   `json:"userLengthUnit"`
	} `json:"unitPreferences"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UserPinStatus    string    `json:"userPinStatus"`
	UserPinUpdatedAt time.Time `json:"userPinUpdatedAt"`
}

type GetCapabilitiesResponse struct {
	Features struct {
		AuxHeat                                     bool `json:"auxHeat"`
		BidirectionalCharging                       bool `json:"bidirectionalCharging"`
		ChargingClockTimer                          bool `json:"chargingClockTimer"`
		ControllableRearWindowBlind                 bool `json:"controllableRearWindowBlind"`
		ControllableSunroof                         bool `json:"controllableSunroof"`
		Convertible                                 bool `json:"convertible"`
		DcCharging                                  bool `json:"dcCharging"`
		DistronicPro                                bool `json:"distronicPro"`
		DoubleDoorLock                              bool `json:"doubleDoorLock"`
		DriverAssistancePackageHigh                 bool `json:"driverAssistancePackageHigh"`
		DriverAssistancePackagePlus                 bool `json:"driverAssistancePackagePlus"`
		EcoCharging                                 bool `json:"ecoCharging"`
		FastCharging                                bool `json:"fastCharging"`
		HepaFilter                                  bool `json:"hepaFilter"`
		Mopf                                        bool `json:"mopf"`
		PictureTransfer                             bool `json:"pictureTransfer"`
		PluggedStateDependingPreEntryClimateControl bool `json:"pluggedStateDependingPreEntryClimateControl"`
		PrecondNow                                  bool `json:"precondNow"`
		RearSunProtectionBlinds                     bool `json:"rearSunProtectionBlinds"`
		RemoteSettingPersonalizedTemperature        bool `json:"remoteSettingPersonalizedTemperature"`
		RemoteSettingTemperature                    bool `json:"remoteSettingTemperature"`
		TwinTire                                    bool `json:"twinTire"`
		UrbanGuard                                  bool `json:"urbanGuard"`
		VariableOpenableSunroof                     bool `json:"variableOpenableSunroof"`
		VariableOpenableWindow                      bool `json:"variableOpenableWindow"`
		WeeklyProfile                               bool `json:"weeklyProfile"`
	} `json:"features"`
	Vehicle struct {
		Baumuster                      string        `json:"baumuster"`
		ChangeYearCodes                interface{}   `json:"changeYearCodes"`
		ControllableSunroofBlindsCount interface{}   `json:"controllableSunroofBlindsCount"`
		DigitalVehicleKeys             []string      `json:"digitalVehicleKeys"`
		DoorsCount                     int           `json:"doorsCount"`
		DoorsHandleType                string        `json:"doorsHandleType"`
		DrivingSide                    string        `json:"drivingSide"`
		ElectricVehicleType            interface{}   `json:"electricVehicleType"`
		ElectricWindowLifts            []string      `json:"electricWindowLifts"`
		FuelTypes                      []string      `json:"fuelTypes"`
		HeadUnitSoftwareVersion        string        `json:"headUnitSoftwareVersion"`
		HeadUnitType                   string        `json:"headUnitType"`
		ModelYearCode                  string        `json:"modelYearCode"`
		PowertrainBatteryModel         []interface{} `json:"powertrainBatteryModel"`
		ProductGroup                   string        `json:"productGroup"`
		RemoteSeatConfiguration        string        `json:"remoteSeatConfiguration"`
		StarArchitecture               string        `json:"starArchitecture"`
		SunroofType                    string        `json:"sunroofType"`
		TcuType                        string        `json:"tcuType"`
		TirePressureMonitorType        string        `json:"tirePressureMonitorType"`
	} `json:"vehicle"`
}

type GetCommandCapabilitiesResponse struct {
	Commands []struct {
		AdditionalInformation interface{} `json:"additionalInformation"`
		CapabilityInformation interface{} `json:"capabilityInformation"`
		CommandName           string      `json:"commandName"`
		IsAvailable           bool        `json:"isAvailable"`
		Parameters            interface{} `json:"parameters"`
	} `json:"commands"`
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
