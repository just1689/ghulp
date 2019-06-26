# Ghulp - Github clone helper

The goal of this command line utility is to make cloning a tiny bit faster by:
- Listing all the projects for a user.
- Getting the clone url and cloning the repo for you.

I often found I had an idea of the repository I wanted to clone but could not remember the clone url. This makes my life a little easier. Maybe it will make yours easier too. Issues / PRs welcome!

## Usage
Ensure you have at least Go 1.12
```bash
go get github.com/just1689/ghulp
# List the repositories for just1689
ghulp just1689
# You will be presented with a list of repos along with a number. Provide the number and hit enter.
1\n

```