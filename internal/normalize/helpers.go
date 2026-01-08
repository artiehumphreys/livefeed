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
	float32ParseBits = 64
)

func clampU64(lo, hi, val uint64) uint64 {
	return min(max(lo, val), hi)
}

func clampF64(lo, hi, val float64) float64 {
	return min(max(lo, val), hi)
}

func stoui(s string, bitSize int) uint64 {
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
	return uint8(clampU64(0, math.MaxUint8, stoui(s, uint8ParseBits)))
}

func stou16(s string) uint16 {
	return uint16(clampU64(0, math.MaxUint16, stoui(s, uint16ParseBits)))
}

func stof32(s string) float32 {
	return float32(clampF64(0, math.MaxFloat32, stof(s, float32ParseBits)))
}
