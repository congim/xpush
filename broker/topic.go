package broker

//type TopicType string

//const (
//	GroupType      TopicType = "g"
//	PrivateType    TopicType = "p"
//	TopicSeparator string    = "/"
//)

//Topic topic订阅情况，用来映射conns
//type Topic struct {
//	topic string
//	conns sync.Map
//}

//
//// topic的详细信息，用来在conn映射topic
//type TopicInfo struct {
//	Type  TopicType
//	Topic string
//}
//
//func topicType(topic string) (TopicType, error) {
//	parts := strings.Split(topic, TopicSeparator)
//	if len(parts) != 3 {
//		return "", fmt.Errorf("illegal topic, topic is %s", topic)
//	}
//
//	switch TopicType(parts[1]) {
//	case GroupType:
//		return GroupType, nil
//	case PrivateType:
//		return PrivateType, nil
//	default:
//		return "", fmt.Errorf("illegal topic type, type is %s, topic is %s", parts[1], topic)
//	}
//}

//
//var _ Topicer = (*topicer)(nil)
//
//type Topicer interface {
//	Subscribe(string) error
//	UnSubscribe(string) error
//	OnMessage(message *message.Message) error
//}
//
//type topicer struct {
//	conns sync.Map
//}
//
//func (t *topicer) Subscribe(string) error {
//	return nil
//}
//
//func (t *topicer) UnSubscribe(string) error {
//	return nil
//}
//
//func (t *topicer) OnMessage(message *message.Message) error {
//	return nil
//}
