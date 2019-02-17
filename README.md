# clean-arch-demo

クリーンアーキテクチャもどき。お勉強。


## Develop

```bash
# Install
$ git clone git@github.com:grandcolline/clean-arch-demo.git

# Serve For Develop
$ docker-compose up -d

# Build
$ docker build -t grandcolline/clean-arch-demo .
```

## Note

### Entities

ビジネスルールの為のデータ構造、もしくはメソッドを持ったオブジェクト

### Use cases

アプリケーション固有のビジネスルール。エンティティとのデータの流れを組み立てる。

### Interface Adapters

外部から、ユースケースとエンティティーで使われる内部形式にデータを変換、  
または内部から外部の機能にもっとも便利な形式に、データを変換するアダプター。

コントローラーはインプットポートとアウトプットポートを組み立てて、インプット ポートを実行するだけ。

### Frameworks & Drivers

フレームワークやツールから構成される。このレイヤーには、多くのコードは書かない。  
ただし、ひとつ内側の円と通信するつなぎのコードは、ここに含まれる。


## 参考
http://nakawatch.hatenablog.com/entry/2018/07/11/181453
https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
