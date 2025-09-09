package types

import (
	"errors"
	"strings"
	"time"

	"github.com/rashintha/nacha/util"
)

// NachaBatchHeader represents the NACHA Batch Header (Type 5)
type NachaBatchHeader struct {
	Type             string // Char Count: 1 | Fixed Value: 5
	ServiceClassCode string // Char Count: 3 | Values: 200 - Credits and Debits, 220 - Credits only, 225 - Debits only

	CompanyName              string // Char Count: 16
	CompanyDiscretionaryData string // Char Count: 20 | Optional
	CompanyIdentification    string // Char Count: 10 | Value: Tax ID or Bank Assigned ID

	StandardEntryClassCode string // Char Count: 3 | Values: PPD or CCD

	CompanyEntryDescription string // Char Count: 10 | Values: General identification term (Payroll etc.)
	CompanyDescriptiveDate  string // Char Count: 6 | Format: YYMMDD | Optional

	EffectiveEntryDate   string // Char Count: 6 | Format: YYMMDD
	SettlementDateJulian string // Char Count: 3 | Fixed Value: blank
	OriginatorStatusCode string // Char Count: 1 | Value Usually: 1
	ODFIIdentification   string // Char Count: 8 | Value: First 8 digits of the ODFI Routing Number
	BatchNumber          string // Char Count: 7 | Values: 0000001 - 9999999
}

// Default sets the default values for the NachaBatchHeader
func (h *NachaBatchHeader) Default() {
	h.Type = "5"

	h.CompanyDiscretionaryData = util.ToFixedWidthString("", 20, false)
	h.CompanyDescriptiveDate = util.ToFixedWidthString("", 6, false)
	h.SettlementDateJulian = util.ToFixedWidthString("", 3, false)
	h.OriginatorStatusCode = "1"
}

// SetType sets the Type to "5"
func (h *NachaBatchHeader) SetType() {
	h.Type = "5"
}

// SetServiceClassCode sets the ServiceClassCode
func (h *NachaBatchHeader) SetServiceClassCode(code string) error {
	if code != "200" && code != "220" && code != "225" {
		return errors.New("ServiceClassCode must be 200, 220, or 225")
	}

	h.ServiceClassCode = code
	return nil
}

// SetCompanyName sets the CompanyName
// If the name is more than 16 characters, it will be truncated
func (h *NachaBatchHeader) SetCompanyName(name string) error {
	if name == "" {
		return errors.New("CompanyName cannot be empty")
	}

	h.CompanyName = util.ToFixedWidthString(strings.ToUpper(name), 16, false)
	return nil
}

// SetCompanyDiscretionaryData sets the CompanyDiscretionaryData
func (h *NachaBatchHeader) SetCompanyDiscretionaryData(data string) {
	h.CompanyDiscretionaryData = util.ToFixedWidthString(strings.ToUpper(data), 20, false)
}

// SetCompanyDiscretionaryDataToDefault sets the CompanyDiscretionaryData to the default value of ""
func (h *NachaBatchHeader) SetCompanyDiscretionaryDataToDefault() {
	h.CompanyDiscretionaryData = util.ToFixedWidthString("", 20, false)
}

// SetCompanyIdentification sets the CompanyIdentification
func (h *NachaBatchHeader) SetCompanyIdentification(id string) error {
	if id == "" {
		return errors.New("CompanyIdentification cannot be empty")
	}
	if len(id) > 10 {
		return errors.New("CompanyIdentification must be 10 characters or less")
	}

	h.CompanyIdentification = util.ToFixedWidthString(id, 10, false)
	return nil
}

// SetStandardEntryClassCode sets the StandardEntryClassCode
func (h *NachaBatchHeader) SetStandardEntryClassCode(code string) error {
	if code != "PPD" && code != "CCD" {
		return errors.New("StandardEntryClassCode must be PPD or CCD")
	}

	h.StandardEntryClassCode = code
	return nil
}

// SetCompanyEntryDescription sets the CompanyEntryDescription
func (h *NachaBatchHeader) SetCompanyEntryDescription(description string) error {
	if description == "" {
		return errors.New("CompanyEntryDescription cannot be empty")
	}
	if len(description) > 10 {
		return errors.New("CompanyEntryDescription must be 10 characters or less")
	}

	h.CompanyEntryDescription = util.ToFixedWidthString(strings.ToUpper(description), 10, false)
	return nil
}

// SetCompanyDescriptiveDate sets the CompanyDescriptiveDate
func (h *NachaBatchHeader) SetCompanyDescriptiveDate(date time.Time) {
	h.CompanyDescriptiveDate = date.Format("060102")
}

// SetCompanyDescriptiveDateToDefault sets the CompanyDescriptiveDate to the default value of ""
func (h *NachaBatchHeader) SetCompanyDescriptiveDateToDefault() {
	h.CompanyDescriptiveDate = util.ToFixedWidthString("", 6, false)
}

// SetEffectiveEntryDate sets the EffectiveEntryDate
func (h *NachaBatchHeader) SetEffectiveEntryDate(date time.Time) {
	h.EffectiveEntryDate = date.Format("060102")
}

// SetSettlementDateJulian sets the SettlementDateJulian to the default value of blanks
func (h *NachaBatchHeader) SetSettlementDateJulian() {
	h.SettlementDateJulian = util.ToFixedWidthString("", 3, false)
}

// SetOriginatorStatusCode sets the OriginatorStatusCode
func (h *NachaBatchHeader) SetOriginatorStatusCode(code string) error {
	if len(code) != 1 {
		return errors.New("OriginatorStatusCode must be 1 character")
	}
	h.OriginatorStatusCode = code
	return nil
}

// SetOriginatorStatusCodeToDefault sets the OriginatorStatusCode to the default value of "1"
func (h *NachaBatchHeader) SetOriginatorStatusCodeToDefault() {
	h.OriginatorStatusCode = "1"
}

// SetODFIIdentification sets the ODFIIdentification
func (h *NachaBatchHeader) SetODFIIdentification(id string) error {
	if len(id) != 8 {
		return errors.New("ODFIIdentification must be 8 characters")
	}

	h.ODFIIdentification = id
	return nil
}

// SetBatchNumber sets the BatchNumber
func (h *NachaBatchHeader) SetBatchNumber(number string) error {
	if number == "" {
		return errors.New("BatchNumber cannot be empty")
	}
	if len(number) > 7 {
		return errors.New("BatchNumber must be 7 characters or less")
	}

	h.BatchNumber = util.ToFixedWidthZeroString(number, 7)
	return nil
}
