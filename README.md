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
<img src=".static/skywalking-UI-alarm.png"/>
<p align="center">-SkyWalking alarm UI-</p>

<img src=".static/skywalking-dingding-notify.png"/>
<p align="center">-dingtalk message body-</p>
