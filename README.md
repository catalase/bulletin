Bulletin
========

![Build Status](https://travis-ci.org/catalase/bulletin.svg?branch=master)

# What is it?
Bulletin package crawls news bulletin on [Yonhap News](http://www.yonhapnews.co.kr/).

# Installment
```
$ go get https://github.com/catalase/bulletin
```

# Usage

```go
func Academy() {
    // Get news bulletin from Yonhap news.
    briefs, _ := Briefs()
    
    for _, brief := range briefs {
        // url of the news
        brief.URL()

        // title of the news
        brief.Title()

        // the time when the news is sent.
        brief.Time()
    }
}
```

# Warning
`Brief.Time` returnes time with Korea timezone ([KST][KST], +9). Be careful to deal with time returned by `Brief.Time` if you are not Korean!

# News
- [North Korea's new time zone to break from 'imperialism'][http://www.bbc.com/news/world-asia-33815049]
- [North Korea welcomes new time zone to break from 'imperialism'][http://www.bbc.com/news/world-asia-33942111]

North Korea has introduced new timezone to break from imperialism of past Japan (In past, Korea was ruled out by Japan for 1910-1945). Disregard it for they has been crazy. the Republic of Korea (Korea) don't use North Korea's timezone.

Anyway, is there newspaper in North Korea? I don't know :)

[KST]: https://en.wikipedia.org/wiki/Korea_Standard_Time