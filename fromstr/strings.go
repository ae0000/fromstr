package fromstr

import "strconv"

// Int64 returns the int64 out of the string, or if the string is not a
// valid int, it returns the default
func Int64(str string, defaultValue int64) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i
}

// Int returns the int out of the string, or if the string is not a
// valid int, it returns the default
func Int(str string, defaultValue int) int {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}

// Float64 returns the int64 out of the string, or if the string is not a
// valid int, it returns the default
func Float64(str string, defaultValue float64) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultValue
	}
	return i
}
