//go:build go1.26

package errutil

import "errors"

func as[T error](err error, target *T) bool {
	if err == nil {
		return false
	}
	if target == nil {
		panic("errutil: target cannot be nil")
	}
	var ok bool
	*target, ok = find[T](err)
	return ok
}

func find[T error](err error) (T, bool) {
	return errors.AsType[T](err)
}
