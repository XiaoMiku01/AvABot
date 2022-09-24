package bilipush

import (
	"crypto/tls"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	bilimetadata "github.com/Mrs4s/go-cqhttp/plugins/bilipush/grpc_api/bilibili/metadata"
	device "github.com/Mrs4s/go-cqhttp/plugins/bilipush/grpc_api/bilibili/metadata/device"
	locale "github.com/Mrs4s/go-cqhttp/plugins/bilipush/grpc_api/bilibili/metadata/locale"
	network "github.com/Mrs4s/go-cqhttp/plugins/bilipush/grpc_api/bilibili/metadata/network"
)

var GrpcClient *grpc.ClientConn

//var grpcCtx context.Context

func GrpcInit() error {
	addr := "grpc.biliapi.net:443"
	creds := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS12,
	}))
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}
	var err error
	GrpcClient, err = grpc.Dial(addr, creds, grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("BiliBili GRPC 服务器链接失败: %v", err)
		return err
	}

	return nil
}

func GetMetadata() metadata.MD {
	dev := &device.Device{
		MobiApp:  "android",
		Device:   "phone",
		Build:    6830300,
		Channel:  "bili",
		Buvid:    "XX82B818F96FB2E312B3A1BA44DB41892FF88",
		Platform: "android",
	}
	dev_siz := dev.XXX_Size()
	b := make([]byte, 0, dev_siz)
	dev_bin, _ := dev.XXX_Marshal(b, false)
	lo := &locale.Locale{
		Timezone: "Asia/Shanghai",
	}
	lo_siz := lo.XXX_Size()
	b = make([]byte, 0, lo_siz)
	lo_bin, _ := lo.XXX_Marshal(b, false)
	me := &bilimetadata.Metadata{
		AccessKey: "",
		MobiApp:   "android",
		Device:    "phone",
		Build:     6830300,
		Channel:   "bili",
		Buvid:     "XX82B818F96FB2E312B3A1BA44DB41892FF88",
		Platform:  "android",
	}
	me_siz := me.XXX_Size()
	b = make([]byte, 0, me_siz)
	me_bin, _ := me.XXX_Marshal(b, false)
	ne := &network.Network{
		Type: network.NetworkType_WIFI,
	}
	ne_siz := ne.XXX_Size()
	b = make([]byte, 0, ne_siz)
	ne_bin, _ := ne.XXX_Marshal(b, false)
	md := metadata.Pairs(
		"x-bili-device-bin", string(dev_bin),
		"x-bili-local-bin", string(lo_bin),
		"x-bili-metadata-bin", string(me_bin),
		"x-bili-network-bin", string(ne_bin),
		"authorization", "",
	)
	return md
}
