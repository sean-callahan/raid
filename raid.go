package raid

import (
	"fmt"
)

// Level represents a standard RAID level
type Level int

// Standard RAID levels
const (
	Level0  Level = 0
	Level1  Level = 1
	Level4  Level = 4
	Level5  Level = 5
	Level6  Level = 6
	Level10 Level = 10
)

// LevelText provides a string representation of a Level
var LevelText = map[Level]string{
	Level0:  "RAID 0 (stripe set)",
	Level1:  "RAID 1 (mirror)",
	Level4:  "RAID 4 (single parity striped set)",
	Level5:  "RAID 5 (striped parity)",
	Level6:  "RAID 6 (RAID 5 w/ double parity)",
	Level10: "RAID 10 (striped mirrors)",
}

// Minimum number of drives per level
var levelMin = map[Level]uint64{
	Level0:  2,
	Level1:  2,
	Level4:  2,
	Level5:  3,
	Level6:  4,
	Level10: 4,
}

// TotalCapacity computes the total capacity of the drives using this
// RAID level.
func (l Level) TotalCapacity(capacity uint64, drives uint64) (uint64, error) {
	if drives < levelMin[l] {
		return 0, fmt.Errorf("Too few drives: need at least [%d]", levelMin[l])
	}
	switch l {
	case Level0:
		return capacity * drives, nil
	case Level1:
		return capacity, nil
	case Level4:
		return capacity * (drives - 1), nil
	case Level5:
		return capacity * (drives - 1), nil
	case Level6:
		return capacity * (drives - 2), nil
	case Level10:
		return capacity * (drives / 2), nil
	default:
		return 0, fmt.Errorf("Unimplemented RAID level %d", l)
	}
}

// SpaceEfficiency computes the effeciency of the RAID level based on the
// provided drives and their capacity.
func (l Level) SpaceEfficiency(capacity uint64, drives uint64) (float64, error) {
	if drives < levelMin[l] {
		return -1, fmt.Errorf("Too few drives: need at least [%d]", levelMin[l])
	}
	switch l {
	case Level0:
		return 1.0, nil
	case Level1:
		return 1.0 / float64(drives), nil
	case Level4:
		return 1.0 - 1.0/float64(drives), nil
	case Level5:
		return 1.0 - 1.0/float64(drives), nil
	case Level6:
		return 1.0 - 2.0/float64(drives), nil
	case Level10:
		return float64(drives) / 2.0 / float64(drives), nil
	default:
		return -1, fmt.Errorf("Unimplemented RAID level %d", l)
	}
}

// FaultTolerance computes the number of drives that can fail using
// this RAID level.
func (l Level) FaultTolerance(drives uint64) (uint64, error) {
	if drives < levelMin[l] {
		return 0, fmt.Errorf("Too few drives: need at least [%d]", levelMin[l])
	}
	switch l {
	case Level0:
		return 0, nil
	case Level1:
		return drives - 1, nil
	case Level4:
		return 1, nil
	case Level5:
		return 1, nil
	case Level6:
		return 2, nil
	case Level10:
		return drives / 2, nil
	default:
		return 0, fmt.Errorf("Unimplemented RAID level %d", l)
	}
}
