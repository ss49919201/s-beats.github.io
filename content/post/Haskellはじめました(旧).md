---
title: "Haskellはじめました"
date: 2022-05-29T19:49:14+09:00
draft: true
comments: false
adsense: false
---


まずは GHCup をインストールします。

```sh
curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh
```

いろいろ聞かれるのでよしなに。

```sh
...(略)
Press ENTER to proceed or ctrl-c to abort.
Note that this script can be re-run at any given time.
```

無心でENTERキーを叩く。

```sh
-------------------------------------------------------------------------------

Detected zsh shell on your system...
Do you want ghcup to automatically add the required PATH variable to "/Users/sakaeshinya/.zshrc"?

[P] Yes, prepend  [A] Yes, append  [N] No  [?] Help (default is "P").

p
```

```sh
-------------------------------------------------------------------------------
Do you want to install haskell-language-server (HLS)?
HLS is a language-server that provides IDE-like functionality
and can integrate with different editors, such as Vim, Emacs, VS Code, Atom, ...
Also see https://github.com/haskell/haskell-language-server/blob/master/README.md

[Y] Yes  [N] No  [?] Help (default is "N").

Y
```

補完が効くようになるのかな？
それは欲しいので`[Y]`

```sh
...(略)
Do you want to install stack?
Stack is a haskell build tool similar to cabal that is used by some projects.
Also see https://docs.haskellstack.org/

[Y] Yes  [N] No  [?] Help (default is "N")

Y
```

なんか必要そうなのでYを選択。

```sh
All done!

To start a simple repl, run:
  ghci

To start a new haskell project in the current directory, run:
  cabal init --interactive

To install other GHC versions and tools, run:
  ghcup tui

If you are new to Haskell, check out https://www.haskell.org/ghcup/steps/
```

インストール完了。
反映する。

```sh
$ source .zshrc
```

コードを書いていく。

```sh
$ touch main.hs
```

```haskell
main = putStrLn "Hello World"
```

ビルドする。

```sh
$ ghc main.hs
[1 of 1] Compiling Main             ( main.hs, main.o )
Linking main ..
```

実行する。

```sh
$ ./main
Hello World
```
