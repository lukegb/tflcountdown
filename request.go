package tflcountdown

type Request struct {
	StopAlso           *bool
	ReturnList         []string
	Circle             []string
	StropPointName     []string
	StopID             []string
	StopCode1          []string
	StopCode2          []string
	StopPointType      []string
	Towards            []string
	Bearing            []int
	StopPointState     []uint
	VisitNumber        []uint
	LineID             []string
	LineName           []string
	DirectionID        *Direction
	DirectionText      []string
	DirectionName      []string
	VehicleID          []string
	TripID             []string
	RegistrationNumber []string
	StopPointIndicator []string
	TflMessageType     []uint
	MessagePriority    []uint
}
