// +build !windows,!nacl,!plan9

// package syslog wraps the builtin syslog package for supported systems
// and provides fallback logging to the default logger for unsupported systems.
package syslog

import (
	"log"
	gsyslog "log/syslog"
)

func NewLogger(p Priority, logFlag int) (*log.Logger, error) { return gsyslog.NewLogger(p, logFlag) }

func New(priority Priority, tag string) (*Writer, error)

func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)

type Priority = gsyslog.Priority

const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG   = gsyslog.LOG_EMERG
	LOG_ALERT   = gsyslog.LOG_ALERT
	LOG_CRIT    = gsyslog.LOG_CRIT
	LOG_ERR     = gsyslog.LOG_ERR
	LOG_WARNING = gsyslog.LOG_WARNING
	LOG_NOTICE  = gsyslog.LOG_NOTICE
	LOG_INFO    = gsyslog.LOG_INFO
	LOG_DEBUG   = gsyslog.LOG_DEBUG
)

const (

	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN     = gsyslog.LOG_KERN
	LOG_USER     = gsyslog.LOG_USER
	LOG_MAIL     = gsyslog.LOG_MAIL
	LOG_DAEMON   = gsyslog.LOG_DAEMON
	LOG_AUTH     = gsyslog.LOG_AUTH
	LOG_SYSLOG   = gsyslog.LOG_SYSLOG
	LOG_LPR      = gsyslog.LOG_LPR
	LOG_NEWS     = gsyslog.LOG_NEWS
	LOG_UUCP     = gsyslog.LOG_UUCP
	LOG_CRON     = gsyslog.LOG_CRON
	LOG_AUTHPRIV = gsyslog.LOG_AUTHPRIV
	LOG_FTP      = gsyslog.LOG_FTP
	_            // unused
	_            // unused
	_            // unused
	_            // unused
	LOG_LOCAL0   = gsyslog.LOG_LOCAL0
	LOG_LOCAL1   = gsyslog.LOG_LOCAL1
	LOG_LOCAL2   = gsyslog.LOG_LOCAL2
	LOG_LOCAL3   = gsyslog.LOG_LOCAL3
	LOG_LOCAL4   = gsyslog.LOG_LOCAL4
	LOG_LOCAL5   = gsyslog.LOG_LOCAL5
	LOG_LOCAL6   = gsyslog.LOG_LOCAL6
	LOG_LOCAL7   = gsyslog.LOG_LOCAL7
)

type Writer = gsyslog.Writer
