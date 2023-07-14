# Code Engine Go SDK Integration tests
# Requires the following env. variables (provided to TravisCI)
# - CE_API_KEY: IBM Cloud API Key
# - CE_PROJECT_ID: GUID of Code Engine project to target
# - CE_PROJECT_REGION: region for API URL
# - CE_ACCOUNT_ID: account id of the current user

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