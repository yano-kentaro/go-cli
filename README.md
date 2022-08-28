# GO-CLI
## 開発環境構築
- 適当なディレクトリにclone
- VSCodeでリモートコンテナのプロジェクトを開く
## コマンド一覧
### 認証情報関連
- 認証情報の追加
```
go-cli auth add --auth ${auth} --key ${key} --token ${token}
```
- 認証情報の修正
```
go-cli auth update --auth ${auth} --key ${key} --token ${token}
```
- 認証情報の削除
```
go-cli auth remove --auth ${auth}
```
### Worksアプリ関連
- アプリ一覧の取得
```
go-cli apps get --auth ${auth}
```
- アプリ詳細の取得
```
go-cli app get --auth ${auth} --app ${app}
```
### Worksアカウント関連
- アカウント一覧の取得
```
go-cli accounts get --auth ${auth}
```
- アカウント詳細の取得
```
go-cli accounts get --auth ${auth} --app ${app}
```
### レコード関連
- レコード一覧取得（固定リスト）
- レコード一覧取得（絞り込み検索）
### 一括処理関連
- レコード一括登録設定の取得
- レコード一括登録の実行
- レコード一括登録作業データの削除
### カスタムリスト関連
- カスタムリスト一覧の取得
### ファイル関連
- ファイルの取得