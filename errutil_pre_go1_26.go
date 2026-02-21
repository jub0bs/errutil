//go:build !go1.26

package errutil

func as[T error](err error, target *T) bool {
	if err == nil {
		return false
	}
	if target == nil {
		panic("errutil: target cannot be nil")
	}
	return asRec(err, target)
}

func asRec[T error](err error, target *T) bool {
	for {
		if x, ok := err.(T); ok {
			*target = x
			return true
		}
		if x, ok := err.(interface{ As(any) bool }); ok && x.As(target) {
			return true
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if err == nil {
					continue
				}
				if asRec(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

func find[T error](err error) (T, bool) {
	if err == nil {
		var zero T
		return zero, false
	}
	var ptr *T
	return findRec(err, &ptr)
}

func findRec[T error](err error, ptr2 **T) (T, bool) {
	for {
		x, ok := err.(T)
		if ok {
			return x, true
		}
		if x, ok := err.(interface{ As(any) bool }); ok {
			if *ptr2 == nil {
				*ptr2 = new(T)
			}
			if x.As(*ptr2) {
				return **ptr2, true
			}
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				var zero T
				return zero, false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if err == nil {
					continue
				}
				if x, ok := findRec(err, ptr2); ok {
					return x, true
				}
			}
			var zero T
			return zero, false
		default:
			var zero T
			return zero, false
		}
	}
}
