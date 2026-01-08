package normalize

import (
	"strconv"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseClock(mins string, secs string) *types.GameClock {
	minStr := strings.TrimSpace(mins)
	secStr := strings.TrimSpace(secs)

	if minStr == "" && secStr == "" {
		return nil
	}

	// minutes and seconds should fit into an 8bit integer,
	// default to 0 on error
	minutes, err := strconv.ParseUint(minStr, 10, 8)

	if err != nil {
		minutes = 0
	}

	seconds, err := strconv.ParseUint(secStr, 10, 8)

	if err != nil || seconds > 59 {
		seconds = 0
	}

	return &types.GameClock{
		Minutes: uint8(minutes),
		Seconds: uint8(seconds),
	}
}
