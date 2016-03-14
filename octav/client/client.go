// DO NOT EDIT. Automatically generated by hsup
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-urlenc"
)

var _ = bytes.MinRead
var _ = json.Decoder{}

type Client struct {
	Client   *http.Client
	Endpoint string
}

func New(s string) *Client {
	return &Client{
		Client:   &http.Client{},
		Endpoint: s,
	}
}

func (c *Client) CreateConference(in *model.CreateConferenceRequest) (ret *model.Conference, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Conference
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateRoom(in *model.CreateRoomRequest) (ret *model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateSession(in *model.CreateSessionRequest) (ret *model.Session, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Session
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateUser(in *model.CreateUserRequest) (ret *model.User, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.User
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateVenue(in *model.CreateVenueRequest) (ret *model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) DeleteConference(in *model.DeleteConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteRoom(in *model.DeleteRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteSession(in *model.DeleteSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteUser(in *model.DeleteUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteVenue(in *model.DeleteVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) ListRooms(in *model.ListRoomRequest) (ret []model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListRooms").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload []model.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListSessionsByConference(in *model.ListSessionsByConferenceRequest) (ret interface{}, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListSessionsByConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/schedule/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload interface{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListVenues(in *model.ListVenueRequest) (ret []model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListVenues").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload []model.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) LookupConference(in *model.LookupConferenceRequest) (ret *model.Conference, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Conference
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupRoom(in *model.LookupRoomRequest) (ret *model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupSession(in *model.LookupSessionRequest) (ret *model.Session, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Session
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupUser(in *model.LookupUserRequest) (ret *model.User, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.User
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupVenue(in *model.LookupVenueRequest) (ret *model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload model.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) UpdateConference(in *model.UpdateConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateRoom(in *model.UpdateRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateSession(in *model.UpdateSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateUser(in *model.UpdateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateVenue(in *model.UpdateVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}
