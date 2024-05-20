package types

type RequestStore interface {
    CreateRequest(request Request) error
    GetAllRequests() ([]Request, error)
    UpdateRequest(request Request) error
}

type Request struct {
    ID string `json:"id"`
    Status string `json:"status"`
}
