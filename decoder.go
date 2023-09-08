package kfaker

import (
	"reflect"

	"github.com/vingarcia/structscanner"
)

func Fake(obj any, customValues map[string]any) error {
	return structscanner.Decode(obj, newDecoder(customValues))
}

// decoder can be used to fill a struct with the values of from url.Values.
type decoder struct {
	customValues map[string]any
}

func newDecoder(customValues map[string]any) decoder {
	return decoder{
		customValues: customValues,
	}
}

// DecodeField implements the TagDecoder interface
func (e decoder) DecodeField(info structscanner.Field) (any, error) {
	value, ok := e.customValues[info.Name]
	if ok {
		return value, nil
	}

	switch info.Kind {
	case
		reflect.Int, reflect.Float32, reflect.Float64, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:

		return 42, nil
	case reflect.String:
		return "Fake" + info.Name, nil
	case reflect.Map:
		return reflect.MakeMap(info.Type), nil
	case reflect.Slice:
		return reflect.MakeSlice(info.Type, 0, 7), nil
	}

	return nil, nil
}
