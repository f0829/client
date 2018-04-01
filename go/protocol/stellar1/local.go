// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/local.avdl

package stellar1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type BalancesLocalArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type SendLocalArg struct {
	Recipient string `codec:"recipient" json:"recipient"`
	Amount    string `codec:"amount" json:"amount"`
	Asset     Asset  `codec:"asset" json:"asset"`
	Note      string `codec:"note" json:"note"`
}

type WalletInitLocalArg struct {
}

type WalletDumpLocalArg struct {
}

type OwnAccountLocalArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
}

type ImportSecretKeyLocalArg struct {
	SecretKey   SecretKey `codec:"secretKey" json:"secretKey"`
	MakePrimary bool      `codec:"makePrimary" json:"makePrimary"`
}

type LocalInterface interface {
	BalancesLocal(context.Context, AccountID) ([]Balance, error)
	SendLocal(context.Context, SendLocalArg) (PaymentResult, error)
	WalletInitLocal(context.Context) error
	WalletDumpLocal(context.Context) (Bundle, error)
	OwnAccountLocal(context.Context, AccountID) (bool, error)
	ImportSecretKeyLocal(context.Context, ImportSecretKeyLocalArg) error
}

func LocalProtocol(i LocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "stellar.1.local",
		Methods: map[string]rpc.ServeHandlerDescription{
			"balancesLocal": {
				MakeArg: func() interface{} {
					ret := make([]BalancesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BalancesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]BalancesLocalArg)(nil), args)
						return
					}
					ret, err = i.BalancesLocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"sendLocal": {
				MakeArg: func() interface{} {
					ret := make([]SendLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SendLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]SendLocalArg)(nil), args)
						return
					}
					ret, err = i.SendLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"walletInitLocal": {
				MakeArg: func() interface{} {
					ret := make([]WalletInitLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.WalletInitLocal(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"walletDumpLocal": {
				MakeArg: func() interface{} {
					ret := make([]WalletDumpLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.WalletDumpLocal(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ownAccountLocal": {
				MakeArg: func() interface{} {
					ret := make([]OwnAccountLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]OwnAccountLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]OwnAccountLocalArg)(nil), args)
						return
					}
					ret, err = i.OwnAccountLocal(ctx, (*typedArgs)[0].AccountID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"importSecretKeyLocal": {
				MakeArg: func() interface{} {
					ret := make([]ImportSecretKeyLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ImportSecretKeyLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ImportSecretKeyLocalArg)(nil), args)
						return
					}
					err = i.ImportSecretKeyLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LocalClient struct {
	Cli rpc.GenericClient
}

func (c LocalClient) BalancesLocal(ctx context.Context, accountID AccountID) (res []Balance, err error) {
	__arg := BalancesLocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.balancesLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) SendLocal(ctx context.Context, __arg SendLocalArg) (res PaymentResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.sendLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) WalletInitLocal(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.walletInitLocal", []interface{}{WalletInitLocalArg{}}, nil)
	return
}

func (c LocalClient) WalletDumpLocal(ctx context.Context) (res Bundle, err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.walletDumpLocal", []interface{}{WalletDumpLocalArg{}}, &res)
	return
}

func (c LocalClient) OwnAccountLocal(ctx context.Context, accountID AccountID) (res bool, err error) {
	__arg := OwnAccountLocalArg{AccountID: accountID}
	err = c.Cli.Call(ctx, "stellar.1.local.ownAccountLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) ImportSecretKeyLocal(ctx context.Context, __arg ImportSecretKeyLocalArg) (err error) {
	err = c.Cli.Call(ctx, "stellar.1.local.importSecretKeyLocal", []interface{}{__arg}, nil)
	return
}