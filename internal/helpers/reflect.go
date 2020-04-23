package helpers

import (
	"fmt"
	"reflect"
)

// CallByName calls object func by its name
func CallByName(o interface{}, name string, params ...interface{}) (out []reflect.Value, err error) {
	v := reflect.ValueOf(o)
	m := v.MethodByName(name)

	if !m.IsValid() {
		return make([]reflect.Value, 0), &ErrorHelper{
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
