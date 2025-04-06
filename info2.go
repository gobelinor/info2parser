package info2parser

import (
	"encoding/binary"
	"fmt"
	"time"
	"unicode/utf16"
)

type Info2 struct {
	Header         uint64
	FileSize       uint64
	DeletionTime   int64
	FileNameLength uint32
	OriginalPath   string
}

func FiletimeToTime(filetime int64) time.Time {
	return time.Unix(0, (filetime-116444736000000000)*100)
}

func decodeUTF16LE(b []byte) string {
	u16 := make([]uint16, len(b)/2)
	for i := 0; i < len(u16); i++ {
		u16[i] = binary.LittleEndian.Uint16(b[i*2:])
	}
	for i, v := range u16 {
		if v == 0 {
			u16 = u16[:i]
			break
		}
	}
	return string(utf16.Decode(u16))
}

func Parse(data []byte) (Info2, error) {
	if len(data) < 28 {
		return Info2{}, fmt.Errorf("file invalid")
	}

	header := binary.LittleEndian.Uint64(data[0:8])
	fileSize := binary.LittleEndian.Uint64(data[8:16])
	deletionTimeRaw := int64(binary.LittleEndian.Uint64(data[16:24]))
	fileNameLength := binary.LittleEndian.Uint32(data[24:28])
	originalPath := decodeUTF16LE(data[28:])

	if len(originalPath)+1 != int(fileNameLength) {
		return Info2{}, fmt.Errorf("file invalid")
	}

	return Info2{
		Header:         header,
		FileSize:       fileSize,
		DeletionTime:   deletionTimeRaw,
		FileNameLength: fileNameLength,
		OriginalPath:   originalPath,
	}, nil
}
