package main

type VehicleBaseInfo struct {
	VIN    string `xml:"KoeretoejOplysningStelNummer"`
	Status string `xml:"KoeretoejOplysningStatus"`

	FirstRegistered string `xml:"KoeretoejOplysningFoersteRegistreringDato"`

	// Weight + Pull
	WeightSelf  int `xml:"KoeretoejOplysningEgenVaegt"`
	WeightTotal int `xml:"KoeretoejOplysningTekniskTotalVaegt"`
	WeightMax   int `xml:"KoeretoejOplysningTotalVaegt"`
	Axes        int `xml:"KoeretoejOplysningAkselAntal"`

	// Brand, model, color
	Brand   string `xml:"KoeretoejBetegnelseStruktur>KoeretoejMaerkeTypeNavn"`
	Model   string `xml:"KoeretoejBetegnelseStruktur>Model>KoeretoejModelTypeNavn"`
	Variant string `xml:"KoeretoejBetegnelseStruktur>Variant>KoeretoejVariantTypeNavn"`
	Type    string `xml:"KoeretoejBetegnelseStruktur>Type>KoeretoejTypeTypeNavn"`
	Color   string `xml:"KoeretoejFarveStruktur>FarveTypeStruktur>FarveTypeNavn"`

	// Engine
	EngineMileage    string  `xml:"KoeretoejMotorStruktur>KoeretoejMotorKmPerLiter"`
	EngineKilometers float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorKilometerstand"`
	EngineFuel       string  `xml:"KoeretoejMotorStruktur>DrivkraftTypeNavn"`
	EngineCylinders  int     `xml:"KoeretoejMotorStruktur>KoeretoejMotorCylinderAntal"`
	EngineVolume     float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorSlagVolumen"`
	EngineEffect     float32 `xml:"KoeretoejMotorStruktur>KoeretoejMotorStoersteEffekt"`

	ChassisType string `xml:"KarrosseriTypeStruktur>KarrosseriTypeNavn"`
}

type Vehicle struct {
	Plate       string `xml:"RegistreringNummerNummer" gorm:"primaryKey"`
	VehicleType string `xml:"KoeretoejArtNavn"`

	BaseInfo VehicleBaseInfo `xml:"KoeretoejOplysningGrundStruktur" gorm:"embedded"`

	End   string `xml:"RegistreringNummerUdloebDato"`
	Usage string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseNavn"`
}
