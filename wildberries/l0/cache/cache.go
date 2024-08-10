package cache

import (
	"main/model"
	"sync"
)

type Cache struct {
	cachedData map[string]model.Order
	m          sync.RWMutex
}

func (c *Cache) SaveData(data model.Order) {
	c.m.RLock()
	defer c.m.RUnlock()
	c.cachedData[data.OrderUid] = data
}

func (c *Cache) GetOrderById(uid string) model.Order {
	c.m.RLock()
	defer c.m.RUnlock()
	if order, found := c.cachedData[uid]; found {
		return order
	}
	return model.Order{}
}

func NewCache() *Cache {
	return &Cache{make(map[string]model.Order), sync.RWMutex{}}
}

// func (c *Cache) RestoreDataFromDB(d db.DataBase) {
// 	var data []model.Order = d.GetAllData()
// 	for _, v := range data {
// 		c.SaveData(v)
// 	}
// }
