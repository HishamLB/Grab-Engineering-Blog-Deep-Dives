# Background 
Admittedly, I was not even sure if I should include and implementation section for this one. 
There's 2 main things discovered in the article: The tool (visualizer) and the idea (solution). 

The former I can simply run on a go binary and look further, the latter requires a whole architecture study of some project.

Let's do both!
But not the same project for both since I do not have a large enough Go project yet.

# go-binsize-viz
Unfortunately I cannot provide background for the go project I will be using for one simple reason: I can't.
I can still provide my findings using the visualizer tool!

It works! 

Surprisingly intuitive setup. Let's take a closer look. 
```crypto/``` is consuming ~35 MB out of the 39 MB binary (90% of it!)

Honestly, I had to double check my go file because I do not remember needing to use anything in crypto/! But oh wait, htttp. Surely that is what pulled crypto. FIPS was also under crypto, which makes sense.
Apparantly Go put FIPS 140 support directly into standard library but I'm not sure if that's why it's pulling it and I probably do need crypto anyways.


# Architecture review
Background is it is a Jetpack Compose project. That's all.

For a part of the project (it is ugly):
├── presentation/
│   ├── admin/
│   ├── auth/
│   │   ├── JoinCommunityScreen.kt
│   │   └── SignUpScreen.kt
│   ├── dashboard/
│   │   ├── CampusDashboardScreen.kt
│   │   ├── DashboardSection.kt
│   │   ├── components/   [THERE ARE 36 FILES HERE]
│   │   ├── models/
│   │   └── utils/
│   ├── manager/
│   │   ├── ManagerAnalyticsScreen.kt
│   │   ├── ManagerAuditLogsScreen.kt
│   │   ├── ManagerCommunityBoundsScreen.kt
│   │   ├── ManagerCommunitySetupScreen.kt
│   │   ├── ManagerConfigureMetricsScreen.kt
│   │   ├── ManagerConfigureSpacesScreen.kt
│   │   ├── ManagerSafetyIncidentsScreen.kt
│   │   └── ManagerTrustControlsScreen.kt
│   ├── navigation/
│   └── reports/

The fix is a bit obvious refactor components and utils so dashboard does not own them! 
