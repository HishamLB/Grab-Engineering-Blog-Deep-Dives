# [Article Link](https://engineering.grab.com/streamlining-grabs-segmentation-platform)

Author(s): Jake Ng

[Implementation](./implementation/README.md)

# Notes:
User Segmentation is the process of dividing passengers, driver partners into sub-groups based on certait attributes.
Marketing campaigns and specific experience can be tailored based on the segment the user belongs to.
1. Segment Creation:
This is powered by Apache's Spark jobs. When a segment is created a job retrieves data from the data lake. After validation and cleansing the spark job calls subsystems to populate the segment with users
2. Segment Serving:
A NoSQL ScyllaDB is used to store the id, and segment name.

There are problems with this as the write QPS (Queries Per Second) became a bottleneck. Teams would have to wait hours to create the segments. 
Solution:
## Segments as bitmaps
Since a segment is stored across multiple rows (user id : segment name) this causes a huge number of writes to the db. 
Storing index/bit as a bitmap would solve this. To check for membership a bitwise operation can be used to check if the bit at the user id's index is 0 or 1. However, if we had 2 users user id 1 and user id 200000 it will require a bitmap with 200000 bits.
## Roaring Bitmaps
Compressed uint32 bitmaps. Uses 3 different data structures based on the distribution within the chunk this achieves good compression.
## Array Containers
These are used when data is spares (<= 4096 values).It is a sorted array of 16 bit integers which is memory efficient for sparse data and provides log-time access.
## Run Containers
When a chunk has long consecutive values. They use RLE (run-length encoding) to reduce storage required. They store 2 values representing the start and end values. 
# Comments/Analysis:

Interesting read, problems at enterprise-level seem to deal a lot with performance and efficiency which makes sense. This blog post had me looking up a few things and revising certain Data Structures mentioned. There are some details I felt I missed but that is good practice for implementation. 


