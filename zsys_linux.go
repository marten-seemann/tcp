// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_linux.go

package tcp

const (
	sysSIOCINQ  = 0x541b
	sysSIOCOUTQ = 0x5411

	sysTCP_CORK          = 0x3
	sysTCP_KEEPIDLE      = 0x4
	sysTCP_KEEPINTVL     = 0x5
	sysTCP_KEEPCNT       = 0x6
	sysTCP_INFO          = 0xb
	sysTCP_NOTSENT_LOWAT = 0x19

	sysTCPI_OPT_TIMESTAMPS = 0x1
	sysTCPI_OPT_SACK       = 0x2
	sysTCPI_OPT_WSCALE     = 0x4
	sysTCPI_OPT_ECN        = 0x8
	sysTCPI_OPT_ECN_SEEN   = 0x10
	sysTCPI_OPT_SYN_DATA   = 0x20

	CAOpen     CAState = 0x0
	CADisorder CAState = 0x1
	CACWR      CAState = 0x2
	CARecovery CAState = 0x3
	CALoss     CAState = 0x4

	sysSizeofTCPInfo = 0x90
)

type sysTCPInfo struct {
	State           uint8
	Ca_state        uint8
	Retransmits     uint8
	Probes          uint8
	Backoff         uint8
	Options         uint8
	Pad_cgo_0       [2]byte
	Rto             uint32
	Ato             uint32
	Snd_mss         uint32
	Rcv_mss         uint32
	Unacked         uint32
	Sacked          uint32
	Lost            uint32
	Retrans         uint32
	Fackets         uint32
	Last_data_sent  uint32
	Last_ack_sent   uint32
	Last_data_recv  uint32
	Last_ack_recv   uint32
	Pmtu            uint32
	Rcv_ssthresh    uint32
	Rtt             uint32
	Rttvar          uint32
	Snd_ssthresh    uint32
	Snd_cwnd        uint32
	Advmss          uint32
	Reordering      uint32
	Rcv_rtt         uint32
	Rcv_space       uint32
	Total_retrans   uint32
	Pacing_rate     uint64
	Max_pacing_rate uint64
	Bytes_acked     uint64
	Bytes_received  uint64
	Segs_out        uint32
	Segs_in         uint32
}
