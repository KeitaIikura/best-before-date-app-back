package crypter

type (
	ICrypter interface {
		EncryptString(source string) (string, error)
		EncryptData(data interface{}) ([]byte, error)
		DecryptString(source string) (string, error)
		DecryptData(data []byte, output interface{}) error
		GenerateSha512Hash(str ...string) string
		GenerateSha256Hash(str ...string) string
	}
)
