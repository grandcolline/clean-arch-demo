# clean-arch-demo

クリーンアーキテクチャもどき。お勉強。


## Develop

```bash
# Install
$ git clone git@github.com:grandcolline/clean-arch-demo.git
$ cd clean-arch-demo

# Serve For Develop
$ docker-compose up -d

# Build
$ docker build -t grandcolline/clean-arch-demo .

# Run
$ docker run grandcolline/clean-arch-demo
```

## API TEST

```bash
# 登録
$ curl -X POST "http://localhost:8080/users" -d "{\"name\":\"John\",\"email\":\"John@example.com\"}"

# 確認
$ curl -X GET "http://localhost:8080/users"
$ curl -X GET "http://localhost:8080/users/1"

# 変更
$ curl -X PUT "http://localhost:8080/users/1" -d "{\"name\":\"Johhhn\",\"email\":\"John@example.com\"}"

# 削除
$ curl -X DELETE "http://localhost:8080/users/1"
```

## Note

![uncle-bob's CleanArchitecture image](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### Entities

ビジネスルールの為のデータ構造、もしくはメソッドを持ったオブジェクト。

```java
entity
 └── user.go  // ユーザエンティティ
```

### Use cases

アプリケーション固有のビジネスルール。エンティティとのデータの流れを組み立てる。

```java
usecase
 ├── user.go  // ユーザエンティティに対するインストラクタの実装
 └── util.go  // 複数のインストラクタで共通のポートの定義
```

### Interface Adapters

外部から、ユースケースとエンティティで使われる内部形式にデータを変換、  
または内部から外部の機能にもっとも便利な形式に、データを変換するアダプタ。

コントローラはインプットポートとアウトプットポートを組み立てて、インプットポートを実行するだけ。

```java
adapter
 ├── controller              // Controller: InputPortの実行。
 │   ├── form
 │   │   └── user_form.go
 │   └── user_controller.go
 ├── gateway                 // Gateway: RegistoryPortの実装。DBとの接続。
 │   ├── model
 │   │   └── user_model.go
 │   └── user_gateway.go
 ├── logger                  // Logger: LoggerPortの実装。ログの出力。
 │   └── logger.go
 └── presenter               // Presenter: OutputPortの実装。ユーザレスポンス。
     ├── cmn_presenter.go
     └── user_presenter.go
```

### Frameworks & Drivers

フレームワークやツールから構成される。このレイヤーには、多くのコードは書かない。  
ただし、ひとつ内側の円と通信するつなぎのコードは、ここに含まれる。

```java
driver
 ├── config
 │   └── config.go
 ├── mysql
 │   └── connection.go
 └── router.go
```

