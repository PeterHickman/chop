#!/bin/sh

if ! [ "$(id -u)" = 0 ]; then
  echo 'You must be root to do this.' 1>&2
  exit 1
fi

##
# When installing we would install as root, however on
# BSD based systems (such as OSX) this is not a valid
# group / user for install so we use the numeric value
##
ROOT="0"

TARGET="/usr/local/bin/chop"

NIM_FOUND=$(command -v nim)

CHOICE="$1"

if [ "$CHOICE" = "nim" ]; then
  if [ -z "$NIM_FOUND" ]; then
    echo "The nim compiler is not found, unable to install"
    exit 1
  fi

  echo "Installing the nim version to $TARGET"
  nim compile -d:release chop.nim
  install -g $ROOT -o $ROOT -m 0755 chop $TARGET
  rm chop
elif [ "$CHOICE" = "ruby" ]; then
  echo "Installing the ruby version to $TARGET"
  install -g $ROOT -o $ROOT -m 0755 chop.rb $TARGET
else
  echo "Pass 'nim' or 'ruby' to install the appropriate version"
  exit 1
fi
