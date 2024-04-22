package completion

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ParamsPalm struct {
	Api    `json:",inline"`
	Prompt struct {
		Text string `json:"text"`
	} `json:"prompt"`
	// SafetySettings []struct {

	// } `json:"safetySettings"`
	Temperature     float32 `json:"temperature"`
	CandidateCount  int32   `json:"candidateCount"`
	MaxOutputTokens int32   `json:"maxOutputTokens"`
	TopP            int32   `json:"topP"`
	TopK            int32   `json:"topK"`
}

func (params ParamsPalm) Completion() (interface{}, error) {

	jsonParams, err := json.Marshal(params)
	if err != nil {
		// handle error
		return nil, err
	}

	url := "https://generativelanguage.googleapis.com/v1beta2/models/text-bison-001:generateText?key=" + params.Api.Key

	payload := bytes.NewReader(jsonParams)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

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
