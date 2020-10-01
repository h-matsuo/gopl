#!/bin/sh

export EDITOR=vim

go run github.com/h-matsuo/gopl/ch04/ex11 \
  -user <USERNAME> \
  -token <PERSONAL_ACCESS_TOKEN> \
  -owner <OWNER> \
  -repo <REPO> \
  -create \
  -title 'My New Issue'

# go run github.com/h-matsuo/gopl/ch04/ex11 \
#   -user <USERNAME> \
#   -token <PERSONAL_ACCESS_TOKEN> \
#   -owner <OWNER> \
#   -repo <REPO> \
#   -get \
#   -number <ISSUE_NUMBER>

# go run github.com/h-matsuo/gopl/ch04/ex11 \
#   -user <USERNAME> \
#   -token <PERSONAL_ACCESS_TOKEN> \
#   -owner <OWNER> \
#   -repo <REPO> \
#   -update \
#   -number <ISSUE_NUMBER> \
#   -title 'My Updated Issue'

# go run github.com/h-matsuo/gopl/ch04/ex11 \
#   -user <USERNAME> \
#   -token <PERSONAL_ACCESS_TOKEN> \
#   -owner <OWNER> \
#   -repo <REPO> \
#   -close \
#   -number <ISSUE_NUMBER>
