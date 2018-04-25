# pre 

Pre is a simple command line tool for adding text to the beginning of a file. 

You previously could do this ["easily"](https://media.giphy.com/media/ANbD1CCdA3iI8/giphy.gif) with `sed`:

```bash
$ sed -i.old '1s;^;to be prepended;' inFile
```

But now you can just do it with `pre`:

```bash 
$ echo "some text in front" | pre someFile.txt | sponge someFile.txt
```

If you don't have sponge, you can get it from `moreutils`:

```
# mac
brew install moreutils

# ubuntu/linux
apt-get install moreutils
```

## Install 

If you have go, you can install this with:

```
go get -u github.com/Flaque/pre
```
