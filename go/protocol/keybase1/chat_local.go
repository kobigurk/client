// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/chat_local.avdl

package keybase1

import (
	chat1 "github.com/keybase/client/go/protocol/chat1"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type MessageText struct {
	Body string `codec:"body" json:"body"`
}

type MessageConversationMetadata struct {
	ConversationTitle string `codec:"conversationTitle" json:"conversationTitle"`
}

type MessageEdit struct {
	MessageID chat1.MessageID `codec:"messageID" json:"messageID"`
	Body      string          `codec:"body" json:"body"`
}

type MessageDelete struct {
	MessageID chat1.MessageID `codec:"messageID" json:"messageID"`
}

type MessageAttachment struct {
	Path string `codec:"path" json:"path"`
}

type MessageBody struct {
	Type                 chat1.MessageType            `codec:"type" json:"type"`
	Text                 *MessageText                 `codec:"text,omitempty" json:"text,omitempty"`
	Attachment           *MessageAttachment           `codec:"attachment,omitempty" json:"attachment,omitempty"`
	Edit                 *MessageEdit                 `codec:"edit,omitempty" json:"edit,omitempty"`
	Delete               *MessageDelete               `codec:"delete,omitempty" json:"delete,omitempty"`
	ConversationMetadata *MessageConversationMetadata `codec:"conversationMetadata,omitempty" json:"conversationMetadata,omitempty"`
}

type MessagePlaintext struct {
	ClientHeader  chat1.MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	MessageBodies []MessageBody             `codec:"messageBodies" json:"messageBodies"`
}

type MessageInfoLocal struct {
	IsNew            bool   `codec:"isNew" json:"isNew"`
	SenderUsername   string `codec:"senderUsername" json:"senderUsername"`
	SenderDeviceName string `codec:"senderDeviceName" json:"senderDeviceName"`
}

type Message struct {
	ServerHeader     chat1.MessageServerHeader `codec:"serverHeader" json:"serverHeader"`
	MessagePlaintext MessagePlaintext          `codec:"messagePlaintext" json:"messagePlaintext"`
	Info             *MessageInfoLocal         `codec:"info,omitempty" json:"info,omitempty"`
}

type ThreadView struct {
	Messages   []Message         `codec:"messages" json:"messages"`
	Pagination *chat1.Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type MessageSelector struct {
	MessageTypes  []chat1.MessageType    `codec:"MessageTypes" json:"MessageTypes"`
	Since         *string                `codec:"Since,omitempty" json:"Since,omitempty"`
	OnlyNew       bool                   `codec:"onlyNew" json:"onlyNew"`
	Limit         int                    `codec:"limit" json:"limit"`
	Conversations []chat1.ConversationID `codec:"conversations" json:"conversations"`
	MarkAsRead    bool                   `codec:"markAsRead" json:"markAsRead"`
}

type ConversationInfoLocal struct {
	Id        chat1.ConversationID `codec:"id" json:"id"`
	TlfName   string               `codec:"tlfName" json:"tlfName"`
	TopicName string               `codec:"topicName" json:"topicName"`
	TopicType chat1.TopicType      `codec:"topicType" json:"topicType"`
}

type ResolvedConversationLocal struct {
	Conversation ConversationInfoLocal `codec:"conversation" json:"conversation"`
	Timestamp    Time                  `codec:"timestamp" json:"timestamp"`
}

type ConversationLocal struct {
	Id       chat1.ConversationID   `codec:"id" json:"id"`
	Info     *ConversationInfoLocal `codec:"info,omitempty" json:"info,omitempty"`
	Messages []Message              `codec:"messages" json:"messages"`
}

type GetInboxLocalArg struct {
	Pagination *chat1.Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetThreadLocalArg struct {
	ConversationID chat1.ConversationID `codec:"conversationID" json:"conversationID"`
	MarkAsRead     bool                 `codec:"markAsRead" json:"markAsRead"`
	Pagination     *chat1.Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostLocalArg struct {
	ConversationID   chat1.ConversationID `codec:"conversationID" json:"conversationID"`
	MessagePlaintext MessagePlaintext     `codec:"messagePlaintext" json:"messagePlaintext"`
}

type ResolveConversationLocalArg struct {
	Conversation ConversationInfoLocal `codec:"conversation" json:"conversation"`
}

type NewConversationLocalArg struct {
	Conversation ConversationInfoLocal `codec:"conversation" json:"conversation"`
}

type UpdateTopicNameLocalArg struct {
	ConversationID chat1.ConversationID `codec:"conversationID" json:"conversationID"`
	NewTopicName   string               `codec:"newTopicName" json:"newTopicName"`
}

type GetMessagesLocalArg struct {
	Selector MessageSelector `codec:"selector" json:"selector"`
}

type CompleteAndCanonicalizeTlfNameArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type ChatLocalInterface interface {
	GetInboxLocal(context.Context, *chat1.Pagination) (chat1.InboxView, error)
	GetThreadLocal(context.Context, GetThreadLocalArg) (ThreadView, error)
	PostLocal(context.Context, PostLocalArg) error
	ResolveConversationLocal(context.Context, ConversationInfoLocal) ([]ResolvedConversationLocal, error)
	NewConversationLocal(context.Context, ConversationInfoLocal) (ConversationInfoLocal, error)
	UpdateTopicNameLocal(context.Context, UpdateTopicNameLocalArg) error
	GetMessagesLocal(context.Context, MessageSelector) ([]ConversationLocal, error)
	CompleteAndCanonicalizeTlfName(context.Context, string) (CanonicalTlfName, error)
}

func ChatLocalProtocol(i ChatLocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.chatLocal",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxLocal(ctx, (*typedArgs)[0].Pagination)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadLocalArg)(nil), args)
						return
					}
					ret, err = i.GetThreadLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postLocal": {
				MakeArg: func() interface{} {
					ret := make([]PostLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostLocalArg)(nil), args)
						return
					}
					err = i.PostLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"resolveConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]ResolveConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ResolveConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]ResolveConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.ResolveConversationLocal(ctx, (*typedArgs)[0].Conversation)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.NewConversationLocal(ctx, (*typedArgs)[0].Conversation)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"updateTopicNameLocal": {
				MakeArg: func() interface{} {
					ret := make([]UpdateTopicNameLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UpdateTopicNameLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]UpdateTopicNameLocalArg)(nil), args)
						return
					}
					err = i.UpdateTopicNameLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesLocalArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesLocal(ctx, (*typedArgs)[0].Selector)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"completeAndCanonicalizeTlfName": {
				MakeArg: func() interface{} {
					ret := make([]CompleteAndCanonicalizeTlfNameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CompleteAndCanonicalizeTlfNameArg)
					if !ok {
						err = rpc.NewTypeError((*[]CompleteAndCanonicalizeTlfNameArg)(nil), args)
						return
					}
					ret, err = i.CompleteAndCanonicalizeTlfName(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ChatLocalClient struct {
	Cli rpc.GenericClient
}

func (c ChatLocalClient) GetInboxLocal(ctx context.Context, pagination *chat1.Pagination) (res chat1.InboxView, err error) {
	__arg := GetInboxLocalArg{Pagination: pagination}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getInboxLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) GetThreadLocal(ctx context.Context, __arg GetThreadLocalArg) (res ThreadView, err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getThreadLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) PostLocal(ctx context.Context, __arg PostLocalArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.postLocal", []interface{}{__arg}, nil)
	return
}

func (c ChatLocalClient) ResolveConversationLocal(ctx context.Context, conversation ConversationInfoLocal) (res []ResolvedConversationLocal, err error) {
	__arg := ResolveConversationLocalArg{Conversation: conversation}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.resolveConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) NewConversationLocal(ctx context.Context, conversation ConversationInfoLocal) (res ConversationInfoLocal, err error) {
	__arg := NewConversationLocalArg{Conversation: conversation}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.newConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) UpdateTopicNameLocal(ctx context.Context, __arg UpdateTopicNameLocalArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.updateTopicNameLocal", []interface{}{__arg}, nil)
	return
}

func (c ChatLocalClient) GetMessagesLocal(ctx context.Context, selector MessageSelector) (res []ConversationLocal, err error) {
	__arg := GetMessagesLocalArg{Selector: selector}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getMessagesLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) CompleteAndCanonicalizeTlfName(ctx context.Context, tlfName string) (res CanonicalTlfName, err error) {
	__arg := CompleteAndCanonicalizeTlfNameArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.completeAndCanonicalizeTlfName", []interface{}{__arg}, &res)
	return
}
