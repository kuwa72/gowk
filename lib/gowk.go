package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/tools/imports"
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
	File     string
}

var baseCode string = `package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	%s
)

func main() {
	// begin
	%s

	// main loop
	// file or stdin
	r:=os.Stdin
	scanner := bufio.NewScanner(r)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := []string{scanner.Text()} // 0: line of input
		s = append(s, strings.Split(s[0], " ")...) //

		// main
		%s
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//end
	%s
}`

//Run code with imports and before/mid/after codes.
func Run(before, mid, after string, imps ...string) error {
	is := buildImports(imps...)
	code := fmt.Sprintf(baseCode, is, before, mid, after)
	//fmt.Println(code)
	fixedCode, err := fixImports(code)
	//fmt.Println(fixedCode)
	if err != nil {
		return err
	}
	fn, err := createFileToTempDir(fixedCode)
	//fmt.Println(fn)
	if err != nil {
		log.Println(fixedCode)
		return err
	}
	return goRun(fn)
}

func buildImports(is ...string) string {
	var ret string
	for _, i := range is {
		ret = ret + fmt.Sprintf("\"%s\" ", i)
	}
	return ret
}

// fixImports formats and adjusts imports.
func fixImports(code string) (string, error) {
	fixed, err := imports.Process("", []byte(code), nil)
	return string(fixed), err
}

func goRun(fn string) error {
	args := append([]string{"run"}, fn)
	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// createFileToTempDir is create temporary directory.
// You must remove directory after usign.
// ex: defer os.RemoveAll(dir)
// return directory name and error.
func createFileToTempDir(code string) (string, error) {
	dir, err := ioutil.TempDir("", "gowk")
	if err != nil {
		return "", err
	}

	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(code), 0666); err != nil {
		return dir, err
	}

	return tmpfn, nil
}