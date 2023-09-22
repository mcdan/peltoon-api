package dto

type Instructor struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Ride struct {
	ClassTypeIDs []string   `json:"class_type_ids"`
	Description  string     `json:"description"`
	HomeID       string     `json:"home_peloton_id"`
	ID           string     `json:"id"`
	StudioID     string     `json:"studio_peloton_id"`
	Title        string     `json:"title"`
	Instructor   Instructor `json:"instructor"`
}
type RideDetails struct {
	Ride Ride `json:"ride"`
}
