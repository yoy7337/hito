package utils

import (
	"time"

	"github.com/jinzhu/copier"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type BaseType interface {
	Integer | Float | ~string | ~bool
}

type CompareType interface {
	Integer | Float
}

type EmptyStruct struct{}
type IndexSearchType map[string]EmptyStruct

// Copy structure from src to dest
// If src is empty, dest will be empty too
func Copy(dest any, src any) error {
	return copier.CopyWithOption(dest, src, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

// Primitive type to pointer
func ToPoint[T Integer | Float | ~string | time.Time | ~bool | map[string]string | map[string]any](s T) *T {
	// 當 s 為常數的時候，直接轉換為指標會有問題，所以需要先將常數轉為變數
	var pointer T
	Copy(&pointer, s)
	return &pointer
}
