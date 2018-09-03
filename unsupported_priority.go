// +build windows nacl plan9

package syslog

type Priority int

const severityMask = 0x07
const facilityMask = 0xf8

const (
	// Severity

	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (
	// Facility

	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN Priority = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_ // unused
	_ // unused
	_ // unused
	_ // unused
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

func parseSeverity(p Priority) string {
	ps := p & severityMask
	var s string
	switch ps {
	case LOG_EMERG:
		s = "EMERG"
	case LOG_ALERT:
		s = "ALERT"
	case LOG_CRIT:
		s = "CRIT"
	case LOG_ERR:
		s = "ERR"
	case LOG_WARNING:
		s = "WARNING"
	case LOG_NOTICE:
		s = "NOTICE"
	case LOG_INFO:
		s = "INFO"
	case LOG_DEBUG:
		s = "DEBUG"
	default:
		s = ""
	}
	return s
}
