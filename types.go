package hellogo

// TypeConvert 用于将一种类型转成另一种类型
type TypeConvert[IT any, OT any] func(IT) OT

// Interface 将输入的类型转为 interface{}
func Interface[T any](input T) interface{} {
	return input
}

// Pointer 将输入的值转成指针
func Pointer[T any](input T) *T {
	return &input
}

// Value 将输入的指针转成值
func Value[T any](input *T) T {
	return *input
}

// Self 返回自身
func Self[T any](input T) T {
	return input
}

// SliceElemType 转换slice 中的元素类型
func SliceElemType[IT any, OT any](input []IT, f TypeConvert[IT, OT]) []OT {
	ret := make([]OT, 0, len(input))
	for _, i := range input {
		ret = append(ret, f(i))
	}
	return ret
}

// NewSliceConvert 转换 slice 类型的转换器，用于多层slice [][]T 使用
func NewSliceConvert[IT any, OT any](input []IT, f TypeConvert[IT, OT]) TypeConvert[[]IT, []OT] {
	return func(input []IT) []OT {
		return SliceElemType(input, f)
	}
}

// MapElemType 将map 中的元素转换成另一种
func MapElemType[KT comparable, IT any, OT any](input map[KT]IT, f TypeConvert[IT, OT]) map[KT]OT {
	ret := make(map[KT]OT, len(input))
	for k, v := range input {
		ret[k] = f(v)
	}
	return ret
}

// NewMapConvert 转换map 类型的转换器，用于处理多层map[KT1]map[KT2]VT
func NewMapConvert[KT comparable, IT any, OT any](input map[KT]IT, f TypeConvert[IT, OT]) func(map[KT]IT) map[KT]OT {
	return func(input map[KT]IT) map[KT]OT {
		return MapElemType(input, f)
	}
}


// KeyF 用于从输入类型，获取一个可当做键盘值得类型，特殊的TypeConvert
type KeyF[T any, KT comparable] func(input T) KT

// Group 对输入的slice进行分组
func Group[T any, KT comparable](input []T, f KeyF[T, KT]) map[KT][]T {
	ret := make(map[KT][]T, len(input))
	for _, v := range input {
		key := f(v)
		ret[key] = append(ret[key], v)
	}
	return ret
}

// SliceToMap 将slice转换为map
func SliceToMap[KT comparable, T any](input []T, f KeyF[T, KT]) map[KT]T {
	ret := make(map[KT]T, len(input))
	for _, i := range input {
		ret[f(i)] = i
	}
	return ret
}


// SliceToSet 将slice转换为set
func SliceToSet[T any, KT comparable](input []T, f KeyF[T, KT]) map[KT]struct{} {
	ret := make(map[KT]struct{}, len(input))
	for _, i := range input {
		ret[f(i)] = struct{}{}
	}
	return ret
}

// MapToSlice 将map转换为slice
func MapToSlice[KT comparable, T any](input map[KT]T) []T {
	ret := make([]T, 0, len(input))
	for _, v := range input {
		ret = append(ret, v)
	}
	return ret
}
