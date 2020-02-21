# データ構造とリファクタリング

## 型の種類

- 値
  - 真偽値、数値、文字列、日付時刻など
  - 列挙型
- ポインタ
- コレクション Collection
    - 配列 Array
    - リスト List
    - 集合 Set
    - 連想配列 Map
- 構造体
- クラス

## 言語でのサポート状況

どの言語でも機能を作ることはできるが、ここでは言語の機能あるいは標準的なライブラリでサポートされているものを挙げる。

型 | Delphi | Go | ES5 | Ruby
--|-------|-------|-------|-------
列挙型 | (....) | *1 | 無 | 無
ポインタ | 有 | 有 | 無 | 無
配列 | array of type | []type | array | Array
リスト | TObjectList, etc | あまり使わない | array | Array
集合 | set of 順序型 | 無 | Set | (Array#uniq?)
連想配列 | TStringsのValues | map[型1]型2 | Object | Hash
構造体 | record | struct | Object | Struct
クラス | TObjectを継承 | *2 | Object | Object

## 関数とメソッド

何らかのクラスやオブジェクトなどのデータ構造に依存せずに存在するのが関数であり、依存するのがメソッドである。
関数のスコープは言語によって異なるが、グローバルなスコープの関数を多用するのはよろしくない。
できるだけ影響範囲を適切に絞って定義するべきである。

関数が特定のデータ構造に関係するのであれば、それはそのデータ構造についてのメソッドに変更するべきである。


## リファクタリングとは

誤解を恐れずに言えば、リファクタリングとは、データ構造を見直し、そのデータ構造に適した関数やメソッドを配置することである。
よって、どのようなデータ構造を持つべきであるのかを知らなければ、良いリファクタリングはできない。

## 良い構造とは

- 良い名前
    - 名前重要！
    - チーム内で合意が取れればOKだが・・・
    - 複数形
- 良い形
    - 適切な型
    - 適切な関数・メソッド
- テストのしやすさ

## テストのしやすさ

- オブジェクトや構造体は状態を持つ
    - 状態を持っているとテストが大変
    - 副作用の少ない形へ
    - mutable から immutable へ

## JSONの例

https://www.sitepoint.com/10-example-json-files/

http://docwiki.embarcadero.com/RADStudio/Rio/ja/JSON_オブジェクト_フレームワーク


## 構造とコレクション

コレクションなしで複雑な構造が作られることは稀。
なので、コレクションをどう扱うかが重要。

コレクションの操作で代表的なものはこちら

- Filter (Select)
- Map
- Reduce


## ライブラリの扱い

現時点のテクノ産業ではライブラリの作成は基本的に禁止されている。
これはライブラリの変更によってアプリケーションをリリースしなければならないという本末転倒な事態が起きていたため。

ただし厳密にはライブラリを作ることを禁止していない。
主要な部分について自動テストが書かれた状態であり、そのリリースが適切にコントロールできる見込みであればOK。


## リファクタリングとライブラリ

本来この２つは別々に扱うべきトピックですが、似たコードが複数のフォームに分散している問題を持つテクノ産業にとって、
リファクタリングで行うべきことは関数やクラスの抽出であり、その結果できるものはアプリケーション内ライブラリというべきもの。
ではライブラリとはどういう形であるべきか。

- 良い名前
    - 名前重要！
- 良い形
    - (アプリケーション内なので提供方法は問題ないはず？)
    - 適切な型
    - 適切な関数・メソッド
- テストのしやすさ

良い構造として挙げたものと同じもの。

アプリケーション内ライブラリを作れば良い構造というわけではないが、重複したものが多数存在する構造よりは良いはず。

## ライブラリとは？

なんらかの構造毎に分割された型と関数・メソッドを何らかの方法で配布できる形にしたもの。

以下は、テクノに関係する言語とそれらの代表的なライブラリのための言語サポートと配布方法。

言語 | ライブラリを作る言語的サポート | 配布方法
-----|------------|--------------
Delphi | ユニット、パッケージ | ファイル
Go | パッケージ | HTTP(組み込み)
ES5 | パッケージ | npm
Ruby | gem | rubygems

## ライブラリとインタフェース

インタフェースとは、クラス（あるいはそれに準じるもの）について、内部の状態を問わずメソッドを呼び出す方法を提供するもので、型の一種である。
DelphiやGo、Javaなど静的型付け言語に用意されていることが多い。
ES5やRubyは動的型付け言語であるためダックタイピングと呼ばれるように型を意識せずメソッドを呼び出すことができる（メソッドが存在しなければ例外）。

さて、ECMAScriptの正統的な型付け言語である、ただしベースがESであるため静的型付け言語とは言い難いTypeScriptはというと、
もちろんインタフェースは提供されている。ただしDelphiやJavaのインタフェースとは異なる。

DelphiやJavaのインタフェースは、そのインタフェースを実装するクラスを定義する際に、どのインタフェースを実装するのかを明示する必要がある。
しかしGoやTypeScriptのインタフェースは、クラスを定義する際にインタフェースを指定する必要はない。
別の言い方をすれば、GoやTypeScriptのインタフェースは、使う側が既存のクラスのメソッドから使用するものを選んで作ることができる型である。



### Delphiでのインタフェースの注意

ガベッジコレクタのないDelphiで、インタフェース型が多用されると、生成したオブジェクトのメモリ管理が難しくなる可能性がある。


## テストのしやすさとは

- 状態を持たない、あるいは持っていても少ない
- 現在時刻や環境変数など外部的な要因に依存しない


## リファクタリングの目指すべき方向性

コレ自身はリファクタリングとは呼びにくいものですが、こんがらがったコードを解き明かしていくために
このような変更を行っていきたいと思っています。
https://github.com/tecowl/data_structures/pull/1
