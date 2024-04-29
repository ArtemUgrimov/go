package server

import (
	"bytes"
	"encoding/json"
	"hw15/internal/users"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerPostUser(t *testing.T) {
	server := Server{}
	mux := server.GetRouter()
	go http.ListenAndServe(":8000", mux)

	userReq := users.UserCreationRequest{
		Nickname: "Artem",
	}
	data, err := json.Marshal(userReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/user", bytes.NewReader(data))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	server.handlePostUser(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	require.NotNil(t, rr.Body)

	bodyBytes, err := io.ReadAll(rr.Body)
	require.NoError(t, err)
	require.Equal(t, string(bodyBytes), `{"status":"ok"}`)
}
