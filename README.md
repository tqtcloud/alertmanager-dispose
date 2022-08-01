```shell
# 安装部署于 10.163.90.49
# 直接打包
go build -ldflags "-s -w" -o alertmanager-dispose .\main.go
nohup  /opt/alertmanager-dispose > /opt/alertmanager-dispose.log 2>&1  &
# docker 打包部署
cd /usr/local/src/alertmanager-dispose
docker build -t harbor.gzstrong.com/library/alertmanager-dispose:v1 .
docker run -d --name alert-dispose -p 18083:18083  harbor.gzstrong.com/library/alertmanager-dispose:v1
# 配置在alertmanager如下
...
receivers:
- name: 'web.hook'
  webhook_configs:
  - url: 'http://10.163.90.49:8060/dingtalk/webhook1/send'
    send_resolved: true
  - url: 'http://10.163.90.49:18083/webhook'
    send_resolved: true
- name: 'kafkaalarm'
...

# 配置防火墙
firewall-cmd --zone=public --add-port=18083/tcp --permanent
firewall-cmd --reload

```