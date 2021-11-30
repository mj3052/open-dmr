package main

type VehicleBaseInfo struct {
	VIN    string `xml:"KoeretoejOplysningStelNummer" json:"vin"`
	Status string `xml:"KoeretoejOplysningStatus" json:"status"`

	FirstRegistered string `xml:"KoeretoejOplysningFoersteRegistreringDato" json:"first_registered"`

	// Weight + Pull
	WeightSelf  int `xml:"KoeretoejOplysningEgenVaegt" json:"weight_self"`
	WeightTotal int `xml:"KoeretoejOplysningTekniskTotalVaegt" json:"weight_total"`
	WeightMax   int `xml:"KoeretoejOplysningTotalVaegt" json:"weight_max"`
	Axes        int `xml:"KoeretoejOplysningAkselAntal" json:"axes"`

	// Brand, model, color
	Brand   string `xml:"KoeretoejBetegnelseStruktur>KoeretoejMaerkeTypeNavn" json:"brand"`
	Model   string `xml:"KoeretoejBetegnelseStruktur>Model>KoeretoejModelTypeNavn" json:"model"`
	Variant string `xml:"KoeretoejBetegnelseStruktur>Variant>KoeretoejVariantTypeNavn" json:"variant"`
	Type    string `xml:"KoeretoejBetegnelseStruktur>Type>KoeretoejTypeTypeNavn" json:"type"`
	Color   string `xml:"KoeretoejFarveStruktur>FarveTypeStruktur>FarveTypeNavn" json:"color"`

	// Engine
	EngineMileage    string  `xml:"KoeretoejMotorStruktur>KoeretoejMotorKmPerLiter" json:"engine_mileage"`
	EngineKilometers float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorKilometerstand" json:"engine_kilometers"`
	EngineFuel       string  `xml:"KoeretoejMotorStruktur>DrivkraftTypeNavn" json:"engine_fuel"`
	EngineCylinders  int     `xml:"KoeretoejMotorStruktur>KoeretoejMotorCylinderAntal" json:"engine_cylinders"`
	EngineVolume     float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorSlagVolumen" json:"engine_volume"`
	EngineEffect     float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorStoersteEffekt" json:"engine_effect"`

	ChassisType string `xml:"KarrosseriTypeStruktur>KarrosseriTypeNavn" json:"chassis_type"`
}

type Vehicle struct {
	Plate       string `xml:"RegistreringNummerNummer" gorm:"primaryKey" json:"plate"`
	VehicleType string `xml:"KoeretoejArtNavn" json:"vehicle_type"`

	BaseInfo VehicleBaseInfo `xml:"KoeretoejOplysningGrundStruktur" gorm:"embedded" json:"info"`

	End   string `xml:"RegistreringNummerUdloebDato" json:"expiry_date"`
	Usage string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseNavn" json:"usage"`
}
