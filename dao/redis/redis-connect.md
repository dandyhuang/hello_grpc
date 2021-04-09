## redis连接池


### 连接
```golang

func DialContext(ctx context.Context, network, address string, options ...DialOption)

```

type Pool struct {
	// 连接
	Dial func() (Conn, error)
	DialContext func(ctx context.Context) (Conn, error)
	TestOnBorrow func(c Conn, t time.Time) error
	// 最大空闲连接 
	MaxIdle int
	// 最大连接数 0为无限
	MaxActive int
	IdleTimeout time.Duration
	Wait bool
	MaxConnLifetime time.Duration
	mu           sync.Mutex    // mu protects the following fields
	closed       bool          // set to true when the pool is closed.
	active       int           // the number of open connections in the pool
	initOnce     sync.Once     // the init ch once func
	ch           chan struct{} // limits open connections when p.Wait is true
	idle         idleList      // idle connections
	waitCount    int64         // total number of connections waited for.
	waitDuration time.Duration // total time waited for new connections.
}
```