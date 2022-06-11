# nnecstasy-backend-go
nnecstasy のバックエンドリポジトリです。

目次
- [nnecstasy-backend-go](#nnecstasy-backend-go)
- [必要ツール](#必要ツール)
- [開発手順](#開発手順)
  - [vscode利用者向け](#vscode利用者向け)
- [参考](#参考)


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

[http://localhost:8081/](http://localhost:8081/)にアクセスして、GraphiQL(Graphql実行環境v)の画面が出ていればOKです。
## vscode利用者向け
vscodeの拡張機能 [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) をインストールしてください。

```sh
# git clone
$ git clone https://github.com/ky-hy/nnecstasy-backend-go.git
$ cd ./nnecstasy-backend-go
# vocode起動 
$ code .
```
ctrl(cmd) + shift + p を押して, 「Remote-Containers: Reopen in Container」を実行してください。
実行後、[http://localhost:8081/](http://localhost:8081/)にアクセスして、GraphiQL(Graphql実行環境v)の画面が出ていればOKです。


# 参考
- https://code.visualstudio.com/docs/remote/containers
