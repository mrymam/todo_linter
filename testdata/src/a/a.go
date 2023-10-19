package a

// TODO @mrymam until:2023-11-25 description
func f() {}

// hoge TODO // want "TODOから始めてください"
func g() {}

// TODO  hoge huga // want "ユーザーを記入してください"
func a() {}

// TODO @mrymam hoge // want "期限の記入フォーマットが異なります"
func b() {}

// TODO @mrymam until:hogeho // want "期限を正しく入力してください"
func c() {}

// TODO @mrymam until:2010-10-10 // want "期限切れです"
func d() {}
