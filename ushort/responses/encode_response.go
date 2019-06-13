package responses

import (
	"encoding/json"
	"time"
)

// EncoderResponse Response for shortUrl
type EncoderResponse struct {
	OriginalURL string    `json:"url"`
	ShortURL    string    `json:"shortUrl"`
	CreatedTime time.Time `json:"createdAt"`
}

// NewEncoderResponse creates new instance of EncoderResponse
func NewEncoderResponse(ourl, surl string) *EncoderResponse {
	return &EncoderResponse{OriginalURL: ourl, ShortURL: surl, CreatedTime: time.Now()}
}

// JSON converts to json
func (er *EncoderResponse) JSON() (string, error) {
	erBytes, err := json.Marshal(er)
	var res string
	if err == nil {
		res = string(erBytes)
	}
	return res, err
}
