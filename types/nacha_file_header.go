package types

import (
	"errors"
	"strings"
	"time"

	"github.com/rashintha/nacha/util"
)

// NachaFileHeader represents the NACHA file header (Type 1)
type NachaFileHeader struct {
	Type         string // Char Count: 1 | Fixed Value: 1
	PriorityCode string // PriorityCode Char Count: 2 | Value Usually: 01

	ImmediateDestination string // ImmediateDestination Char Count: 10 (including leading space if 9-digit routing)
	ImmediateOrigin      string // ImmediateOrigin Char Count: 10 (including leading space if 9-digit routing)

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

// Default sets the default values for the NachaFileHeader
func (h *NachaFileHeader) Default() {
	h.Type = "1"
	h.PriorityCode = "01"

	h.FileCreationDate = time.Now().UTC().Format("060102")
	h.FileCreationTime = time.Now().UTC().Format("1504")
	h.FileIDModifier = "A"

	h.RecordSize = "094"
	h.BlockingFactor = "10"
	h.FormatCode = "1"

	h.ReferenceCode = util.ToFixedWidthString("", 8, false)
}

// SetType sets the Type to 1
func (h *NachaFileHeader) SetType() {
	h.Type = "1"
}

// SetPriorityCode sets the PriorityCode
func (h *NachaFileHeader) SetPriorityCode(code string) error {
	if len(code) != 2 {
		return errors.New("PriorityCode must be 2 characters")
	}

	h.PriorityCode = code
	return nil
}

// SetImmediateDestination sets the ImmediateDestination routing number
func (h *NachaFileHeader) SetImmediateDestination(routingNumber string) error {
	if routingNumber == "" {
		return errors.New("ImmediateDestination cannot be empty")
	}
	if len(routingNumber) > 10 {
		return errors.New("ImmediateDestination must be 10 characters or less")
	}

	h.ImmediateDestination = util.ToFixedWidthString(routingNumber, 10, true)
	return nil
}

// SetImmediateOrigin sets the ImmediateOrigin routing number
func (h *NachaFileHeader) SetImmediateOrigin(routingNumber string) error {
	if routingNumber == "" {
		return errors.New("ImmediateOrigin cannot be empty")
	}
	if len(routingNumber) > 10 {
		return errors.New("ImmediateOrigin must be 10 characters or less")
	}

	h.ImmediateOrigin = util.ToFixedWidthString(routingNumber, 10, true)
	return nil
}

// SetImmediateDestinationName sets the ImmediateDestinationName.
// If the name is more than 23 characters, it will be truncated
func (h *NachaFileHeader) SetImmediateDestinationName(name string) error {
	if name == "" {
		return errors.New("ImmediateDestinationName cannot be empty")
	}

	h.ImmediateDestinationName = util.ToFixedWidthString(strings.ToUpper(name), 23, false)
	return nil
}

// SetImmediateOriginName sets the ImmediateOriginName
// If the name is more than 23 characters, it will be truncated
func (h *NachaFileHeader) SetImmediateOriginName(name string) error {
	if name == "" {
		return errors.New("ImmediateOriginName cannot be empty")
	}

	h.ImmediateOriginName = util.ToFixedWidthString(strings.ToUpper(name), 23, false)
	return nil
}

// SetFileCreationDate sets the FileCreationDate
func (h *NachaFileHeader) SetFileCreationDate(date time.Time) {
	h.FileCreationDate = date.Format("060102")
}

// SetFileCreationDateToDefault sets the FileCreationDate to the current date (UTC)
func (h *NachaFileHeader) SetFileCreationDateToDefault() {
	h.FileCreationDate = time.Now().UTC().Format("060102")
}

// SetFileCreationTime sets the FileCreationTime
func (h *NachaFileHeader) SetFileCreationTime(time time.Time) {
	h.FileCreationTime = time.Format("1504")
}

// SetFileCreationTimeToDefault sets the FileCreationTime to the current time (UTC)
func (h *NachaFileHeader) SetFileCreationTimeToDefault() {
	h.FileCreationTime = time.Now().UTC().Format("1504")
}

// SetFileIDModifier sets the FileIDModifier
func (h *NachaFileHeader) SetFileIDModifier(modifier string) error {
	if modifier == "" {
		return errors.New("FileIDModifier cannot be empty")
	}
	if len(modifier) != 1 {
		return errors.New("FileIDModifier must be 1 character")
	}

	h.FileIDModifier = modifier
	return nil
}

// SetFileIDModifierToDefault sets the FileIDModifier to the default value of "A"
func (h *NachaFileHeader) SetFileIDModifierToDefault() {
	h.FileIDModifier = "A"
}

// SetRecordSizeToDefault sets the RecordSize to the default value of "094"
func (h *NachaFileHeader) SetRecordSizeToDefault() {
	h.RecordSize = "094"
}

// SetBlockingFactorToDefault sets the BlockingFactor to the default value of "10"
func (h *NachaFileHeader) SetBlockingFactorToDefault() {
	h.BlockingFactor = "10"
}

// SetFormatCodeToDefault sets the FormatCode to the default value of "1"
func (h *NachaFileHeader) SetFormatCodeToDefault() {
	h.FormatCode = "1"
}

// SetReferenceCode sets the ReferenceCode
func (h *NachaFileHeader) SetReferenceCode(code string) error {
	if len(code) > 8 {
		return errors.New("ReferenceCode must be 8 characters or less")
	}

	h.ReferenceCode = util.ToFixedWidthString(code, 8, false)
	return nil
}

// SetReferenceCodeToDefault sets the ReferenceCode to the default value of ""
func (h *NachaFileHeader) SetReferenceCodeToDefault() {
	h.ReferenceCode = util.ToFixedWidthString("", 8, false)
}
