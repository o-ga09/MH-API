package batch

import (
	"context"
	"fmt"
	"log/slog"

	fieldDomain "mh-api/app/internal/domain/fields"
	itemDomain "mh-api/app/internal/domain/items"
	monsterDomain "mh-api/app/internal/domain/monsters"
	musicDomain "mh-api/app/internal/domain/music"
	partDomain "mh-api/app/internal/domain/part"
	productDomain "mh-api/app/internal/domain/products"
	rankingDomain "mh-api/app/internal/domain/ranking"
	tribeDomain "mh-api/app/internal/domain/tribes"
	weaknessDomain "mh-api/app/internal/domain/weakness"
	weaponDomain "mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/presenter/middleware"

	fieldDriver "mh-api/app/internal/driver/fields"
	itemDriver "mh-api/app/internal/driver/items"
	monsterDriver "mh-api/app/internal/driver/monsters"
	musicDriver "mh-api/app/internal/driver/music"

	"mh-api/app/internal/driver/mysql"
	partDriver "mh-api/app/internal/driver/parts"
	productDriver "mh-api/app/internal/driver/products"
	rankingDriver "mh-api/app/internal/driver/ranking"
	tribeDriver "mh-api/app/internal/driver/tribes"
	weaknessDriver "mh-api/app/internal/driver/weakness"
	weaponDriver "mh-api/app/internal/driver/weapons"
)

type BatchService struct {
	monsterService  monsterDomain.Repository
	itemService     itemDomain.Repository
	fieldService    fieldDomain.Repository
	musicService    musicDomain.Repository
	partService     partDomain.Repository
	productService  productDomain.Repository
	tribeService    tribeDomain.Repository
	rankingService  rankingDomain.Repository
	weaknessService weaknessDomain.Repository
	weaponService   weaponDomain.Repository
}

func NewBatchService(
	monsterService monsterDomain.Repository,
	itemService itemDomain.Repository,
	fieldService fieldDomain.Repository,
	musicService musicDomain.Repository,
	partService partDomain.Repository,
	productService productDomain.Repository,
	tribeService tribeDomain.Repository,
	rankingRepo rankingDomain.Repository,
	weaknessService weaknessDomain.Repository,
	weaponService weaponDomain.Repository,
) *BatchService {
	return &BatchService{
		monsterService:  monsterService,
		itemService:     itemService,
		fieldService:    fieldService,
		musicService:    musicService,
		partService:     partService,
		productService:  productService,
		tribeService:    tribeService,
		rankingService:  rankingRepo,
		weaknessService: weaknessService,
		weaponService:   weaponService,
	}
}

func Exec(ctx context.Context, batchName string) error {
	batchService := BatchDI()

	switch batchName {
	case "createMonster":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
		err := Create(ctx, batchService)
		if err != nil {
			return err
		}
	case "removeMonster":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "createItem":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "removeItem":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "createWeapon":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "removeWeapon":
		slog.Log(ctx, middleware.SeverityInfo, fmt.Sprintf("[Batch Name]: %s", batchName))
	}

	return nil
}

func BatchDI() *BatchService {
	db := mysql.New(context.Background())
	Monsterrepo := monsterDriver.NewMonsterRepository(db)
	Fieldrepo := fieldDriver.NewfieldRepository(db)
	Musicrepo := musicDriver.NewmusicRepository(db)
	Itemrepo := itemDriver.NewMonsterRepository(db)
	Partrepo := partDriver.NewMonsterRepository(db)
	Productrepo := productDriver.NewMonsterRepository(db)
	Rankingrepo := rankingDriver.NewMonsterRepository(db)
	Triberepo := tribeDriver.NewtribeRepository(db)
	Weaknessrepo := weaknessDriver.NewweakRepository(db)
	Weaponrepo := weaponDriver.NewweaponRepository(db)

	return NewBatchService(Monsterrepo, Itemrepo, Fieldrepo, Musicrepo, Partrepo, Productrepo, Triberepo, Rankingrepo, Weaknessrepo, Weaponrepo)
}
