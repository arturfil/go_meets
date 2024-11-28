package integration_tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arturfil/meetings_app_server/services/user"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersIntegration(t *testing.T) {
	// Verify testDB is not nil
	require.NotNil(t, testDB, "testDB should not be nil")

	store := user.NewStore(testDB)
	handler := user.NewHandler(store)
	router := chi.NewRouter()
	handler.RegisterRoutes(router)

	tests := []struct {
		name           string
		method         string
		path           string
		body           string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:   "Sign Up User - Success",
			method: "POST",
			path:   "/v1/auth/signup",
			body: `{
                "first_name": "John",
                "last_name": "Doe",
                "email": "john@example.com",
                "password": "password123"
            }`,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:   "Login User - Success",
			method: "POST",
			path:   "/v1/auth/login",
			body: `{
                "email": "john@example.com",
                "password": "password123"
            }`,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response map[string]string
				err := json.NewDecoder(rr.Body).Decode(&response)
				require.NoError(t, err)
				assert.Contains(t, response, "token")
				assert.NotEmpty(t, response["token"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req, err := http.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			require.NoError(t, err)

			if tt.body != "" {
				req.Header.Set("Content-Type", "application/json")
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(rr, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr)
			}
		})
	}
}
