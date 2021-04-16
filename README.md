Comet Backup
====

I solved the extensions from 1 to 5. I had a look at others and came across some ideas but no time to implement them. 

NOTE: I found there are somewhere can be refactored to be more structural and extendable, but I paid more attention on solving the problems directly rather than designing good interfaces.    

# Usage

### Build

```shell script
make build
```

The build output is a binary executable file located under directory target.

### Test

```shell script
make test
```

### Run

Show help information:

```shell script
./target/traffic-controller -h
```

Passing arguments from command line:

```shell script
./target/traffic-controller --n 10 --e 20 --s 4 --w 1 
```

Loading data from CSV files:

```shell script
./target/traffic-controller --input ./data/input.csv
```

### Clean

```shell script
make clean
```



