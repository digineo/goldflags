package goldflags

import (
	"fmt"
	"strconv"
	"time"
)

const banner = `
,____  __       __
|    \|__|-----|__|-----.-----,-----.
|  -  |  |  _  |  |     |  .__|  _  |
|____/|__|___  |__|__|__|_____|_____|
         |_____|  %s`
const extraBannerLine = `
                  %s %s`

var (
	commitID   string // (via -ldflags) Commit ID
	commitUnix string // (via -ldflags) Commit date as Unix timestamp in sec
	buildCount string // (via -ldflags) Number of commits
	buildUnix  string // (via -ldflags) Compile date as Unix timestamp in sec

	commitDate string // computed commit date (formatted string)
	buildDate  string // computed compile date (formatted string)
)

func init() {
	format := "2006-01-02 15:04:05"
	if commitUnix != "" && commitDate == "" {
		if t, err := strconv.ParseInt(commitUnix, 10, 64); err == nil {
			commitDate = time.Unix(t, 0).Format(format)
		}
	}
	if buildUnix != "" && buildDate == "" {
		if t, err := strconv.ParseInt(buildUnix, 10, 64); err == nil {
			buildDate = time.Unix(t, 0).Format(format)
		}
	}
}

// VersionString returns a combined version information string
func VersionString() string {
	if commitDate != "" {
		if commitID != "" {
			return fmt.Sprintf("%s (%s)", commitID, commitDate)
		}
		return fmt.Sprintf("development (%s)", commitDate)
	}
	return fmt.Sprintf("development")
}

// BuildString returns string containing build information (build date and number)
func BuildString() string {
	if buildCount != "" && buildDate != "" {
		return fmt.Sprintf("%s (%s)", buildCount, buildDate)
	}
	return fmt.Sprintf("development")
}

// Banner returns a pretty formatted version information banner
func Banner(appName string) string {
	banner := fmt.Sprintf(banner, appName)
	banner += fmt.Sprintf(extraBannerLine, "Version", VersionString())
	banner += fmt.Sprintf(extraBannerLine, "Build  ", BuildString())
	return banner
}
