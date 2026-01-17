package normalize

import (
	"math"
	"strconv"
	"strings"
)

// parse with more bits than destination to avoid premature rejection
const (
	uint8ParseBits   = 16
	uint16ParseBits  = 32
	uint32ParseBits  = 64
	float32ParseBits = 64
)

func clampU64(lo, hi, val uint64) uint64 {
	return min(max(lo, val), hi)
}

func clampI64(lo, hi, val int64) int64 {
	return min(max(lo, val), hi)
}

func clampF64(lo, hi, val float64) float64 {
	return min(max(lo, val), hi)
}

func Stoui(s string, bitSize int) uint64 {
	// in case of bad data, we default to 0
	ans, err := strconv.ParseUint(strings.TrimSpace(s), 10, bitSize)
	if err != nil {
		return 0
	}
	return ans
}

func stof(s string, bitSize int) float64 {
	ans, err := strconv.ParseFloat(strings.TrimSpace(s), bitSize)
	if err != nil {
		return 0
	}
	return ans
}

func stou8(s string) uint8 {
	// parse then clamp to limits of correct type
	return uint8(clampU64(0, math.MaxUint8, Stoui(s, uint8ParseBits)))
}

func stou16(s string) uint16 {
	return uint16(clampU64(0, math.MaxUint16, Stoui(s, uint16ParseBits)))
}

func stou32(s string) uint32 {
	return uint32(clampU64(0, math.MaxUint32, Stoui(s, uint32ParseBits)))
}

func stof32(s string) float32 {
	return float32(clampF64(0, math.MaxFloat32, stof(s, float32ParseBits)))
}

func stofPct(s string) float32 {
	s = strings.TrimSpace(strings.TrimSuffix(s, "%"))
	f, err := strconv.ParseFloat(s, float32ParseBits)
	if err != nil {
		return 0
	}
	return float32(f)
}
