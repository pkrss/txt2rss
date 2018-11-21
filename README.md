# txt2rss
txt to rss software, it can start one rss server.

*** useage ***

put your multiple content to html/txt/{fileName}.txt, txt template your can see html/txt/1.txt,  

then access your url:

http://127.0.0.1:28081/rss/?file={fileName}

if your want random access rss item:

http://127.0.0.1:28081/rss/?random=1&file={fileName}

*** startup ***

```
$ ./txt2rss
```
or custom bind address:
```
$ ./txt2rss -addr=127.0.0.1:28081
```

*** advance **

usage this tool with pkrss: https://sourceforge.net/projects/pkrss/
add http://127.0.0.1:28081/rss/?file=1 to your custom rss source, then listen and see lyric in windows.

