#!/bin/bash

DIST_PATH="./dist" # Define the path where you want the "dist" file to be generated.

# Ensure the "dist" directory exists, and remove the previous "dist" file if it exists.
mkdir -p "$DIST_PATH"
rm -f "$DIST_PATH/dist"

# Call "go build" with the appropriate build constraint to set the output file path.
go build -o "$DIST_PATH/gocli" -tags=dist

echo "Build completed. The dist file is generated at $DIST_PATH/gocli."
