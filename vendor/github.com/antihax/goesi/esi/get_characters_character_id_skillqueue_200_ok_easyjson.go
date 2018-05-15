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

func easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillqueue200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdSkillqueue200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdSkillqueue200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdSkillqueue200Ok
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
func easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdSkillqueue200OkList) {
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
func (v GetCharactersCharacterIdSkillqueue200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillqueue200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillqueue200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillqueue200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillqueue200Ok) {
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
		case "skill_id":
			out.SkillId = int32(in.Int32())
		case "finish_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.FinishDate).UnmarshalJSON(data))
			}
		case "start_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartDate).UnmarshalJSON(data))
			}
		case "finished_level":
			out.FinishedLevel = int32(in.Int32())
		case "queue_position":
			out.QueuePosition = int32(in.Int32())
		case "training_start_sp":
			out.TrainingStartSp = int32(in.Int32())
		case "level_end_sp":
			out.LevelEndSp = int32(in.Int32())
		case "level_start_sp":
			out.LevelStartSp = int32(in.Int32())
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
func easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdSkillqueue200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SkillId != 0 {
		const prefix string = ",\"skill_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SkillId))
	}
	if true {
		const prefix string = ",\"finish_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.FinishDate).MarshalJSON())
	}
	if true {
		const prefix string = ",\"start_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.StartDate).MarshalJSON())
	}
	if in.FinishedLevel != 0 {
		const prefix string = ",\"finished_level\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FinishedLevel))
	}
	if in.QueuePosition != 0 {
		const prefix string = ",\"queue_position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.QueuePosition))
	}
	if in.TrainingStartSp != 0 {
		const prefix string = ",\"training_start_sp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TrainingStartSp))
	}
	if in.LevelEndSp != 0 {
		const prefix string = ",\"level_end_sp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LevelEndSp))
	}
	if in.LevelStartSp != 0 {
		const prefix string = ",\"level_start_sp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LevelStartSp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdSkillqueue200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillqueue200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbf16696EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillqueue200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillqueue200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbf16696DecodeGithubComAntihaxGoesiEsi1(l, v)
}
