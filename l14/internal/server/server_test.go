package server

import (
	"bytes"
	"encoding/json"
	"hw15/internal/primitives"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupServer() *Server {
	server := Server{}
	mux := server.GetRouter()
	go http.ListenAndServe(":8000", mux)
	return &server
}

func newUser(t *testing.T) (*http.Request, *httptest.ResponseRecorder) {
	userReq := UserCreationRequest{
		Nickname: "Artem",
	}
	data, err := json.Marshal(userReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/user", bytes.NewReader(data))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	return req, rr
}

// router.HandleFunc("POST /user", s.handlePostUser) // register a user
func TestHandlerPostUser(t *testing.T) {
	server := setupServer()
	req, rec := newUser(t)

	server.handlePostUser(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	require.NotNil(t, rec.Body)

	bodyBytes, err := io.ReadAll(rec.Body)
	require.NoError(t, err)
	idx := strings.LastIndex(string(bodyBytes), `"status":"ok"`)
	require.NotEqual(t, idx, -1)
}

// router.HandleFunc("POST /user/{id}/routes", s.handlePostRoute) // create route
func TestCreateRoute(t *testing.T) {
	server := setupServer()
	userReq, rec := newUser(t)

	// Create user
	server.handlePostUser(rec, userReq)
	require.Equal(t, http.StatusOK, rec.Code)
	require.NotNil(t, rec.Body)
	bodyBytes, err := io.ReadAll(rec.Body)
	require.NoError(t, err)

	// Read response with user id
	type UserResp struct {
		Id     int64  `json:"id"`
		Status string `json:"status"`
	}
	user := UserResp{}
	err = json.Unmarshal(bodyBytes, &user)
	require.NoError(t, err)

	// Prepare and send route request
	routeReq := RouteRequest{
		Start: primitives.RoutePoint{
			Latitude:   123.0,
			Longtitude: 456.0,
		},
		Finish: primitives.RoutePoint{
			Latitude:   789.0,
			Longtitude: 654.0,
		},
	}
	data, err := json.Marshal(routeReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/user/id/routes", bytes.NewReader(data))
	req.SetPathValue("id", strconv.FormatInt(user.Id, 10))
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	server.handlePostRoute(recorder, req)

	// check that route is created
	require.Equal(t, http.StatusOK, recorder.Code)
	require.NotNil(t, recorder.Body)

	bodyBytes, err = io.ReadAll(recorder.Body)
	require.NoError(t, err)
	require.Equal(t, string(bodyBytes), `{"status":"ok"}`)
	require.Equal(t, len(server.Users[user.Id].Routes), 1)
}
