package cartoon

type CreateCartoonDTO struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}
