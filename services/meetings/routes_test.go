// go:build unittests
package meetings

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

func TestMeetingHandlers(t *testing.T) {
    meetingStore := &mockMeetingStore{}
    handler := NewHandler(meetingStore)

    t.Run("should return all meetings", func(t *testing.T) {

    })

    req, err := http.NewRequest(http.MethodGet, "/v1/meetings", nil)
    if err != nil {
        log.Println("error")
        return 
    }

    rr := httptest.NewRecorder()
    router := chi.NewRouter()

    router.HandleFunc("/v1/meetings", handler.getAllMeetings)
    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("[register] expected status code %d, got %d, -> %v", http.StatusOK, rr.Code, rr.Body)
    }
}


type mockMeetingStore struct {}

func (m *mockMeetingStore) GetAllMeetings() ([]types.MeetingResponse, error) {
    return nil, nil
}

func (m *mockMeetingStore) GetMeetingByID(id string) (types.Meeting, error) {
    return types.Meeting{}, nil
}

func (m *mockMeetingStore) CreateMeeting(meeting types.Meeting) error {
    return nil
}


func (m *mockMeetingStore) UpdateMeeting(meeting types.Meeting) error {
    return nil
}


func (m *mockMeetingStore) DeleteMeeting(id string) error {
    return nil
}
