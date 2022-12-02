package encoding

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
)

// 「data:*/*;base64,*****」形式の文字列をDecodeする
//	返り値 byte[]:デコード後のバイト配列、string:ファイルのmimeType(e.g. image/png)
func DecodeString(ctx context.Context, str string) ([]byte, string, error) {
	// 接頭辞「data:*/*;base64,」の部分を除いてDecodeする
	targetStr := str[strings.Index(str, ",")+1:]
	binaryFileData, err := base64.StdEncoding.DecodeString(targetStr)
	if err != nil {
		return nil, "", fmt.Errorf("err Base64.DecodeString: %w", err)
	}

	// 接頭辞から「*/*」の部分を取り出す。(e.g. image/png)
	mimeType := str[strings.Index(str, ":")+1 : strings.Index(str, ";")]

	return binaryFileData, mimeType, nil
}
