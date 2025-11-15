package metrics

import (
	"math/big"

	"github.com/prometheus/client_golang/prometheus"
)

type SyncerMetricer interface {
	RecordChainBlockHeight(chainId string, height *big.Int)
	RecordEventBlockHeight(chainId string, height *big.Int)
	RecordNativeTokenBalance(chainId string, balance *big.Int)
	RecordChainAddressNonce(chainId string, nonce *big.Int)
}

type SyncerMetrics struct {
	chainBlockHeight   *prometheus.GaugeVec
	eventBlockHeight   *prometheus.GaugeVec
	nativeTokenBalance *prometheus.GaugeVec
	chainAddressNonce  *prometheus.GaugeVec
}

func NewSyncerMetrics(registry *prometheus.Registry, subsystem string) *SyncerMetrics {
	chainBlockHeight := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "chain_block_height",
		Help:      "Different chain block height",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	eventBlockHeight := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "event_block_height",
		Help:      "Different chain event block height",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	nativeTokenBalance := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "native_token_balance",
		Help:      "Different chain native token balance",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	chainAddressNonce := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "chain_address_nonce",
		Help:      "Different chain address nonce",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	registry.MustRegister(chainBlockHeight)
	registry.MustRegister(eventBlockHeight)
	registry.MustRegister(nativeTokenBalance)
	registry.MustRegister(chainAddressNonce)

	return &SyncerMetrics{
		chainBlockHeight:   chainBlockHeight,
		eventBlockHeight:   eventBlockHeight,
		nativeTokenBalance: nativeTokenBalance,
		chainAddressNonce:  chainAddressNonce,
	}
}

func (rm *SyncerMetrics) RecordChainBlockHeight(chainId string, height *big.Int) {
	rm.chainBlockHeight.WithLabelValues(chainId).Set(float64(height.Uint64()))
}

func (rm *SyncerMetrics) RecordEventBlockHeight(chainId string, height *big.Int) {
	rm.eventBlockHeight.WithLabelValues(chainId).Set(float64(height.Uint64()))
}

func (rm *SyncerMetrics) RecordNativeTokenBalance(chainId string, balance *big.Int) {
	rm.nativeTokenBalance.WithLabelValues(chainId).Set(float64(balance.Uint64()))
}

func (rm *SyncerMetrics) RecordChainAddressNonce(chainId string, nonce *big.Int) {
	rm.chainAddressNonce.WithLabelValues(chainId).Set(float64(nonce.Uint64()))
}
