Go Monero RPC Client
====================

<p align="center">
<img src="https://github.com/boomhut/go-monero-rpc-client/raw/master/media/img/monero_gopher.png" alt="Monero Gopher" width="200" />
</p>

A comprehensive Go client library for interacting with Monero's Wallet RPC and Daemon RPC interfaces. This package provides type-safe, idiomatic Go bindings for all Monero RPC methods, making it easy to build applications that interact with Monero wallets and nodes.

**Key Features:**
- 🔐 Full Wallet RPC support (95+ methods)
- 🌐 Complete Daemon/Node RPC support (60+ methods)
- 🔒 Digest authentication support
- 📦 Type-safe request/response structs
- ✅ Compatible with Monero v0.18.3.x
- 🔄 Regular updates to match latest Monero RPC API

This package is based on [go-monero-rpc-client](https://github.com/omani/go-monero-rpc-client). 

## Wallet RPC Client

[![GoDoc](https://godoc.org/github.com/boomhut/go-monero-rpc-client/wallet?status.svg)](https://godoc.org/github.com/boomhut/go-monero-rpc-client/wallet)

### Monero RPC Version
The ```go-monero-rpc-client/wallet``` package is fully compatible with the latest [Monero Wallet RPC API](https://docs.getmonero.org/rpc-library/wallet-rpc/) as of November 2025.

### Latest Updates (v2.0)
This client now supports all methods from the latest Monero Wallet RPC API, including:
- **New Methods**: `auto_refresh`, `describe_transfer`, `edit_address_book`, `estimate_tx_size_and_weight`, `exchange_multisig_keys`, `freeze`, `frozen`, `thaw`, `scan_tx`, `set_daemon`, `setup_background_sync`, `start_background_sync`, `stop_background_sync`, `get_default_fee_priority`
- **Enhanced Transfer Methods**: Added support for `subtract_fee_from_outputs`, `amounts_by_dest`, `spent_key_images`, and improved weight tracking
- **Improved Balance Info**: Added `time_to_unlock` and `blocks_to_unlock` fields
- **Better Transfer Tracking**: Added `amounts`, `subaddr_indices`, and `locked` status fields
- **Enhanced Multisig**: Full support for the new `exchange_multisig_keys` workflow

All changes are backwards compatible with existing code.

### Installation

Install the package using Go modules:

```sh
go get -u github.com/boomhut/go-monero-rpc-client
```

### Prerequisites

Before using the wallet client, you need to run `monero-wallet-rpc`. This is the RPC server that comes with Monero and allows external applications to interact with your wallet.

#### Starting monero-wallet-rpc (Development Mode - No Authentication)

For development and testing, you can start the wallet RPC without authentication:

```sh
./monero-wallet-rpc \
  --wallet-file /home/$user/stagenetwallet/stagenetwallet \
  --daemon-address YOUR_STAGENET_NODE:38081 \
  --stagenet \
  --rpc-bind-port 6061 \
  --password 'mystagenetwalletpassword' \
  --disable-rpc-login
```

**Parameter explanations:**
- `--wallet-file`: Path to your wallet file
- `--daemon-address`: Monero node to connect to (use stagenet for testing)
- `--stagenet`: Use Monero's staging network (safe for testing with fake XMR)
- `--rpc-bind-port`: Port where the RPC server will listen
- `--password`: Your wallet password
- `--disable-rpc-login`: Disable RPC authentication (⚠️ only for development!)

**Network Options:**
- For testing: Use `--stagenet` or `--testnet` with public nodes from [monero.fail](https://monero.fail/?chain=monero&network=stagenet)
- For production: Use `--mainnet` with your own node or trusted remote node

### Basic Usage Example

```go
package main

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/boomhut/go-monero-rpc-client/wallet"
)

func main() {
  // Create a new wallet client instance
  // The Address should point to your running monero-wallet-rpc instance
  client := wallet.New(wallet.Config{
    Address: "http://127.0.0.1:6061/json_rpc",
  })

  // Example 1: Check wallet balance
  // AccountIndex 0 is the primary account
  balance, err := client.GetBalance(&wallet.RequestGetBalance{
    AccountIndex: 0,
  })
  if err != nil {
    log.Fatalf("Failed to get balance: %v", err)
  }

  // Display balance information
  fmt.Printf("Balance: %d atomic units\n", balance.Balance)
  fmt.Printf("Unlocked Balance: %d atomic units\n", balance.UnlockedBalance)
  fmt.Printf("Time to unlock: %d seconds\n", balance.TimeToUnlock)
  fmt.Printf("Blocks to unlock: %d blocks\n", balance.BlocksToUnlock)

  // Example 2: Get wallet address
  address, err := client.GetAddress(&wallet.RequestGetAddress{
    AccountIndex: 0,
  })
  if err != nil {
    log.Fatalf("Failed to get address: %v", err)
  }
  fmt.Printf("Primary Address: %s\n", address.Address)

  // Example 3: Get incoming transfers
  // This retrieves all confirmed incoming transactions
  transfers, err := client.GetTransfers(&wallet.RequestGetTransfers{
    AccountIndex: 0,
    In:           true,  // Get incoming transfers
    Out:          false, // Don't get outgoing transfers
    Pending:      false, // Don't get pending transfers
    Pool:         false, // Don't get transfers in the pool
  })
  if err != nil {
    log.Fatalf("Failed to get transfers: %v", err)
  }

  // Display incoming transfers
  fmt.Printf("\nIncoming Transfers (%d total):\n", len(transfers.In))
  for i, transfer := range transfers.In {
    fmt.Printf("\n--- Transfer %d ---\n", i+1)
    fmt.Printf("  TX ID: %s\n", transfer.Txid)
    fmt.Printf("  Amount: %d atomic units\n", transfer.Amount)
    fmt.Printf("  Confirmations: %d\n", transfer.Confirmations)
    fmt.Printf("  Height: %d\n", transfer.Height)
    fmt.Printf("  Timestamp: %d\n", transfer.Timestamp)
    
    // Pretty print full details as JSON
    details, _ := json.MarshalIndent(transfer, "  ", "  ")
    fmt.Printf("  Details:\n  %s\n", string(details))
  }
}
```

### Production Usage with Authentication

For production environments, **always enable RPC authentication** to secure your wallet:

```sh
./monero-wallet-rpc \
  --wallet-file /home/$user/mainnetwallet/mainnetwallet \
  --daemon-address localhost:18081 \
  --mainnet \
  --rpc-bind-port 18082 \
  --password 'mysecurewalletpassword' \
  --rpc-login myuser:mystrongpassword
```

**Security Note:** The `--rpc-login` parameter enables HTTP Digest Authentication. Always use strong credentials!

#### Authenticated Client Example

```go
package main

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/boomhut/go-monero-rpc-client/wallet"
  "github.com/xinsnake/go-http-digest-auth-client"
)

func main() {
  // Create HTTP Digest authentication transport
  // This matches the --rpc-login credentials from monero-wallet-rpc
  transport := httpdigest.New("myuser", "mystrongpassword")

  // Create authenticated wallet client
  client := wallet.New(wallet.Config{
    Address:   "http://127.0.0.1:18082/json_rpc",
    Transport: transport, // Add the auth transport
  })

  // Now all requests will be authenticated
  balance, err := client.GetBalance(&wallet.RequestGetBalance{
    AccountIndex: 0,
  })
  if err != nil {
    log.Fatalf("Failed to get balance: %v", err)
  }

  fmt.Printf("Authenticated request successful!\n")
  fmt.Printf("Balance: %d atomic units\n", balance.Balance)
}
```

### Advanced Wallet Examples

#### Sending Monero

```go
// Send XMR to one or more destinations
transfer, err := client.Transfer(&wallet.RequestTransfer{
  Destinations: []*wallet.Destination{
    {
      Address: "recipient_address_here",
      Amount:  1000000000000, // Amount in atomic units (1 XMR = 1e12 atomic units)
    },
  },
  AccountIndex: 0,
  Priority:     wallet.PriorityDefault, // Transaction priority (affects fee)
  GetTxKey:     true,                   // Return transaction key for proof
})
if err != nil {
  log.Fatalf("Transfer failed: %v", err)
}

fmt.Printf("Transaction sent!\n")
fmt.Printf("TX Hash: %s\n", transfer.TxHash)
fmt.Printf("TX Key: %s\n", transfer.TxKey)
fmt.Printf("Fee: %d atomic units\n", transfer.Fee)
```

#### Creating Subaddresses

```go
// Create a new subaddress for receiving payments
// Subaddresses improve privacy by not reusing addresses
subaddress, err := client.CreateAddress(&wallet.RequestCreateAddress{
  AccountIndex: 0,
  Label:        "Payment from Customer #123", // Optional label for organization
})
if err != nil {
  log.Fatalf("Failed to create subaddress: %v", err)
}

fmt.Printf("New subaddress created: %s\n", subaddress.Address)
fmt.Printf("Subaddress index: %d\n", subaddress.AddressIndex)
```

#### Output Management (Freezing/Thawing)

```go
// Freeze an output to prevent it from being spent
// Useful for reserving specific outputs or compliance scenarios
err := client.Freeze(&wallet.RequestFreeze{
  KeyImage: "key_image_of_output",
})
if err != nil {
  log.Fatalf("Failed to freeze output: %v", err)
}

// Check if an output is frozen
frozen, err := client.Frozen(&wallet.RequestFrozen{
  KeyImage: "key_image_of_output",
})
if err != nil {
  log.Fatalf("Failed to check frozen status: %v", err)
}
fmt.Printf("Output is frozen: %v\n", frozen.Frozen)

// Thaw (unfreeze) the output when ready to use it
err = client.Thaw(&wallet.RequestThaw{
  KeyImage: "key_image_of_output",
})
```

#### Background Sync for Mobile Wallets

```go
// Setup background sync mode - useful for mobile wallets
// Allows wallet to sync efficiently while app is in background
setupResp, err := client.SetupBackgroundSync(&wallet.RequestSetupBackgroundSync{
  BackgroundSyncType: "reuse-wallet",
  WalletPassword:     "your_wallet_password",
})
if err != nil {
  log.Fatalf("Failed to setup background sync: %v", err)
}

// Start background syncing
err = client.StartBackgroundSync()
if err != nil {
  log.Fatalf("Failed to start background sync: %v", err)
}

// When foreground resumes, stop background sync
err = client.StopBackgroundSync()
```

## Daemon RPC Client

[![GoDoc](https://godoc.org/github.com/boomhut/go-monero-rpc-client/daemon?status.svg)](https://godoc.org/github.com/boomhut/go-monero-rpc-client/daemon)

### Overview

The daemon client lets you interact with a Monero node (`monerod`). This is useful for building:
- **Block explorers** - Query blockchain data and transaction history
- **Monitoring tools** - Track node health, sync status, and network activity
- **Mining applications** - Control and monitor mining operations
- **Network analysis** - Study peer connections and transaction pool

### Monero RPC Version
The ```go-monero-rpc-client/daemon``` package is fully compatible with the latest [Monero Daemon RPC API](https://docs.getmonero.org/rpc-library/monerod-rpc/) as of November 2025.

### Features
- **Blockchain Operations**: Query blocks, headers, transactions, and chain information
- **Network Management**: Monitor peers, connections, and network statistics
- **Mining Support**: Start/stop mining, check mining status, and get miner data
- **Transaction Pool**: View and manage mempool transactions
- **Node Administration**: Configure limits, logging, bootstrap daemon, and more
- **Advanced Features**: Output distribution, fee estimation, PoW calculation, auxiliary PoW support

### Prerequisites

You need a running Monero daemon (`monerod`). Start it with:

```sh
# For mainnet
./monerod --rpc-bind-port 18081

# For stagenet (testing)
./monerod --stagenet --rpc-bind-port 38081

# With RPC authentication (recommended for production)
./monerod --rpc-bind-port 18081 --rpc-login myuser:mypassword
```

### Basic Usage Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/boomhut/go-monero-rpc-client/daemon"
)

func main() {
  // Create a new daemon client
  // Connect to your local node or a remote node
  client := daemon.New(daemon.Config{
    Address: "http://127.0.0.1:18081", // Mainnet default port
  })

  // Example 1: Get node information
  info, err := client.GetInfo()
  if err != nil {
    log.Fatalf("Failed to get node info: %v", err)
  }

  fmt.Printf("=== Node Information ===\n")
  fmt.Printf("Height: %d\n", info.Height)
  fmt.Printf("Target Height: %d\n", info.TargetHeight)
  fmt.Printf("Synchronized: %v\n", info.Synchronized)
  fmt.Printf("Network: %s\n", info.Nettype)
  fmt.Printf("Incoming Connections: %d\n", info.IncomingConnectionsCount)
  fmt.Printf("Outgoing Connections: %d\n", info.OutgoingConnectionsCount)
  fmt.Printf("TX Pool Size: %d\n", info.TxPoolSize)
  fmt.Printf("Database Size: %d bytes\n", info.DatabaseSize)

  // Example 2: Get blockchain height
  height, err := client.GetHeight()
  if err != nil {
    log.Fatalf("Failed to get height: %v", err)
  }
  fmt.Printf("\nCurrent blockchain height: %d\n", height.Height)

  // Example 3: Get last block header
  header, err := client.GetLastBlockHeader()
  if err != nil {
    log.Fatalf("Failed to get last block header: %v", err)
  }

  fmt.Printf("\n=== Last Block ===\n")
  fmt.Printf("Hash: %s\n", header.BlockHeader.Hash)
  fmt.Printf("Height: %d\n", header.BlockHeader.Height)
  fmt.Printf("Timestamp: %d\n", header.BlockHeader.Timestamp)
  fmt.Printf("Difficulty: %d\n", header.BlockHeader.Difficulty)
  fmt.Printf("Reward: %d atomic units\n", header.BlockHeader.Reward)
  fmt.Printf("Number of TXs: %d\n", header.BlockHeader.NumTxes)

  // Example 4: Check sync status
  if info.Height < info.TargetHeight {
    remaining := info.TargetHeight - info.Height
    fmt.Printf("\n⚠️  Node is syncing: %d blocks remaining\n", remaining)
  } else {
    fmt.Printf("\n✓ Node is fully synchronized\n")
  }
}
```

### Advanced Daemon Examples

#### Querying Specific Blocks

```go
// Get a specific block by height
block, err := client.GetBlock(2000000, false) // height, fillPowHash
if err != nil {
  log.Fatalf("Failed to get block: %v", err)
}

fmt.Printf("Block at height 2000000:\n")
fmt.Printf("Hash: %s\n", block.BlockHeader.Hash)
fmt.Printf("Miner TX Hash: %s\n", block.MinerTxHash)
fmt.Printf("Transaction Hashes: %v\n", block.TxHashes)

// Get a block by hash
block2, err := client.GetBlock("block_hash_here", false)
if err != nil {
  log.Fatalf("Failed to get block: %v", err)
}
```

#### Transaction Pool Monitoring

```go
// Get full transaction pool contents
pool, err := client.GetTransactionPool()
if err != nil {
  log.Fatalf("Failed to get tx pool: %v", err)
}

fmt.Printf("Transactions in pool: %d\n", len(pool.Transactions))
for i, tx := range pool.Transactions {
  fmt.Printf("\nTX %d:\n", i+1)
  fmt.Printf("  Hash: %s\n", tx.IDHash)
  fmt.Printf("  Fee: %d atomic units\n", tx.Fee)
  fmt.Printf("  Size: %d bytes\n", tx.BlobSize)
  fmt.Printf("  Receive Time: %d\n", tx.ReceiveTime)
  fmt.Printf("  Relay Count: %d\n", tx.RelayedCount)
}

// Get transaction pool statistics
stats, err := client.GetTransactionPoolStats()
if err != nil {
  log.Fatalf("Failed to get pool stats: %v", err)
}

fmt.Printf("\n=== Pool Statistics ===\n")
fmt.Printf("Total TXs: %d\n", stats.PoolStats.TxsTotal)
fmt.Printf("Total Bytes: %d\n", stats.PoolStats.BytesTotal)
fmt.Printf("Total Fees: %d atomic units\n", stats.PoolStats.FeeTotal)
fmt.Printf("Oldest TX: %d seconds\n", stats.PoolStats.Oldest)
```

#### Network Monitoring

```go
// Get peer connections
connections, err := client.GetConnections()
if err != nil {
  log.Fatalf("Failed to get connections: %v", err)
}

fmt.Printf("Active Connections: %d\n", len(connections.Connections))
for i, conn := range connections.Connections {
  fmt.Printf("\nPeer %d:\n", i+1)
  fmt.Printf("  Address: %s\n", conn.Address)
  fmt.Printf("  Incoming: %v\n", conn.Incoming)
  fmt.Printf("  Height: %d\n", conn.Height)
  fmt.Printf("  Download: %d bytes/s\n", conn.CurrentDownload)
  fmt.Printf("  Upload: %d bytes/s\n", conn.CurrentUpload)
  fmt.Printf("  Live Time: %d seconds\n", conn.LiveTime)
}

// Get network statistics
netStats, err := client.GetNetStats()
if err != nil {
  log.Fatalf("Failed to get network stats: %v", err)
}

fmt.Printf("\n=== Network Statistics ===\n")
fmt.Printf("Total Bytes In: %d\n", netStats.TotalBytesIn)
fmt.Printf("Total Bytes Out: %d\n", netStats.TotalBytesOut)
fmt.Printf("Total Packets In: %d\n", netStats.TotalPacketsIn)
fmt.Printf("Total Packets Out: %d\n", netStats.TotalPacketsOut)
```

#### Fee Estimation

```go
// Get recommended fee per byte
feeEstimate, err := client.GetFeeEstimate(0) // grace blocks (0 for immediate)
if err != nil {
  log.Fatalf("Failed to get fee estimate: %v", err)
}

fmt.Printf("Recommended fee: %d atomic units per byte\n", feeEstimate.Fee)
fmt.Printf("Fee levels available: %v\n", feeEstimate.Fees)
```

#### Mining Control (For Local Nodes)

```go
// Check mining status
miningStatus, err := client.MiningStatus()
if err != nil {
  log.Fatalf("Failed to get mining status: %v", err)
}

fmt.Printf("Mining active: %v\n", miningStatus.Active)
if miningStatus.Active {
  fmt.Printf("Mining speed: %d H/s\n", miningStatus.Speed)
  fmt.Printf("Mining threads: %d\n", miningStatus.ThreadsCount)
  fmt.Printf("Mining address: %s\n", miningStatus.Address)
}

// Start mining (only works on local node)
err = client.StartMining(
  "your_wallet_address",
  2,     // Number of threads
  false, // Background mining
  false, // Ignore battery
)
if err != nil {
  log.Printf("Failed to start mining: %v", err)
}

// Stop mining
err = client.StopMining()
if err != nil {
  log.Printf("Failed to stop mining: %v", err)
}
```

#### Authenticated Daemon Connection

```go
import "github.com/xinsnake/go-http-digest-auth-client"

// Create authenticated transport
transport := httpdigest.New("myuser", "mypassword")

// Create authenticated daemon client
client := daemon.New(daemon.Config{
  Address:   "http://127.0.0.1:18081",
  Transport: transport,
})

// All requests now include authentication
info, err := client.GetInfo()
```

### Using Remote Nodes

You can connect to remote public nodes for quick access without running your own node:

```go
// Connect to a remote node (example)
client := daemon.New(daemon.Config{
  Address: "http://node.moneroworld.com:18089",
})

// Always verify you trust the remote node operator!
// Remote nodes can see your requests but cannot spend your funds
```

**Security Warning:** When using remote nodes:
- They can see your IP address and transaction queries
- Use your own node for maximum privacy
- Only use trusted remote nodes
- Consider using Tor for additional privacy

## API Documentation

For complete API reference, see:
- [Wallet RPC Documentation](https://godoc.org/github.com/boomhut/go-monero-rpc-client/wallet)
- [Daemon RPC Documentation](https://godoc.org/github.com/boomhut/go-monero-rpc-client/daemon)
- [Official Monero RPC Documentation](https://docs.getmonero.org/rpc-library/)

## Testing

Run the test suite:

```sh
go test ./...
```

## Troubleshooting

### "connection refused" errors
- Ensure `monero-wallet-rpc` or `monerod` is running
- Verify the port matches your configuration
- Check firewall settings

### "unauthorized" errors
- Verify RPC credentials match between daemon/wallet and your client code
- Ensure you're using the correct authentication transport

### "method not found" errors
- Update to the latest version: `go get -u github.com/boomhut/go-monero-rpc-client`
- Ensure your Monero wallet/daemon version is compatible (v0.18.3.x+)

## Contributing

Contributions are welcome! Here's how you can help:

1. **Report Bugs**: Open an issue with details about the problem
2. **Suggest Features**: Propose new features or improvements
3. **Submit PRs**: Fork the repo, make changes, and submit a pull request
4. **Improve Docs**: Help make the documentation clearer and more comprehensive

### Development Setup

```sh
# Clone the repository
git clone https://github.com/boomhut/go-monero-rpc-client.git
cd go-monero-rpc-client

# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build ./...
```

## License

MIT License - see LICENSE file for details

## Support

- **Issues**: https://github.com/boomhut/go-monero-rpc-client/issues
- **Monero Community**: https://www.reddit.com/r/Monero/
- **Monero Documentation**: https://docs.getmonero.org/

## Disclaimer

This software is provided as-is. Always test with testnet/stagenet before using on mainnet. The authors are not responsible for any loss of funds.
