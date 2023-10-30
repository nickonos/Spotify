#!/bin/bash

# read the workflow template
WORKFLOW_TEMPLATE=$(cat .github/workflow-template.yml)
SERVICE_TEMPLATE=$(cat .github/service-template.yml)
DEPLOYMENT_TEMPLATE=$(cat .github/deployment-template.yml)


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

        # replace template route placeholder with route name
        SERVICE=$(echo "${SERVICE_TEMPLATE}" | sed "s/{{ROUTE}}/${ROUTE}/g")

        # create workflow file if it doesn't exist
        touch .github/manifests/${ROUTE}-service.yaml

        # save workflow to file
        echo "${SERVICE}" > .github/manifests/${ROUTE}-service.yaml

        # replace template route placeholder with route name
        DEPLOYMENT=$(echo "${DEPLOYMENT_TEMPLATE}" | sed "s/{{ROUTE}}/${ROUTE}/g")

        # create workflow file if it doesn't exist
        touch .github/manifests/${ROUTE}-deployment.yaml

        # save workflow to file
        echo "${DEPLOYMENT}" > .github/manifests/${ROUTE}-deployment.yaml
    fi
done