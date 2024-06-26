package dadjokes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jsphbtst/spongebobcase/pkg/types"
)

func Get() (*types.DadJoke, error) {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dadJoke types.DadJoke
	err = json.Unmarshal(body, &dadJoke)
	if err != nil {
		return nil, err
	}

	return &dadJoke, nil
}
