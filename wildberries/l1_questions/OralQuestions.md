# L1 Устные вопросы

## 1. Какой самый эффективный способ конкатенации строк

![alt text](images/benchResults.png)

## 2. Что такое интерфейсы, как они применяются в Go?


Интерфейсный тип в Go — это своего рода определение. Он определяет и описывает конкретные методы, которые должны быть у какого-то другого типа.

Одним из интерфейсных типов из стандартной библиотеки является интерфейс fmt.Stringer:

```go
type Stringer interface {
    String() string
}
```

Мы говорим, что что-то удовлетворяет этому интерфейсу (или реализует этот интерфейс), если у этого «что-то» есть метод с конкретным сигнатурным строковым значением String().

Например, тип Book удовлетворяет интерфейсу, потому что у него есть строковый метод String():

```go
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```

Неважно, каким типом является Book или что он делает. Важно лишь, что у него есть метод под названием String(), который возвращает строковое значение.

А вот другой пример. Тип Count тоже удовлетворяет интерфейсу fmt.Stringer, потому что у него есть метод с тем же сигнатурным строковым значением String().


```go
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```

Здесь важно понять, что у нас есть два разных типа Book и Count, которые действуют по-разному. Но их объединяет то, что они оба удовлетворяют интерфейсу fmt.Stringer.

Можете посмотреть на это с другой стороны. Если вы знаете, что объект удовлетворяет интерфейсу fmt.Stringer, то можете считать, что у него есть метод с сигнатурным строковым значением String(), которое вы можете вызывать.

А теперь самое важное.

Когда вы видите в Go объявление (переменной, параметра функции или поля структуры), имеющее интерфейсный тип, вы можете использовать объект любого типа, пока он удовлетворяет интерфейсу.

Допустим, у нас есть функция:


```go
func WriteLog(s fmt.Stringer) {
    log.Println(s.String())
}
```

Поскольку WriteLog() использует в объявлении параметра интерфейсный тип fmt.Stringer, мы можем передавать любой объект, удовлетворяющий интерфейсу fmt.Stringer. Например, можем передать типы Book и Count, которые создали ранее в методе WriteLog(), и код будет нормально работать.

Кроме того, поскольку передаваемый объект удовлетворяет интерфейсу fmt.Stringer, мы знаем, что у него есть строковый метод String(), который может быть безопасно вызван функцией WriteLog().

Давайте соберём всё сказанное в один пример, демонстрирующий мощь интерфейсов.

```go
package main

import (
    "fmt"
    "strconv"
    "log"
)

// Объявляем тип Book, который удовлетворяет интерфейсу fmt.Stringer.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Объявляем тип Count, который удовлетворяет интерфейсу fmt.Stringer.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Объявляем функцию WriteLog(), которая берёт любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
    log.Println(s.String())
}

func main() {
    // Инициализируем объект Book и передаём в WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Инициализируем объект Count и передаём в WriteLog().
    count := Count(3)
    WriteLog(count)
}
```

Это круто. В основной функции мы создали разные типы Book и Count, но передали их одной функции WriteLog(). А та вызвала соответствующие функции String() и записала результаты в журнал.

Есть целый ряд причин, по которым вы можете начать использовать интерфейсы в Go. И по моему опыту, самые важные из них такие:

1. Интерфейсы помогают уменьшить дублирование, то есть количество шаблонного кода.
2. Они облегчают использование в модульных тестах заглушек вместо реальных объектов.
3. Будучи архитектурным инструментом, интерфейсы помогают отвязывать части вашей кодовой базы.

Рассмотрим подробнее эти способы использования интерфейсов.

**Уменьшение количества шаблонного кода**

Пусть у нас есть структура Customer, содержащая какие-то данные о клиенте. В одной части кода мы хотим записывать эту информацию в bytes.Buffer, а в другой части хотим записывать данные о клиенте в os.File на диске. Но, в обоих случаях, мы хотим сначала сериализовать структуру Сustomer в JSON.

При таком сценарии мы можем с помощью интерфейсов Go уменьшить количество шаблонного кода.

В Go есть интерфейсный тип io.Writer:

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

И мы можем воспользоваться тем, что bytes.Buffer и тип os.File удовлетворяют этому интерфейсу, поскольку имеют, соответственно, методы bytes.Buffer.Write() и os.File.Write().

```go
package main

import (
    "encoding/json"
    "io"
    "log"
    "os"
)

// Создаём тип Customer.
type Customer struct {
    Name string
    Age  int
}

// Реализуем метод WriteJSON, который берёт io.Writer в виде параметра.
// Он отправляет структуру Сustomer в JSON, и если всё отрабатывает 
// успешно, то вызывается соответствующий метод Write() из io.Writer.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    // Инициализируем структуру Customer.
    c := &Customer{Name: "Alice", Age: 21}

    // Затем с помощью Buffer можем вызвать метод WriteJSON
    var buf bytes.Buffer
    err := c.WriteJSON(buf)
    if err != nil {
        log.Fatal(err)
    }

    // или воспользоваться файлом.
    f, err := os.Create("/tmp/customer")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


    err = c.WriteJSON(f)
    if err != nil {
        log.Fatal(err)
    }
}
```

Конечно, это лишь выдуманный пример (мы можем по-разному структурировать код, чтобы добиться того же результата). Но он хорошо иллюстрирует преимущества использования интерфейсов: мы можем один раз создать метод Customer.WriteJSON() и вызывать его каждый раз, когда нужно записать во что-то, удовлетворяющее интерфейсу io.Writer.

Но если вы новичок в Go, у вас возникнет пара вопросов: «Как узнать, что интерфейс io.Writer вообще существует? И как заранее узнать, что ему удовлетворяют bytes.Buffer и os.File?»

Боюсь, простого решения тут нет. Вам просто нужно набраться опыта, познакомиться с интерфейсами и разными типами из стандартной библиотеки. В этом поможет чтение документации по этой библиотеке и просмотр чужого кода. А для быстрого ознакомления я добавил в конец статьи самые полезные типы интерфейсные типы.

Но даже если вы не используете интерфейсы из стандартной библиотеки, ничто не мешает вам создать и использовать собственные интерфейсные типы. Об этом мы поговорим ниже.

**Модульное тестирование и заглушки**

Чтобы понять, как интерфейсы помогают в модульном тестировании, давайте рассмотрим пример посложнее.

Допустим, у вас есть магазин, и вы храните в PostgreSQL информацию о продажах и количестве клиентов. Вы хотите написать код, вычисляющий долю продаж (удельное количество продаж на одного клиента) за последние сутки, округлённую до двух знаков после запятой.

```go
package main

import (
    "fmt"
    "log"
    "time"
    "database/sql"
    _ "github.com/lib/pq"
)

type ShopDB struct {
    *sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func main() {
    db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    shopDB := &ShopDB{db}
    sr, err := calculateSalesRate(shopDB)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf(sr)
}

func calculateSalesRate(sdb *ShopDB) (string, error) {
    since := time.Now().Sub(24 * time.Hour)

    sales, err := sdb.CountSales(since)
    if err != nil {
        return "", err
    }

    customers, err := sdb.CountCustomers(since)
    if err != nil {
        return "", err
    }

    rate := float64(sales) / float64(customers)
    return fmt.Sprintf("%.2f", rate), nil
}
```

Теперь мы хотим создать модульный тест для функции calculateSalesRate(), чтобы проверить корректность вычислений.

Сейчас это проблематично. Нам понадобится настроить тестовый экземпляр PostgreSQL, а также создать и удалить скрипты для наполнения базы фальшивыми данными. Придётся проделать много работы, если мы действительно хотим протестировать наши вычисления.

И на помощь приходят интерфейсы!

Мы создадим собственный интерфейсный тип, описывающий методы CountSales() и CountCustomers(), на которые опирается функция calculateSalesRate(). Затем обновим сигнатуру calculateSalesRate(), чтобы использовать этот интерфейсный тип в качестве параметра вместо прописанного типа *ShopDB.

```go
package main

import (
    "fmt"
    "log"
    "time"
    "database/sql"
    _ "github.com/lib/pq"
)

// Создаём свой интерфейс ShopModel. Он прекрасно подходит для
// интерфейса с описанием нескольких методов, и он должен описывать
// входные параметры-типы, а также типы возвращаемых значений.
type ShopModel interface {
    CountCustomers(time.Time) (int, error)
    CountSales(time.Time) (int, error)
}

// Тип ShopDB удовлетворяет новому интерфейсу ShopModel, потому что
// у него есть два необходимых метода -- CountCustomers() и CountSales().
type ShopDB struct {
    *sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func main() {
    db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    shopDB := &ShopDB{db}

    sr := calculateSalesRate(shopDB)
    fmt.Printf(sr)
}

// Заменим это для использования интерфейсного типа ShopModel в виде параметра
// вместо прописанного типа *ShopDB.
func calculateSalesRate(sm ShopModel) string {
    since := time.Now().Sub(24 * time.Hour)

    sales, err := sm.CountSales(since)
    if err != nil {
        return "", err
    }

    customers, err := sm.CountCustomers(since)
    if err != nil {
        return "", err
    }

    rate := float64(sales) / float64(customers)
    return fmt.Sprintf("%.2f", rate), nil
}
```

После того как мы это сделали, нам будет просто создать заглушку, которая удовлетворяет интерфейсу ShopModel. Затем можно использовать её в ходе модульного тестирования корректной работы математической логики в функции calculateSalesRate(). Вот так:

```go
package main

import (
    "testing"
)

type MockShopDB struct{}

func (m *MockShopDB) CountCustomers() (int, error) {
    return 1000, nil
}

func (m *MockShopDB) CountSales() (int, error) {
    return 333, nil
}

func TestCalculateSalesRate(t *testing.T) {
    // Инициализируем заглушку.
    m := &MockShopDB{}
    // Передаём заглушку в функцию calculateSalesRate().
    sr := calculateSalesRate(m)

    // Проверяем, соответствует ли возвращаемое значение ожиданиям на основе
    // фальшивых входных данных.
    exp := "0.33"
    if sr != exp {
        t.Fatalf("got %v; expected %v", sr, exp)
    }
}
```

**Архитектура приложения**


В предыдущем примере мы видели, как можно использовать интерфейсы для отвязки определённых частей кода от использования конкретных типов. Например, функции calculateSalesRate() совершенно не важно, что вы ей передадите, лишь бы оно удовлетворяло интерфейсу ShopModel.

Вы можете расширить эту идею и создавать в крупных проектах целые «отвязанные» уровни.
Допустим, вы создаёте веб-приложение, взаимодействующее с базой данных. Если сделаете интерфейс, описывающий определённые методы для взаимодействия с БД, то сможете ссылаться на него вместо конкретного типа через HTTP-обработчики. Поскольку HTTP-обработчики ссылаются только на интерфейс, это поможет отвязать друг от друга HTTP-уровень и уровень взаимодействия с базой данных. Будет проще работать с уровнями независимо, и в будущем вы сможете заменять какие-то уровни, не влияя на работу остальных.

Я написал об этом паттерне в одном из предыдущих постов, там больше подробностей и приведены практические примеры.

**Что такое пустой интерфейс?**

Если вы уже какое-то время программируете на Go, то наверняка сталкивались с пустым интерфейсным типом interface{}. Попробую объяснить, что это такое. В начале этой статьи я написал:

Интерфейсный тип в Go — это своего рода определение. Он определяет и описывает конкретные методы, которые должны быть у какого-то другого типа.

Пустой интерфейсный тип не описывает методы. У него нет правил. И поэтому любой объект удовлетворяет пустому интерфейсу.

По сути, пустой интерфейсный тип interface{} — своего рода джокер. Если вы встретили его в объявлении (переменной, параметра функции или поля структуры), то можете использовать объект любого типа.


```go
package main

import "fmt"


func main() {
    person := make(map[string]interface{}, 0)

    person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

    fmt.Printf("%+v", person)
}
```

Когда речь заходит об извлечении и использовании значений из map’ы, важно помнить вот о чём. Допустим, вы хотите получить значение age и увеличить его на 1. Если вы напишете подобный код, то он не скомпилируется:


```go
package main

import "log"

func main() {
    person := make(map[string]interface{}, 0)

    person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

    person["age"] = person["age"] + 1

    fmt.Printf("%+v", person)
}
```

Вы получите сообщение об ошибке:

```invalid operation: person["age"] + 1 (mismatched types interface {} and int)```

Причина в том, что значение, хранящееся в map, принимает тип interface{} и теряет свой исходный, базовый тип int. И поскольку значение больше не целочисленное, мы не можем прибавить к нему 1.

Чтобы это обойти, вам нужно сделать значение снова целочисленным, и только потом его использовать:

```go
package main

import "log"

func main() {
    person := make(map[string]interface{}, 0)

    person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

    age, ok := person["age"].(int)
    if !ok {
        log.Fatal("could not assert value to int")
        return
    }

    person["age"] = age + 1

    log.Printf("%+v", person)
}
```


# 3. Чем отличаюютя RWMutex от Mutex

Mutex означает mutual exclusion(взаимное исключение) и является способом защиты critical section(критическая секция) вашей программы.

![alt text](images/mutex.webp)

Критическая секция — это область вашей программы, которая требует эксклюзивного доступа к общему ресурсу. При нахождении в критической секции двух (или более) потоков возникает состояние race(гонки). Так же возможны проблемы взаимной блокировки(deadlock).

Mutex обеспечивает безопасный доступ к общим ресурсам.

RWMutex концептуально то же самое, что и Mutex: он защищает доступ к памяти. Тем не менее, RWMutex дает вам немного больше контроля над памятью. Вы можете запросить блокировку для чтения, и в этом случае вам будет предоставлен доступ, если блокировка не удерживается для записи.

Это означает, что произвольное число читателей может удерживать блокировку читателя, пока ничто другое не удерживает блокировку писателя.

Посмотрим как это работает:

```go
func (c *counter) CountV1() int {
   c.Lock()
   defer c.Unlock()
   return c.count
}
func (c *counter) CountV2() int {
   c.RLock()
   defer c.RUnlock()
   return c.count
}
```

A little bit of testing benchmark

```go
func BenchmarkCountV1(b *testing.B) {
   c := new(counter)
   var wg sync.WaitGroup
   for i := 0; i < b.N; i++ {
      for j := 0; j < 1000; j++ {
         wg.Add(1)
         go func() {
            defer wg.Done()
            c.CountV1()
         }()
      }
      wg.Wait()
   }
}

func BenchmarkCountV2(b *testing.B) {
   c := new(counter)
   var wg sync.WaitGroup
   for i := 0; i < b.N; i++ {
      for j := 0; j < 1000; j++ {
         wg.Add(1)
         go func() {
            defer wg.Done()
            c.CountV2()
         }()
      }
      wg.Wait()
   }
}
```

Results:

BenchmarkCountV1-8           2132            501896 ns/op
BenchmarkCountV2-8           3358            306254 ns/op


# 4. Каналы

Что такое каналы?

Канал — это объект связи, с помощью которого горутины обмениваются данными. Технически это конвейер (или труба), откуда можно считывать или помещать данные. То есть одна горутина может отправить данные в канал, а другая — считать помещенные в этот канал данные.

Создание канала

Go для создания канала предоставляет ключевое слово chan. Канал может передавать данные только одного типа, данные других типов через это канал передавать невозможно.

```go
package main

import "fmt"

func main() {
    var c chan int
    fmt.Println(c)
} 
```

Программа выше создает канал c, который будет передавать int. Данная программа выведет <nil>, потому что нулевое значение канала — это nil. Такой канал абсолютно бесполезен. Вы не можете передать или получить данные из канала, так как он не был создан (инициализирован). Для его создания необходимо использовать make.

```go
package main

import "fmt"

func main() {
    c := make(chan int)

    fmt.Printf("type of `c` is %T\n", c)
    fmt.Printf("value of `c` is %v\n", c)
}
```

В данном примере используется короткий синтаксис := для создания канала с использованием функции make. Программа выше выводит следующий результат:

```
type of `c` is chan int
value of `c` is 0xc0420160c0
```

Обратите внимание на значение переменной c, это адрес в памяти. В go каналы являются указателями. В большинстве своем, когда вам необходимо взаимодействовать с горутиной, вы помещаете канал как аргумент в функцию или метод. Горутина получает этот канал как аргумент, и вам не нужно разыменовывать его для того, чтобы извлечь или передать данные через этот канал.

Запись и чтение данных

Go предоставляет простой синтаксис для чтения <- и записи в канал

```go
c <- data
```
В этом примере мы передаем данные в канал c. Направление стрелки указывает на то, что мы извлекаем данные из data и помещаем в канал c.

```go
<- c
```
А здесь мы считываем данные с канала c. Эта операция не сохраняет данные в переменную и она является корректной. Если вам необходимо сохранить данные с канала в переменную, вы можете использовать следующий синтаксис:

```go
var data int
data = <- c
```
Теперь данные из канала c, который имеет тип int, могут быть записаны в переменную data. Так же можно упростить запись, используя короткий синтаксис:

```go
data := <- c
```

Go определит тип данных, передаваемый каналу c, и предоставит data корректный тип данных.

Все вышеобозначенные операции с каналом являются блокируемыми. Когда вы помещаете данные в канал, горутина блокируется до тех пор, пока данные не будут считаны другой горутиной из этого канала. В то же время операции канала говорят планировщику о планировании другой горутины, поэтому программа не будет заблокирована полностью. Эти функции весьма полезны, так как отпадает необходимость писать блокировки для взаимодействия горутин.

Каналы на практике

```go
package main

import "fmt"

func greet(c chan string) {
    fmt.Println("Hello " + <-c + "!")
}

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    go greet(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

Разберем программу по шагам:


1. Мы объявили функцию greet, которая принимает канал c как аргумент. В этой функции мы считываем данные из канала c и выводим в консоль.
2. В функции main программа сначала выводит "main() started".
3. Затем мы, используя make, создаем канал c с типом даных string.
4. Помещаем канал с в функцию greet и запускаем функцию как горутину, используя ключевое слово go.
5. Теперь у нас имеется две горутины main и greet, main по-прежнему остается активной.
6. Помещаем данные в канал с и в этот момент main блокируется до тех пор, пока другая горутина (greet) не считает данные из канала c. Планировщик Go планирует запуск greet и выполняет описанное в первом пункте.
7. После чего main снова становится активной и выводит в консоль "main() stopped".

Deadlock (Взаимная блокировка)

Как уже ранее говорилось, чтение или запись данных в канал блокирует горутину и контроль передается свободной горутине. Представим, что такие горутины отсутствуют, либо они все "спят". В этот момент возникает deadlock, который приведет к аварийному завершению программы.


Если вы попытаетесь считать данные из канала, но в канале будут отсутствовать данные, планировщик заблокирует текущую горутину и разблокирует другую в надежде, что какая-либо горутина передаст данные в канал. То же самое произойдет в случае отправки данных: планировщик заблокирует передающую горутину, пока другая не считает данные из канала.
Примером deadlock может быть main горутина, которая эксклюзивно производит операции с каналом.

package main

import "fmt"

func main() {
    fmt.Println("main() started")

    c := make(chan string)
    c <- "John"

    fmt.Println("main() stopped")
}

Программа выше выведет следующее при попытке ее исполнить:

```
main() started
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
        program.go:10 +0xfd
exit status 2
```

Закрытие канала

В Go так же можно закрыть канал, через закрытый канал невозможно будет передать или принять данные. Горутина может проверить закрыт канал или нет, используя следующую конструкцию: val, ok := <- channel, где ok будет истиной в случае, если канал открыт или операция чтения может быть выполнена, иначе ok будет false, если канал закрыт и отсутствуют данных для чтения из него. Закрыть канал можно, используя встроенную функцию close, используя следующий синтаксис close(channel). Давайте рассмотрим следующий пример:

```go
package main

import "fmt"

func greet(c chan string) {
    <-c // for John
    <-c // for Mike
}

func main() {
    fmt.Println("main() started")

    c := make(chan string, 1)

    go greet(c)
    c <- "John"

    close(c) // closing channel

    c <- "Mike"
    fmt.Println("main() stopped")
}
```

Для понимания концепта блокировки первая операция отправки c <- "John" будет блокирующей, и другая горутина должна будет считать данные из канала, следовательно greet горутина будет запланирована планировщиком. Затем первая операция чтения будет неблокируемой, поскольку присутствуют данные для чтения в канале c. Вторая операция чтения будет блокируемой, потому что в канале c отсутствуют данные, поэтому планировщик переключится на main горутину и программа выполнит закрытие канала close(c).

Вывод программы:

```
main() started
panic: send on closed channel

goroutine 1 [running]:
main.main()
    program.go:20 +0x120
exit status 2
```

Как вы можете заметить, программа завершилась с ошибкой, которая говорит, что запись в закрытый канал невозможна. Для дальнейшего понимания закрытия каналов давайте рассмотрим пример с циклом for.

**Размер буфера канала**

Как вы уже заметили, каждая операция отправки данных в канал блокирует текущую горутину. Но мы еще не использовали функцию make с 2-мя аргументами. Второй аргумент — это размер буфера канала. По-умолчанию размер буфера канала равен 0, такой канал называется небуферизированным каналом. То есть все, что мы пишем в канал, сразу доступно для чтения.


Когда размер буфера больше 0, горутина не блокируется до тех пор, пока буфер не будет заполнен. Когда буфер заполнен, любые значения отправляемые через канал, добавляются к буферу, отбрасывая предыдущее значение, которое доступно для чтения (где горутина будет заблокирована). Но есть один подвох, операция чтения на буферизированном канале является жадной, таким образом, как только операция чтения началась, она не будет завершена до полного опустошения буфера. Это означает, что горутина будет считывать буфер канала без блокировки до тех пор, пока буфер не станет пустым.


Для объявления буферизированного канала мы можем использовать следующий синтаксис:

```go
c := make(chan Type, n)
```

Это выражение создаст канал с типом данных Type и размером буфера n. Текущая горутина не будет заблокирована, пока в канал не будет передано n+1 данных.


Давайте докажем, что горутина не блокируется, пока буфер не заполнится и не переполнится:

```go
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)

    go squares(c)

    c <- 1
    c <- 2
    c <- 3

    fmt.Println("main() stopped")
}
```

В этом примере канал c имеет размер буфера равным 3. Это означает, что он может содержать 3 значения(c <- 3), но поскольку буфер не переполняется (так как мы не поместили новое значение в буфер), main не будет блокироваться, и программа будет успешно завершена без вывода чисел. Вывод программы:


main() started
main() stopped


Теперь давайте поместим еще одно значение в канал:

```go
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)

    go squares(c)

    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("main() stopped")
}
```

Как упоминалось ранее, теперь мы помещаем дополнительное значение в буфер и main блокируется, затем стартует горутина squares, которая вычитывает все значения из буфера, пока он не станет пустым.

**Длина и емкость канала**

Подобно срезам, буферизированный канал имеет длину и емкость. Длина канала — это количество значений в очереди (не считанных) в буфере канала, емкость — это размер самого буфера канала. Для того, чтобы вычислить длину, мы используем функцию len, а, используя функцию cap, получаем размер буфера.

```go
package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
    fmt.Println()
}
```

Вывод программы:


Length of channel c is 2 and capacity of channel c is 3

Вышеприведенная программа работает нормально и deadlock не возникает, потому что размер буфера канала равен 3, а мы записали только 2 значения в буфер, поэтому планировщик не попытался запланировать другую горутину и не заблокировал main. Вы даже можете считать эти данные в main, если вам это необходимо, потому что буфер не заполнен.

```go
package main

import "fmt"

func sender(c chan int) {
    c <- 1 // len 1, cap 3
    c <- 2 // len 2, cap 3
    c <- 3 // len 3, cap 3
    c <- 4 // <- goroutine blocks here
    close(c)
}

func main() {
    c := make(chan int, 3)

    go sender(c)

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

    // read values from c (blocked here)
    for val := range c {
        fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
    }
}
```


Вывод программы:


Length of channel c is 0 and capacity of channel c is 3
Length of channel c after value '1' read is 3
Length of channel c after value '2' read is 2
Length of channel c after value '3' read is 1
Length of channel c after value '4' read is 0


```go
Дополнительный пример с буферизированным каналом:


package main

import (
    "fmt"
    "runtime"
)

func squares(c chan int) {
    for i := 0; i < 4; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)
    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())
    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())

    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())

    c <- 5
    c <- 6
    c <- 7
    c <- 8 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())
    fmt.Println("main() stopped")
}
```

```
Вывод программы:


main() started
active goroutines 2
1
4
9
16
active goroutines 1
active goroutines 2
25
36
49
64
active goroutines 1
main() stopped
```

Используя буферизованный канал и цикл for range, мы можем читать с закрытых каналов. Поскольку у закрытых каналов данные все еще живут в буфере, их можно считать:

```go
package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2
    c <- 3
    close(c)

    // iteration terminates after receiving 3 values
    for elem := range c {
        fmt.Println(elem)
    }
}
```

Работа с несколькими горутинами

Давайте напишем 2 горутины, одна для вычисления квадрата целого числа, а другая для вычисления куба:

```go
package main

import "fmt"

func square(c chan int) {
    fmt.Println("[square] reading")
    num := <-c
    c <- num * num
}

func cube(c chan int) {
    fmt.Println("[cube] reading")
    num := <-c
    c <- num * num * num
}

func main() {
    fmt.Println("[main] main() started")

    squareChan := make(chan int)
    cubeChan := make(chan int)

    go square(squareChan)
    go cube(cubeChan)

    testNum := 3
    fmt.Println("[main] sent testNum to squareChan")

    squareChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] sent testNum to cubeChan")

    cubeChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] reading from channels")

    squareVal, cubeVal := <-squareChan, <-cubeChan
    sum := squareVal + cubeVal

    fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
    fmt.Println("[main] main() stopped")
}
```

Однонаправленные каналы

До сих пор мы видели каналы, которые могут передавать и принимать данные. Но мы также можем создать канал, который будет однонаправленным. Например, канал, который сможет только считывать данные, и канал который сможет только записывать их.


Однонаправленный канал также создается с использованием make, но с дополнительным стрелочным синтаксисом.

```go
roc := make(<-chan int)
soc := make(chan<- int)
```

Где roc канал для чтения, а soc канал для записи. Следует заметить, что каналы также имеют разный тип.

```go
package main

import "fmt"

func main() {
    roc := make(<-chan int)
    soc := make(chan<- int)

    fmt.Printf("Data type of roc is `%T`\n", roc)
    fmt.Printf("Data type of soc is `%T\n", soc)
}
```

Но в чем смысл использования однонаправленного канала? Использование однонаправленного канала улучшает безопасность типов в программe, что, как следствие, порождает меньше ошибок.


Но допустим, что у вас есть программа, в которой вам нужно только читать данные из канала, а основная программа должна иметь возможность читать и записывать данные из/в тот же канал. Как это будет работать?


К счастью Go предоставляет простой синтаксис для преобразования двунаправленного канала в однонаправленный канал.

```go
import "fmt"

func greet(roc <-chan string) {
    fmt.Println("Hello " + <-roc + "!")
}

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    go greet(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

Анонимные горутины

Каналы также могут работать и с анонимными горутинами. Давайте изменим предыдущий пример, используя анонимные горутины.
Вот что у нас получилось:

```go
package main

import "fmt"

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    // launch anonymous goroutine
    go func(c chan string) {
        fmt.Println("Hello " + <-c + "!")
    }(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

Канал с типом данных канала

Каналы являются объектами первого класса, то есть они могут быть использованы как значение элемента структуры, или аргументы функции, как возврат значения из функции/метода и даже как тип для другого канала. В примере ниже мы используем канал в качестве типа данных для другого канала:

```go
package main

import "fmt"

// gets a channel and prints the greeting by reading from channel
func greet(c chan string) {
    fmt.Println("Hello " + <-c + "!")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
    c := make(chan string)
    cc <- c
}

func main() {
    fmt.Println("main() started")

    // make a channel `cc` of data type channel of string data type
    cc := make(chan chan string)

    go greeter(cc) // start `greeter` goroutine using `cc` channel

    // receive a channel `c` from `greeter` goroutine
    c := <-cc

    go greet(c) // start `greet` goroutine using `c` channel

    // send data to `c` channel
    c <- "John"

    fmt.Println("main() stopped")
}
```

**select**

select похож на switch без аргументов, но он может использоваться только для операций с каналами. Оператор select используется для выполнения операции только с одним из множества каналов, условно выбранного блоком case.

Давай взглянем на пример ниже, и обсудим как он работает:

```go
package main

import (
    "fmt"
    "time"
)

var start time.Time
func init() {
    start = time.Now()
}

func service1(c chan string) {
    time.Sleep(3 * time.Second)
    c <- "Hello from service 1"
}

func service2(c chan string) {
    time.Sleep(5 * time.Second)
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

В этом примере мы используем оператор select как switch, но вместо булевых операций, мы используем операции для чтения данных из канала. Оператор select также является блокируемым, за исключением использования default(позже вы увидите пример с его использованием). После выполнения одного из блоков case, горутина main будет разблокирована. Задались вопросом когда case условие выполнится?


Если все блоки case являются блокируемыми, тогда select будет ждать до момента, пока один из блоков case разблокируется и будет выполнен. Если несколько или все канальные операции не блокируемы, тогда один из неблокируемых case будет выбран случайным образом (Примечание переводчика: имеется ввиду случай, когда пришли одновременно данные из двух и более каналов).


Давайте наконец разберем программу, которую написали ранее. Мы запустили 2 горутины с независимыми каналами. Затем мы использовали оператор select c двумя case операторами. Один case считывает данные из chan1 а другой из chan2. Так как каналы не используют буфер, операция чтения будет блокируемой. Таким образом оба case будут блокируемыми и select будет ждать до тех пор, пока один из case не разблокируется.


Когда программа находится в блоке select горутина main будет заблокирована и будут запланированы все горутины (по одной за раз), которые используются в блоке select, в нашем случае это service1 и service2. service1 ждет 3 секунды, после чего будет разблокирован и сможет записать данные в chan1. Таким же образом как и service1 действует service2, только он ожидает 5 секунд и осуществляет запись в chan2. Так как service1 разблокируется раньше, чем service2, первый case разблокируется раньше и произведет чтение из chan1, а второй case будет проигнорирован. После чего управление вернется в main, и программа завершится после вывода в консоль.

Вышеприведенная программа имитирует реальный веб-сервис, в котором балансировщик нагрузки получает миллионы запросов и должен возвращать ответ от одной из доступных служб. Используя стандартные горутины, каналы и select, мы можем запросить ответ у нескольких сервисов, и тот, который ответит раньше всех, может быть использован.

Для того, чтобы сделать все блоки case неблокируемыми, мы можем использовать каналы с буфером.

```go
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func main() {
    fmt.Println("main() started", time.Since(start))
    chan1 := make(chan string, 2)
    chan2 := make(chan string, 2)

    chan1 <- "Value 1"
    chan1 <- "Value 2"
    chan2 <- "Value 1"
    chan2 <- "Value 2"

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from chan2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

default case

Так же как и switch, оператор select поддерживает оператор default. Оператор default является неблокируемым, но это еще не все, оператор default делает блок select всегда неблокируемым. Это означает, что операции отправки и чтение на любом канале (не имеет значения будет ли канал с буфером или без) всегда будут неблокируемыми.


Если значение будет доступно на каком-либо канале, то select выполнит этот case. Если нет, то он немедленно выполнит default.

# 4.1 Чем отличаются буферизированные и не буферизированные каналы?

![alt text](images/channels.png)


# 5. Какой размер у структуры struct{}{}?

Пустая структура - это тип struct, не имеющий полей. Вот несколько примеров в именованной и анонимной формах

Прежде чем перейти к рассмотрению самой пустой структуры, я хотел бы сделать небольшое отступление и поговорить о ширине.

Термин width, как и большинство других терминов, пришел из компилятора gc, хотя его этимология, вероятно, уходит на десятки лет назад.

Ширина описывает количество байт памяти, которое занимает экземпляр типа. Поскольку адресное пространство процесса одномерно, я считаю, что ширина - более подходящий термин, чем размер.

Ширина - это свойство типа. Поскольку каждое значение в программе на Go имеет тип, ширина значения определяется его типом и всегда кратна 8 битам.

Мы можем узнать ширину любого значения, а значит, и ширину его типа с помощью функции unsafe.Sizeof().

Пустая структура
Теперь, когда мы изучили ширину, должно быть очевидно, что пустая структура имеет ширину, равную нулю. Она занимает ноль байт памяти.

Поскольку пустая структура занимает ноль байт, из этого следует, что она не нуждается в заполнении. Таким образом, структура, состоящая из пустых структур, также не занимает места.

В соответствии с принципом ортогональности (orthogonality) Go, пустая структура - это такой же тип структуры, как и любой другой. Все свойства, привычные для обычных структур, одинаково применимы к пустой структуре.

Вы можете объявить массив struct{}, но они, конечно, не занимают места в памяти.

Слайс struct{} занимает только места сколько занимает его заголовок(header). Как было показано выше, их массив, содержащий пустые структуры не занимает места.

Все это справедливо и для подслайса (subslice). len и cap работают, как и ожидается.

Теперь, когда мы продемонстрировали, что пустые структуры ведут себя так же, как и любые другие типы, из этого следует, что мы можем использовать их в качестве получателей методов.

# 6. Есть ли в Go перегрузка методов или операторов?

No

# 7. В какой последовательности будут выведены элементы map[int]int?

Пример:
m[0]=1 m[1]=124 m[2]=281

Рандомно, превалирует 1 124 281

# 8. В чем разница make и new

The new() function in Go is a built-in function that allocates memory for a new zeroed value of a specified type and returns a pointer to it. It is primarily used for initializing and obtaining a pointer to a newly allocated zeroed value of a given type, usually for data types like structs.

Here's a simple example:

```go
package main

import "fmt"

type Person struct {
    Name 	string
    Age  	int
    Gender 	string
}

func main() {
    // Using new() to allocate memory for a Person struct
    p := new(Person)

    // Initializing the fields
    p.Name = "John Doe"
    p.Age = 30
    p.Gender = "Male"

    fmt.Println(p)
}
```

In this example, new(Person) allocates memory for a new Person struct, and p is a pointer to the newly allocated zeroed value.

On the other hand, the make() function is used for initializing slices, maps, and channels – data structures that require runtime initialization. Unlike new(), make() returns an initialized (non-zeroed) value of a specified type.

Let's look at an example using a slice:

```go
package main

import "fmt"

func main() {
    // Using make() to create a slice with a specified length and capacity
    s := make([]int, 10, 15)

    // Initializing the elements
    for i := 0; i < 10; i++ {
        s[i] = i + 1
    }

    fmt.Println(s)
}
```

When dealing with value types like structs, you can use new() to allocate memory for a new zeroed value. This is suitable for scenarios where you want a pointer to an initialized structure.

For slices, maps, and channels, where initialization involves setting up data structures and internal pointers, use make() to create an initialized instance.

Keep in mind that new() returns a pointer, while make() returns a non-zeroed value. Choose the appropriate method based on whether you need a pointer or an initialized value.

Understanding the distinction between new() and make() in Go is crucial for writing clean and efficient code. By using the right method for the appropriate data types, you can ensure proper memory allocation and initialization in your Go programs.

# 9. Сколько существует способов задать перенную типа slice или map

# 10. Что выведет данная программа и почему?

```go
func update(p *int) { 
    b := 2
    p = &b
}

func main() { 
    var (
    a=1
    p = &a )
    fmt.Println(*p) 
    update(p) 
    fmt.Println(*p)
}
```

# 11. Что выведет данная программа и почему?

```go
package main

import (
	"fmt"
	"sync"
)

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

Выведет 0...4 в случайном порядке, потом случится deadlock, т.к. в горутину передается WaitGroup по значению, поэтому сигнализация об выполнении элемента группы. Метод Wait() ожидает завершения всех горутин из группы, но внутренний счетчик в WaitGroup будет равен 5, поэтому главный поток блокируется навсегда.

# 12. Что выведет данная программа и почему?

```go
package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
```

Выведется 0. Создается переменная в локальной области видимости, она и будет инкрементироваться в блоке if.

# 13. Что выведет данная программа и почему?

```go
package main

import "fmt"

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

Выведется 100 2 3 4 5. Чтобы добавился новый элемент, нужно вернуть из функции слайс, он имеет уже новый адрес из реалокации памяти но новый адрес. Однако значение в нулевом индексе изменится, т.к. слайс изначально передается по указателю.


# 14. Что выведет данная программа и почему?

```go
package main

import "fmt"

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

Выведется [b b a] [a a]. Элемент "a" не добавится, т.к. возвращаемый слайс после операции `slice = append(slice, "a")` будет иметь другой адрес, т.к. происходит реаллокация слайса на новую область памяти.