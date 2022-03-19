# go-malproxy
## 概要
悪性Captive Portalサイトを用いて，個人情報を盗む．
## 使用言語
[Golang](https://go.dev/)

## 設計
### `Server`
- `handler`層
  > サーバ側で受け取ったパラメータを取得・整理する．
- `params`層
  > 具体的なパラメータのみをまとめて管理している．
- `public`層
  > ログイン画面等のデザイン部分を保存している．
- `service`層
  > 所謂ビジネスロジック部分．パラメータの具体的な処理を行う．
- `template`層
  > 表示する`HTML`ファイルを自動生成する．
  > 
  > ひな形を含めた`HTML`ファイルは`templates`層で自動生成される．また，スクリーンショットは`templates/img`内に自動生成される．

## 環境構築
- [これ](https://github.com/KeiTaylor0606/How-to-built-environment/blob/main/VSCode/GolangForVM.md)を参照してください．

