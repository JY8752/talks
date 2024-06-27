---
title: Goを嫌いにならないためのメンタルモデル
class: text-center
transition: fade-out
mdc: true
---

# Goを嫌いにならないためのメンタルモデル

<style>
  h1 {
    font-size: 4em !important;
  }
</style>

<div class="pt-10 opacity-50">
  2024/06/28 Yamanaka Junichi
</div>

---

## 自己紹介

<div class="grid grid-cols-2 h-screen mt-5">
  <div class="pr-5">
    <h3 class="pt-5">Yamanaka Junichi</h3>
    <h3 class="py-5">経歴</h3>
    <ul>
      <li> 2013年 ~ 2020年　Jcomで営業</li>
      <li> 2020年 ~ 2022年 Javaで受託開発</li>
      <li> 2022年 ~ 2024年 Coconeでブロックチェーンゲームの開発をKotlinやGoで</li>
      <li> 2024年 ~ フリーランス</li>
    </ul>
    <h3 class="py-5">好きな技術</h3>
    Go, Protobuf, Unit Test
  </div>
  <div class="pl-5">
    <img 
      src="https://github.com/JY8752.png"
      alt="アイコン"
      class="rounded-full w-80 h-80 mt-15"
    />
  </div>
</div>

---

## 今日話したいこと

<v-click>

- 🚀 Goを好きになるには**今までの言語からのパラダイムシフト**が必要(かもしれない)

</v-click>

<v-click>

- 🥑 Goの**言語仕様が少ないことには理由が必ずある**

</v-click>

<v-click>

- 🐢 Goは**独自の文化やルールを強制しているわけではない**

</v-click>

<v-click>
<span text-4xl text-blue-300 py-10>
Goが合わないとならない!!とならないようにこれらのことを知っておきましょう😇
</span>
</v-click>

<style>
  ul {
    list-style-type: none !important;
    padding-left: 0;
  }
  li {
    font-size: 2rem;
    padding: 1rem;
  }
</style>

<!--
- Goはオブジェクト指向言語や関数型言語なんかとは違った世界線の言語
- そのため、こういったことを知らないとGoが合わないと思ってしまう人がけっこう多い気がする
- 今回はそういった悲しい誤解が少しでも減ればいいなと思います
-->

---

## Java脳でGoを書くとどうなるか

<br />

### オブジェクトのカプセル化ができなくて悩む

<div grid grid-cols-2>
<Transform :scale="0.9">
Java

```java {all|3,4}
public class User {

    private String name;
    private int age;

    // コンストラクタ
    public User(String name, int age) {
        this.name  = name;
        this.age = age;
    }

    // getter
    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }
}
```

</Transform>

<Transform :scale="0.9">
Go

```go {2|all|3-6|20-}
package user

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	return User{name: name, age: age}
}

func (u User) Name() string {
	return u.name
}

func (u User) Age() int {
	return u.age
}

// 同じパッケージからだと
// プライベートなフィールドにもアクセスできてしまう
func UpdateName(name string, user *User) {
	user.name = name
}

```

</Transform>
</div>

<!--
- 一例としてオブジェクトのカプセル化についてお話しできればと思います
- こちらはJavaとGoでUserというオブジェクトを表現したコード例です
-->

---

### nullチェックがしたくなる

<div grid grid-cols-2>

<Transform :scale="0.9">

Java

```java {all|8-13|16-}
public class Item {
    private final String id;

    private Item(String id) {
        this.id = id;
    }

    public static Item createItem(String id) {
        if (id.startsWith("d-")) {
            return null;
        }
        return new Item(id);
    }
}

// Main.java
var item = Item.createItem("d-1111");
if(item == null) {
    // 何か処理
}
```

</Transform>

<Transform :scale="0.9">

Go

```go {7|all|2-4|8-}
func NewItem(id string) (Item, error) {
	if strings.HasPrefix(id, "d-") {
		return Item{}, fmt.Errorf("specific item not starting `d-` prefix ")
	}
	return Item{Id: id}, nil
}

func main() {
	item, err := NewItem("d-111")
	if err != nil {
		panic(err)
	}
}
```

</Transform>

</div>

<!--
もう一つ、　Java的な例でnullチェックの話です。
- 先ほどと同じようにJavaとGoの二つのコードを用意してます。
-->

---

### ここまでのまとめ

<br />

<div v-click px-5 text-3xl>

Goは**オブジェクト指向の言語ではない**

<div v-after p-5 text-2xl text-indigo-200>
  👉<span pl-5>Goはもともと<span font-bold>システムプログラミングを書く言語</span></span>
</div>

</div>

<div v-click text-3xl px-5 pt-10>

Goでは**変数に値があるかは気にしなくてもいい**

<div v-after p-5 text-2xl text-indigo-200><span pr-5>👉</span>ポインタ以外はデフォルト値が使われるためpanicしない</div>
<div v-after p-5 text-2xl text-indigo-200><span pr-5>👉</span>ポインタ型はnilになり得るがerrorを返すことでnil参照が起こらない</div>
<div v-after p-5 text-2xl text-indigo-200><span pr-5>👉</span>Goではnillチェックの代わりにerrorチェックをする</div>

</div>

<!--
- Goはシステムプログラミングをするための言語でどちらかといえばCやC++に近い言語です。
- Goでは適切にエラーハンドリングをすることで基本的には変数に値があるかどうかは気にしなくていいい、はずです
-->

---

## Goにないもの

<br />

### 配列・List操作

<br />

<div grid grid-cols-2>

<Transform :scale="0.9">
Kotlin

```kotlin {all|10-14}
data class Student(val math: Int, val english: Int)

fun main(args: Array<String>) {
  val students = listOf(
  Student(math = 70, english = 81),
  Student(math = 70, english = 81),
  Student(math = 70, english = 81),
  )

    students.filter { it.math > 70 }
        .filter { it.english > 80 }
        .map { it.math + it.english }
        .forEach { println("total: $it") }

}

```

</Transform>

<Transform :scale="0.9">
Go

```go {5|all|13-}
type Student struct {
	math    int
	english int
}

func main() {
	students := []Student{
		{math: 71, english: 81},
		{math: 71, english: 81},
		{math: 71, english: 81},
	}

	for _, s := range students {
		if s.math <= 70 || s.english <= 80 {
			continue
		}
		sum := s.math + s.english
		fmt.Printf("total: %d\n", sum)
	}
}
```

</Transform>

</div>

<!--
次に他の言語にあってGoにないモノの話でこれもめちゃくちゃあるんですが有名なので配列・Listの話をします。
-->

---

### enum

<br />

<div grid grid-cols-2>
<Transform :scale="0.9">
Kotlin

```kotlin {all|-5|7-|all}
enum class Status{
    Success,
    Failed
    ;
}

fun main() {
    // code 0 name Success
    // code 1 name Failed
    for (status in Status.entries) {
        println("code ${status.ordinal} name ${status.name}")
    }
}
```

</Transform>
<Transform :scale="0.9">
Go

```go {7|all|-5|6-}
type status int
const (
	_ status = iota
	success
	failed
)

func main() {
	// 1 2
	fmt.Printf("%v %v", success, failed)
}
```

</Transform>
</div>

<div v-click text-3xl mt-5 text-blue-300>
他の言語にあるようなenumの機能が欲しければstringerやenumerというモジュールもある
</div>

<!--
enumは飛ばします
-->

---

### ここまでのまとめ

<br />

<div v-click px-3 text-3xl>

Goにはなぜmapやfilterがないのか

<div v-after p-3 text-2xl text-indigo-300><span pr-5>👉</span>遅くなるから</div>
<div v-after p-3 text-2xl text-indigo-300><span pr-5>👉</span>Goは他の言語にあるからといってパフォーマンスを落とすような機能を取り入れない</div>

</div>

<div v-click px-3 text-3xl mt-8>

Goにはなぜenumがないのか

<div v-after p-3 text-2xl text-indigo-300><span pr-5>👉</span>Goにある既存の機能で十分enumとしての機能を提供することができる</div>
<div v-after p-3 text-2xl text-indigo-300><span pr-5>👉</span>enumがもたらす複雑さに比べて得られるものが少ない</div>

</div>

<!--
-
-->

---
layout: center
---

<span text-4xl>Goはベストよりもベターを採用する言語😇</span>

<!--
最後にGoの命名の話を少しできればと思います。
ここにのせたのはGoの単体テストのコード例なのですが...
-->

---

## Goの短い変数名について

<br />

<Transform :scale="0.8">
Go

```go {all|1|19-}
func TestExample(t *testing.T) {
	tests := map[string]struct {
		x        int
		y        int
		expected int
	}{
		"1 + 1 = 2": {
			x:        1,
			y:        1,
			expected: 2,
		},
		"0 + (-1) = -1": {
			x:        0,
			y:        -1,
			expected: -1,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			if tt.x+tt.y != tt.expected {
				t.Errorf("expected %d but %d", tt.expected, tt.x+tt.y)
			}
		})
	}
}
```

</Transform>

<!--
最後にGoの命名の話を少しできればと思います。 ここにのせたのはGoの単体テストのコード例なのですが...
-->

---

### パッケージ名を使った命名

syncパッケージの例

```go {all|1|12-14|1,12-14}
var once sync.Once

func Hello(ch chan struct{}) {
	time.Sleep(time.Second)
	ch<-struct{}{}
	close(ch)
}

func main() {
	ch := make(chan struct{})
	defer func ()  {
		once.Do(func() {
			close(ch)
		})
	}()

	go Hello(ch)

	<-ch
}
```

<!--
次にGoはパッケージ名込みで命名をする言語です。
これに関してはsyncパッケージがとてもよくできてるのでその紹介なんですが...
-->

---

## まとめ

<br />

<div v-click text-3xl>Goはシステムプログラミングのための言語</div>
<div v-after text-2xl p-5 text-indigo-300><span pr-5>👉</span>Goが合わないと思ったのなら今までの言語のメンタルモデルを一度壊しましょう😇</div>

<div v-click text-3xl mt-6>Goはベストよりもベターを選択する言語</div>
<div v-after text-2xl p-5 text-indigo-300><span pr-5>👉</span>なぜ？と思うものにはだいたい理由があるので調べてみるといいです 🔍</div>

<div v-click text-3xl mt-6>Goは命名についてどの言語よりも真剣に考えている</div>
<div v-after text-2xl p-5 text-indigo-300><span pr-5>👉</span>Goの短い命名は独自の文化でもルールでもないです</div>

---
layout: end
---

---

## 参考

<br />

- [Goのなぜ問答](https://zenn.dev/nobonobo/articles/9a9f12b27bfde9)
