package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rashintha/nacha/util"
)

// NachaEntry represents the NACHA Entry (Type 6)
type NachaEntry struct {
	Type string // Char Count: 1 | Fixed Value: 6

	// Char Count: 2 | Values:
	// Checking Accounts - 22 (Credit), 23 (Prenote Credit), 27 (Debits), 28 (Prenote Debit)
	// Savings Accounts - 32 (Credit), 33 (Prenote Credit), 37 (Debits), 38 (Prenote Debit)
	TransactionCode            string
	ReceivingDFIIdentification string // Char Count: 8 | value: First 8 digits of the Receiving DFI Routing Number
	CheckDigit                 string // Char Count: 1 | Value: Last digit of the Receiving DFI Routing Number
	DFIAccountNumber           string // Char Count: 17 | Value: DFI Account Number
	Amount                     string // Char Count: 10 | Value: Amount of the Entry

	IndividualIDNumber string // Char Count: 15 | Value: Individual ID Number (Employee Number etc.)
	IndividualName     string // Char Count: 22 | Value: Individual Name

	DiscretionaryData      string // Char Count: 2 | Optional
	AddendaRecordIndicator string // Char Count: 1 | Value: 0 - No Addenda Record, 1 - Addenda Record
	TraceNumber            string // Char Count: 15 | Value: First 8 digits of the ODFI Routing Number plus Entry Detail Sequence Number
}

// Default sets the default values for the NachaEntry
func (e *NachaEntry) Default() {
	e.Type = "6"
	e.DiscretionaryData = util.ToFixedWidthString("", 2, false)
}

// SetTransactionCode sets the TransactionCode
func (e *NachaEntry) SetTransactionCode(code string) error {
	if code != "22" && code != "23" && code != "27" && code != "28" && code != "32" && code != "33" && code != "37" && code != "38" {
		return errors.New("TransactionCode must be 22, 23, 27, 28, 32, 33, 37, or 38")
	}

	e.TransactionCode = code
	return nil
}

// SetReceivingDFIIdentification sets the ReceivingDFIIdentification
func (e *NachaEntry) SetReceivingDFIIdentification(id string) error {
	if len(id) != 8 {
		return errors.New("ReceivingDFIIdentification must be 8 characters")
	}

	e.ReceivingDFIIdentification = id
	return nil
}

// SetCheckDigit sets the CheckDigit
func (e *NachaEntry) SetCheckDigit(digit string) error {
	if len(digit) != 1 {
		return errors.New("CheckDigit must be 1 character")
	}

	e.CheckDigit = digit
	return nil
}

// SetDFIAccountNumber sets the DFIAccountNumber
func (e *NachaEntry) SetDFIAccountNumber(number string) error {
	if number == "" {
		return errors.New("DFIAccountNumber cannot be empty")
	}

	e.DFIAccountNumber = util.ToFixedWidthString(number, 17, false)
	return nil
}

// SetAmount sets the Amount
func (e *NachaEntry) SetAmount(amount float32) error {
	if amount < 1 {
		return errors.New("Amount must be greater than 0")
	}

	e.Amount = util.ToFixedWidthZeroString(fmt.Sprintf("%d", int(amount*100)), 10)
	return nil
}

// SetIndividualIDNumber sets the IndividualIDNumber
func (e *NachaEntry) SetIndividualIDNumber(id string) error {
	if id == "" {
		return errors.New("IndividualIDNumber cannot be empty")
	}
	if len(id) > 15 {
		return errors.New("IndividualIDNumber must be 15 characters or less")
	}

	e.IndividualIDNumber = util.ToFixedWidthString(id, 15, false)
	return nil
}

// SetIndividualName sets the IndividualName
// If the name is more than 22 characters, it will be truncated
func (e *NachaEntry) SetIndividualName(name string) error {
	if name == "" {
		return errors.New("IndividualName cannot be empty")
	}

	e.IndividualName = strings.ToUpper(util.ToFixedWidthString(name, 22, false))
	return nil
}

// SetDiscretionaryData sets the DiscretionaryData
// If the data is more than 2 characters, it will be truncated
func (e *NachaEntry) SetDiscretionaryData(data string) {
	e.DiscretionaryData = strings.ToUpper(util.ToFixedWidthString(data, 2, false))
}

// SetDiscretionaryDataToDefault sets the DiscretionaryData to the default value of ""
func (e *NachaEntry) SetDiscretionaryDataToDefault() {
	e.DiscretionaryData = util.ToFixedWidthString("", 2, false)
}

// SetAddendaRecordIndicator sets the AddendaRecordIndicator
func (e *NachaEntry) SetAddendaRecordIndicator(indicator string) error {
	if indicator != "0" && indicator != "1" {
		return errors.New("AddendaRecordIndicator must be 0 or 1")
	}

	e.AddendaRecordIndicator = indicator
	return nil
}

func (e *NachaEntry) SetTraceNumber(odfiId string, number string) error {
	if odfiId == "" {
		return errors.New("ODFIId cannot be empty")
	}
	if number == "" {
		return errors.New("TraceNumber cannot be empty")
	}
	if len(number) > 7 {
		return errors.New("TraceNumber must be 7 characters or less")
	}
	if len(odfiId) != 8 {
		return errors.New("ODFIId must be 8 characters")
	}
	e.TraceNumber = odfiId + util.ToFixedWidthZeroString(number, 7)
	return nil
}
