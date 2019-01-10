package main

import (
	"context"
	"flag"
	"strings"
	"fmt"
	"time"
	"strconv"
	"./clustering"
	"./messages"
	"github.com/perlin-network/noise/crypto/ed25519"
	"github.com/perlin-network/noise/log"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/perlin-network/noise/types/opcode"
)


type node struct {
    port int
	shard int
}

var nodeList [4]node
var shardresp [10]int
var n=0

type ClusterPlugin struct{ *network.Plugin }

func (state *ClusterPlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	
	
	case *messages.ClusterMessage:
		index:=int(msg.Shard)
		shardresp[index-1]++
		
	case *messages.ClusterResponse:
		for i:=0 ; i<n ; i++ {
			if nodeList[i].port==int(msg.Port){
				nodeList[i].shard=int(msg.Shard)
				break
			}
		}	
		
	}

	return nil
}


func main() {
	// process other flags
	portFlag := flag.Int("port", 8002, "port to listen to")
	hostFlag := flag.String("host", "localhost", "host to listen to")
	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
	peersFlag := flag.String("peers", "", "peers to connect to")
	flag.Parse()

	port := int(*portFlag)
	host := *hostFlag
	protocol := *protocolFlag
	peers := strings.Split(*peersFlag, ",")

	keys := ed25519.RandomKeyPair()

	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.ClusterMessage{})
	opcode.RegisterMessageType(opcode.Opcode(1001), &messages.ClusterResponse{})

	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(protocol, host, uint16(port)))
	// Register peer discovery plugin.
	builder.AddPlugin(new(discovery.Plugin))

	// Add custom chat plugin.
	builder.AddPlugin(new(ClusterPlugin))

	net, err := builder.Build()
	
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	go net.Listen()

	if len(peers) > 0 {
		net.Bootstrap("tcp://localhost:8001","tcp://localhost:8000","tcp://localhost:8003")
	}
	
	p:=time.Now()
	ctx := network.WithSignMessage(context.Background(), true)
	clusters := clustering.GetClusters()
	time.Sleep(1*time.Second)
	told:=false
	for {
		//t:=int(time.Now().Unix())
		
		//if t%600==0{
			// fmt.Println("n now is")
			// fmt.Println(n)
		if n<4 && !told{

		
			fmt.Println(len(clusters))
			for i := 0; i < len(clusters); i++ {
				for j := 0; j < len(clusters[i]); j++ {
					if port != clusters[i][j] {
						nodeList[n]=node{port:clusters[i][j],shard:i+1}
						n++
						fmt.Println("Printing 2D array")
						fmt.Println(clusters[i][j])
						client, err := net.Client("tcp://localhost:"+strconv.Itoa(clusters[i][j]));
						 if err != nil {
						 	fmt.Println(err)
						 }else{
							fmt.Println("Communicating Shard List")
							client.Tell(ctx,&messages.ClusterMessage{Port: int32(clusters[i][j]),Shard: int32(i+1)})
						 }
						 
						
					} else{
						nodeList[n]=node{port:port, shard:i+1}
						n++
					}
				}
			}
			told=true
		}
		//}else 
			if time.Since(p)>30*time.Second{
				fmt.Println("Now calculating vote")
			max:=0
			for i := 1; i < 10; i++ {
				if shardresp[i]>shardresp[max]{
					max=i
				}
			}
			for i:=0 ; i<n ; i++ {
				if nodeList[i].port==port{
					nodeList[i].shard=max+1
					break
				}
			}	
			net.Broadcast(ctx,&messages.ClusterResponse{Port:int32(port),Shard:int32(max+1)})
			time.Sleep(15*time.Second)
			fmt.Println(nodeList)
			break
		}
	}
}
