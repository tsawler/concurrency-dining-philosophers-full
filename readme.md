# Concurrency in Go: The Dining Philosophers problem

The Dining Philosophers problem is well known in computer science 
circles. Five philosophers, numbered from 0 through 4, live in a 
house where the table is laid for them; each philosopher has their 
own place at the table. Their only difficulty – besides those of 
philosophy – is that the dish served is a very difficult kind of 
spaghetti, that has to be eaten with two forks. There are two f
orks next to each plate, so that presents no difficulty: 
as a consequence, however, no two neighbours may be eating simultaneously,
since there are five philosophers and five forks.

This is a simple implementation of Dijkstra's solution to the 
"Dining Philosophers" problem. 