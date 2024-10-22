package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/deque"
)

const (
	INF = 1 << 60
	MOD = int(1e9) + 7
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type point struct {
	x, y int
}

var dx []int = []int{0, 1, -1, 0}
var dy []int = []int{1, 0, 0, -1}
var dxx []int = []int{-2, -2, -2, -2, -2, -1, -1, -1, -1, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2}
var dyy []int = []int{-2, -1, 0, 1, 2, -2, -1, 1, 2, -2, 2, -2, -1, 1, 2, -2, -1, 0, 1, 2}

func main() {
	defer out.Flush()

	h, w := ni2()
	s, t := point{x: ni() - 1, y: ni() - 1}, point{x: ni() - 1, y: ni() - 1}
	grid := make([][]string, h)
	for i := 0; i < h; i++ {
		grid[i] = strings.Split(ns(), "")
	}

	dist, used := make([][]int, h), make([][]bool, h)
	for i := 0; i < h; i++ {
		dist[i], used[i] = make([]int, w), make([]bool, w)
		for j := 0; j < w; j++ {
			dist[i][j] = INF
		}
	}
	dist[s.x][s.y] = 0

	deq := deque.New[point]()
	deq.PushBack(point{x: s.x, y: s.y})
	for deq.Size() > 0 {
		u := deq.PopFront()
		now := point{x: u.x, y: u.y}
		used[now.x][now.y] = true

		for i := 0; i < 4; i++ {
			nx := now.x + dx[i]
			ny := now.y + dy[i]
			if !(0 <= nx && nx < h && 0 <= ny && ny < w) {
				continue
			}
			if grid[nx][ny] == "#" {
				continue
			}
			if used[nx][ny] {
				continue
			}

			if dist[nx][ny] > dist[now.x][now.y] {
				dist[nx][ny] = dist[now.x][now.y]
				deq.PushFront(point{x: nx, y: ny})
			}
		}
		for i := 0; i < len(dxx); i++ {
			nx := now.x + dxx[i]
			ny := now.y + dyy[i]
			if !(0 <= nx && nx < h && 0 <= ny && ny < w) {
				continue
			}
			if grid[nx][ny] == "#" {
				continue
			}
			if used[nx][ny] {
				continue
			}

			if dist[nx][ny] > dist[now.x][now.y]+1 {
				dist[nx][ny] = dist[now.x][now.y] + 1
				deq.PushBack(point{x: nx, y: ny})
			}
		}
	}

	if dist[t.x][t.y] == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[t.x][t.y])
	}
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

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
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

/*
単調増加部分列の長さを返す
strict = falseのとき、同じ値が連続することを許す
*/
func lis(a []int, strict bool) int {
	dp := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = INF
	}

	for i := 0; i < len(a); i++ {
		var it int
		if strict {
			it = lowerBound(dp, a[i])
		} else {
			it = upperBound(dp, a[i])
		}
		dp[it] = a[i]
	}

	return lowerBound(dp, INF)
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

// nを素因数分解 => map[素因数]個数
func factorize(n int) map[int]int {
	result := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}

		result[i] = 0
		for n%i == 0 {
			result[i]++
			n /= i
		}
	}

	if n != 1 {
		result[n] = 1
	}

	return result
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
	return mint((int(m) + MOD) % MOD)
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
	if n < k || n < 0 || k < 0 {
		return 0
	}
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
	if n < k || n < 0 || k < 0 {
		return 0
	}
	v := mint(1)
	for i := 0; i < k; i++ {
		v = v.mul(n - i).div(i + 1)
	}
	return int(v)
}

type modcomb struct {
	n          int
	factorizes []int
}

func newModcomb(n int) *modcomb {
	factorizes := make([]int, n+1)
	factorizes[0] = 1
	for i := 1; i <= n; i++ {
		factorizes[i] = (factorizes[i-1] * i) % MOD
	}

	return &modcomb{n: n, factorizes: factorizes}
}

func (m *modcomb) combination(n, k int) int {
	if n < k || n < 0 || k < 0 {
		return 0
	}
	result := mint(m.factorizes[n]).div(m.factorizes[k]).div(m.factorizes[n-k])
	return int(result)
}
