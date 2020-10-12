# Code Engine Go SDK Integration tests
# Requires the following env. variables (provdied to TravisCI)
# - CE_API_KEY: IBM Cloud API Key
# - CE_PROJECT_ID: GUID of Code Engine project to target
# - CE_PROJECT_REGION: region for API URL

echo "Running integration tests..."

# Run example, get exit code
cd example
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
    echo "Intergration test output is incorrect:"
    echo "Expected '$exampleoutput' to contain '$outputcheck'"
    exit 1
fi

echo "Success!"