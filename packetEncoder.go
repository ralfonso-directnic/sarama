package kafka

type packetEncoder interface {
	putInt16(in int16)
	putInt32(in int32)
	putError(in kafkaError)
	putString(in *string)
	putBytes(in *[]byte)
	putArrayCount(in int)
}