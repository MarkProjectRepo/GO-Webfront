# GO-Zeroshot
Quick sentiment classification webfront with python and GO

This was originally going to be a zero-shot classification, but I decided to go with sentiment for the sake of keeping it light.

The point of this project is to try and handle creating a webserver in GO which retrieves data from an autoscaling api on GCP.

See it for yourself running on [GCP](http://34.123.217.221:8080/), at least for as long as I leave this up! If it's not up, here's an example screenshot.

![Sentiment analysis webfront](https://github.com/MarkProjectRepo/GO-Zeroshot/blob/master/images/sentiment_webfront.PNG?raw=true)

What it lacks in style, the backend makes up for in functionality.

## The Cluster
![Screenshot of a gcp dashboard with 2 kubernetes node pools, totalling 4 nodes](https://github.com/MarkProjectRepo/GO-Zeroshot/blob/master/images/nodes.PNG?raw=true)

The above screenshot displays the resources that are being used; Considering it's relatively light, I only have 4 nodes actively running, but if the demand were to increaes,
it would scale only up to a maximum of 10 total.

I made heavy use of [Cortex](https://docs.cortex.dev) which allowed me to point it to the clusters and use them as its execution environment. 
Below you can see, the VM it's pointing to, from which it orchestrates the Kubernetes nodes!
 
![CLI output from Cortex](https://github.com/MarkProjectRepo/GO-Zeroshot/blob/master/images/cortex_env.PNG?raw=true)

And here's a screenshot of Cortex's output for the running service! It's amazing how simple this can be once the environment is hooked up.

![CLI output from Cortex](https://github.com/MarkProjectRepo/GO-Zeroshot/blob/master/images/cortex_output.PNG?raw=true)

Never mind some default names left in place *cough* this project isn't text generation *cough*

## Webfront

The entire webfront is written in pure html and GO! Though this is a pretty hands-on re-introduction to GO, the main point was to try and make use of native functionality
like templates, and have the server handle everything needed to communicate with the backend. I think the code speaks for itself in terms of what it does, but I'll
just finish by saying that it's such an interesting language, and I'd really like to dive deeper into working with it for async processes and seeing how I could
make it handle larger traffic gracefully.
