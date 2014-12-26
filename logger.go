// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "fmt"
    "log"
    "os"
    "path"
    "runtime"
    "sync"
)

const (
    // log message level
    LevelInfo = iota
    LevelDebug
    LevelWarning
    LevelError
)

type Logger struct {
    log          *log.Logger
    lock         sync.Mutex
    level        int
    enableCaller bool
    depthCaller  int
}

func NewLogger() *Logger {
    logger := log.New(os.Stderr, "信息：", log.LstdFlags)
    return &Logger{
        log:         logger,
        depthCaller: 2,
    }
}

// Change log level
func (bl *Logger) SetLevel(l int) {
    bl.level = l
}

// Enable caller
func (bl *Logger) EnableCaller(b bool) {
    bl.enableCaller = b
}

// Depth caller
func (bl *Logger) DepthCaller(d int) {
    bl.depthCaller = d
}

// Write log at last
func (bl *Logger) WriteLog(level int, msg string) error {
    if bl.level > level {
        return nil
    }
    output := msg
    if bl.enableCaller {
        _, file, line, ok := runtime.Caller(bl.depthCaller)
        if ok {
            _, filename := path.Split(file)
            output = fmt.Sprintf("[%s:%d] %s", filename, line, msg)
        }
    }
    bl.log.Println(output)
    return nil
}

// Debug log
func (bl *Logger) Debug(format string, v ...interface{}) {
    msg := fmt.Sprintf("[DEBUG]"+format, v...)
    bl.WriteLog(LevelDebug, msg)
}

// Info log
func (bl *Logger) Info(format string, v ...interface{}) {
    f := fmt.Sprintf("[INFO]"+format, v...)
    bl.WriteLog(LevelInfo, f)
}

// Warning log
func (bl *Logger) Warning(format string, v ...interface{}) {
    f := fmt.Sprintf("[WARN]"+format, v...)
    bl.WriteLog(LevelWarning, f)
}

// Error log
func (bl *Logger) Error(format string, v ...interface{}) {
    f := fmt.Sprintf("[Error]"+format, v...)
    bl.WriteLog(LevelError, f)
}
