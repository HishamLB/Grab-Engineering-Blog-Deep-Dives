# Background 
So I'll be using a past project to apply the principles I learned from the article.

The project is a Spring Boot Kotlin project and I examined the Controller classes to look for any examples I can apply the 5 questions on. 
This function was chosen: 
```Kotlin
@PostMapping("/logs/search")
fun searchLogs(@RequestBody filter: ManagerAuditLogFilterDTO): ResponseEntity<ManagerAuditLogSearchResponseDTO> {
    val managerId = authenticatedManagerId()
        val logs = auditService.searchLogs(
                managerUUID = managerId,
                actionType = filter.actionType,
                outcome = filter.outcome,
                keyword = filter.keyword,
                from = filter.from,
                to = filter.to
                )
        val actionTypes = logs.mapNotNull { it.actionType }.distinct().sorted()
        val outcomes = logs.mapNotNull { it.outcome }.distinct().sorted()
        return ResponseEntity.ok(
                ManagerAuditLogSearchResponseDTO(
                    logs = logs.map { log ->
                    ManagerAuditLogSummaryDTO(
                            logId = log.logId,
                            timestamp = log.timestamp.toString(),
                            title = auditService.toDisplayTitle(log.actionType),
                            actionType = log.actionType,
                            category = auditService.toCategory(log.actionType),
                            outcome = log.outcome,
                            resource = auditService.toResourceLabel(log.targetType),
                            actor = log.actorId ?: "anonymous",
                            actorType = auditService.toActorType(log.actionType),
                            detailsPreview = log.details
                            )
                    },
                    totalItems = logs.size.toLong(),
                    availableActionTypes = actionTypes,
                    availableOutcomes = outcomes
                    )
                )
}
```

The issue here is that the controller is doing more than necessary by manually assembling the raw logs.
Lets apply the 5 questions:
1. Who/What is the user? - The controller using ```auditService.searchLogs()``` and future me!
2. What do they want to achieve? - Search audit logs and return something usable by the frontend
3. What are they capable of? - The consumer (controller in this case) is burdened with knowing a lot this invites future misuse by developers
4. What can I do to make their life easier? - Move mapping to service layer.
5. Is there anything similar users already know how to use? - The controllers usually dont use their own mappings.

Solution: 
Give the ```Controller``` the mapped logs by introducing a helper function in the service layer.
For example:

```Kotlin 
fun toSummaryDTO(log: AuditLog): ManagerAuditLogSummaryDTO
```
