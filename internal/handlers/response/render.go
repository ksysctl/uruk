package response

import (
	"fmt"
	"reflect"
)

func render(o interface{}, name string, params ...interface{}) (out []reflect.Value, err error) {
	v := reflect.ValueOf(o)
	m := v.MethodByName(name)

	if !m.IsValid() {
		return make([]reflect.Value, 0), &ErrorResponse{
			msg: fmt.Sprintf("Method not found %s", name),
		}
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	out = m.Call(in)

	return
}
