package basic

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BatchCheckMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Msgs":
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zbai) {
				z.Msgs = (z.Msgs)[:zbai]
			} else {
				z.Msgs = make([]*CheckUsersMsg, zbai)
			}
			for zxvk := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zxvk] = nil
				} else {
					if z.Msgs[zxvk] == nil {
						z.Msgs[zxvk] = new(CheckUsersMsg)
					}
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
						case "System":
							z.Msgs[zxvk].System, err = dc.ReadInt()
							if err != nil {
								return
							}
						case "Session":
							z.Msgs[zxvk].Session, err = dc.ReadString()
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
func (z *BatchCheckMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Msgs"
	err = en.Append(0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Msgs)))
	if err != nil {
		return
	}
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "System"
			err = en.Append(0x82, 0xa6, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d)
			if err != nil {
				return err
			}
			err = en.WriteInt(z.Msgs[zxvk].System)
			if err != nil {
				return
			}
			// write "Session"
			err = en.Append(0xa7, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Msgs[zxvk].Session)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BatchCheckMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Msgs"
	o = append(o, 0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "System"
			o = append(o, 0x82, 0xa6, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d)
			o = msgp.AppendInt(o, z.Msgs[zxvk].System)
			// string "Session"
			o = append(o, 0xa7, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e)
			o = msgp.AppendString(o, z.Msgs[zxvk].Session)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchCheckMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Msgs":
			var zwht uint32
			zwht, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zwht) {
				z.Msgs = (z.Msgs)[:zwht]
			} else {
				z.Msgs = make([]*CheckUsersMsg, zwht)
			}
			for zxvk := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zxvk] = nil
				} else {
					if z.Msgs[zxvk] == nil {
						z.Msgs[zxvk] = new(CheckUsersMsg)
					}
					var zhct uint32
					zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zhct > 0 {
						zhct--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "System":
							z.Msgs[zxvk].System, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Session":
							z.Msgs[zxvk].Session, bts, err = msgp.ReadStringBytes(bts)
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
func (z *BatchCheckMsg) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 7 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.Msgs[zxvk].Session)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CheckUsersMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "System":
			z.System, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Session":
			z.Session, err = dc.ReadString()
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
func (z CheckUsersMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "System"
	err = en.Append(0x82, 0xa6, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.System)
	if err != nil {
		return
	}
	// write "Session"
	err = en.Append(0xa7, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Session)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CheckUsersMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "System"
	o = append(o, 0x82, 0xa6, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d)
	o = msgp.AppendInt(o, z.System)
	// string "Session"
	o = append(o, 0xa7, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Session)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CheckUsersMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "System":
			z.System, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Session":
			z.Session, bts, err = msgp.ReadStringBytes(bts)
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
func (z CheckUsersMsg) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.Session)
	return
}
