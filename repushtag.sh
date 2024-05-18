#!/usr/bin/env sh

if [[ $# -eq 0 ]] ; then
    echo 'no tag provide!'
    exit 1
fi

git push origin
git push github
git tag -d $1
git push -d origin $1
git tag $1
git push origin $1