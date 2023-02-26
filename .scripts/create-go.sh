#!/bin/bash

if [[ ! $# -eq 1 ]]; then
    echo "please provide a project description like 'symmetric-difference'"
    exit 1
fi

if [[ ! $1 =~ ^[a-z\-]*$ ]]; then
    echo "project discription should only contain 'a-z' or '-'"
    exit 1
fi

root=$(dirname $(dirname $(readlink -fn $0)))
dir="$(date +%Y-%m-%d)-$1"

echo "creating new project dir:"
echo "  $root/$dir"

mkdir -p "$root/$dir"
cd "$root/$dir"
go mod init "github.com/svenliebig/coding-puzzles/$dir"
echo "done!"
