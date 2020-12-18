# fishcraft
[![Build Status](https://travis-ci.com/80-am/fishcraft.svg?branch=master&status=started)](https://travis-ci.com/80-am/fishcraft)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/80-am/fishcraft)](https://golang.org/)
[![License: GPL-3.0](https://img.shields.io/github/license/80-am/fishcraft)](https://opensource.org/licenses/GPL-3.0)

World of Warcraft fishing bot 🎣

<img src="./demo.gif" height="200">

Fishcraft uses pixel recognition and detects when fish is on the hook by checking for movements in the bobber.

## Getting Started

These instructions will get you up and running on your local machine.

Copy [config.yml.sample](config.yml.sample) into config.yml and fill in your desired keybind as string.

```yml
key: "0"
debug: false
```

Start the application and focus your WoW window with the character standing at a fishing spot fully zoomed in (FPV).

For now the World of Warcraft window will need to be focused at all time.

## Pipeline

Things to be added in the future:
- [ ] Have the progam run "unfocused".
- [ ] Add option to sell using a mount / repair bot.
- [ ] Log loot.
- [ ] Attach lure.

## Disclaimer

This bot might interfere with Blizzards policy against Third-Party software.
Using it might suspend your account.

You assume all responsibility and liability.