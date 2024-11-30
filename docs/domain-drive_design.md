# このプロジェクトにおけるアーキテクチャ仕様

## ドメインにおける扱い

### /pkg/domain

プロジェクトにおけるオブジェクトをまとめるディレクトリ

#### /pkg/domain/entity

DBへ保存するための構造体で外部に依存せず、外に持ち込まない

#### /pkg/domain/model

プロジェクトで活かすためにここでキャストさせる

#### /pkg/domain/repository

DBへのクエリのインターフェイス

#### /pkg/domain/service

model内で別モデルを入れる際にここで呼び出す

#### /pkg/domain/value_object

外部へ型をそのまま持ち込ませ内容にキャストさせる値オブジェクト
