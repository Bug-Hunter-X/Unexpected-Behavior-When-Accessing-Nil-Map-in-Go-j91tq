# Unexpected Behavior When Accessing Nil Map in Go

This example demonstrates the behavior of accessing elements in a nil map in Go.  Unlike some languages, Go does not throw an immediate error when you attempt to access a key in a nil map.  Instead, it returns the zero value for the map's value type.

This can lead to unexpected behavior if not handled correctly.  The solution demonstrates safe ways to avoid this potential issue.

## Bug
The original code attempts to access a key in an uninitialized map. The expected behavior might be a panic; however, it silently returns 0.

## Solution
The solution offers ways to handle nil maps effectively: explicitly checking for nil, using the comma ok idiom, or utilizing the `make()` function to initialize the map before accessing its elements.