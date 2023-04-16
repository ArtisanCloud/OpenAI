package openaitype

// PointInt 将int转换为*int
func PointInt(i int) *int {
	return &i
}

// PointInt32 将int32转换为*int32
func PointInt32(i int32) *int32 {
	return &i
}

// PointInt64 将int64转换为*int64
func PointInt64(i int64) *int64 {
	return &i
}

// PointBool 将bool转换为*bool
func PointBool(b bool) *bool {
	return &b
}

// PointFloat32 将float32转换为*float32
func PointFloat32(f float32) *float32 {
	return &f
}

// PointFloat64 将float64转换为*float64
func PointFloat64(f float64) *float64 {
	return &f
}

// PointString 将string转换为*string
func PointString(s string) *string {
	return &s
}
