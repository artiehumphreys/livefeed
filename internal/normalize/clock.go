package normalize

import (
	"math"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseClockFromString(s string) *types.GameClock {
	if s == "" {
		return nil
	}

	i := strings.Index(s, ":")
	if i == -1 {
		return nil
	}

	var minutes, seconds string = s[:i], s[i+1:]

	secs := stou8(seconds)
	if secs > 59 {
		secs = 0
	}

	return &types.GameClock{
		Minutes: stou8(minutes),
		Seconds: secs,
	}
}

func parseClockFromBoxScore(raw *types.RawBoxScore) *types.GameClock {
	if raw.Seconds == nil && raw.Minutes == nil {
		return nil
	}

	var minutes, seconds uint8

	if raw.Minutes != nil {
		if m, err := raw.Minutes.Int64(); err == nil {
			minutes = uint8(clampI64(0, math.MaxUint8, m))
		}
	}

	if raw.Seconds != nil {
		if s, err := raw.Seconds.Int64(); err == nil {
			seconds = uint8(clampI64(0, 59, s))
		}
	}

	return &types.GameClock{
		Minutes: minutes,
		Seconds: seconds,
	}
}
