package modal

type BookModal struct {
	ID           int                 `json:"id"`
	BookName     string              `json:"bookName"`
	Details      DetailsModal        `json:"details"`
	Available    bool                `json:"available"`
	Availability []AvailabilityModal `json:"availability"`
}

type AvailabilityModal struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
}

type DetailsModal struct {
	Author string `json:"author"`
	Seller string `json:"seller"`
}
