// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson979d6e15DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdOrders200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdOrders200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdOrders200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdOrders200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson979d6e15EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdOrders200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdOrders200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson979d6e15EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdOrders200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson979d6e15EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdOrders200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson979d6e15DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdOrders200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson979d6e15DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson979d6e15DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdOrders200Ok) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "order_id":
			out.OrderId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "region_id":
			out.RegionId = int32(in.Int32())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "range":
			out.Range_ = string(in.String())
		case "price":
			out.Price = float64(in.Float64())
		case "volume_total":
			out.VolumeTotal = int32(in.Int32())
		case "volume_remain":
			out.VolumeRemain = int32(in.Int32())
		case "issued":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Issued).UnmarshalJSON(data))
			}
		case "is_buy_order":
			out.IsBuyOrder = bool(in.Bool())
		case "min_volume":
			out.MinVolume = int32(in.Int32())
		case "escrow":
			out.Escrow = float64(in.Float64())
		case "duration":
			out.Duration = int32(in.Int32())
		case "wallet_division":
			out.WalletDivision = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson979d6e15EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdOrders200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OrderId != 0 {
		const prefix string = ",\"order_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrderId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.RegionId != 0 {
		const prefix string = ",\"region_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RegionId))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.Range_ != "" {
		const prefix string = ",\"range\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Range_))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Price))
	}
	if in.VolumeTotal != 0 {
		const prefix string = ",\"volume_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeTotal))
	}
	if in.VolumeRemain != 0 {
		const prefix string = ",\"volume_remain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeRemain))
	}
	if true {
		const prefix string = ",\"issued\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Issued).MarshalJSON())
	}
	if in.IsBuyOrder {
		const prefix string = ",\"is_buy_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsBuyOrder))
	}
	if in.MinVolume != 0 {
		const prefix string = ",\"min_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MinVolume))
	}
	if in.Escrow != 0 {
		const prefix string = ",\"escrow\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Escrow))
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Duration))
	}
	if in.WalletDivision != 0 {
		const prefix string = ",\"wallet_division\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.WalletDivision))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdOrders200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson979d6e15EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdOrders200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson979d6e15EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdOrders200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson979d6e15DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdOrders200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson979d6e15DecodeGithubComAntihaxGoesiEsi1(l, v)
}
