## Mini EVM

This is a mini Ethereum runtime virutal machine to test and verify the smart contracts. The core runtime codes are from [Go Ethereum](https://github.com/ethereum/go-ethereum).

[![Build Status](https://travis-ci.org/sec-bit/minievm.svg?branch=master)](https://travis-ci.org/sec-bit/minievm)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/sec-bit?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

## Running

You could put your .sol files into the `sols`directory and need to install the [Solidity Tool](https://github.com/ethereum/solidity). Make sure you have the `solc` in the $PATH
```bash
	which solc
```

Then you can write the features file with `Behavior driven development` style and put it into the `internal/features/contract` directory. This is an example to test the `Ballot.sol`
```
Feature: Create the contract
  Scenario: Test the contract is deployed
    Given Create the evm
    And Deploy the contract "Ballot.sol"
```
Now you can get the [gucumber](https://github.com/gucumber/gucumber) and run all the feature tests. We have the Makefile to keep these things easily. So you just to run
```bash
make
```

## License

The Mini EVM is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html),
also included in our repository in the `COPYING` file.
