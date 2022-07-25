package impl

import (
	log "github.com/sirupsen/logrus"
	"github.com/tqtcloud/alertmanager-dispose/pkg/sshd"
	"os"
)

func ConnectHost(user string, password string, host string, port int) error {
	_, session, err := sshd.Connect(user, password, host, port)
	if err != nil {
		log.Infof("连接远端服务器失败：%s", err)
		return err
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	defer session.Close()
	cmd := `cd /usr/local/gzstrong/jt808Gps && sh restart.sh  && echo $(date +%Y-%m-%d" "%H:%M:%S) >> alert.log`
	//cmd := `cd /usr/local/gzstrong/jt808Gps && docker restart  node-exporter && echo $(date +%Y-%m-%d" "%H:%M:%S) >> alert.log`
	err = session.Run(cmd)
	if err != nil {
		log.Errorf("远端命令执行失败：%s", err)
		return err
	}
	return nil
}
