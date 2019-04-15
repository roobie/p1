#!/usr/bin/env bash

set -e
set -x

source env.bash

cd $2

sigint_handler() {
  kill $(jobs -p)
  exit
}

trap sigint_handler SIGINT

while true; do
  go build
  ./$1 &
  inotifywait -e modify -r `pwd`
  kill $(jobs -p)
done
