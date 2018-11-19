#!/bin/sh
#export LD_LIBRARY_PATH=/opt/glibc2.14/lib:$LD_LIBRARY_PATH
#go build -a -o jyj_blog main.go
kill -9 $(pidof /var/gowww/src/jyj-go/jyj_blog)
nohup /var/gowww/src/jyj-go/jyj_blog -c /var/gowww/src/jyj-go/conf/conf.yaml -l /var/gowww/src/jyj-go/conf/log.xml > /var/gowww/src/jyj-go/sys.log  2>&1 &
