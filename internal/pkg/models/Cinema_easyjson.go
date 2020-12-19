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

func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(in *jlexer.Lexer, out *TicketPlace) {
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
		case "row":
			out.Row = int(in.Int())
		case "place":
			out.Place = int(in.Int())
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(out *jwriter.Writer, in TicketPlace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"row\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Row))
	}
	{
		const prefix string = ",\"place\":"
		out.RawString(prefix)
		out.Int(int(in.Place))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TicketPlace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TicketPlace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TicketPlace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TicketPlace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(l, v)
}
func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(in *jlexer.Lexer, out *TicketInput) {
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
		case "login":
			out.Login = string(in.String())
		case "scheduleID":
			out.ScheduleID = uint64(in.Uint64())
		case "placeField":
			if in.IsNull() {
				in.Skip()
				out.PlaceField = nil
			} else {
				in.Delim('[')
				if out.PlaceField == nil {
					if !in.IsDelim(']') {
						out.PlaceField = make([]TicketPlace, 0, 4)
					} else {
						out.PlaceField = []TicketPlace{}
					}
				} else {
					out.PlaceField = (out.PlaceField)[:0]
				}
				for !in.IsDelim(']') {
					var v1 TicketPlace
					(v1).UnmarshalEasyJSON(in)
					out.PlaceField = append(out.PlaceField, v1)
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(out *jwriter.Writer, in TicketInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"scheduleID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ScheduleID))
	}
	{
		const prefix string = ",\"placeField\":"
		out.RawString(prefix)
		if in.PlaceField == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.PlaceField {
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
func (v TicketInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TicketInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TicketInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TicketInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(l, v)
}
func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(in *jlexer.Lexer, out *Ticket) {
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
		case "Login":
			out.Login = string(in.String())
		case "Schedule":
			(out.Schedule).UnmarshalEasyJSON(in)
		case "TransactionDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.TransactionDate).UnmarshalJSON(data))
			}
		case "PlaceField":
			(out.PlaceField).UnmarshalEasyJSON(in)
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(out *jwriter.Writer, in Ticket) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"Login\":"
		out.RawString(prefix)
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"Schedule\":"
		out.RawString(prefix)
		(in.Schedule).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"TransactionDate\":"
		out.RawString(prefix)
		out.Raw((in.TransactionDate).MarshalJSON())
	}
	{
		const prefix string = ",\"PlaceField\":"
		out.RawString(prefix)
		(in.PlaceField).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ticket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ticket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ticket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ticket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(l, v)
}
func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(in *jlexer.Lexer, out *SearchCinema) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(out *jwriter.Writer, in SearchCinema) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchCinema) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchCinema) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchCinema) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchCinema) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels3(l, v)
}
func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(in *jlexer.Lexer, out *GetCinemaList) {
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
		case "Limit":
			out.Limit = int(in.Int())
		case "Page":
			out.Page = int(in.Int())
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(out *jwriter.Writer, in GetCinemaList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Limit\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Limit))
	}
	{
		const prefix string = ",\"Page\":"
		out.RawString(prefix)
		out.Int(int(in.Page))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCinemaList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCinemaList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCinemaList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCinemaList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels4(l, v)
}
func easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(in *jlexer.Lexer, out *Cinema) {
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
		case "Name":
			out.Name = string(in.String())
		case "Address":
			out.Address = string(in.String())
		case "HallCount":
			out.HallCount = int(in.Int())
		case "PathToAvatar":
			out.PathToAvatar = string(in.String())
		case "AuthorID":
			out.AuthorID = uint64(in.Uint64())
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
func easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(out *jwriter.Writer, in Cinema) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"HallCount\":"
		out.RawString(prefix)
		out.Int(int(in.HallCount))
	}
	{
		const prefix string = ",\"PathToAvatar\":"
		out.RawString(prefix)
		out.String(string(in.PathToAvatar))
	}
	{
		const prefix string = ",\"AuthorID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.AuthorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Cinema) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cinema) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316958ddEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cinema) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cinema) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316958ddDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels5(l, v)
}
