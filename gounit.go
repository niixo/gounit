package gounit

import "reflect"

func AssertEqual(e, a interface{}) bool {
	return reflect.DeepEqual(e, a)
}

