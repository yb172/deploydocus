#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

VERSION="$1"
if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

# Check for staged changes in git
if ! git diff --quiet --cached; then
  echo "Error: There are staged changes in git. Commit or stash them before running this script."
  exit 1
fi

# Append version to prod.txt
echo "$(date +%Y-%m-%d), $VERSION" >> prod.txt

git checkout -b "push-prod-$VERSION"

git add prod.txt
git commit -m "push $VERSION to prod"
git push origin "push-prod-$VERSION"

echo "Merge this PR to push $VERSION to production: https://github.com/yb172/deploydocus/compare/main...push-prod-$VERSION"
