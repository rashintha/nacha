# NACHA Generator

![Go](https://img.shields.io/badge/go-badge?style=for-the-badge&logo=go&logoColor=white&color=%2300ADD8)

## Introduction

NACHA Generator is a Go library designed to facilitate the generation of NACHA (National Automated Clearing House
Association) files. NACHA files are used for processing ACH (Automated Clearing House) transactions, which are
electronic fund transfers between banks in the United States.

## Features

- Easy-to-use API for NACHA file generation
- Support for File Header and Control Records
- Support for Batch Header and Control Records
- Support for Entry Detail Records and Addenda Records
- Validation of required fields and data formats
- Automatic calculation of control totals and hash values

## Installation

To install the NACHA Generator library, use the following command:

```bash
go get github.com/rashintha/nacha
```

## Basic Usage
This is a basic example of how to create a new NACHA file and add records to it.
```go
package main

import (
	"fmt"
	"time"

	"github.com/rashintha/nacha"
)

func main() {
	// Create a new NACHA file
	file := nacha.NewFile()

	// Set the file header fields
	err := file.Header.SetImmediateDestination("123456789")
	if err != nil {
		panic(err)
	}

	err = file.Header.SetImmediateDestinationName("Destination Bank")
	if err != nil {
		panic(err)
	}

	err = file.Header.SetImmediateOrigin("987654321")
	if err != nil {
		panic(err)
	}

	err = file.Header.SetImmediateOriginName("Origin Bank")
	if err != nil {
		panic(err)
	}

	// Add a batch to the file
	batch := file.NewBatch()

	// Set the batch header fields
	err = batch.Header.SetServiceClassCode(220)
	if err != nil {
		panic(err)
	}

	err = batch.Header.SetCompanyName("ABC Company")
	if err != nil {
		panic(err)
	}

	err = batch.Header.SetCompanyIdentification("1122334455")
	if err != nil {
		panic(err)
	}

	err = batch.Header.SetStandardEntryClassCode("CCD")
	if err != nil {
		panic(err)
	}

	err = batch.Header.SetCompanyEntryDescription("Payroll")
	if err != nil {
		panic(err)
	}

	batch.Header.SetEffectiveEntryDate(time.Now().UTC())

	err = batch.Header.SetODFIIdentification("12345678")
	if err != nil {
		panic(err)
	}

	err = batch.Header.SetBatchNumber(1)
	if err != nil {
		panic(err)
	}

	// Add an entry to the batch
	entry := batch.AddEntry()

	err = entry.SetTransactionCode(27)
	if err != nil {
		panic(err)
	}

	err = entry.SetReceivingDFIIdentification("02120770")
	if err != nil {
		panic(err)
	}

	err = entry.SetCheckDigit("7")
	if err != nil {
		panic(err)
	}

	err = entry.SetDFIAccountNumber("29079117")
	if err != nil {
		panic(err)
	}

	err = entry.SetAmount(1364)
	if err != nil {
		panic(err)
	}

	err = entry.SetIndividualIDNumber("392344")
	if err != nil {
		panic(err)
	}

	err = entry.SetIndividualName("BBC Company")
	if err != nil {
		panic(err)
	}

	err = entry.SetTraceNumber("12345678", 1)
	if err != nil {
		panic(err)
	}

	// Add another entry to the batch
	entry2 := batch.AddEntry()

	err = entry2.SetTransactionCode(27)
	if err != nil {
		panic(err)
	}

	err = entry2.SetReceivingDFIIdentification("02120770")
	if err != nil {
		panic(err)
	}

	err = entry2.SetCheckDigit("7")
	if err != nil {
		panic(err)
	}

	err = entry2.SetDFIAccountNumber("18076850")
	if err != nil {
		panic(err)
	}

	err = entry2.SetAmount(982.50)
	if err != nil {
		panic(err)
	}

	err = entry2.SetIndividualIDNumber("392353")
	if err != nil {
		panic(err)
	}

	err = entry2.SetIndividualName("LMN Company")
	if err != nil {
		panic(err)
	}

	err = entry2.SetTraceNumber("12345678", 2)
	if err != nil {
		panic(err)
	}

	// Add an addenda to the entry
	addenda := entry2.NewAddenda()
	addenda.SetPaymentRelatedInformation("Bill Payment for 2021")

	err = addenda.SetAddendaSequenceNumber(1)
	if err != nil {
		panic(err)
	}

	// Generate the file content
	file.GenerateFile()

	// Print the file content
	fmt.Print(file.String())
}

```

