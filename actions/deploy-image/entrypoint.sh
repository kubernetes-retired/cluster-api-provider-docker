#!/bin/sh -l

echo $GITHUB_WORKSPACE
cd $GITHUB_WORKSPACE || true
./scripts/publish-manager