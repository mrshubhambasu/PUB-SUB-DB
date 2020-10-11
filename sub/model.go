package main

type Data struct {
	Offers []Offers `json:"offers"`
}
type Hotel struct {
	HotelID     string  `json:"hotel_id"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Telephone   string  `json:"telephone"`
	Description string  `json:"description"`
	RoomCount   int     `json:"room_count"`
	Currency    string  `json:"currency"`
}
type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}
type Room struct {
	HotelID     string `json:"hotel_id"`
	RoomID      string `json:"room_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
type CancellationPolicy struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

type RatePlan struct {
	HotelID    string `json:"hotel_id"`
	RatePlanID string `json:"rate_plan_id"`

	Name string `json:"name"`

	MealPlan string `json:"meal_plan"`
}
type GuaranteePolicy struct {
	Required bool `json:"Required"`
}
type OriginalData struct {
	GuaranteePolicy GuaranteePolicy `json:"GuaranteePolicy"`
}
type Fees struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float64 `json:"percent"`
}
type Offers struct {
	CmOfferID    string       `json:"cm_offer_id"`
	Hotel        Hotel        `json:"hotel"`
	Room         Room         `json:"room"`
	RatePlan     RatePlan     `json:"rate_plan"`
	OriginalData OriginalData `json:"original_data"`
	Capacity     Capacity     `json:"capacity"`
	Number       int          `json:"number"`
	Price        int          `json:"price"`
	Currency     string       `json:"currency"`
	CheckIn      string       `json:"check_in"`
	CheckOut     string       `json:"check_out"`
	Fees         []Fees       `json:"fees"`
}
