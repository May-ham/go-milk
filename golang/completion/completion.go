package completion

type Completion interface {
	Completion() (interface{}, error)
}

// TODO: suss out will api key be fetched as a struct or as a plain variable if I'll send a JSON to my server

type Api struct {
	Key string `json:"apiKey"`
}

func Complete(c Completion) {
	c.Completion()
}
