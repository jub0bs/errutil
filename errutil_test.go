package errutil_test

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
	"slices"
	"testing"

	"github.com/jub0bs/errutil"
)

// 	cases := []struct {
// 		desc   string
// 		target *simpleError
// 	}{
// 		{
// 			desc:   "nil target",
// 			target: nil,
// 		}, {
// 			desc:   "non-nil target",
// 			target: new(simpleError),
// 		},
// 	}
// 	for _, tc := range cases {
// 		f := func(t *testing.T) {
// 			var err error
// 			got := errutil.As(err, tc.target)
// 			if got {
// 				const tmpl = "As(%v, %T(%v)): got true; want false"
// 				t.Errorf(tmpl, err, tc.target, tc.target)
// 			}
// 		}
// 		t.Run(tc.desc, f)
// 	}
// }

func TestAsPanicsForNonNilErrAndNilTarget(t *testing.T) {
	err := errors.New("oh no!")
	var target *simpleError
	defer func() {
		if r := recover(); r == nil {
			const tmpl = "As(%v, %T(%v)) did not panic"
			t.Errorf(tmpl, err, target, target)
		}
	}()
	errutil.As(err, target)
}

func TestAs(t *testing.T) {
	for _, tc := range cases {
		f := func(t *testing.T) {
			match := errutil.As(tc.err, tc.target)
			if match != tc.match {
				const tmpl = "errutil.As(err, %[1]T(%[1]v)): got %t; want %t"
				t.Fatalf(tmpl, tc.target, match, tc.match)
			}
			if !match {
				return
			}
			if got := *tc.target; got != tc.want {
				t.Fatalf("*target: got %#v; want %#v", got, tc.want)
			}
			if match != errors.As(tc.err, tc.target) { // sanity check
				const tmpl = "errutil.As(err, %[1]T(%[1]v)) !=  errors.As(err, %[1]T(%[1]v))"
				t.Fatalf(tmpl, tc.target)
			}
		}
		t.Run(tc.desc, f)
	}
}

// see https://github.com/golang/go/issues/66455#issuecomment-2018372473
func TestAsTargetWiderThanError(t *testing.T) {
	err := new(net.DNSError)
	type timeouter interface {
		Timeout() bool
		error
	}
	var _ timeouter = err
	var target = new(timeouter)
	match := errutil.As(err, target)
	if !match {
		const tmpl = "errutil.As(err, %[1]T(%[1]v)): got false; want true"
		t.Fatalf(tmpl, target)
	}
	if got := *target; got != err {
		t.Fatalf("*target: got %#v; want %#v", got, err)
	}
	if match != errors.As(err, target) { // sanity check
		const tmpl = "errutil.As(err, %[1]T(%[1]v)) !=  errors.As(err, %[1]T(%[1]v))"
		t.Fatalf(tmpl, target)
	}
}

func ExampleAs() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errutil.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// Failed at path: non-existing
}

// In this example, the target's desired type is an interface type other than
// error:
//
//	interface { Timeout() bool }
//
// A simple workaround for coaxing [As] into accepting such a target simply
// consists in [embedding] error in the target's desired type.
//
// [embedding]: https://go.dev/ref/spec#Embedded_interfaces
func ExampleAs_interface() {
	fakeLookupIP := func(_ string) ([]net.IP, error) {
		return nil, &net.DNSError{IsTimeout: true}
	}
	if _, err := fakeLookupIP("invalid-TLD.123"); err != nil {
		var to interface {
			Timeout() bool
			error // for errutil.As to accept &to as its second argument
		}
		if errutil.As(err, &to) {
			fmt.Printf("Timed out: %t\n", to.Timeout())
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// Timed out: true
}

func BenchmarkAs(b *testing.B) {
	for _, bc := range cases {
		f := func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				errutil.As(bc.err, bc.target)
			}
		}
		b.Run(bc.desc, f)
	}
}

func BenchmarkAsAgainstErrorsPkg(b *testing.B) {
	for _, bc := range cases {
		f := func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				errors.As(bc.err, bc.target)
			}
		}
		b.Run("v=errors/"+bc.desc, f)

		f = func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				errutil.As(bc.err, bc.target)
			}
		}
		b.Run("v=errutil/"+bc.desc, f)
	}
}

func TestFind(t *testing.T) {
	for _, tc := range cases {
		f := func(t *testing.T) {
			got, match := errutil.Find[simpleError](tc.err)
			if match != tc.match || got != tc.want {
				const tmpl = "errutil.Find(err): got %#v, %t; want %#v, %t"
				t.Fatalf(tmpl, got, match, tc.want, tc.match)
			}
		}
		t.Run(tc.desc, f)
	}
}

// see https://github.com/golang/go/issues/66455#issuecomment-2018372473
func TestFindTargetWiderThanError(t *testing.T) {
	err := new(net.DNSError)
	type timeouter interface {
		Timeout() bool
		error
	}
	var _ timeouter = err
	got, match := errutil.Find[timeouter](err)
	want := timeouter(err)
	if !match || got != want {
		const tmpl = "errutil.Find(err): got %#v, %t; want %#v, true"
		t.Fatalf(tmpl, got, match, want)
	}
}

func ExampleFind() {
	if _, err := os.Open("non-existing"); err != nil {
		if pathError, ok := errutil.Find[*fs.PathError](err); ok {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// Failed at path: non-existing
}

// In this example, the result's desired type is an interface type other than
// error:
//
//	interface { Timeout() bool }
//
// A simple workaround for coaxing [Find] into accepting such a type argument
// simply consists in [embedding] error in the result's desired type.
//
// [embedding]: https://go.dev/ref/spec#Embedded_interfaces
func ExampleFind_interface() {
	fakeLookupIP := func(_ string) ([]net.IP, error) {
		return nil, &net.DNSError{IsTimeout: true}
	}
	if _, err := fakeLookupIP("invalid-TLD.123"); err != nil {
		type timeouter interface {
			Timeout() bool
			error // for errutil.Find to accept timeouter as its type argument
		}
		if to, ok := errutil.Find[timeouter](err); ok {
			fmt.Printf("Timed out: %t\n", to.Timeout())
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// Timed out: true
}

func BenchmarkFind(b *testing.B) {
	for _, bc := range cases {
		f := func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				errutil.Find[simpleError](bc.err)
			}
		}
		b.Run(bc.desc, f)
	}
}

func BenchmarkFindAgainstErrorsPkg(b *testing.B) {
	for _, bc := range cases {
		f := func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				findErrorsPkg[simpleError](bc.err)
			}
		}
		b.Run("v=errors/"+bc.desc, f)

		f = func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				errutil.Find[simpleError](bc.err)
			}
		}
		b.Run("v=errutil/"+bc.desc, f)
	}
}

// A version of errors.Find implemented in terms of errors.As;
// useful for benchmarks.
func findErrorsPkg[T error](err error) (T, bool) {
	if err == nil {
		var zero T
		return zero, false
	}
	target := new(T)
	ok := errors.As(err, target)
	return *target, ok
}

type TestCase[T error] struct {
	desc   string
	err    error
	target *T
	match  bool
	want   T
}

var cases = []TestCase[simpleError]{
	{
		desc:   "nil error, nil target",
		err:    nil,
		target: nil,
		match:  false,
	}, {
		desc:   "nil error, non-nil target",
		err:    nil,
		target: new(simpleError),
		match:  false,
	}, {
		desc:   "no match",
		err:    errors.New("oh no!"),
		target: new(simpleError),
	}, {
		desc:   "simple match",
		err:    simpleError{msg: "foo"},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc:   "aser",
		err:    aser{msg: "foo", f: masqueradeAsSimpleError},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc:   "wrapper that wraps nil error",
		err:    wrapper{},
		target: new(simpleError),
		match:  false,
	}, {
		desc: "wrapper that contains match",
		err: wrapper{
			simpleError{msg: "foo"},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "deeply nested wrapper that contains match",
		err: wrapper{
			wrapper{
				wrapper{simpleError{msg: "foo"}},
			},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "wrapper that contains aser",
		err: wrapper{
			aser{msg: "foo", f: masqueradeAsSimpleError},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc:   "empty joiner",
		err:    joiner{},
		target: new(simpleError),
		match:  false,
	}, {
		desc:   "joiner that contains nil",
		err:    joiner{nil},
		target: new(simpleError),
		match:  false,
	}, {
		desc: "joiner that contains nil and match",
		err: joiner{
			nil,
			simpleError{msg: "foo"},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "joiner that contains non-nil and match",
		err: joiner{
			errors.New("oh no!"),
			simpleError{msg: "foo"},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "joiner that contains match and non-nil",
		err: joiner{
			simpleError{msg: "foo"},
			errors.New("oh no!"),
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "joiner that contains two matches",
		err: joiner{
			simpleError{msg: "foo"},
			simpleError{msg: "bar"},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "deeply nested joiner that contains non-nil and three matches",
		err: joiner{
			simpleError{msg: "foo"},
			joiner{
				errors.New("oh no!"),
				simpleError{msg: "bar"},
				simpleError{msg: "baz"},
			},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "mix of wrappers and joiners",
		err: joiner{
			wrapper{
				simpleError{msg: "foo"},
			},
			joiner{
				errors.New("oh no!"),
				wrapper{simpleError{msg: "bar"}},
				simpleError{msg: "baz"},
			},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc: "mix of wrappers and joiners that contains asers",
		err: joiner{
			wrapper{
				aser{msg: "foo", f: masqueradeAsSimpleError},
			},
			joiner{
				errors.New("oh no!"),
				wrapper{aser{msg: "bar", f: masqueradeAsSimpleError}},
				aser{msg: "baz", f: masqueradeAsSimpleError},
			},
		},
		target: new(simpleError),
		match:  true,
		want:   simpleError{msg: "foo"},
	}, {
		desc:   "joiner that contains many false asers",
		err:    joiner(slices.Repeat([]error{aser{msg: "foo"}}, 16)),
		target: new(simpleError),
		match:  false,
	},
}

type simpleError struct {
	msg string
}

func (s simpleError) Error() string {
	return s.msg
}

type wrapper struct {
	err error
}

func (w wrapper) Error() string {
	return ""
}

func (w wrapper) Unwrap() error {
	return w.err
}

type joiner []error

func (j joiner) Error() string {
	return ""
}

func (j joiner) Unwrap() []error {
	return j
}

type aser struct {
	msg string
	f   func(aser, any) bool
}

func (a aser) Error() string {
	return a.msg
}

func (a aser) As(target any) bool {
	if a.f == nil {
		return false
	}
	return a.f(a, target)
}

func masqueradeAsSimpleError(a aser, target any) bool {
	switch x := target.(type) {
	case *simpleError:
		*x = simpleError{msg: a.msg}
		return true
	default:
		return false
	}
}
