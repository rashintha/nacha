package types

// NachaFile represents the NACHA file
type NachaFile struct {
	Header  NachaFileHeader
	Batches []*NachaBatch
}

// NewBatch creates a new NachaBatch and appends it to the file's Batches'
func (f *NachaFile) NewBatch() *NachaBatch {
	batch := &NachaBatch{}
	batch.Control.Default()
	f.Batches = append(f.Batches, batch)
	return batch
}
