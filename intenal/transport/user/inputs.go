package user

type UserCreateInput struct {
	ID        int64
	Username  string   `json:"username,omitempty"`
	Location  Location `json:"location,omitempty"`
	DateBirth string   `json:"date_birth,omitempty"`
}
type Location struct {
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	City    string `json:"city,omitempty"`
}
