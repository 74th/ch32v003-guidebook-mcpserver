# CH32V003 Guidebook MCP Server

@74th が執筆した技術同人誌『CH32V003 Guidebook』を利用したCH32V003ファームウェア開発用のMCP Serverプログラムです。

技術同人誌は、2025年5月31日スタートの技術書典18にて頒布します！

> CH32V003開発ガイドブック：74th
>
> https://techbookfest.org/product/qH8sBZPnJAREuL45aDSfpa

本リポジトリにはGitサブモジュールで、技術同人誌の本文及び抽出ドキュメントを参照していますが、こちらは書籍購入者向けとしてオープンソースにはしておりません。
技術同人誌の内容を含むMCP Serverを、フリーソフトウェアとして配布しています。

## 概要

ツール`read_ch32v003_guide_book`を提供します。

引数のframeworkとitemを指定することで、同人誌のコンテンツと、サンプルコードをMCP Serverに与えます。

- framework:
  - WCH SDK
  - ch32fun
- item:
  - GPIO
  - Timer
  - DMA
  - ADC
  - PWM
  - I2C
  - SPI
  - UART
  - watchdogtimer

frameworkを指定する必要があるため、LLMにframeworkを指示しておいてください。

## 利用方法

[リリースページ](https://github.com/74th/ch32v003-guidebook-mcpserver/releases)からダウンロードして、PC内に設置してください。

### VS Codeで、GitHub Copilotを利用する場合

#### MCP Serverの有効化

VS Codeで拡張機能GitHub Copilotをインストールしてください。

リポジトリにて`.vscode/mcp.json`を編集、もしくは作成して、以下のように、設置したMCP Serverのパスを指定してください。

```json
{
	"servers": {
		"guidebook": {
			"command": "/Users/nnyn/ghq/github.com/74th/ch32v003-guidebook-mcpserver/ch32v003-guidebook-mcpserver",
			"args": [],
			"env": {}
		}
	}
}
```

以下のように、GitHub Copilot Chatを開き、Agent Modeを選択した上で、Toolを有効化してください。

![enable copilot](./docs/using_mcp_server_with_github_copilot.drawio.svg)

Toolに表示されない場合、MCP Serverの再起動も試みてください。

Chatのプロンプトにて、以下のようにTool`read_ch32v003_guide_book`を使うように指示して利用してください。
この時フレームワーク名を`WCH SDK`か`ch32fun`を教えておく必要があります。
以下はプロンプトの例です。

```
CH32V003マイコンのファームウェアの開発です
フレームワークはch32funを使います
ch32funでの実装方法は、Tool read_ch32v003_guide_bookを参照してください。

UARTを使う準備をして。ボーレート 115200で1バイト読み書きするコードも用意して。
```

#### MCP ServerをGitHub Copilotに使わせるようにカスタムプロンプトの指定

毎回Toolを指定する指示をしなくても、自動でプロンプトに追加できます。

リポジトリにて`.vscode/settings.json`を編集、もしくは作成して、`"chat.promptFiles": true`を追加してください。

```json
{
  "chat.promptFiles": true
}
```

`.github/copilot-instructions.md`に以下のように記述してください。

```markdown
CH32V003マイコンのファームウェアの開発です
フレームワークはch32funを使います
ch32funでの実装方法は、Tool read_ch32v003_guide_bookを参照してください。
```

### サンプルのプロジェクトファイル

サンプルのプロジェクトフォルダを用意してあります。

#### ch32fun

[./sample_project_directory/ch32fun](./sample_project_directory/ch32fun/)

ch32fun/examples/templateに上記設定を追加したものです。

Makefile、.vscode/c_cpp_properties.json、.vscode/mcp.json中の、ch32fun及びch32v003-guidebook-mcpserverのパスを書き換えて利用してください。

## 生成サンプル

プロンプトのサンプル例です。

[./example_prompt.md](./example_prompt.md)

## ご注意

このMCP Serverを利用することで、完全に動作するCH32V003のファームウェアを作成を約束はしません。

実験的なプロダクトです。
期待通り動作しなくてもご容赦ください。

## LICENSE

本プログラムには以下の3種類の著作物が含まれます。

1. ソフトウェア部分 … 本リポジトリに含まれるコード
2. 同人誌コンテンツ部分 …… プログラム内部に格納された、また同人誌購入者向けに提供される
3. サンプルコード部分 …… プログラム内部に格納された、または公開されているサンプルコード

1.はMITライセンスが適用されます。原文は[./LICENSEのSECTION I](./LICENSE)を参照してください。

2.は個人利用のみのライセンスです。詳しくは[./LICENSEのSECTION II](./LICENSE)を参照してください。

3.はCC0ライセンスが適用されます。

MCP Serverのプログラム自体には、2.のライセンスが含まれます。