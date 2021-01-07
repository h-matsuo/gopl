#!/usr/bin/env bash

curl 'https://gopl.io/ch1/helloworld?go-get=1'
# ==>
# <!DOCTYPE html>
# <html>
# <head>
# <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
# <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
# </head>
# <body>
# </body>
# </html>

# The `<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">` part
# indicates the source code hosted on `https://github.com/adonovan/gopl.io`.
