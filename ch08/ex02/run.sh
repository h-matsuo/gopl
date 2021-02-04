#!/usr/bin/env bash

go run github.com/h-matsuo/gopl/ch08/ex02 -port 8000 &

# Supports following FTP commands:
# - USER
# - PORT
# - LIST
# - RETR
# - QUIT

ftp localhost 8000

wait
