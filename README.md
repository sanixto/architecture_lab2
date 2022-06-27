# Software architecture lab4
## Evenloop

### Task

In Go, implement a command interpreter that executes instructions in an event loop.

The interpreter must read commands from the input file. 

Each command occupies one line and has a format

{instruction} {argument} ...

### Commands:

`print {argument}`

The instruction must result in the output of the specified argument on the screen.

Example of a call in the input file:
 
`print Hello!`

Expected result:
 
`mmmmm`

`printc <count> <symbol>`

The instruction should generate a string consisting of symbols of length count.

Example of a call in the input file:
 
`printc 5 m`
 
Expected result:
 
`mmmmm`

### Instructions
  
Install:
```
git clone https://github.com/sanixto/architecture_lab4.git
```
Run:
```
go run main.go
```
Test:
```
go test eventloop/engine
```  

Completed:
- Hakavyi Oleksandr [@sanix_to](https://t.me/sanix_to)
- Lukianenko Mikhail [@lukianenko78](https://t.me/lukianenko78)
