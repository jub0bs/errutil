# jub0bs/errutil

[![tag](https://img.shields.io/github/tag/jub0bs/errutil.svg)](https://github.com/jub0bs/errutil/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.25.0-%23007d9c)
[![Go Reference](https://pkg.go.dev/badge/github.com/jub0bs/errutil.svg)](https://pkg.go.dev/github.com/jub0bs/errutil)
[![license](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](https://github.com/jub0bs/errutil/raw/main/LICENSE)
[![build](https://github.com/jub0bs/errutil/actions/workflows/errutil.yml/badge.svg)](https://github.com/jub0bs/errutil/actions/workflows/errutil.yml)
[![codecov](https://codecov.io/gh/jub0bs/errutil/branch/main/graph/badge.svg?token=N208BHWQTM)](https://app.codecov.io/gh/jub0bs/errutil/tree/main)
[![goreport](https://goreportcard.com/badge/jub0bs/errutil)](https://goreportcard.com/report/jub0bs/errutil)

A collection of utility functions for working with [Go][golang] errors.

## Installation

```shell
go get github.com/jub0bs/errutil
```

jub0bs/errutil requires [Go 1.25.0][go1.25] or above. Note that it only
[supports][release-policy] the two most recent minor versions of Go.

## Documentation

The documentation is available on [pkg.go.dev][pkgsite].

## Code coverage

![coverage](https://codecov.io/gh/jub0bs/errutil/branch/main/graphs/sunburst.svg?token=N208BHWQTM)

## Benchmarks

Here are some results of benchmarks pitting `errutil.Find` against [`errors.As`][as]:

```text
goos: darwin
goarch: arm64
pkg: github.com/jub0bs/errutil
cpu: Apple M4
                                                                      │   errors.As   │             errutil.Find             │
                                                                      │    sec/op     │    sec/op     vs base                │
All/c=nil_error,_nil_target-10                                           3.717n ± 17%   4.058n ±  1%   +9.19% (p=0.041 n=10)
All/c=nil_error,_non-nil_target-10                                      21.210n ± 26%   3.519n ±  1%  -83.41% (p=0.000 n=10)
All/c=no_match-10                                                       66.650n ±  7%   7.743n ±  1%  -88.38% (p=0.000 n=10)
All/c=simple_match-10                                                   80.370n ±  4%   5.771n ±  0%  -92.82% (p=0.000 n=10)
All/c=aser-10                                                            69.14n ± 11%   28.26n ±  7%  -59.12% (p=0.000 n=10)
All/c=wrapper_that_wraps_nil_error-10                                   66.480n ±  7%   8.457n ±  1%  -87.28% (p=0.000 n=10)
All/c=wrapper_that_contains_match-10                                    87.830n ±  2%   9.075n ±  2%  -89.67% (p=0.000 n=10)
All/c=deeply_nested_wrapper_that_contains_match-10                      109.65n ±  2%   19.20n ±  5%  -82.49% (p=0.000 n=10)
All/c=wrapper_that_contains_aser-10                                      80.50n ±  5%   30.69n ±  2%  -61.87% (p=0.000 n=10)
All/c=empty_joiner-10                                                   71.935n ±  6%   8.698n ± 16%  -87.91% (p=0.000 n=10)
All/c=joiner_that_contains_nil-10                                       71.015n ±  3%   9.144n ±  0%  -87.12% (p=0.000 n=10)
All/c=joiner_that_contains_nil_and_match-10                              92.06n ±  3%   11.33n ±  1%  -87.69% (p=0.000 n=10)
All/c=joiner_that_contains_non-nil_and_match-10                         106.25n ±  3%   17.88n ± 11%  -83.18% (p=0.000 n=10)
All/c=joiner_that_contains_match_and_non-nil-10                          91.91n ±  1%   11.26n ±  7%  -87.75% (p=0.000 n=10)
All/c=joiner_that_contains_two_matches-10                                91.69n ±  5%   11.37n ±  6%  -87.60% (p=0.000 n=10)
All/c=deeply_nested_joiner_that_contains_non-nil_and_three_matches-10    91.21n ±  3%   11.28n ±  0%  -87.63% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners-10                                    106.80n ±  3%   15.64n ±  8%  -85.35% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners_that_contains_asers-10                 94.85n ±  2%   37.70n ±  4%  -60.26% (p=0.000 n=10)
All/c=joiner_that_contains_many_false_asers-10                           300.6n ±  1%   160.8n ±  1%  -46.48% (p=0.000 n=10)
geomean                                                                  71.62n         13.14n        -81.65%

                                                                      │  errors.As   │              errutil.Find               │
                                                                      │     B/op     │    B/op     vs base                     │
All/c=nil_error,_nil_target-10                                          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
All/c=nil_error,_non-nil_target-10                                      16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=no_match-10                                                       16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=simple_match-10                                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=aser-10                                                           16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
All/c=wrapper_that_wraps_nil_error-10                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=wrapper_that_contains_match-10                                    16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=deeply_nested_wrapper_that_contains_match-10                      16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=wrapper_that_contains_aser-10                                     16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
All/c=empty_joiner-10                                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_nil-10                                       16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_nil_and_match-10                             16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_non-nil_and_match-10                         16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_match_and_non-nil-10                         16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_two_matches-10                               16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=deeply_nested_joiner_that_contains_non-nil_and_three_matches-10   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners-10                                    16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners_that_contains_asers-10                16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
All/c=joiner_that_contains_many_false_asers-10                          16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                                            ²               ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                                                                      │  errors.As   │              errutil.Find               │
                                                                      │  allocs/op   │ allocs/op   vs base                     │
All/c=nil_error,_nil_target-10                                          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
All/c=nil_error,_non-nil_target-10                                      1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=no_match-10                                                       1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=simple_match-10                                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=aser-10                                                           1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
All/c=wrapper_that_wraps_nil_error-10                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=wrapper_that_contains_match-10                                    1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=deeply_nested_wrapper_that_contains_match-10                      1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=wrapper_that_contains_aser-10                                     1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
All/c=empty_joiner-10                                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_nil-10                                       1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_nil_and_match-10                             1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_non-nil_and_match-10                         1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_match_and_non-nil-10                         1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=joiner_that_contains_two_matches-10                               1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=deeply_nested_joiner_that_contains_non-nil_and_three_matches-10   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners-10                                    1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
All/c=mix_of_wrappers_and_joiners_that_contains_asers-10                1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
All/c=joiner_that_contains_many_false_asers-10                          1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                                            ²               ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean
```

## License

All source code is covered by the [MIT License][license].

[as]: https://pkg.go.dev/errors#As
[go1.25]: https://tip.golang.org/doc/go1.25
[golang]: https://go.dev/
[license]: https://github.com/jub0bs/errutil/blob/main/LICENSE
[pkgsite]: https://pkg.go.dev/github.com/jub0bs/errutil
[release-policy]: https://go.dev/doc/devel/release#policy
