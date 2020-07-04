#!/bin/bash

kubectl apply -k ./mysql/base/ >/dev/null 2>&1

if [ $? = 0 ]; then
  echo "mysql start"
else
  echo "Error"
fi

kubectl apply -k ./go/github.com/ke6ch/api/base/ >/dev/null 2>&1 

if [ $? = 0 ]; then
  echo "api start"
else
  echo "Error"
fi

kubectl apply -k ./next/base/ >/dev/null 2>&1

if [ $? = 0 ]; then
  echo "app start"
else
  echo "Error"
fi

kubectl apply -k ./nginx/base/ >/dev/null 2>&1

if [ $? = 0 ]; then
  echo "proxy start"
else
  echo "Error"
fi

