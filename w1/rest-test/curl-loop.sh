#!/bin/bash

while true
do
	curl -sX GET http://${1}:8080/2020/myinfo |json_pp
	sleep 2
done
