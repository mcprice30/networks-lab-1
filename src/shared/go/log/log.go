package log

import (
	"errors"
	"fmt"
	"os"
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

var level LogLevel = ERROR

func Level(lvlStr string) error {
	lvl, exists := levelMap[lvlStr]
	if !exists {
		return errors.New("Level does not exist")
	}
	return SetLevel(lvl)
} 

func SetLevel(lvl LogLevel) error {
	if lvl > FATAL || lvl < TRACE {
		return errors.New("Invalid log level")
	}
	level = lvl
	return nil
}

func GetLevel() LogLevel {
	return level
}

func Trace(str string, elems ... interface{}) {
	if level <= TRACE {
		s := fmt.Sprintf(str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Print(s)
		} else {
			fmt.Print(s[:MaxMessageSize-3] + "...")
		}
	}
}

func Info(str string, elems ... interface{}) {
	if level <= INFO {
		s := fmt.Sprintf(str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Print(s)
		} else {
			fmt.Print(s[:MaxMessageSize-3] + "...")
		}
	}
}

func Warn(str string, elems ... interface{}) {
	if level <= WARN {
		s := fmt.Sprintf(str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
	}
}

func Error(str string, elems ... interface{}) {
	if level <= ERROR {
		s := fmt.Sprintf(str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
	}
}

func Fatal(str string, elems ... interface{}) {
	if level <= FATAL {
		s := fmt.Sprintf(str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Fprint(os.Stderr, s)
		} else {
			fmt.Fprint(os.Stderr, s[:MaxMessageSize-3] + "...")
		}
		os.Exit(1)
	}
}
