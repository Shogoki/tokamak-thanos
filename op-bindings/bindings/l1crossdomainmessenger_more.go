// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"249\",\"type\":\"t_contract(OptimismPortal)1019\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"nativeTokenAddress\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1019\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101965760003560e01c80635644cfdf116100e15780639fce812c1161008a578063b28ade2511610064578063b28ade25146104bf578063d764ad0b146104df578063e0e593c5146104f2578063ecc704281461051257600080fd5b80639fce812c1461042b578063a4e7f8bd1461045f578063b1b1b2091461048f57600080fd5b806383a74074116100bb57806383a74074146103e65780638cbeeef2146102f757806392a162cf146103fd57600080fd5b80635644cfdf146103905780636425666b146103a65780636e296e45146103d157600080fd5b80633f827a5a116101435780634c1d6a691161011d5780634c1d6a69146102f75780634d0047ee1461030d57806354fd4d501461033a57600080fd5b80633f827a5a1461028f5780634273ca16146102b7578063485cc955146102d757600080fd5b80630ff754ea116101745780630ff754ea146102135780632828d7e8146102655780633dbb202b1461027a57600080fd5b806301ffc9a71461019b578063028f85f7146101d05780630c568498146101fe575b600080fd5b3480156101a757600080fd5b506101bb6101b63660046122bd565b610577565b60405190151581526020015b60405180910390f35b3480156101dc57600080fd5b506101e5601081565b60405167ffffffffffffffff90911681526020016101c7565b34801561020a57600080fd5b506101e5603f81565b34801561021f57600080fd5b5060f9546102409073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c7565b34801561027157600080fd5b506101e5604081565b61028d610288366004612386565b610610565b005b34801561029b57600080fd5b506102a4600181565b60405161ffff90911681526020016101c7565b3480156102c357600080fd5b506101bb6102d23660046123ed565b610874565b3480156102e357600080fd5b5061028d6102f2366004612460565b610ad5565b34801561030357600080fd5b506101e5619c4081565b34801561031957600080fd5b5060fa546102409073ffffffffffffffffffffffffffffffffffffffff1681565b34801561034657600080fd5b506103836040518060400160405280600581526020017f312e372e3100000000000000000000000000000000000000000000000000000081525081565b6040516101c7919061250f565b34801561039c57600080fd5b506101e561138881565b3480156103b257600080fd5b5060f95473ffffffffffffffffffffffffffffffffffffffff16610240565b3480156103dd57600080fd5b50610240610d2e565b3480156103f257600080fd5b506101e562030d4081565b34801561040957600080fd5b5061041d610418366004612551565b610e15565b6040516101c7929190612620565b34801561043757600080fd5b506102407f000000000000000000000000000000000000000000000000000000000000000081565b34801561046b57600080fd5b506101bb61047a36600461263f565b60ce6020526000908152604090205460ff1681565b34801561049b57600080fd5b506101bb6104aa36600461263f565b60cb6020526000908152604090205460ff1681565b3480156104cb57600080fd5b506101e56104da366004612658565b610e5e565b61028d6104ed3660046126ac565b610ece565b3480156104fe57600080fd5b5061028d61050d366004612732565b611838565b34801561051e57600080fd5b5061056960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101c7565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4273ca1600000000000000000000000000000000000000000000000000000000148061060a57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6107497f000000000000000000000000000000000000000000000000000000000000000061063f858585610e5e565b347fd764ad0b000000000000000000000000000000000000000000000000000000006106ab60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016106c797969594939291906127ec565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611a20565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856107ce60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516107e095949392919061284b565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60fa5460009073ffffffffffffffffffffffffffffffffffffffff163314610923576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f6f6e6c7920616363657074206e617469766520746f6b656e20617070726f766560448201527f2063616c6c6261636b000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60008061096585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e1592505050565b90925090508515610a365760fa546109959073ffffffffffffffffffffffffffffffffffffffff16893089611beb565b60fa5460f9546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810189905291169063095ea7b3906044016020604051808303816000875af1158015610a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a349190612899565b505b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c4290610a95908b908a90879060009088906004016128bb565b600060405180830381600087803b158015610aaf57600080fd5b505af1158015610ac3573d6000803e3d6000fd5b5060019b9a5050505050505050505050565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610b23575060005460ff8083167401000000000000000000000000000000000000000090920416105b610baf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161091a565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff90911617750100000000000000000000000000000000000000000017905560f9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fa805492851692909116919091179055610c8c611c86565b60f95473ffffffffffffffffffffffffffffffffffffffff1615801590610cca575060fa5473ffffffffffffffffffffffffffffffffffffffff1615155b50600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f74207365740000000000000000000000606482015260840161091a565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b60006060600483511015610e3a57505060408051602081019091526000815262030d40905b825160208401805160e01c93509060048114610e57576004820192505b5050915091565b6000611388619c4080603f610e7a604063ffffffff8816612933565b610e849190612963565b610e8f601088612933565b610e9c9062030d406129b1565b610ea691906129b1565b610eb091906129b1565b610eba91906129b1565b610ec491906129b1565b90505b9392505050565b8315610eff5760f95460fa54610eff9173ffffffffffffffffffffffffffffffffffffffff91821691163087611beb565b60f087901c60028110610fba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161091a565b8061ffff166000036110af57600061100b878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611d5f915050565b600081815260cb602052604090205490915060ff16156110ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161091a565b505b60006110f5898989898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d7e92505050565b90506110ff611da1565b15611136573415611112576111126129dd565b600081815260ce602052604090205460ff1615611131576111316129dd565b611288565b34156111ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161091a565b600081815260ce602052604090205460ff16611288576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161091a565b61129187611e97565b15611344576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161091a565b600081815260cb602052604090205460ff16156113e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161091a565b611404856113f5611388619c406129b1565b67ffffffffffffffff16611eda565b158061142a575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561154357600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161153c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161091a565b505061182f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556001861561162b5760fa546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a90529091169063095ea7b3906044016020604051808303816000875af1158015611604573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116289190612899565b90505b600061167d89619c405a61163f9190612a0c565b600089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611ef892505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508080156116b55750815b1561171d57600083815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261182a565b600083815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161182a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161091a565b505050505b50505050505050565b6118ef7f0000000000000000000000000000000000000000000000000000000000000000611867858585610e5e565b867fd764ad0b000000000000000000000000000000000000000000000000000000006118d360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b8b898c8c6040516024016106c797969594939291906127ec565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561197460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161198695949392919061284b565b60405180910390a260405184815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b905090565b3415611a88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f44656e79206465706f736974696e672045544800000000000000000000000000604482015260640161091a565b8115611b545760fa54611ab39073ffffffffffffffffffffffffffffffffffffffff16333085611beb565b60fa5460f9546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810185905291169063095ea7b3906044016020604051808303816000875af1158015611b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b529190612899565b505b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c4290611bb3908790869088906000908890600401612a23565b600060405180830381600087803b158015611bcd57600080fd5b505af1158015611be1573d6000803e3d6000fd5b5050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611c80908590611f12565b50505050565b6000547501000000000000000000000000000000000000000000900460ff16611d31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161091a565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b6000611d6d85858585612023565b805190602001209050949350505050565b6000611d8e8787878787876120bc565b8051906020012090509695505050505050565b60f95460009073ffffffffffffffffffffffffffffffffffffffff1633148015611a1b575060f954604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691639bf62d829160048083019260209291908290030181865afa158015611e57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e7b9190612a70565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff821630148061060a57505060f95473ffffffffffffffffffffffffffffffffffffffff90811691161490565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6000611f74826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661215b9092919063ffffffff16565b80519091501561201e5780806020019051810190611f929190612899565b61201e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161091a565b505050565b60608484848460405160240161203c9493929190612a8d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016120d996959493929190612ad7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b6060610ec484846000858573ffffffffffffffffffffffffffffffffffffffff85163b6121e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161091a565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161220d9190612b2e565b60006040518083038185875af1925050503d806000811461224a576040519150601f19603f3d011682016040523d82523d6000602084013e61224f565b606091505b509150915061225f82828661226a565b979650505050505050565b60608315612279575081610ec7565b8251156122895782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091a919061250f565b6000602082840312156122cf57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610ec757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461232157600080fd5b50565b60008083601f84011261233657600080fd5b50813567ffffffffffffffff81111561234e57600080fd5b60208301915083602082850101111561236657600080fd5b9250929050565b803563ffffffff8116811461238157600080fd5b919050565b6000806000806060858703121561239c57600080fd5b84356123a7816122ff565b9350602085013567ffffffffffffffff8111156123c357600080fd5b6123cf87828801612324565b90945092506123e290506040860161236d565b905092959194509250565b60008060008060006080868803121561240557600080fd5b8535612410816122ff565b94506020860135612420816122ff565b935060408601359250606086013567ffffffffffffffff81111561244357600080fd5b61244f88828901612324565b969995985093965092949392505050565b6000806040838503121561247357600080fd5b823561247e816122ff565b9150602083013561248e816122ff565b809150509250929050565b60005b838110156124b457818101518382015260200161249c565b83811115611c805750506000910152565b600081518084526124dd816020860160208601612499565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610ec760208301846124c5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561256357600080fd5b813567ffffffffffffffff8082111561257b57600080fd5b818401915084601f83011261258f57600080fd5b8135818111156125a1576125a1612522565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156125e7576125e7612522565b8160405282815287602084870101111561260057600080fd5b826020860160208301376000928101602001929092525095945050505050565b63ffffffff83168152604060208201526000610ec460408301846124c5565b60006020828403121561265157600080fd5b5035919050565b60008060006040848603121561266d57600080fd5b833567ffffffffffffffff81111561268457600080fd5b61269086828701612324565b90945092506126a390506020850161236d565b90509250925092565b600080600080600080600060c0888a0312156126c757600080fd5b8735965060208801356126d9816122ff565b955060408801356126e9816122ff565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561271357600080fd5b61271f8a828b01612324565b989b979a50959850939692959293505050565b60008060008060006080868803121561274a57600080fd5b8535612755816122ff565b945060208601359350604086013567ffffffffffffffff81111561277857600080fd5b61278488828901612324565b909450925061279790506060870161236d565b90509295509295909350565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261283e60c0830184866127a3565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815260806020820152600061287b6080830186886127a3565b905083604083015263ffffffff831660608301529695505050505050565b6000602082840312156128ab57600080fd5b81518015158114610ec757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015263ffffffff84166040820152821515606082015260a06080820152600061225f60a08301846124c5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168185168183048111821515161561295a5761295a612904565b02949350505050565b600067ffffffffffffffff808416806129a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff8083168185168083038211156129d4576129d4612904565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015612a1e57612a1e612904565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a06080820152600061225f60a08301846124c5565b600060208284031215612a8257600080fd5b8151610ec7816122ff565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152612ac660808301856124c5565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152612b2260c08301846124c5565b98975050505050505050565b60008251612b40818460208701612499565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
