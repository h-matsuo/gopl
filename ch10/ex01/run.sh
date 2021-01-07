#!/usr/bin/env bash

go run gopl.io/ch3/mandelbrot \
  | go run github.com/h-matsuo/gopl/ch10/ex01 -out jpeg \
  | go run github.com/h-matsuo/gopl/ch10/ex01 -out gif \
  | go run github.com/h-matsuo/gopl/ch10/ex01 -out png \
  > chained-out.png
