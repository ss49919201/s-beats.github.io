---
title: "SDKからSESのメール送信APIを呼び出す際の注意点(Go)"
date: 2022-06-02T07:50:24+09:00
draft: true
comments: false
adsense: false
---

Sourceに非ASCIIを含めると文字化けする。
mimeパッケージでエンコードする。
RFC2047によると、エンコード対象のほとんどがASCIIの場合はQエンコード、そうでない場合はBエンコードが推奨されるそう。