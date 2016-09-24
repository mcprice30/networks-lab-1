// Lab 1 - Group 11
// 
// The log package includes a basic logger to be used by the other elements of
// the project.
//
package log

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	TRACE LogLevel = 0
	INFO LogLevel = 1
	WARN  LogLevel = 2 
	ERROR LogLevel = 3
	FATAL LogLevel = 4
)

var levelMap map[string]LogLevel = map[string]LogLevel {
	"FATAL": FATAL,
	"ERROR": ERROR,
	"WARN": WARN,
	"INFO": INFO,
	"TRACE": TRACE,
}

var MaxMessageSize = 250

// What level to log at.
var level LogLevel = ERROR

// Used in formatting the timestamp for each log message.
const fmtStr string = "01-02-2006 15:04:05"

// Level takes a string and tries to set the log level to be that level.
func Level(lvlStr string) error {
	lvl, exists := levelMap[lvlStr]
	if !exists {
		return errors.New("Level does not exist")
	}
	return SetLevel(lvl)
} 

// SetLevel sets the current logging level to be the given logging level.
func SetLevel(lvl LogLevel) error {
	if lvl > FATAL || lvl < TRACE {
		return errors.New("Invalid log level")
	}
	level = lvl
	return nil
}

// GetLevel returns the current logging level.
func GetLevel() LogLevel {
	return level
}

// Trace logs a message at TRACE level.
// It is formatted like sprintf, with a timestamp and stamp at the beginning.
func Trace(str string, elems ... interface{}) {
	if level <= TRACE {
		logPrefix := fmt.Sprintf("[TRACE %s] ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Print(s)
		} else {
			fmt.Print(s[:MaxMessageSize-3] + "...")
		}
	}
}

// Info logs a message at INFO level.
// It is formatted like sprintf, with a timestamp and stamp at the beginning.
func Info(str string, elems ... interface{}) {
	if level <= INFO {
		logPrefix := fmt.Sprintf("[INFO %s]  ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Print(s)
		} else {
			fmt.Print(s[:MaxMessageSize-3] + "...")
		}
	}
}

// Warn logs a message at WARN level.
// It is formatted like sprintf, with a timestamp and stamp at the beginning.
func Warn(str string, elems ... interface{}) {
	if level <= WARN {
		logPrefix := fmt.Sprintf("[WARN %s]  ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
	}
}

// Error logs a message at ERROR level.
// It is formatted like sprintf, with a timestamp and stamp at the beginning.
func Error(str string, elems ... interface{}) {
	if level <= ERROR {
		logPrefix := fmt.Sprintf("[ERROR %s] ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
	}
}

// Fatal logs a message at FATAL level, then terminates the program with error.
// It is formatted like sprintf, with a timestamp and stamp at the beginning.
func Fatal(str string, elems ... interface{}) {
	if level <= FATAL {
		logPrefix := fmt.Sprintf("[FATAL %s] ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
		os.Exit(1)
	}
}
