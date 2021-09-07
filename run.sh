#!/bin/sh
squidclient mgr:active_requests | egrep -E '^Connection: 0x*' -A 15 > /etc/squid/OwlActivesRequest