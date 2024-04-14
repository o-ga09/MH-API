package parts

import (
	"context"
	"mh-api/app/internal/domain/part"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type partRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *partRepository {
	return &partRepository{
		conn: conn,
	}
}

func (r *partRepository) Get(ctx context.Context, monsterId string) (part.Parts, error) {
	p := []mysql.Part{}
	err := r.conn.Find(&p).Error
	if err != nil {
		return nil, err
	}

	res := part.Parts{}
	for _, r := range p {
		res = append(res, *part.NewPart(r.PartId, r.MonsterId, r.Name, r.Description))
	}

	return res, nil
}

func (r *partRepository) Save(ctx context.Context, p part.Part) error {
	data := mysql.Part{
		PartId:      p.GetID(),
		MonsterId:   p.GetMonsterID(),
		Name:        p.GetName(),
		Description: p.GetDescription(),
	}
	err := r.conn.Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *partRepository) Remove(ctx context.Context, partId string) error {
	data := mysql.Part{
		PartId: partId,
	}
	err := r.conn.Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
