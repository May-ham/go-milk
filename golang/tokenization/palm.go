package tokenization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type TokenResponsePalm struct {
	TokenCount int `json:"tokenCount"`
}

func TokenizePalm(w http.ResponseWriter, r *http.Request) {

	var data RequestData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(data.ApiKey)
	fmt.Println(data.Text)

	url := "https://generativelanguage.googleapis.com/v1beta3/models/text-bison-001:countTextTokens?key=" + data.ApiKey

	payload := strings.NewReader("{\"prompt\":{\"text\":\"" + data.Text + "\"}}")

	fmt.Println(payload)

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	var tokenRes TokenResponsePalm

	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		fmt.Println(err)
		return
	}

	tokenCount := tokenRes.TokenCount

	fmt.Printf(`The "%s" sentence has %d tokens`, data.Text, tokenCount)

	fmt.Fprintf(w, "Tokens in text: "+strconv.Itoa(tokenCount))
}
