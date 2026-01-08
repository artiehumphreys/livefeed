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
