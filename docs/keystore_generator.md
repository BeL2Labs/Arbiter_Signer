# Keystore Generator

A command-line tool for generating Ethereum and Bitcoin keystore files.

## Features

- Generate Ethereum keystore files (UTC format)
- Generate Bitcoin keystore files (WIF format)
- Support for password protection
- Simple command-line interface

## Usage

### Generate Ethereum Keystore

#### Generation Rules
- Uses go-ethereum's keystore package
- Private key is encrypted using AES-128-CTR
- Key derivation uses scrypt with standard parameters:
  - N = 262144
  - r = 8  
  - p = 1
- Generates a JSON format file containing:
  - Encrypted private key
  - Encryption parameters (iv, salt)
  - Scrypt parameters
  - Public address
  - Version information

#### Command
```bash
./keystore-generator -c eth -s <private_key> -p <password> [-o <output_file>]
```

Examples:
```bash
# Basic usage
./keystore-generator -c eth -s 4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d -p password123

# With custom output filename
./keystore-generator -c eth -s 4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d -p password123 -o my_eth_wallet
```

### Generate Bitcoin Keystore

#### Generation Rules  
- Uses btcutil to generate WIF format private key
- Private key is encrypted using custom encryption (btc.Encrypt)
- Generates an encrypted text file containing:
  - Encrypted WIF private key
  - Encryption metadata

#### Command
```bash
./keystore-generator -c btc -s <private_key> -p <password> [-o <output_file>]
```

Examples:
```bash
# Basic usage
./keystore-generator -c btc -s 4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d -p password123

# With custom output filename
./keystore-generator -c btc -s 4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d -p password123 -o my_bitcoin_wallet
```

## Parameters

| Flag | Description                     | Required |
|------|---------------------------------|----------|
| -c   | Chain type (eth or btc)         | Yes      |
| -s   | Private key (64 hex characters) | Yes (for generation) |
| -f   | Keystore file path              | Yes (for parsing) |
| -p   | Password for keystore           | Yes (for parsing), No (for generation) |
| -o   | Output filename (optional)      | No       |

## Output

- By default, keystore files are saved in the current working directory:
  - Ethereum keystore files use UTC format
  - Bitcoin keystore files use WIF format
- When using the `-o` parameter, files are saved in the specified path:
  - Relative paths are resolved from current working directory
  - Absolute paths are supported
  - Parent directories will be created automatically if needed


### Parse Keystore Files

#### Parse Ethereum Keystore
```bash
./keystore-generator -c eth -f <keystore_file> -p <password>
```

Example:
```bash
./keystore-generator -c eth -f ./my_eth_wallet -p password123
```

#### Parse Bitcoin Keystore
```bash
./keystore-generator -c btc -f <keystore_file> -p <password>
```

Example:
```bash
./keystore-generator -c btc -f ./my_bitcoin_wallet -p password123
```
