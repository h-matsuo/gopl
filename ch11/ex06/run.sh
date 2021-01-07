#!/usr/bin/env bash

go test -v -bench=. github.com/h-matsuo/gopl/ch11/ex06
# ==>
# goos: darwin
# goarch: amd64
# pkg: github.com/h-matsuo/gopl/ch11/ex06
# BenchmarkPopCount_0
# BenchmarkPopCount_0-8                           315641566                3.82 ns/op
# BenchmarkLSBPopCount_0
# BenchmarkLSBPopCount_0-8                        20075414                60.6 ns/op
# BenchmarkBitClearPopCount_0
# BenchmarkBitClearPopCount_0-8                   619923427                1.65 ns/op
# BenchmarkPopCount_FFFF
# BenchmarkPopCount_FFFF-8                        313331206                3.82 ns/op
# BenchmarkLSBPopCount_FFFF
# BenchmarkLSBPopCount_FFFF-8                     20127182                59.3 ns/op
# BenchmarkBitClearPopCount_FFFF
# BenchmarkBitClearPopCount_FFFF-8                163099068                7.44 ns/op
# BenchmarkPopCount_FFFFFFFF
# BenchmarkPopCount_FFFFFFFF-8                    312571825                3.88 ns/op
# BenchmarkLSBPopCount_FFFFFFFF
# BenchmarkLSBPopCount_FFFFFFFF-8                 20026806                59.8 ns/op
# BenchmarkBitClearPopCount_FFFFFFFF
# BenchmarkBitClearPopCount_FFFFFFFF-8            54023569                21.4 ns/op
# BenchmarkPopCount_FFFFFFFFFFFFFFFF
# BenchmarkPopCount_FFFFFFFFFFFFFFFF-8            315649844                3.85 ns/op
# BenchmarkLSBPopCount_FFFFFFFFFFFFFFFF
# BenchmarkLSBPopCount_FFFFFFFFFFFFFFFF-8         20129488                59.1 ns/op
# BenchmarkBitClearPopCount_FFFFFFFFFFFFFFFF
# BenchmarkBitClearPopCount_FFFFFFFFFFFFFFFF-8    24566768                48.4 ns/op
# PASS
# ok      github.com/h-matsuo/gopl/ch11/ex06      17.181s
