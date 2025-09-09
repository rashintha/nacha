package types

import (
	"errors"
	"strconv"

	"github.com/rashintha/nacha/util"
)

// NachaFileControl represents the NACHA file control (Type 9)
type NachaFileControl struct {
	Type              string // Char Count: 1 | Fixed Value: 9
	BatchCount        string // Char Count: 6 | Values: 1 - 999999
	BlockCount        string // Char Count: 6 | Values: 1 - 999999
	EntryAddendaCount string // Char Count: 8 | Values: 1 - 99999999
	EntryHash         string // Char Count: 10 | Value: Total of all entry hashes in batch control records

	TotalDebits  string // Char Count: 12 | Value: Total of all debits in batch control records
	TotalCredits string // Char Count: 12 | Value: Total of all credits in batch control records

	Reserved string // Char Count: 39 | Value: Blank
}

// Default sets the default values for the NachaFileControl
func (f *NachaFileControl) Default() {
	f.Type = "9"
	f.Reserved = util.ToFixedWidthString("", 39, false)
}

// SetType sets the Type to "9"
func (f *NachaFileControl) SetType() {
	f.Type = "9"
}

// SetBatchCount sets the BatchCount
func (f *NachaFileControl) SetBatchCount(count int) error {
	if count < 1 || count > 999999 {
		return errors.New("BatchCount must be between 1 and 999999")
	}

	f.BatchCount = util.ToFixedWidthZeroString(strconv.Itoa(count), 6)
	return nil
}

// SetBlockCount sets the BlockCount
func (f *NachaFileControl) SetBlockCount(count int) error {
	if count < 1 || count > 999999 {
		return errors.New("BlockCount must be between 1 and 999999")
	}

	f.BlockCount = util.ToFixedWidthZeroString(strconv.Itoa(count), 6)
	return nil
}

// SetEntryHash sets the EntryHash
func (f *NachaFileControl) SetEntryHash(hash int) error {
	if hash < 0 {
		return errors.New("EntryHash must be greater than or equal to 0")
	}

	f.EntryHash = util.ToFixedWidthZeroString(strconv.Itoa(hash), 10)
	return nil
}

// SetEntryAddendaCount sets the EntryAddendaCount
func (f *NachaFileControl) SetEntryAddendaCount(count int) error {
	if count < 0 || count > 99999999 {
		return errors.New("EntryAddendaCount must be between 0 and 99999999")
	}

	f.EntryAddendaCount = util.ToFixedWidthZeroString(strconv.Itoa(count), 8)
	return nil
}

// SetTotalDebits sets the TotalDebits
func (f *NachaFileControl) SetTotalDebits(amount float64) error {
	if amount < 0 {
		return errors.New("TotalDebits must be greater than or equal to 0")
	}
	if amount > 9999999999.99 {
		return errors.New("TotalDebits must be less than or equal to 9999999999.99")
	}

	f.TotalDebits = util.ToFixedWidthZeroString(strconv.Itoa(int(amount*100)), 12)
	return nil
}

// SetTotalCredits sets the TotalCredits
func (f *NachaFileControl) SetTotalCredits(amount float64) error {
	if amount < 0 {
		return errors.New("TotalCredits must be greater than or equal to 0")
	}
	if amount > 9999999999.99 {
		return errors.New("TotalCredits must be less than or equal to 9999999999.99")
	}

	f.TotalCredits = util.ToFixedWidthZeroString(strconv.Itoa(int(amount*100)), 12)
	return nil
}
