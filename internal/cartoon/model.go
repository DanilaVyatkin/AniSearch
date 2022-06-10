package cartoon

type Cartoon struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}
