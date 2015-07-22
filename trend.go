package anaconda

import (
	"time"
)

type Trend struct {
	Events          string `json:"events"`
	Name            string `json:"name"`
	PromotedContent string `json:"promoted_content"`
	Query           string `json:"query"`
	Url             string `json:"url"`
}

type TrendResponse struct {
	AsOf      string              `json:"as_of"`
	CreatedAt string              `json:"created_at"`
	Locations []LocationPlaceType `json:"locations"`
	Trends    []Trend             `json:"trends"`
}

// CreatedAtTime is a convenience wrapper that returns the CreatedAt time, parsed as a time.Time struct
func (tr TrendResponse) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, tr.CreatedAt)
}
