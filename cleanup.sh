#!/usr/bin/env bash
set -e

if [ -f /usr/local/bin/precious ]; then
  rm /usr/local/bin/precious
fi

rm -rf /usr/local/precious