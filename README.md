# 軍儀

## 軍儀棋(日語:軍儀(發音[ぐんぎ](https://translate.google.com.tw/?sl=ja&tl=zh-TW&text=%E3%81%90%E3%82%93%E3%81%8E&op=translate)))是 **_《HUNTER×HUNTER 獵人》蟻王篇_** 漫畫第 244 至 260 話，與動畫第 102 至 112 所出現的虛構遊戲

## 在 2022 的 2 月情人節時由[官方](https://store.universal-music.co.jp/s/hunter-gungi/)釋出實體化訊息並開放預購

### 根據官方的規則介紹與部分有預購實體商品而收到成品的紛絲，所提供的訊息而做出這份筆記，以下分別先說明與定義各部件的規範：

-   棋盤：規格為 $9\times 9$ 的正方形棋盤，共有 81 格的棋子進行區，類似將棋的棋盤。
-   駒台：用於放置尚未入場的棋子，類似將期的駒台。
-   棋子：兩邊各有 14 種共 25 個的棋子，如[詳細](https://github.com/littletrainee/GunGiResourceAndNote/blob/main/%E6%A3%8B%E7%A8%AE.md)
-   移除區(可有可無?)：被吃掉的棋子放置區，可有可無，類似象棋?

### 根據官方實體化的解說書第[6](https://github.com/littletrainee/GunGiResourceAndNote/blob/main/%E8%A7%A3%E8%AA%AA%E6%9B%B8/6.jpg)頁的敘述，流程如下：

1. 雙方各選一個顏色作為自己的顏色(黑或白)
2. 待選定顏色後雙方擲棋決定先後手順序，由任意一名玩家將黑棋與白棋的帥握在手中，並在棋盤中心上方讓棋子自然落下，不論是否正反面，離中心點最近的顏色為先手
3. 在最初布陣的 **第一巡** 雙方玩家皆只能放 **帥**，第二巡後可以依照個人喜好佈置棋子，由先手先行放置，當先手玩家決定不再配置時可以宣告"配置完成"或"佈陣完成"，布陣期間雙方僅能在離自身最近的三列內放置。
4. 先手佈陣完成後，後手可以持續進行配置，直到後手宣告“佈陣完成”，後，進入正式對弈的階段
5. 未佈陣進入棋盤中的棋需放置於駒台上，以當作待機的棋子用

正式對弈階段

在每一巡當中，雙方都可以從兩個動作擇一進行，分別**新**與**移駒**

-   **新(あらた)**：將自己備用區的其中一枚棋子放置於棋盤上，最遠距離為不超過離自己的最遠棋子範圍，亦可以堆疊在非 **帥** 的棋子上方。

-   **移駒**：選擇場上任一格子當中最上方的己方棋子做移動，每個棋子的移動方式需參照各個棋子移動的方式為主<a href="https://github.com/littletrainee/GunGiResourceAndNote/blob/main/%E8%A7%A3%E8%AA%AA%E6%9B%B8/3.jpg"><sup>[P3]</sup></a><a href="https://github.com/littletrainee/GunGiResourceAndNote/blob/main/%E8%A7%A3%E8%AA%AA%E6%9B%B8/4.jpg"><sup>[P4]</sup></a><a href="https://github.com/littletrainee/GunGiResourceAndNote/blob/main/%E8%A7%A3%E8%AA%AA%E6%9B%B8/5.jpg"><sup>[P5]</sup></a>，可以堆疊在除了帥以外的所有棋子上方，但不可超過當前級別限制最高的層數，移動至堆疊於非帥的棋子上方後，可以有兩種動作分別為**俘獲**或**控制**。
    -   **俘獲**：當堆疊在超過已經有兩層的棋子上方時可以選擇俘獲的方式將對方的棋子移出遊戲(類似**象棋**的吃或是**西洋棋**的 capture、吃子或取子)
    -   **控制**：將棋子移動到目標棋子的上方，使下方的棋子被控制住無法移動或進行動作

遊戲進行期間除了**砲**、**筒**與**弓**可以進行跳棋以外其餘棋子皆無法進行跳棋，若這三個特殊棋的前方有高於自己段數的棋子在自己前方時，擇無法進行跳棋方式的前進，類似**象棋**的**柺馬腳**。

**謀**同樣屬於特殊棋，雖然沒有像是**砲**、**筒**與**弓**一樣有跳棋功能，但是卻可以使用背叛，若自己的備牌區有與被堆疊的棋子一樣的棋子時，對其進行**俘獲**可以將己方駒台上的旗子替換為被堆疊的對方棋子，使對方棋子移出遊戲並增加我方棋子。

在遊戲進行期間雙方輪流進行移駒或新的動作直到將對方的**帥**移出遊戲為止，才算結束這局對弈。

## 遊戲報讀棋子移動的方式被定義為[橫-縱-段-駒]，橫為水平從右到左，縱為垂直從上到下，段為當前棋子的最後的段數，駒為當前的棋子。

---

本儲存庫為軍儀棋的棋譜產生器，可以直接匯入在代碼內使用

```Go
import "github.com/littletrainee/GunGiNotationGenerator"
```

並在 main()下面使用以下行即可自動運行本儲存庫代碼

```Go
var {想要的變數名稱} GunGiNotationGenerator.GunGiNotationGenerator = GunGiNotationGenerator.GunGiNotationGenerator
{想要的變數名稱}.Initilization()
{想要的變數名稱}.Start()
```

或直接使用儲存庫下的[build.exe](https://github.com/littletrainee/GunGiNotationGenerator/blob/main/build.exe)檔案。
