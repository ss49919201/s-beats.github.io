---
title: "デフォルト"
date: 2021-12-23T21:04:28+09:00
draft: true
comments: false
adsense: false
---

```sh
touch archetypes/default.md
```

```md
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
comments: false
adsense: false
---
```

```sh
hugo new post/title.md
```