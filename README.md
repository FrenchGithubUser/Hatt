<h1 align="center"> Hatt </h1>
<p align="center"> Stop searching, start finding </p>

This tool brings a graphical interface to search for files through multiple websites. Similarily to what [Jackett](https://github.com/Jackett/Jackett) does, but for DDL (direct download).

## Goals

- [ ] List files together with their informations (size, thumbnail preview, length (pages/duration etc), publication date, etc)

- [ ] Sort files in the list

- [ ] Add categories/subcategories for each website and only search for files in appropriate categories/subcategories

- [ ] Handle multiple search links if one website provides multiple categories/subcategories, with only 1 config file (no useless redundancy)

- [ ] Create a score system to rank sources (in case of multiple sources for the same file) (based on dl speed, accuracy etc.)

- [ ] Support multiple languages

- [ ] Support websites which require authentication to search (usually forums)

- [ ] Allow the user to create custom lists of websites for custom categories (only search on those websites if the custom category is selected)

## Not Goals (to this day)

- All-in-one tool to download/stream media and display it nicely at the same time. Many softwares already do that very well ([Kodi](https://github.com/xbmc/xbmc) for movies/TV shows, [Pegasus-fe](https://github.com/mmatyas/pegasus-frontend) for games and programs, just to name a few)

- Easy "download" option. Some programs already allow to download files very well ([JDownloader](https://jdownloader.org/), [Youtube-dl](https://github.com/ytdl-org/youtube-dl), [Lux](https://github.com/iawia002/lux), just to name a few). There might be an implementation of such a feature by adding those programs as dependencies to Hatt later.


### Categories

ebooks, audio books, newspapers/magazines, courses, movies, tv shows, anime, music, computer software, android apks, console games, pc games, pictures

<details> <summary> <b> Supported sources </b> </summary>

 * batflixmovies
 * bilibili
 * f2movies
 * gomovies
 * himovies
 * kupdf
 * magazinerack
 * nsw2u
 * openloadmov
 * pdfdrive
 * rarefilmm
 * romulation
 * wawacity

</details>
