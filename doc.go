// Package errutil provides utility functions for working with errors.
//
// The advent of [errors.As] in the standard library predates that of
// parametric polymorphism (generics) in the language.
// As a result, [errors.As] is not as ergonomic, type-safe, or efficient as it
// ideally could be.
// Functions [As] and [Find] are inspired by several unaccepted proposals
// (see issues [51945], [56949], and [64771]) and aim to address those
// limitations.
//
// In most cases, [As] can be used as a drop-in replacement for [errors.As].
//
// [Find] is a more efficient and arguably more ergonomic alternative to [As].
// Incidentally, [the error-inspection draft design proposal] suggests that
// [errors.As] would have been very similar to [Find] if the Go team had
// cracked the parametric-polymorphism nut in time for [errors.As]'s inception
// in the standard library.
// In many cases, a call to [errors.As] can advantageously be refactored to a
// call to [Find].
//
// If you can use Go 1.26 or above, simply rely on [errors.AsType] instead of
// [As] or [Find].
//
// [51945]: https://github.com/golang/go/issues/51945
// [56949]: https://github.com/golang/go/issues/56949
// [64771]: https://github.com/golang/go/issues/64771
// [the error-inspection draft design proposal]: https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md#the-is-and-as-functions
package errutil
