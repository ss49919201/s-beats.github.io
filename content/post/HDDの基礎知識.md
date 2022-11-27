---
title: "HDDの基礎知識"
date: 2022-11-27T13:46:33+09:00
draft: false
comments: false
adsense: false
---

HDDの基本的な用語や構成を全然覚えられず、都度検索しているので備忘録としてまとめます。

## プラッタ

- データを保存する円盤
- プラッタは複数枚搭載
- プラッタの片面または両面に情報を保持

## ヘッド

- プラッタ上の情報を読み書きする部品
- ヘッドはプラッタの数だけ搭載(両面に情報が保持される場合各面にヘッドが搭載される)

## トラック

- プラッタを同じ円心上に分けた領域
- プラッタの各面に存在

## セクタ

- トラックを放射上に分けた領域
- 最小の記録単位

## シリンダ

- 複数のプラッタにおいて同じ位置に存在するトラックの集まり
- 同じシリンダ上にあるトラックへの読み書きはヘッドを動かす必要なし

## シーク

- 読み書き対象のトラックへのヘッドの移動
- アームを動かすことで実現

## サーチ

- 読み書き対象のセクタへのヘッドの移動
- ディスクの回転で実現

## 読み書きにかかる時間

- トラックへの移動時間(シークの時間) + セクタへの移動時間(サーチの時間) + データ転送の時間

## 参考記事

- https://gigazine.net/news/20151104-how-hard-drive-work/
- https://elite-lane.com/hdd/
- https://gigazine.net/news/20151104-how-hard-drive-work/
- http://park12.wakwak.com/~eslab/pcmemo/hdisk/index.html
- https://www.seplus.jp/dokushuzemi/ec/fe/fenavi/easy_calc/disc/
