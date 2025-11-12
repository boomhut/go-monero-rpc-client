# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0] - 2025-11-12

### Added - Daemon RPC Client
Complete implementation of the Monero Daemon RPC client with 60+ methods:

#### JSON-RPC Methods
- Blockchain queries: `GetBlockCount`, `GetBlock`, `GetBlockHeaderBy{Hash,Height}`, `GetBlockHeadersRange`, `GetLastBlockHeader`
- Network info: `GetInfo`, `GetConnections`, `HardForkInfo`, `GetVersion`, `SyncInfo`
- Mining: `GetBlockTemplate`, `SubmitBlock`, `GenerateBlocks`, `GetMinerData`, `CalcPow`
- Transaction pool: `GetTxpoolBacklog`, `FlushTxpool`, `RelayTx`
- Outputs & fees: `GetOutputHistogram`, `GetOutputDistribution`, `GetFeeEstimate`, `GetCoinbaseTxSum`
- Chain management: `GetAlternateChains`, `PruneBlockchain`, `FlushCache`
- Ban management: `SetBans`, `GetBans`, `Banned`
- Auxiliary PoW: `AddAuxPow` for merged mining support

#### Other RPC Endpoints
- Basic info: `GetHeight`, `GetTransactions`, `GetAltBlocksHashes`
- Key images: `IsKeyImageSpent`
- Broadcasting: `SendRawTransaction`
- Mining control: `StartMining`, `StopMining`, `MiningStatus`
- Peer management: `GetPeerList`, `GetPublicNodes`
- Transaction pool: `GetTransactionPool`, `GetTransactionPoolHashes`, `GetTransactionPoolStats`
- Configuration: `SetLogLevel`, `SetLogCategories`, `SetLogHashRate`, `SetBootstrapDaemon`
- Network limits: `GetLimit`, `SetLimit`, `InPeers`, `OutPeers`, `GetNetStats`
- Advanced: `GetOuts`, `SaveBC`, `Update`, `PopBlocks`, `StopDaemon`

### Added - Wallet RPC Methods

### Added - New RPC Methods
- `SetDaemon()` - Set daemon that the wallet connects to
- `AutoRefresh()` - Set auto-refresh mode for the wallet
- `DescribeTransfer()` - Describe a transaction from unsigned_txset or multisig_txset
- `EditAddressBook()` - Edit an existing address book entry
- `EstimateTxSizeAndWeight()` - Estimate size and weight of a transaction
- `ExchangeMultisigKeys()` - Exchange multisig keys with other participants (N-1/N multisig)
- `Freeze()` - Freeze a single output by key image so it will not be used
- `Frozen()` - Check whether a given output is currently frozen by key image
- `Thaw()` - Thaw a single output by key image so it may be used again
- `ScanTx()` - Scan blockchain for specific transactions
- `SetupBackgroundSync()` - Set up background sync mode
- `StartBackgroundSync()` - Start background sync mode
- `StopBackgroundSync()` - Stop background sync mode
- `GetDefaultFeePriority()` - Get the default fee priority setting

### Enhanced - Existing Methods

#### Transfer Methods
- `RequestTransfer`
  - Added `SubtractFeeFromOutputs []uint64` - Choose which destinations to fund the tx fee from
- `ResponseTransfer`
  - Added `AmountsByDest` struct - Amounts transferred per destination
  - Added `Weight uint64` - Transaction weight metric
  - Added `SpentKeyImages` struct - Key images of spent outputs

#### Balance Methods
- `ResponseGetBalance`
  - Added `TimeToUnlock uint64` - Time in seconds before balance is safe to spend
  - Added `BlocksToUnlock uint64` - Number of blocks before balance is safe to spend
  - Enhanced `PerSubaddress` with `TimeToUnlock` field

#### Payment Methods
- `ResponseGetPayments` and `ResponseGetBulkPayments`
  - Added `Locked bool` field to payment structs - Indicates if output is spendable

#### Transfer Tracking
- `Transfer` struct (used by GetTransfers, GetTransferByTxID)
  - Added `Amounts []uint64` - List of amounts for each destination
  - Added `SubaddrIndices []struct` - List of subaddress indices involved
- `ResponseGetTransferByTxID`
  - Added `Transfers []Transfer` - Array of transfers with the same txid (if multiple found)

#### Split Transfer Methods
- `ResponseTransferSplit`
  - Added `WeightList []uint64` - Metric used to calculate transaction fee
  - Added `SpentKeyImagesList` - Key images of spent outputs for each transaction

### API Compatibility
- Fully compatible with Monero Wallet RPC as documented at https://docs.getmonero.org/rpc-library/wallet-rpc/
- Tested against Monero v0.18.3.x

### Breaking Changes
None - All changes are backwards compatible. New fields are optional and existing code will continue to work without modification.

### Migration Guide
If you want to use the new features:

#### Using Auto-Refresh
```go
// Enable auto-refresh with 30 second interval
err := client.AutoRefresh(&wallet.RequestAutoRefresh{
    Enable: true,
    Period: 30,
})
```

#### Using Output Freezing
```go
// Freeze an output to prevent it from being spent
err := client.Freeze(&wallet.RequestFreeze{
    KeyImage: "your_key_image_here",
})

// Check if output is frozen
resp, err := client.Frozen(&wallet.RequestFrozen{
    KeyImage: "your_key_image_here",
})

// Thaw to allow spending again
err := client.Thaw(&wallet.RequestThaw{
    KeyImage: "your_key_image_here",
})
```

#### Using Enhanced Transfer with Fee Subtraction
```go
// Subtract fees from specific outputs instead of change
resp, err := client.Transfer(&wallet.RequestTransfer{
    Destinations: []*wallet.Destination{
        {Address: "address1", Amount: 1000000000000},
        {Address: "address2", Amount: 2000000000000},
    },
    SubtractFeeFromOutputs: []uint64{0, 1}, // Subtract fees evenly from both outputs
    AccountIndex: 0,
    Priority: wallet.PriorityDefault,
})
```

#### Using Background Sync
```go
// Setup background sync
setupResp, err := client.SetupBackgroundSync(&wallet.RequestSetupBackgroundSync{
    BackgroundSyncType: "reuse-wallet",
    WalletPassword: "your_password",
})

// Start background sync
err = client.StartBackgroundSync()

// Stop background sync
err = client.StopBackgroundSync()
```

## [1.3.0] - Previous Version
- Initial implementation with core wallet RPC methods
