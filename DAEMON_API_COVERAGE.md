# Daemon RPC API Coverage

Complete implementation status of Monero Daemon RPC methods as documented at https://docs.getmonero.org/rpc-library/monerod-rpc/

## JSON-RPC Methods (via /json_rpc endpoint)

### Block Operations
- ✅ `get_block_count` - Get the node's current height
- ✅ `on_get_block_hash` - Look up a block's hash by height
- ✅ `get_block_template` - Get a block template for mining
- ✅ `submit_block` - Submit a mined block to the network
- ✅ `generateblocks` - Generate blocks (regtest/testnet only)
- ✅ `get_last_block_header` - Get the block header of the last block
- ✅ `get_block_header_by_hash` - Get block header information by hash
- ✅ `get_block_header_by_height` - Get block header information by height
- ✅ `get_block_headers_range` - Get block headers in a range
- ✅ `get_block` - Get full block information by hash or height

### Network & Node Info
- ✅ `get_connections` - Information about incoming and outgoing connections
- ✅ `get_info` - General information about the state of the node
- ✅ `hard_fork_info` - Information about the current hard fork
- ✅ `get_version` - Daemon version and hard fork schedule
- ✅ `sync_info` - Synchronization status and peer information

### Ban Management
- ✅ `set_bans` - Ban or unban peers
- ✅ `get_bans` - Get list of banned peers
- ✅ `banned` - Check if an address is banned

### Transaction Pool
- ✅ `flush_txpool` - Flush transactions from the transaction pool
- ✅ `get_txpool_backlog` - Get transaction pool backlog
- ✅ `relay_tx` - Relay transactions to other nodes

### Outputs & Fees
- ✅ `get_output_histogram` - Get histogram of output amounts
- ✅ `get_coinbase_tx_sum` - Get the coinbase amount and fee from blocks
- ✅ `get_fee_estimate` - Get current fee per kB
- ✅ `get_output_distribution` - Get output distribution

### Mining
- ✅ `get_miner_data` - Get mining data needed for mining the next block
- ✅ `calc_pow` - Calculate proof-of-work hash for a block
- ✅ `add_aux_pow` - Add auxiliary proof-of-work for merged mining

### Chain Management
- ✅ `get_alternate_chains` - Display alternative chains seen by the node
- ✅ `prune_blockchain` - Prune the blockchain
- ✅ `flush_cache` - Flush bad transactions / blocks from the cache

## Other RPC Methods (Direct HTTP endpoints)

### Basic Information
- ✅ `/get_height` - Get the node's current height
- ✅ `/get_transactions` - Look up transactions by hash
- ✅ `/get_alt_blocks_hashes` - Get alternative blocks hashes

### Key Images & Transactions
- ✅ `/is_key_image_spent` - Check if outputs have been spent
- ✅ `/send_raw_transaction` - Broadcast a raw transaction to the network

### Mining Control
- ✅ `/start_mining` - Start mining on the daemon
- ✅ `/stop_mining` - Stop mining on the daemon
- ✅ `/mining_status` - Get mining status

### Transaction Pool
- ✅ `/get_transaction_pool` - Get transaction pool contents
- ✅ `/get_transaction_pool_hashes` - Get transaction pool hashes
- ✅ `/get_transaction_pool_stats` - Get transaction pool statistics

### Peer Management
- ✅ `/get_peer_list` - Get the known peers list
- ✅ `/get_public_nodes` - Get public nodes on the network

### Configuration
- ✅ `/set_log_hash_rate` - Set log hash rate display
- ✅ `/set_log_level` - Set the daemon log level
- ✅ `/set_log_categories` - Set the daemon log categories
- ✅ `/set_bootstrap_daemon` - Set bootstrap daemon

### Network & Limits
- ✅ `/get_limit` - Get bandwidth limits
- ✅ `/set_limit` - Set bandwidth limits
- ✅ `/out_peers` - Set maximum number of outgoing peers
- ✅ `/in_peers` - Set maximum number of incoming peers
- ✅ `/get_net_stats` - Get network statistics

### Advanced Operations
- ✅ `/get_outs` - Get outputs
- ✅ `/save_bc` - Save the blockchain
- ✅ `/update` - Check for daemon updates
- ✅ `/pop_blocks` - Pop blocks from the blockchain
- ✅ `/stop_daemon` - Stop the daemon

## Coverage Statistics

**Total Methods**: 60+
**Implemented**: 60+ (100%)
**Not Implemented**: 0 (0%)

## Notes

- All methods are implemented with full request/response type definitions
- Supports both JSON-RPC 2.0 (via `/json_rpc`) and direct HTTP endpoints
- Custom HTTP transport support for authentication (digest auth, etc.)
- Complete struct definitions with all fields documented in the official API

## Testing

To test the daemon client:

```go
package main

import (
    "fmt"
    "log"

    "github.com/boomhut/go-monero-rpc-client/daemon"
)

func main() {
    client := daemon.New(daemon.Config{
        Address: "http://localhost:18081",
    })

    // Test blockchain queries
    count, err := client.GetBlockCount()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Block count: %d\n", count.Count)

    // Test node info
    info, err := client.GetInfo()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Height: %d, Synchronized: %v\n", info.Height, info.Synchronized)

    // Test transaction pool
    pool, err := client.GetTransactionPool()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("TX pool size: %d\n", len(pool.Transactions))
}
```

## Compatibility

- **Monero Version**: v0.18.3.x
- **RPC Documentation**: https://docs.getmonero.org/rpc-library/monerod-rpc/
- **Last Updated**: November 2025
