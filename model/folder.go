package model

import "time"

type Folder struct {
	Id         int64     `gorm:"primary_key" json:"id"`                                     // ID
	Files      []*File   `json:"files"`                                                     // many2many
	Folders    []*Folder `gorm:"foreignkey:ParentId" json:"folders"`                        // one2many 当前目录下的目录
	UserId     int64     `gorm:"index:user_id_folder_name_unique_index" json:"user_id"`     // 创建者(组合唯一)
	FolderName string    `gorm:"index:user_id_folder_name_unique_index" json:"folder_name"` // 目录名称(组合唯一)
	ParentId   int64     `gorm:"default:0" json:"parent_id"`                                // 父目录
	Key        string    `gorm:"default:''" json:"key"`                                     // 辅助键
	Level      int64     `gorm:"default:1" json:"level"`                                    // 辅助键
	CreatedAt  time.Time `json:"created_at"`                                                // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`                                                // 更新时间
}

const (
	FolderKeyPrefix = "-"
)

type FolderStore interface {
	// 创建一个目录
	CreateFolder(folder *Folder) (err error)
	// 目录是否存在
	ExistFolder(userId int64, folderName string) (isExist bool)
	// 当 id != 0 则表示加载指定目录, 当 id == 0 则表示加载根目录
	LoadFolder(id, userId int64, isLoadRelated bool) (folder *Folder, err error)
	// 删除指定目录
	DeleteFolder(ids []int64, userId int64) (err error)
	// 移动目录
	MoveFolder(to *Folder, ids []int64) (err error)
	// 复制目录
	CopyFolder(to *Folder, ids []int64) (err error)
	// 重命名目录
	RenameFolder(id int64, newName string) (err error)
}

type FolderService interface {
	FolderStore
}
