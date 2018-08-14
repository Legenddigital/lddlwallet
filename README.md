lddlwallet
=========

lddlwallet is a daemon handling Legenddigital wallet functionality.  All interaction
with the wallet is performed over RPC.

Public and private keys are derived using the hierarchical
deterministic format described by
[BIP0032](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki).
Unencrypted private keys are not supported and are never written to
disk.  lddlwallet uses the
`m/44'/<coin type>'/<account>'/<branch>/<address index>`
HD path for all derived addresses, as described by
[BIP0044](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki).

Due to the sensitive nature of public data in a BIP0032 wallet,
lddlwallet provides the option of encrypting not just private keys, but
public data as well.  This is intended to thwart privacy risks where a
wallet file is compromised without exposing all current and future
addresses (public keys) managed by the wallet. While access to this
information would not allow an attacker to spend or steal coins, it
does mean they could track all transactions involving your addresses
and therefore know your exact balance.  In a future release, public data
encryption will extend to transactions as well.

lddlwallet provides two modes of operation to connect to the Legenddigital
network.  The first (and default) is to communicate with a single
trusted `lddld` instance using JSON-RPC.  The second is a
privacy-preserving Simplified Payment Verification (SPV) mode (enabled
with the `--spv` flag) where the wallet connects either to specified
peers (with `--spvconnect`) or peers discovered from seeders and other
peers. Both modes can be switched between with just a restart of the
wallet.  It is advised to avoid SPV mode for heavily-used wallets
which require downloading most blocks regardless.

Not all functionality is available when running in SPV mode.  Some of
these features may become available in future versions, but only if a
consensus vote passes to activate the required changes.  Currently,
the following features are disabled or unavailable to SPV wallets:

  * Voting

  * Revoking tickets before expiry

  * Determining exact number of live and missed tickets (as opposed to
    simply unspent).

Wallet clients interact with the wallet using one of two RPC servers:

  1. A legacy JSON-RPC server inspired by the Bitcoin Core rpc server

     The JSON-RPC server exists to ease the migration of wallet applications
     from Core, but complete compatibility is not guaranteed.  Some portions of
     the API (and especially accounts) have to work differently due to other
     design decisions (mostly due to BIP0044).  However, if you find a
     compatibility issue and feel that it could be reasonably supported, please
     report an issue.  This server is enabled by default as long as a username
     and password are provided.

  2. A gRPC server

     The gRPC server uses a new API built for lddlwallet, but the API is not
     stabilized.  This server is enabled by default and may be disabled with
     the config option `--nogrpc`.  If you don't mind applications breaking
     due to API changes, don't want to deal with issues of the legacy API, or
     need notifications for changes to the wallet, this is the RPC server to
     use. The gRPC server is documented [here](./rpc/documentation/README.md).

## Installation and updating

### Windows - MSIs Available

Install the latest MSIs available here:

https://github.com/Legenddigital/Legenddigital-release/releases

### Windows/Linux/BSD/POSIX - Build from source

Building or updating from source requires the following build dependencies:

- **Go 1.9 or 1.10**

  Installation instructions can be found here: http://golang.org/doc/install.
  It is recommended to add `$GOPATH/bin` to your `PATH` at this point.

- **Dep**

  Dep is used to manage project dependencies and provide reproducible builds.
  It is recommended to use the latest Dep release, unless a bug prevents doing
  so.  The latest releases (for both binary and source) can be found
  [here](https://github.com/golang/dep/releases).

Unfortunately, the use of `dep` prevents a handy tool such as `go get` from
automatically downloading, building, and installing the source in a single
command.  Instead, the latest project and dependency sources must be first
obtained manually with `git` and `dep`, and then `go` is used to build and
install the project.

**Getting the source**:

For a first time installation, the project and dependency sources can be
obtained manually with `git` and `dep` (create directories as needed):

```
git clone https://github.com/Legenddigital/lddlwallet $GOPATH/src/github.com/Legenddigital/lddlwallet
cd $GOPATH/src/github.com/Legenddigital/lddlwallet
dep ensure
```

To update an existing source tree, pull the latest changes and install the
matching dependencies:

```
cd $GOPATH/src/github.com/Legenddigital/lddlwallet
git pull
dep ensure
```

**Building/Installing**:

The `go` tool is used to build or install (to `GOPATH`) the project.  Some
example build instructions are provided below (all must run from the `lddlwallet`
project directory).

To build and install `lddlwallet` and all helper commands (in the `cmd`
directory) to `$GOPATH/bin/`, as well as installing all compiled packages to
`$GOPATH/pkg/` (**use this if you are unsure which command to run**):

```
go install . ./cmd/...
```

To build a `lddlwallet` executable and install it to `$GOPATH/bin/`:

```
go install
```

To build a `lddlwallet` executable and place it in the current directory:

```
go build
```

## Docker

All tests and linters may be run in a docker container using the script `run_tests.sh`.  This script defaults to using the current supported version of go.  You can run it with the major version of go you would like to use as the only arguement to test a previous on a previous version of go (generally Legenddigital supports the current version of go and the previous one).

```
./run_tests.sh 1.9
```

To run the tests locally without docker:

```
./run_tests.sh local
```

## Getting Started

The following instructions detail how to get started with lddlwallet connecting
to a localhost lddld.  Commands should be run in `cmd.exe` or PowerShell on
Windows, or any terminal emulator on *nix.

- Run the following command to start lddld:

```
lddld -u rpcuser -P rpcpass
```

- Run the following command to create a wallet:

```
lddlwallet -u rpcuser -P rpcpass --create
```

- Run the following command to start lddlwallet:

```
lddlwallet -u rpcuser -P rpcpass
```

If everything appears to be working, it is recommended at this point to
copy the sample lddld and lddlwallet configurations and update with your
RPC username and password.

PowerShell (Installed from MSI):
```
PS> cp "$env:ProgramFiles\Legenddigital\Lddld\sample-lddld.conf" $env:LOCALAPPDATA\Lddld\lddld.conf
PS> cp "$env:ProgramFiles\Legenddigital\Dcrwallet\sample-lddlwallet.conf" $env:LOCALAPPDATA\Dcrwallet\lddlwallet.conf
PS> $editor $env:LOCALAPPDATA\Lddld\lddld.conf
PS> $editor $env:LOCALAPPDATA\Dcrwallet\lddlwallet.conf
```

PowerShell (Installed from source):
```
PS> cp $env:GOPATH\src\github.com\Legenddigital\lddld\sample-lddld.conf $env:LOCALAPPDATA\Lddld\lddld.conf
PS> cp $env:GOPATH\src\github.com\Legenddigital\lddlwallet\sample-lddlwallet.conf $env:LOCALAPPDATA\Dcrwallet\lddlwallet.conf
PS> $editor $env:LOCALAPPDATA\Lddld\lddld.conf
PS> $editor $env:LOCALAPPDATA\Dcrwallet\lddlwallet.conf
```

Linux/BSD/POSIX (Installed from source):
```bash
$ cp $GOPATH/src/github.com/Legenddigital/lddld/sample-lddld.conf ~/.lddld/lddld.conf
$ cp $GOPATH/src/github.com/Legenddigital/lddlwallet/sample-lddlwallet.conf ~/.lddlwallet/lddlwallet.conf
$ $EDITOR ~/.lddld/lddld.conf
$ $EDITOR ~/.lddlwallet/lddlwallet.conf
```

## Issue Tracker

The [integrated github issue tracker](https://github.com/Legenddigital/lddlwallet/issues)
is used for this project.

## License

lddlwallet is licensed under the liberal ISC License.
