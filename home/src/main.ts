import './style.css'
import 'uno.css'

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <h1 class="m-0">My Talks</h1>
  <div class="text-gray-500 mb-10">Yamanaka Junichi</div>
  <img 
    src="https://github.com/JY8752.png"
    alt="アイコン"
    class="rounded-full w-60 h-60"
  />

  <div class="flex flex-col">
    <h2>2024/6</h2>
    <a href="/2024/not-hate-go">Goを嫌いにならないためのメンタルモデル</a>
  </div>
`