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

// define type, const, global var, method here
%s

func main() {
	// begin
	%s

	// main
	%s

	//end
	%s
}`

var loopCode string = `
	// main loop
	// file or stdin
	r:=os.Stdin
	scanner := bufio.NewScanner(r)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := []string{scanner.Text()} // 0: line of input
		s = append(s, strings.Fields(s[0])...) //

		// inner main
		%s
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
`

//Run code with imports and before/mid/after codes.
func Run(define, before, mid, after string, withLoop bool, printCode bool, imps ...string) error {
	is := buildImports(imps...)
	var code string
	if withLoop {
		code = fmt.Sprintf(baseCode, is, define, before, fmt.Sprintf(loopCode, mid), after)
	} else {
		code = fmt.Sprintf(baseCode, is, define, before, mid, after)
	}
	//fmt.Println(code)
	fixedCode, err := fixImports(code)
	//fmt.Println(fixedCode)
	if err != nil {
		log.Println(err)
		log.Println(code)
		fixedCode = code
		// Ignore error, cause imports return lazy errors
		// compiler return more error detail.
		// return err
	}
	if printCode {
		// verbose printing
		log.Println(fixedCode)
	}
	fn, err := createFileToTempDir(fixedCode)
	//fmt.Println(fn)
	if err != nil {
		log.Println("Failed to create temporary file")
		return err
	}
	if err != goRun(fn) {
		log.Println(fixedCode)
		return err
	}
	return nil
}

func buildImports(is ...string) string {
	var ret string
	for _, i := range is {
		ret = ret + fmt.Sprintf("\"%s\";", i)
	}
	return ret
}

// fixImports formats and adjusts imports.
func fixImports(code string) (string, error) {
	imports.Debug = true
	fixed, err := imports.Process("", []byte(code), &imports.Options{AllErrors: true})
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
