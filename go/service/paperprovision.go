// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package service

import (
	"github.com/keybase/client/go/engine"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	"golang.org/x/net/context"
)

type PaperProvisionHandler struct {
	*BaseHandler
	libkb.Contextified
}

func NewPaperProvisionHandler(xp rpc.Transporter, g *libkb.GlobalContext) *PaperProvisionHandler {
	return &PaperProvisionHandler{
		BaseHandler:  NewBaseHandler(xp),
		Contextified: libkb.NewContextified(g),
	}
}

func (h *PaperProvisionHandler) PaperProvision(_ context.Context, arg keybase1.PaperProvisionArg) error {

	ctx := engine.Context{
		LogUI:       h.getLogUI(arg.SessionID),
		SecretUI:    h.getSecretUI(arg.SessionID, h.G()),
		LoginUI:     h.getLoginUI(arg.SessionID),
		ProvisionUI: h.getProvisionUI(arg.SessionID),
		SessionID:   arg.SessionID,
	}
	eng := engine.NewPaperProvisionEngine(h.G(), arg.Username, arg.DeviceName, arg.PaperKey)
	err := engine.RunEngine(eng, &ctx)
	if err != nil {
		return err
	}
	return nil
}