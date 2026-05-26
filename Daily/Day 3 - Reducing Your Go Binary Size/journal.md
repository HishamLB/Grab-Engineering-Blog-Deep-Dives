# [Article Link](https://engineering.grab.com/reducing-your-go-binary-size)

Author(s): Stan Halka, Samuel Thomas

[Implementation](./implementation/README.md)

# Notes:

Grab has over 250+ microservies and some teams manage their own build test and deployment systems.
This leads to broken dependencies, non-reproducible builds, and missing security permissions.
Downtime + managing IaaS gets harder too.

Grab wanted to move to containerisation for these reasons: 

- Support to build and push container image during CI
- Create a standard VM image capable of running container workloads. 
- A deployment method to allow existing services to migrate to containers safely.

They found a large binary and use go-binsize-viz to inspect it.

Steps to analyze the golang binary:

1. Build using:
```go build -a -o service_name ./path/to/main.go```
2. Copy the binary over to ```go-binsize-viz``` repository.
3. Run:
``` 
#!/usr/bin/env bash
#
# This script needs more input parsing, but it serves the needs for now.
#
mkdir dist
# step 1
go tool nm -size $1 | c++filt > dist/$1.symtab
# step 2
python3 tab2pydic.py dist/$1.symtab > dist/$1-map.py
# step 3
# must be data.js
python3 simplify.py dist/$1-map.py > dist/$1-data.js
rm data.js
ln -s dist/$1-data.js data.js
```
4. Run a local python server to visualize the components
```python3 -m http.server```
5. Open http://localhost:8000/treemap_v3.html
6. Analyze!

# Comments/Analysis:
I don't have much experience with Golang as of writing this but a treemap tool to visualize binaries sounds very handy. In this case, it appeared every message format for any service at Grab was included in every service binary so that's a huge reduction in binary sizes overall. There is a very important lesson to learn here and it is the last sentence in the blog. It was hard at first for me to understand but after some digging basically don't mix generic platform code with company-specific implementations in the same connected codebase. Generic and minimal = good. Service-specific logic = outside platform! 

