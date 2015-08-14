package macvlan

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/docker/libnetwork/ipallocator"
	"github.com/gorilla/mux"
	"github.com/samalba/dockerclient"
	"github.com/vishvananda/netlink"
)

const(
	MethodReceiver = "NetworkDriver"
	bridgeMode + "bridge"
	defaultRoute = "0.0.0.0/0"
	containerIfacePrefix = "eth"
	defaultMTU = 1500
	minMTU = 68	
)

type Driver interface{
	Listen(string) error	
}

type bridgeOpts struct{
	brName string
	brSubnet net.IPNet
	brIP net.IPNet	
}

type driver struct{
	dockerer
	pluginConfig
	ipallocator *ipallocator.IPAllocator
	version string
	network string
	cidr *net.IPNet
	nameserver string
}

type pluginConfig struct {
	mtu int
	mode string
	hostIface string
	containerSubnet *net.IPNet
	gateway net.IP
}

func New(version string,ctx *cli.Context)(Driver,error){
	docker,err := dockerclient.NewDockerClient("unix:///var/run/docker.sock",nil)
	if err != nil{
		return nil,Errorf("could not connect to docker: %s",err)	
	}
	if ctx.String("host-interface") == "" {
		log.Fatalf("Required flag [host-interface] that is used for off box communication was not defined.")
	}	
	
	if ok := validateHostIface(ctx.String("host-interface")); !ok{
		log.Fatal("Requird field host-interface was not found")
	}
	macvlanEthIface = ctx.String("host-interface")
	if ctx.Int("mtu") <=0 {
		cliMTU = cliMTU
	}else if ctx.Int("mtu")>=minMTU{
		cliMTU = ctx.Int("mtu")	
	}else {
		log.Fatal("the MTU value must greater than %d",minMTU)	
	}
	
	containerGW,containerNet,err := net.ParseCIDR(ctx.String("macvlan-subnet"))
	if err!=nil {
		log.Fatal("subnet have errors")
	}
	if ctx.String("mode")==""{
		macvlanMode = bridgeMode
	}
	switch ctx.String("mode"){
		case bridgeMode:
			macvlanMode = bridgeMode
		default:
			log.Fatal("Invalid macvlan mode supplied")
	}
	if ctx.String("gateway") !=""{
		cliGateway := net.ParseIP(ctx.String("gateway"))
		if err!=nil{
			log.Fatal("have error")	
		}
		containerGW = cliGateway
	}else{
		containerGW = ipIncrement(containerGW)	
	}
	pluginOpts := &pluginConfig{
		mtu: cliMTU,
		mode: macvlanMode,
		containerSubnet: containerNet,
		gatewayIP : containerGW,
		hostIface: macvlanEthIface,	
	}
	log.Infof("plugin configuration options are")
	
	ipallocator = :ipallocator.New()
	d:=driver{
		dockerer:dockerer{
			client:docker,
		},
		ipallocator : ipallocator,
		version: version,
		pluginConfig: *pluginOpts\
	}
	return d,nil
}

func (driver *driver) Listen(socket string) error {
	router :=mux.NewRouter()
	router.NotfoundHandler = http.HandlerFunc(notFound)
	
	router.Method("GET").Path("/status").HandlerFunc(driver.status)
	router.Methods("post").Path("plugin.Activate").HandlerFunc(driver.handshake)
	
	handleMethod := func(method string,h http.HandlerFunc){
		router.Methods("Post").Path
	}
	handleMethod("CreateNetwork",driver.createNetwork)
	handleMethod("deleteNetwork",driver.deleteNetwork)
	handleMethod("EndpointOperInfo",driver.infoEndpoint)
	listener,err = net.listen("unix",socket)
	if err != nil{
		return err
	}
	return http.Serve(listener,router)
}

func notFound(w http.ResponseEriter, r *http.Request){
	log.warf("not found")
	http.NotFound(w,r)	
}

func sendError(w http.ResponseWriter,msg string,code int){
	log.Errorf("%d %s",code,msg)
	http.Error(w,msg,code)	
}

func errorResponsef(){
	json.NewEncoder(w).Encode(map[string]string{
		"err": fmt.sprintf(fmtsring,item...)\,
	})
}

func objectResponse(w http.responsewriter,obj interface{}){
	if err:=json.NewEncoder(w).encode(obj)
}

func (driver *driver) handshake(w http.ResponseWriter,r *http.Request){
	err := json.NewEncoder(w).Encode(&handshakeResp){
		[]string{"NetworkDriver"},
	}
	if err != nil{
		sendError(w,"encode error",http.StatusInternalServerError)	
		return
	}
	log.Debug("Handshake completed")
}


