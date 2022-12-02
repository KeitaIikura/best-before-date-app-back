package session

import (
	"bbdate/pkg/config"
	"encoding/json"

	"github.com/gin-contrib/sessions"
)

const (
	TmxSessionKey      = "tmx-session-id"
	TmxSessionValueKey = "tmx-sessionv-id"
)

type TmxSession struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func NewTmxSession(userID, userName string) *TmxSession {
	return &TmxSession{
		UserID:   userID,
		UserName: userName,
	}
}

// byteへの変換
func MarshalTmxSession(ts TmxSession) ([]byte, error) {
	b, err := json.Marshal(ts)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

// bytesからの変換
func UnmarshalTmxSession(b []byte) (*TmxSession, error) {
	session := new(TmxSession)
	if err := json.Unmarshal(b, session); err != nil {
		return nil, err
	}
	return session, nil
}

// 取得
func GetFromTmxStrore(s sessions.Session) []byte {
	sdata := s.Get(TmxSessionValueKey)
	if sdata == nil {
		return nil // no session
	}
	ret, _ := sdata.([]byte)
	return ret
}

// 保存
func SaveToTmxStore(s sessions.Session, target []byte) error {
	s.Set(TmxSessionValueKey, target)
	return saveTmxSession(s, config.EnvStore.TmxSessionTimeout)
}

// 削除
func DeleteSession(s sessions.Session) error {
	s.Delete(TmxSessionValueKey)
	return saveTmxSession(s, -1)
}

func saveTmxSession(s sessions.Session, maxAge int) error {
	// TODO: Secureフラグの設定
	s.Options(sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   false,
	})
	return s.Save()
}
