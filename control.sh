#!/usr/bin/env bash

source env.bash

dir=`pwd`

services=(authentication authorization accounting)

function build () {
  for item in ${services[*]}
  do
    cd $dir/src/$item
    echo "Running 'go build' in src/$item"
    go build
    cd $dir
  done
}

function mykillall () {
  kill $(jobs -p)
}

function run () {
  mykillall

  for item in ${services[*]}
  do
    echo "Running $item"
    $dir/src/$item/$item &
  done
}

case $1 in
  build)
    build
    ;;
  run)
    run
    ;;
  killall)
    mykillall
    ;;
esac
