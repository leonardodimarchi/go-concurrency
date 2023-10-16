# Golang Concurrency

## What i've used

### Goroutines

Goroutines allow us to perform concurrent operations without the need for explicit thread management. Goroutines are created using the `go` keyword, followed by a function call.

### Channels

Channels are a way to communicate and synchronize data between goroutines. 

### Wait Groups

Wait Groups are a synchronization mechanism in Go that allow the main goroutine to wait for a group of goroutines to complete their execution before proceeding. They are particularly useful when you have multiple goroutines performing independent tasks and you want to wait for all of them to finish before continuing.

## Running it

To run the program you will need a `.env` file containing a api key from [open weather](https://openweathermap.org/api).

```
API_KEY=<KEY>
```

Then, you can use `go run *.go` to compile and run the program

## Results

```
------Sync------
Data for Sorocaba {Sorocaba Clouds {21.16}}
Data for São Paulo {São Paulo Clouds {18.68}}
Data for Paraná {Paraná Clouds {22.68}}
Data for Itu {Itu Clear {21.23}}

Took: 1.767177273s

------Async------
Data for Sorocaba {Sorocaba Clouds {21.16}}
Data for Paraná {Paraná Clouds {22.68}}
Data for Itu {Itu Clear {21.23}}
Data for São Paulo {São Paulo Clouds {18.68}}

Took: 888.033896ms
```
