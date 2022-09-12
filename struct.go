package pvtago

type PublicMessage struct {
	Cause               int64             `json:"Cause"`
	CauseReportLabel    CauseReportLabel  `json:"CauseReportLabel"`
	ChannelMessages     []interface{}     `json:"ChannelMessages"`
	DaysOfWeek          int64             `json:"DaysOfWeek"`
	Effect              int64             `json:"Effect"`
	EffectReportLabel   EffectReportLabel `json:"EffectReportLabel"`
	FromDate            string            `json:"FromDate"`
	FromTime            string            `json:"FromTime"`
	Message             string            `json:"Message"`
	MessageID           int64             `json:"MessageId"`
	MessageTranslations []interface{}     `json:"MessageTranslations"`
	Priority            int64             `json:"Priority"`
	PublicAccess        int64             `json:"PublicAccess"`
	Published           bool              `json:"Published"`
	Routes              []int64           `json:"Routes"`
	Signs               []int64           `json:"Signs"`
	ToDate              string            `json:"ToDate"`
	ToTime              string            `json:"ToTime"`
	URL                 interface{}       `json:"URL"`
	DetourID            interface{}       `json:"Detour_Id"`
}

type CauseReportLabel string

const (
	UnknownCause CauseReportLabel = "UnknownCause"
)

type EffectReportLabel string

const (
	UnknownEffect EffectReportLabel = "UnknownEffect"
)

type RouteDetail struct {
	Color                    string          `json:"Color"`
	Directions               []Direction     `json:"Directions"`
	GoogleDescription        string          `json:"GoogleDescription"`
	Group                    interface{}     `json:"Group"`
	IncludeInGoogle          bool            `json:"IncludeInGoogle"`
	IsHeadway                bool            `json:"IsHeadway"`
	IsHeadwayMonitored       bool            `json:"IsHeadwayMonitored"`
	IsVisible                bool            `json:"IsVisible"`
	IvrDescription           string          `json:"IvrDescription"`
	LongName                 string          `json:"LongName"`
	Messages                 []PublicMessage `json:"Messages"`
	RouteAbbreviation        string          `json:"RouteAbbreviation"`
	RouteID                  int64           `json:"RouteId"`
	RouteRecordID            int64           `json:"RouteRecordId"`
	RouteStops               []RouteStop     `json:"RouteStops"`
	RouteTraceFilename       string          `json:"RouteTraceFilename"`
	RouteTraceHash64         interface{}     `json:"RouteTraceHash64"`
	ShortName                string          `json:"ShortName"`
	SortOrder                int64           `json:"SortOrder"`
	Stops                    []Stop          `json:"Stops"`
	TextColor                string          `json:"TextColor"`
	Vehicles                 []Vehicle       `json:"Vehicles"`
	DetourActiveMessageCount int64           `json:"DetourActiveMessageCount"`
}

type Direction struct {
	Dir                   Dir         `json:"Dir"`
	DirectionDesc         interface{} `json:"DirectionDesc"`
	DirectionIconFileName interface{} `json:"DirectionIconFileName"`
}

type RouteStop struct {
	Direction Dir   `json:"Direction"`
	RouteID   int64 `json:"RouteId"`
	SortOrder int64 `json:"SortOrder"`
	StopID    int64 `json:"StopId"`
}

type Stop struct {
	Description  string  `json:"Description"`
	IsTimePoint  bool    `json:"IsTimePoint"`
	Latitude     float64 `json:"Latitude"`
	Longitude    float64 `json:"Longitude"`
	Name         string  `json:"Name"`
	StopID       int64   `json:"StopId"`
	StopRecordID int64   `json:"StopRecordId"`
}

type Vehicle struct {
	BlockFareboxID             int64       `json:"BlockFareboxId"`
	CommStatus                 string      `json:"CommStatus"`
	Destination                string      `json:"Destination"`
	Deviation                  int64       `json:"Deviation"`
	Direction                  Dir         `json:"Direction"`
	DirectionLong              string      `json:"DirectionLong"`
	DisplayStatus              string      `json:"DisplayStatus"`
	StopID                     int64       `json:"StopId"`
	CurrentStatus              interface{} `json:"CurrentStatus"`
	DriverName                 interface{} `json:"DriverName"`
	GPSStatus                  int64       `json:"GPSStatus"`
	Heading                    int64       `json:"Heading"`
	LastStop                   string      `json:"LastStop"`
	LastUpdated                string      `json:"LastUpdated"`
	Latitude                   float64     `json:"Latitude"`
	Longitude                  float64     `json:"Longitude"`
	Name                       string      `json:"Name"`
	OccupancyStatus            int64       `json:"OccupancyStatus"`
	OnBoard                    int64       `json:"OnBoard"`
	OpStatus                   string      `json:"OpStatus"`
	RouteID                    int64       `json:"RouteId"`
	RunID                      int64       `json:"RunId"`
	Speed                      interface{} `json:"Speed"`
	TripID                     int64       `json:"TripId"`
	VehicleID                  int64       `json:"VehicleId"`
	SeatingCapacity            int64       `json:"SeatingCapacity"`
	TotalCapacity              int64       `json:"TotalCapacity"`
	PropertyName               string      `json:"PropertyName"`
	OccupancyStatusReportLabel string      `json:"OccupancyStatusReportLabel"`
}

type Dir string

const (
	N Dir = "N"
	S Dir = "S"
)
