package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicTest(t *testing.T) {
	// setup
	data := []byte{}
	req, err := http.NewRequest("POST", "/items", bytes.NewReader(data))
	require.NoError(t, err, "Error creating a request")

	// execution
	recorder := httptest.NewRecorder()
	// call the handler function
	// handler(recorder, req)

	// assertion
	require.Equal(t, http.StatusOK, recorder.Code)

	// assert.Equal(t, add(1, 1), 2, "They should be equal")

	// if add(5, 5) != 10 {
	// 	t.Error("Somthing went wrong with function add")
	// }

	// cases := []struct {
	// 	name string
	// 	a    int
	// 	b    int
	// 	sum  int
	// }{
	// 	{
	// 		name: "Adding two small numbers",
	// 		a:    1,
	// 		b:    2,
	// 		sum:  3,
	// 	},
	// 	{
	// 		name: "Adding two big numbers",
	// 		a:    5,
	// 		b:    5,
	// 		sum:  10,
	// 	},
	// 	{
	// 		name: "Adding two negative numbers",
	// 		a:    -100,
	// 		b:    -200,
	// 		sum:  -300,
	// 	},
	// }

	// for _, c := range cases {
	// 	t.Run(c.name, func(t *testing.T) {
	// 		if add(c.a, c.b) != c.sum {
	// 			t.Error("Sum does not match")
	// 		}
	// 	})
	// }
}
