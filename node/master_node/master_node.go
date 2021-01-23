package node

import (
	"bytes"
	"context"
	"crypto/rand"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	maddr "github.com/multiformats/go-multiaddr"
	"io"
	"os"
	"yu/config"
	. "yu/node"
	"yu/storage/kv"
)

var MasterNodeInfoKey = []byte("master-node-info")

type MasterNode struct {
	info    *MasterNodeInfo
	p2pHost host.Host
	metadb  kv.KV
	ctx     context.Context
}

func NewMasterNode(cfg *config.Conf) (*MasterNode, error) {
	metadb, err := kv.NewKV(&cfg.NodeDB)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	p2pHost, err := makeP2pHost(ctx, &cfg.NodeConf)
	if err != nil {
		return nil, err
	}
	data, err := metadb.Get(MasterNodeInfoKey)
	if err != nil {
		return nil, err
	}

	var info *MasterNodeInfo
	if data == nil {
		info = &MasterNodeInfo{
			Name:        cfg.NodeConf.NodeName,
			WorkerNodes: cfg.NodeConf.WorkerNodes,
		}
		infoByt, err := info.EncodeMasterNodeInfo()
		if err != nil {
			return nil, err
		}
		err = metadb.Set(MasterNodeInfoKey, infoByt)
		if err != nil {
			return nil, err
		}
	} else {
		info, err = DecodeMasterNodeInfo(data)
		if err != nil {
			return nil, err
		}
	}

	info.P2pID = p2pHost.ID().String()

	return &MasterNode{
		info,
		p2pHost,
		metadb,
		ctx,
	}, nil
}

func (m *MasterNode) ConnectP2PNetwork(cfg *config.NodeConf) error {
	m.p2pHost.SetStreamHandler(protocol.ID(cfg.ProtocolID), m.handleStream)

	for _, addrStr := range cfg.ConnectAddrs {
		addr, err := maddr.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		err = m.p2pHost.Connect(m.ctx, *peer)
		if err != nil {
			return err
		}
	}
	return nil
}

func makeP2pHost(ctx context.Context, cfg *config.NodeConf) (host.Host, error) {
	r, err := loadNodeKeyReader(cfg)
	if err != nil {
		return nil, err
	}
	priv, _, err := crypto.GenerateKeyPairWithReader(cfg.NodeKeyType, cfg.NodeKeyBits, r)
	if err != nil {
		return nil, err
	}
	return libp2p.New(
		ctx,
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(cfg.P2pListenAddrs...),
	)
}

func loadNodeKeyReader(cfg *config.NodeConf) (io.Reader, error) {
	if cfg.NodeKey != "" {
		return bytes.NewBufferString(cfg.NodeKey), nil
	}
	if cfg.NodeKeyFile != "" {
		return os.Open(cfg.NodeKeyFile)
	}
	return rand.Reader, nil
}
