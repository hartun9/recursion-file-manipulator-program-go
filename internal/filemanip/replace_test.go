package filemanip

import (
	"os"
	"testing"
)

func TestReplaceString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		needle    string
		newString string
		expected  string
	}{
		{
			name:      "通常の文字列置換",
			input:     "Hello World",
			needle:    "World",
			newString: "Go",
			expected:  "Hello Go",
		},
		{
			name:      "複数箇所の置換",
			input:     "apple apple banana apple",
			needle:    "apple",
			newString: "orange",
			expected:  "orange orange banana orange",
		},
		{
			name:      "置換対象が存在しない場合",
			input:     "Hello World",
			needle:    "Go",
			newString: "Rust",
			expected:  "Hello World",
		},
		{
			name:      "日本語文字列の置換",
			input:     "こんにちは世界",
			needle:    "世界",
			newString: "Go言語",
			expected:  "こんにちはGo言語",
		},
		{
			name:      "特殊文字の置換",
			input:     "記号: !@#$%^&*()",
			needle:    "!@#",
			newString: "ABC",
			expected:  "記号: ABC$%^&*()",
		},
		{
			name:      "空のファイル",
			input:     "",
			needle:    "test",
			newString: "replaced",
			expected:  "",
		},
		{
			name:      "改行を含む文字列の置換",
			input:     "line1\nline2\nline1\n",
			needle:    "line1",
			newString: "modified",
			expected:  "modified\nline2\nmodified\n",
		},
		{
			name:      "置換後の文字列が空",
			input:     "remove this text",
			needle:    "remove ",
			newString: "",
			expected:  "this text",
		},
		{
			name:      "全文置換",
			input:     "全て置換",
			needle:    "全て置換",
			newString: "完全に新しいテキスト",
			expected:  "完全に新しいテキスト",
		},
		{
			name:      "短い文字列を長い文字列に置換",
			input:     "a b c",
			needle:    "a",
			newString: "longer",
			expected:  "longer b c",
		},
		{
			name:      "長い文字列を短い文字列に置換",
			input:     "longer text here",
			needle:    "longer",
			newString: "s",
			expected:  "s text here",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 一時的な入力ファイルを作成
			inputFile, err := os.CreateTemp("", "input-*.txt")
			if err != nil {
				t.Fatalf("入力ファイルの作成に失敗: %v", err)
			}
			defer os.Remove(inputFile.Name())

			// 入力ファイルにテストデータを書き込む
			if _, err := inputFile.WriteString(tt.input); err != nil {
				t.Fatalf("入力ファイルへの書き込みに失敗: %v", err)
			}
			inputFile.Close()

			// replaceString関数を実行
			ReplaceString(inputFile.Name(), tt.needle, tt.newString)

			// ファイルの内容を読み込む
			result, err := os.ReadFile(inputFile.Name())
			if err != nil {
				t.Fatalf("ファイルの読み込みに失敗: %v", err)
			}

			// 期待される結果と比較
			if string(result) != tt.expected {
				t.Errorf("期待される結果と異なります\n期待: %q\n実際: %q", tt.expected, string(result))
			}
		})
	}
}

func TestReplaceStringFileNotFound(t *testing.T) {
	// 存在しないファイルを指定した場合、panicが発生することを確認
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("存在しないファイルでpanicが発生することを期待")
		}
	}()

	ReplaceString("/non/existent/file.txt", "needle", "newstring")
}
