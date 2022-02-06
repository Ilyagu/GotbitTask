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

func easyjson79a0a577DecodeGotbittaskInternalAppTaskModels(in *jlexer.Lexer, out *Task) {
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
		case "id":
			out.Id = int64(in.Int64())
		case "creator":
			out.Creator = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "executor":
			out.Executor = string(in.String())
		case "created":
			out.Created = string(in.String())
		case "completed":
			out.Completed = bool(in.Bool())
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
func easyjson79a0a577EncodeGotbittaskInternalAppTaskModels(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Id))
	}
	if in.Creator != "" {
		const prefix string = ",\"creator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Creator))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Executor != "" {
		const prefix string = ",\"executor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Executor))
	}
	if in.Created != "" {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Created))
	}
	{
		const prefix string = ",\"completed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Completed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Task) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson79a0a577EncodeGotbittaskInternalAppTaskModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79a0a577EncodeGotbittaskInternalAppTaskModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Task) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson79a0a577DecodeGotbittaskInternalAppTaskModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79a0a577DecodeGotbittaskInternalAppTaskModels(l, v)
}