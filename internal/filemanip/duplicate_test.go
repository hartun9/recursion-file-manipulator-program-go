package filemanip

import (
	"os"
	"testing"
)

func TestDeplicateContents(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		n        int
		expected string
	}{
		{
			name:     "通常のテキストをn=1で複製",
			input:    "Hello\n",
			n:        1,
			expected: "Hello\nHello\n",
		},
		{
			name:     "通常のテキストをn=2で複製",
			input:    "Test\n",
			n:        2,
			expected: "Test\nTest\nTest\n",
		},
		{
			name:     "通常のテキストをn=3で複製",
			input:    "Go\n",
			n:        3,
			expected: "Go\nGo\nGo\nGo\n",
		},
		{
			name:     "複数行のファイルを複製",
			input:    "line1\nline2\n",
			n:        1,
			expected: "line1\nline2\nline1\nline2\n",
		},
		{
			name:     "日本語テキストを複製",
			input:    "こんにちは\n",
			n:        2,
			expected: "こんにちは\nこんにちは\nこんにちは\n",
		},
		{
			name:     "特殊文字を含むファイルを複製",
			input:    "記号: !@#$%^&*()\n",
			n:        1,
			expected: "記号: !@#$%^&*()\n記号: !@#$%^&*()\n",
		},
		{
			name:     "末尾に改行がないファイル",
			input:    "no newline",
			n:        1,
			expected: "no newlineno newline",
		},
		{
			name:     "n=0の場合（複製しない）",
			input:    "original\n",
			n:        0,
			expected: "original\n",
		},
		{
			name:     "空のファイルを複製",
			input:    "",
			n:        2,
			expected: "",
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

			// deplicateContents関数を実行
			DeplicateContents(inputFile.Name(), tt.n)

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

func TestDeplicateContentsFileNotFound(t *testing.T) {
	// 存在しないファイルを指定した場合、panicが発生することを確認
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("存在しないファイルでpanicが発生することを期待")
		}
	}()

	DeplicateContents("/non/existent/file.txt", 1)
}
