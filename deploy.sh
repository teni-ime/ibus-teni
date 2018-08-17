#!/bin/sh
if [ -z "$OSC_USER" ] || [ -z "$OSC_PASS" ] || [ -z "$OSC_PATH" ]
then
  echo OSC_USER|OSC_PASS|OSC_PATH is not defined, do not deploy
  exit 0
fi

sudo apt-get update
sudo apt-get install -y osc
osc --version

echo "[general]" >> ~/.oscrc
echo "apiurl = https://api.opensuse.org" >> ~/.oscrc
echo "[https://api.opensuse.org]" >> ~/.oscrc
echo "user = $OSC_USER" >> ~/.oscrc
echo "pass = $OSC_PASS" >> ~/.oscrc

SRC_DIR=$(pwd)
mkdir ../obs
cd ../obs
osc checkout $OSC_PATH
cd $SRC_DIR
make build src DESTDIR=../obs/$OSC_PATH
cd ../obs/$OSC_PATH
osc addremove
osc st

