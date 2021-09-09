Build
---

To build, simply run the following. You will need Go installed to build the executable.

    $ make build

Getting started
---

To see the help and options, run the following command

    $ ./shortest-path --help

For convinence, a random flag is included that creates a sample graph and find the distance between two random vertices as following

```sh
./shortest-path --random
Adjacency Distance Matrix
---
    	Node1	Node2	Node3	Node4	Node5	Node6
Node1|	    0	   96	   31	   64	   96	   60
Node2|	   61	    0	   69	   69	   28	    9
Node3|	   67	   12	    0	   72	   34	   48
Node4|	   78	   48	   82	    0	   31	   59
Node5|	   64	   97	   47	   11	    0	   75
Node6|	   17	   67	   63	   68	   66	    0
Fastest route between Node3 to Node2
---
Node3 --(12)--> Node2
```

Input Format
---

The input format is as following 

```
SIZE
SOURCE
VERTEX
ROW 1
ROW 2
.
.
.
ROW N
```

The `SIZE` is the size of the graph or the number of the vertices. SIZE > 0

The `SOURCE` is the source node. (0 < SOURCE <= SIZE)

The `DESTINATION` is the destination node. (0 < DESTINATION <= SIZE)

`Rows 1 through N` are the Adjacancy Distance Matrix where each row represents the distance to from that node to all the other nodes. Distance between the node and itself must be equal to 0. (N == SIZE)

For an exaample of the sample code see the [samplefile.txt](./samplefile.txt)

Read from File
---

If you have a file that matches the expected input format, run

    $ ./shortest-path --file /path/to/file

Read from Stdin
---

You also have the option to pass the input to the program through Stdin. For instance

```sh
$ cat <<EOF  | ./shortest-path
6
2
1
0   2  3  4  5  6
7   0  9 10 11 12
13 14  0 16 17 18
19 20 21  0 23 24
25 26 27 28  0 30
31 32 33 34 35  0
EOF
```

