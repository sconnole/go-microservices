# go-microservices

Is it overkill to have a kubernetes pod that makes one http request with fake sentence data and then terminates? Sure! On a production environment. I'm just over here building something on a weeknight to show it's pretty quick to standup microservices on local environment. 

is this how I would build a production service that handled thousands of requests? Obviously not. Are there optimizations that can be done? Yeah for sure. Did I have fun making it? Yeah I did. 

I haven't built anything in golang for a few months and just wanted a quick refresher. 

I haven't used helm charts for a few months. Wanted a quick refresher. 

Would I build a database layer if I were to make this at scale? Yeah. 

Is it a better practice to have migrations for your database tables? Yeah. 

Should I have structs hanging out in the middle of my `main.go` file? Probably fine for the scope and scale of this local project. 

You get the idea. This wasn't really built to have perfect systems architecting. Just a quick refresher for me on some concepts. ✌️

## 31-Jul-2025 

I've add an LLM model. I've created a python cron job to pull some data that was created with the job creator (clever name I know). Yeah it's a little strange to have random names as the "data" coming back. But here's an example output 

```
Topic: {"id":1,"data":"Mr. Brenden Stracke","status":"PENDING","created_at":"2025-07-31T18:50:10Z"}

Poem:
Still shimmering, still dancin' in the moonlight,
I am Mr. Brende's perfect stranger.
Lost in a dream where he lives and breathes,
Intoxicated by the magic of his eyes.

Murmurs rise from his lips like whispers in hushed tones,
As though the stars above him whisper their secrets.
I am drawn to the warmth of his gentle smile,
A light that shines so brightly it's almost divine.

In his heart, there is something eternal and true,
An anchor that holds us all in a web of peace.
He teaches me how to be more, to find the grace,
To let my soul sing like the night sky above.

I am his reflection, a mirror of his essence,
A symbol of hope, of love, and of life's journey.
As I walk with him, he takes me by the hand,
And we journey through dreams, waking moments bright.

In this perfect stranger, I find my own true home,
Where light shines like a thousand stars above.
stream closed EOF for default/poem-generator-manual-1753987874-6kl62 (poem-generator)
```

And another one: 

```
Topic: {"id":7,"data":"King Craig Quitzon","status":"PENDING","created_at":"2025-07-31T18:50:20Z"}

Poem:
With a mighty heart and a fiery spirit,
Craig Quitzon ruled with grace and might.
In his land, he built a kingdom fair,
And made it shine like never before.

His reign was long and glorious,
A testament to the power of steel.
His wars were victories won in sight,
And his people flourished far beyond.

But alas, time stood still as death drew near,
Craig Quitzon passed away without a say.
A legend of strength and valor he knew,
As his spirit soars on in the sky above.
stream closed EOF for default/poem-generator-29233135-wrkxv (poem-generator)
```