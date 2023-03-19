package city

type City struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Region     string `json:"region,omitempty"`
	District   string `json:"district,omitempty"`
	Population int    `json:"population,omitempty"`
	Foundation int    `json:"foundation,omitempty"`
}
