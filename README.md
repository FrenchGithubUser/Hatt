# <img width="50px" style="margin-bottom:-12px;" src="./frontend/public/images/hatt-logo.png" alt="Hatt"></img> Hatt

This tool allows to search for files through multiple websites, with an intuitive user interface. Similarly to what [Jackett](https://github.com/Jackett/Jackett) does, but for DDL (direct download) and streaming.

<p align="center">
  <img src="https://img.shields.io/github/downloads/FrenchGithubUser/Hatt/total" alt="Downloads"/>
  <img src="https://img.shields.io/badge/Coded%20By%20Human-100%25-brightgreen"/>
</p>
<p align="center">
  <a href="https://discord.gg/88NbZrwmZk" target="_blank"><img height="30px" alt="Discord" src="https://img.shields.io/discord/1088442023582904390?label=Discord&logo=discord"></a>
  <a href="https://www.reddit.com/r/Hatt/" target="_blank"><img height="30px" alt="Subreddit subscribers" src="https://img.shields.io/reddit/subreddit-subscribers/hatt?label=Reddit&style=social"></a>
</p>

![Hatt - Home](.meta/home.jpg)

If you have suggestions (new features, new sources etc.), find a bug or want to notify about something, do not hesitate to open an issue.

If a source is broken, down or should not be trusted anymore, please open an issue about it !

<details> <summary> <b> Supported sources </b> </summary>

 * 9anime
 * androeed
 * animekaizoku
 * animepahe
 * animetosho
 * apkmb
 * audiobb
 * audiobookbay
 * audiobooksbee
 * audiobookslab
 * audiobookss
 * bigaudiobooks
 * bilibili
 * comicextra
 * coomer
 * ddlbase
 * diakov
 * dodi-repacks
 * dosgamesarchive
 * ebook-hunter
 * edgeemu
 * emugames
 * f2movies
 * fapachi
 * fapello
 * fapeza
 * filecr
 * fitgirl-repacks
 * forcoder
 * free-mp3-download
 * g4u
 * galaxyaudiobook
 * game-2u
 * gamedrive
 * getcomics
 * gload
 * gog-games
 * gogoanime
 * goldenaudiobook
 * gomovies
 * hdaudiobooks
 * himovies
 * hotaudiobooks
 * hotleak
 * kayoanime
 * kemono
 * kupdf
 * libgenli
 * library genesis (.rs)
 * lrepack
 * magazinerack
 * memoryoftheworld
 * mobilism
 * nesgm
 * nsw2u
 * online-courses
 * openloadmov
 * ovagames
 * pdfdrive
 * rarefilmm
 * readcomicsonline
 * repackme
 * revdl
 * romulation
 * sflix
 * slavart
 * softarchive
 * steamrip
 * tokybook
 * trantor
 * udemy24
 * uhdmovies
 * vimm
 * watchcartoonsonline
 * wawacity
 * xoxocomics
 * yourserie
 * youtube
 * zoro

</details>

## Features

- Parallel searching on many websites, all at the same time
- Custom website lists to search on
- Search on websites which require an account
- Quickly filter results with a double search and different sorting methods
- Big and ever growing list of supported websites
- Multiple languages support
- Clean and intuitive interface
- Dark mode
- Almost no js rendering, mostly static parsing for high performances

## Installation

Here are the different ways you can get Hatt :

### Linux

- From the [releases tab](https://github.com/FrenchGithubUser/Hatt/releases)

- AUR : [hatt-bin](https://aur.archlinux.org/packages/hatt-bin)


### Windows

- From the [releases tab](https://github.com/FrenchGithubUser/Hatt/releases)


### MacOS

- From the [releases tab](https://github.com/FrenchGithubUser/Hatt/releases) (works on M1 chip). Not always available for the latest release as it requires apple hardware, which I don't personally have.

- Build from source

- Install linux


## Build from source

```
wails build
```

See CONTRIBUTING.md for more details

## Not Goals (to this day)

- All-in-one tool to download/stream media and display it nicely at the same time. Many softwares already do that very well ([Kodi](https://github.com/xbmc/xbmc) for movies/TV shows, [Pegasus-fe](https://github.com/mmatyas/pegasus-frontend) for games and programs, just to name a few)

- Easy "download" option. Some programs already allow to download files very well ([JDownloader](https://jdownloader.org/), [Youtube-dl](https://github.com/ytdl-org/youtube-dl), [Lux](https://github.com/iawia002/lux), just to name a few). There might be an implementation of such a feature by adding those programs as dependencies to Hatt later.


## FAQ

<details> <summary> Why not making a website instead of a program ? </summary>

This would allow more accessibility and less trust needed from the users, however :

- It would require a server to do the scraping, as it is not possible to do it directly in the browser because of the CORS policy that most modern browsers have. This would mean extra costs, which I don't want to bother with, and could impact the project.

- The server's IP address could easily be blocked by most of the sources. Another solution would be to maintain a local database, refreshed every x days, but this requires a lot of extra work (to maintain it, and to create a full-content scraper for every source)

- Having a tool running on everyone's computer allows for more decentralization

- I don't want to be held responsible of running such a service. This repository only provides some code, that you are free to run or not. The user takes the responsibility of what is done with it, not me.

</details>

<details> <summary> Is there a mobile version ? </summary>

Hatt is built on top of the wails framework. As soon as wails support android/iOS builds, Hatt will be available for those platforms.

</details>

## Support

If you like Hatt, you can donate to support its development here :
<details> <summary> <b> Cryptocurrency </b> </summary>

Monero : 46NLLW7dzu5jo2eZ3SiAKgQuVL1Jw8wPMSBAYA3eh4h334HzwMNFSXQ3V3PmXYEoMFXkt24pTHcD1X57KRePN8ehQXn3Ggt

Bitcoin : bc1qf4a44ae76txhmfxfsa875u8mv6murdwawt7msx

Ethereum : 0x134a0974f2fefF0F76276fBdD44439717B2b587B

</details>

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/W7W7JUGNI)

## Disclaimer

All the sources linked in this tool are not intended to support copyright infringement. I am not responsible for and in no way associated to any external links or their content linked, all the links available through this tool are publicly available over the internet. I have no control over the nature, content and availability of other websites. If you dislike the information this tool provides then please contact the corresponding website's owner/webmaster/hoster directly and fill a DMCA takedown request.
