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

## Note

![uncle-bob's CleanArchitecture image](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### Entities

ビジネスルールの為のデータ構造、もしくはメソッドを持ったオブジェクト

```java
entity
└── user.go
```

### Use cases

アプリケーション固有のビジネスルール。エンティティとのデータの流れを組み立てる。

```java
usecase
├── logger.go
└── user.go
```

### Interface Adapters

外部から、ユースケースとエンティティーで使われる内部形式にデータを変換、  
または内部から外部の機能にもっとも便利な形式に、データを変換するアダプター。

コントローラーはインプットポートとアウトプットポートを組み立てて、インプット ポートを実行するだけ。

```java
adapter
├── controller
│   ├── error.go
│   └── user_controller.go
├── gateway
│   └── user_repository.go
├── logger
│   └── logger.go
└── presenter
    └── user_presenter.go
```

### Frameworks & Drivers

フレームワークやツールから構成される。このレイヤーには、多くのコードは書かない。  
ただし、ひとつ内側の円と通信するつなぎのコードは、ここに含まれる。

```java
driver
├── logger.go
├── mysql
│   └── connection.go
└── router.go
```

