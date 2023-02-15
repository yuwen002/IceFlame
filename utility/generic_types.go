package utility

// IfInt
// @Description: int合集
type IfInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// IfUint
// @Description: uint合集
type IfUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}

// IfFloat
// @Description:float合集
type IfFloat interface {
	~float32 | ~float64
}
