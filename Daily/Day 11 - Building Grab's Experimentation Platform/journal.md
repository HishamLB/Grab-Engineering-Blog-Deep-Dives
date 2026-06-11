# [Article Link](https://engineering.grab.com/building-grab-s-experimentation-platform)

Author(s): Abeesh Thomas, Roman Atachiants

[Implementation](./implementation/README.md)

# Notes:

A lot of the improvements at Grab come from successful experiments that they run. 
To make it easier Grab built a platform to:
- Create a unified experimentation platform 
- Allow simple experiments
- Automate cohort selection.
- Support power analysis for tests.
- Create triggers/alerts to notify them of specific effects from an experiment.
- Centralized UI making it easier to manage and create an experiment.

A geo-temporal segmentation for concurrent experiments is used to prevent any experiment interference. Related work here: [Google's Paper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/36500.pdf)

An experiment is abstracted into: 
- A variable: which is something that can be changed like a different payment method enabled for a specific city or user.
- A metric: which is something that can be observed and tracked. Like cancellation rate or revenue.

An experiment is formalized as a time-bound configuration. 
The apps SDKs apply the experiment and avoid excessive network calls.


# Comments/Analysis:

I've heard that Grab runs a lot of experiments and with the nature of how they operate it makes sense to have a centralized system and procedure to be able to observe these experiments.
The architecture described here felt similar to the one described in [Day 7's](../Day%207%20-%20Reliable%20and%20Scalable%20Feature%20Toggles%20and%20A%20B%20Testing%20SDK%20at%20Grab/) blog which makes sense because one of the authors (Roman Atachiants) was on that blog! 
I guess I should've read this before Day 7's, maybe I will take more into account the blogs' release date since some of them seem like they are a part of a series.
I wish more was explained about how the SDK "intelligently" applies experiments without doing any costly network calls. I assume it's some method locally that checks a hash or so. 

