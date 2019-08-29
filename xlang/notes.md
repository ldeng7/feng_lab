### callees

#### go
```
go build -buildmode=c-shared -o ../lib/libxlanggo.so && rm ../lib/libxlanggo.h
```

### callers

#### go
```
go build && LD_LIBRARY_PATH=../../callees/lib ./go
```

#### lua
```
LD_LIBRARY_PATH=../../callees/lib luajit main.lua
```

#### python
```
LD_LIBRARY_PATH=../../callees/lib python3 main.py
```

#### ruby
```
LD_LIBRARY_PATH=../../callees/lib ruby main.rb
```
