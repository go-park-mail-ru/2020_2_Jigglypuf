// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(in *jlexer.Lexer, out *HallPlace) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Place":
			out.Place = int(in.Int())
		case "Row":
			out.Row = int(in.Int())
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
func easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(out *jwriter.Writer, in HallPlace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Place\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Place))
	}
	{
		const prefix string = ",\"Row\":"
		out.RawString(prefix)
		out.Int(int(in.Row))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HallPlace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HallPlace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HallPlace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HallPlace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(l, v)
}
func easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(in *jlexer.Lexer, out *HallConfig) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Levels":
			if in.IsNull() {
				in.Skip()
				out.Levels = nil
			} else {
				in.Delim('[')
				if out.Levels == nil {
					if !in.IsDelim(']') {
						out.Levels = make([]HallPlace, 0, 4)
					} else {
						out.Levels = []HallPlace{}
					}
				} else {
					out.Levels = (out.Levels)[:0]
				}
				for !in.IsDelim(']') {
					var v1 HallPlace
					(v1).UnmarshalEasyJSON(in)
					out.Levels = append(out.Levels, v1)
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
func easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(out *jwriter.Writer, in HallConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Levels\":"
		out.RawString(prefix[1:])
		if in.Levels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Levels {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HallConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HallConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HallConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HallConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(l, v)
}
func easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(in *jlexer.Lexer, out *CinemaHall) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint64(in.Uint64())
		case "PlaceAmount":
			out.PlaceAmount = int(in.Int())
		case "PlaceConfig":
			(out.PlaceConfig).UnmarshalEasyJSON(in)
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
func easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(out *jwriter.Writer, in CinemaHall) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"PlaceAmount\":"
		out.RawString(prefix)
		out.Int(int(in.PlaceAmount))
	}
	{
		const prefix string = ",\"PlaceConfig\":"
		out.RawString(prefix)
		(in.PlaceConfig).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CinemaHall) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CinemaHall) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b68e925EncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CinemaHall) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CinemaHall) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b68e925DecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(l, v)
}
