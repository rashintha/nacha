package types

import (
	"fmt"
	"math"
	"strconv"

	"github.com/rashintha/nacha/util"
)

// NachaFile represents the NACHA file
type NachaFile struct {
	Header       NachaFileHeader
	Batches      []*NachaBatch
	Control      NachaFileControl
	BlockFillers []*NachaBlockFiller
}

// NewBatch creates a new NachaBatch and appends it to the file's Batches'
func (f *NachaFile) NewBatch() *NachaBatch {
	batch := &NachaBatch{}
	batch.Header.Default()
	batch.Control.Default()
	f.Batches = append(f.Batches, batch)
	return batch
}

// NewBlockFiller creates a new NachaBlockFiller and appends it to the file's BlockFillers'
func (f *NachaFile) NewBlockFiller() {
	filler := &NachaBlockFiller{}
	filler.Default()
	f.BlockFillers = append(f.BlockFillers, filler)
}

// GenerateFileControl generates the FileControl
func (f *NachaFile) GenerateFileControl() {
	f.Control.BatchCount = util.ToFixedWidthZeroString(strconv.Itoa(len(f.Batches)), 6)

	blockCount := float64(2)
	entryAddendaCount := 0
	entryHashTotal := int64(0)
	totalDebits := int64(0)
	totalCredits := int64(0)

	for _, batch := range f.Batches {
		blockCount += float64(2 + len(batch.Entries))
		entryAddendaCount += len(batch.Entries)

		entryHash, _ := strconv.ParseInt(batch.Control.EntryHash, 10, 64)
		entryHashTotal += entryHash

		for _, entry := range batch.Entries {
			blockCount += float64(len(entry.Addenda))
			entryAddendaCount += len(entry.Addenda)
		}

		debitAmount, _ := strconv.ParseInt(batch.Control.TotalDebits, 10, 64)
		totalDebits += debitAmount

		creditAmount, _ := strconv.ParseInt(batch.Control.TotalCredits, 10, 64)
		totalCredits += creditAmount
	}

	f.Control.BlockCount = util.ToFixedWidthZeroString(strconv.Itoa(int(math.Ceil(blockCount/10))), 6)
	f.Control.EntryAddendaCount = util.ToFixedWidthZeroString(strconv.Itoa(entryAddendaCount), 8)
	f.Control.EntryHash = util.ToFixedWidthZeroString(strconv.Itoa(int(entryHashTotal)), 10)
	f.Control.TotalDebits = util.ToFixedWidthZeroString(strconv.Itoa(int(totalDebits)), 12)
	f.Control.TotalCredits = util.ToFixedWidthZeroString(strconv.Itoa(int(totalCredits)), 12)

	for range 10 - (int(blockCount) % 10) {
		f.NewBlockFiller()
	}
}

// String returns the NACHA file as a string
func (f *NachaFile) String() string {
	// File Header
	nachaString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s\n",
		f.Header.Type, f.Header.PriorityCode, f.Header.ImmediateDestination, f.Header.ImmediateOrigin,
		f.Header.FileCreationDate, f.Header.FileCreationTime, f.Header.FileIDModifier,
		f.Header.RecordSize, f.Header.BlockingFactor, f.Header.FormatCode,
		f.Header.ImmediateDestinationName, f.Header.ImmediateOriginName, f.Header.ReferenceCode)

	for _, batch := range f.Batches {
		// Batch Header
		nachaString += fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s\n",
			batch.Header.Type, batch.Header.ServiceClassCode, batch.Header.CompanyName,
			batch.Header.CompanyDiscretionaryData, batch.Header.CompanyIdentification,
			batch.Header.StandardEntryClassCode, batch.Header.CompanyEntryDescription, batch.Header.CompanyDescriptiveDate,
			batch.Header.EffectiveEntryDate, batch.Header.SettlementDateJulian, batch.Header.OriginatorStatusCode,
			batch.Header.ODFIIdentification, batch.Header.BatchNumber)

		for _, entry := range batch.Entries {
			nachaString += fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s\n",
				entry.Type, entry.TransactionCode, entry.ReceivingDFIIdentification, entry.CheckDigit,
				entry.DFIAccountNumber, entry.Amount, entry.IndividualIDNumber, entry.IndividualName,
				entry.DiscretionaryData, entry.AddendaRecordIndicator, entry.TraceNumber)

			for _, addenda := range entry.Addenda {
				nachaString += fmt.Sprintf("%s%s%s%s%s\n",
					addenda.Type, addenda.AddendaTypeCode, addenda.PaymentRelatedInformation,
					addenda.AddendaSequenceNumber, addenda.EntryDetailSequenceNumber)
			}
		}

		nachaString += fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s\n",
			batch.Control.Type, batch.Control.ServiceClassCode, batch.Control.EntryAddendaCount,
			batch.Control.EntryHash, batch.Control.TotalDebits, batch.Control.TotalCredits,
			batch.Control.CompanyIdentification, batch.Control.MessageAuthenticationCode, batch.Control.Reserved,
			batch.Control.ODFIIdentification, batch.Control.BatchNumber)
	}

	// File Control
	nachaString += fmt.Sprintf("%s%s%s%s%s%s%s%s\n",
		f.Control.Type, f.Control.BatchCount, f.Control.BlockCount, f.Control.EntryAddendaCount,
		f.Control.EntryHash, f.Control.TotalDebits, f.Control.TotalCredits, f.Control.Reserved)

	// Block Fillers
	for _, filler := range f.BlockFillers {
		nachaString += fmt.Sprintf("%s\n", filler.Reserved)
	}

	return nachaString
}
