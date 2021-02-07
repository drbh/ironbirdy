
# Iron Birdy

### Single Line Install
```
curl -sLJO https://github.com/drbh/ironbirdy/releases/download/origin/ironbirdy --output ironbirdy && chmod +x ironbirdy && ./ironbirdy
```

<img src="https://media.giphy.com/media/xTka02cClo5HtVqpC8/giphy.gif"/>


**WARNING THIS IS A WORK IN PROGRESS AND SHOULD NOT BE USED!!**

This project's api is evolving and may not work as described below. This repo is a location for the code in it's current state.



### SETTING UP YOUR CONFIG FILE

```
# make a new folder
mkdir ~/.robintools

# make a new file
touch ~/.robintools/config.yml 

# open and edit the file
open -a TextEdit ~/.robintools/config.yml
```

Now copy this into the file and update the values to your user/pass

```
account:
  email: "MYCOOLEMAIL"
  password: "MYSUPERSECRETPASSWORD"
```




### Help Menu

```bash
./ironbirdy help  
```
```
Usage:
  ironbirdy [command]

Available Commands:
  bearcall    Generate all bear call spreads from the current market prices
  bullput     Generate all bull put spreads from the current market prices
  condors     Generate iron condors
  help        Help about any command
  histvol     Use historical daily pricing to find max postive and negative deltas intra week

Flags:
  -h, --help   help for ironbirdy

Use "ironbirdy [command] --help" for more information about a command.
```


### Bear Call Spreads

```bash
./ironbirdy bearcall -t IWM -x 2021-02-12 -o bearcalls.csv 
```
```
Usage:
  ironbirdy bearcall [flags]

Flags:
  -c, --config string   path to config file (default "/Users/drbh/.robintools/config.yml")
  -x, --expire string   expiration date (default "2021-02-05")
  -h, --help            help for bearcall
  -o, --output string   write csv [filename]
  -t, --ticker string   underlying symbol (default "IWM")
```


### Bull Put Spread

```bash
./ironbirdy bullput -t IWM -x 2021-02-12 -o bullputs.csv 
```
```
Usage:
  ironbirdy bullput [flags]

Flags:
  -c, --config string   path to config file (default "/Users/drbh/.robintools/config.yml")
  -x, --expire string   expiration date (default "2021-02-05")
  -h, --help            help for bullput
  -o, --output string   write csv [filename]
  -t, --ticker string   underlying symbol (default "IWM")
```


### Iron Condors

```bash
./ironbirdy condors -t IWM -x 2021-02-12 -o condors.csv 
```
```
Usage:
  ironbirdy condors [flags]

Flags:
  -c, --config string    path to config file (default "/Users/drbh/.robintools/config.yml")
  -d, --distance float   min distance from price (default 0.05)
  -e, --end string       end date (default "2021-02-04")
  -x, --expire string    expiration date (default "2021-02-05")
  -h, --help             help for condors
  -o, --output string    write csv [filename]
  -s, --start string     start date (default "1800-01-01")
  -t, --ticker string    underlying symbol (default "IWM")
```

##  Development

```bash
go build -o ironbirdy cli/cli.go
```
