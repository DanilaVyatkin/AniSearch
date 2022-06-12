package cartoon

type Cartoon struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}
