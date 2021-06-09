package modal

type BookModal struct {
	ID           int                 `json:"id"`
	Book         string              `json:"book"`
	Author       AuthorModal         `json:"author"`
	Available    bool                `json :"available"`
	Availability []AvailabilityModal `json:"availability"`
}

type AvailabilityModal struct {
	ID       int    `json:"id"`
	location string `json:"location"`
}

type AuthorModal struct {
	Name string `json:"name"`
}
