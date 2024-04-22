package tokenization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type TokenResponseAI21 struct {
	Text   string `json:"text"`
	Tokens []struct {
		Token     string `json:"token"`
		TextRange struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"textRange"`
	} `json:"tokens"`
}

func TokenizeAI21(w http.ResponseWriter, r *http.Request) {

	var data RequestData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(data.ApiKey)
	fmt.Println(data.Text)

	url := "https://api.ai21.com/studio/v1/tokenize"

	payload := strings.NewReader("{\"text\":\"" + data.Text + "\"}")

	fmt.Println(payload)

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", data.ApiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	var tokenRes TokenResponseAI21

	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		fmt.Println(err)
		return
	}

	tokenCount := len(tokenRes.Tokens)

	fmt.Printf(`The "%s" sentence has %d tokens`, tokenRes.Text, tokenCount)

	fmt.Fprintf(w, "Tokens in text: "+strconv.Itoa(tokenCount))
}
