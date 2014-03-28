package server

// Contains code that handles Raft related logic
// for each of the KVStore methods

// Currently the only purpose of methods in this
// files is to ensure that message is replicated
// in RaftServer's log

// raftGet sends a Get message to raft
// and waits for message to come back
// TODO: remove local panic
func (s *kvStore) raftGet(key string) {
	packet := Get{Key: key}
	s.raftLeader.Outbox() <- packet
	le := <-s.raftLeader.Inbox()
	data, ok := le.Data.(Get)
	if !ok || data != packet {
		panic("Received different Get message than expected")
	}
}

func (s *kvStore) raftPut(key string, value string) {
	packet := Put{Key: key, Value: value}
	s.raftLeader.Outbox() <- packet
	le := <-s.raftLeader.Inbox()
	data, ok := le.Data.(Put)
	if !ok || data != packet {
		panic("Received different Get message than expected")
	}
}

func (s *kvStore) raftDelete(key string) {
	packet := Delete{Key: key}
	s.raftLeader.Outbox() <- packet
	le := <-s.raftLeader.Inbox()
	data, ok := le.Data.(Delete)
	if !ok || data != packet {
		panic("Received different Get message than expected")
	}
}
