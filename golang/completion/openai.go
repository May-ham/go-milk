package completion

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ParamsOpenAI struct {
	Api      `json:",inline"`
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature       float32 `json:"temperature"`
	Top_p             float32 `json:"top_p,omitempty"`
	Stream            bool    `json:"stream,omitempty"`
	Presence_penalty  float32 `json:"presence_penalty,omitempty"`
	Frequency_penalty float32 `json:"frequency_penalty,omitempty"`
}

func (params ParamsOpenAI) Completion() (interface{}, error) {

	url := "https://api.openai.com/v1/chat/completions"

	jsonParams, err := json.Marshal(params)
	if err != nil {
		// handle error
		return nil, err
	}

	payload := bytes.NewReader(jsonParams)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+params.Api.Key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	bodyRes, _ := io.ReadAll(res.Body)

	// Create an empty interface
	var result interface{}
	// Unmarshal the response body into the empty interface
	err = json.Unmarshal(bodyRes, &result)
	if err != nil {
		// handle error
		return nil, err
	}

	return result, nil
}
