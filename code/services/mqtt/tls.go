package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"serveos-datasync/config"
)

func NewTLSConfig(cfg config.Config) *tls.Config {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: !cfg.ServerMQTTValidateCert,
	}
	if cfg.ServerMQTTEnableTLS {
		certpool := x509.NewCertPool()
		pemCerts, err := os.ReadFile(cfg.ServerMQTTCAPath)
		if err == nil {
			certpool.AppendCertsFromPEM(pemCerts)
		}
		tlsConfig.RootCAs = certpool
	}
	return tlsConfig
}
