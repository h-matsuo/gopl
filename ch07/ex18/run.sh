#!/usr/bin/env bash

curl https://www.w3.org/TR/2006/REC-xml11-20060816/ 2> /dev/null \
  | go run github.com/h-matsuo/gopl/ch07/ex18
