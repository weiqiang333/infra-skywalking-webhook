# infra-skywalking-webhook
skywalking alarm webhook


- install
```bash
go get -u github.com/weiqiang333/infra-skywalking-webhook
cd $GOPATH/src/github.com/weiqiang333/infra-skywalking-webhook/
bash build/build.sh
./bin/infra-skywalking-webhook help
```

- Example
```bash
./bin/infra-skywalking-webhook --config configs/production.yml --address 0.0.0.0:8000
```

# Demonstration
![](https://github.com/weiqiang333/infra-skywalking-webhook/.static/skywalking-UI-alarm.png)
<p align="center">-SkyWalking alarm UI-</p>

![](https://github.com/weiqiang333/infra-skywalking-webhook/.static/skywalking-dingding-notify.png)
<p align="center">-dingtalk message body-</p>