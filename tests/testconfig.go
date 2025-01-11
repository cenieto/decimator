package test

// testToleranceRelative defines the relative tolerance used for comparing floating-point numbers in tests.
// It represents the maximum allowed relative difference between the expected and actual values for the test to pass.
// A value of 1e-12 ensures that the comparison is precise up to 12 decimal places.
const testToleranceRelative = 1e-12

// testToleranceAbsolute defines the absolute tolerance used for comparing floating-point numbers in tests.
// It represents the maximum allowed absolute difference between the expected and actual values for the test to pass.
// A value of 1e-9 ensures that the comparison allows small discrepancies in the least significant digits, useful for large numbers.
const testToleranceAbsolute = 1e-9