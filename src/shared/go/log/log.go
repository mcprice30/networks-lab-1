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

var level LogLevel = ERROR
const fmtStr string = "01-02-2006 15:04:05"

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
		logPrefix := fmt.Sprintf("[TRACE %s] ", time.Now().Format(fmtStr))
		s := fmt.Sprintf(logPrefix + str + "\n", elems...)
		if len(s) < MaxMessageSize {
			fmt.Print(s)
		} else {
			fmt.Print(s[:MaxMessageSize-3] + "...")
		}
	}
}

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
