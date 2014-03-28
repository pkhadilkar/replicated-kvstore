package server

// message types for each type of operation in Raft

type Get struct {
	Key string
}

type Put struct {
	Key   string
	Value string
}

type Delete struct {
	Key string
}
