#!/usr/bin/bash
APP=$1

cd "$(dirname "${0}")/.."

if [[ $APP == "SENDER" ]]
then
    echo "Starting sender application"
    go run cmd/sender-app/main.go $2
elif [[ $APP == "RECEIVER" ]]
then
    CONSUMERS=${2:-0}
    echo "Starting receiver application"
    go run cmd/receiver-app/main.go $CONSUMERS
else 
    echo "Invalid Option"
fi