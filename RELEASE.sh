#!/bin/bash
# USAGE: release 1.1.1

DIR=$(dirname $0)
VERSION=$1

git stash
git checkout develop -b release/$VERSION

sed -i -e "s/Version: .*/Version: $VERSION/" $DIR/README.md

git commit -a -m "version $VERSION"
git checkout master
git merge --no-ff release/$VERSION
git tag v$VERSION
git checkout develop
git merge --no-ff release/$VERSION
git branch -d release/$VERSION
git push
git push --tags
