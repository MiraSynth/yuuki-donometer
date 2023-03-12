# Yuuki Donometer
A little app that connect to GoFundMe, gets donation data and displays a meter and an alert box with sounds and everything!

> Consider donating to Yuuki: https://gofund.me/60c780d1

<p align="center">
  <img src="assets/headpat-yuuki.gif" width=300 />
</p>


<br />

## Installing
- Go to the [latest release](https://github.com/MiraSynth/yuuki-donometer/releases/tag/latest) page
- Download the `goyuuki.exe` file
- Double click the file to run it
- When windows defeder complains, click on `More info` and then `Run anyway`
- A command line will appear, follow it's instruction.

## Pages
### *Alert Box*
Copy the url into an OBS browser source.
- http://localhost:8969/alert

To style add the following to the Custom CSS field
```css
:root {
    --alert-text-color: #ffffff;
    --alert-text-size: 3em;
}
```

<br />

### *Donation Meter*
Copy the url into an OBS browser source.
- http://localhost:8969/meter

 To style add the following to the Custom CSS field
```css
:root {
    --meter-height: 3em;
    --meter-color: #ffb9ea;
    --meter-progress-color: #ff61d0;
    --meter-text-color: #000000;
    --meter-text-size: 0.8em;
}
```

<br />

## Notes:
Remember to uncheck the following in the OBS Browser Source
- Shutdown source when not visible
- Refresh browser when scene becomes active

> ***Do not close this Application, just minimize it!***
<br />
---
<br />
<br />
Written by [@MiraSynth](https://www.twitter.com/mirasynth), for [Yuuki](https://www.twitter.com/bananyuuki) with love! <3