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

# Create and checkout the fixes branch
git checkout -b fixes/$VERSION $VERSION

# Create the cherrypick branch and push it to origin
git checkout -b cherrypick/$VERSION

git push -u origin cherrypick/$VERSION

git checkout fixes/$VERSION

echo "Pull request URL: https://github.com/yb172/deploydocus/compare/cherrypick/$VERSION...fixes/$VERSION"
