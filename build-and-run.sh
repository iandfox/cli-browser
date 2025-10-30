#!/bin/sh

set -e # exit immediately if a command exits with a nonzero code

delay=1

echo "--- Building browser.go..."

go build browser.go

sleep $delay

# echo "--- Running browser with no args..."

# ./browser

# sleep $delay

echo "--- Running browser against random quote..."

sleep $delay

./browser https://quotes.toscrape.com/random
