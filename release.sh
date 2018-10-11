#!/usr/bin/env bash

echo "Building Wordpress Sanitizer for Windows..."
GOOS=windows go build -v -o target/windows/wordpress-sanitizer.exe

echo "Building Wordpress Sanitizer for Linux..."
GOOS=linux go build -v -o target/linux/wordpress-sanitizer

echo "Building Wordpress Sanitizer for OSX..."
GOOS=darwin go build -v -o target/osx/wordpress-sanitizer

echo "Packaging Wordpress Sanitizer for Windows..."
zip target/wordpress-sanitizer-windows.zip target/windows/wordpress-sanitizer.exe

echo "Packaging Wordpress Sanitizer for Linux..."
zip target/wordpress-sanitizer-linux.zip target/linux/wordpress-sanitizer

echo "Packaging Wordpress Sanitizer for OSX..."
zip target/wordpress-sanitizer-osx.zip target/osx/wordpress-sanitizer