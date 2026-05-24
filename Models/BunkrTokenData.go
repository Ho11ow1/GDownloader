package Models

type BunkrTokenData struct {
    Encrypted bool `json:"encrypted"`
    Timestamp int64 `json:"timestamp"`
    URL string `json:"url"`
}
