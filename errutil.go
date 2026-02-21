package errutil

// As finds the first error in err's tree that matches target,
// and if one is found, sets target to that error value and returns true.
// Otherwise, it returns false.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, As examines err followed by a depth-first traversal of its children.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target
// or if the error has a method As(any) bool such that As(target) returns true.
// In the latter case, the As method is responsible for setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if err is not nil and target is nil.
//
// Note that an instantiation of the form As[error] constitutes a (harmless)
// programming mistake, as it is never useful;
// such a mistake is similar to passing a value of type *error as the second
// argument of [errors.As], a mistake which is covered by a vet check.
func As[T error](err error, target *T) bool {
	if err == nil {
		return false
	}
	if target == nil {
		panic("errutil: target cannot be nil")
	}
	return as(err, target)
}

func as[T error](err error, target *T) bool {
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
				if as(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

// Find finds the first error in err's tree that matches type T,
// and if so, returns the corresponding value and true.
// Otherwise, it returns the zero value and false.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, Find examines err followed by a depth-first traversal of its children.
//
// An error matches type T if type-asserting it to T succeeds,
// or if the error has a method As(any) bool such that As(target),
// where target is any non-nil value of type *T, returns true.
// In the latter case, the As method is responsible for setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// Note that an instantiation of the form Find[error] constitutes a (harmless)
// programming mistake, as it is never useful;
// such a mistake is similar to passing a value of type *error as the second
// argument of [errors.As], a mistake which is covered by a vet check.
func Find[T error](err error) (T, bool) {
	if err == nil {
		var zero T
		return zero, false
	}
	var ptr *T
	return find(err, &ptr)
}

func find[T error](err error, ptr2 **T) (T, bool) {
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
				if x, ok := find(err, ptr2); ok {
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
