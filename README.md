# MH-API (モンスターハンター API)

[![lint and test](https://github.com/o-ga09/MH-API/actions/workflows/lint_and_test.yml/badge.svg)](https://github.com/o-ga09/MH-API/actions/workflows/lint_and_test.yml)

English version is [here](./README_EN.md)

MH-API は、モンスターハンターシリーズに関する攻略情報等を提供するオープンソースのプロジェクトです。このプロジェクトでは、モンスターハンターシリーズのプレイヤーがゲームの攻略情報等にアクセスし、二次創作やツール等の制作を簡単にできるようにするための API を開発しています。

## はじめに

この README.md は、MH-API プロジェクトのガイドラインと使い方について説明します。以下のガイドラインに従ってプロジェクトに参加していただくことを歓迎します。

本プロジェクトにおける行動規範は[こちら](./CODE_OF_CONDUCT_JA.md)

## Getting Started

MH-API プロジェクトに参加するためには、以下の手順に従ってください。

### 事前準備

- コントリビュートガイドを確認する。[コントリビュートガイドはこちら](./CONTRIBUTING_JA.md)

### 環境構築

1. リポジトリのディレクトリに移動

   ```bash
   cd MH-API
   ```

2. エディタでディレクトリを開く
3. 新規ブランチを作成する

   ```bash
       git checkout -b "[new branch]"
   ```

4. 動作確認する

   ```bash
        # dockerを立ち上げる
        make up

        # {"message": "ok"}とレスポンスが返ること
        curl http://localhost:8080/v1/system/health
   ```

5. テストを実行する

   ```bash
       # 全てのテストがokとなること
       make test
   ```

## コミュニティとコミュニケーション

本コミュニティでは[こちら](./CODE_OF_CONDUCT_JA.md)の行動規範に従ってください。

MH-API プロジェクトに参加するには、以下のコミュニケーションチャンネルを利用できます。

- Slack チャンネル: [slack.mhapi.org](https://mh-api.slack.com) に参加して、他のコントリビューターやメンバーと交流しましょう。

- Issue トラッカー: [MH-API Issue トラッカー](https://github.com/o-ga09/MH-API/issues) を使用して、バグ報告や新しい機能の提案を行ってください。

- メーリングリスト: [mhapi-dev@groups.com](mailto:mhapiadm@gmail.com) に参加して、メールを通じたディスカッションや重要なお知らせを受け取りましょう。

## ライセンス

MH-API プロジェクトは、[MIT ライセンス](https://opensource.org/licenses/MIT)のもとで公開されています。詳細なライセンス情報については、プロジェクト内の LICENSE ファイルをご確認ください。

## 貢献ガイドライン

MH-API プロジェクトへの貢献に関するガイドラインについては、[CONTRIBUTING.md](./CONTRIBUTING_JA.md)を参照してください。プロジェクトにコードやドキュメントの貢献をする前に、ガイドラインをお読みください。

## サポート

MH-API プロジェクトに関するサポートが必要な場合は、[mhapiadm@gmail.com](mailto:mhapiadm@gmail.com)までお問い合わせください。

## API Spec

[API doc](https://o-ga09.github.io/MH-API/spec.html)

## 謝辞

このプロジェクトは、オープンソースコミュニティの貢献者や MH-API ユーザーのご協力によって成り立っています。多くの人々に感謝の意を表します。

<!-- プロジェクトの詳細や最新情報については、[MH-API公式ウェブサイト](https://mhapi.org)をご覧ください。 -->

このプロジェクトは、株式会社カプコンの商標および登録商標であるモンスターハンターシリーズ ™ を利用しています。モンスターハンターシリーズ ™ は株式会社カプコンの知的財産です。ここにカプコン様へ感謝の意を表します。

なお、このプロジェクトは非公式のものであり、株式会社カプコンとは関係ありません。

**Happy coding!**

### References

<https://opensource.guide/ja/starting-a-project/>

this project has started from 2023/5/21
