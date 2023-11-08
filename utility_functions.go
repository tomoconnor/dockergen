package main

import (
	"fmt"
	"strings"
)

// This function takes a command line string, and returns it as a string like [ "command", "arg1", "arg2", ... ] format
func commandLineToDockerFormattedCmd(commandLine string) string {
	// Split the command line into a slice of strings
	commandLineSlice := strings.Split(commandLine, " ")
	// Add double quotes around each element of the slice
	for i, v := range commandLineSlice {
		commandLineSlice[i] = fmt.Sprintf("\"%s\"", v)
	}

	// Join the slice back into a string
	inner := strings.Join(commandLineSlice, ", ")
	// trim trailing comma space
	inner = strings.TrimSuffix(inner, ", ")
	// add square brackets
	result := fmt.Sprintf("[ %s ]", inner)
	return result
}
