package items

import (
	"context"
	"mh-api/internal/domain/items" // ドメイン層のitemsパッケージをインポート
)

// ItemDTO はAPIレスポンスのためのアイテムデータ転送オブジェクトです。
type ItemDTO struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
}

// ItemListResponseDTO はアイテムリストAPIのレスポンス形式です。
type ItemListResponseDTO struct {
	Items []ItemDTO `json:"items"`
}

// Service はアイテムに関するビジネスロジックを提供するサービスインターフェースです。
// (今回は具象型のみ実装しますが、将来的な拡張性を考慮してインターフェースを定義しておくことも可能です)
type Service struct {
	itemRepo items.Repository // items.Repositoryインターフェースへの依存
}

// NewService は Service の新しいインスタンスを生成します。
func NewService(itemRepo items.Repository) *Service {
	return &Service{
		itemRepo: itemRepo,
	}
}

// GetAllItems は全てのアイテム情報を取得し、DTOに変換して返します。
func (s *Service) GetAllItems(ctx context.Context) (*ItemListResponseDTO, error) {
	domainItems, err := s.itemRepo.FindAll(ctx)
	if err != nil {
		// TODO: エラーの種類に応じたハンドリング (例: RecordNotFoundなど)
		return nil, err // エラーをそのまま返すか、サービス層独自のエラーに変換するか検討
	}

	var itemDTOs []ItemDTO
	for _, domainItem := range domainItems {
		itemDTOs = append(itemDTOs, ItemDTO{
			ItemID:   domainItem.GetID(),   // ドメインオブジェクトのメソッド経由で値を取得
			ItemName: domainItem.GetName(), // ドメインオブジェクトのメソッド経由で値を取得
		})
	}

	return &ItemListResponseDTO{Items: itemDTOs}, nil
}
