package core

import (
	"fmt"
	"strings"
)

type LogLevel string

var LogLevels = struct {
	Silence LogLevel
	Verbose LogLevel
	Info    LogLevel
	Warn    LogLevel
	Error   LogLevel
	Fatal   LogLevel
}{
	Silence: "silence",
	Verbose: "debug",
	Info:    "info",
	Warn:    "warn",
	Error:   "error",
	Fatal:   "fatal",
}

func (c LogLevel) String() string {
	switch c {
	case "silence":
		return "Silence"
	case "debug":
		return "Verbose"
	case "info":
		return "Info"
	case "warn":
		return "Warn"
	case "error":
		return "Error"
	case "fatal":
		return "Fatal"
	default:
		return "Unknown"
	}
}

func LogLevelEnumFromString(enum string) (LogLevel, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Silence"):
		return LogLevels.Silence, nil
	case strings.ToLower("Verbose"):
		return LogLevels.Verbose, nil
	case strings.ToLower("Info"):
		return LogLevels.Info, nil
	case strings.ToLower("Warn"):
		return LogLevels.Warn, nil
	case strings.ToLower("Error"):
		return LogLevels.Error, nil
	case strings.ToLower("Fatal"):
		return LogLevels.Fatal, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := LogLevelEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func LogLevelToArray() []LogLevel {
	return []LogLevel{
		LogLevels.Silence,
		LogLevels.Verbose,
		LogLevels.Info,
		LogLevels.Warn,
		LogLevels.Error,
		LogLevels.Fatal,
	}
}

type LogLevelEnumFlag struct {
	Value LogLevel
}

func (e *LogLevelEnumFlag) Type() string {
	return "LogLevel"
}

func (e *LogLevelEnumFlag) String() string {
	return e.Value.String()
}

func (e *LogLevelEnumFlag) Set(value string) error {
	validEnumValues := LogLevelToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
