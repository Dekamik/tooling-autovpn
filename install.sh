#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

apt update
apt install -y terraform ansible openvpn
ln -s ./autovpn /usr/bin/autovpn
