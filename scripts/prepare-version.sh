#!/bin/bash
set -e

# test

echo "Running prepare version script..."

git config --global user.email "$GITHUB_EMAIL"
git config --global user.name "$GITHUB_NAME"

CURRENT_BASE_VERSION=$(grm flow phase "$CHANNEL" --current --print=base)
TAG_NAME=$(grm flow phase "$CHANNEL" --current)
BUILD_VERSION=$TAG_NAME

git checkout "$BRANCH"

echo "version=$BUILD_VERSION" > version_output.txt
echo "base_version=$CURRENT_BASE_VERSION" >> version_output.txt

mkdir -p CHANGELOG
echo "changelog_path=CHANGELOG/$BUILD_VERSION.md" >> version_output.txt
