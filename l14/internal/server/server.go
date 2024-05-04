package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"hw15/internal/routes"
	"hw15/internal/users"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	Users map[int64]*users.User
}

func (s *Server) GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	// presentation layer
	router.HandleFunc("GET /user/{id}/routes", s.handleGetRoutes)           // get user's routes
	router.HandleFunc("POST /user/{id}/routes", s.handlePostRoute)          // create route
	router.HandleFunc("POST /user/{id}/routes/next", s.handlePostRouteNext) // start or continue the route route
	router.HandleFunc("POST /user", s.handlePostUser)                       // register a user

	return router
}

func (s *Server) handleGetRoutes(w http.ResponseWriter, r *http.Request) {
	user, err := s.readUser(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	routes, err := json.Marshal(user.Routes)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "cannot marshal routes")
	}
	w.Write(routes)
}

func (s *Server) handlePostUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := s.readBody(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	userReq := UserCreationRequest{}
	err = json.Unmarshal(bodyBytes, &userReq)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Body unmarshalling error: %v", err))
		return
	}
	user := users.NewUser(userReq.Nickname)
	if s.Users == nil {
		s.Users = make(map[int64]*users.User)
	}
	s.Users[user.Id] = &user
	w.Write([]byte(fmt.Sprintf(`{"status":"ok", "id": %d}`, user.Id)))
}

func (s *Server) handlePostRoute(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := s.readBody(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.readUser(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if user.Routes != nil && len(user.Routes) > 0 {
		sendError(w, http.StatusBadRequest, "user already has some routes")
		return
	}
	rr := RouteRequest{}
	err = json.Unmarshal(bodyBytes, &rr)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Body unmarshalling error: %v", err))
		return
	}
	route := routes.NewRoute(rr.Start, rr.Finish)
	user.Routes = append(user.Routes, route)
	w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) handlePostRouteNext(w http.ResponseWriter, r *http.Request) {
	user, err := s.readUser(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(user.Routes) == 0 {
		sendError(w, http.StatusBadRequest, "User has no routes")
		return
	}
	route := user.Routes[len(user.Routes)-1]
	err = route.Next()
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if route.IsFinished() {
		w.Write([]byte(`{"status":"ok", "msg": "You are at the destination"}`))
	} else {
		bytes, err := json.Marshal(route.Points[route.NextDestination])
		if err != nil {
			sendError(w, http.StatusInternalServerError, "cannot marshal route")
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"status":"ok", "next": %s}`, string(bytes))))
	}
}

func (s *Server) readBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, errors.New("no body in request")
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return bodyBytes, nil
}

func (s *Server) readUser(r *http.Request) (*users.User, error) {
	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse user id: %w", err)
	}
	user, ok := s.Users[userId]
	if !ok {
		return nil, errors.New("cannot find user")
	}
	return user, nil
}

func sendError(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"status":"error","message":%s"}`, text)))
}
