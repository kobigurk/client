// Auto-generated by avdl-compiler v1.3.7 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/tlf.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type CryptKeysArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type TlfInterface interface {
	// CryptKeys returns TLF crypt keys from all generations.
	CryptKeys(context.Context, string) (TLFCryptKeys, error)
}

func TlfProtocol(i TlfInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.tlf",
		Methods: map[string]rpc.ServeHandlerDescription{
			"CryptKeys": {
				MakeArg: func() interface{} {
					ret := make([]CryptKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CryptKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]CryptKeysArg)(nil), args)
						return
					}
					ret, err = i.CryptKeys(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TlfClient struct {
	Cli rpc.GenericClient
}

// CryptKeys returns TLF crypt keys from all generations.
func (c TlfClient) CryptKeys(ctx context.Context, tlfName string) (res TLFCryptKeys, err error) {
	__arg := CryptKeysArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlf.CryptKeys", []interface{}{__arg}, &res)
	return
}
