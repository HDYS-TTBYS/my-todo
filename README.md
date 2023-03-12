# my-todo

作成中

## dev

1. docker compose で開発環境を立ち上げる

```bash
make up
```

- go_api
  - vscode remote container で編集
- web
  - vscode remote container で編集

## staging

1. local kubernetes cluster を用意(docker desktop 等)

2. pgo をインストールする

```bash
make install-pgo
```

3. nginx-ingress-controller をインストールする

```bash
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
helm install nginx-ingress nginx-stable/nginx-ingress
```

4. 環境変数を設定

```
export DOCKERHUB_USERNAME="your_dockerhub_username"
export GITHUB_SHA="latest"
```
