# Simulated Evolution

Adapted from A.K. Dewdney's article "Simulated evolution: wherein bugs learn to hunt bacteria" in the "Computer Recreations" column of Scientific American (May 1989: 138-141). I gave it a little twist and replaced bugs and bacteria with animals and trees. The motion of the animals is still a bit random, not exactly taking the genes into consideration. My objective was mostly to play with [Pixel](github.com/faiface/pixel), a 2D game library.

The code is still a tiny tad messy and lacking documentation, but hopefully it's easy to understand and modify.

# Dependencies

The project uses Go modules, so you'll want to use a vergion of Go more recent than [1.11](https://blog.golang.org/using-go-modules).

# Building

```
$ make bin
go build -o evolution cmd/simulated-evolution/*.go
```

# The world

When you launch the program, an animal will be spawned at the center of the world. Each animal has 8 genes and a certain amount of energy. Some of the genes will determine their predisposition for agressiveness. When they reproduce (actually this is pretty basic, they essentially clone themselves), there is a gene mutation on the offspring.

There are 6 types of trees in the world, and when eaten they either give or take energy, depending on the type of tree.

# Controls

- Up, Down, Left, Right => controls the camera motion.
- Mouse / TouchPad scroll up/down => Camera zoom in and out

# Limitations and Caveats

- All the values are still hardcoded: chance of reproduction, necessary energy for reproduction, probability of trees growing, etc.
- Path for sprites are hardcoded.
- Animal motion is still random, perhaps it should depend not only on the genes but on the outcome of exploring certain areas and their interaction with other animals.
- The code isn't tested at all

# Contributing

I'll be tackling the limitations and caveats as I find the time, but contributions are more than welcome. Just open a Pull-request.

# License

[MIT](https://github.com/csixteen/simulated-evolution/blob/master/LICENSE), as always.
