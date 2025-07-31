# go-microservices

Is it overkill to have a kubernetes pod that makes one http request with fake sentence data and then terminates? Sure! On a production environment. I'm just over here building something on a weeknight to show it's pretty quick to standup microservices on local environment. 

is this how I would build a production service that handled thousands of requests? Obviously not. Are there optimizations that can be done? Yeah for sure. Did I have fun making it? Yeah I did. 

I haven't built anything in golang for a few months and just wanted a quick refresher. 

I haven't used helm charts for a few months. Wanted a quick refresher. 

Would I build a database layer if I were to make this at scale? Yeah. 

Is it a better practice to have migrations for your database tables? Yeah. 

Should I have structs hanging out in the middle of my `main.go` file? Probably fine for the scope and scale of this local project. 

You get the idea. This wasn't really built to have perfect systems architecting. Just a quick refresher for me on some concepts. ✌️