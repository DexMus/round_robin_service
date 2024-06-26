new:
	hertztool new --psm=github.DexMus.round_robin_service -force --idl=idl/round_robin_service.thrift

gen:
	hertztool update -idl=idl/round_robin_service.thrift --upgrade_to_hertztool_v3 && go mod tidy && go mod verify