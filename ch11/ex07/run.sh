#!/usr/bin/env bash

go test -v -bench=. github.com/h-matsuo/gopl/ch11/ex07/intset
# ==>
# goos: darwin
# goarch: amd64
# pkg: github.com/h-matsuo/gopl/ch11/ex07/intset
# BenchmarkIntSetAdd_num1_max1
# BenchmarkIntSetAdd_num1_max1-8                  305733897                3.95 ns/op
# BenchmarkMapAdd_num1_max1
# BenchmarkMapAdd_num1_max1-8                     150945777                7.85 ns/op
# BenchmarkIntSetAdd_num1000_max1000
# BenchmarkIntSetAdd_num1000_max1000-8              338953              3426 ns/op
# BenchmarkMapAdd_num1000_max1000
# BenchmarkMapAdd_num1000_max1000-8                  87490             13257 ns/op
# BenchmarkIntSetAdd_num1000_max10000
# BenchmarkIntSetAdd_num1000_max10000-8              34885             34264 ns/op
# BenchmarkMapAdd_num10000_max10000
# BenchmarkMapAdd_num10000_max10000-8                 4263            264732 ns/op
# BenchmarkIntSetUnionWith_num1_max1
# BenchmarkIntSetUnionWith_num1_max1-8            1000000000               0.000001 ns/op
# BenchmarkMapUnionWith_num1_max1
# BenchmarkMapUnionWith_num1_max1-8               1000000000               0.000001 ns/op
# BenchmarkIntSetUnionWith_num1000_max1000
# BenchmarkIntSetUnionWith_num1000_max1000-8      1000000000               0.000000 ns/op
# BenchmarkMapUnionWith_num1000_max1000
# BenchmarkMapUnionWith_num1000_max1000-8         1000000000               0.000046 ns/op
# BenchmarkIntSetUnionWith_num1000_max10000
# BenchmarkIntSetUnionWith_num1000_max10000-8     1000000000               0.000002 ns/op
# BenchmarkMapUnionWith_num10000_max10000
# BenchmarkMapUnionWith_num10000_max10000-8       1000000000               0.000565 ns/op
# PASS
# ok      github.com/h-matsuo/gopl/ch11/ex07/intset       10.341s
