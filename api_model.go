package mercedes

import "time"

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
