package types

type RequestStore interface {
	CreateRequest(request Request) error
	GetAllRequests() ([]RequestResponse, error)
	GetRequestById(id, queryType string) (*RequestResponse, error)
	UpdateRequest(id string, request Request) error
}

type Request struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Status string `json:"status,omitempty"`
	Type   string `json:"type"`
	Value  string `json:"value,omitempty"`
}

type RequestResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Value     string `json:"value,omitempty"`
}
