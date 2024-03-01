# todo_list
Full stack todo list using go, htmx and pico css

All executable files are placed in the bin folder by default


```makefile
make dev
```
builds and runs the application is dev mode, allowing you to make changes to views and static assets without restarting the application, the dev executable is suffixed with _dev

```makefile
make build
```
builds the application in embedded mode, embedding web assets in the executable - producing a stand alone executable with no suffix

```makefile
make clean
```
deletes the bin folder and it's contents

config for the bind address/port can be found under the config folder.  The application looks for this folder in the current working directory, loading a default of localhost:3000 if it can't find it.