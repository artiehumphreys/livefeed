package normalize

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseBoxScore(data []byte) (*types.RawBoxScore, error) {
	var res types.RawBoxScore
	// short initialization to reduce scope of error, neat pattern
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func normalizeBoxScore(raw *types.RawBoxScore) (*types.BoxScore, error) {
	contestStr := strings.TrimSpace(raw.ContestID)
	contestID, err := strconv.ParseInt(contestStr, 10, 32)

	if err != nil {
		return nil, fmt.Errorf("invalid contest ID %q: %w", contestStr, err)
	}

	clock := parseClock(raw)

	// TODO: parse team

	return &types.BoxScore{
		ContestID: uint32(contestID),
		Status:    strings.TrimSpace(raw.Status),
		Period:    strings.TrimSpace(raw.Period),
		Clock:     clock,
	}, nil
}

func parseClock(raw *types.RawBoxScore) *types.GameClock {
	minStr := strings.TrimSpace(raw.Minutes)
	secStr := strings.TrimSpace(raw.Seconds)

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
