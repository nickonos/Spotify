#!/bin/bash
cd services
# Loop through each subdirectory
for dir in */ ; do
    # Change into the directory
    cd "$dir"
    
    # Run go mod tidy
    go mod tidy -e
    
    # Change back to the parent directory
    cd ..
done