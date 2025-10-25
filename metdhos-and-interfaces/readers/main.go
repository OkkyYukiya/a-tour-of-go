package main

import (
	"io"
	"os"
	"strings"
)

// Readers
// io packageはデータストリームを読むことを表現する
// io.Reader interfaceを規定している

// Goの標準ライブラリにはfile, network接続、圧縮、暗号化などで、このinterfaceの多くの実装がある
// io.ReaderのinterfaceはReadメソッドを持つ

// io.Reader interfaceはReadメソッドを持つ
// func main() {
// 	// strings.Reader を作成し、 8 byte毎に読み出しています。
// 	r := strings.NewReader("Hello Reader")
// 	b := make([]byte, 8)

// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n]= %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// 	// n = 8 err = <nil> b = [72 101 108 108 111 32 82 101]
// 	// b[:n]= "Hello Re"
// 	// n = 4 err = <nil> b = [97 100 101 114 111 32 82 101]
// 	// b[:n]= "ader"
// 	// n = 0 err = EOF b = [97 100 101 114 111 32 82 101]
// 	// b[:n]= ""
// }

// type MyReader struct{}

// // Read は b 全体を 'A' で埋め、書き込んだバイト数を返す。
// // 無限ストリームなのでエラーは返さない。
// func (MyReader) Read(b []byte) (int, error) {
// 	for i := range b {
// 		b[i] = 'A'
// 	}
// 	return len(b), nil
// }

// func main() {
// 	reader.Validate(MyReader{})
// }

type rot13Reader struct {
	r io.Reader
}

// ROT13変換関数
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

// Readメソッドの実装
func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}