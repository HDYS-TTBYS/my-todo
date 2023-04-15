# my-todo

aws lightsail メモリ2G インスタンス k3sで公開中

api:<https://todo.tthd-app.link/api/todos>

frontend:<https://todo.tthd-app.link>

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

1. local kubernetes cluster を用意(k3s 仮想マシン等)

```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION="v1.24.10+k3s1" sh -s - server --cluster-init --disable traefik
sudo cat /etc/rancher/k3s/k3s.yaml
```

2. pgo をインストールする

```bash
make install-pgo
```

3. postgresをdeployする

```bash
kubectl apply -f manifests/postgres.yaml
```

4. cert-manager をインストールする

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.11.0 \
  --set installCRDs=true
```

5. nginx-ingress-controllerをインストールする

```bash
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
kubectl create namespace nginx-ingress
helm install nginx nginx-stable/nginx-ingress -n nginx-ingress
```

6. kubernetes secretを作成

```bash
kubectl -n  cert-manager create secret generic prod-route53-credentials-secret \
  --from-literal=secret-access-key=<SECRET_ACCESS_KEY>
```

7. 環境変数を設定

```
export DOCKERHUB_USERNAME="hdys"
export GITHUB_SHA="latest"
export ACCESSKEYID="*************"
```

letsencrypte-issuer.yamlを編集する(email, server)

8. 証明書作成

```bash
cat manifests/letsencrypte-issuer.yaml | envsubst '${ACCESSKEYID}' | kubectl apply -f -
kubectl apply -f manifests/certificate.yaml
```

9. デプロイ

```bash
make docker-build-and-push
make deploy
```

## production

1. kubernetes cluster を用意(k3s)

```bash
curl -sfL https://get.k3s.io | sh -s - --disable traefik --tls-san "todo.tthd-app.link"
```

2. pgo をインストールする

```bash
make install-pgo
```

3. postgresをdeployする

```bash
kubectl apply -f manifests/postgres.yaml
```

4. cert-manager をインストールする

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.crds.yaml
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.11.0 \
  #--set installCRDs=true
```

5. nginx-ingress-controllerをインストールする

```bash
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
helm install \
  nginx nginx-stable/nginx-ingress \
  --namespace nginx-ingress \
  --create-namespace
```

6. kubernetes secretを作成

```bash
kubectl -n  cert-manager create secret generic prod-route53-credentials-secret \
  --from-literal=secret-access-key=<SECRET_ACCESS_KEY>
```

7. 環境変数を設定

```
export ACCESSKEYID="*************"
```

8. 証明書作成

```bash
cat manifests/letsencrypte-issuer.yaml | envsubst '${ACCESSKEYID}' | kubectl apply -f -
kubectl apply -f manifests/certificate.yaml
```

9. github actionsにsecretを追加する

```bash
DOCKERHUB_USERNAME=*************

DOCKERHUB_PASSWORD=*************

KUBE_CONFIG=cat $HOME/.kube/config | base64
```

10. CD

main branch にpush する
