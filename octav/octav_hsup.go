package octav

// DO NOT EDIT. Automatically generated by hsup
import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/validator"
	"github.com/gorilla/mux"
	"github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-urlenc"
	"golang.org/x/net/context"
)

var _ = json.Decoder{}
var _ = urlenc.Marshal

type Server struct {
	*mux.Router
}

// NewContext creates a cteonxt.Context object from the request.
// If you are using appengine, for example, you probably want to set this
// function to something that create a context, and then sets
// the appengine context to it so it can be referred to later.
var NewContext func(*http.Request) context.Context = func(r *http.Request) context.Context {
	return context.Background()
}

func Run(l string) error {
	return http.ListenAndServe(l, New())
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.SetupRoutes()
	return s
}

func httpError(w http.ResponseWriter, message string, st int, err error) {
	if pdebug.Enabled {
		if err == nil {
			pdebug.Printf("HTTP Error %s", message)
		} else {
			pdebug.Printf("HTTP Error %s: %s", message, err)
		}
	}
	http.Error(w, http.StatusText(st), st)
}

func getInteger(v url.Values, f string) ([]int64, error) {
	x, ok := v[f]
	if !ok {
		return nil, nil
	}

	ret := make([]int64, len(x))
	for i, e := range x {
		p, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil, err
		}
		ret[i] = p
	}

	return ret, nil
}

func httpAddConferenceDates(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpAddConferenceDates")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.AddConferenceDatesRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPAddConferenceDatesRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doAddConferenceDates(NewContext(r), w, r, payload)
}

func httpCreateConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.CreateConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doCreateConference(NewContext(r), w, r, payload)
}

func httpCreateRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doCreateRoom(NewContext(r), w, r, payload)
}

func httpCreateSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.CreateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doCreateSession(NewContext(r), w, r, payload)
}

func httpCreateUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doCreateUser(NewContext(r), w, r, payload)
}

func httpCreateVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.CreateVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doCreateVenue(NewContext(r), w, r, payload)
}

func httpDeleteConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteConference(NewContext(r), w, r, payload)
}

func httpDeleteConferenceDates(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteConferenceDates")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteConferenceDatesRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteConferenceDatesRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteConferenceDates(NewContext(r), w, r, payload)
}

func httpDeleteRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteRoom(NewContext(r), w, r, payload)
}

func httpDeleteSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteSession(NewContext(r), w, r, payload)
}

func httpDeleteUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteUser(NewContext(r), w, r, payload)
}

func httpDeleteVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.DeleteVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doDeleteVenue(NewContext(r), w, r, payload)
}

func httpListConferences(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListConferences")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.ListConferencesRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListConferencesRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doListConferences(NewContext(r), w, r, payload)
}

func httpListRooms(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListRooms")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.ListRoomRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListRoomsRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doListRooms(NewContext(r), w, r, payload)
}

func httpListSessionsByConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListSessionsByConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.ListSessionsByConferenceRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListSessionsByConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doListSessionsByConference(NewContext(r), w, r, payload)
}

func httpListVenues(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListVenues")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.ListVenueRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListVenuesRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doListVenues(NewContext(r), w, r, payload)
}

func httpLookupConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.LookupConferenceRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doLookupConference(NewContext(r), w, r, payload)
}

func httpLookupRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.LookupRoomRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doLookupRoom(NewContext(r), w, r, payload)
}

func httpLookupSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.LookupSessionRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doLookupSession(NewContext(r), w, r, payload)
}

func httpLookupUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.LookupUserRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doLookupUser(NewContext(r), w, r, payload)
}

func httpLookupVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.LookupVenueRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doLookupVenue(NewContext(r), w, r, payload)
}

func httpUpdateConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.UpdateConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doUpdateConference(NewContext(r), w, r, payload)
}

func httpUpdateRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doUpdateRoom(NewContext(r), w, r, payload)
}

func httpUpdateSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.UpdateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doUpdateSession(NewContext(r), w, r, payload)
}

func httpUpdateUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doUpdateUser(NewContext(r), w, r, payload)
}

func httpUpdateVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload model.UpdateVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid JSON input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input (validation failed)`, http.StatusInternalServerError, err)
		return
	}
	doUpdateVenue(NewContext(r), w, r, payload)
}

func (s *Server) SetupRoutes() {
	r := s.Router
	r.HandleFunc(`/v1/conference/create`, httpCreateConference)
	r.HandleFunc(`/v1/conference/date/add`, httpAddConferenceDates)
	r.HandleFunc(`/v1/conference/date/delete`, httpDeleteConferenceDates)
	r.HandleFunc(`/v1/conference/delete`, httpDeleteConference)
	r.HandleFunc(`/v1/conference/list`, httpListConferences)
	r.HandleFunc(`/v1/conference/lookup`, httpLookupConference)
	r.HandleFunc(`/v1/conference/update`, httpUpdateConference)
	r.HandleFunc(`/v1/room/create`, httpCreateRoom)
	r.HandleFunc(`/v1/room/delete`, httpDeleteRoom)
	r.HandleFunc(`/v1/room/list`, httpListRooms)
	r.HandleFunc(`/v1/room/lookup`, httpLookupRoom)
	r.HandleFunc(`/v1/room/update`, httpUpdateRoom)
	r.HandleFunc(`/v1/schedule/list`, httpListSessionsByConference)
	r.HandleFunc(`/v1/session/create`, httpCreateSession)
	r.HandleFunc(`/v1/session/delete`, httpDeleteSession)
	r.HandleFunc(`/v1/session/lookup`, httpLookupSession)
	r.HandleFunc(`/v1/session/update`, httpUpdateSession)
	r.HandleFunc(`/v1/user/create`, httpCreateUser)
	r.HandleFunc(`/v1/user/delete`, httpDeleteUser)
	r.HandleFunc(`/v1/user/lookup`, httpLookupUser)
	r.HandleFunc(`/v1/user/update`, httpUpdateUser)
	r.HandleFunc(`/v1/venue/create`, httpCreateVenue)
	r.HandleFunc(`/v1/venue/delete`, httpDeleteVenue)
	r.HandleFunc(`/v1/venue/list`, httpListVenues)
	r.HandleFunc(`/v1/venue/lookup`, httpLookupVenue)
	r.HandleFunc(`/v1/venue/update`, httpUpdateVenue)
}