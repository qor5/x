#!/usr/bin/env bash

set -e

# Download the zip file
curl -L https://codeload.github.com/qor5/vuetify-pro-tiptap/zip/refs/heads/release -o vuetify-pro-tiptap.zip

# Clean up the destination directory
rm -rf src/lib/TiptapEditor/lib

# Create temp directory and unzip
mkdir -p vuetify-pro-tiptap
unzip vuetify-pro-tiptap.zip -d vuetify-pro-tiptap

# Create destination directory
mkdir -p src/lib/TiptapEditor/lib

# Move contents from the extracted directory to the destination
mv vuetify-pro-tiptap/vuetify-pro-tiptap-release/* src/lib/TiptapEditor/lib/

# Clean up
rm -rf vuetify-pro-tiptap
rm vuetify-pro-tiptap.zip

echo "Done! The vuetify-pro-tiptap has been updated."
