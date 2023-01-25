package simple

import (
	"fmt"
	"strings"
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	// waBinary "go.mau.fi/whatsmeow/binary"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"github.com/bolaxd/dumn/helper"
	)

type S struct {
	Conn *whatsmeow.Client 
	M *events.Message
}

type DownloadableMessage interface {
	proto.Message
	GetDirectPath() string
	GetMediaKey() []byte
	GetFileSha256() []byte
	GetFileEncSha256() []byte
}

var id = helper.GRandomString()
func SimpleGo(X *whatsmeow.Client, B *events.Message) *S  {
	return &S {
		Conn: X,
		M: B,
	}
}
func (ball *S) Reply(teks string, quoted bool) {
		quot := &waProto.ContextInfo{
				StanzaId:      &ball.M.Info.ID,
				Participant:   proto.String(ball.M.Info.Sender.String()),
				QuotedMessage: ball.M.Message,
			}
	if !quoted {
		quot = &waProto.ContextInfo{
				StanzaId:      &ball.M.Info.ID,
				Participant:   proto.String(ball.M.Info.Sender.String()),
				QuotedMessage: ball.M.Message,
			}
	}
	ball.Conn.SendMessage(context.Background(), ball.M.Info.Chat, &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: proto.String(teks),
			ContextInfo: quot,
		},
	}, whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) SendImg(jid types.JID, img []byte, teks string, mime string) {
	resp := ball.Up(img, whatsmeow.MediaImage);
	struk := &waProto.ImageMessage{
		Caption: proto.String(teks),
		Mimetype: proto.String(mime),
		Url: &resp.URL,
		DirectPath: &resp.DirectPath,
		MediaKey: resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256: resp.FileSHA256,
		FileLength: &resp.FileLength,
	}
	ball.Conn.SendMessage(context.Background(), jid, &waProto.Message{ ImageMessage: struk }, whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) SendVid(jid types.JID, vid []byte, teks string, mime string) {
	resp := ball.Up(vid, whatsmeow.MediaVideo);
	struk := &waProto.VideoMessage{
		Caption: proto.String(teks),
		Mimetype: proto.String(mime),
		Url: &resp.URL,
		DirectPath: &resp.DirectPath,
		MediaKey: resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256: resp.FileSHA256,
		FileLength: &resp.FileLength,
	}
	ball.Conn.SendMessage(context.Background(), jid, &waProto.Message{ VideoMessage: struk }, whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) SendAud(jid types.JID, aud []byte, mime string) {
	resp := ball.Up(aud, whatsmeow.MediaAudio);
	struk := &waProto.AudioMessage{
		Mimetype: proto.String(mime),
		Url: &resp.URL,
		DirectPath: &resp.DirectPath,
		MediaKey: resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256: resp.FileSHA256,
		FileLength: &resp.FileLength,
	}
	ball.Conn.SendMessage(context.Background(), jid, &waProto.Message{ AudioMessage: struk }, whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) SendStik(jid types.JID, stik []byte) {
	resp := ball.Up(stik, whatsmeow.MediaImage);
	struk := &waProto.StickerMessage{
		Mimetype: proto.String("image/webp"),
		Url: &resp.URL,
		DirectPath: &resp.DirectPath,
		MediaKey: resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256: resp.FileSHA256,
		FileLength: &resp.FileLength,
	}
	ball.Conn.SendMessage(context.Background(), jid, &waProto.Message{ StickerMessage: struk }, whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) LeaveGc(jid types.JID)  {
	ball.Conn.LeaveGroup(jid)
}

func (ball *S) GetMetadata(jid types.JID) (*types.GroupInfo)  {
	val, _ := ball.Conn.GetGroupInfo(jid)
	return val
}

func (ball *S) Joining(code string) {
	ball.Conn.JoinGroupWithLink(code)
}

func (ball *S) GetInfoLink(code string) (*types.GroupInfo) {
	val, _ := ball.Conn.GetGroupInfoFromLink(code)
	return val
}

// func (ball *S) GcUpdate(jid types.JID, typ string, map[types.JID]typ) (*waBinary.Node)  {
// 	val, _ := ball.Conn.UpdateGroupParticipants(jid, typ)
// 	return val
// }

func (ball *S) SetGcPP(jid types.JID, pp []byte) string {
	val, err := ball.Conn.SetGroupPhoto(jid, pp)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func (ball *S) SetGcName(jid types.JID, name string) {
	ball.Conn.SetGroupName(jid, name)
}

func (ball *S) SetGcChat(jid types.JID, open bool) {
	ball.Conn.SetGroupAnnounce(jid, open)
}

func (ball *S) SetGcLock(jid types.JID, lock bool) {
	ball.Conn.SetGroupLocked(jid, lock)
}
func (ball *S) SetDesc(jid types.JID, topic string)  {
	ball.Conn.SetGroupTopic(jid, "", "", topic)
}

// func (ball *S) DelMsg(send types.JID, "" types.MessageID) {
// 	ball.Conn.SendMessage(context.Background(), send, "", ball.Conn.BuildRevoke(ball.M.Info.Chat, send, ""))
// }

func (ball *S) SetLink(jid types.JID, revoke bool) string {
	val, _ := ball.Conn.GetGroupInviteLink(jid, revoke)
	return val
}

func (ball *S) PollMsg(jid types.JID, teks string, isi []string)	 {
	ball.Conn.SendMessage(context.Background(), jid, ball.Conn.BuildPollCreation(teks, isi, 1), whatsmeow.SendRequestExtra{ ID: id })
}

func (ball *S) DL(msg DownloadableMessage) []byte  {
	resp, _ := ball.Conn.Download(msg)
	return resp
}

func (ball *S) Up(media []byte, info whatsmeow.MediaType) (resp whatsmeow.UploadResponse)  {
	val, err := ball.Conn.Upload(context.Background(), media, info)
	if err != nil {
		panic("Byte kosong")
	}
	return val
}

func (ball *S) FetchGroupAdmin(jid types.JID) ([]string, error) {
	var Admin []string
	resp, err := ball.Conn.GetGroupInfo(jid)
	if err != nil {
		return Admin, err
	} else {
		for _, group := range resp.Participants {
			if group.IsAdmin || group.IsSuperAdmin {
				Admin = append(Admin, group.JID.String())
			}
		}
	}
	return Admin, nil
}

func (ball *S) GetGroupAdmin(jid types.JID, sender string) bool {
	if !ball.M.Info.IsGroup {
		return false
	}
	admin, err := ball.FetchGroupAdmin(jid)
	if err != nil {
		return false
	}
	for _, v := range admin {
		if strings.Contains(v, sender) {
			return true
		}
	}
	return false
}