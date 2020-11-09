#!/bin/bash
#chmod u+x this_script.sh
echo "inside goBuild.sh ..."
echo "..."
#rm main main.zip

echo ""
echo "build into Linux 64-bits binary..."
platform='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
  platform='linux'
  echo "linux os is detected"
  if [[ "$(uname -m)" == 'x86_64' ]]; then
    echo "linux 64-bit os is detected"
    go build -o main *.go
  else
    echo "non linux 64-bit os is detected"
    GOOS=linux GOARCH-amd64 go build *.go -o main 
  fi
else
  echo "non linux 64-bit os is detected"
  GOOS=linux GOARCH-amd64 go build *.go -o main 
fi

zip main.zip main .env
echo "confirm the zip file include the .env file!"
