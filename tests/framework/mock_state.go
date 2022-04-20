package framework

type TestState struct {
	Oracles []*OracleState `json:"oracles"`
}

type OracleState struct {
	Oracle   string `json:"oracle"`
	BlockNum int64  `json:"block"`
	Feed     string `json:"feed"`
}

type StateStore struct {
	Oracle OracleStateStore
}

func NewStateStore() *StateStore {
	return &StateStore{
		Oracle: NewOracleStateStore(),
	}
}

func NewOracleStateStore() OracleStateStore {
	return OracleStateStore{
		Data:  make(map[string][]OracleState),
		Index: make(map[string]int),
	}
}

type OracleStateStore struct {
	Data  map[string][]OracleState
	Index map[string]int
}

func (s *OracleStateStore) GetIndex(oracle string, blockNum int64) int {
	objs := s.Data[oracle]
	var index int
	for i := s.Index[oracle]; i < len(objs); i++ {
		if objs[i].BlockNum <= blockNum {
			index = i
		} else {
			break
		}
	}
	s.Index[oracle] = index
	return index
}

func (s *OracleStateStore) GetState(oracle string, index int) OracleState {
	return s.Data[oracle][index]
}

func (s *OracleStateStore) AddState(obj *OracleState) {
	s.Data[obj.Oracle] = append(s.Data[obj.Oracle], *obj)
}
