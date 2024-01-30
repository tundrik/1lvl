1. Какой самый эффективный способ конкатенации строк?

Ответ:

```
strings.Builder или bytes.Buffer если нужна манипуляция байтами.
```

2. Что такое интерфейсы, как они применяются в Go?
   Ответ:

```
...
```

3. Чем отличаются RWMutex от Mutex?
   Ответ:

```
...
```

4. Чем отличаются буферизированные и не буферизированные каналы?
   Ответ:

```
...
```

5. Какой размер у структуры struct{}{}?
   Ответ:

```
...
```

6. Есть ли в Go перегрузка методов или операторов?
   Ответ:

```
...
```

7. В какой последовательности будут выведены элементы map[int]int?
   Пример:

```go
m[0]=1
m[1]=124
m[2]=281
```

Ответ:

```
...
```

8. В чем разница make и new?
   Ответ:

```
...
```

9. Сколько существует способов задать переменную типа slice или map?
   Ответ:

```
...
```

10. Что выведет данная программа и почему?

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```

Ответ:

```
...
```

11. Что выведет данная программа и почему?

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

Ответ:

```
...
```

12. Что выведет данная программа и почему?

```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```

Ответ:

```
...
```

13. Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```

Ответ:

```
...
```

14. Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

```

Ответ:

```
...
```
