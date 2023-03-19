# Contributing to Hatt

Thanks for taking some time to contribute to Hatt, this is what makes open source projects alive !

Feel free to open an issue or a github disucssion if you have any question.


## Local setup

### Dependencies 

The following dependencies are needed to run the program :

- Go (1.18+)
- NPM (Node 15+)
- [Wails](https://wails.io/docs/gettingstarted/installation)

The other dependencies (npm and go libraries) will be installed in the first run

### Start developpement server

```
wails dev
```

### Build Hatt

```
wails build
```

More informations [here](https://wails.io/docs/reference/cli#build)


## Adding sources

Sources can be scraped with 2 different methods :

- By using the `ScrapePlainHtml` function, which parses the html response to a request
- By creating a specific scraper. This method is used only if the `ScrapePlainHtml` function can't be used for the source, or if there is an API. This is to prevent code redundancy.