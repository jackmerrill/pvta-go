package pvtago

import "fmt"

type Client struct {
	http *NetworkClient
}

func NewClient() *Client {
	return &Client{
		http: &NetworkClient{
			baseURL: "https://bustracker.pvta.com/InfoPoint/rest",
		},
	}
}

// Gets all public messages, including those that have expired.
func (c *Client) GetPublicMessages() ([]PublicMessage, error) {
	var messages []PublicMessage
	err := c.http.Get("/PublicMessages/GetAllMessages", &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Gets a specific public message by ID.
func (c *Client) GetPublicMessage(id int64) (*PublicMessage, error) {
	var message PublicMessage
	err := c.http.Get(fmt.Sprintf("/PublicMessages/Get/%d", id), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// Gets all active public messages.
func (c *Client) GetCurrentMessages() ([]PublicMessage, error) {
	var messages []PublicMessage
	err := c.http.Get("/PublicMessages/GetCurrentMessages", &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Gets all routes.
func (c *Client) GetRoutes() ([]RouteDetail, error) {
	var routes []RouteDetail
	err := c.http.Get("/RouteDetails/GetAllRouteDetails", &routes)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

// Get a route by ID.
func (c *Client) GetRoute(id int64) (*RouteDetail, error) {
	var route RouteDetail
	err := c.http.Get(fmt.Sprintf("/RouteDetails/Get/%d", id), &route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

// Gets all visible routes.
func (c *Client) GetVisibleRoutes() ([]RouteDetail, error) {
	var routes []RouteDetail
	err := c.http.Get("/Routes/GetVisibleRoutes", &routes)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

// Gets all stops.
func (c *Client) GetStops() ([]Stop, error) {
	var stops []Stop
	err := c.http.Get("/Stops/GetAllStops", &stops)
	if err != nil {
		return nil, err
	}

	return stops, nil
}

// Gets a stop by ID.
func (c *Client) GetStop(id int64) (*Stop, error) {
	var stop Stop
	err := c.http.Get(fmt.Sprintf("/Stops/Get/%d", id), &stop)
	if err != nil {
		return nil, err
	}

	return &stop, nil
}

// Searches for a stop by a selected query.
func (c *Client) SearchStops(query string) ([]Stop, error) {
	var stops []Stop
	err := c.http.Get(fmt.Sprintf("/Stops/Search?query=%s", query), &stops)
	if err != nil {
		return nil, err
	}

	return stops, nil
}

// Gets all vehicles.
func (c *Client) GetVehicles() ([]Vehicle, error) {
	var vehicles []Vehicle
	err := c.http.Get("/Vehicles/GetAllVehicles", &vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

// Gets a vehicle by ID.
func (c *Client) GetVehicle(id int64) (*Vehicle, error) {
	var vehicle Vehicle
	err := c.http.Get(fmt.Sprintf("/Vehicles/Get/%d", id), &vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

// Gets all vehicles for a route.
func (c *Client) GetVehiclesForRoute(id int64) ([]Vehicle, error) {
	var vehicles []Vehicle
	err := c.http.Get(fmt.Sprintf("/Vehicles/GetAllVehiclesForRoute?routeID=%d", id), &vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

// Gets all vehicles for routes.
func (c *Client) GetVehiclesForRoutes(ids []int64) ([]Vehicle, error) {
	var vehicles []Vehicle
	queryparams := ""
	for i, id := range ids {
		if i == 0 {
			queryparams += fmt.Sprintf("?routeID=%d", id)
		} else {
			queryparams += fmt.Sprintf("&routeID=%d", id)
		}
	}
	err := c.http.Get(fmt.Sprintf("/Vehicles/GetAllVehiclesForRoutes%s", queryparams), &vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
