package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stephenwithav/weather/forecast"
)

// New
func New(appId string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 2,
		},
		appId: appId,
	}
}

type Client struct {
	client *http.Client
	appId  string
}

// Retrieve
func (c *Client) Retrieve(w http.ResponseWriter, r *http.Request) {
	resp, err := c.client.Get(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=imperial",
			r.FormValue("lat"), r.FormValue("long"), c.appId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fcast, err := forecast.New(buf)
	if err := json.NewEncoder(w).Encode(fcast); err != nil {
		panic(err)
	}
}
