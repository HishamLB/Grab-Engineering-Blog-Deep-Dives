There is quite a few things I can implement from this article but I'll expand the scope when necessary:


1. User generator for testing.
2. Segmentation Service that creates segments with presets (Account age > 30 days, etc.)
3. Run Containers.
4. Some sort of profiling tool



# Comments

I couldn't get Roaring Bitmaps to work but I did create data structures for run containers, array contaienrs, and normal bitmaps. 
I also tried to do a Greedy conditional to pick DS based on cost but Roaring Bitmaps need to be implemented for this to be tested.
I will definitely be continuing this project as I have learned a lot
