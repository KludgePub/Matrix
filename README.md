# Workshop Matrix Demoscene
___

### Goal is to achieve something like that

![Rainfall of symbols like in the Matrix](./assets/example.gif "matrix_example")

## Implementation steps

### Construction of matrix with: `factory/matrix.go`
File responsible to create the structure of the matrix, to create a letter object you will need to think about making it like a table.

1. Use function `createNewLetterActor` it will be required to give window/screen configuration and position coordinates where initially it will be placed in the screen.
2. To fill symbols like in the movie "The Matrix" take a look at function `setSymbol` it will randomly assign a symbol to your actor and that will be displayed on the screen.
3. Return slice of generated symbols/actors `[]entity.SceneObject` from `func CreateMatrix(...)` it will be called in `main.go` for you.

### Make effect of rain with: `actor/letter.go`
This file contains basic implementation on how it will behave during screen updates.
Basically in method `OnUpdate`, you will try to reproduce effects from the movie "a rainfall of symbols".
With structure `MatrixLetter` you have access to symbol `MatrixLetter.Text.Position` also you could access to the color `MatrixLetter.color`.
So you could make movements of letters and color transitions during the update of the screen - each frame `OnUpdate` is called.

If simplify it, implement logic in `OnUpdate` method:
1. Symbol movement by updating `MatrixLetter.Text.Position`.
2. Symbol color transition by changing `MatrixLetter.color`.

### Build

Before building by simple `go build` you need have `SDL2` library with needed dependencies.

Debian and friends: 
* ```apt install libsdl2{,-image,-gfx}-dev```

Darwin:
* ```brew install sdl2{,_image,_gfx} pkg-config```

___
> Tested only on Debian 10 x64

