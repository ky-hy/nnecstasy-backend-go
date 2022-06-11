# nnecstasy-backend-go
nnecstasy のバックエンドリポジトリです。

目次
- [nnecstasy-backend-go](#nnecstasy-backend-go)
- [必要ツール](#必要ツール)
- [開発手順](#開発手順)


# 必要ツール

docker: 20.10.11
docker-compose: 1.29.2

# 開発手順

```sh
# git clone
$ git clone https://github.com/ky-hy/nnecstasy-backend-go.git
$ cd ./nnecstasy-backend-go
# コンテナ起動
$ docker-compose up -d --build
# expressのログの確認
$ docker-compose logs -f golang-demo
```

[http://localhost:8081/](http://localhost:8081/)にアクセスして、GraphiQL(Graphql実行環境)の画面が出ていればOKです。
