# taskapp

## docker-compose

```
docker-compose up
localhost:80
```

## kubernetes

```
docker build -t taskapp/proxy ./nginx
docker build -t taskapp/app ./next
docker build -t taskapp/api ./go/github.com/ke6ch/api
docker build -t taskapp/mysql ./mysql

kubectl apply -k ./nginx/base
kubectl apply -k ./next/base
kubectl apply -k ./go/github.com/ke6ch/api/base
kubectl apply -k ./mysql/base
```
