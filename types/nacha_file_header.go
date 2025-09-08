package types

import "time"

// NachaFileHeader represents the NACHA file header (Type 1)
type NachaFileHeader struct {
	Type         string // Char Count: 1 | Fixed Value: 1
	PriorityCode string // PriorityCode Char Count: 2 | Value Usually: 01

	ImmediateDestination string // ImmediateDestination Char Count: 10 (including leading space if 9-digit routing)
	ImmediateOrigin      string // ImmediateOrigin Char Count: 10 (including leading space if 9-digit id)

	FileCreationDate string // FileCreationDate Char Count: 6 | Format: YYMMDD
	FileCreationTime string // FileCreationTime Char Count: 4 | Format: HHMM
	FileIDModifier   string // FileIDModifier Char Count: 1 |Values: A-Z or 0-9

	RecordSize     string // RecordSize Char Count: 3 | Fixed Value: 094
	BlockingFactor string // BlockingFactor Char Count: 2 | Fixed Value: 10
	FormatCode     string // FormatCode Char Count: 1 | Fixed Value: 1

	ImmediateDestinationName string // ImmediateDestinationName Char Count: 23
	ImmediateOriginName      string // ImmediateOriginName Char Count: 23
	ReferenceCode            string // ReferenceCode Char Count: 8 | Optional
}

func (h *NachaFileHeader) Default() {
	h.Type = "1"
	h.PriorityCode = "01"

	h.FileCreationDate = time.Now().UTC().Format("060102")
	h.FileCreationTime = time.Now().UTC().Format("1504")
	h.FileIDModifier = "A"

	h.RecordSize = "094"
	h.BlockingFactor = "10"
	h.FormatCode = "1"

	h.ReferenceCode = ""
}
