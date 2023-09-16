package main

import (
	// Register connectors
	_ "github.com/go-gost/x/connector/direct"
	_ "github.com/go-gost/x/connector/forward"
	_ "github.com/go-gost/x/connector/http"
	_ "github.com/go-gost/x/connector/http2"
	_ "github.com/go-gost/x/connector/relay"
	_ "github.com/go-gost/x/connector/sni"
	_ "github.com/go-gost/x/connector/socks/v4"
	_ "github.com/go-gost/x/connector/socks/v5"
	_ "github.com/go-gost/x/connector/ss"
	_ "github.com/go-gost/x/connector/ss/udp"
	_ "github.com/go-gost/x/connector/sshd"

	// Register dialers
	_ "github.com/go-gost/x/dialer/direct"
	_ "github.com/go-gost/x/dialer/dtls"
	_ "github.com/go-gost/x/dialer/ftcp"
	_ "github.com/go-gost/x/dialer/grpc"
	_ "github.com/go-gost/x/dialer/http2"
	_ "github.com/go-gost/x/dialer/http2/h2"
	_ "github.com/go-gost/x/dialer/http3"
	_ "github.com/go-gost/x/dialer/icmp"
	_ "github.com/go-gost/x/dialer/kcp"
	_ "github.com/go-gost/x/dialer/mtls"
	_ "github.com/go-gost/x/dialer/mws"
	_ "github.com/go-gost/x/dialer/obfs/http"
	_ "github.com/go-gost/x/dialer/obfs/tls"
	_ "github.com/go-gost/x/dialer/pht"
	_ "github.com/go-gost/x/dialer/quic"
	_ "github.com/go-gost/x/dialer/ssh"
	_ "github.com/go-gost/x/dialer/sshd"
	_ "github.com/go-gost/x/dialer/tcp"
	_ "github.com/go-gost/x/dialer/tls"
	_ "github.com/go-gost/x/dialer/udp"
	_ "github.com/go-gost/x/dialer/ws"

	// Register handlers
	_ "github.com/go-gost/x/handler/auto"
	_ "github.com/go-gost/x/handler/serial"
	_ "github.com/go-gost/x/handler/dns"
	_ "github.com/go-gost/x/handler/forward/local"
	_ "github.com/go-gost/x/handler/forward/remote"
	_ "github.com/go-gost/x/handler/http"
	_ "github.com/go-gost/x/handler/http2"
	_ "github.com/go-gost/x/handler/http3"
	_ "github.com/go-gost/x/handler/redirect/tcp"
	_ "github.com/go-gost/x/handler/redirect/udp"
	_ "github.com/go-gost/x/handler/relay"
	_ "github.com/go-gost/x/handler/sni"
	_ "github.com/go-gost/x/handler/socks/v4"
	_ "github.com/go-gost/x/handler/socks/v5"
	_ "github.com/go-gost/x/handler/ss"
	_ "github.com/go-gost/x/handler/ss/udp"
	_ "github.com/go-gost/x/handler/sshd"
	_ "github.com/go-gost/x/handler/tap"
	_ "github.com/go-gost/x/handler/tun"

	// Register listeners
	_ "github.com/go-gost/x/listener/serial"
	_ "github.com/go-gost/x/listener/dns"
	_ "github.com/go-gost/x/listener/dtls"
	_ "github.com/go-gost/x/listener/ftcp"
	_ "github.com/go-gost/x/listener/grpc"
	_ "github.com/go-gost/x/listener/http2"
	_ "github.com/go-gost/x/listener/http2/h2"
	_ "github.com/go-gost/x/listener/http3"
	_ "github.com/go-gost/x/listener/http3/h3"
	_ "github.com/go-gost/x/listener/icmp"
	_ "github.com/go-gost/x/listener/kcp"
	_ "github.com/go-gost/x/listener/mtls"
	_ "github.com/go-gost/x/listener/mws"
	_ "github.com/go-gost/x/listener/obfs/http"
	_ "github.com/go-gost/x/listener/obfs/tls"
	_ "github.com/go-gost/x/listener/pht"
	_ "github.com/go-gost/x/listener/quic"
	_ "github.com/go-gost/x/listener/redirect/tcp"
	_ "github.com/go-gost/x/listener/redirect/udp"
	_ "github.com/go-gost/x/listener/rtcp"
	_ "github.com/go-gost/x/listener/rudp"
	_ "github.com/go-gost/x/listener/ssh"
	_ "github.com/go-gost/x/listener/sshd"
	_ "github.com/go-gost/x/listener/tap"
	_ "github.com/go-gost/x/listener/tcp"
	_ "github.com/go-gost/x/listener/tls"
	_ "github.com/go-gost/x/listener/tun"
	_ "github.com/go-gost/x/listener/udp"
	_ "github.com/go-gost/x/listener/ws"
)
