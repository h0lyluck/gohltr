#!/bin/bash

last_tag=$(git describe --tags --abbrev=0)

clean="${last_tag#v}"

major=$(echo $clean | cut -d. -f1)
minor=$(echo $clean | cut -d. -f2)
patch=$(echo $clean | cut -d. -f3)

patch=$((patch + 1))

new_tag="v$major.$minor.$patch"

git tag $new_tag

git push origin $new_tag


read -p "Tag created: $new_tag, do you want to remove the last tag $last_tag? (y/n) " answer
if [[ $answer == "y" ]]; then
    git tag -d $last_tag
    git push origin :refs/tags/$last_tag
    echo "Tag $last_tag removed."
fi