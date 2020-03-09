# check-package

Toy project to make a command line utility written in Go to check the status of a package you are waiting for. 

Currently USPS is supported

### Installation

```
go get github.com/claytonblythe/check-package
```

It should be installed at $GOPATH/bin and be in your $PATH for you to run with autocomplete
```
ls $GOPATH/bin/check*
```

### Usage
Here is an example of using it to retrieve the status of a USPS package
```
check-package 9405510202061045238692
```

### Output
```
Package 9405510202061045238692 In-Transit
Expected Delivery by Monday March 9 at 8:00pm
LAST UPDATE:
March 8, 2020
 at 12:21 amArrived at USPS Regional Origin FacilityNEWARK NJ DISTRIBUTION CENTER
```