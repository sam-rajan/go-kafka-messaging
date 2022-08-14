#!/usr/bin/bash
APP=$1

cd "$(dirname "${0}")/.."

if [[ $APP == "SENDER" ]]
then
    echo "Starting sender application"
    go run cmd/sender-app/main.go
elif [[ $APP == "RECEIVER" ]]
then
    echo "Starting receiver application"
    go run cmd/receiver-app/main.go
else 
    echo "Invalid Option"
fi