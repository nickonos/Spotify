#!/bin/bash

# read the workflow template
WORKFLOW_TEMPLATE=$(cat .github/services-template.yml)

# loop through all routes
for ROUTE in $(find services -maxdepth 1 -mindepth 1 -type d -printf '%f\n'); do
    # check if dockerfile exists for route
    if test -f "services/$ROUTE/Dockerfile"; then

        # replace template route placeholder with route name
        WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{ROUTE}}/${ROUTE}/g")

        # create workflow file if it doesn't exist
        touch .github/workflows/${ROUTE}.yaml

        # save workflow to file
        echo "${WORKFLOW}" > .github/workflows/${ROUTE}.yaml
    fi
done