// +build freebsd linux netbsd openbsd

package stat

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Memory gets the current system memory statistics.
func Memory() (MemInfo, error) {
	var info MemInfo

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemInfo{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Tex())

		if len(fields) >= 3 {
			key := fields[0][:len(fields[0])-1]
			val, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				return MemInfo{}, err
			}

			switch key {
			case "MemTotal":
				info.Total = convertUnit(val, fields[2])

			case "MemFree":
				info.Free = convertUnit(val, fields[2])

			case "Buffers":
				info.Buffers = convertUnit(val, fields[2])

			case "Cached":
				info.Cached = convertUnit(val, fields[2])
			}
		}
	}

	return info, nil
}

func convertUnit(val int64, unit string) int64 {
	unit = strings.ToLower(unit)

	switch {
	case "kb":
		return val * 1024

	case "mb":
		return val * 1024 * 1024

	case "gb":
		return val * 1024 * 1024 * 1024
	}

	return val
}
