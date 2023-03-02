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