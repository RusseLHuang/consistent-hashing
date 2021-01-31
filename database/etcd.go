package database

import (
	"context"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
)

var etcdClient *clientv3.Client

func getDBEndpoint() *clientv3.Client {
	etcdConfig := viper.Get("etcd").(map[string]interface{})
	host := etcdConfig["host"].(string)
	port := etcdConfig["port"].(string)

	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{host + ":" + port},
			DialTimeout: 2 * time.Second,
		},
	)

	if err == context.DeadlineExceeded {
		panic(err)
	}

	return cli
}

func GetEtcdClient() *clientv3.Client {
	if etcdClient == nil {
		etcdClient = getDBEndpoint()
	}

	return etcdClient
}

func FetchDBKeyList() []string {
	newCtx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	client := GetEtcdClient()

	resp, err := client.Get(newCtx, "db_key_list")
	if err != nil {
		panic(err)
	}
	cancel()

	if len(resp.Kvs) < 1 {
		panic("No db list found")
	}

	dbKeyListValue := string(resp.Kvs[0].Value)
	dbKeyList := strings.Split(dbKeyListValue, " ")

	return dbKeyList
}

func FetchDBConnectionInfo(key string) string {
	newCtx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	client := GetEtcdClient()

	resp, err := client.Get(newCtx, key)
	if err != nil {
		panic(err)
	}
	cancel()

	if len(resp.Kvs) < 1 {
		panic("No db list found")
	}

	connectionInfo := string(resp.Kvs[0].Value)

	return connectionInfo
}
