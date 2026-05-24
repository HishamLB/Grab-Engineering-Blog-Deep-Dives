# [Article Link](https://engineering.grab.com/programmers-beware-ux-is-not-just-for-designers)
Author(s): [Corey Scott](https://coreyscott.dev/) 

[Implementation](./implementation/README.md)

# Notes:

UX can be discovered through asking these 5 questions
- Who/What is the user?
- What do they want to achieve? - Apply 80/20 rule: 80% of users will want to do 1 thing while the rest 20%
- What are they capable of? - Skills/experience/domain knowledge
- What can I do to make their life easier? 
- Is there anything similar out there that the user already knows how to use?

Designing a REST API we can think of it in the context of the above 5 questions. 
Doing this may have us arrive at a non RESTful approach but with better UX both explicit (seeing the data/using the app) and implicit (less cost and less battery)
 
Context of Code:

Consider: 
```AddBalance(5, false)```
What does false mean?

Apply the 5 questions:
1. Who/What is the user? - Your future self/teammates
2. What do they want to achieve? - Use your code
3. What are they capable of? - Assume lower skill level and a bit less domain knowledge
4. What can I do to make their life easier? - Ask: what interface would allow future me to use this without thinking?
5. Is there anything similar out there? - Consistency is important

Solution:
Replace with 2 functions:
```AddBalanceCreateIfMissing(5)```
```AddBalanceFailOnMissing(5)```

# Comments/Analysis:
Very insightful. A lot of this resonated with me since UX is something that I've always had in the back of my mind but also never bothered really considering at every level/scale.
UX is often linked to Product Designers/UI Designers but Corey Scott gives an outlook that abstracts who the "User" is in UX and by doing that makes it possible to apply UX principles to a wide variety of stages.
The engineers reading your code in the future are users of your code and so asking the template 5 questions when designing code nets interesting design considerations.
Something I found interesting too was how sacrificing design patterns can sometimes result in better UX. In the example given where 2 different "entities" are to be served, Corey considered sacrificing RESTful correctness just for performance and/or better UX.

UX of APIs is something that I'm definitely now more inclined to practice and get down but like it was mentioned in the article iteration is what brews the best UX. 
