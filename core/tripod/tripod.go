package tripod

import (
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/core/env"
	. "github.com/yu-org/yu/core/tripod/dev"
	. "github.com/yu-org/yu/core/types"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Tripod struct {
	*ChainEnv
	*Land

	BlockVerifier BlockVerifier
	TxnChecker    TxnChecker

	Init       Init
	BlockCycle BlockCycle

	Committer Committer

	Instance interface{}

	name string
	// Key: Writing Name
	writings map[string]Writing
	// Key: Reading Name
	readings map[string]Reading
	// key: p2p-handler type code
	P2pHandlers map[int]P2pHandler
}

func NewTripod() *Tripod {
	return NewTripodWithName("")
}

func NewTripodWithName(name string) *Tripod {
	return &Tripod{
		name:        name,
		writings:    make(map[string]Writing),
		readings:    make(map[string]Reading),
		P2pHandlers: make(map[int]P2pHandler),

		BlockVerifier: new(DefaultBlockVerifier),
		TxnChecker:    new(DefaultTxnChecker),
		Init:          new(DefaultInit),
		BlockCycle:    new(DefaultBlockCycle),
		Committer:     new(DefaultCommitter),
	}
}

func (t *Tripod) SetInstance(tripodInstance any) {
	if t.name == "" {
		pkgStruct := reflect.TypeOf(tripodInstance).String()
		strArr := strings.Split(pkgStruct, ".")
		tripodName := strings.ToLower(strArr[len(strArr)-1])
		t.name = tripodName
	}

	if isImplementInterface(tripodInstance, (*TxnChecker)(nil)) {
		t.SetTxnChecker(tripodInstance.(TxnChecker))
	}
	if isImplementInterface(tripodInstance, (*BlockCycle)(nil)) {
		t.SetBlockCycle(tripodInstance.(BlockCycle))
	}
	if isImplementInterface(tripodInstance, (*Committer)(nil)) {
		t.SetCommitter(tripodInstance.(Committer))
	}
	if isImplementInterface(tripodInstance, (*BlockVerifier)(nil)) {
		t.SetBlockVerifier(tripodInstance.(BlockVerifier))
	}
	if isImplementInterface(tripodInstance, (*Init)(nil)) {
		t.SetInit(tripodInstance.(Init))
	}

	for name, _ := range t.writings {
		logrus.Infof("register Writing (%s) into Tripod(%s) \n", name, t.name)
	}

	for name, _ := range t.readings {
		logrus.Infof("register Reading (%s) into Tripod(%s) \n", name, t.name)
	}

	t.Instance = tripodInstance
}

func isImplementInterface(value any, ifacePtr interface{}) bool {
	iface := reflect.TypeOf(ifacePtr).Elem()
	return reflect.TypeOf(value).Implements(iface)
}

func (t *Tripod) Name() string {
	return t.name
}

func (t *Tripod) GetCurrentBlock() (*CompactBlock, error) {
	return t.Chain.GetEndBlock()
}

func (t *Tripod) SetChainEnv(env *ChainEnv) {
	t.ChainEnv = env
}

func (t *Tripod) SetLand(land *Land) {
	t.Land = land
}

func (t *Tripod) SetInit(init Init) {
	t.Init = init
}

func (t *Tripod) SetCommitter(c Committer) {
	t.Committer = c
}

func (t *Tripod) SetBlockCycle(bc BlockCycle) {
	t.BlockCycle = bc
}

func (t *Tripod) SetBlockVerifier(bv BlockVerifier) {
	t.BlockVerifier = bv
}

func (t *Tripod) SetTxnChecker(tc TxnChecker) {
	t.TxnChecker = tc
}

func (t *Tripod) SetWritings(wrs ...Writing) {
	for _, wr := range wrs {
		name := getFuncName(wr)
		t.writings[name] = wr
	}
}

func (t *Tripod) SetReadings(readings ...Reading) {
	for _, r := range readings {
		name := getFuncName(r)
		t.readings[name] = r
	}
}

func (t *Tripod) SetP2pHandler(code int, handler P2pHandler) *Tripod {
	t.P2pHandlers[code] = handler
	logrus.Infof("register P2pHandler(%d) into Tripod(%s) \n", code, t.name)
	return t
}

func getFuncName(i interface{}) string {
	ptr := reflect.ValueOf(i).Pointer()
	nameFull := runtime.FuncForPC(ptr).Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")
	return strings.TrimSuffix(funcName, "-fm")
}

func (t *Tripod) ExistWriting(name string) bool {
	_, ok := t.writings[name]
	return ok
}

func (t *Tripod) GetWriting(name string) Writing {
	return t.writings[name]
}

func (t *Tripod) GetReading(name string) Reading {
	return t.readings[name]
}

func (t *Tripod) AllReadingNames() []string {
	allNames := make([]string, 0)
	for name, _ := range t.readings {
		allNames = append(allNames, name)
	}
	return allNames
}

func (t *Tripod) AllWritingNames() []string {
	allNames := make([]string, 0)
	for name, _ := range t.writings {
		allNames = append(allNames, name)
	}
	return allNames
}
