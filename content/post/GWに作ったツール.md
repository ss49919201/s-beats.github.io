---
title: "GWに作ったツール"
date: 2022-05-07T22:37:49+09:00
draft: false
comments: false
adsense: false
---

今回のGW期間で作ったツールを雑に列挙する。

## idgen

ダミーのUUIDとかULID欲しいなーって時にコマンド打ち込むだけで出力してくれるツールを作った。
個人開発の中で結構使ってて気に入っている。
いずれはVersion指定とかもできたらいいかなと思っている。

https://github.com/s-beats/idgen

## ssmparam

https://github.com/aws/aws-sdk-go の SSM のAPIの薄いラッパーを作った。
`~Input`と`~Output`の形の構造体を扱うのが面倒な時があったのがモチベーションである。
ユースケースを絞りすぎている感はあるけど割と気に入っている。
簡素な作りだけど開発自体は楽しかった。

https://github.com/s-beats/ssmparam

## まとめ

やっぱり開発って楽しい。
最近はWEBアプリとかを作るよりも便利ツールやライブラリを作ることに喜びを感じます。
