# <img width="50px" style="margin-bottom:-12px;" src="./frontend/public/images/hatt-logo.png" alt="Hatt"></img> Hatt

This tool brings a graphical interface to search for files through multiple websites. Similarily to what [Jackett](https://github.com/Jackett/Jackett) does, but for DDL (direct download).

![Hatt - Home](.meta/home.jpg)

## Goals

- [ ] List files together with their informations (size, thumbnail preview, length (pages/duration etc), publication date, etc)

- [ ] Sort files in the list

- [ ] Add categories/subcategories for each website and only search for files in appropriate categories/subcategories

- [ ] Handle multiple search links if one website provides multiple categories/subcategories, with only 1 config file (no useless redundancy)

- [ ] Create a score system to rank sources (in case of multiple sources for the same file) (based on dl speed, accuracy etc.)

- [ ] List the tools that are known to work to download files for each website

- [ ] Display more information on mouse hover (summary of the item etc)

- [ ] Multiple display settings (main color, thumbnail sizes etc.)

- [ ] Support multiple languages

- [ ] Support websites which require authentication to search (usually forums)

- [ ] Allow the user to create custom lists of websites for custom categories (only search on those websites if the custom category is selected)

## Not Goals (to this day)

- All-in-one tool to download/stream media and display it nicely at the same time. Many softwares already do that very well ([Kodi](https://github.com/xbmc/xbmc) for movies/TV shows, [Pegasus-fe](https://github.com/mmatyas/pegasus-frontend) for games and programs, just to name a few)

- Easy "download" option. Some programs already allow to download files very well ([JDownloader](https://jdownloader.org/), [Youtube-dl](https://github.com/ytdl-org/youtube-dl), [Lux](https://github.com/iawia002/lux), just to name a few). There might be an implementation of such a feature by adding those programs as dependencies to Hatt later.


## Installation

- From the [releases tab](https://github.com/FrenchGithubUser/Hatt/releases)


- From your OS' specific repository

AUR : Coming soon


- Build from source

```
wails build
```

See CONTRIBUTING.md for more details


<details> <summary> <b> Supported sources </b> </summary>

 * androeed
 * apkmb
 * batflixmovies
 * bilibili
 * edgeemu
 * emugames
 * f2movies
 * free-mp3-download
 * gamedrive
 * getintopc
 * gload
 * gog-games
 * gomovies
 * himovies
 * kupdf
 * library genesis (.rs)
 * magazinerack
 * nesgm
 * nsw2u
 * online-courses
 * openloadmov
 * pdfdrive
 * rarefilmm
 * revdl
 * romulation
 * sflix
 * steamrip
 * tokybook
 * vimm
 * wawacity
 * youtube

</details>
