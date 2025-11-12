package daemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

// Client is the interface to interact with Monero daemon RPC
type Client interface {
	// JSON-RPC Methods
	GetBlockCount() (*ResponseGetBlockCount, error)
	OnGetBlockHash(height uint64) (string, error)
	GetBlockTemplate(walletAddress string, reserveSize uint64) (*ResponseGetBlockTemplate, error)
	SubmitBlock(blockBlob string) (*ResponseSubmitBlock, error)
	GenerateBlocks(amountOfBlocks uint64, walletAddress string) (*ResponseGenerateBlocks, error)
	GetLastBlockHeader() (*ResponseGetLastBlockHeader, error)
	GetBlockHeaderByHash(hash string, fillPowHash bool) (*ResponseGetBlockHeaderByHash, error)
	GetBlockHeaderByHeight(height uint64, fillPowHash bool) (*ResponseGetBlockHeaderByHeight, error)
	GetBlockHeadersRange(startHeight, endHeight uint64, fillPowHash bool) (*ResponseGetBlockHeadersRange, error)
	GetBlock(hashOrHeight interface{}, fillPowHash bool) (*ResponseGetBlock, error)
	GetConnections() (*ResponseGetConnections, error)
	GetInfo() (*ResponseGetInfo, error)
	HardForkInfo() (*ResponseHardForkInfo, error)
	SetBans(bans []BanRequest) (*ResponseSetBans, error)
	GetBans() (*ResponseGetBans, error)
	Banned(address string) (*ResponseBanned, error)
	FlushTxpool(txIDs []string) (*ResponseFlushTxpool, error)
	GetOutputHistogram(amounts []uint64, minCount, maxCount uint64, unlocked bool, recentCutoff uint64) (*ResponseGetOutputHistogram, error)
	GetCoinbaseTxSum(height, count uint64) (*ResponseGetCoinbaseTxSum, error)
	GetVersion() (*ResponseGetVersion, error)
	GetFeeEstimate(graceBlocks uint64) (*ResponseGetFeeEstimate, error)
	GetAlternateChains() (*ResponseGetAlternateChains, error)
	RelayTx(txIDs []string) (*ResponseRelayTx, error)
	SyncInfo() (*ResponseSyncInfo, error)
	GetTxpoolBacklog() (*ResponseGetTxpoolBacklog, error)
	GetOutputDistribution(amounts []uint64, cumulative bool, fromHeight, toHeight uint64) (*ResponseGetOutputDistribution, error)
	GetMinerData() (*ResponseGetMinerData, error)
	PruneBlockchain(check bool) (*ResponsePruneBlockchain, error)
	CalcPow(majorVersion, height uint64, blockBlob, seedHash string) (*ResponseCalcPow, error)
	FlushCache(badTxs, badBlocks bool) (*ResponseFlushCache, error)
	AddAuxPow(blocktemplateBlob string, auxPow []AuxPow) (*ResponseAddAuxPow, error)

	// Other RPC Methods
	GetHeight() (*ResponseGetHeight, error)
	GetTransactions(txHashes []string, decodeAsJSON, prune, split bool) (*ResponseGetTransactions, error)
	GetAltBlocksHashes() (*ResponseGetAltBlocksHashes, error)
	IsKeyImageSpent(keyImages []string) (*ResponseIsKeyImageSpent, error)
	SendRawTransaction(txAsHex string, doNotRelay bool) (*ResponseSendRawTransaction, error)
	StartMining(minerAddress string, threadsCount uint64, doBackgroundMining, ignoreBattery bool) (*ResponseStartMining, error)
	StopMining() (*ResponseStopMining, error)
	MiningStatus() (*ResponseMiningStatus, error)
	SaveBC() (*ResponseSaveBC, error)
	GetPeerList() (*ResponseGetPeerList, error)
	GetPublicNodes(gray, white, includeBlocked bool) (*ResponseGetPublicNodes, error)
	SetLogHashRate(visible bool) (*ResponseSetLogHashRate, error)
	SetLogLevel(level uint64) (*ResponseSetLogLevel, error)
	SetLogCategories(categories string) (*ResponseSetLogCategories, error)
	SetBootstrapDaemon(address, username, password, proxy string) (*ResponseSetBootstrapDaemon, error)
	GetTransactionPool() (*ResponseGetTransactionPool, error)
	GetTransactionPoolHashes() (*ResponseGetTransactionPoolHashes, error)
	GetTransactionPoolStats() (*ResponseGetTransactionPoolStats, error)
	StopDaemon() (*ResponseStopDaemon, error)
	GetLimit() (*ResponseGetLimit, error)
	SetLimit(limitDown, limitUp int64) (*ResponseSetLimit, error)
	OutPeers(outPeers uint64) (*ResponseOutPeers, error)
	InPeers(inPeers uint64) (*ResponseInPeers, error)
	GetNetStats() (*ResponseGetNetStats, error)
	GetOuts(outputs []OutputIndex, getTxID bool) (*ResponseGetOuts, error)
	Update(command, path string) (*ResponseUpdate, error)
	PopBlocks(nBlocks uint64) (*ResponsePopBlocks, error)
}

type client struct {
	config Config
}

// New creates a new Monero daemon RPC client
func New(config Config) Client {
	return &client{
		config: config,
	}
}

// Helper method for JSON-RPC calls
func (c *client) do(method string, req, res interface{}) error {
	message, err := json2.EncodeClientRequest(method, req)
	if err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.config.Address+"/json_rpc", bytes.NewReader(message))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for key, value := range c.config.CustomHeaders {
		httpReq.Header.Set(key, value)
	}

	transport := c.config.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	httpClient := &http.Client{Transport: transport}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if httpRes.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", httpRes.StatusCode, string(body))
	}

	if err := json2.DecodeClientResponse(bytes.NewReader(body), res); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// Helper method for non JSON-RPC calls
func (c *client) doOther(endpoint string, req, res interface{}) error {
	var httpReq *http.Request
	var err error

	if req != nil {
		data, marshalErr := json.Marshal(req)
		if marshalErr != nil {
			return fmt.Errorf("failed to marshal request: %w", marshalErr)
		}
		httpReq, err = http.NewRequest("POST", c.config.Address+endpoint, bytes.NewReader(data))
	} else {
		httpReq, err = http.NewRequest("GET", c.config.Address+endpoint, nil)
	}

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for key, value := range c.config.CustomHeaders {
		httpReq.Header.Set(key, value)
	}

	transport := c.config.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	httpClient := &http.Client{Transport: transport}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if httpRes.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", httpRes.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, res); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// JSON-RPC Method Implementations

func (c *client) GetBlockCount() (*ResponseGetBlockCount, error) {
	var res ResponseGetBlockCount
	if err := c.do("get_block_count", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) OnGetBlockHash(height uint64) (string, error) {
	req := RequestOnGetBlockHash{height}
	var res ResponseOnGetBlockHash
	if err := c.do("on_get_block_hash", req, &res); err != nil {
		return "", err
	}
	return string(res), nil
}

func (c *client) GetBlockTemplate(walletAddress string, reserveSize uint64) (*ResponseGetBlockTemplate, error) {
	req := &RequestGetBlockTemplate{
		WalletAddress: walletAddress,
		ReserveSize:   reserveSize,
	}
	var res ResponseGetBlockTemplate
	if err := c.do("get_block_template", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SubmitBlock(blockBlob string) (*ResponseSubmitBlock, error) {
	req := RequestSubmitBlock{blockBlob}
	var res ResponseSubmitBlock
	if err := c.do("submit_block", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GenerateBlocks(amountOfBlocks uint64, walletAddress string) (*ResponseGenerateBlocks, error) {
	req := &RequestGenerateBlocks{
		AmountOfBlocks: amountOfBlocks,
		WalletAddress:  walletAddress,
	}
	var res ResponseGenerateBlocks
	if err := c.do("generateblocks", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetLastBlockHeader() (*ResponseGetLastBlockHeader, error) {
	var res ResponseGetLastBlockHeader
	if err := c.do("get_last_block_header", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetBlockHeaderByHash(hash string, fillPowHash bool) (*ResponseGetBlockHeaderByHash, error) {
	req := &RequestGetBlockHeaderByHash{
		Hash:        hash,
		FillPowHash: fillPowHash,
	}
	var res ResponseGetBlockHeaderByHash
	if err := c.do("get_block_header_by_hash", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetBlockHeaderByHeight(height uint64, fillPowHash bool) (*ResponseGetBlockHeaderByHeight, error) {
	req := &RequestGetBlockHeaderByHeight{
		Height:      height,
		FillPowHash: fillPowHash,
	}
	var res ResponseGetBlockHeaderByHeight
	if err := c.do("get_block_header_by_height", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetBlockHeadersRange(startHeight, endHeight uint64, fillPowHash bool) (*ResponseGetBlockHeadersRange, error) {
	req := &RequestGetBlockHeadersRange{
		StartHeight: startHeight,
		EndHeight:   endHeight,
		FillPowHash: fillPowHash,
	}
	var res ResponseGetBlockHeadersRange
	if err := c.do("get_block_headers_range", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetBlock(hashOrHeight interface{}, fillPowHash bool) (*ResponseGetBlock, error) {
	req := &RequestGetBlock{
		FillPowHash: fillPowHash,
	}

	switch v := hashOrHeight.(type) {
	case string:
		req.Hash = v
	case uint64:
		req.Height = v
	case int:
		req.Height = uint64(v)
	default:
		return nil, fmt.Errorf("hashOrHeight must be string (hash) or uint64/int (height)")
	}

	var res ResponseGetBlock
	if err := c.do("get_block", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetConnections() (*ResponseGetConnections, error) {
	var res ResponseGetConnections
	if err := c.do("get_connections", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetInfo() (*ResponseGetInfo, error) {
	var res ResponseGetInfo
	if err := c.do("get_info", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) HardForkInfo() (*ResponseHardForkInfo, error) {
	var res ResponseHardForkInfo
	if err := c.do("hard_fork_info", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetBans(bans []BanRequest) (*ResponseSetBans, error) {
	req := &RequestSetBans{Bans: bans}
	var res ResponseSetBans
	if err := c.do("set_bans", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetBans() (*ResponseGetBans, error) {
	var res ResponseGetBans
	if err := c.do("get_bans", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) Banned(address string) (*ResponseBanned, error) {
	req := &RequestBanned{Address: address}
	var res ResponseBanned
	if err := c.do("banned", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) FlushTxpool(txIDs []string) (*ResponseFlushTxpool, error) {
	req := &RequestFlushTxpool{TxIDs: txIDs}
	var res ResponseFlushTxpool
	if err := c.do("flush_txpool", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetOutputHistogram(amounts []uint64, minCount, maxCount uint64, unlocked bool, recentCutoff uint64) (*ResponseGetOutputHistogram, error) {
	req := &RequestGetOutputHistogram{
		Amounts:      amounts,
		MinCount:     minCount,
		MaxCount:     maxCount,
		Unlocked:     unlocked,
		RecentCutoff: recentCutoff,
	}
	var res ResponseGetOutputHistogram
	if err := c.do("get_output_histogram", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetCoinbaseTxSum(height, count uint64) (*ResponseGetCoinbaseTxSum, error) {
	req := &RequestGetCoinbaseTxSum{
		Height: height,
		Count:  count,
	}
	var res ResponseGetCoinbaseTxSum
	if err := c.do("get_coinbase_tx_sum", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetVersion() (*ResponseGetVersion, error) {
	var res ResponseGetVersion
	if err := c.do("get_version", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetFeeEstimate(graceBlocks uint64) (*ResponseGetFeeEstimate, error) {
	req := &RequestGetFeeEstimate{GraceBlocks: graceBlocks}
	var res ResponseGetFeeEstimate
	if err := c.do("get_fee_estimate", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetAlternateChains() (*ResponseGetAlternateChains, error) {
	var res ResponseGetAlternateChains
	if err := c.do("get_alternate_chains", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) RelayTx(txIDs []string) (*ResponseRelayTx, error) {
	req := &RequestRelayTx{TxIDs: txIDs}
	var res ResponseRelayTx
	if err := c.do("relay_tx", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SyncInfo() (*ResponseSyncInfo, error) {
	var res ResponseSyncInfo
	if err := c.do("sync_info", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetTxpoolBacklog() (*ResponseGetTxpoolBacklog, error) {
	var res ResponseGetTxpoolBacklog
	if err := c.do("get_txpool_backlog", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetOutputDistribution(amounts []uint64, cumulative bool, fromHeight, toHeight uint64) (*ResponseGetOutputDistribution, error) {
	req := &RequestGetOutputDistribution{
		Amounts:    amounts,
		Cumulative: cumulative,
		FromHeight: fromHeight,
		ToHeight:   toHeight,
	}
	var res ResponseGetOutputDistribution
	if err := c.do("get_output_distribution", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetMinerData() (*ResponseGetMinerData, error) {
	var res ResponseGetMinerData
	if err := c.do("get_miner_data", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) PruneBlockchain(check bool) (*ResponsePruneBlockchain, error) {
	req := &RequestPruneBlockchain{Check: check}
	var res ResponsePruneBlockchain
	if err := c.do("prune_blockchain", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) CalcPow(majorVersion, height uint64, blockBlob, seedHash string) (*ResponseCalcPow, error) {
	req := &RequestCalcPow{
		MajorVersion: majorVersion,
		Height:       height,
		BlockBlob:    blockBlob,
		SeedHash:     seedHash,
	}
	var res ResponseCalcPow
	if err := c.do("calc_pow", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) FlushCache(badTxs, badBlocks bool) (*ResponseFlushCache, error) {
	req := &RequestFlushCache{
		BadTxs:    badTxs,
		BadBlocks: badBlocks,
	}
	var res ResponseFlushCache
	if err := c.do("flush_cache", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) AddAuxPow(blocktemplateBlob string, auxPow []AuxPow) (*ResponseAddAuxPow, error) {
	req := &RequestAddAuxPow{
		BlocktemplateBlob: blocktemplateBlob,
		AuxPow:            auxPow,
	}
	var res ResponseAddAuxPow
	if err := c.do("add_aux_pow", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Other RPC Method Implementations

func (c *client) GetHeight() (*ResponseGetHeight, error) {
	var res ResponseGetHeight
	if err := c.doOther("/get_height", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetTransactions(txHashes []string, decodeAsJSON, prune, split bool) (*ResponseGetTransactions, error) {
	req := &RequestGetTransactions{
		TxHashes:     txHashes,
		DecodeAsJSON: decodeAsJSON,
		Prune:        prune,
		Split:        split,
	}
	var res ResponseGetTransactions
	if err := c.doOther("/get_transactions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetAltBlocksHashes() (*ResponseGetAltBlocksHashes, error) {
	var res ResponseGetAltBlocksHashes
	if err := c.doOther("/get_alt_blocks_hashes", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) IsKeyImageSpent(keyImages []string) (*ResponseIsKeyImageSpent, error) {
	req := &RequestIsKeyImageSpent{KeyImages: keyImages}
	var res ResponseIsKeyImageSpent
	if err := c.doOther("/is_key_image_spent", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SendRawTransaction(txAsHex string, doNotRelay bool) (*ResponseSendRawTransaction, error) {
	req := &RequestSendRawTransaction{
		TxAsHex:    txAsHex,
		DoNotRelay: doNotRelay,
	}
	var res ResponseSendRawTransaction
	if err := c.doOther("/send_raw_transaction", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) StartMining(minerAddress string, threadsCount uint64, doBackgroundMining, ignoreBattery bool) (*ResponseStartMining, error) {
	req := &RequestStartMining{
		MinerAddress:       minerAddress,
		ThreadsCount:       threadsCount,
		DoBackgroundMining: doBackgroundMining,
		IgnoreBattery:      ignoreBattery,
	}
	var res ResponseStartMining
	if err := c.doOther("/start_mining", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) StopMining() (*ResponseStopMining, error) {
	var res ResponseStopMining
	if err := c.doOther("/stop_mining", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) MiningStatus() (*ResponseMiningStatus, error) {
	var res ResponseMiningStatus
	if err := c.doOther("/mining_status", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SaveBC() (*ResponseSaveBC, error) {
	var res ResponseSaveBC
	if err := c.doOther("/save_bc", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetPeerList() (*ResponseGetPeerList, error) {
	var res ResponseGetPeerList
	if err := c.doOther("/get_peer_list", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetPublicNodes(gray, white, includeBlocked bool) (*ResponseGetPublicNodes, error) {
	req := &RequestGetPublicNodes{
		Gray:           gray,
		White:          white,
		IncludeBlocked: includeBlocked,
	}
	var res ResponseGetPublicNodes
	if err := c.doOther("/get_public_nodes", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetLogHashRate(visible bool) (*ResponseSetLogHashRate, error) {
	req := &RequestSetLogHashRate{Visible: visible}
	var res ResponseSetLogHashRate
	if err := c.doOther("/set_log_hash_rate", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetLogLevel(level uint64) (*ResponseSetLogLevel, error) {
	req := &RequestSetLogLevel{Level: level}
	var res ResponseSetLogLevel
	if err := c.doOther("/set_log_level", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetLogCategories(categories string) (*ResponseSetLogCategories, error) {
	req := &RequestSetLogCategories{Categories: categories}
	var res ResponseSetLogCategories
	if err := c.doOther("/set_log_categories", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetBootstrapDaemon(address, username, password, proxy string) (*ResponseSetBootstrapDaemon, error) {
	req := &RequestSetBootstrapDaemon{
		Address:  address,
		Username: username,
		Password: password,
		Proxy:    proxy,
	}
	var res ResponseSetBootstrapDaemon
	if err := c.doOther("/set_bootstrap_daemon", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetTransactionPool() (*ResponseGetTransactionPool, error) {
	var res ResponseGetTransactionPool
	if err := c.doOther("/get_transaction_pool", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetTransactionPoolHashes() (*ResponseGetTransactionPoolHashes, error) {
	var res ResponseGetTransactionPoolHashes
	if err := c.doOther("/get_transaction_pool_hashes", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetTransactionPoolStats() (*ResponseGetTransactionPoolStats, error) {
	var res ResponseGetTransactionPoolStats
	if err := c.doOther("/get_transaction_pool_stats", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) StopDaemon() (*ResponseStopDaemon, error) {
	var res ResponseStopDaemon
	if err := c.doOther("/stop_daemon", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetLimit() (*ResponseGetLimit, error) {
	var res ResponseGetLimit
	if err := c.doOther("/get_limit", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) SetLimit(limitDown, limitUp int64) (*ResponseSetLimit, error) {
	req := &RequestSetLimit{
		LimitDown: limitDown,
		LimitUp:   limitUp,
	}
	var res ResponseSetLimit
	if err := c.doOther("/set_limit", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) OutPeers(outPeers uint64) (*ResponseOutPeers, error) {
	req := &RequestOutPeers{OutPeers: outPeers}
	var res ResponseOutPeers
	if err := c.doOther("/out_peers", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) InPeers(inPeers uint64) (*ResponseInPeers, error) {
	req := &RequestInPeers{InPeers: inPeers}
	var res ResponseInPeers
	if err := c.doOther("/in_peers", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetNetStats() (*ResponseGetNetStats, error) {
	var res ResponseGetNetStats
	if err := c.doOther("/get_net_stats", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) GetOuts(outputs []OutputIndex, getTxID bool) (*ResponseGetOuts, error) {
	req := &RequestGetOuts{
		Outputs: outputs,
		GetTxID: getTxID,
	}
	var res ResponseGetOuts
	if err := c.doOther("/get_outs", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) Update(command, path string) (*ResponseUpdate, error) {
	req := &RequestUpdate{
		Command: command,
		Path:    path,
	}
	var res ResponseUpdate
	if err := c.doOther("/update", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) PopBlocks(nBlocks uint64) (*ResponsePopBlocks, error) {
	req := &RequestPopBlocks{NBlocks: nBlocks}
	var res ResponsePopBlocks
	if err := c.doOther("/pop_blocks", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
