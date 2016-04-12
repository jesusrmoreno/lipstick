# Lipstick
:lipstick: A simple app to make your git commit messages more expressive.
Instead of remembering complicated and unrelated emoji names you can write
things like:
```
:tests: Added testing
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

# Setup
By default lipstick uses the following mappings
```toml
[commitKinds]
format = ":art:"
performance = ":racehorse:"
docs = ":books:"
bugfix = ":bug:"
crucial = ":ambulance:"
remove = ":fire:"
tests = ":white_check_mark:"
security = ":lock:"
ui = ":lipstick:"
wip = ":construction:"
tags = ":bookmark:"
initial = ":tada:"
logging = ":speaker:"
removeLogging = ":mute:"
feature = ":sparkles:"
configuration = ":snowflake:"
```

To override these you can create a .lipstickrc file (this file must follow toml
  syntax) in the same directory as your git folder. For instance the file used
  in this config is as follows:
```toml
[commitKinds]
format = ":art:"
performance = ":racehorse:"
docs = ":books:"
bugfix = ":bug:"
crucial = ":ambulance:"
remove = ":fire:"
tests = ":white_check_mark:"
security = ":lock:"
ui = ":lipstick:"
wip = ":construction:"
tags = ":bookmark:"
init = ":tada:"
logging = ":speaker:"
removeLogging = ":mute:"
feature = ":sparkles:"
configuration = ":snowflake:"
license = ":copyright:"
release = ":gem:"
vendor = ":package:"
chore = ":information_source:"
```


# Uninstall
To remove the hook simply run:
```bash
lipstick uninstall
```
