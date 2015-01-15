package tflcountdown

import (
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	StopAlso           *bool
	ReturnList         *FieldMap
	Circle             []string
	StopPointName      []string
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
	MessageType     []uint
	MessagePriority    []uint
}

func (r *Request) Encode() url.Values {
	v := url.Values{}

	if r.StopAlso != nil {
		if *r.StopAlso {
			v.Add("StopAlso", "true")
		} else {
			v.Add("StopAlso", "false")
		}
	}

	maybeAddValue(v, "ReturnList", strings.Join(r.ReturnList.Keys(), ","))
	maybeAddValue(v, "Circle", strings.Join(r.Circle, ","))
	maybeAddValue(v, "StopPointName", strings.Join(r.StopPointName, ","))
	maybeAddValue(v, "StopID", strings.Join(r.StopID, ","))
	maybeAddValue(v, "StopCode1", strings.Join(r.StopCode1, ","))
	maybeAddValue(v, "StopCode2", strings.Join(r.StopCode2, ","))
	maybeAddValue(v, "StopPointType", strings.Join(r.StopPointType, ","))
	maybeAddValue(v, "Towards", strings.Join(r.Towards, ","))

	maybeAddValue(v, "Bearing", strings.Join(encodeIntArray(r.Bearing), ","))
	maybeAddValue(v, "StopPointState", strings.Join(encodeUintArray(r.StopPointState), ","))
	maybeAddValue(v, "VisitNumber", strings.Join(encodeUintArray(r.VisitNumber), ","))

	maybeAddValue(v, "LineID", strings.Join(r.LineID, ","))
	maybeAddValue(v, "LineName", strings.Join(r.LineName, ","))

	if r.DirectionID != nil {
		if *r.DirectionID == 1 {
			v.Add("DirectionID", "1")
		} else {
			v.Add("DirectionID", "2")
		}
	}

	maybeAddValue(v, "DirectionText", strings.Join(r.DirectionText, ","))
	maybeAddValue(v, "DirectionName", strings.Join(r.DirectionName, ","))
	maybeAddValue(v, "VehicleID", strings.Join(r.VehicleID, ","))
	maybeAddValue(v, "TripID", strings.Join(r.TripID, ","))
	maybeAddValue(v, "RegistrationNumber", strings.Join(r.RegistrationNumber, ","))
	maybeAddValue(v, "StopPointIndicator", strings.Join(r.StopPointIndicator, ","))

	maybeAddValue(v, "MessageType", strings.Join(encodeUintArray(r.MessageType), ","))
	maybeAddValue(v, "MessagePriority", strings.Join(encodeUintArray(r.MessagePriority), ","))

	return v
}

func maybeAddValue(d url.Values, n string, v string) {
	if v != "" {
		d.Add(n, v)
	}
}

func encodeIntArray(x []int) []string {
	arr := make([]string, len(x))
	for i, v := range x {
		arr[i] = strconv.FormatInt(int64(v), 10)
	}
	return arr
}
func encodeUintArray(x []uint) []string {
	arr := make([]string, len(x))
	for i, v := range x {
		arr[i] = strconv.FormatUint(uint64(v), 10)
	}
	return arr
}
