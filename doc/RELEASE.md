# リリース管理ガイド

MH-API では [tagpr](https://github.com/Songmu/tagpr) を使用してリリースフローを自動化しています。

---

## 概要

tagpr は「**PR をマージするだけでリリースが完了する**」フローを実現するツールです。

- `main` ブランチに差分が積まれると、次のバージョンに向けたリリース PR を自動生成・更新
- リリース PR をマージすると、マージコミットへの **自動タグ付け** と **GitHub Releases の自動作成** が行われます

---

## リリースフロー

```
feature PR マージ → main 更新
        ↓
  tagpr が自動実行
        ↓
  リリース PR を作成 or 更新（タイトル例: "Release v1.2.2"）
        ↓
  リリースしたいタイミングでリリース PR をマージ
        ↓
  タグ付け + GitHub Releases 作成 完了 🎉
```

### 1. feature PR をマージする

通常の開発フロー（feature ブランチ → PR → `main` へマージ）を行います。

### 2. リリース PR が自動生成される

`main` への push をトリガーに `.github/workflows/tagpr.yml` が実行され、tagpr がリリース用 PR を自動生成します。

- PR には前回リリース以降のマージコミット一覧が記載されます
- `main` に追加のマージがあるたびに PR が自動更新されます

### 3. リリース PR をマージする

リリースしたいタイミングでリリース PR をマージするだけで完了です。

- `version.go` の `Version` 定数が新しいバージョンに自動更新されます
- マージコミットに `v1.x.x` 形式のタグが自動付与されます
- GitHub Releases が自動作成されます

---

## バージョニング規則

[Semantic Versioning (semver)](https://semver.org/lang/ja/) に従います。

| バージョン種別 | PR に付けるラベル | 例 |
|---|---|---|
| メジャー（破壊的変更） | `major` | v1.2.1 → v2.0.0 |
| マイナー（後方互換の機能追加） | `minor` | v1.2.1 → v1.3.0 |
| パッチ（バグ修正など） | ラベルなし（デフォルト） | v1.2.1 → v1.2.2 |

**ラベルはリリース PR 自体ではなく、そのリリースに含まれる各 feature PR に付与します。**  
tagpr が feature PR のラベルを自動集計し、リリース PR に適切なラベルを付与します。

---

## 設定ファイル

### `.tagpr`

```ini
[tagpr]
vPrefix = true          # タグに v プレフィックスを付ける（例: v1.2.3）
releaseBranch = main    # リリース対象ブランチ
versionFile = version.go  # バージョンを同期するファイル
release = true          # GitHub Releases を自動作成する
```

### `.github/workflows/tagpr.yml`

```yaml
on:
  push:
    branches: [main]   # main への push でトリガー

jobs:
  tagpr:
    permissions:
      contents: write       # タグ・コミットの作成
      pull-requests: write  # リリース PR の作成・更新
      issues: write         # ラベルの作成
```

### `version.go`

tagpr が自動更新するバージョン管理ファイルです。

```go
package mhapi

const Version = "v1.2.1"  // tagpr がリリース時に自動更新
```

---

## よくある質問

### Q. リリース PR をすぐにマージしたくない場合は？

リリース PR はそのまま放置しておけば問題ありません。`main` が更新されるたびに自動追従します。

### Q. リリースをスキップしたい場合は？

リリース PR をクローズするとスキップできます。ただし次に `main` が更新されたタイミングで再度 PR が作成されます。

### Q. 手動でリリースタグを打ちたい場合は？

通常の `git tag` + `git push` でタグを付与してください。tagpr は最新の semver タグを基準にするため、手動タグも正しく認識されます。

---

## 関連リンク

- [tagpr 公式リポジトリ](https://github.com/Songmu/tagpr)
- [Semantic Versioning](https://semver.org/lang/ja/)
- [GitHub Releases](https://github.com/o-ga09/MH-API/releases)
- [Actions ワークフロー](./../.github/workflows/tagpr.yml)
