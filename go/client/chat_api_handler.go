// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/keybase/client/go/protocol/chat1"
	"golang.org/x/net/context"
)

type ErrInvalidOptions struct {
	method  string
	version int
	err     error
}

func (e ErrInvalidOptions) Error() string {
	return fmt.Sprintf("invalid %s v%d options: %s", e.method, e.version, e.err)
}

type Call struct {
	Jsonrpc string
	ID      int
	Method  string
	Params  Params
}

type Params struct {
	Version int
	Options json.RawMessage
}

type CallError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Reply struct {
	Jsonrpc string      `json:"jsonrpc,omitempty"`
	ID      int         `json:"id,omitempty"`
	Error   *CallError  `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type APICallStatus struct {
	Code int    `json:"code"`
	Desc string `json:"desc,omitempty"`
}

type ChatAPIHandler interface {
	ListV1(context.Context, Call, io.Writer) error
	ReadV1(context.Context, Call, io.Writer) error
	SendV1(context.Context, Call, io.Writer) error
}

type ChatServiceHandler interface {
	ListV1(context.Context) Reply
	ReadV1(context.Context, readOptionsV1) Reply
	SendV1(context.Context, sendOptionsV1) Reply
}

type ChatAPI struct {
	svcHandler ChatServiceHandler
	indent     bool
}

type ChatChannel struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

func (c ChatChannel) Valid() bool {
	return len(c.Name) > 0
}

type ChatMessage struct {
	Body string
}

func (c ChatMessage) Valid() bool {
	return len(c.Body) > 0
}

type sendOptionsV1 struct {
	Channel ChatChannel
	Message ChatMessage
}

func (s sendOptionsV1) Check() error {
	if !s.Channel.Valid() {
		return ErrInvalidOptions{version: 1, method: "send", err: errors.New("invalid channel")}
	}
	if !s.Message.Valid() {
		return ErrInvalidOptions{version: 1, method: "send", err: errors.New("invalid message")}
	}

	return nil
}

type readOptionsV1 struct {
	Channel        ChatChannel
	ConversationID chat1.ConversationID `json:"conversation_id"`
	Limit          string
}

func (r readOptionsV1) Check() error {
	if !r.Channel.Valid() && r.ConversationID == 0 {
		return ErrInvalidOptions{version: 1, method: "read", err: errors.New("need channel or conversation_id")}
	}
	if r.Channel.Valid() && r.ConversationID > 0 {
		return ErrInvalidOptions{version: 1, method: "read", err: errors.New("include channel or conversation_id, not both")}
	}

	return nil
}

func (a *ChatAPI) ListV1(ctx context.Context, c Call, w io.Writer) error {
	if len(c.Params.Options) != 0 {
		return ErrInvalidOptions{version: 1, method: "list", err: errors.New("unexpected options, should be empty")}
	}

	return a.encodeReply(c, a.svcHandler.ListV1(ctx), w)
}

func (a *ChatAPI) ReadV1(ctx context.Context, c Call, w io.Writer) error {
	if len(c.Params.Options) == 0 {
		return ErrInvalidOptions{version: 1, method: "read", err: errors.New("empty options")}
	}
	var opts readOptionsV1
	if err := json.Unmarshal(c.Params.Options, &opts); err != nil {
		return err
	}
	if err := opts.Check(); err != nil {
		return err
	}

	// opts are valid for read v1

	return a.encodeReply(c, a.svcHandler.ReadV1(ctx, opts), w)
}

func (a *ChatAPI) SendV1(ctx context.Context, c Call, w io.Writer) error {
	if len(c.Params.Options) == 0 {
		return ErrInvalidOptions{version: 1, method: "send", err: errors.New("empty options")}
	}
	var opts sendOptionsV1
	if err := json.Unmarshal(c.Params.Options, &opts); err != nil {
		return err
	}
	if err := opts.Check(); err != nil {
		return err
	}

	// opts are valid for send v1

	return a.encodeReply(c, a.svcHandler.SendV1(ctx, opts), w)
}

func (a *ChatAPI) encodeReply(call Call, reply Reply, w io.Writer) error {
	// copy jsonrpc fields from call to reply
	reply.Jsonrpc = call.Jsonrpc
	reply.ID = call.ID

	enc := json.NewEncoder(w)
	if a.indent {
		enc.SetIndent("", "\t")
	}
	return enc.Encode(reply)
}
