package models

type LocationsResponse struct {
	Locations Locations `json:"locations"`
}
type Locations struct {
	Edges []struct {
		Node struct {
			ID string `json:"id"`
		} `json:"node"`
	} `json:"edges"`
}
