package basic

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *NewPushPacket) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "badge":
			z.Badge, err = dc.ReadString()
			if err != nil {
				return
			}
		case "data":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Data = nil
			} else {
				if z.Data == nil {
					z.Data = new(PushPacket)
				}
				err = z.Data.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewPushPacket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "badge"
	err = en.Append(0x82, 0xa5, 0x62, 0x61, 0x64, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Badge)
	if err != nil {
		return
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return err
	}
	if z.Data == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Data.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewPushPacket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "badge"
	o = append(o, 0x82, 0xa5, 0x62, 0x61, 0x64, 0x67, 0x65)
	o = msgp.AppendString(o, z.Badge)
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	if z.Data == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Data.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewPushPacket) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "badge":
			z.Badge, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "data":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Data = nil
			} else {
				if z.Data == nil {
					z.Data = new(PushPacket)
				}
				bts, err = z.Data.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewPushPacket) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Badge) + 5
	if z.Data == nil {
		s += msgp.NilSize
	} else {
		s += z.Data.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PushMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mt":
			z.MsgType, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "id":
			z.SessionID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "tc":
			z.Topic, err = dc.ReadString()
			if err != nil {
				return
			}
		case "m":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Message) >= int(zajw) {
				z.Message = (z.Message)[:zajw]
			} else {
				z.Message = make([][]byte, zajw)
			}
			for zbai := range z.Message {
				z.Message[zbai], err = dc.ReadBytes(z.Message[zbai])
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PushMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "mt"
	err = en.Append(0x84, 0xa2, 0x6d, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.MsgType)
	if err != nil {
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SessionID)
	if err != nil {
		return
	}
	// write "tc"
	err = en.Append(0xa2, 0x74, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Topic)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Message)))
	if err != nil {
		return
	}
	for zbai := range z.Message {
		err = en.WriteBytes(z.Message[zbai])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PushMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "mt"
	o = append(o, 0x84, 0xa2, 0x6d, 0x74)
	o = msgp.AppendInt(o, z.MsgType)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.SessionID)
	// string "tc"
	o = append(o, 0xa2, 0x74, 0x63)
	o = msgp.AppendString(o, z.Topic)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Message)))
	for zbai := range z.Message {
		o = msgp.AppendBytes(o, z.Message[zbai])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PushMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mt":
			z.MsgType, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "id":
			z.SessionID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "tc":
			z.Topic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "m":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Message) >= int(zhct) {
				z.Message = (z.Message)[:zhct]
			} else {
				z.Message = make([][]byte, zhct)
			}
			for zbai := range z.Message {
				z.Message[zbai], bts, err = msgp.ReadBytesBytes(bts, z.Message[zbai])
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PushMessage) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 3 + msgp.StringPrefixSize + len(z.SessionID) + 3 + msgp.StringPrefixSize + len(z.Topic) + 2 + msgp.ArrayHeaderSize
	for zbai := range z.Message {
		s += msgp.BytesPrefixSize + len(z.Message[zbai])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PushPacket) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "tts":
			var zdaf uint32
			zdaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.TTopics) >= int(zdaf) {
				z.TTopics = (z.TTopics)[:zdaf]
			} else {
				z.TTopics = make([]string, zdaf)
			}
			for zcua := range z.TTopics {
				z.TTopics[zcua], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "tpids":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.TOpIDs) >= int(zpks) {
				z.TOpIDs = (z.TOpIDs)[:zpks]
			} else {
				z.TOpIDs = make([]string, zpks)
			}
			for zxhx := range z.TOpIDs {
				z.TOpIDs[zxhx], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "fid":
			z.FromOpID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ftop":
			z.FromTopic, err = dc.ReadString()
			if err != nil {
				return
			}
		case "pt":
			z.PackType, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "mid":
			z.MsgID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "i":
			z.IosMsg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "e":
			z.Expire, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "l":
			z.Label, err = dc.ReadString()
			if err != nil {
				return
			}
		case "mt":
			z.MsgType, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "ms":
			z.Message, err = dc.ReadBytes(z.Message)
			if err != nil {
				return
			}
		case "ij":
			z.IosJSON, err = dc.ReadBytes(z.IosJSON)
			if err != nil {
				return
			}
		case "nk":
			z.Nick, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PushPacket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 13
	// write "tts"
	err = en.Append(0x8d, 0xa3, 0x74, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.TTopics)))
	if err != nil {
		return
	}
	for zcua := range z.TTopics {
		err = en.WriteString(z.TTopics[zcua])
		if err != nil {
			return
		}
	}
	// write "tpids"
	err = en.Append(0xa5, 0x74, 0x70, 0x69, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.TOpIDs)))
	if err != nil {
		return
	}
	for zxhx := range z.TOpIDs {
		err = en.WriteString(z.TOpIDs[zxhx])
		if err != nil {
			return
		}
	}
	// write "fid"
	err = en.Append(0xa3, 0x66, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FromOpID)
	if err != nil {
		return
	}
	// write "ftop"
	err = en.Append(0xa4, 0x66, 0x74, 0x6f, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FromTopic)
	if err != nil {
		return
	}
	// write "pt"
	err = en.Append(0xa2, 0x70, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PackType)
	if err != nil {
		return
	}
	// write "mid"
	err = en.Append(0xa3, 0x6d, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.MsgID)
	if err != nil {
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteString(z.IosMsg)
	if err != nil {
		return
	}
	// write "e"
	err = en.Append(0xa1, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Expire)
	if err != nil {
		return
	}
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Label)
	if err != nil {
		return
	}
	// write "mt"
	err = en.Append(0xa2, 0x6d, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.MsgType)
	if err != nil {
		return
	}
	// write "ms"
	err = en.Append(0xa2, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Message)
	if err != nil {
		return
	}
	// write "ij"
	err = en.Append(0xa2, 0x69, 0x6a)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.IosJSON)
	if err != nil {
		return
	}
	// write "nk"
	err = en.Append(0xa2, 0x6e, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nick)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PushPacket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "tts"
	o = append(o, 0x8d, 0xa3, 0x74, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TTopics)))
	for zcua := range z.TTopics {
		o = msgp.AppendString(o, z.TTopics[zcua])
	}
	// string "tpids"
	o = append(o, 0xa5, 0x74, 0x70, 0x69, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TOpIDs)))
	for zxhx := range z.TOpIDs {
		o = msgp.AppendString(o, z.TOpIDs[zxhx])
	}
	// string "fid"
	o = append(o, 0xa3, 0x66, 0x69, 0x64)
	o = msgp.AppendString(o, z.FromOpID)
	// string "ftop"
	o = append(o, 0xa4, 0x66, 0x74, 0x6f, 0x70)
	o = msgp.AppendString(o, z.FromTopic)
	// string "pt"
	o = append(o, 0xa2, 0x70, 0x74)
	o = msgp.AppendInt(o, z.PackType)
	// string "mid"
	o = append(o, 0xa3, 0x6d, 0x69, 0x64)
	o = msgp.AppendString(o, z.MsgID)
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendString(o, z.IosMsg)
	// string "e"
	o = append(o, 0xa1, 0x65)
	o = msgp.AppendInt64(o, z.Expire)
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendString(o, z.Label)
	// string "mt"
	o = append(o, 0xa2, 0x6d, 0x74)
	o = msgp.AppendInt(o, z.MsgType)
	// string "ms"
	o = append(o, 0xa2, 0x6d, 0x73)
	o = msgp.AppendBytes(o, z.Message)
	// string "ij"
	o = append(o, 0xa2, 0x69, 0x6a)
	o = msgp.AppendBytes(o, z.IosJSON)
	// string "nk"
	o = append(o, 0xa2, 0x6e, 0x6b)
	o = msgp.AppendString(o, z.Nick)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PushPacket) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "tts":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TTopics) >= int(zcxo) {
				z.TTopics = (z.TTopics)[:zcxo]
			} else {
				z.TTopics = make([]string, zcxo)
			}
			for zcua := range z.TTopics {
				z.TTopics[zcua], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "tpids":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TOpIDs) >= int(zeff) {
				z.TOpIDs = (z.TOpIDs)[:zeff]
			} else {
				z.TOpIDs = make([]string, zeff)
			}
			for zxhx := range z.TOpIDs {
				z.TOpIDs[zxhx], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "fid":
			z.FromOpID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ftop":
			z.FromTopic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "pt":
			z.PackType, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "mid":
			z.MsgID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "i":
			z.IosMsg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "e":
			z.Expire, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "l":
			z.Label, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "mt":
			z.MsgType, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ms":
			z.Message, bts, err = msgp.ReadBytesBytes(bts, z.Message)
			if err != nil {
				return
			}
		case "ij":
			z.IosJSON, bts, err = msgp.ReadBytesBytes(bts, z.IosJSON)
			if err != nil {
				return
			}
		case "nk":
			z.Nick, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PushPacket) Msgsize() (s int) {
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zcua := range z.TTopics {
		s += msgp.StringPrefixSize + len(z.TTopics[zcua])
	}
	s += 6 + msgp.ArrayHeaderSize
	for zxhx := range z.TOpIDs {
		s += msgp.StringPrefixSize + len(z.TOpIDs[zxhx])
	}
	s += 4 + msgp.StringPrefixSize + len(z.FromOpID) + 5 + msgp.StringPrefixSize + len(z.FromTopic) + 3 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.MsgID) + 2 + msgp.StringPrefixSize + len(z.IosMsg) + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Label) + 3 + msgp.IntSize + 3 + msgp.BytesPrefixSize + len(z.Message) + 3 + msgp.BytesPrefixSize + len(z.IosJSON) + 3 + msgp.StringPrefixSize + len(z.Nick)
	return
}
