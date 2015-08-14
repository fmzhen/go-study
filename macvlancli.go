package macvlan

import "github.com/codegangsta/cli"

var (
	FlagMacvlanMode  = cli.StringFlag{Name: "mode", Value: macvlanMode, Usage: "name of the macvlan mode bridge|private|passthrough|vepa. By default,bridge mode is implicit: --bridge-name=bridge"}
	FlagGateway      = cli.StringFlag{Name: "gateway", Valur: gatewayIP, Usage: "IP of the default gateway.default : --bridge-ip=172.18.40.1/24"}
	FlagBridgeSubnet = cli.StringFlag{Name: "macvlan-subnet", Value: defaultsubnet, Usage: "subnet for containes"}
	FlagMacvlanEth   = cli.StringFlag{Name: "host-interface", Value: macvlanEthIface, Usage: "the ethernet interface on the underlying OS that will be used as the parent interface that the container will use for external communications"}
)

// Unexported variables
var (
	// TODO: align with dnet-ctl for bridge properties.
	macvlanMode     = "bridge"         // currently only mode supported. Does anyone use the others?
	macvlanEthIface = "eth1"           // parent interface to the macvlan iface
	defaultSubnet   = "192.168.1.0/24" // magic default /24 for demo/testing
	gatewayIP       = "192.168.1.1"    // this is the address of an external route
	cliMTU          = 1500             // generally accepted default MTU
)
