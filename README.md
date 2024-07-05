# open main.go file

# package main

# import "fmt"

# go mod init go_tutorials  ## module ismini kendimiz verdik

# go run main.go

# https://go.dev/ref/spec#Numeric_types


# susluparantez alt 7


# 
````bash
func updateName(x string) string {
    x = "wedge"
    return x
}
````
Bu fonksiyonun çalışma mantığı şu şekildedir:

> updateName adında bir fonksiyon tanımlanır.
> x adında bir string parametre alır.
> Fonksiyon içinde x değeri "wedge" olarak güncellenir.
> Güncellenmiş x string'i return ifadesi ile fonksiyon dışına döndürülür.

Örneğin, bu fonksiyonu çağırdığınızda
fmt.Println(updateName("Luke"))

Çıktı olarak "wedge" elde edersiniz, çünkü fonksiyon parametresi olarak verilen "Luke" değeri fonksiyon içinde "wedge" olarak değiştirilmiş ve geri döndürülmüştür.

# 
````bash
func updateName() string {
    x := "wedge"
    return x
}
````
Bu fonksiyonun çalışma mantığı şu şekildedir:

> updateName adında bir fonksiyon tanımlanır.
> x adında bir string değişkeni tanımlanır ve "wedge" değeri atanır.
> Tanımlanan x string değişkeni return ifadesi ile fonksiyon dışına döndürülür.

Örneğin, bu fonksiyonu çağırdığınızda:
fmt.Println(updateName())

Çıktı olarak "wedge" elde edersiniz, çünkü fonksiyon her zaman "wedge" değerini döndürmektedir. Bu fonksiyon parametre almaz ve her zaman aynı değeri döndürür.

# 
İlk fonksiyon ve ikinci fonksiyon arasındaki farklar şunlardır:

Parametre Alıp Almama:

İlk fonksiyon updateName(x string) string bir parametre alır (x adında bir string) ve bu parametreyi günceller.
İkinci fonksiyon updateName() string ise herhangi bir parametre almaz, sadece sabit bir string değerini döndürür.
Değişken Kullanımı:

İlk fonksiyonda x adında bir değişken tanımlanır ve bu değişken fonksiyon içinde güncellenir.
İkinci fonksiyonda ise x sabit bir değer olarak fonksiyonun başında tanımlanır ve değişmez bir şekilde döndürülür.
Kullanım Senaryosu:

İlk fonksiyon genellikle dışarıdan alınan bir değeri işleyip dönüştürmek veya güncellemek için kullanılır. Örneğin, bir kullanıcının adını veya bir nesnenin adını güncellemek gibi.
İkinci fonksiyon ise genellikle sabit bir değeri döndürmek için kullanılır. Örneğin, bir sabit metin veya bir sabit sayı gibi.
Özetle, fonksiyonlar parametre alma ve döndürdükleri değerler açısından farklılık gösterir. İlk fonksiyon parametre alır ve verilen değeri güncellerken, ikinci fonksiyon parametre almaz ve her zaman sabit bir değer döndürür.



# group A types --> strings, ints, bools, floats, arrays, structs: These are non-pointer values
# group B types --> slices, maps, functions: These are pointer values. 



\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

````bash
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name)
````


Bu iki satır arasındaki fark, Go programlama dilinde değişken tanımlama ve atama yöntemlerinin farklı kullanımıyla ilgilidir.

İlk satırda := operatörü kullanılarak bir değişken tanımlanıyor ve değer atanıyor. Bu operatör, Go dilinde "kısa değişken tanımlama" olarak bilinir. Örneğin:


> name, _ := reader.ReadString('\n')
B
urada name değişkeni reader.ReadString('\n') fonksiyonunun döndürdüğü değere atanıyor. := operatörü, değişkenin tipini belirtmeden kullanılabilir ve değişkenin zaten tanımlandığı veya daha önce kullanılmadığı durumlarda kullanılır. Bu operatör, Go'da yaygın olarak kullanılan bir pratiktir ve kodun daha okunabilir ve kısa olmasını sağlar.

İkinci satırda ise sadece = operatörü kullanılarak bir değişkene değer atanıyor. Bu durumda, değişkenin zaten tanımlandığı ve sadece yeni bir değer atandığı varsayılır. Örneğin:


> name = strings.TrimSpace(name)
Burada name değişkeninin mevcut değeri, strings.TrimSpace(name) fonksiyonunun döndürdüğü değerle değiştiriliyor.

Özetle:

> := operatörü, değişken tanımlama ve değer atama işlemlerini aynı anda yapmak için kullanılır ve değişkenin daha önce tanımlanmadığı durumlarda tercih edilir.

> = operatörü, zaten tanımlı bir değişkene yeni bir değer atamak için kullanılır.

Go dilinde tipik bir kullanım, değişken tanımlama ve atama işlemlerini := operatörüyle yapmaya özen göstermek ve sadece değer atamalarında = operatörünü kullanmaktır.

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\


