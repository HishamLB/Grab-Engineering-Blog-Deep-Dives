The plan is simple! 
Fix the bug in the Redis Cluster Library.
Lucky for me Grab already mentioned a way to fix it so I want to try to. 

# Reproduce The Issue

Before we tackle fixing anything we need to be able to reproduce the behaviour that we're trying to fix. 
Lucky for me, once again, Grab provided some of that.
We do need to write our own Go script however.

1. Create a bash script file to create and start the Redis Cluster
2. Create a go script that populates and checks the status of the Redisc Cluster
3. Observe

I instantly ran into an issue, or well, lack of an issue. 
Apparantly, or at least according to what I could find, this whole issue was fixed! 
Well then, we don't have to fix it but let's at least replicate the issue and look closer. 

To do this I had to grab and old version of ```Redis-Go```, ```v6``` to be exact. (7 years old). 
I ran into some issues using that old of a version mostly about Valkey but other than that I could replicate it! 

I could also replicate the ~60s till the Cluster reads the state of a slave as well as shown here: 

