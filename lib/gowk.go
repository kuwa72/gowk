package lib

import (
	"go/printer"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	FS       string            // 入力フィールドセパレータ(空白とタブ)
	CNVFMT   string            // 数値を文字列に変換するフォーマット
	OFMT     string            // 数字の出力フォーマット(%.6g)
	OFS      string            // 出力フィールドセパレータ（空白）
	ORS      string            // 出力レコードセパレータ（\n）
	RS       string            // 入力レコードセパレータ（\n）
	SUBSEP   string            // 配列添字セパレータ(\034)
	ARGC     int               // コマンド行の引数の数+1
	ARGV     []string          // コマンド行の引数の配列
	ENVIRON  map[string]string // 環境変数の値
	FILENAME string            // 入力ファイル名
	FNR      int               // 入力ファイルの通算レコード
	NF       int               // 入力レコードのフィールド数
	NR       int               // 入力レコード総数
	RLENGTH  int               // matchで適合した文字列の長さ
	RSTART   int               // matchで適合した文字列の開始位置

	FilePath string
	Fset     []string
	File     string
}

var (
	baseCode string = `
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// begin
	%s

	// main loop
	// file or stdin
	r:=os.Stdin
	scanner := bufio.NewScanner(r)
	var s = []string{}
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s = append(s, scanner.Text()) // 0: line of input
		s = append(s, strings.Split(s[0], " ") //

		// main
		%s
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//end
	%s
}
`
)

func (s *Session) Run() error {
	f, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}

	err = printer.Fprint(f, s.Fset, s.File)
	if err != nil {
		return err
	}

	return goRun(append(s.ExtraFilePaths, s.FilePath))
}

func goRun(files []string) error {
	args := append([]string{"run"}, files...)
	debugf("go %s", strings.Join(args, " "))
	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
