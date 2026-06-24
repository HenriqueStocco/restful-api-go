#!/bin/env bash

# Interrompe script se qualquer comando falhar
set -e

# Descobre o diretório onde o script está
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$DIR/../pretty-log.sh"

# Variaveis globais
COMMIT_MSG=$(cat "$1")
REGEX="^(feat|refactor|style|docs|chore|fix|ci|test|perf|build):.+"
ERR_MSG="'feat:', 'refactor:', 'style:', 'docs:', 'chore:', 'fix:', 'ci:', 'test:', 'perf:' ou 'build:'"

# Verifica se a mensagem segue os padrões de commit
if ! echo "$COMMIT_MSG" | grep -qE "${REGEX}"; then
 logErr "Erro: A mensagem do commit deve começar com $ERR_MSG."
  exit 1
fi

# Verifica se o commit tem menos de 72 caracteres, que é a padronização
if [ ${#COMMIT_MSG} -gt 72 ]; then
  logErr "Erro: A mensagem de commit deve conter no máximo 72 caracteres."
  exit 1
fi

logSucc "Commit padronizado"