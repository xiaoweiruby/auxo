package cast

import (
	"reflect"
	"strings"

	"github.com/cuigh/auxo/ext/reflects"
)

func TryToValue(i interface{}, t reflect.Type) (v reflect.Value, err error) {
	var value interface{}
	switch t {
	case reflects.TypeBool:
		value, err = TryToBool(i)
	case reflects.TypeString:
		value = ToString(i)
	case reflects.TypeInt:
		value, err = TryToInt(i)
	case reflects.TypeInt8:
		value, err = TryToInt8(i)
	case reflects.TypeInt16:
		value, err = TryToInt16(i)
	case reflects.TypeInt32:
		value, err = TryToInt32(i)
	case reflects.TypeInt64:
		value, err = TryToInt64(i)
	case reflects.TypeUint:
		value, err = TryToUint(i)
	case reflects.TypeUint8:
		value, err = TryToUint8(i)
	case reflects.TypeUint16:
		value, err = TryToUint16(i)
	case reflects.TypeUint32:
		value, err = TryToUint32(i)
	case reflects.TypeUint64:
		value, err = TryToUint64(i)
	case reflects.TypeFloat32:
		value, err = TryToFloat32(i)
	case reflects.TypeFloat64:
		value, err = TryToFloat64(i)
	case reflects.TypeDuration:
		value, err = TryToDuration(i)
	case reflects.TypeTime:
		value, err = TryToTime(i)
	default:
		if t.Kind() == reflect.Slice {
			return TryToSliceValue(i, t.Elem())
		}
		err = castError(i, t.String())
	}
	if err == nil {
		v = reflect.ValueOf(value)
	}
	return
}

// TryToSliceValue cast interface value to a slice.
// Argument t is element type of slice.
func TryToSliceValue(i interface{}, t reflect.Type) (slice reflect.Value, err error) {
	if s, ok := i.(string); ok {
		i = strings.Split(s, ",")
	}

	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		err = castError(i, "[]"+t.String())
		return
	}

	if t == v.Type().Elem() {
		return v, nil
	}

	length := v.Len()
	slice = reflect.MakeSlice(reflect.SliceOf(t), length, length)
	for k := 0; k < length; k++ {
		var value reflect.Value
		value, err = TryToValue(v.Index(k).Interface(), t)
		if err != nil {
			return
		}
		slice.Index(k).Set(value)
	}
	return slice, nil
}

func TryToSlice(i interface{}, t reflect.Type) (interface{}, error) {
	v, err := TryToSliceValue(i, t)
	if err != nil {
		return nil, err
	}
	return v.Interface(), nil
}