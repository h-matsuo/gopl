#!/usr/bin/env bash

# Run web server
go run github.com/h-matsuo/gopl/ch12/ex12 &

sleep 1

# Test validations

curl 'localhost:12345/search?uid=9E0D5D7E-FB45-49A9-9446-703DEA56760F&e=test%40example.com&e=foo%40bar.com'
# --> OK

curl 'localhost:12345/search?uid=invalid'
# --> NG: invalid uuid

curl 'localhost:12345/search?e=invalid'
# --> NG: invalid email

wait
