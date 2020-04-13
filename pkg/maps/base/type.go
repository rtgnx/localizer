package base

import (
	"github.com/hashicorp/go-hclog"

	"github.com/the-maldridge/shadow"
)

// Base provides a virtual component that handles the passwd and group
// maps.
type Base struct {
	hclog.Logger

	// The baseDir is the location that files will be read from
	// and written to.  On most systems this should be specified
	// as /etc.
	baseDir string

	// The minUID and minGID specify the numeric lower bound for
	// remote values to be loaded into the system.  These values
	// should be set with a decent amount of headroom above the
	// local namespace on the machine.  A default of 2000 is
	// recommended for both.
	minUID int32
	minGID int32

	// The fallbackShell is a mix between convenience and
	// security.  On a secure system this will be /bin/false or
	// /sbin/nologin, whereas on a convenient system this will be
	// /bin/sh or /bin/bash.  This shell will be substituted in if
	// the shell specified for a user isn't present in the list of
	// shells that are known to the system.
	fallbackShell string

	// This is the list of shells that are permitted on a given
	// host.  This list should normally be populated with the list
	// from /etc/shells.
	shellWhitelist map[string]struct{}

	passwd *shadow.PasswdMap
	group  *shadow.GroupMap
}
