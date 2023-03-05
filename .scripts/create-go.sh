#!/bin/bash

scriptDir=$(dirname $(readlink -fn $0))
templateDir=$(dirname $scriptDir)/.templates/go

templates=( $(ls -1 $templateDir) )

if [[ ! $# -eq 1 ]]; then
    echo "please provide a project template like '${templates[0]}'"
    echo ""
    echo "this is a list of possible go templates:"
    for template in ${templates[@]}; do
        echo " * $template"
    done
    exit 1
fi

chosenTemplate=$1

if [[ ! " ${templates[@]} " =~ " $chosenTemplate " ]]; then
    echo "template '$chosenTemplate' does not exist"
    echo ""
    echo "this is a list of possible go templates:"
    for template in ${templates[@]}; do
        echo " * $template"
    done
    exit 1
fi

if [[ ! $1 =~ ^[a-z\-]*$ ]]; then
    echo "project discription should only contain 'a-z' or '-'"
    exit 1
fi

root=$(dirname $(dirname $(readlink -fn $0)))
dir="$(date +%Y-%m-%d)-$1"
projectDir="$root/$dir-go"

echo "creating new project dir:"
echo "  $projectDir"

mkdir -p "$projectDir"
cp -r "$templateDir/$chosenTemplate/." "$projectDir"

go mod init "coding-puzzles/$chosenTemplate"
echo "done!"
echo ""
echo "run: 'cd $projectDir'"
echo ""
