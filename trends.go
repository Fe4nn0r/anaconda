package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetTrendsByPlace(id int64, v url.Values) (trendResponse []TrendResponse, err error) {
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/trends/place.json", v, &trendResponse, _GET, response_ch}
	return trendResponse, (<-response_ch).err
}

func (a TwitterApi) GetTrendsAvailable() (location []Location, err error) {
	response_ch := make(chan response)
	v := url.Values{}
	a.queryQueue <- query{BaseUrl + "/trends/available.json", v, &location, _GET, response_ch}
	return location, (<-response_ch).err
}

func (a TwitterApi) GetTrendsClosest(v url.Values) (location []Location, err error) {
	response_ch := make(chan response)

	v = cleanValues(v)

	a.queryQueue <- query{BaseUrl + "/trends/closest.json", v, &location, _GET, response_ch}
	return location, (<-response_ch).err
}
