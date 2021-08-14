package util

const (
	New = ud(0)
)

type ud int

func (u ud) String(s string) string {
	var data = s
	return data
}

func (u ud) Int8(d int8) int8 {
	var data = d
	return data
}

func (u ud) Int16(d int16) int16 {
	var data = d
	return data
}

func (u ud) Int32(d int32) int32 {
	var data = d
	return data
}

func (u ud) Int64(d int64) int64 {
	var data = d
	return data
}

func (u ud) UInt8(d uint8) uint8 {
	var data = d
	return data
}

func (u ud) UInt16(d uint16) uint16 {
	var data = d
	return data
}

func (u ud) UInt32(d uint32) uint32 {
	var data = d
	return data
}

func (u ud) UInt64(d uint64) uint64 {
	var data = d
	return data
}

func (u ud) Float32(f float32) float32 {
	var data = f
	return data
}

func (u ud) Float64(f float64) float64 {
	var data = f
	return data
}

func (u ud) StringPtr(s string) *string {
	var data = s
	return &data
}

func (u ud) Int8Ptr(d int8) *int8 {
	var data = d
	return &data
}

func (u ud) Int16Ptr(d int16) *int16 {
	var data = d
	return &data
}

func (u ud) Int32Ptr(d int32) *int32 {
	var data = d
	return &data
}

func (u ud) Int64Ptr(d int64) *int64 {
	var data = d
	return &data
}

func (u ud) UInt8Ptr(d uint8) *uint8 {
	var data = d
	return &data
}

func (u ud) UInt16Ptr(d uint16) *uint16 {
	var data = d
	return &data
}

func (u ud) UInt32Ptr(d uint32) *uint32 {
	var data = d
	return &data
}

func (u ud) UInt64Ptr(d uint64) *uint64 {
	var data = d
	return &data
}

func (u ud) Float32Ptr(f float32) *float32 {
	var data = f
	return &data
}

func (u ud) Float64Ptr(f float64) *float64 {
	var data = f
	return &data
}
