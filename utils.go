package tflcountdown

import (
	"strings"
	"time"
)

type FieldMap map[string]bool

func NewFieldMap() FieldMap {
	return make(FieldMap)
}

func NewDefaultFieldMap() FieldMap {
	fm := NewFieldMap()
	fm.Add("StopPointName")
	fm.Add("LineName")
	fm.Add("EstimatedTime")
	return fm
}

func (f FieldMap) Add(x string) {
	x = strings.ToLower(x)
	f[x] = true
}
func (f FieldMap) Remove(x string) {
	x = strings.ToLower(x)
	f[x] = false
}
func (f FieldMap) Contains(name string) bool {
	name = strings.ToLower(name)
	x, y := f[name]
	return x && y
}
func (f FieldMap) AddAll(x []string) {
	for _, v := range x {
		f.Add(v)
	}
}

func (f FieldMap) Keys() []string {
	keys := make([]string, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}
	return keys
}

type TflArray struct {
	currElem int
	arr      []interface{}
}

func NewTflArray(iarr []interface{}) *TflArray {
	return &TflArray{0, iarr}
}

func (t *TflArray) AsTflMessageType() TflMessageType {
	v := TflMessageType(t.arr[t.currElem].(float64))
	t.currElem++
	return v
}
func (t *TflArray) AsStr() *string {
	v := t.arr[t.currElem].(string)
	t.currElem++
	return &v
}
func (t *TflArray) AsInt() *int {
	v := t.arr[t.currElem].(int)
	t.currElem++
	return &v
}
func (t *TflArray) AsUint() *uint {
	v := t.arr[t.currElem].(uint)
	t.currElem++
	return &v
}
func (t *TflArray) AsFloat64() *float64 {
	v := t.arr[t.currElem].(float64)
	t.currElem++
	return &v
}
func (t *TflArray) AsTime() *time.Time {
	v := int64(t.arr[t.currElem].(float64) / 1000)
	t.currElem++
	time := time.Unix(v, 0)
	return &time
}

func (t *TflArray) Len() int {
	return len(t.arr)
}
func (t *TflArray) Rewind() {
	t.currElem = 0
}
