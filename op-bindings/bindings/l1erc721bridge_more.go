// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const L1ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"spacer_0_2_30\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_bytes30\"},{\"astId\":1003,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1004,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_contract(StandardBridge)1008\"},{\"astId\":1005,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)46_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"49\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes30\":{\"encoding\":\"inplace\",\"label\":\"bytes30\",\"numberOfBytes\":\"30\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)1008\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e mapping(uint256 =\u003e bool)))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L1ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100c95760003560e01c8063761f449311610081578063aa5574521161005b578063aa5574521461020c578063c4d66de81461021f578063c89701a21461023257600080fd5b8063761f4493146101bd5780637f46ddb2146101d0578063927ede2d146101ee57600080fd5b806354fd4d50116100b257806354fd4d501461012d5780635c975abb146101765780635d93a3fc1461018957600080fd5b80633687011a146100ce5780633cb747bf146100e3575b600080fd5b6100e16100dc366004610e68565b610252565b005b6001546101039073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101696040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516101249190610f56565b60005b6040519015158152602001610124565b610179610197366004610f70565b603160209081526000938452604080852082529284528284209052825290205460ff1681565b6100e16101cb366004610fb1565b6102fe565b60025473ffffffffffffffffffffffffffffffffffffffff16610103565b60015473ffffffffffffffffffffffffffffffffffffffff16610103565b6100e161021a366004611049565b610736565b6100e161022d3660046110c0565b6107f2565b6002546101039073ffffffffffffffffffffffffffffffffffffffff1681565b333b156102e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102f686863333888888886109b7565b505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331480156103d35750600254600154604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691636e296e45916004808201926020929091908290030181865afa158015610397573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103bb91906110dd565b73ffffffffffffffffffffffffffffffffffffffff16145b61045f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016102dd565b3073ffffffffffffffffffffffffffffffffffffffff881603610504576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016102dd565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603160209081526040808320938a1683529281528282208683529052205460ff1615156001146105d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c31204272696467650000000000000060648201526084016102dd565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526031602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b15801561069357600080fd5b505af11580156106a7573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107259493929190611143565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff85166107d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016102dd565b6107e987873388888888886109b7565b50505050505050565b600054610100900460ff16158080156108125750600054600160ff909116105b8061082c5750303b15801561082c575060005460ff166001145b6108b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102dd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561091657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61093482734200000000000000000000000000000000000014610cf7565b801561099757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff8716610a5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016102dd565b600063761f449360e01b888a8989898888604051602401610a819796959493929190611183565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603186528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610bc157600080fd5b505af1158015610bd5573d6000803e3d6000fd5b50506001546002546040517f3dbb202b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9283169450633dbb202b9350610c389290911690859089906004016111e0565b600060405180830381600087803b158015610c5257600080fd5b505af1158015610c66573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610ce49493929190611143565b60405180910390a4505050505050505050565b600054610100900460ff16610d8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102dd565b6001805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560028054929093169116179055565b73ffffffffffffffffffffffffffffffffffffffff81168114610e0357600080fd5b50565b803563ffffffff81168114610e1a57600080fd5b919050565b60008083601f840112610e3157600080fd5b50813567ffffffffffffffff811115610e4957600080fd5b602083019150836020828501011115610e6157600080fd5b9250929050565b60008060008060008060a08789031215610e8157600080fd5b8635610e8c81610de1565b95506020870135610e9c81610de1565b945060408701359350610eb160608801610e06565b9250608087013567ffffffffffffffff811115610ecd57600080fd5b610ed989828a01610e1f565b979a9699509497509295939492505050565b6000815180845260005b81811015610f1157602081850181015186830182015201610ef5565b81811115610f23576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f696020830184610eeb565b9392505050565b600080600060608486031215610f8557600080fd5b8335610f9081610de1565b92506020840135610fa081610de1565b929592945050506040919091013590565b600080600080600080600060c0888a031215610fcc57600080fd5b8735610fd781610de1565b96506020880135610fe781610de1565b95506040880135610ff781610de1565b9450606088013561100781610de1565b93506080880135925060a088013567ffffffffffffffff81111561102a57600080fd5b6110368a828b01610e1f565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561106457600080fd5b873561106f81610de1565b9650602088013561107f81610de1565b9550604088013561108f81610de1565b9450606088013593506110a460808901610e06565b925060a088013567ffffffffffffffff81111561102a57600080fd5b6000602082840312156110d257600080fd5b8135610f6981610de1565b6000602082840312156110ef57600080fd5b8151610f6981610de1565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006111796060830184866110fa565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526111d360c0830184866110fa565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061120f6060830185610eeb565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1ERC721BridgeStorageLayoutJSON), L1ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC721Bridge"] = L1ERC721BridgeStorageLayout
	deployedBytecodes["L1ERC721Bridge"] = L1ERC721BridgeDeployedBin
	immutableReferences["L1ERC721Bridge"] = false
}
