#!/bin/bash

# Check for Go installation
if ! command -v go &> /dev/null
then
    echo "Go could not be found. You can install Go using Homebrew with the following command:"
    echo "brew install go"
    exit 1
fi

# Check for Node installation
if ! command -v node &> /dev/null
then
    echo "Node could not be found. You can install Node using Homebrew with the following command:"
    echo "brew install node"
    exit 1
fi


echo "Setting up your new torque project..."
echo "Please report issues with this script to https://github.com/tylermmorton/create-torque-app/issues/"

# Get the absolute path of the current directory
CURRENT_DIR_ABSOLUTE_PATH=$(cd "$(dirname "$0")"; pwd)

# Remove the '$GOPATH/src/' prefix from the absolute path to get the new package name
NEW_PACKAGE_NAME=${CURRENT_DIR_ABSOLUTE_PATH#$GOPATH/src/}

# The old package name based on the template
OLD_PACKAGE_NAME="github.com/tylermmorton/create-torque-app"

echo "Updating module name in go.mod..."
# Update go.mod
sed -i '' -e "s|$OLD_PACKAGE_NAME|$NEW_PACKAGE_NAME|g" go.mod

echo "Updating import statements in Go files..."
# Update import statements in Go files
find . -name "*.go" -type f | xargs sed -i '' -e "s|$OLD_PACKAGE_NAME|$NEW_PACKAGE_NAME|g"

echo "Updating struct references in template files..."
find . -name "*.tmpl.html" -type f | xargs sed -i '' -e "s|$OLD_PACKAGE_NAME|$NEW_PACKAGE_NAME|g"

# Resolve dependenciese
echo "Tidying up dependencies..."
go mod tidy

# Install task runner
echo "Installing task runner..."
go install github.com/go-task/task/v3/cmd/task@latest

touch .env
mkdir -p .build/assets/
# Delete the one time setup script
rm setup.sh

echo "✅  One time setup complete. Thank you for using the template and trying out the torque framework!"
echo "💾 Source code: https://github.com/tylermmorton/torque"
echo "📖 Documentation: https://lbft.dev/getting-started"
echo "... or jump right in!"
echo "🚀 Start by running 'task dev' from the root directory of your project."
