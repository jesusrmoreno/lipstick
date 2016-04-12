# Lipstick
:lipstick: A simple app to make your git commit messages more expressive.  
Instead of remembering complicated and unrelated emoji names you can write
things like:
```
:tests Added testing
```
and it'll get turned into:
```
:white_check_mark: Added testing
```
Which will show up as
:white_check_mark: Added testing


# Install
```bash
go get github.com/jesusrmoreno/lipstick
```

After go getting the package cd into a git repo and run
```bash
lipstick install
```
This will add the git commit message hook to turn your keywords into github
emoji.
