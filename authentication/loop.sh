#!/bin/sh
set -e
set -x

sigint_handler() {
  kill $(jobs -p)
  exit
}

trap sigint_handler SIGINT

while true; do
  go build
  ./authentication &
  inotifywait -e modify -r `pwd`
  kill $(jobs -p)
done
