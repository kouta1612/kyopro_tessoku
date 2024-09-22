package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = 1 << 60
	MOD = int(1e9) + 7
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	n, k := ni2()

	res := 0.0
	for i := 1; i <= n; i++ {
		point := i
		cnt := 0.0

		for point <= k-1 {
			point *= 2
			cnt++
		}

		res += float64(math.Pow(0.5, cnt))
	}

	res /= float64(n)

	fmt.Println(res)
}

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func ni() int {
	sc.Scan()
	return atoi(sc.Text())
}

func ni2() (int, int) {
	return ni(), ni()
}

func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func ni1d(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = ni()
	}
	return res
}

func ni2d(n, m int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = ni()
		}
	}
	return res
}

func nf() float64 {
	sc.Scan()
	return atof(sc.Text())
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func nb() byte {
	sc.Scan()
	return sc.Bytes()[0]
}

func print(v ...interface{}) {
	_, e := fmt.Fprint(out, v...)
	if e != nil {
		panic(e)
	}
}

func println(v ...interface{}) {
	_, e := fmt.Fprintln(out, v...)
	if e != nil {
		panic(e)
	}
}

func printf(format string, v ...interface{}) {
	fmt.Fprintf(out, format, v...)
}

func printIntLn(ns []int) {
	res := make([]interface{}, 0, len(ns))
	for _, v := range ns {
		res = append(res, v)
	}
	println(res...)
}

func printStringLn(ns []string) {
	res := make([]interface{}, 0, len(ns))
	for _, v := range ns {
		res = append(res, v)
	}
	println(res...)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func atof(s string) float64 {
	res, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

/*
指定したスライスの順列を生成する。（要素が重複していてもOK）
ex:

	for {
		// aの処理

		if !next_permutation(a) {
			break
		}
	}
*/
func next_permutation(a []int) bool {
	// a[l] < a[l+1]を満たす最大のlを求める
	l := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			l = i
		}
	}

	if l == -1 {
		return false
	}

	// a[l] < a[r]を満たす最大のrを求める
	r := len(a) - 1
	for i := r; i > l; i-- {
		if a[l] < a[i] {
			r = i
			break
		}
	}

	a[l], a[r] = a[r], a[l]
	// [l+1,n-1]が降順なので逆順にする
	for i := l + 1; i <= (l+len(a))/2; i++ {
		a[i], a[len(a)-(i-l)] = a[len(a)-(i-l)], a[i]
	}

	return true
}

// a[i] >= v を満たす最小のインデックスを取得する
func lowerBound(a []int, v int) int {
	ng, ok := -1, len(a)
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if a[mid] >= v {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// a[i] > v を満たす最小のインデックスを取得する
func upperBound(a []int, v int) int {
	ng, ok := -1, len(a)
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if a[mid] > v {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func ceil(a, b int) int {
	return (a + b - 1) / b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func chmax(x *int, y int) {
	*x = max(*x, y)
}

func chmin(x *int, y int) {
	*x = min(*x, y)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

type unionfind struct {
	// par[x]: 要素 x の親頂点の番号 (自身が根の場合は −1)
	// rank[x]: 要素 x の属する根付き木の高さ
	// siz[x]: 要素 x の属する根付き木に含まれる頂点数
	par, rank, siz []int
}

func newUnionFind(n int) *unionfind {
	result := new(unionfind)
	result.par, result.rank, result.siz = make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		result.par[i] = -1
		result.rank[i] = 0
		result.siz[i] = 1
	}
	return result
}

func (u *unionfind) root(x int) int {
	if u.par[x] == -1 {
		return x
	}
	// 経路圧縮
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u *unionfind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func (u *unionfind) unite(x, y int) bool {
	rx, ry := u.root(x), u.root(y)
	if rx == ry {
		return false
	}

	// union by rank
	if u.rank[x] < u.rank[y] {
		rx, ry = ry, rx
	}
	u.par[ry] = rx
	if u.rank[rx] == u.rank[ry] {
		u.rank[rx]++
	}
	u.siz[rx] += u.siz[ry]

	return true
}

func (u *unionfind) size(x int) int {
	return u.siz[u.root(x)]
}

type mint int

func (m mint) add(n int) mint {
	n = n % MOD
	return mint(int(m) + n).mod()
}

func (m mint) sub(n int) mint {
	return mint(int(m) - n + MOD).mod()
}

func (m mint) mul(n int) mint {
	n = n % MOD
	return mint(int(m) * n).mod()
}

func (m mint) div(n int) mint {
	n = n % MOD
	d := inv(n)
	return m.mul(d).mod()
}

func (m mint) mod() mint {
	return mint(int(m) % MOD)
}

func inv(n int) int {
	return mpow(n, MOD-2, MOD)
}

func mpow(a, b, p int) int {
	if b == 0 {
		return 1
	}

	t := mpow(a, b/2, p)
	if b%2 == 0 {
		return (t * t) % p
	} else {
		return (((t * t) % p) * a) % p
	}
}

func combination(n int, k int) int {
	if n-k < k {
		k = n - k
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
		v /= (i + 1)
	}
	return v
}

func modcombination(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	if n-k < k {
		k = n - k
	}
	v := mint(1)
	for i := 0; i < k; i++ {
		v = v.mul(n - i).div(i + 1)
	}
	return int(v)
}