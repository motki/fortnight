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

func easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetIndustryFacilities200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetIndustryFacilities200OkList, 0, 2)
			} else {
				*out = GetIndustryFacilities200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetIndustryFacilities200Ok
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
func easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetIndustryFacilities200OkList) {
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
func (v GetIndustryFacilities200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIndustryFacilities200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIndustryFacilities200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIndustryFacilities200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetIndustryFacilities200Ok) {
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
		case "facility_id":
			out.FacilityId = int64(in.Int64())
		case "tax":
			out.Tax = float32(in.Float32())
		case "owner_id":
			out.OwnerId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "region_id":
			out.RegionId = int32(in.Int32())
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
func easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetIndustryFacilities200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FacilityId != 0 {
		const prefix string = ",\"facility_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FacilityId))
	}
	if in.Tax != 0 {
		const prefix string = ",\"tax\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Tax))
	}
	if in.OwnerId != 0 {
		const prefix string = ",\"owner_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OwnerId))
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
	if in.SolarSystemId != 0 {
		const prefix string = ",\"solar_system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SolarSystemId))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetIndustryFacilities200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIndustryFacilities200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1840f9EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIndustryFacilities200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIndustryFacilities200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1840f9DecodeGithubComAntihaxGoesiEsi1(l, v)
}
