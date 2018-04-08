#!/bin/bash

if [ $# != 2 ] ; then
  echo "para error"
  exit 1;
fi

if [ $? != 0 ] ; then
  echo "para error"$?
  exit 1;
fi

cat `find $1 -name "*$2"`