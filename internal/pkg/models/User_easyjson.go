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

func easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(in *jlexer.Lexer, out *User) {
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
		case "Login":
			out.Login = string(in.String())
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
func easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Login\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Login))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels(l, v)
}
func easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(in *jlexer.Lexer, out *RegistrationInput) {
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
		case "password":
			out.Password = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
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
func easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(out *jwriter.Writer, in RegistrationInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RegistrationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RegistrationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RegistrationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RegistrationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels1(l, v)
}
func easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(in *jlexer.Lexer, out *AuthInput) {
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
		case "password":
			out.Password = string(in.String())
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
func easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(out *jwriter.Writer, in AuthInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson750e3c9dEncodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson750e3c9dDecodeGithubComGoParkMailRu20202JigglypufInternalPkgModels2(l, v)
}