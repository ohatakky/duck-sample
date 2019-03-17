- interfasceを満たす
特定のinterfaceのメソッドを全て実装すること

- ダックタイピング
interfaceAを満たした実体はAである。
ダックタイピングが、is - aの機能をもたらす。
goではインターフェースを用いてis-a関係を型の間に築く。

- 委譲と埋め込みの違い
委譲も埋め込みもinterfaceを満たすことについては変わりはないが、
埋め込みは既存のstructを持たないことが異なる。
既存の型に依存する処理がある場合、予期せぬ動作を起こすことがある。

- わからないこと
- 委譲と継承の違い
委譲は、has - a
継承は、is - a
どんな問題がある？？？？？？？
継承のデメリット : SRPに違反することがある。
じゃあ委譲にすればどうなる？ : SRPに違反しなくなる。

ansとsvで、同じ料金計算クラスを継承したとき、料金計算クラスは修正できなくなる。はどう説明する？
SRP違反なので別々の料金計算クラスを用意する。ここに委譲、継承という言葉は介入しない？

まず、SRP違反にならないようクラスを2つ作る。
委譲にすることで料金計算に関する責務を負わない。
テストをするために、interfaceに依存するようDIする。(委譲とDIは異なるのか？)

- 委譲とDIの違い

- 委譲と埋め込みの違い
委譲は、has - a
埋め込みは、is - a
goの埋め込みはis-aだが継承ではない。

- phpの委譲とgoの委譲
php : class , field, construct
go : struct, field

- SOLID原則とは
前提として、SOLID原則はOOPに関する話
S : 対象のアクターは1つ。適切なドメインモデルに分割する。
O : 修正するとき、追加のみで要件を満たせる。
L : ?
I : 
D : レイヤー抽象に依存させる。

- SOLID原則の観点から、委譲、継承を見る。

- 何がどの原則を満たしているのか。

- 継承の問題点と、goではどのようにis-a関係を表現しているのか
