package types

import (
	"errors"
	"strconv"

	"github.com/rashintha/nacha/util"
)

type NachaBatchControl struct {
	Type              string // Char Count: 1 | Fixed Value: 8
	ServiceClassCode  string // Char Count: 3 | Values: Same as the Batch Header
	EntryAddendaCount string // Char Count: 6 | Values: 1 - 999999
	EntryHash         string // Char Count: 10 | Value: Hash of the Entry Detail Records

	TotalDebits  string // Char Count: 12 | Value: Total of all Debits
	TotalCredits string // Char Count: 12 | Value: Total of all Credits

	CompanyIdentification string // Char Count: 10 | Value: Company ID used in the batch header

	MessageAuthenticationCode string // Char Count: 19 | Value: Blank
	Reserved                  string // Char Count: 6 | Value: Blank

	ODFIIdentification string // Char Count: 8 | Value: First 8 digits of the ODFI Routing Number
	BatchNumber        string // Char Count: 7 | Value: Same as the Batch Header
}

// Default sets the default values for the NachaBatchControl
func (b *NachaBatchControl) Default() {
	b.Type = "8"
	b.MessageAuthenticationCode = util.ToFixedWidthString("", 19, false)
	b.Reserved = util.ToFixedWidthString("", 6, false)
}

// SetType sets the Type to "9"
func (b *NachaBatchControl) SetType() {
	b.Type = "8"
}

// SetServiceClassCode sets the ServiceClassCode
func (b *NachaBatchControl) SetServiceClassCode(code string) error {
	if code != "200" && code != "220" && code != "225" {
		return errors.New("ServiceClassCode must be 200, 220, or 225")
	}

	b.ServiceClassCode = code
	return nil
}

// SetEntryAddendaCount sets the EntryAddendaCount
func (b *NachaBatchControl) SetEntryAddendaCount(count int) error {
	if count < 1 || count > 999999 {
		return errors.New("EntryAddendaCount must be between 1 and 999999")
	}
	b.EntryAddendaCount = util.ToFixedWidthZeroString(strconv.Itoa(count), 6)
	return nil
}

// SetEntryHash sets the EntryHash
func (b *NachaBatchControl) SetEntryHash(hash int64) error {
	if hash < 0 || hash > 9999999999 {
		return errors.New("EntryHash must be between 0 and 9999999999")
	}

	b.EntryHash = util.ToFixedWidthZeroString(strconv.Itoa(int(hash)), 10)
	return nil
}

// SetTotalDebits sets the TotalDebits
func (b *NachaBatchControl) SetTotalDebits(amount float64) error {
	if amount < 0 {
		return errors.New("TotalDebits must be greater than or equal to 0")
	}
	if amount > 99999999999999.99 {
		return errors.New("TotalDebits must be less than or equal to 99999999999999.99")
	}
	b.TotalDebits = util.ToFixedWidthZeroString(strconv.Itoa(int(amount*100)), 12)
	return nil
}

// SetTotalCredits sets the TotalCredits
func (b *NachaBatchControl) SetTotalCredits(amount float64) error {
	if amount < 0 {
		return errors.New("TotalCredits must be greater than or equal to 0")
	}
	if amount > 99999999999999.99 {
		return errors.New("TotalCredits must be less than or equal to 99999999999999.99")
	}
	b.TotalCredits = util.ToFixedWidthZeroString(strconv.Itoa(int(amount*100)), 12)
	return nil
}

// SetCompanyIdentification sets the CompanyIdentification
func (b *NachaBatchControl) SetCompanyIdentification(id string) error {
	if id == "" {
		return errors.New("CompanyIdentification cannot be empty")
	}
	if len(id) > 10 {
		return errors.New("CompanyIdentification must be 10 characters or less")
	}

	b.CompanyIdentification = util.ToFixedWidthString(id, 10, false)
	return nil
}

// SetMessageAuthenticationCode sets the MessageAuthenticationCode to blank
func (b *NachaBatchControl) SetMessageAuthenticationCode() {
	b.MessageAuthenticationCode = util.ToFixedWidthString("", 19, false)
}

// SetReserved sets the Reserved to blank
func (b *NachaBatchControl) SetReserved() {
	b.Reserved = util.ToFixedWidthString("", 6, false)
}

// SetODFIIdentification sets the ODFIIdentification
func (b *NachaBatchControl) SetODFIIdentification(id string) error {
	if id == "" {
		return errors.New("ODFIIdentification cannot be empty")
	}
	if len(id) > 7 {
		return errors.New("ODFIIdentification must be 7 characters or less")
	}

	b.ODFIIdentification = util.ToFixedWidthZeroString(id, 7)
	return nil
}
