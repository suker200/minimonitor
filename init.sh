#!/bin/sh

cd /
./minimonitor &

echo "10"
/build/nginx/sbin/nginx -c /nginx.conf 
