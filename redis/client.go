package redis

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/ssh"
)

func Client1() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "dbdd6fea7a96845b",
		DB:       0, // use default database
	})
}

func Client2() *redis.Client {
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

func Client3() *redis.Client {
	sshConfig := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("password")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         15 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", "remoteIP:22", sshConfig)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort("127.0.0.1", "6379"),
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshClient.Dial(network, addr)
		},
		// Disable timeouts, because SSH does not support deadlines.
		ReadTimeout:  -1,
		WriteTimeout: -1,
	})
}
