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

type PostBook struct {
	BookName  *string `json:"bookName"`
	Available *bool   `json:"available"`
	Location  *string `json:"location"`
	Author    *string `json:"author"`
	Seller    *string `json:"seller"`
}
