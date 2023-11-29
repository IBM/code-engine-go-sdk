#!/bin/bash
##################################################
# Licensed Materials - Property of IBM
# IBM Cloud Code Engine, 5900-AB0
# © Copyright IBM Corp. 2020, 2023
# US Government Users Restricted Rights - Use, duplication or
# disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
##################################################

set -e

# Code Engine Go SDK Integration tests
# Requires the following env. variables (provided to TravisCI)
# - CE_API_KEY: IBM Cloud API Key
# - CE_PROJECT_ID: GUID of Code Engine project to target
# - CE_PROJECT_REGION: region for API URL
# - CE_ACCOUNT_ID: account id of the current user
# - CE_TLS_KEY_FILE_PATH: path to TLS key file
# - CE_TLS_CERT_FILE_PATH: path to TLS crt file

function get_repo {
    if [ ! -d "$apiDirectory" ]; then
        printf "Cloning github.ibm.com/coligo/api...\n"
        git clone https://github.ibm.com/coligo/api.git "$apiDirectory"
    else
        printf "github.ibm.com/coligo/api already cloned, getting latest...\n"
        cd "$apiDirectory"
        if [[ $(git status --porcelain) ]]; then
            printf "Local working tree contains changes... stashing them\n"
            git stash
        fi
        git checkout main
        git pull
        cd "$rootDirectory"
    fi
}

echo ""
echo "----------------------------------"
echo "Getting test dependencies ..."
echo "----------------------------------"
rootDirectory=$(pwd)
apiDirectory=$rootDirectory/api
get_repo
export CE_TLS_KEY_FILE_PATH=$apiDirectory/test/integration/v1beta/domainmappings/tls-files/demohero.key
export CE_TLS_CERT_FILE_PATH=$apiDirectory/test/integration/v1beta/domainmappings/tls-files/demohero.crt

echo ""
echo "----------------------------------"
echo "Running integration v1 tests ..."
echo "----------------------------------"

# Run example, get exit code
cd example

cd v1
exampleoutput=$(go run example.go)
exampleexit=$?
if [ $exampleexit -ne 0 ]; then
    echo "Integration tests failed with exit code $exampleexit"
    echo $exampleoutput
    exit $exampleexit
fi

# Check if output is expected
outputcheck="2 configmaps"
if [[ $exampleoutput != *$outputcheck* ]]; then
    echo ""
    echo "Integration test output is incorrect:"
    echo "Expected '$exampleoutput' to contain '$outputcheck'"
    echo "================================"
    echo "FAILED"
    echo "================================"
    exit 1
fi

echo "$exampleoutput"
echo "Success!"

echo ""
echo "----------------------------------"
echo "Running integration v2 tests ..."
echo "----------------------------------"


# Run example, get exit code
cd ../v2
exampleoutput=$(go run example_v2.go)
exampleexit=$?
if [ $exampleexit -ne 0 ]; then
    echo ""
    echo "Integration tests failed with exit code $exampleexit"
    echo $exampleoutput
    exit $exampleexit
    echo "================================"
    echo "FAILED"
    echo "================================"
    exit 1
fi

echo "$exampleoutput"
echo "Success!"

echo ""
echo "================================"
echo "SUCCEEDED"
echo "================================"