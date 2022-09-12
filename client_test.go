package pvtago_test

import (
	"math/rand"
	"testing"
	"time"

	pvtago "github.com/jackmerrill/pvta-go"
)

func TestGetPublicMessages(t *testing.T) {
	client := pvtago.NewClient()
	messages, err := client.GetPublicMessages()
	if err != nil {
		t.Error(err)
	}
	if len(messages) == 0 {
		t.Error("Expected at least one message")
	}

	for _, message := range messages {
		if message.Message == "" {
			t.Error("Expected message to have a message")
		}
	}
}

func TestGetPublicMessage(t *testing.T) {
	client := pvtago.NewClient()
	messages, err := client.GetPublicMessages()
	if err != nil {
		t.Error(err)
	}
	if len(messages) == 0 {
		t.Error("Expected at least one message")
	}

	rand.Seed(time.Now().Unix())

	message, err := client.GetPublicMessage(messages[rand.Intn(len(messages))].MessageID)
	if err != nil {
		t.Error(err)
	}
	if message.Message == "" {
		t.Error("Expected message to have a message")
	}
}

func TestGetRoutes(t *testing.T) {
	client := pvtago.NewClient()
	routes, err := client.GetRoutes()
	if err != nil {
		t.Error(err)
	}
	if len(routes) == 0 {
		t.Error("Expected at least one route")
	}

	for _, route := range routes {
		if route.RouteID == 0 {
			t.Error("Expected route to have an ID")
		}
	}
}

func TestGetRoute(t *testing.T) {
	client := pvtago.NewClient()
	routes, err := client.GetRoutes()
	if err != nil {
		t.Error(err)
	}
	if len(routes) == 0 {
		t.Error("Expected at least one route")
	}

	rand.Seed(time.Now().Unix())

	route, err := client.GetRoute(routes[rand.Intn(len(routes))].RouteID)
	if err != nil {
		t.Error(err)
	}
	if route.RouteID == 0 {
		t.Error("Expected route to have an ID")
	}
}

func TestGetVisibleRoutes(t *testing.T) {
	client := pvtago.NewClient()
	routes, err := client.GetVisibleRoutes()
	if err != nil {
		t.Error(err)
	}
	if len(routes) == 0 {
		t.Error("Expected at least one route")
	}

	for _, route := range routes {
		if route.RouteID == 0 {
			t.Error("Expected route to have an ID")
		}
	}
}

func TestGetStops(t *testing.T) {
	client := pvtago.NewClient()
	stops, err := client.GetStops()
	if err != nil {
		t.Error(err)
	}
	if len(stops) == 0 {
		t.Error("Expected at least one stop")
	}

	for _, stop := range stops {
		if stop.StopID == 0 {
			t.Error("Expected stop to have an ID")
		}
	}
}

func TestGetStop(t *testing.T) {
	client := pvtago.NewClient()
	stops, err := client.GetStops()
	if err != nil {
		t.Error(err)
	}
	if len(stops) == 0 {
		t.Error("Expected at least one stop")
	}

	rand.Seed(time.Now().Unix())

	stop, err := client.GetStop(stops[rand.Intn(len(stops))].StopID)
	if err != nil {
		t.Error(err)
	}
	if stop.StopID == 0 {
		t.Error("Expected stop to have an ID")
	}
}

func TestSearchStop(t *testing.T) {
	client := pvtago.NewClient()
	stops, err := client.SearchStops("Hampshire")
	if err != nil {
		t.Error(err)
	}
	if len(stops) == 0 {
		t.Error("Expected at least one stop")
	}

	for _, stop := range stops {
		if stop.StopID == 0 {
			t.Error("Expected stop to have an ID")
		}
	}
}

func TestGetVehicles(t *testing.T) {
	client := pvtago.NewClient()
	vehicles, err := client.GetVehicles()
	if err != nil {
		t.Error(err)
	}
	if len(vehicles) == 0 {
		t.Error("Expected at least one vehicle")
	}

	for _, vehicle := range vehicles {
		if vehicle.VehicleID == 0 {
			t.Error("Expected vehicle to have an ID")
		}
	}
}

func TestGetVehicle(t *testing.T) {
	client := pvtago.NewClient()
	vehicles, err := client.GetVehicles()
	if err != nil {
		t.Error(err)
	}
	if len(vehicles) == 0 {
		t.Error("Expected at least one vehicle")
	}

	rand.Seed(time.Now().Unix())

	vehicle, err := client.GetVehicle(vehicles[rand.Intn(len(vehicles))].VehicleID)
	if err != nil {
		t.Error(err)
	}
	if vehicle.VehicleID == 0 {
		t.Error("Expected vehicle to have an ID")
	}
}

func TestGetVehiclesForRoute(t *testing.T) {
	client := pvtago.NewClient()
	routes, err := client.GetRoutes()
	if err != nil {
		t.Error(err)
	}
	if len(routes) == 0 {
		t.Error("Expected at least one route")
	}

	rand.Seed(time.Now().Unix())

	route, err := client.GetRoute(routes[rand.Intn(len(routes))].RouteID)
	if err != nil {
		t.Error(err)
	}
	if route.RouteID == 0 {
		t.Error("Expected route to have an ID")
	}

	vehicles, err := client.GetVehiclesForRoute(route.RouteID)
	if err != nil {
		t.Error(err)
	}
	if len(vehicles) == 0 {
		// a route may not have any vehicles
		return
	}
}

func TestGetVehiclesForRoutes(t *testing.T) {
	client := pvtago.NewClient()
	routes, err := client.GetRoutes()
	if err != nil {
		t.Error(err)
	}
	if len(routes) == 0 {
		t.Error("Expected at least one route")
	}

	rand.Seed(time.Now().Unix())

	route1, err := client.GetRoute(routes[rand.Intn(len(routes))].RouteID)
	if err != nil {
		t.Error(err)
	}
	if route1.RouteID == 0 {
		t.Error("Expected route to have an ID")
	}

	route2, err := client.GetRoute(routes[rand.Intn(len(routes))].RouteID)
	if err != nil {
		t.Error(err)
	}
	if route2.RouteID == 0 {
		t.Error("Expected route to have an ID")
	}

	vehicles, err := client.GetVehiclesForRoutes([]int64{route1.RouteID, route2.RouteID})
	if err != nil {
		t.Error(err)
	}
	if len(vehicles) == 0 {
		// a route may not have any vehicles
		return
	}
}
