package app

type MapAuth = map[string]interface{}
type MapWeb = map[string]interface{}
type MapRouter = map[string]interface{}
type MapDB = map[string]interface{}

type IVoontiService interface {
	ServiceType() string
	Stat() interface{}
	Policies() interface{}
}

type BaseService struct {
	IVoontiConfig
	UUID        string
	Name        string
	Version     string
	Description string
	Owner       interface{}
	master      IVoontiService
}

type AuthService struct {
	*BaseService
	Name     string
	Version  string
	Policies MapAuth
}

type AppService struct {
	*BaseService
	Name    string
	Version string
}

type DataStoreService struct {
	*BaseService
	Services MapDB
}

type RequestService struct {
	*BaseService
	RequestType string
	Requesters  []string
	Timeout     time.Duration
	Handlers    Map
	Accepts     Map
	Replies     Map
}

type WebService struct {
	*BaseService
	Domain    string
	Port      []string
	SSL       bool
	BaseDir   string
	Container string
	Content   struct {
		Type    []string
		Index   string
		Assets  []string
		Builder string
	}
	Websocket struct {
		Enabled  bool
		Path     string
		Protocol string
	}

	Middlewares []string
}
