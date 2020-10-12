# Code Engine Go SDK Integration tests
# Requires the following env. variables (provdied to TravisCI)
# - CE_API_KEY: IBM Cloud API Key
# - CE_PROJECT_ID: GUID of Code Engine project to target
# - CE_PROJECT_REGION: region for API URL

echo "Running integration tests..."

# Enter example directory
cd example

# Run example, get exit code
output=$(go run example.go)
exitcode=$?

# Print results
if [ $exitcode = 0 ]; then
    echo "Success!"
else
    echo "Integration tests failed with exit code $exitcode"
    echo $output
fi

exit $exitcode