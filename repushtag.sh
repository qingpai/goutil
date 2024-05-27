#!/usr/bin/env sh

if [[ $# -eq 0 ]] ; then
    echo 'no tag provide!'
    exit 1
fi

git push origin
git tag $1
git push origin $1

export ALL_PROXY=socks5://192.168.0.11:7448
git push github
git push github $1
unset ALL_PROXY