# Ringbl [WIP]
A consistent hash ring with bounded loads
Consistent hashing with bounded loads implementation using Red Black Tree 
![ringbl](https://3.bp.blogspot.com/-pgZ4b9H7VlM/WOJ91rDe_XI/AAAAAAAABqw/wIjtyPHheFgyHpXIqY4qNLhd_H9DnHsXACLcB/s640/image00.png)

## Example Usage

```go
ring:=NewRing([]string{"server-1","server-2","server-3"},1)
node:=ring.Get("foo")
```


## TODO

- Test cases
- Performance test for xxhash

## Paper
https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html

https://www.akamai.com/es/es/multimedia/documents/technical-publication/consistent-hashing-and-random-trees-distributed-caching-protocols-for-relieving-hot-spots-on-the-world-wide-web-technical-publication.pdf


Ringbl Image source: https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
