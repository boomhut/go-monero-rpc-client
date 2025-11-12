package daemon

// TxPoolBacklogEntry represents transaction pool backlog data
type TxPoolBacklogEntry struct {
	BlobSize   uint64 `json:"blob_size"`
	Fee        uint64 `json:"fee"`
	TimeInPool uint64 `json:"time_in_pool"`
}
