package coordinator

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_beacon"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_coordinator"
)

type topics struct {
	randomnessRequestedTopic            common.Hash
	randomnessFulfillmentRequestedTopic common.Hash
	randomWordsFulfilledTopic           common.Hash
	configSetTopic                      common.Hash
	newTransmissionTopic                common.Hash
}

func newTopics() topics {
	return topics{
		randomnessRequestedTopic:            vrf_coordinator.VRFCoordinatorRandomnessRequested{}.Topic(),
		randomnessFulfillmentRequestedTopic: vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested{}.Topic(),
		randomWordsFulfilledTopic:           vrf_coordinator.VRFCoordinatorRandomWordsFulfilled{}.Topic(),
		configSetTopic:                      vrf_beacon.VRFBeaconConfigSet{}.Topic(),
		newTransmissionTopic:                vrf_beacon.VRFBeaconNewTransmission{}.Topic(),
	}
}
