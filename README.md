
# Iron Birdy

<img src="https://media.giphy.com/media/xTka02cClo5HtVqpC8/giphy.gif"/>


**WARNING THIS IS A WORK IN PROGRESS AND SHOULD NOT BE USED!!**

This project's api is evolving and may not work as described below. This repo is a location for the code in it's current state.


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