package types

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rashintha/nacha/util"
)

type NachaAddenda struct {
	Type                      string // Char Count: 1 | Fixed Value: 7
	AddendaTypeCode           string // Char Count: 2 | Values: 05 - PPD & CCD
	PaymentRelatedInformation string // Char Count: 80 | Optional
	AddendaSequenceNumber     string // Char Count: 4 | Values: 1 - 9999
	EntryDetailSequenceNumber string // Char Count: 7 | Values: Same as Entry Detail Record Sequence Number
}

// Default sets the default values for the NachaAddenda
func (a *NachaAddenda) Default() {
	a.Type = "7"
	a.AddendaTypeCode = "05"
	a.PaymentRelatedInformation = util.ToFixedWidthString("", 80, false)
}

func (a *NachaAddenda) SetType() {
	a.Type = "7"
}

// SetAddendaTypeCode sets the AddendaTypeCode
func (a *NachaAddenda) SetAddendaTypeCode(code string) error {
	if len(code) != 2 {
		return errors.New("AddendaTypeCode must be 2 characters")
	}

	a.AddendaTypeCode = code
	return nil
}

// SetAddendaTypeCodeToDefault sets the AddendaTypeCode to the default value of "05"
func (a *NachaAddenda) SetAddendaTypeCodeToDefault() {
	a.AddendaTypeCode = "05"
}

// SetPaymentRelatedInformation sets the PaymentRelatedInformation
func (a *NachaAddenda) SetPaymentRelatedInformation(info string) {
	a.PaymentRelatedInformation = util.ToFixedWidthString(strings.ToUpper(info), 80, false)
}

// SetAddendaSequenceNumber sets the AddendaSequenceNumber
func (a *NachaAddenda) SetAddendaSequenceNumber(seq int) error {
	if seq < 1 || seq > 9999 {
		return errors.New("AddendaSequenceNumber must be between 1 and 9999")
	}

	a.AddendaSequenceNumber = util.ToFixedWidthZeroString(strconv.Itoa(seq), 4)
	return nil
}

// SetEntryDetailSequenceNumber sets the EntryDetailSequenceNumber
func (a *NachaAddenda) SetEntryDetailSequenceNumber(seq string) error {
	if seq == "" {
		return errors.New("EntryDetailSequenceNumber cannot be empty")
	}
	if len(seq) > 7 {
		return errors.New("EntryDetailSequenceNumber must be 7 characters or less")
	}
	a.EntryDetailSequenceNumber = util.ToFixedWidthZeroString(seq, 7)
	return nil
}
