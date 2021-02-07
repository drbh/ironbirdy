
# Iron Birdy

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




### Bear Call Spreads

```
./ironbirdy bearcall -t IWM -x 2021-02-12 -o bearcalls.csv 
```

### Bull Put Spread

```
./ironbirdy bullput -t IWM -x 2021-02-12 -o bullputs.csv 
```

### Iron Condors

```
./ironbirdy condors -t IWM -x 2021-02-12 -o condors.csv 
```

##  Development

```
go build -o ironbirdy cli/cli.go
```