package rfc5424

const severityMask = 0x07
const facilityMask = 0xf8

type Severity int

const (
	DefaultSeverity = iota
	Emergency       = iota
	Alert           = iota
	Critical        = iota
	Error           = iota
	Warning         = iota
	Notice          = iota
	Info            = iota
	Debug           = iota
)

var severityNames = map[string]Severity{
	"emergency":     Emergency,
	"emerg":         Emergency,
	"alert":         Alert,
	"critical":      Critical,
	"crit":          Critical,
	"error":         Error,
	"warning":       Warning,
	"warn":          Warning,
	"notice":        Notice,
	"informational": Info,
	"info":          Info,
	"debug":         Debug,
}

type Facility int

const (
	DefaultFacility = iota
	Kernel          = iota
	User            = iota
	Mail            = iota
	Daemon          = iota
	Auth            = iota
	Syslog          = iota
	LPR             = iota
	News            = iota
	UUCP            = iota
	Clock           = iota
	AuthPriv        = iota
	FTP             = iota
	NTP             = iota
	Audit           = iota
	LogAlert        = iota
	Cron            = iota
	Local0          = iota
	Local1          = iota
	Local2          = iota
	Local3          = iota
	Local4          = iota
	Local5          = iota
	Local6          = iota
	Local7          = iota
)

var facilityNames = map[string]Facility{
	"kernel":   Kernel,
	"user":     User,
	"mail":     Mail,
	"daemon":   Daemon,
	"auth":     Auth,
	"syslog":   Syslog,
	"lpr":      LPR,
	"news":     News,
	"uucp":     UUCP,
	"clock":    Clock,
	"authpriv": AuthPriv,
	"ftp":      FTP,
	"ntp":      NTP,
	"audit":    Audit,
	"logalert": LogAlert,
	"cron":     Cron,
	"local0":   Local0,
	"local1":   Local1,
	"local2":   Local2,
	"local3":   Local3,
	"local4":   Local4,
	"local5":   Local5,
	"local6":   Local6,
	"local7":   Local7,
}
