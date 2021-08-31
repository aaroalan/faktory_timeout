# Timeout issue in K8s

## Steps

1. Build docker image, this will build `app.go` into a deployable image: 
```shell
docker build -t worker-test-faktory .
```
3. With kubernetes running apply the k8s configs: 
```shell
kubectl apply -f k8s.yaml
```

Both pods should be running, you can get the pod names with the command:
```shell
kubectl get pods
```
with this command you should be able to see the pod names e.g.
```shell
NAME                              READY   STATUS    RESTARTS   AGE
faktory-server-5bdb4ff565-8pmzr   1/1     Running   0          11s
faktory-worker-7bbc44bdb8-dlt7m   1/1     Running   0          11s
```

note: faktory-worker deployment has a "wait" time, worker will be running until ready column shows 1/1.

You can bind the faktory dashboard port to your local host with the command:

```shell
kubectl port-forward <FAKTORY_POD_NAME> <LOCAL_PORT>:7420
# e.g.
kubectl port-forward faktory-server-5bdb4ff565-8pmzr 7422:7420
```

then you can access dashboard in "http://localhost:7422". PASSWORD: "foobar" (it is set in the k8s.yaml file)

You can see the log errors if you access the logs of the running worker.

```shell
kubectl logs -f <FAKTORY_WORKER_POD_NAME>
# e.g.
kubectl logs -f faktory-worker-7bbc44bdb8-dlt7m
```

you should see the error logs e.g.:

```shell
2021/08/31 16:27:42.551203 read tcp 10.1.0.17:59932->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:42.585250 read tcp 10.1.0.17:59936->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:43.491308 read tcp 10.1.0.17:59944->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:43.950229 read tcp 10.1.0.17:59950->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:44.132562 read tcp 10.1.0.17:59952->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:44.244863 read tcp 10.1.0.17:59954->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:44.557779 read tcp 10.1.0.17:59958->10.103.89.27:7419: i/o timeout
2021/08/31 16:27:44.626693 read tcp 10.1.0.17:59962->10.103.89.27:7419: i/o timeout
```
