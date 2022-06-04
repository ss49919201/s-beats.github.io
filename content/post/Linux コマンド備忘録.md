---
title: "Linux コマンド備忘録"
date: 2022-05-19T10:11:54+09:00
draft: true
comments: false
adsense: false
---

## head

標準入力の内容を先頭から標準出力に吐き出す。

```sh
$ head -n 7 ./content/post/'Linux コマンド備忘録.md'
```

```sh
---
title: "Linux コマンド備忘録"
date: 2022-05-19T10:11:54+09:00
draft: true
comments: false
adsense: false
---
```

## tail

標準入力の内容を末尾から標準出力に吐き出す。

```sh
$ tail -n 7 ./content/post/'Linux コマンド備忘録.md'
```

```sh
---
title: "Linux コマンド備忘録"
date: 2022-05-19T10:11:54+09:00
draft: true
comments: false
adsense: false
---
```

## grep

標準入力から特定文字列を含む行を抜き出して、標準出力に吐き出す。

```sh
$ grep コマンド ./content/post/'Linux コマンド備忘録.md'
```

```sh
title: "Linux コマンド備忘録"
$ grep コマンド ./content/post/'Linux コマンド備忘録.md'
title: "Linux コマンド備忘録"
title: "Linux コマンド備忘録"
title: "Linux コマンド備忘録"
```