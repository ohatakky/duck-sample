package main

import "fmt"

type SmartPhone interface {
	Read()
	Play()
	Search()
}

type IPhone6 struct {
	Size int64
}
type Animal struct{}

// IPhone6はRead(), Play(), Search()を実装したのでSmartPhoneを満たす
func (i IPhone6) Read() {
	fmt.Println("read a book")
}

func (i IPhone6) Play() {
	fmt.Println("play a music")
}

func (i IPhone6) Search() {
	fmt.Println("search for information")
}

func duckSmartPhone(s SmartPhone) {
	switch x := s.(type) {
	case IPhone6:
		fmt.Println(x.Size)
	default:
		fmt.Println("not IPhone6")
	}
}

func not_is_a(i IPhone6) {
	fmt.Println("this is an iphone6")
}

type deligate_IPhone7 struct {
	i6     IPhone6
	FeliCa bool
}

type embedded_IPhone7 struct {
	IPhone6
	FeliCa bool
}

func main() {
	iphone6 := IPhone6{Size: 10}
	duckSmartPhone(iphone6)

	// neko := Animal{}
	// duckSmartPhone(neko) // SmartPhoneを満たしていないのでエラー

	// 委譲によるダックタイピング
	de_iphone7 := deligate_IPhone7{i6: iphone6, FeliCa: true}
	duckSmartPhone(de_iphone7.i6)

	// 埋め込みによるダックタイピング
	em_iphone7 := embedded_IPhone7{}
	duckSmartPhone(em_iphone7)
	duckSmartPhone(em_iphone7.IPhone6) // 予期せぬ動作

	not_is_a(iphone6)
	// not_is_a(em_iphone7) // エラー。埋め込みは is - a ではない

	// アドレスを確認
	var de_iphone7_2 *deligate_IPhone7
	de_iphone7_2 = &deligate_IPhone7{i6: iphone6, FeliCa: true}
	fmt.Println(de_iphone7_2.i6)

}
