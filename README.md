# XRing
A consistent hash ring with bounded loads
Consistent hashing with bounded loads implementation using Red Black Tree 
![XRing](https://3.bp.blogspot.com/-pgZ4b9H7VlM/WOJ91rDe_XI/AAAAAAAABqw/wIjtyPHheFgyHpXIqY4qNLhd_H9DnHsXACLcB/s640/image00.png)

## Example Usage

```go
	nodes := []string{"a", "b", "c"}
	cnf := &xring.Config{
		VirtualNodes: 300,
		LoadFactor:   2,
	}
	hashRing := xring.NewRing(nodes, cnf)
	node,err:=hashRing.Get("foo")
```


## TODO

- Add more test cases
- Performance test for xxhash

## Paper
https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html

https://www.akamai.com/es/es/multimedia/documents/technical-publication/consistent-hashing-and-random-trees-distributed-caching-protocols-for-relieving-hot-spots-on-the-world-wide-web-technical-publication.pdf


XRing Image source: https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
