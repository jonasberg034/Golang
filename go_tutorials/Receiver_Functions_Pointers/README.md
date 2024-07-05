###

````bash

func (b *bill) addItem(name string, price float64) {
    b.items[name] = price
}

````

Bu fonksiyon bir bill tipi için tanımlanmış bir metottur (func (b *bill)). *bill kullanılarak tanımlanan bir pointer alıcıdır, yani bir bill yapısının bir işaretçisi üzerinde işlem yapmasını sağlar.

>Alıcı Tipi: 
addItem fonksiyonu bill tipi ile ilişkilendirilmiştir. *bill, işaretçi alıcı olduğunu gösterir; bu da metodu doğrudan bill yapısında değişiklik yapabilmesini sağlar.

> Parametreler:

>name string: 
Bu parametre, faturaya eklenen öğenin adını temsil eder.
price float64: Bu parametre, öğenin fiyatını temsil eder.
Fonksiyon Gövdesi:

> b.items[name] = price: 
Bu satır, price değerini bill yapısının (b) items haritasındaki name anahtarına atar. Temelde, items haritasına öğenin adı (name) anahtar olarak ve öğenin fiyatı (price) değer olarak ekler veya günceller.

Bu satır, b işaretçisinin gösterdiği bill yapısının items alanına (b.items) yeni bir öğe ekler. name parametresi, öğenin adını temsil eder ve bu adı price parametresi ile ilişkilendirir, yani öğenin fiyatını belirler.


###

````bash
func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

````

> Alıcı Tipi: 

updateTip fonksiyonu, bill tipi ile ilişkilendirilmiş bir metottur (func (b *bill)). b parametresi bir *bill türündedir, yani bill yapısının bir işaretçisidir. Bu sayede fonksiyon, bill yapısının içindeki verileri doğrudan değiştirebilir.

> Parametre:

tip float64: Bu parametre, faturaya eklenmek veya güncellenmek istenen bahşiş miktarını temsil eder.

> Fonksiyon Gövdesi:

b.tip = tip: Bu satır, b işaretçisinin gösterdiği bill yapısının tip alanına (b.tip) tip değerini atar. Böylece, faturanın bahşiş alanını güncellemiş oluruz.
