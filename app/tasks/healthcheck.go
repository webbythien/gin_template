package tasks

import "fmt"

// HealthCheck is a task function used for health checks.
// It takes an integer input, prints a health check message with the input value,
// and returns nil to indicate successful execution.
// Parameters:
//   - input: An integer value representing the health check number.
//
// Returns:
//   - An error (always nil) to indicate successful execution.
func HealthCheck(input int64) error {
	// Print a health check message with the input value.
	fmt.Printf("HealthCheck Number: %d\n", input)

	// Return nil to indicate successful execution.
	return nil
}
