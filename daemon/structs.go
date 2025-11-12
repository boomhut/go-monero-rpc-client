package daemon

// Helper structs
type BlockHeader struct {
	BlockSize                 uint64 `json:"block_size"`
	BlockWeight               uint64 `json:"block_weight"`
	CumulativeDifficulty      uint64 `json:"cumulative_difficulty"`
	CumulativeDifficultyTop64 uint64 `json:"cumulative_difficulty_top64"`
	Depth                     uint64 `json:"depth"`
	Difficulty                uint64 `json:"difficulty"`
	DifficultyTop64           uint64 `json:"difficulty_top64"`
	Hash                      string `json:"hash"`
	Height                    uint64 `json:"height"`
	LongTermWeight            uint64 `json:"long_term_weight"`
	MajorVersion              uint64 `json:"major_version"`
	MinerTxHash               string `json:"miner_tx_hash"`
	MinorVersion              uint64 `json:"minor_version"`
	Nonce                     uint64 `json:"nonce"`
	NumTxes                   uint64 `json:"num_txes"`
	OrphanStatus              bool   `json:"orphan_status"`
	PowHash                   string `json:"pow_hash"`
	PrevHash                  string `json:"prev_hash"`
	Reward                    uint64 `json:"reward"`
	Timestamp                 uint64 `json:"timestamp"`
	WideCumulativeDifficulty  string `json:"wide_cumulative_difficulty"`
	WideDifficulty            string `json:"wide_difficulty"`
}

type Connection struct {
	Address           string `json:"address"`
	AddressType       uint64 `json:"address_type"`
	AvgDownload       uint64 `json:"avg_download"`
	AvgUpload         uint64 `json:"avg_upload"`
	ConnectionID      string `json:"connection_id"`
	CurrentDownload   uint64 `json:"current_download"`
	CurrentUpload     uint64 `json:"current_upload"`
	Height            uint64 `json:"height"`
	Host              string `json:"host"`
	Incoming          bool   `json:"incoming"`
	IP                string `json:"ip"`
	LiveTime          uint64 `json:"live_time"`
	LocalIP           bool   `json:"local_ip"`
	Localhost         bool   `json:"localhost"`
	PeerID            string `json:"peer_id"`
	Port              string `json:"port"`
	PruningSeed       uint64 `json:"pruning_seed"`
	RecvCount         uint64 `json:"recv_count"`
	RecvIdleTime      uint64 `json:"recv_idle_time"`
	RPCCreditsPerHash uint64 `json:"rpc_credits_per_hash"`
	RPCPort           uint64 `json:"rpc_port"`
	SendCount         uint64 `json:"send_count"`
	SendIdleTime      uint64 `json:"send_idle_time"`
	State             string `json:"state"`
	SupportFlags      uint64 `json:"support_flags"`
}

type Ban struct {
	Host    string `json:"host"`
	IP      uint64 `json:"ip"`
	Seconds uint64 `json:"seconds"`
}

type AuxPow struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

type ChainInfo struct {
	BlockHash            string   `json:"block_hash"`
	BlockHashes          []string `json:"block_hashes"`
	Difficulty           uint64   `json:"difficulty"`
	Height               uint64   `json:"height"`
	Length               uint64   `json:"length"`
	MainChainParentBlock string   `json:"main_chain_parent_block"`
	WideDifficulty       string   `json:"wide_difficulty"`
}

type HistogramEntry struct {
	Amount            uint64 `json:"amount"`
	TotalInstances    uint64 `json:"total_instances"`
	UnlockedInstances uint64 `json:"unlocked_instances"`
	RecentInstances   uint64 `json:"recent_instances"`
}

type OutKey struct {
	Height   uint64 `json:"height"`
	Key      string `json:"key"`
	Mask     string `json:"mask"`
	Txid     string `json:"txid"`
	Unlocked bool   `json:"unlocked"`
}

type TxInfo struct {
	AsHex           string   `json:"as_hex"`
	AsJSON          string   `json:"as_json"`
	BlockHeight     uint64   `json:"block_height"`
	BlockTimestamp  uint64   `json:"block_timestamp"`
	DoubleSpendSeen bool     `json:"double_spend_seen"`
	InPool          bool     `json:"in_pool"`
	OutputIndices   []uint64 `json:"output_indices"`
	PrunableAsHex   string   `json:"prunable_as_hex"`
	PrunableHash    string   `json:"prunable_hash"`
	PrunedAsHex     string   `json:"pruned_as_hex"`
	TxHash          string   `json:"tx_hash"`
}

type Peer struct {
	ID                uint64 `json:"id"`
	Host              string `json:"host"`
	IP                uint64 `json:"ip"`
	Port              uint64 `json:"port"`
	RPCPort           uint64 `json:"rpc_port"`
	RPCCreditsPerHash uint64 `json:"rpc_credits_per_hash"`
	LastSeen          uint64 `json:"last_seen"`
	PruningSeed       uint64 `json:"pruning_seed"`
}

type SpanInfo struct {
	ConnectionID     string `json:"connection_id"`
	NBlocks          uint64 `json:"nblocks"`
	Rate             uint64 `json:"rate"`
	RemoteAddress    string `json:"remote_address"`
	Size             uint64 `json:"size"`
	Speed            uint64 `json:"speed"`
	StartBlockHeight uint64 `json:"start_block_height"`
}

type PeerInfo struct {
	Info  Connection `json:"info"`
	Spans []SpanInfo `json:"spans,omitempty"`
}

// JSON-RPC Methods

// GetBlockCount
type ResponseGetBlockCount struct {
	Count     uint64 `json:"count"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// OnGetBlockHash
type RequestOnGetBlockHash []uint64

type ResponseOnGetBlockHash string

// GetBlockTemplate
type RequestGetBlockTemplate struct {
	WalletAddress string `json:"wallet_address"`
	ReserveSize   uint64 `json:"reserve_size"`
	PrevBlock     string `json:"prev_block,omitempty"`
}

type ResponseGetBlockTemplate struct {
	Blockhashing_blob string `json:"blockhashing_blob"`
	BlocktemplateBlob string `json:"blocktemplate_blob"`
	Difficulty        uint64 `json:"difficulty"`
	DifficultyTop64   uint64 `json:"difficulty_top64"`
	ExpectedReward    uint64 `json:"expected_reward"`
	Height            uint64 `json:"height"`
	NextSeedHash      string `json:"next_seed_hash"`
	PrevHash          string `json:"prev_hash"`
	ReservedOffset    uint64 `json:"reserved_offset"`
	SeedHash          string `json:"seed_hash"`
	SeedHeight        uint64 `json:"seed_height"`
	Status            string `json:"status"`
	Untrusted         bool   `json:"untrusted"`
	WideDifficulty    string `json:"wide_difficulty"`
}

// SubmitBlock
type RequestSubmitBlock []string

type ResponseSubmitBlock struct {
	Status string `json:"status"`
}

// GenerateBlocks
type RequestGenerateBlocks struct {
	AmountOfBlocks uint64 `json:"amount_of_blocks"`
	WalletAddress  string `json:"wallet_address"`
	StartingNonce  uint64 `json:"starting_nonce,omitempty"`
	PrevBlock      string `json:"prev_block,omitempty"`
}

type ResponseGenerateBlocks struct {
	Blocks    []string `json:"blocks"`
	Height    uint64   `json:"height"`
	Status    string   `json:"status"`
	Untrusted bool     `json:"untrusted"`
}

// GetLastBlockHeader
type ResponseGetLastBlockHeader struct {
	BlockHeader BlockHeader `json:"block_header"`
	Status      string      `json:"status"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeaderByHash
type RequestGetBlockHeaderByHash struct {
	Hash        string `json:"hash"`
	FillPowHash bool   `json:"fill_pow_hash,omitempty"`
}

type ResponseGetBlockHeaderByHash struct {
	BlockHeader BlockHeader `json:"block_header"`
	Status      string      `json:"status"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeaderByHeight
type RequestGetBlockHeaderByHeight struct {
	Height      uint64 `json:"height"`
	FillPowHash bool   `json:"fill_pow_hash,omitempty"`
}

type ResponseGetBlockHeaderByHeight struct {
	BlockHeader BlockHeader `json:"block_header"`
	Status      string      `json:"status"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeadersRange
type RequestGetBlockHeadersRange struct {
	StartHeight uint64 `json:"start_height"`
	EndHeight   uint64 `json:"end_height"`
	FillPowHash bool   `json:"fill_pow_hash,omitempty"`
}

type ResponseGetBlockHeadersRange struct {
	Headers   []BlockHeader `json:"headers"`
	Status    string        `json:"status"`
	Untrusted bool          `json:"untrusted"`
}

// GetBlock
type RequestGetBlock struct {
	Hash        string `json:"hash,omitempty"`
	Height      uint64 `json:"height,omitempty"`
	FillPowHash bool   `json:"fill_pow_hash,omitempty"`
}

type ResponseGetBlock struct {
	Blob        string      `json:"blob"`
	BlockHeader BlockHeader `json:"block_header"`
	JSON        string      `json:"json"`
	MinerTxHash string      `json:"miner_tx_hash"`
	Status      string      `json:"status"`
	TxHashes    []string    `json:"tx_hashes"`
	Untrusted   bool        `json:"untrusted"`
}

// GetConnections
type ResponseGetConnections struct {
	Connections []Connection `json:"connections"`
	Status      string       `json:"status"`
	Untrusted   bool         `json:"untrusted"`
}

// GetInfo
type ResponseGetInfo struct {
	AdjustedTime              uint64 `json:"adjusted_time"`
	AltBlocksCount            uint64 `json:"alt_blocks_count"`
	BlockSizeLimit            uint64 `json:"block_size_limit"`
	BlockSizeMedian           uint64 `json:"block_size_median"`
	BlockWeightLimit          uint64 `json:"block_weight_limit"`
	BlockWeightMedian         uint64 `json:"block_weight_median"`
	BootstrapDaemonAddress    string `json:"bootstrap_daemon_address"`
	BusySyncing               bool   `json:"busy_syncing"`
	Cumulativeifficulty       uint64 `json:"cumulative_difficulty"`
	CumulativeDifficultyTop64 uint64 `json:"cumulative_difficulty_top64"`
	DatabaseSize              uint64 `json:"database_size"`
	Difficulty                uint64 `json:"difficulty"`
	DifficultyTop64           uint64 `json:"difficulty_top64"`
	FreeSpace                 uint64 `json:"free_space"`
	GreyPeerlistSize          uint64 `json:"grey_peerlist_size"`
	Height                    uint64 `json:"height"`
	HeightWithoutBootstrap    uint64 `json:"height_without_bootstrap"`
	IncomingConnectionsCount  uint64 `json:"incoming_connections_count"`
	Mainnet                   bool   `json:"mainnet"`
	Nettype                   string `json:"nettype"`
	Offline                   bool   `json:"offline"`
	OutgoingConnectionsCount  uint64 `json:"outgoing_connections_count"`
	RPCConnectionsCount       uint64 `json:"rpc_connections_count"`
	Stagenet                  bool   `json:"stagenet"`
	StartTime                 uint64 `json:"start_time"`
	Status                    string `json:"status"`
	Synchronized              bool   `json:"synchronized"`
	Target                    uint64 `json:"target"`
	TargetHeight              uint64 `json:"target_height"`
	Testnet                   bool   `json:"testnet"`
	TopBlockHash              string `json:"top_block_hash"`
	TxCount                   uint64 `json:"tx_count"`
	TxPoolSize                uint64 `json:"tx_pool_size"`
	Untrusted                 bool   `json:"untrusted"`
	UpdateAvailable           bool   `json:"update_available"`
	Version                   string `json:"version"`
	WasBootstrapEverUsed      bool   `json:"was_bootstrap_ever_used"`
	WhitePeerlistSize         uint64 `json:"white_peerlist_size"`
	WideCumulativeDifficulty  string `json:"wide_cumulative_difficulty"`
	WideDifficulty            string `json:"wide_difficulty"`
}

// HardForkInfo
type ResponseHardForkInfo struct {
	EarliestHeight uint64 `json:"earliest_height"`
	Enabled        bool   `json:"enabled"`
	State          uint64 `json:"state"`
	Status         string `json:"status"`
	Threshold      uint64 `json:"threshold"`
	Untrusted      bool   `json:"untrusted"`
	Version        uint64 `json:"version"`
	Votes          uint64 `json:"votes"`
	Voting         uint64 `json:"voting"`
	Window         uint64 `json:"window"`
}

// BanRequest represents a single ban request
type BanRequest struct {
	Host    string `json:"host,omitempty"`
	IP      uint64 `json:"ip,omitempty"`
	Ban     bool   `json:"ban"`
	Seconds uint64 `json:"seconds"`
}

// SetBans
type RequestSetBans struct {
	Bans []BanRequest `json:"bans"`
}

type ResponseSetBans struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// GetBans
type ResponseGetBans struct {
	Bans      []Ban  `json:"bans"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// Banned
type RequestBanned struct {
	Address string `json:"address"`
}

type ResponseBanned struct {
	Banned  bool   `json:"banned"`
	Seconds uint64 `json:"seconds"`
	Status  string `json:"status"`
}

// FlushTxpool
type RequestFlushTxpool struct {
	TxIDs []string `json:"txids,omitempty"`
}

type ResponseFlushTxpool struct {
	Status string `json:"status"`
}

// GetOutputHistogram
type RequestGetOutputHistogram struct {
	Amounts      []uint64 `json:"amounts"`
	MinCount     uint64   `json:"min_count,omitempty"`
	MaxCount     uint64   `json:"max_count,omitempty"`
	Unlocked     bool     `json:"unlocked,omitempty"`
	RecentCutoff uint64   `json:"recent_cutoff,omitempty"`
}

type ResponseGetOutputHistogram struct {
	Histogram []HistogramEntry `json:"histogram"`
	Status    string           `json:"status"`
	Untrusted bool             `json:"untrusted"`
}

// GetCoinbaseTxSum
type RequestGetCoinbaseTxSum struct {
	Height uint64 `json:"height"`
	Count  uint64 `json:"count"`
}

type ResponseGetCoinbaseTxSum struct {
	EmissionAmount      uint64 `json:"emission_amount"`
	EmissionAmountTop64 uint64 `json:"emission_amount_top64"`
	FeeAmount           uint64 `json:"fee_amount"`
	FeeAmountTop64      uint64 `json:"fee_amount_top64"`
	Status              string `json:"status"`
	Untrusted           bool   `json:"untrusted"`
	WideEmissionAmount  string `json:"wide_emission_amount"`
	WideFeeAmount       string `json:"wide_fee_amount"`
}

// GetVersion
type ResponseGetVersion struct {
	CurrentHeight uint64 `json:"current_height"`
	HardForks     []struct {
		Height    uint64 `json:"height"`
		HFVersion uint64 `json:"hf_version"`
	} `json:"hard_forks"`
	Release   bool   `json:"release"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
	Version   uint64 `json:"version"`
}

// GetFeeEstimate
type RequestGetFeeEstimate struct {
	GraceBlocks uint64 `json:"grace_blocks,omitempty"`
}

type ResponseGetFeeEstimate struct {
	Fee              uint64   `json:"fee"`
	Fees             []uint64 `json:"fees"`
	QuantizationMask uint64   `json:"quantization_mask"`
	Status           string   `json:"status"`
	Untrusted        bool     `json:"untrusted"`
}

// GetAlternateChains
type ResponseGetAlternateChains struct {
	Chains    []ChainInfo `json:"chains"`
	Status    string      `json:"status"`
	Untrusted bool        `json:"untrusted"`
}

// RelayTx
type RequestRelayTx struct {
	TxIDs []string `json:"txids"`
}

type ResponseRelayTx struct {
	Status string `json:"status"`
}

// SyncInfo
type ResponseSyncInfo struct {
	Height                uint64     `json:"height"`
	NextNeededPruningSeed uint64     `json:"next_needed_pruning_seed"`
	Overview              string     `json:"overview"`
	Peers                 []PeerInfo `json:"peers"`
	Spans                 []SpanInfo `json:"spans,omitempty"`
	Status                string     `json:"status"`
	TargetHeight          uint64     `json:"target_height"`
	Untrusted             bool       `json:"untrusted"`
}

// GetTxpoolBacklog
type ResponseGetTxpoolBacklog struct {
	Backlog   []TxPoolBacklogEntry `json:"backlog"`
	Status    string               `json:"status"`
	Untrusted bool                 `json:"untrusted"`
}

// GetOutputDistribution
type RequestGetOutputDistribution struct {
	Amounts    []uint64 `json:"amounts"`
	Cumulative bool     `json:"cumulative,omitempty"`
	FromHeight uint64   `json:"from_height,omitempty"`
	ToHeight   uint64   `json:"to_height,omitempty"`
	Binary     bool     `json:"binary,omitempty"`
	Compress   bool     `json:"compress,omitempty"`
}

type ResponseGetOutputDistribution struct {
	Distributions []struct {
		Amount       uint64   `json:"amount"`
		Base         uint64   `json:"base"`
		Distribution []uint64 `json:"distribution"`
		StartHeight  uint64   `json:"start_height"`
	} `json:"distributions"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// GetMinerData
type ResponseGetMinerData struct {
	AlreadyGeneratedCoins uint64 `json:"already_generated_coins"`
	Difficulty            string `json:"difficulty"`
	Height                uint64 `json:"height"`
	MajorVersion          uint64 `json:"major_version"`
	MedianWeight          uint64 `json:"median_weight"`
	PrevID                string `json:"prev_id"`
	SeedHash              string `json:"seed_hash"`
	Status                string `json:"status"`
	TxBacklog             []struct {
		Fee    uint64 `json:"fee"`
		ID     string `json:"id"`
		Weight uint64 `json:"weight"`
	} `json:"tx_backlog"`
	Untrusted bool `json:"untrusted"`
}

// PruneBlockchain
type RequestPruneBlockchain struct {
	Check bool `json:"check,omitempty"`
}

type ResponsePruneBlockchain struct {
	Pruned      bool   `json:"pruned"`
	PruningSeed uint64 `json:"pruning_seed"`
	Status      string `json:"status"`
	Untrusted   bool   `json:"untrusted"`
}

// CalcPow
type RequestCalcPow struct {
	MajorVersion uint64 `json:"major_version"`
	Height       uint64 `json:"height"`
	BlockBlob    string `json:"block_blob"`
	SeedHash     string `json:"seed_hash"`
}

type ResponseCalcPow struct {
	Result string `json:"result"`
}

// FlushCache
type RequestFlushCache struct {
	BadTxs    bool `json:"bad_txs,omitempty"`
	BadBlocks bool `json:"bad_blocks,omitempty"`
}

type ResponseFlushCache struct {
	Status string `json:"status"`
}

// AddAuxPow
type RequestAddAuxPow struct {
	BlocktemplateBlob string   `json:"blocktemplate_blob"`
	AuxPow            []AuxPow `json:"aux_pow"`
}

type ResponseAddAuxPow struct {
	AuxPow            []AuxPow `json:"aux_pow"`
	BlockhashingBlob  string   `json:"blockhashing_blob"`
	BlocktemplateBlob string   `json:"blocktemplate_blob"`
	MerkleRoot        string   `json:"merkle_root"`
	MerkleTreeDepth   uint64   `json:"merkle_tree_depth"`
	Status            string   `json:"status"`
	Untrusted         bool     `json:"untrusted"`
}

// Other RPC Methods (non JSON-RPC)

// GetHeight
type ResponseGetHeight struct {
	Hash      string `json:"hash"`
	Height    uint64 `json:"height"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// GetTransactions
type RequestGetTransactions struct {
	TxHashes     []string `json:"txs_hashes"`
	DecodeAsJSON bool     `json:"decode_as_json,omitempty"`
	Prune        bool     `json:"prune,omitempty"`
	Split        bool     `json:"split,omitempty"`
}

type ResponseGetTransactions struct {
	MissedTx  []string `json:"missed_tx,omitempty"`
	Status    string   `json:"status"`
	Txs       []TxInfo `json:"txs"`
	TxsAsHex  []string `json:"txs_as_hex,omitempty"`
	TxsAsJSON []string `json:"txs_as_json,omitempty"`
	Untrusted bool     `json:"untrusted"`
}

// GetAltBlocksHashes
type ResponseGetAltBlocksHashes struct {
	BlkHashes []string `json:"blks_hashes"`
	Status    string   `json:"status"`
	Untrusted bool     `json:"untrusted"`
}

// IsKeyImageSpent
type RequestIsKeyImageSpent struct {
	KeyImages []string `json:"key_images"`
}

type ResponseIsKeyImageSpent struct {
	SpentStatus []uint64 `json:"spent_status"`
	Status      string   `json:"status"`
	Untrusted   bool     `json:"untrusted"`
}

// SendRawTransaction
type RequestSendRawTransaction struct {
	TxAsHex    string `json:"tx_as_hex"`
	DoNotRelay bool   `json:"do_not_relay,omitempty"`
}

type ResponseSendRawTransaction struct {
	DoubleSpend       bool   `json:"double_spend"`
	FeeTooLow         bool   `json:"fee_too_low"`
	InvalidInput      bool   `json:"invalid_input"`
	InvalidOutput     bool   `json:"invalid_output"`
	LowMixin          bool   `json:"low_mixin"`
	NotRct            bool   `json:"not_rct"`
	NotRelayed        bool   `json:"not_relayed"`
	Overspend         bool   `json:"overspend"`
	Reason            string `json:"reason"`
	SanityCheckFailed bool   `json:"sanity_check_failed"`
	Status            string `json:"status"`
	TooBig            bool   `json:"too_big"`
	TooFewOutputs     bool   `json:"too_few_outputs"`
	TxExtraTooBig     bool   `json:"tx_extra_too_big"`
	Untrusted         bool   `json:"untrusted"`
}

// StartMining
type RequestStartMining struct {
	DoBackgroundMining bool   `json:"do_background_mining"`
	IgnoreBattery      bool   `json:"ignore_battery"`
	MinerAddress       string `json:"miner_address"`
	ThreadsCount       uint64 `json:"threads_count"`
}

type ResponseStartMining struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// StopMining
type ResponseStopMining struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// MiningStatus
type ResponseMiningStatus struct {
	Active                    bool   `json:"active"`
	Address                   string `json:"address"`
	BGIdleThreshold           uint64 `json:"bg_idle_threshold"`
	BGIgnoreBattery           bool   `json:"bg_ignore_battery"`
	BGMinIdleSeconds          uint64 `json:"bg_min_idle_seconds"`
	BGTarget                  uint64 `json:"bg_target"`
	BlockReward               uint64 `json:"block_reward"`
	BlockTarget               uint64 `json:"block_target"`
	Difficulty                uint64 `json:"difficulty"`
	DifficultyTop64           uint64 `json:"difficulty_top64"`
	IsBackgroundMiningEnabled bool   `json:"is_background_mining_enabled"`
	PowAlgorithm              string `json:"pow_algorithm"`
	Speed                     uint64 `json:"speed"`
	Status                    string `json:"status"`
	ThreadsCount              uint64 `json:"threads_count"`
	Untrusted                 bool   `json:"untrusted"`
	WideDifficulty            string `json:"wide_difficulty"`
}

// SaveBC
type ResponseSaveBC struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// GetPeerList
type ResponseGetPeerList struct {
	GrayList  []Peer `json:"gray_list"`
	Status    string `json:"status"`
	WhiteList []Peer `json:"white_list"`
}

// GetPublicNodes
type RequestGetPublicNodes struct {
	Gray           bool `json:"gray,omitempty"`
	White          bool `json:"white,omitempty"`
	IncludeBlocked bool `json:"include_blocked,omitempty"`
}

type ResponseGetPublicNodes struct {
	Gray      []Peer `json:"gray"`
	Status    string `json:"status"`
	White     []Peer `json:"white"`
	Untrusted bool   `json:"untrusted"`
}

// SetLogHashRate
type RequestSetLogHashRate struct {
	Visible bool `json:"visible"`
}

type ResponseSetLogHashRate struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// SetLogLevel
type RequestSetLogLevel struct {
	Level uint64 `json:"level"`
}

type ResponseSetLogLevel struct {
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// SetLogCategories
type RequestSetLogCategories struct {
	Categories string `json:"categories,omitempty"`
}

type ResponseSetLogCategories struct {
	Categories string `json:"categories"`
	Status     string `json:"status"`
	Untrusted  bool   `json:"untrusted"`
}

// SetBootstrapDaemon
type RequestSetBootstrapDaemon struct {
	Address  string `json:"address"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
}

type ResponseSetBootstrapDaemon struct {
	Status string `json:"status"`
}

// GetTransactionPool
type ResponseGetTransactionPool struct {
	SpentKeyImages []struct {
		IDHash    string   `json:"id_hash"`
		TxsHashes []string `json:"txs_hashes"`
	} `json:"spent_key_images"`
	Status       string `json:"status"`
	Transactions []struct {
		BlobSize           uint64 `json:"blob_size"`
		DoNotRelay         bool   `json:"do_not_relay"`
		DoubleSpendSeen    bool   `json:"double_spend_seen"`
		Fee                uint64 `json:"fee"`
		IDHash             string `json:"id_hash"`
		KeptByBlock        bool   `json:"kept_by_block"`
		LastFailedHeight   uint64 `json:"last_failed_height"`
		LastFailedIDHash   string `json:"last_failed_id_hash"`
		LastRelayedTime    uint64 `json:"last_relayed_time"`
		MaxUsedBlockHeight uint64 `json:"max_used_block_height"`
		MaxUsedBlockIDHash string `json:"max_used_block_id_hash"`
		ReceiveTime        uint64 `json:"receive_time"`
		RelayedCount       uint64 `json:"relayed"`
		TxBlob             string `json:"tx_blob"`
		TxJSON             string `json:"tx_json"`
		Weight             uint64 `json:"weight"`
	} `json:"transactions"`
	Untrusted bool `json:"untrusted"`
}

// GetTransactionPoolHashes
type ResponseGetTransactionPoolHashes struct {
	Status    string   `json:"status"`
	TxHashes  []string `json:"tx_hashes"`
	Untrusted bool     `json:"untrusted"`
}

// GetTransactionPoolStats
type ResponseGetTransactionPoolStats struct {
	PoolStats struct {
		BytesMax   uint64 `json:"bytes_max"`
		BytesMed   uint64 `json:"bytes_med"`
		BytesMin   uint64 `json:"bytes_min"`
		BytesTotal uint64 `json:"bytes_total"`
		FeeTotal   uint64 `json:"fee_total"`
		Histo      []struct {
			Bytes uint64 `json:"bytes"`
			Txs   uint64 `json:"txs"`
		} `json:"histo"`
		Histo98pc       uint64 `json:"histo_98pc"`
		Num10m          uint64 `json:"num_10m"`
		NumDoubleSpends uint64 `json:"num_double_spends"`
		NumFailing      uint64 `json:"num_failing"`
		NumNotRelayed   uint64 `json:"num_not_relayed"`
		Oldest          uint64 `json:"oldest"`
		TxsTotal        uint64 `json:"txs_total"`
	} `json:"pool_stats"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// StopDaemon
type ResponseStopDaemon struct {
	Status string `json:"status"`
}

// GetLimit
type ResponseGetLimit struct {
	LimitDown uint64 `json:"limit_down"`
	LimitUp   uint64 `json:"limit_up"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// SetLimit
type RequestSetLimit struct {
	LimitDown int64 `json:"limit_down"`
	LimitUp   int64 `json:"limit_up"`
}

type ResponseSetLimit struct {
	LimitDown uint64 `json:"limit_down"`
	LimitUp   uint64 `json:"limit_up"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// OutPeers
type RequestOutPeers struct {
	OutPeers uint64 `json:"out_peers"`
}

type ResponseOutPeers struct {
	OutPeers  uint64 `json:"out_peers"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// InPeers
type RequestInPeers struct {
	InPeers uint64 `json:"in_peers"`
}

type ResponseInPeers struct {
	InPeers   uint64 `json:"in_peers"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

// GetNetStats
type ResponseGetNetStats struct {
	StartTime       uint64 `json:"start_time"`
	Status          string `json:"status"`
	TotalBytesIn    uint64 `json:"total_bytes_in"`
	TotalBytesOut   uint64 `json:"total_bytes_out"`
	TotalPacketsIn  uint64 `json:"total_packets_in"`
	TotalPacketsOut uint64 `json:"total_packets_out"`
	Untrusted       bool   `json:"untrusted"`
}

// OutputIndex represents an output to retrieve
type OutputIndex struct {
	Amount uint64 `json:"amount"`
	Index  uint64 `json:"index"`
}

// GetOuts
type RequestGetOuts struct {
	Outputs []OutputIndex `json:"outputs"`
	GetTxID bool          `json:"get_txid,omitempty"`
}

type ResponseGetOuts struct {
	Outs      []OutKey `json:"outs"`
	Status    string   `json:"status"`
	Untrusted bool     `json:"untrusted"`
}

// Update
type RequestUpdate struct {
	Command string `json:"command"`
	Path    string `json:"path,omitempty"`
}

type ResponseUpdate struct {
	AutoURI   string `json:"auto_uri"`
	Hash      string `json:"hash"`
	Path      string `json:"path"`
	Status    string `json:"status"`
	Update    bool   `json:"update"`
	UserURI   string `json:"user_uri"`
	Version   string `json:"version"`
	Untrusted bool   `json:"untrusted"`
}

// PopBlocks
type RequestPopBlocks struct {
	NBlocks uint64 `json:"nblocks"`
}

type ResponsePopBlocks struct {
	Height    uint64 `json:"height"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}
