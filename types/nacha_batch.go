package types

import (
	"strconv"

	"github.com/rashintha/nacha/util"
)

type NachaBatch struct {
	Header  NachaBatchHeader
	Entries []*NachaEntry
	Control NachaBatchControl
}

// AddEntry adds a new NachaEntry to the batch and appends it to the batch's Entries'
func (b *NachaBatch) AddEntry() *NachaEntry {
	entry := &NachaEntry{}
	entry.Default()
	b.Entries = append(b.Entries, entry)
	return entry
}

// GenerateBatchControl generates the BatchControl
func (b *NachaBatch) GenerateBatchControl() {
	b.Control.ServiceClassCode = b.Header.ServiceClassCode

	entriesAddendaCount := len(b.Entries)
	entryHash := int64(0)
	totalDebits := int64(0)
	totalCredits := int64(0)

	for _, entry := range b.Entries {
		entriesAddendaCount += len(entry.Addenda)

		RDFINumber, _ := strconv.ParseInt(entry.ReceivingDFIIdentification, 10, 64)
		entryHash += RDFINumber

		if entry.TransactionCode == "27" || entry.TransactionCode == "28" || entry.TransactionCode == "37" || entry.TransactionCode == "38" {
			debitAmount, _ := strconv.ParseInt(entry.Amount, 10, 64)
			totalDebits += debitAmount
		}

		if entry.TransactionCode == "22" || entry.TransactionCode == "23" || entry.TransactionCode == "32" || entry.TransactionCode == "33" {
			creditAmount, _ := strconv.ParseInt(entry.Amount, 10, 64)
			totalCredits += creditAmount
		}
	}

	b.Control.EntryAddendaCount = util.ToFixedWidthZeroString(strconv.Itoa(entriesAddendaCount), 6)
	b.Control.EntryHash = util.ToFixedWidthZeroString(strconv.Itoa(int(entryHash)), 10)
	b.Control.TotalDebits = util.ToFixedWidthZeroString(strconv.Itoa(int(totalDebits)), 12)
	b.Control.TotalCredits = util.ToFixedWidthZeroString(strconv.Itoa(int(totalCredits)), 12)

	b.Control.CompanyIdentification = b.Header.CompanyIdentification
	b.Control.ODFIIdentification = b.Header.ODFIIdentification
	b.Control.BatchNumber = b.Header.BatchNumber
}
