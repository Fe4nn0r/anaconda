package anaconda

type Location struct {
	Country     string            `json:"country"`
	CountryCode string            `json:"countryCode"`
	Name        string            `json:"name"`
	parentid    int64             `json:"parentid"`
	PlaceType   LocationPlaceType `json:"placeType"`
	Url         string            `json:"url"`
	Woeid       int64             `json:"woeid"`
}

type LocationPlaceType struct {
	Name string `json:"name"`
	Code int64  `json:"code"`
}
