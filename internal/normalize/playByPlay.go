package normalize

import (
	"encoding/json"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func ParsePlayByPlay(data []byte) (*types.RawPlayByPlaySummary, error) {
	var res types.RawPlayByPlaySummary

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func NormalizePlayByPlay(raw *types.RawPlayByPlaySummary) (types.PlayByPlaySummary, error) {

}
