#!/usr/bin/env bash

TZ=US/Eastern    go run github.com/h-matsuo/gopl/ch08/ex01/clock -port 8010 &
TZ=Asia/Tokyo    go run github.com/h-matsuo/gopl/ch08/ex01/clock -port 8020 &
TZ=Europe/London go run github.com/h-matsuo/gopl/ch08/ex01/clock -port 8030 &

sleep 1

go run github.com/h-matsuo/gopl/ch08/ex01/clockwall \
  NewYork=localhost:8010 \
    Tokyo=localhost:8020 \
   London=localhost:8030
