package common

type RouteMetadata struct {
	RequiresAuth bool   `json:"requiresAuth"`
	Icon         string `json:"icon,omitempty"`
	Order        int    `json:"order"`
}
