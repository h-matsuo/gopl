#!/usr/bin/env bash

go run github.com/h-matsuo/gopl/ch08/ex02 -port 8000 &

sleep 1

ftp $USER@localhost 8000
# Supports: pwd, cd, ls, get, put, ascii, binary

wait
