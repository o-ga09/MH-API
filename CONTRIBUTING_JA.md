# コントリビュートガイド

バグフィックス、データ提供、推奨事項など、すべての貢献を歓迎します。

プルリクエストを提出したり問題を提起したりする前に、[issue on GitHub](https://github.com/o-ga09/MH-API/issues) を見てください。他の人に先を越されたかもしれません。

このリポジトリに貢献するには

- プロジェクトを自分の GitHub プロファイルに[フォークする](https://help.github.com/articles/fork-a-repo/)

- git clone を使ってプロジェクトをダウンロードします

```bash
    git clone git@github.com:<YOUR_USERNAME>/MH-API.git
```

- 新しいブランチを作成し、わかりやすい名前を付けます：

```bash
    git checkout -b my_new_branch
```

- コードを書き、何かを修正し、それが動作することを証明するためにテストを追加します。**テストに合格しない限り、あるいは新しい機能が追加された場合に新しいテストがない限り、プルリクエストは受理されません**。

- テストを実行する

```bash
    # 全てのテストがokとなること
    make test
```

- コードをコミットし、GitHubにプッシュします。
  - このとき、コミットメッセージにissue番号を付けること

- [Open a new pull request](https://help.github.com/articles/creating-a-pull-request/) と、あなたが行った変更を記述してください。

- レビュー後、あなたの変更を受け入れます。

## 質問

質問がある場合は、[issue](issue)を作成してください(プロヒント: 同じ質問をした人がいないか、まず検索してみてください!)。
また、@o-ga09でも受け付けています。

### 開発者の方へ

本プロジェクト(MH-API)をご覧いただき、ありがとうございます。ぜひ、プロジェクトに貢献しませんか。
