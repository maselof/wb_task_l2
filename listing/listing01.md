Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
[77 78 79]
```

В программе создается массив `a` из пяти элементов, затем по этому массиву делается срез `b` [1:4], его емкость будет равна `4`, а длина `3 = 4 - 1`, в результате в срезе окажутся элементы `[77 78 79]`. Срез `b` ссылается на массив `a`, если изменить его элемент, эти изменения будут видны в срезе `b`.