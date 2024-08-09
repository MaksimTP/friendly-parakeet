package cache

import (
	"main/db"
	"main/model"
	"sync"
)

type Cache struct {
	cachedData map[int]model.Order
	m          sync.RWMutex
}

func (c *Cache) SaveData(data model.Order) {
	c.m.RLock()
	defer c.m.RUnlock()
	c.cachedData[data.Id] = data
}

func (c *Cache) GetOrderById(i int) model.Order {
	c.m.RLock()
	defer c.m.RUnlock()
	if order, found := c.cachedData[i]; found {
		return order
	}
	return model.Order{}
}

func (c *Cache) RestoreDataFromDB(d db.DataBase) {
	var data []model.Order = d.GetAllData()
	for _, v := range data {
		c.SaveData(v)
	}
}
