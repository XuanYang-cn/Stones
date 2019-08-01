# Resources

**VIM**

###### Search and substitue: substitute `foo` with `bar` 

**%** for searching in all lines, **c** for asking confirmation, **I** for case sensitive

	:%s/foo/bar/gcI

---
**Wonderful blogs**

|Priority| Website | Bref | Author |
|:--------:|---------|------|--------|
|\* \* \* \*|[Personal Blog](https://www.benfrederickson.com/blog/)|Recommendation algo and wonderful projects|Ben Frederickson([@Github](https://github.com/benfred))|
---
**Git**

|Priority| Website | Bref | Author |
|:--------:|---------|------|--------|
|\* \*|[Git conceptions](https://www.atlassian.com/git/tutorials/)|This article Describe git thoroughly|UNKNOW|


---
**ORM**

|Priority| Website | Bref | Author |
|:--------:|---------|------|--------|
|\* \*|[SQLAlchemy](http://aosabook.org/en/sqlalchemy.html)|SQLAlchemy's approach to database abstraction|Michael Bayer(Twitter@zzzeek)|


---
**Twisted**

|Priority| Website | Bref | Author |
|:--------:|---------|------|--------|
|\*|[Twisted introduction](http://krondo.com/in-which-we-begin-at-the-beginning/)|Series Introduction of Twisted|Dave Peticolas([blog address](http://krondo.com/))|
|\*|[Twisted Reconnecting thrift client Imp](https://www.skitoy.com/p/twisted-code-review/261/)|Codes, 3mins|David Koblas([@Github](https://github.com/koblas))|

---
**Asynchronous**

|Priority| Website | Bref | Author |
|:------:|---------|------|--------|
|\* \* \*|[How does asyncio actually work](https://stackoverflow.com/questions/49005651/how-does-asyncio-actually-work/51116910#51116910)|stackoverflow answer|Brad Solomon([@Github](https://github.com/bsolomon1124))|
|\* \* \*|[Async IO complete walkthrough](https://realpython.com/async-io-python/)|Real Python|Brad Solomon([@Github](https://github.com/bsolomon1124))|
|\* \* \*|[Think Outside the GIL with AsyncIO and MP](https://www.youtube.com/watch?v=0kXaLh8Fz3k)|Youtube video|speaker: John Reese|
|\* \* \*|[Python asyncio module](https://docs.python.org/3/library/asyncio.html)|python doc|python|


---
**Thread**

|Priority| Website | Bref | Author |
|:------:|---------|------|--------|
|\* \*|[Inter-thread communication with queue](http://code.activestate.com/recipes/491281/)|python codes|recipes[@Github](https://github.com/ActiveState/code/tree/master/recipes/Python)|


---
**Distribute tracing**


|Priority| Website | Bref | Author |
|:------:|---------|------|--------|
|\* \* \*|[OpenTracing](https://opentracing.io/)|opentracing official website |Ben Sigelman|

---
**Linux Shell**

- Run a script multiple times
```shell
$ for ((i=1; i<=10; i++))
do
    python script.py "$i"
done
```
- Stress Test

[`apache2-utils` ab tests](https://pkgs.org/download/apache2-utils)
