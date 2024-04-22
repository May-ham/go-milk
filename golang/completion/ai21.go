package completion

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ParamsAI21 struct {
	Api              `json:",inline"`
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int32   `json:"maxTokens"`
	MinTokens        int32   `json:"minTokens"`
	Temperature      float32 `json:"temperature"`
	TopP             float32 `json:"topP"`
	FrequencyPenalty struct {
		Scale               float32 `json:"scale"`
		ApplyToWhitespaces  bool    `json:"applyToWhitespaces"`
		ApplyToPunctuations bool    `json:"applyToPunctuations"`
		ApplyToNumbers      bool    `json:"applyToNumbers"`
		ApplyToStopwords    bool    `json:"applyToStopwords"`
		ApplyToEmojis       bool    `json:"applyToEmojis"`
	} `json:"frequencyPenalty"`
	PresencePenalty struct {
		Scale               float32 `json:"scale"`
		ApplyToWhitespaces  bool    `json:"applyToWhitespaces"`
		ApplyToPunctuations bool    `json:"applyToPunctuations"`
		ApplyToNumbers      bool    `json:"applyToNumbers"`
		ApplyToStopwords    bool    `json:"applyToStopwords"`
		ApplyToEmojis       bool    `json:"applyToEmojis"`
	} `json:"presencePenalty"`
	CountPenalty struct {
		Scale               float32 `json:"scale"`
		ApplyToWhitespaces  bool    `json:"applyToWhitespaces"`
		ApplyToPunctuations bool    `json:"applyToPunctuations"`
		ApplyToNumbers      bool    `json:"applyToNumbers"`
		ApplyToStopwords    bool    `json:"applyToStopwords"`
		ApplyToEmojis       bool    `json:"applyToEmojis"`
	} `json:"countPenalty"`
}

func (params ParamsAI21) Completion() (interface{}, error) {

	url := "https://api.ai21.com/studio/v1/" + params.Model + "/complete"

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
