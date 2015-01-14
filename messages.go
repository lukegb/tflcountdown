package tflcountdown

import (
	"time"
)

type Direction uint

const (
	DIRECTION_ONE Direction = iota
	DIRECTION_TWO
)

type TflMessageType uint

const (
	MESSAGETYPE_STOP TflMessageType = iota
	MESSAGETYPE_PREDICTION
	MESSAGETYPE_FLEXIBLE_MESSAGE
	MESSAGETYPE_BASEVERSION
	MESSAGETYPE_URA_VERSION
)


type Message interface {
	TflMessageType() TflMessageType
	Decode(*TflArray, FieldMap) (Message, error)
}

var registeredTflMessageTypes map[TflMessageType]Message

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

type BaseStopData struct {
	StopPointName      *string
	StopID             *string
	StopCode1          *string
	StopCode2          *string
	StopPointType      *string
	Towards            *string
	Bearing            *uint
	StopPointIndicator *string
	StopPointState     *uint
	Latitude           *float64
	Longitude          *float64
}

func (b *BaseStopData) Decode(inp *TflArray, fields FieldMap) error {
	if fields.Contains("StopPointName") {
		b.StopPointName = inp.AsStr()
	}

	if fields.Contains("StopID") {
		b.StopID = inp.AsStr()
	}

	if fields.Contains("StopCode1") {
		b.StopCode1 = inp.AsStr()
	}

	if fields.Contains("StopCode2") {
		b.StopCode2 = inp.AsStr()
	}

	if fields.Contains("StopPointType") {
		b.StopPointType = inp.AsStr()
	}

	if fields.Contains("Towards") {
		b.Towards = inp.AsStr()
	}

	if fields.Contains("Bearing") {
		b.Bearing = inp.AsUint()
	}

	if fields.Contains("StopPointIndicator") {
		b.StopPointIndicator = inp.AsStr()
	}

	if fields.Contains("StopPointState") {
		b.StopPointState = inp.AsUint()
	}

	if fields.Contains("Latitude") {
		b.Latitude = inp.AsFloat64()
	}

	if fields.Contains("Longitude") {
		b.Longitude = inp.AsFloat64()
	}

	return nil
}

type StopData struct {
	BaseStopData
}

func (StopData) Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.AsTflMessageType() != MESSAGETYPE_STOP {
		return nil, ERROR_INVALID
	}

	msg := StopData{}

	err := msg.BaseStopData.Decode(inp, fields)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

func (StopData) TflMessageType() TflMessageType {
	return MESSAGETYPE_STOP
}

type PredictionData struct {
	BaseStopData

	VisitNumber        *int
	LineID             *string
	LineName           *string
	DirectionID        *int
	DestinationText    *string
	DestinationName    *string
	VehicleID          *string
	TripID             *int
	RegistrationNumber *string
	EstimatedTime      *time.Time
	ExpireTime         *time.Time
}

func (PredictionData) Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.AsTflMessageType() != MESSAGETYPE_PREDICTION {
		return nil, ERROR_INVALID
	}

	msg := PredictionData{}
	err := msg.BaseStopData.Decode(inp, fields)

	if fields.Contains("VisitNumber") {
		msg.VisitNumber = inp.AsInt()
	}

	if fields.Contains("LineID") {
		msg.LineID = inp.AsStr()
	}

	if fields.Contains("LineName") {
		msg.LineName = inp.AsStr()
	}

	if fields.Contains("DirectionID") {
		msg.DirectionID = inp.AsInt()
	}

	if fields.Contains("DestinationText") {
		msg.DestinationText = inp.AsStr()
	}

	if fields.Contains("DestinationName") {
		msg.DestinationName = inp.AsStr()
	}

	if fields.Contains("VehicleID") {
		msg.VehicleID = inp.AsStr()
	}

	if fields.Contains("TripID") {
		msg.TripID = inp.AsInt()
	}

	if fields.Contains("RegistrationNumber") {
		msg.RegistrationNumber = inp.AsStr()
	}

	if fields.Contains("EstimatedTime") {
		msg.EstimatedTime = inp.AsTime()
	}

	if fields.Contains("ExpireTime") {
		msg.ExpireTime = inp.AsTime()
	}

	return msg, err
}

func (PredictionData) TflMessageType() TflMessageType {
	return MESSAGETYPE_PREDICTION
}

type FlexibleMessage struct {
	BaseStopData

	MessageUUID     *string
	MessageType     *uint
	MessagePriority *uint
	MessageText     *string
	StartTime       *time.Time
	ExpireTime      *time.Time
}

func (FlexibleMessage) Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.AsTflMessageType() != MESSAGETYPE_FLEXIBLE_MESSAGE {
		return nil, ERROR_INVALID
	}

	msg := FlexibleMessage{}
	err := msg.BaseStopData.Decode(inp, fields)
	if err != nil {
		return msg, err
	}

	if fields.Contains("MessageUUID") {
		msg.MessageUUID = inp.AsStr()
	}

	if fields.Contains("MessageType") {
		msg.MessageType = inp.AsUint()
	}

	if fields.Contains("MessagePriority") {
		msg.MessagePriority = inp.AsUint()
	}

	if fields.Contains("MessageText") {
		msg.MessageText = inp.AsStr()
	}

	if fields.Contains("StartTime") {
		msg.StartTime = inp.AsTime()
	}

	if fields.Contains("ExpireTime") {
		msg.ExpireTime = inp.AsTime()
	}

	return msg, nil
}

func (FlexibleMessage) TflMessageType() TflMessageType {
	return MESSAGETYPE_FLEXIBLE_MESSAGE
}

type BaseVersion struct {
	Version *string
}

func (BaseVersion) TflMessageType() TflMessageType {
	return MESSAGETYPE_BASEVERSION
}

func (BaseVersion) Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.AsTflMessageType() != MESSAGETYPE_BASEVERSION {
		return nil, ERROR_INVALID
	}

	msg := BaseVersion{}
	if fields.Contains("BaseVersion") {
		msg.Version = inp.AsStr()
	}
	return msg, nil
}

type UraVersion struct {
	Version   *string
	TimeStamp *time.Time
}

func (UraVersion) Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.AsTflMessageType() != MESSAGETYPE_URA_VERSION {
		return nil, ERROR_INVALID
	}

	msg := UraVersion{}
	msg.Version = inp.AsStr()
	msg.TimeStamp = inp.AsTime()
	return msg, nil
}

func (UraVersion) TflMessageType() TflMessageType {
	return MESSAGETYPE_URA_VERSION
}

func register(message TflMessageType, msg Message) {
	registeredTflMessageTypes[message] = msg
}

func init() {
	registeredTflMessageTypes = make(map[TflMessageType]Message)
	register(MESSAGETYPE_STOP, StopData{})
	register(MESSAGETYPE_PREDICTION, PredictionData{})
	register(MESSAGETYPE_FLEXIBLE_MESSAGE, FlexibleMessage{})
	register(MESSAGETYPE_BASEVERSION, BaseVersion{})
	register(MESSAGETYPE_URA_VERSION, UraVersion{})
}
