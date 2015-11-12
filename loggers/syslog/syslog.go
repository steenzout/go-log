//
// Copyright 2015 Rakuten Marketing LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package syslog

import (
	"fmt"
	"log/syslog"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"
)

type Logger struct {
	gol.Log
	writer *syslog.Writer
}

// New creates a syslog logger.
func New(network, raddr string, priority syslog.Priority, app string, lfmt gol.LogFormatter) gol.Logger {

	if w, err := syslog.Dial(network, raddr, syslog.LOG_INFO, app); err != nil {
		fmt.Println("syslog.Dial() failed: %s", err)
		return nil
	} else {
		l := &Logger{
			Log:    gol.Log{},
			writer: w,
		}
		l.SetFormatter(lfmt)
		return l
	}
}

// Send process log message.
func (l *Logger) Send(m *gol.LogMessage) (err error) {
	if m == nil {
		fmt.Println("message is nil")
		return
	}
	if l.Formatter() == nil {
		fmt.Println("formatter is nil")
		return fmt.Errorf("log formatter is nil")
	}

	var msg string
	if msg, err = l.Formatter().Format(m); err != nil {
		fmt.Println("message format error")
		return
	}

	var lvl severity.Type
	if lvl, err = m.Severity(); err != nil {
		fmt.Println("severity error")
		return
	}

	fmt.Println("message=", msg)
	switch lvl {
	case severity.Emergency:
		fmt.Println("level=Emergency")
		return l.writer.Emerg(msg)
	case severity.Alert:
		fmt.Println("level=Alert")
		return l.writer.Alert(msg)
	case severity.Critical:
		fmt.Println("level=Critical")
		return l.writer.Crit(msg)
	case severity.Error:
		fmt.Println("level=Error")
		return l.writer.Err(msg)
	case severity.Warning:
		fmt.Println("level=Warning")
		return l.writer.Warning(msg)
	case severity.Notice:
		fmt.Println("level=Notice")
		return l.writer.Notice(msg)
	case severity.Info:
		fmt.Println("level=Info")
		return l.writer.Info(msg)
	case severity.Debug:
		fmt.Println("level=Debug")
		return l.writer.Debug(msg)
	}

	return nil
}

var _ gol.Logger = (*Logger)(nil)