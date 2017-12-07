#!/usr/bin/env bash

# curl https://raw.githubusercontent.com/daomor/precious/v0.1-alpha/install.sh | sh

wget https://github.com/daomor/precious/releases/download/v0.1-alpha/precious 2>/dev/null || curl -O  https://github.com/daomor/precious/releases/download/v0.1-alpha/precious
mv ./precious /usr/local/bin
chmod +x /usr/local/bin/precious