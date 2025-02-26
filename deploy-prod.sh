#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

VERSION="$1"
if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

BRANCH_NAME="deploy-prod-$VERSION"

# Sync to head (also checks that there are no staged files)
git pull origin main

# Append version to prod.txt
echo "$(date '+%Y-%m-%d %H:%M:%S'), $VERSION" | cat - deploy-prod.txt > temp.txt && mv temp.txt deploy-prod.txt

git checkout -b "$BRANCH_NAME"

git add deploy-prod.txt
git commit -m "Deploy $VERSION to prod"
git push origin "$BRANCH_NAME"

echo "Merge this PR to push $VERSION to production: https://github.com/yb172/deploydocus/compare/main...$BRANCH_NAME"
