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

	for i := range b.Entries {
		entriesAddendaCount += len(b.Entries[i].Addenda)

		RDFINumber, _ := strconv.ParseInt(b.Entries[i].ReceivingDFIIdentification, 10, 64)
		entryHash += RDFINumber

		if b.Entries[i].TransactionCode == "27" || b.Entries[i].TransactionCode == "28" || b.Entries[i].TransactionCode == "37" || b.Entries[i].TransactionCode == "38" {
			debitAmount, _ := strconv.ParseInt(b.Entries[i].Amount, 10, 64)
			totalDebits += debitAmount
		}

		if b.Entries[i].TransactionCode == "22" || b.Entries[i].TransactionCode == "23" || b.Entries[i].TransactionCode == "32" || b.Entries[i].TransactionCode == "33" {
			creditAmount, _ := strconv.ParseInt(b.Entries[i].Amount, 10, 64)
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
