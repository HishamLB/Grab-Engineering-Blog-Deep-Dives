# [Article Link](https://engineering.grab.com/domain-driven-development-in-golang)

Author(s): Kapil Chaurasia, Preeti Karkera


# Notes:
Files had responsibilities mixed together with imports everywhere and new devs struggling to understand the system.
The plan was to reorganize the code around the business domain using DDD.
DDD means structuring the code around real business concepts instead of technical details. So instead of "service and repository" we focus on what the business actually does.
1. Set and figure out the business entities: Project, Developer, Product.
2. Gather domain knowledge using elicitation/talking to them.
3. Create bounded contexts to see what each section is responsible for.
4. Identify entities.

Repositories also are meant to abstract data access so instead of ```db.Query ``` it was made to be ``` GetProjectByUUID() ```

# Comments/Analysis:
Seems like shying away from these technical patterns and standards ends up working better in big enterprises, at least that's the pattern I've been seeing so far.
DDD seems like it ignores the fancy patterns and just focuses on purely understanding the business deeply. 
I think the key takeway is to organize software around business behaviour instead of technical layers. So much for learning these architectural design patterns!
