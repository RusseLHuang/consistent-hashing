# Consistent Hashing Implementation

## Story
While it is easier to scale when using NoSQL database, it's a little bit tricky when you need to scale for write operation in Relational Database. The core idea it pretty much the same to use sharding, it needs to implement a load balancer on application layer.

The problem that consistent hashing solved is to minimize data movement on data distribution when adding or removing a node.

## MySQL Check Open Connections
1. mysql cli
Login
> mysql --user=${username} --password payment

> SHOW STATUS WHERE `variable_name` = 'Threads_connected';


## Hash Value
nodes-1 : 0
nodes-2 : 64700580

National Id
- NATIONALIDA : 98255809
- NATIONALIDB : 47922952
- NATIONALIDC : 64700571
