#!/usr/bin/env sh

if [[ $# -eq 0 ]] ; then
    echo 'no tag provide!'
    exit 1
fi

git push
git tag -d $1
git push -d origin $1
git tag $1
git push origin $1