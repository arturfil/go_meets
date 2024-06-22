package types

type RequestStore interface {
	CreateRequest(request Request) error
	GetAllRequests() ([]RequestResponse, error)
    GetRequestById(id string) (*RequestResponse, error)
	UpdateRequest(id string, request Request) error
}

type Request struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type RequestResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}
