package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bigwhite/tcp-stream-proto/demo4/pkg/frame"
	"github.com/bigwhite/tcp-stream-proto/demo4/pkg/packet"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type customCodecServer struct {
	*gnet.EventServer
	addr       string
	multicore  bool
	async      bool
	codec      gnet.ICodec
	workerPool *goroutine.Pool
}

func (cs *customCodecServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("custom codec server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (cs *customCodecServer) React(framePayload []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// packet decode
	var p packet.Packet
	var ackFramePayload []byte
	p, err := packet.Decode(framePayload)
	if err != nil {
		fmt.Println("react: packet decode error:", err)
		action = gnet.Close // close the connection
		return
	}

	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			action = gnet.Close // close the connection
			return
		}
		out = []byte(ackFramePayload)
		return
	default:
		return nil, gnet.Close // close the connection
	}
}

func customCodecServe(addr string, multicore, async bool, codec gnet.ICodec) {
	var err error
	codec = frame.Frame{}
	cs := &customCodecServer{addr: addr, multicore: multicore, async: async, codec: codec, workerPool: goroutine.Default()}
	err = gnet.Serve(cs, addr, gnet.WithMulticore(multicore), gnet.WithTCPKeepAlive(time.Minute*5), gnet.WithCodec(codec))
	if err != nil {
		panic(err)
	}
}

func main() {
	var port int
	var multicore bool

	// Example command: go run server.go --port 8888 --multicore=true
	flag.IntVar(&port, "port", 8888, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()
	addr := fmt.Sprintf("tcp://:%d", port)
	customCodecServe(addr, multicore, false, nil)
}
