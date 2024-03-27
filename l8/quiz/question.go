package quiz

type Question struct {
	Text    string   `json:"Text"`
	Options []string `json:"Options"`
}
