---
title: "【Go】rateパッケージの使い方"
date: 2022-11-20T18:19:17+09:00
draft: false
comments: false
adsense: false
---

レート制限のある処理の制御に役立つ[golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)の存在を知りました。
繰り返し処理の中で AWS や Stripe のAPIを使いたい時などに採用できるのではないかと思い調査しました。

## はじめにサンプルコード

```go
func main() {
	ctx := context.Background()

	// 1秒間に0.5回(2秒間に1回)の処理を許可する
	l := rate.NewLimiter(0.5, 1)
	for i := 0; i < 10; i++ {
		if err := l.Wait(ctx); err != nil {
			panic(err)
		}
		println(i)
	}

	for i := 0; i < 10; i++ {
		if l.Allow() {
			println(i)
		} else {
			// 許可されなかったら少し待つ
			println("not allowed")
			time.Sleep(time.Second)
		}
	}
}
```

## Allow

即時実行可能かを返します。
内部実装としては引数を固定化して`AllowN`を呼び出しています。
呼び出されている`AllowN`は引数を固定化して`reserveN`を呼び出し戻り値の構造体`Reservation`のbool型のフィールド`ok`の値を返します。
即時実行可能でなければスッと諦めたいときや、別の処理に移ってそちらを終えてからリトライしたい時などに使えそうでしょうか。

## Wait

処理の回数に応じた待機時間だけブロックします。
内部実装としてはは引数を固定化して`WaitN`を呼び出しています。
`context.Context`の`Deadline`を考慮した待機時間を計算し、キャンセル時には内部で検知しエラーを返してくれます。
また、`context.Context`がキャンセルされれば待機時間の途中であってもエラーを返します。
基本は`Wait`を使うことが多そうでしょうか。

## Reserve

処理の実行可能フラグ、待機時間などを持った構造体を返します。
内部実装としてはは引数を固定化して`ReserveN`を呼び出しています。
現時点の制御情報をもとに自前で自由にロジックを組みたい時に使えそうでしょうか。

## reserveN

プライベートメソッドなので利用側から呼び出すことはないですが、大事な処理なので備忘も兼ねて超簡単にまとめておきます。

`reserveN`は時刻、処理を実行したい回数、許容可能な待機時間を引数に受け取り構造体Reservationを返します。
メソッドのコメントには`reserveN`は`AllowN`、`WaitN`、`ReservationN`でヘルパーメソッドであると記載されており、`AllowN`、`WaitN`においては`reserveN`が返す構造体`Reservation`の値をもとロジックが組まれています。
そのことから構造体`Reservation`を生成しているこのメソッドの実装がかなり重要であることがわかります。

## 結び

rateパッケージは400行程度のソースコードで構成されています。
久しぶりにコードリーディングもしたかったので、全部読んで詳細にまとめることを試みましたが挫折しました。(正確には大方読んだけどまとめて記事ににおこす気力が湧かなかった)
