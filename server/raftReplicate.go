// package raftClient contains code that communicates
// with Raft servers
package server

import (
	"encoding/gob"
	"github.com/pkhadilkar/cluster"
	"github.com/pkhadilkar/raft"
	"github.com/pkhadilkar/raft/llog"
	"github.com/pkhadilkar/raft/raftImpl"
	"strconv"
	"time"
)

const serverCount = 5

// Start with the simplest implementation
// Assume that KVStore is only on leader's
// box

func CreateRaftCluster() ([]raft.Raft, error) {
	raftConf, err := raftImpl.ReadConfig("./config.json")
	if err != nil {
		return nil, err
	}

	cluster.NewProxyWithConfig(raftImpl.RaftToClusterConf(raftConf))

	raftServers := make([]raft.Raft, serverCount)

	for i := 0; i < serverCount; i += 1 {
		// create cluster.Server
		clusterServer, err := cluster.NewWithConfig(i, "127.0.0.1", 5000+i, raftImpl.RaftToClusterConf(raftConf))
		if err != nil {
			return nil, err
		}

		logStore, err := llog.Create(raftConf.RaftLogDirectoryPath + "/" + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}

		s, err := raftImpl.NewWithConfig(clusterServer, logStore, raftConf)
		if err != nil {
			return nil, err
		}
		raftServers[i] = s
	}
	// wait for some time to ensure that the servers
	// have elected a leader
	time.Sleep(5 * time.Second)
	return raftServers, err
}

func getRaftLeader(raftServers []raft.Raft) raft.Raft {
	leaderId := raftServers[1].Leader()
	for _, s := range raftServers {
		if s.Pid() == leaderId {
			return s
		}
	}
	return nil
}

func Init() {
	gob.Register(Get{})
	gob.Register(Put{})
	gob.Register(Delete{})
}
