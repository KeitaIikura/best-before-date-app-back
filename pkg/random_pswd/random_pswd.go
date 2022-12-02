package random_pswd

import (
	"github.com/sethvargo/go-password/password"
)

func GenerateRandomPassword() (string, error) {
	// パスワード規約として設定している記号群を含むように明示指定
	symbols := "!\"#$%&'()*+,-./:;<=>?@[]^_`{|}~"

	pwgInput := &password.GeneratorInput{
		Symbols: symbols,
	}

	gen, err := password.NewGenerator(pwgInput)
	if err != nil {
		return "", err
	}

	pw, err := gen.Generate(12, 1, 1, false, false)
	if err != nil {
		return "", err
	}

	return pw, nil
}
