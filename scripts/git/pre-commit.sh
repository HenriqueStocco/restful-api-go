#!/bin/env bash

# Interrompe script se qualquer comando falhar
set -e

# Descobre o diretório onde o script está
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$DIR/../pretty-log.sh"

log "Pre-Commit"