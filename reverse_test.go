package main

import (
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "複数行のファイルを反転",
			input:    "line1\nline2\nline3\n",
			expected: "line3\nline2\nline1\n",
		},
		{
			name:     "1行のファイル",
			input:    "single line\n",
			expected: "single line\n",
		},
		{
			name:     "空のファイル",
			input:    "",
			expected: "",
		},
		{
			name:     "末尾に改行がない複数行",
			input:    "line1\nline2\nline3",
			expected: "line3\nline2\nline1\n",
		},
		{
			name:     "空行を含むファイル",
			input:    "line1\n\nline3\n",
			expected: "line3\n\nline1\n",
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

			// 一時的な出力ファイルを作成
			outputFile, err := os.CreateTemp("", "output-*.txt")
			if err != nil {
				t.Fatalf("出力ファイルの作成に失敗: %v", err)
			}
			outputPath := outputFile.Name()
			outputFile.Close()
			defer os.Remove(outputPath)

			// reverse関数を実行
			reverse(inputFile.Name(), outputPath)

			// 出力ファイルの内容を読み込む
			result, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("出力ファイルの読み込みに失敗: %v", err)
			}

			// 期待される結果と比較
			if string(result) != tt.expected {
				t.Errorf("期待される結果と異なります\n期待: %q\n実際: %q", tt.expected, string(result))
			}
		})
	}
}

func TestReverseFileNotFound(t *testing.T) {
	// 存在しないファイルを指定した場合、panicが発生することを確認
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("存在しないファイルでpanicが発生することを期待")
		}
	}()

	outputFile, err := os.CreateTemp("", "output-*.txt")
	if err != nil {
		t.Fatalf("出力ファイルの作成に失敗: %v", err)
	}
	defer os.Remove(outputFile.Name())
	outputFile.Close()

	reverse("/non/existent/file.txt", outputFile.Name())
}
