# jub0bs/errutil

[![tag](https://img.shields.io/github/tag/jub0bs/errutil.svg)](https://github.com/jub0bs/errutil/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23.2-%23007d9c)
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

jub0bs/errutil requires Go 1.23.1 or above.

## Documentation

The documentation is available on [pkg.go.dev][pkgsite].

## Code coverage

![coverage](https://codecov.io/gh/jub0bs/errutil/branch/main/graphs/sunburst.svg?token=N208BHWQTM)

## Benchmarks

Here are some results of benchmarks pitting `errutil` against [`errors.As`][as]:

```text
goos: linux
goarch: amd64
pkg: github.com/jub0bs/errutil
cpu: AMD EPYC 7763 64-Core Processor
                                                                                    │     errors     │                errutil                │
                                                                                    │     sec/op     │    sec/op      vs base                │
AsAgainstErrorsPkg/nil_error,_nil_target-4                                             3.5220n ± 13%   0.7939n ± 58%  -77.46% (p=0.000 n=10)
AsAgainstErrorsPkg/nil_error,_non-nil_target-4                                         3.9450n ± 19%   0.7905n ± 39%  -79.96% (p=0.000 n=10)
AsAgainstErrorsPkg/no_match-4                                                          59.660n ± 15%    6.746n ± 86%  -88.69% (p=0.000 n=10)
AsAgainstErrorsPkg/simple_match-4                                                      71.765n ± 10%    4.199n ± 26%  -94.15% (p=0.000 n=10)
AsAgainstErrorsPkg/aser-4                                                              62.555n ±  8%    9.154n ± 13%  -85.37% (p=0.000 n=10)
AsAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                      60.700n ± 13%    8.177n ± 32%  -86.53% (p=0.000 n=10)
AsAgainstErrorsPkg/wrapper_that_contains_match-4                                       84.925n ± 14%    9.410n ± 12%  -88.92% (p=0.000 n=10)
AsAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                         113.25n ± 12%    21.67n ±  8%  -80.87% (p=0.000 n=10)
AsAgainstErrorsPkg/wrapper_that_contains_aser-4                                         82.95n ± 74%    15.78n ± 70%  -80.98% (p=0.000 n=10)
AsAgainstErrorsPkg/empty_joiner-4                                                      70.365n ± 25%    7.448n ± 15%  -89.42% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_nil-4                                          61.220n ±  2%    7.865n ± 10%  -87.15% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                                 88.53n ±  1%    10.87n ±  6%  -87.72% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                            101.65n ±  1%    17.18n ± 10%  -83.09% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                             88.95n ±  1%    12.19n ± 10%  -86.30% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_two_matches-4                                   88.12n ± 15%    11.94n ±  8%  -86.45% (p=0.000 n=10)
AsAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4       88.73n ± 14%    12.06n ± 11%  -86.40% (p=0.000 n=10)
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                       109.60n ±  8%    16.46n ± 15%  -84.98% (p=0.000 n=10)
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                    99.48n ±  8%    23.73n ± 15%  -76.15% (p=0.000 n=10)
AsAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                              333.1n ±  3%    138.9n ± 16%  -58.29% (p=0.000 n=10)
FindAgainstErrorsPkg/nil_error,_nil_target-4                                            3.095n ± 30%    3.716n ± 33%  +20.05% (p=0.035 n=10)
FindAgainstErrorsPkg/nil_error,_non-nil_target-4                                        3.993n ± 10%    3.444n ± 28%  -13.75% (p=0.001 n=10)
FindAgainstErrorsPkg/no_match-4                                                       107.050n ±  8%    7.732n ± 11%  -92.78% (p=0.000 n=10)
FindAgainstErrorsPkg/simple_match-4                                                   108.650n ± 15%    5.369n ± 25%  -95.06% (p=0.000 n=10)
FindAgainstErrorsPkg/aser-4                                                            110.45n ±  9%    48.29n ± 13%  -56.28% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                    104.35n ± 15%    10.47n ± 34%  -89.97% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_match-4                                     149.75n ± 14%    10.29n ± 19%  -93.13% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                       160.55n ± 10%    17.63n ± 15%  -89.02% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_aser-4                                      113.80n ±  3%    50.63n ± 11%  -55.51% (p=0.000 n=10)
FindAgainstErrorsPkg/empty_joiner-4                                                    103.50n ±  6%    10.21n ± 15%  -90.14% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil-4                                        105.65n ± 13%    10.64n ± 10%  -89.93% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                              139.60n ± 12%    13.81n ±  9%  -90.11% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                          157.35n ±  8%    17.79n ± 17%  -88.70% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                          135.95n ± 13%    12.89n ± 14%  -90.51% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_two_matches-4                                138.90n ±  8%    13.71n ± 12%  -90.13% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4    132.85n ±  5%    11.67n ±  6%  -91.22% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                     144.85n ± 12%    15.79n ± 15%  -89.10% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                 130.35n ±  3%    51.39n ±  8%  -60.58% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                            368.7n ±  4%    178.4n ± 21%  -51.61% (p=0.000 n=10)
geomean                                                                                 76.25n          11.95n        -84.33%

                                                                                    │    errors    │                 errutil                 │
                                                                                    │     B/op     │    B/op     vs base                     │
AsAgainstErrorsPkg/nil_error,_nil_target-4                                            0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/nil_error,_non-nil_target-4                                        0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/no_match-4                                                         0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/simple_match-4                                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/aser-4                                                             0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_contains_match-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                        0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_contains_aser-4                                       0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/empty_joiner-4                                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_nil-4                                         0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                               0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                           0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                           0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_two_matches-4                                 0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                  0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                            0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/nil_error,_nil_target-4                                          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/nil_error,_non-nil_target-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/no_match-4                                                       16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/simple_match-4                                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/aser-4                                                           16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_match-4                                    16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                      16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_aser-4                                     16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/empty_joiner-4                                                   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil-4                                       16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                             16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                         16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                         16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_two_matches-4                               16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4   16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                    16.00 ± 0%      0.00 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                          16.00 ± 0%     16.00 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                                                          ²               ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                                                                                    │    errors    │                 errutil                 │
                                                                                    │  allocs/op   │ allocs/op   vs base                     │
AsAgainstErrorsPkg/nil_error,_nil_target-4                                            0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/nil_error,_non-nil_target-4                                        0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/no_match-4                                                         0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/simple_match-4                                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/aser-4                                                             0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_contains_match-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                        0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/wrapper_that_contains_aser-4                                       0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/empty_joiner-4                                                     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_nil-4                                         0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                               0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                           0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                           0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_two_matches-4                                 0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4     0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                  0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AsAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                            0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/nil_error,_nil_target-4                                          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/nil_error,_non-nil_target-4                                      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/no_match-4                                                       1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/simple_match-4                                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/aser-4                                                           1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/wrapper_that_wraps_nil_error-4                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_match-4                                    1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_wrapper_that_contains_match-4                      1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/wrapper_that_contains_aser-4                                     1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/empty_joiner-4                                                   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil-4                                       1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_nil_and_match-4                             1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_non-nil_and_match-4                         1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_match_and_non-nil-4                         1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/joiner_that_contains_two_matches-4                               1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/deeply_nested_joiner_that_contains_non-nil_and_three_matches-4   1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners-4                                    1.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
FindAgainstErrorsPkg/mix_of_wrappers_and_joiners_that_contains_asers-4                1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
FindAgainstErrorsPkg/joiner_that_contains_many_false_asers-4                          1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                                                          ²               ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean
```

## License

All source code is covered by the [MIT License][license].

[as]: https://pkg.go.dev/errors#As
[golang]: https://go.dev/
[license]: https://github.com/jub0bs/errutil/blob/main/LICENSE
[pkgsite]: https://pkg.go.dev/github.com/jub0bs/errutil
