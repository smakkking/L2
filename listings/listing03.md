# listing03

Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```

Ответ:

```
Вывод:
<nil>
false

В первом случае мы выводим nil указатель.
Вывод содержит false так как функция Foo возращает nil для *os.PathError,
а в main сравнивается уже непосредственно с nil

Интерфейс в Go является контрактом. Пустой же интерфейс не содержит методов, и согласно этому – под пустой интерфейс подходит любой тип.
```