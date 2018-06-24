#!/bin/bash
trap control_c SIGTERM
trap control_c SIGINT

function control_c() {
	echo -en "\nBringing docker-compose down...\n"
	docker-compose down
	exit $?
}

while true
do 
	docker build --rm -q -t market . && docker-compose up -d
	inotifywait -q -e modify --timefmt '%d-%m-%Y %H:%M' --format '%T %w %f %e' *.go */*.go >> log.txt
	echo -e "\n\nmodification detected... one moment"
	docker-compose down
done
