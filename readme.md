# Consistent Hashing

Assume we have "n" server, the basic way to select server is `hash_key % n`

| server   | key |
|----------| ---|
| server 1 | 0, 3|
| server 2 | 1, 4|
| server 3 | 2, 5|

If we add one more server into pool,

| server   | key |
|----------|-----|
| server 1 | 0, 4 |
| server 2 | 1, 5 |
| server 3 | 2,  |
| server 4 | 3   |

3,4,5 are remapped

If we remove one server 

| server   | key     |
|----------|---------|
| server 1 | 0, 2, 4 |
| server 2 | 1, 3, 5 |

2,3,4,5 are remapped, it seems bad when the server was removed

## Hash Ring
The hash ring is essentially a circular array of nodes, where each node is assigned a position on the ring based on the result of a hash function applied to its identifier. This ensures that each node has a unique position on the ring, and that the positions are distributed evenly around the circle.

## Virtual Node
In a hash-based partitioning scheme, the partitioning of data is determined by the result of applying a hash function to the key of the data. Each node in the system is responsible for a subset of the total hash range, and any data with a key that hashes to a value within that range is stored on that node.

With virtual nodes, each physical node in the system is associated with multiple virtual nodes, each of which is responsible for a portion of the hash range. This allows for finer-grained partitioning of the data and better load balancing, since each physical node can be responsible for a larger number of virtual nodes.

## Summary
1. Minimized keys are redistributed when serves are added/removed
2. It's easy to scale horizontally because data are more evenly distributed
3. Mitigate hotspot key problem.


## Appendix
### Jump Consistent Hash
Jump Consistent Hash is an efficient hash function used for mapping data to different nodes in distributed systems.

Unlike traditional hash functions, Jump Consistent Hash does not calculate a hash value, but instead maps data keys to nodes using a series of jump operations. The main idea behind the function is that given a data key and a number of nodes n, the key can be mapped to a node in the range [0, n) through a series of simple jumps.

Specifically, Jump Consistent Hash first combines the data key with a random number, and then maps the result to a node through a series of jumps. In each jump operation, the function calculates an offset pointing to the next node, and adds it to the current node's number to get the next node's number. This process repeats until a final node number is obtained.

The main advantage of Jump Consistent Hash is that it maintains good load balancing while requiring only O(1) time and space complexity, making it very suitable for use in large-scale distributed systems for hash tables and data sharding. It can also handle node addition and removal, as well as changes in network topology, effectively because it does not need to recompute hash values for all keys.

code: 
[lithammer / go-jump-consistent-hash](https://github.com/lithammer/go-jump-consistent-hash)

