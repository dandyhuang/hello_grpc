package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/yangwenmai/ratelimit/leakybucket"
	"github.com/yangwenmai/ratelimit/simpleratelimit"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

func TestUberLimter(t *testing.T) {
	i := 10
	ip := &i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)

	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

func testContext() {
	ctx, cacel := context.WithTimeout(context.Background(), 3)
	defer func() {
		cacel()
		fmt.Println("canclee=====")
	}()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("cacel:== ")
			time.Sleep(time.Second)
			// cancle 已经退出了，还在处理
			fmt.Println("cacel:==sdljdflkjdklsfjdfskl ")
		}
	}()
	fmt.Println("cacel:== d")
}

// 令牌桶分析 https://cloud.tencent.com/developer/article/1900260
func TestLimter(t *testing.T) {
	testContext()
	tmp := time.Millisecond * 31
	now := time.Now()
	fmt.Println("now:", now)
	time.Sleep(time.Second)
	last := now
	// now = time.Now()
	fmt.Println("now 111 :", now, "d1:", time.Duration(1<<63-1))
	fmt.Println("now:", now.Before(last), " ddsd:", now)
	fmt.Println("now:", now.Sub(last), " ddsd2===22:", now)
	fmt.Println(tmp.Seconds(), 1/tmp.Seconds(), 1/float64(31/1000.0))

	limiter := rate.NewLimiter(rate.Every(time.Millisecond*31), 2)
	//time.Sleep(time.Second)
	for i := 0; i < 4; i++ {
		var ok bool
		if limiter.Allow() {
			ok = true
		}
		time.Sleep(time.Millisecond * 20)
		fmt.Println(ok, limiter.Burst())
	}
}

func TestLimter1(t *testing.T) {
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*10), 2)
	for i := 0; i < 10; i++ {
		var ok bool
		if limiter.Allow() {
			ok = true
		}
		time.Sleep(time.Millisecond * 3)
		fmt.Println(ok, limiter.Burst())
	}
}

func TestSimple(t *testing.T) {
	var limits uint64
	var dec uint64
	dec = 3
	limits = 2
	log.Println("limit===", atomic.AddUint64(&limits, -dec))
	// rate limit: simple
	rl := simpleratelimit.New(10, time.Second)

	for i := 0; i < 100; i++ {
		log.Printf("limit result: %v\n", rl.Limit())
	}
	log.Printf("limit result: %v\n", rl.Limit())

	// rate limit: leaky-bucket
	lb := leakybucket.New()
	b, err := lb.Create("leaky_bucket", 10, time.Second)
	if err != nil {
		log.Println(err)
	}
	log.Printf("bucket capacity:%v", b.Capacity())

	// rate limit: token-bucket
}

//[10,1,2,7,6,1,5]
//8
// [[1,1,6],[1,2,5],[1,7],[1,2,5],[1,7],[2,6]]
// 输出 [[1,2,5],[1,7],[2,6]] 有去重
// [[1,1,6],[1,2,5],[1,7],[2,6]]
// 输出 [[1,1,6],[1,2,5],[1,7],[1,2,5],[1,7],[2,6]]
// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
// s = "aab"
//返回 s 所有可能的分割方案。 输出：[["a","a","b"],["aa","b"]]
var res [][]string
func dfs(start, n int,  s string, arr []string) {
	if start == n {
		tmp:=make([]string,len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		return
	}
	for i:= start; i < n; i++ {
		//if !isPartition(s, start, i) {
		//	continue
		//}
		arr = append(arr, s[start:i+1])
		dfs(i+1, n, s, arr)
		arr = arr[:len(arr)-1]
	}
}

func partition(s string) [][]string {
	arr:=make([]string, 0)
	res =make([][]string, 0)
	dfs(0, len(s), s, arr)
	return res
}

func isPartition(s string,startIndex,end int)bool{
	left:=startIndex
	right:=end

	for ;left<right;{
		if s[left]!=s[right]{
			return false
		}
		left++
		right--
	}
	return true
}
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func TestRedis(t *testing.T) {
	str := "123"
	index, _ := strconv.Atoi(string(str[0]))
	fmt.Println("index", index, string(str[0:len(str)-1]))
	var address []string
	addr:="12"
	flag.StringVar(&addr, "addr", "addrss", "配置文件")
//	addr:="game-redis-pre-pre0.redis.dba.vivo.lan.:11282"
	address = append(address, addr)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
		//MaxRedirects: c.config.MaxRetries,
		//ReadOnly:     c.config.ReadOnly,
		//Password:     c.config.Password,
		//MaxRetries:   c.config.MaxRetries,
		//DialTimeout:  c.config.DialTimeout,
		//ReadTimeout:  c.config.ReadTimeout,
		//WriteTimeout: c.config.WriteTimeout,
		//PoolSize:     c.config.PoolSize,
		//MinIdleConns: c.config.MinIdleConns,
		//IdleTimeout:  c.config.IdleTimeout,
	})
	ctx:=context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis ping", err)
		os.Exit(1)
	}
	pool, _ := ants.NewPool(10000, ants.WithNonblocking(true))
	defer pool.Release()
	for i:=0; i < 1000000; i++ {
		err := ants.Submit(func() {
			key:=RandStringBytesMaskImprSrc(10)
			rdb.SetEX(ctx,key, strings.Repeat("A",100000), time.Second * 100)
		})
		if err != nil {
			fmt.Println("err:", err)
		}
	}
}
