package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Desc   string    `gorm:"size:255;not null;" json:"desc"`
}

func (p *Vote) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Desc = html.EscapeString(strings.TrimSpace(p.Desc))
}

func (p *Vote) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Desc == "" {
		return errors.New("Required Description")
	}
	return nil
}

func (p *Vote) SaveVote(db *gorm.DB) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Create(&p).Error
	if err != nil {
		return &Vote{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

func (p *Vote) FindAllVotes(db *gorm.DB) (*[]Vote, error) {
	var err error
	votes := []Vote{}
	err = db.Debug().Model(&Vote{}).Limit(100).Find(&votes).Error
	if err != nil {
		return &[]Vote{}, err
	}
	if len(votes) > 0 {
		for i, _ := range votes {
			if err != nil {
				return &[]Vote{}, err
			}
			i++
		}
	}
	return &votes, nil
}

func (p *Vote) FindVoteByID(db *gorm.DB, pid uint64) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Vote{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

func (p *Vote) UpdateAVote(db *gorm.DB) (*Vote, error) {

	var err error

	if err != nil {
		return &Vote{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

func (p *Vote) DeleteAVote(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Vote{}).Where("id = ? and author_id = ?", pid, uid).Take(&Vote{}).Delete(&Vote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Vote not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}