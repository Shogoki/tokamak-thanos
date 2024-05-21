// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const LegacyERC20NativeTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/legacy/LegacyERC20NativeToken.sol:LegacyERC20NativeToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/legacy/LegacyERC20NativeToken.sol:LegacyERC20NativeToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/legacy/LegacyERC20NativeToken.sol:LegacyERC20NativeToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/legacy/LegacyERC20NativeToken.sol:LegacyERC20NativeToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/legacy/LegacyERC20NativeToken.sol:LegacyERC20NativeToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var LegacyERC20NativeTokenStorageLayout = new(solc.StorageLayout)

var LegacyERC20NativeTokenDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610387578063e78cea921461033b578063ee9a31a2146103cd57600080fd5b8063ae1f6aaf1461033b578063c01e1bd614610361578063d6c0b2c41461036157600080fd5b80639dc29fac116100bd5780639dc29fac14610302578063a457c2d714610315578063a9059cbb1461032857600080fd5b806370a08231146102d257806395d89b41146102fa57600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004610952565b6103f4565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104e5565b60405161019b919061099b565b61018f610213366004610a37565b610577565b6002545b60405190815260200161019b565b61018f610238366004610a61565b610607565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004610a37565b610692565b61029461028f366004610a37565b61071d565b005b6101f86040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61021c6102e0366004610a9d565b73ffffffffffffffffffffffffffffffffffffffff163190565b6101f86107a5565b610294610310366004610a37565b6107b4565b61018f610323366004610a37565b61083c565b61018f610336366004610a37565b6108c7565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c610395366004610ab8565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000085168314806104ad57507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104dc57507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104f490610aeb565b80601f016020809104026020016040519081016040528092919081815260200182805461052090610aeb565b801561056d5780601f106105425761010080835404028352916020019161056d565b820191906000526020600020905b81548152906001019060200180831161055057829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c656761637945524332304e6174697665546f6b656e3a20617070726f76652060448201527f69732064697361626c656400000000000000000000000000000000000000000060648201526000906084015b60405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4c656761637945524332304e6174697665546f6b656e3a207472616e7366657260448201527f46726f6d2069732064697361626c65640000000000000000000000000000000060648201526000906084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4c656761637945524332304e6174697665546f6b656e3a20696e63726561736560448201527f416c6c6f77616e63652069732064697361626c6564000000000000000000000060648201526000906084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c656761637945524332304e6174697665546f6b656e3a206d696e742069732060448201527f64697361626c656400000000000000000000000000000000000000000000000060648201526084016105fe565b6060600480546104f490610aeb565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c656761637945524332304e6174697665546f6b656e3a206275726e2069732060448201527f64697361626c656400000000000000000000000000000000000000000000000060648201526084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4c656761637945524332304e6174697665546f6b656e3a20646563726561736560448201527f416c6c6f77616e63652069732064697361626c6564000000000000000000000060648201526000906084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4c656761637945524332304e6174697665546f6b656e3a207472616e7366657260448201527f2069732064697361626c6564000000000000000000000000000000000000000060648201526000906084016105fe565b60006020828403121561096457600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461099457600080fd5b9392505050565b600060208083528351808285015260005b818110156109c8578581018301518582016040015282016109ac565b818111156109da576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a3257600080fd5b919050565b60008060408385031215610a4a57600080fd5b610a5383610a0e565b946020939093013593505050565b600080600060608486031215610a7657600080fd5b610a7f84610a0e565b9250610a8d60208501610a0e565b9150604084013590509250925092565b600060208284031215610aaf57600080fd5b61099482610a0e565b60008060408385031215610acb57600080fd5b610ad483610a0e565b9150610ae260208401610a0e565b90509250929050565b600181811c90821680610aff57607f821691505b602082108103610b38577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(LegacyERC20NativeTokenStorageLayoutJSON), LegacyERC20NativeTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["LegacyERC20NativeToken"] = LegacyERC20NativeTokenStorageLayout
	deployedBytecodes["LegacyERC20NativeToken"] = LegacyERC20NativeTokenDeployedBin
	immutableReferences["LegacyERC20NativeToken"] = true
}
