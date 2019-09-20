# infra-skywalking-webhook
skywalking alarm webhook


- install
```bash
go get -u github.com/weiqiang333/infra-skywalking-webhook
bash build/build.sh
```


- user
```bash
./bin/infra-skywalking-webhook help
```


- Example
```bash
./bin/infra-skywalking-webhook --config configs/production.yml --address 0.0.0.0:8000
```
