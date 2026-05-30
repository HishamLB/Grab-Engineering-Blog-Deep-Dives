# [Article Link](https://engineering.grab.com/preventing-pipeline-calls-from-crashing-redis-clusters)

Author(s): Micheal Cartmell, Jihao Huang, Sandeep Kumar 

[Implementation](./implementation/README.md)

# Notes:
A redis node failed, required replacement, and caused significant downtime.
Appolo is Grab's internal driver-side state machine and is important to booking transport etc. It stores driver availability in an AWS ElastiCache Redis Cluster. 
Since Appolo's availability is crucial, its Redis Cluster has 3 shards each with 2 slaves. It hashes all keys and replicates into multiple partitions for reliability.

Steps to replicate the outage startinf from setting a local Redis Cluster: 

1. Download Redis
2. Setup config files for each node: 
```   
port 600x

cluster-enabled yes

cluster-config-file cluster-node-x.conf

cluster-node-timeout 5000

appendonly yes

appendfilename node-x.aof

dbfilename dump-x.rdb
```
3. Initiate each node: 
4. Use a Ruby script to create cluster 
```
$PATH/redis-4.0.9/src/redis-trib.rb create --replicas 2127.0.0.1:6001..... 127.0.0.1:6009

>>> Performing Cluster Check (using node 127.0.0.1:6001)

M: 7b4a5d9a421d45714e533618e4a2b3becc5f8913 127.0.0.1:6001

   slots:0-5460 (5461 slots) master

   2 additional replica(s)

S: 07272db642467a07d515367c677e3e3428b7b998 127.0.0.1:6007

   slots: (0 slots) slave

   replicates 05363c0ad70a2993db893434b9f61983a6fc0bf8

S: 65a9b839cd18dcae9b5c4f310b05af7627f2185b 127.0.0.1:6004

   slots: (0 slots) slave

   replicates 7b4a5d9a421d45714e533618e4a2b3becc5f8913

M: 05363c0ad70a2993db893434b9f61983a6fc0bf8 127.0.0.1:6003

   slots:10923-16383 (5461 slots) master

   2 additional replica(s)

S: a78586a7343be88393fe40498609734b787d3b01 127.0.0.1:6006

   slots: (0 slots) slave

   replicates 72306f44d3ffa773810c810cfdd53c856cfda893

S: e94c150d910997e90ea6f1100034af7e8b3e0cdf 127.0.0.1:6005

   slots: (0 slots) slave

   replicates 05363c0ad70a2993db893434b9f61983a6fc0bf8

M: 72306f44d3ffa773810c810cfdd53c856cfda893 127.0.0.1:6002

   slots:5461-10922 (5462 slots) master

   2 additional replica(s)

S: ac6ffbf25f48b1726fe8d5c4ac7597d07987bcd7 127.0.0.1:6009

   slots: (0 slots) slave

   replicates 7b4a5d9a421d45714e533618e4a2b3becc5f8913

S: bc56b2960018032d0707307725766ec81e7d43d9 127.0.0.1:6008

   slots: (0 slots) slave

   replicates 72306f44d3ffa773810c810cfdd53c856cfda893

[OK] All nodes agree about slots configuration.
```
5. Try to send a query: 
```$PATH/redis-4.0.9/src/redis-cli -c -p 6001 hset driverID 100 state available updated_at 11111```

When a slave node is unreachable the cluster boradcasts to all nodes to not send anything to that port. When a master node is unreachable the cluster promotes a slave to a master. 
So why did the outage happen?
Apparantly the implementation of the ```Go-Redis``` library as well as some ```err != nil``` caused the entire pipeline to fail when a slave is unreachable.
It took a whole minute too because the Redis Cluster only refreshes the state once a minute! Lucky for Grab no data was lost as write queries are only sent to Master nodes and master nodes have a watcher attached to them instead of being polled/refreshed every minute.

This can be fixed in a number of ways including reloading the nodes' status when an error happens. 

# Comments/Analysis:

Very cool article. Learned a bit more about Redis and that sometimes libraries implementing technologies might have their own quirks that don't seem obvious/need fixing. 
