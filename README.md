# talks

スライド資料置き場。

公開スライドは[こちら](https://talks-3l6.pages.dev/)

## 2024-06-28 社内LT
- [Goを嫌いにならないためのメンタルモデル](./2024-06-28/2024-06-28-not-hate-go.pdf)

## build

```
pnpm run build
```

## export

PDFでスライドを出力。

PDF化に必要なためインストール。

```
pnpm i -D playwright-chromium
```

```
pnpm run export
```

## deploy

Cloudflare pagesにデプロイ

```
pnpm i -D -w wrangler
```

cloudflareにログイン

```
pnpx wrangler login
```

```
pnpx wrangler deploy dist
```