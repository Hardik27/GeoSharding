// // // package main

// // // import (

// // // 	"bufio"
// // // 	"context"
// // // 	"flag"
// // // 	"os"
// // // 	"strings"

// // // 	"github.com/perlin-network/noise/crypto/ed25519"
// // // 	"./messages"
// // // 	"github.com/perlin-network/noise/log"
// // // 	"github.com/perlin-network/noise/network"
// // // 	"github.com/perlin-network/noise/network/discovery"
// // // 	"github.com/perlin-network/noise/types/opcode"
// // // )

// // // type ChatPlugin struct{ *network.Plugin }

// // // func (state *ChatPlugin) Receive(ctx *network.PluginContext) error {
// // // 	switch msg := ctx.Message().(type) {
// // // 	case *messages.ChatMessage:
// // // 		log.Info().Msgf("<%s> %s", ctx.Client().ID.Address, msg.Message)
// // // 	}

// // // 	return nil
// // // }

// // // func main() {
// // // 	// process other flags
// // // 	portFlag := flag.Int("port", 8000, "port to listen to")
// // // 	hostFlag := flag.String("host", "localhost", "host to listen to")
// // // 	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
// // // 	peersFlag := flag.String("peers", "", "peers to connect to")
// // // 	flag.Parse()

// // // 	port := uint16(*portFlag)
// // // 	host := *hostFlag
// // // 	protocol := *protocolFlag
// // // 	peers := strings.Split(*peersFlag, ",")

// // // 	keys := ed25519.RandomKeyPair()

// // // 	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
// // // 	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

// // // 	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.ChatMessage{})
// // // 	builder := network.NewBuilder()
// // // 	builder.SetKeys(keys)
// // // 	builder.SetAddress(network.FormatAddress(protocol, host, port))

// // // 	// Register peer discovery plugin.
// // // 	builder.AddPlugin(new(discovery.Plugin))

// // // 	// Add custom chat plugin.
// // // 	builder.AddPlugin(new(ChatPlugin))

// // // 	net, err := builder.Build()
// // // 	if err != nil {
// // // 		log.Fatal().Err(err)
// // // 		return
// // // 	}

// // // 	go net.Listen()

// // // 	if len(peers) > 0 {
// // // 		net.Bootstrap("tcp://localhost:8001","tcp://localhost:8003","tcp://localhost:8002")
// // // 	}

// // // 	reader := bufio.NewReader(os.Stdin)
// // // 	for {
// // // 		input, _ := reader.ReadString('\n')

// // // 		// skip blank lines
// // // 		if len(strings.TrimSpace(input)) == 0 {
// // // 			continue
// // // 		}

// // // 		log.Info().Msgf("<%s> %s", net.Address, input)

// // // 		ctx := network.WithSignMessage(context.Background(), true)
// // // 		net.Broadcast(ctx, &messages.ChatMessage{Message: input})
// // // 	}
// // // }


// // package main

// // import (
// //     //"bufio"
// // 	"context"
// // 	"flag"
// // 	//"os"
// // 	"strings"
// // 	"strconv"
// // 	"./pow"
// // 	"github.com/perlin-network/noise/crypto/ed25519"
// // 	"./messages"
// // 	"./gomeans-master/gomeans"
// // 	"github.com/perlin-network/noise/log"
// // 	"github.com/perlin-network/noise/network"
// // 	"github.com/perlin-network/noise/network/discovery"
// // 	"github.com/perlin-network/noise/types/opcode"
// // )

// // type POWPlugin struct{ *network.Plugin }

// // func (state *POWPlugin) Receive(ctx *network.PluginContext) error {
// // 	switch msg := ctx.Message().(type) {
// // 	case *messages.POWMessage:
// // 		h := pow.Encode(msg.Addr+msg.Pubkey+strconv.Itoa(int(msg.Blocknum))+strconv.Itoa(int(msg.Nonce)))
// // 		if h==msg.Result && pow.IsValidResult(h){
// // 			log.Info().Msgf("<%s> true",ctx.Client().ID.Address)
// // 		}else{
// // 			log.Info().Msgf("<%s> false",ctx.Client().ID.Address)
// // 		}
// // 		log.Info().Msgf("<%s> %s", ctx.Client().ID.Address, strconv.Itoa(int(msg.Nonce)) +"   "+ msg.Pubkey +"   " + msg.Addr+"   " + msg.Result)
// // 	}

// // 	return nil
// // }

// // func main() {
// // 	// process other flags
// // 	portFlag := flag.Int("port", 3000, "port to listen to")
// // 	hostFlag := flag.String("host", "localhost", "host to listen to")
// // 	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
// // 	peersFlag := flag.String("peers", "", "peers to connect to")
// // 	flag.Parse()

// // 	port := uint16(*portFlag)
// // 	host := *hostFlag
// // 	protocol := *protocolFlag
// // 	peers := strings.Split(*peersFlag, ",")

// // 	keys := ed25519.RandomKeyPair()

// // 	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
// // 	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

// // 	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.POWMessage{})
// // 	builder := network.NewBuilder()
// // 	builder.SetKeys(keys)
// // 	builder.SetAddress(network.FormatAddress(protocol, host, port))

// // 	// Register peer discovery plugin.
// // 	builder.AddPlugin(new(discovery.Plugin))

// // 	// Add custom chat plugin.
// // 	builder.AddPlugin(new(POWPlugin))

// // 	net, err := builder.Build()
// // 	if err != nil {
// // 		log.Fatal().Err(err)
// // 		return
// // 	}

// // 	go net.Listen()

// // 	if len(peers) > 0 {
// // 		net.Bootstrap("tcp://localhost:3001","tcp://localhost:3002","tcp://localhost:3003")
// // 	}



// // 	//reader := bufio.NewReader(os.Stdin)
// // 	nonce,result,blocknum,difficulty := pow.Pow(keys.PublicKeyHex())
// // 	ctx := network.WithSignMessage(context.Background(), true)
// // 		net.Broadcast(ctx, &messages.POWMessage{Nonce: int32(nonce),Pubkey:keys.PublicKeyHex(),Addr:pow.GetOutboundIP().String(),Blocknum:int32(blocknum),Difficulty:int32(difficulty),Result:result})
// // 	for {
// // 		//input, _ := reader.ReadString('\n')

// // 		// skip blank lines
// // 		/*if len(strings.TrimSpace(input)) == 0 {
// // 			continue
// // 		}*/

// // 		//log.Info().Msgf("<%s> %s", net.Address, input)

		
// // 	}

// // }

// // func ReadCsvFile(filePath string) (peers string[]) {
// //     // Load a csv file.
// //     f, _ := os.Open(filePath)

// //     // Create a new reader.
// //     r := csv.NewReader(bufio.NewReader(f))
// //     for {
// //         record, err := r.Read()
// //         // Stop at EOF.
// //         if err == io.EOF {
// //             break
// //         }
// //         // Display record.
// //         // ... Display record length.
// //         // ... Display all individual elements of the slice.
// //         fmt.Println(record)
// //         fmt.Println(len(record))
// //         for value := range record {
// //             fmt.Printf("  %v\n", record[value])
// //         }
// //     }
// // }


// package main

// import (
// 	"context"
// 	"flag"
// 	"strings"
// 	"fmt"
// 	"time"
// 	"strconv"
// 	"./clustering"
// 	"./messages"
// 	"github.com/perlin-network/noise/crypto/ed25519"
// 	"github.com/perlin-network/noise/log"
// 	"github.com/perlin-network/noise/network"
// 	"github.com/perlin-network/noise/network/discovery"
// 	"github.com/perlin-network/noise/types/opcode"
// )


// type node struct {
//     port int
// 	shard int
// }

// var nodeList [4]node
// var shardresp [10]int
// var n=0

// type ClusterPlugin struct{ *network.Plugin }

// func (state *ClusterPlugin) Receive(ctx *network.PluginContext) error {
// 	switch msg := ctx.Message().(type) {
	
	
// 	case *messages.ClusterMessage:
// 		index:=int(msg.Shard)
// 		shardresp[index-1]++
		
// 	case *messages.ClusterResponse:
// 		for i:=0 ; i<n ; i++ {
// 			if nodeList[i].port==int(msg.Port){
// 				nodeList[i].shard=int(msg.Shard)
// 				break
// 			}
// 		}	
		
// 	}

// 	return nil
// }


// func main() {
// 	// process other flags
// 	portFlag := flag.Int("port", 8001, "port to listen to")
// 	hostFlag := flag.String("host", "localhost", "host to listen to")
// 	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
// 	peersFlag := flag.String("peers", "", "peers to connect to")
// 	flag.Parse()

// 	port := int(*portFlag)
// 	host := *hostFlag
// 	protocol := *protocolFlag
// 	peers := strings.Split(*peersFlag, ",")

// 	keys := ed25519.RandomKeyPair()

// 	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
// 	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

// 	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.ClusterMessage{})
// 	opcode.RegisterMessageType(opcode.Opcode(1001), &messages.ClusterResponse{})

// 	builder := network.NewBuilder()
// 	builder.SetKeys(keys)
// 	builder.SetAddress(network.FormatAddress(protocol, host, uint16(port)))
// 	// Register peer discovery plugin.
// 	builder.AddPlugin(new(discovery.Plugin))

// 	// Add custom chat plugin.
// 	builder.AddPlugin(new(ClusterPlugin))

// 	net, err := builder.Build()
	
// 	if err != nil {
// 		log.Fatal().Err(err)
// 		return
// 	}

// 	go net.Listen()

// 	if len(peers) > 0 {
// 		net.Bootstrap("tcp://localhost:8000","tcp://localhost:8002","tcp://localhost:8003")
// 	}
	
// 	p :=time.Now()
// 	ctx := network.WithSignMessage(context.Background(), true)
// 	clusters := clustering.GetClusters()
// 	time.Sleep(1*time.Second)
// 	told:=false
// 	for {
		
// 		//t:=int(time.Now().Unix())
		
// 		//if t%600==0{
// 			// fmt.Println("n now is")
// 			// fmt.Println(n)
// 			if n<4 && !told{	
				
// 			//fmt.Println(len(clusters))
			
			
// 			for i := 0; i < len(clusters); i++ {
// 				for j := 0; j < len(clusters[i]); j++ {
// 					if port != clusters[i][j] {
// 						nodeList[n]=node{port:clusters[i][j],shard:i+1}
// 						n++
// 						fmt.Println("Printing 2D array")
// 						fmt.Println(clusters[i][j])
// 						client, err := net.Client("tcp://localhost:"+strconv.Itoa(clusters[i][j]));
// 						 if err != nil {
// 						 	fmt.Println(err)
// 						 }
// 						 if err==nil{
// 							fmt.Println("Communicating Shard List")
// 							client.Tell(ctx,&messages.ClusterMessage{Port: int32(clusters[i][j]),Shard: int32(i+1)})
// 						 }
						 
// 					}/*else{
// 						nodeList[n]=node{port:port,shard:i+1}
// 						n++
// 					}*/
// 				}
// 			}
// 			told=true
// 		}
// 		//}else 
// 			if time.Since(p)>30*time.Second{
// 			fmt.Println("Now calculating vote")
// 			max:=0
// 			for i := 1; i < 10; i++ {
// 				if shardresp[i]>shardresp[max]{
// 					max=i
// 				}
// 			}
// 			for i:=0 ; i<n ; i++ {
// 				if nodeList[i].port==port{
// 					nodeList[i].shard=max+1
// 					break
// 				}
// 			}	
// 			net.Broadcast(ctx,&messages.ClusterResponse{Port:int32(port),Shard:int32(max+1)})
// 			time.Sleep(15*time.Second)
// 			fmt.Println(nodeList)
// 			break
// 		}
// 	}
// }


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
	portFlag := flag.Int("port", 8001, "port to listen to")
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
		net.Bootstrap("tcp://localhost:8002","tcp://localhost:8000","tcp://localhost:8003")
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
