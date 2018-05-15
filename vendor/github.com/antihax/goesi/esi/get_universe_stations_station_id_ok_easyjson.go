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

func easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseStationsStationIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStationsStationIdOkList, 0, 1)
			} else {
				*out = GetUniverseStationsStationIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStationsStationIdOk
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
func easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseStationsStationIdOkList) {
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
func (v GetUniverseStationsStationIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStationsStationIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStationsStationIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStationsStationIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseStationsStationIdOk) {
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
		case "station_id":
			out.StationId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "owner":
			out.Owner = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "race_id":
			out.RaceId = int32(in.Int32())
		case "position":
			easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi2(in, &out.Position)
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "reprocessing_efficiency":
			out.ReprocessingEfficiency = float32(in.Float32())
		case "reprocessing_stations_take":
			out.ReprocessingStationsTake = float32(in.Float32())
		case "max_dockable_ship_volume":
			out.MaxDockableShipVolume = float32(in.Float32())
		case "office_rental_cost":
			out.OfficeRentalCost = float32(in.Float32())
		case "services":
			if in.IsNull() {
				in.Skip()
				out.Services = nil
			} else {
				in.Delim('[')
				if out.Services == nil {
					if !in.IsDelim(']') {
						out.Services = make([]string, 0, 4)
					} else {
						out.Services = []string{}
					}
				} else {
					out.Services = (out.Services)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Services = append(out.Services, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseStationsStationIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StationId != 0 {
		const prefix string = ",\"station_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.StationId))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Owner != 0 {
		const prefix string = ",\"owner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Owner))
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
	if in.RaceId != 0 {
		const prefix string = ",\"race_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RaceId))
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi2(out, in.Position)
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	if in.ReprocessingEfficiency != 0 {
		const prefix string = ",\"reprocessing_efficiency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.ReprocessingEfficiency))
	}
	if in.ReprocessingStationsTake != 0 {
		const prefix string = ",\"reprocessing_stations_take\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.ReprocessingStationsTake))
	}
	if in.MaxDockableShipVolume != 0 {
		const prefix string = ",\"max_dockable_ship_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.MaxDockableShipVolume))
	}
	if in.OfficeRentalCost != 0 {
		const prefix string = ",\"office_rental_cost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.OfficeRentalCost))
	}
	if len(in.Services) != 0 {
		const prefix string = ",\"services\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Services {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseStationsStationIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStationsStationIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStationsStationIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStationsStationIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonBde2b678DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetUniverseStationsStationIdPosition) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjsonBde2b678EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetUniverseStationsStationIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}
