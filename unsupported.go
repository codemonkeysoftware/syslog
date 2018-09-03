// +build windows nacl plan9

package syslog

import (
	"errors"
	"log"
	"os"
	"time"
)

// NewLogger returns a logger wrapping a new fallback syslog writer
func NewLogger(p Priority, logFlag int) (*log.Logger, error) {
	s, err := New(p, "")
	if err != nil {
		return nil, err
	}

	return log.New(s, "", logFlag), nil
}

type Writer struct {
	priority Priority
	tag      string
}

// New creates a new fallback writer that will output to the default logger
// Each write to the returned writer outputs a log message with the given
// priority (severity only, facility is ignored) and
// prefix tag. If tag is empty, the os.Args[0] is used
func New(priority Priority, tag string) (*Writer, error) {
	if priority < 0 || priority > LOG_LOCAL7|LOG_DEBUG {
		return nil, errors.New("log/syslog: invalid priority")
	}

	if tag == "" {
		tag = os.Args[0]
	}

	return &Writer{
		priority: priority,
		tag:      tag,
	}, nil
}

// Dial is included for compatibility and is the same as calling New
func Dial(_, _ string, priority Priority, tag string) (*Writer, error) {
	return New(priority, tag)
}

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error) {
	return w.write(w.priority, string(b))
}

// Close closes a connection to the syslog daemon.
func (w *Writer) Close() error {
	return nil
}

// Emerg logs a message with severity LOG_EMERG, ignoring the severity
// passed to New.
func (w *Writer) Emerg(m string) error {
	_, err := w.write(LOG_EMERG, m)
	return err
}

// Alert logs a message with severity LOG_ALERT, ignoring the severity
// passed to New.
func (w *Writer) Alert(m string) error {
	_, err := w.write(LOG_ALERT, m)
	return err
}

// Crit logs a message with severity LOG_CRIT, ignoring the severity
// passed to New.
func (w *Writer) Crit(m string) error {
	_, err := w.write(LOG_CRIT, m)
	return err
}

// Err logs a message with severity LOG_ERR, ignoring the severity
// passed to New.
func (w *Writer) Err(m string) error {
	_, err := w.write(LOG_ERR, m)
	return err
}

// Warning logs a message with severity LOG_WARNING, ignoring the
// severity passed to New.
func (w *Writer) Warning(m string) error {
	_, err := w.write(LOG_WARNING, m)
	return err
}

// Notice logs a message with severity LOG_NOTICE, ignoring the
// severity passed to New.
func (w *Writer) Notice(m string) error {
	_, err := w.write(LOG_NOTICE, m)
	return err
}

// Info logs a message with severity LOG_INFO, ignoring the severity
// passed to New.
func (w *Writer) Info(m string) error {
	_, err := w.write(LOG_INFO, m)
	return err
}

// Debug logs a message with severity LOG_DEBUG, ignoring the severity
// passed to New.
func (w *Writer) Debug(m string) error {
	_, err := w.write(LOG_DEBUG, m)
	return err
}

func (w *Writer) write(p Priority, msg string) (int, error) {
	severity := parseSeverity(p)
	timestamp := time.Now().Format(time.Stamp)
	log.Printf("<%s>%s %s[%d]: %s", severity, timestamp, w.tag, os.Getpid(), msg)
	return 0, nil
}
