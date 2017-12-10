#!/usr/bin/env bash
set -e

# curl https://raw.githubusercontent.com/daomor/precious/v0.1-alpha/install.sh | sh

DOWNLOAD_URL=https://github.com/daomor/precious/releases/download/v0.1-alpha/precious

wget ${DOWNLOAD_URL} 2>/dev/null || curl -OL  ${DOWNLOAD_URL}

# Create directory.
if [ ! -d /usr/local/precious ]; then
  echo "Info: creating directory '/usr/local/precious'."
  mkdir -p /usr/local/precious
  chown $USER /usr/local/precious
fi

# Check for download.
if [ ! -f ./precious ]; then
  echo "Warn: the program was not downloaded, or, was already moved to '/usr/local/bin' previously."
else

  # Check for program already installed.
  if [ -f /usr/local/bin/precious ]; then
      echo "Info: existing version of the program will be overridden."
      mv ./precious /usr/local/bin
  fi

  echo "Info: moving ./precious to '/usr/local/bin'."
  mv ./precious /usr/local/bin
  chmod +x /usr/local/bin/precious
fi

# Create servers yaml.
if [ ! -f /usr/local/precious/servers.yaml ]; then
  echo "Info: creating server data file at /usr/local/precious/servers.yaml."
  touch /usr/local/precious/servers.yaml
else
  echo "Info: server.yaml already exists."
fi