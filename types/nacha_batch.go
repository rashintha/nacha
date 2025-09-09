package types

type NachaBatch struct {
	Header  NachaBatchHeader
	Entries []*NachaEntry
}

// AddEntry adds a new NachaEntry to the batch and append it to the batch's Entries'
func (b *NachaBatch) AddEntry() *NachaEntry {
	entry := &NachaEntry{}
	b.Entries = append(b.Entries, entry)
	return entry
}
