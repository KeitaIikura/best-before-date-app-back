package calculate

import (
	"testing"
)

// 解説用コメント
// ファイル名は「~_test.go」とすること。~の部分はテスト対象のファイル名と同じで良い。
// テストファイルはテスト対象のファイルと同一パッケージ内に作成する。

// Division関数のテスト。テスト関数は「Test~という関数名にする」
func TestDivision(t *testing.T) {
	// Divsion関数の引数を定義する
	type args struct {
		top     int
		bottom  int
		baseNum int
	}

	// テストの基本となるスライス。このスライス内に定義したテストが全て実行されていく。
	tests := []struct {
		name string  // テスト名
		args args    // テスト対象の関数に渡す引数
		want float64 // 関数の実行結果として予想される返り値（実際にテストを実行した時の返り値がこれと一致すればテスト成功となる）
	}{
		{
			// テストケース１つめ
			name: "小数点2桁表示（切り捨て）",
			args: args{
				top:     10,
				bottom:  3,
				baseNum: 2,
			},
			want: 3.33,
		},
		{
			// テストケース２つめ
			name: "小数点3桁表示（繰り上げ）",
			args: args{
				top:     20,
				bottom:  3,
				baseNum: 3,
			},
			want: 6.667,
		},
		{
			// テストケース３つめ
			name: "小数点0桁表示（切り捨て）",
			args: args{
				top:     10,
				bottom:  3,
				baseNum: 0,
			},
			want: 3,
		},
	}

	// テストコードの実行部分。以下2行(t.Runの行まで)はどのテストでも共通の記述
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 実際にDivision関数を実行して返り値を取得している。
			got := Division(tt.args.top, tt.args.bottom, tt.args.baseNum)

			// 得られた返り値(got)と予想していた返り値(want)が異なる場合はエラーを出す。
			if got != tt.want {
				t.Errorf("failed test %s: got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDivisionFloat(t *testing.T) {
	type args struct {
		top     float64
		bottom  float64
		baseNum int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "小数点2桁表示（切り捨て）",
			args: args{
				top:     10,
				bottom:  3,
				baseNum: 2,
			},
			want: 3.33,
		},
		{
			name: "小数点3桁表示（繰り上げ）",
			args: args{
				top:     20,
				bottom:  3,
				baseNum: 3,
			},
			want: 6.667,
		},
		{
			name: "小数点0桁表示（切り捨て）",
			args: args{
				top:     10,
				bottom:  3,
				baseNum: 0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DivisionFloat(tt.args.top, tt.args.bottom, tt.args.baseNum)
			if got != tt.want {
				t.Errorf("failed test %s: got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestRate(t *testing.T) {
	type args struct {
		top     int
		bottom  int
		baseNum int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "小数点2桁表示（切り捨て）",
			args: args{
				top:     1,
				bottom:  3,
				baseNum: 2,
			},
			want: 33.33,
		},
		{
			name: "小数点3桁表示（繰り上げ）",
			args: args{
				top:     2,
				bottom:  3,
				baseNum: 3,
			},
			want: 66.667,
		},
		{
			name: "小数点0桁表示（切り捨て）",
			args: args{
				top:     1,
				bottom:  3,
				baseNum: 0,
			},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rate(tt.args.top, tt.args.bottom, tt.args.baseNum)
			if got != tt.want {
				t.Errorf("failed test %s: got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
