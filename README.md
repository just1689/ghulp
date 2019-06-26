# Ghulp - Github clone helper

<img src="https://goreportcard.com/badge/github.com/just1689/ghulp">&nbsp;<a href="https://codebeat.co/projects/github-com-just1689-ghulp-master"><img alt="codebeat badge" src="https://codebeat.co/badges/fdacaa8f-93ec-4138-a810-b5521dd51a2c" /></a>
&nbsp;<a href="https://codeclimate.com/github/just1689/ghulp/maintainability"><img src="https://api.codeclimate.com/v1/badges/4ccbe11fba6a8037fa76/maintainability" /></a>&nbsp;
<br />


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