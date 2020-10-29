`tempconv.CelsiusFlag(name string, value Celsius, usage string)` の
第 2 引数に渡されたデフォルト値リテラルは、`Celsius` 型に暗黙的に変換される。

`Celsius` 型は `Celsius.String()` が定義されており、
この中で値に `°` を付加して文字列化する処理が書かれている。
また、`fmt` 系の関数で値を出力する際には `String()` メソッドが定義されていればそれが利用される。

ここで、`flag` パッケージでデフォルトメッセージを出力する `PrintDefaults()` メソッドでは、
以下のように `fmt.Sprintf()` を利用している。
https://github.com/golang/go/blob/go1.15.3/src/flag/flag.go#L526


上記をまとめて、以下の理由により `°` が自動的にヘルプメッセージに付加される：

- `flag` パッケージのヘルプメッセージは `PrintDefaults()` メソッドを呼んでおり、
- その中で `String()` メソッドをさらに呼んでおり、
- `String()` メソッドの中で `°` が付加されているため。
