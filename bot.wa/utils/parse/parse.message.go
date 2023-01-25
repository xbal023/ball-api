package parseMessage

import (
	"os"
	"time"
	"strings"

	"go.mau.fi/whatsmeow/types/events"
	"go.mau.fi/whatsmeow/types"
	x "github.com/bolaxd/dumn/bot.wa/utils/simple"
	)
var (
	prefix = "."
	noown = os.Getenv("OWNER_NUMBER") + "@s.whatsapp.net"
	nobot = os.Getenv("BOT_NUMBER") + "@s.whatsapp.net"
)

type QuotedStruct struct {
	Id string
	Sender *types.JID
	Body string
	Media x.DownloadableMessage
	TypeM string
}

type Parse struct {
	Id string
	Chat *types.JID
	Sender *types.JID 
	Pushname string
	Timestamp *time.Time
	Body string
	Media x.DownloadableMessage
	TypeM string
	Quoted QuotedStruct
	Pref string
	Cmd string
	CmdP string
	Query string
	IsGc bool 
	IsMe bool
	IsOwn bool
	IsAdmin bool
	IsBotAdmin bool
}

func Parser(ball *x.S, up *events.Message) *Parse {
	chat := types.JID(up.Info.Chat)
	sender := types.JID(up.Info.Sender)
	senderQ, _ := types.ParseJID(up.Message.GetExtendedTextMessage().GetContextInfo().GetParticipant())
	times := time.Time(up.Info.Timestamp)
	extendedText := up.Message.GetExtendedTextMessage().GetText()
	conversationText := up.Message.GetConversation()
	imageMessageText := up.Message.GetImageMessage().GetCaption()
	videoMessageText := up.Message.GetVideoMessage().GetCaption()
	buttonMessageText := up.Message.GetTemplateButtonReplyMessage().GetSelectedId()
	var body string
	if conversationText != "" {
		body = conversationText
	} else if extendedText != "" {
		body = extendedText
	} else if imageMessageText != "" {
		body = imageMessageText
	} else if videoMessageText != "" {
		body = videoMessageText
	} else if buttonMessageText != "" {
		body = buttonMessageText
	}
	var cmds string
	if conversationText != "" && strings.HasPrefix(conversationText, prefix) {
		cmds = conversationText
	} else if extendedText != "" && strings.HasPrefix(extendedText, prefix) {
		cmds = extendedText
	} else if imageMessageText != "" && strings.HasPrefix(imageMessageText, prefix) {
		cmds = imageMessageText
	} else if videoMessageText != "" && strings.HasPrefix(videoMessageText, prefix) {
		cmds = videoMessageText
	} else if buttonMessageText != "" && strings.HasPrefix(buttonMessageText, prefix) {
		cmds = buttonMessageText
	}
	extendedQuotedText := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetExtendedTextMessage().GetText()
	conversationQuotedText := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetConversation()
	imageQuotedText := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetImageMessage().GetCaption()
	videoQuotedText := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetVideoMessage().GetCaption()
	var quotedBody string 
	if conversationQuotedText != "" {
		quotedBody = conversationQuotedText
	} else if extendedQuotedText != "" {
		quotedBody = extendedQuotedText
	} else if imageQuotedText != "" {
		quotedBody = imageQuotedText
	} else if videoQuotedText != "" {
		quotedBody = videoQuotedText
	}
	imageMessage := up.Message.GetImageMessage()
	videoMessage := up.Message.GetVideoMessage()
	audioMessage := up.Message.GetAudioMessage()
	documentMessage := up.Message.GetDocumentMessage()
	stickerMessage := up.Message.GetStickerMessage()
	var medii x.DownloadableMessage
	var typeM string
	if imageMessage != nil {
		medii = imageMessage
		typeM = "image"
	} else if videoMessage != nil {
		medii = videoMessage
		typeM = "video"
	} else if audioMessage != nil {
		medii = audioMessage
		typeM = "audio"
	} else if documentMessage != nil {
		medii = documentMessage
		typeM = "document"
	} else if stickerMessage != nil {
		medii = stickerMessage
		typeM = "sticker"
	}
	imageQMessage := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetImageMessage()
	videoQMessage := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetVideoMessage()
	audioQMessage := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetAudioMessage()
	documentQMessage := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetDocumentMessage()
	stickerQMessage := up.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetStickerMessage()
	var mediia x.DownloadableMessage
	var typeQM string
	if imageQMessage != nil {
		mediia = imageQMessage
		typeQM = "image"
	} else if videoQMessage != nil {
		mediia = videoQMessage
		typeQM = "video"
	} else if audioQMessage != nil {
		mediia = audioQMessage
		typeQM = "audio"
	} else if documentQMessage != nil {
		mediia = documentQMessage
		typeQM = "document"
	} else if stickerQMessage != nil {
		mediia = stickerQMessage
		typeQM = "sticker"
	}
	_, command, _ := strings.Cut(strings.ToLower(strings.Split(cmds, " ")[0]), prefix)
	return &Parse {
	Id: up.Info.ID,
	Sender: &sender,
	Chat: &chat,
	Pushname: up.Info.PushName,
	Timestamp: &times,
	Body: body,
	Media: medii,
	TypeM: typeM,
	Quoted: QuotedStruct{
		Id: up.Message.GetExtendedTextMessage().GetContextInfo().GetStanzaId(),
		Sender: &senderQ,
		Body: quotedBody,
		Media: mediia,
		TypeM: typeQM,
	},
	Pref: prefix,
	Cmd : command,
	CmdP: strings.ToLower(strings.Split(cmds, " ")[0]),
	Query: strings.Join(strings.Split(cmds, " ")[1:], " "),
	IsGc: up.Info.IsGroup,
	IsMe: up.Info.IsFromMe,
	IsOwn: strings.Contains(up.Info.Sender.String(), noown),
	IsAdmin: ball.GetGroupAdmin(up.Info.Chat, up.Info.Sender.String()),
	IsBotAdmin: ball.GetGroupAdmin(up.Info.Chat, nobot),
	}
}