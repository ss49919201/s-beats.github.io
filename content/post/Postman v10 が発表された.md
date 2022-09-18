---
title: "Postman v10 が発表された"
date: 2022-09-10T08:32:36+09:00
draft: true
comments: false
adsense: false
---

Postman v10 のリリースが発表されました。
https://blog.postman.com/announcing-postman-v10/

面白そうなアップデートが結構あったので、内容を簡単にまとめてみます。

## Postman CLI

- Postman と同じランタイムで実行される。
- Newman の基盤の上に乗っている。
- https://blog.postman.com/introducing-the-postman-cli-to-automate-your-api-testing/

自動APIテストとかに使えるかも。

## gRPC サポート

- GraphQL に続き gRPC がサポートされた。
- gRPC のURLをURLバーに打ち込むだけで、使えるメソッドをセットしてくれる。
    - サーバーリフレクションというらしいんだけど、サーバーに問い合わせて使えるメソッドを取得しているのかな？
- https://blog.postman.com/postman-v10-and-grpc-what-you-can-do/

重い腰を上げて gRPC を始める機が熟したか...。

## パートナーワークスペース

- チーム外のメンバーを招待できるようになった。
- 権限を絞ったり、特定のワークスペースにのみアクセスを許可するなどの制御ができる。
    - これらを管理する新しいロールができた。
- https://blog.postman.com/introducing-partner-workspaces/

「チームに入れるとコストが気になるけど、同じワークスペースは使ってほしい業務委託メンバーが複数人いる」みたいなケースで嬉しいかな？

## API Security

- セキュリティルールに違反してしないかをチェックできる。
- 定義だけじゃなくてリクエストもチェックしてくれる。
- Postman が提供してくれるルールをON/OFFして設定する。
- https://blog.postman.com/introducing-api-security-in-postman-v10/

API 定義のレビューってミスを見逃しがちな気がするから良いかも。
というかそもそも脆弱な定義とか全然知らないからこれを使って学んでいきたい。

## API Governance

- 組織の API のルールに違反してしないかをチェックできる。
- 重大度とかも設定できる。
- Postman rule library に含まれる違反の場合は影響や修正方法へのリンクを掲示してくれる。
- https://blog.postman.com/api-governance-with-postman-v10/

スタイルガイドがきちんとある組織ならかなり活きそう。
スタイルガイド作るのが大変そうだけど。

## あまり理解できなかった

- API Builder の改善
- Private API Network

使ったことなくてイメージ付きませんでした。
