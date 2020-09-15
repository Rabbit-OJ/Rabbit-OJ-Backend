package channel

var (
	JudgeRequestDeliveryChan  chan []byte
	JudgeRequeueDeliveryChan  chan []byte
	JudgeResponseDeliveryChan chan []byte

	JudgeRequestBridgeChan  chan *JudgeRequestBridgeMessage
	MQPublishMessageChannel chan *MQMessage
)

type JudgeRequestBridgeMessage struct {
	Data        []byte
	SuccessChan chan bool
}

type MQMessage struct {
	Async bool
	Topic []string
	Key   []byte
	Value []byte
}
