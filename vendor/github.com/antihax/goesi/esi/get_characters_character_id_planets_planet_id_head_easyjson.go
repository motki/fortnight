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

func easyjson45497e3DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdHeadList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdHeadList, 0, 5)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdHeadList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdHead
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
func easyjson45497e3EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdHeadList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdHeadList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson45497e3EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdHeadList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson45497e3EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdHeadList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson45497e3DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdHeadList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson45497e3DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson45497e3DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdHead) {
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
		case "head_id":
			out.HeadId = int32(in.Int32())
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
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
func easyjson45497e3EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdHead) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HeadId != 0 {
		const prefix string = ",\"head_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.HeadId))
	}
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Longitude))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdHead) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson45497e3EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdHead) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson45497e3EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdHead) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson45497e3DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdHead) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson45497e3DecodeGithubComAntihaxGoesiEsi1(l, v)
}