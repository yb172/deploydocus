#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

VERSION="$1"
if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

# Fetch latest tags
git fetch --tags

git checkout main
git pull origin main

# Create and checkout the fix branch with code as of <version>
git checkout -b cherrypick-draft-$VERSION $VERSION

# Create the cherrypick branch and push it to origin
git checkout -b release-cherrypick-$VERSION
git push -u origin release-cherrypick-$VERSION

# Switch back to fix branch to work on fix
git checkout cherrypick-draft-$VERSION

echo "Create cherrypick PR: https://github.com/yb172/deploydocus/compare/release-cherrypick-$VERSION...cherrypick-draft-$VERSION"
