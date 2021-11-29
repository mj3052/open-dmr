package main

type VehicleBaseInfo struct {
	VIN         string `xml:"KoeretoejOplysningStelNummer"`
	Status      string `xml:"KoeretoejOplysningStatus"`
	WeightSelf  int    `xml:"KoeretoejOplysningEgenVaegt"`
	WeightTotal int    `xml:"KoeretoejOplysningTekniskTotalVaegt"`
	WeightMax   int    `xml:"KoeretoejOplysningTotalVaegt"`
	Axes        int    `xml:"KoeretoejOplysningAkselAntal"`

	VehicleBrand   string `xml:"KoeretoejBetegnelseStruktur>KoeretoejMaerkeTypeNavn"`
	VehicleModel   string `xml:"KoeretoejBetegnelseStruktur>Model>KoeretoejModelTypeNavn"`
	VehicleVariant string `xml:"KoeretoejBetegnelseStruktur>Variant>KoeretoejVariantTypeNavn"`
	VehicleType    string `xml:"KoeretoejBetegnelseStruktur>Type>KoeretoejTypeTypeNavn"`

	VehicleColor string `xml:"KoeretoejFarveStruktur>FarveTypeStruktur>FarveTypeNavn"`
}

type Vehicle struct {
	Plate       string `xml:"RegistreringNummerNummer" gorm:"primaryKey"`
	VehicleType string `xml:"KoeretoejArtNavn"`

	BaseInfo VehicleBaseInfo `xml:"KoeretoejOplysningGrundStruktur" gorm:"embedded"`

	End string `xml:"RegistreringNummerUdloebDato"`

	Usage string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseNavn"`
}
