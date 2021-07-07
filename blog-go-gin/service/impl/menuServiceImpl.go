package impl

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"sort"
	"sync"
)

type MenuServiceImpl struct {
	wg sync.WaitGroup
}

func (m *MenuServiceImpl) GetUserMenus(roleid int) ([]*pb.ScUserMenuMessage, error) {
	userMenus, err := model.GetUserMenus(roleid)
	if err != nil {
		return nil, err
	}
	catalog := make([]*model.Menu, 0)
	subMenu := make([]*model.Menu, 0)
	for _, menu := range userMenus {
		if menu.ParentID == 0 {
			catalog = append(catalog, menu)
		} else {
			subMenu = append(subMenu, menu)
		}
	}
	//TODO 这里可以并行处理
	catalogList := m.SortByOrderNum(catalog)
	//获取目录下的子菜单
	logging.Logger.Debug(catalogList)
	subMenuMap := m.GetSubMenu(subMenu)
	logging.Logger.Debug(subMenuMap)
	//转换前端菜单格式
	return m.ConvertUserMenuList(catalogList, subMenuMap), nil

}

func (m *MenuServiceImpl) SortByOrderNum(menus []*model.Menu) []*model.Menu {
	sort.Slice(menus, func(i, j int) bool {
		if menus[i].OrderNum != menus[j].OrderNum {
			return menus[i].OrderNum < menus[j].OrderNum
		}
		return menus[i].Name < menus[j].Name
	})
	return menus
}

func (m *MenuServiceImpl) GetSubMenu(menus []*model.Menu) map[int][]*model.Menu {
	logging.Logger.Debug(menus)
	subMenuMap := make(map[int][]*model.Menu)
	for _, menu := range menus {
		if subMenuMap[menu.ParentID] == nil {
			subMenu := make([]*model.Menu, 0)
			subMenu = append(subMenu, menu)
			subMenuMap[menu.ParentID] = subMenu
		} else {
			subMenuMap[menu.ParentID] = append(subMenuMap[menu.ParentID], menu)
		}
	}
	return subMenuMap
}

func (m *MenuServiceImpl) ConvertUserMenuList(catalogList []*model.Menu, subMenuMap map[int][]*model.Menu) []*pb.ScUserMenuMessage {
	// 获取目录下的子菜单
	var scUserMenuList []*pb.ScUserMenuMessage
	for _, menu := range catalogList {
		var scUserMenuMessage *pb.ScUserMenuMessage
		var childrenList []*pb.ScUserMenuMessage
		subMenus, ok := subMenuMap[menu.ID]
		if ok {
			//多级菜单处理
			scUserMenuMessage = &pb.ScUserMenuMessage{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Icon:      menu.Icon,
				IsHidden:  menu.IsHidden == 1,
			}
			subMenus = m.SortByOrderNum(subMenus)

			for _, subMenu := range subMenus {
				childrenList = append(childrenList, &pb.ScUserMenuMessage{
					Name:      subMenu.Name,
					Path:      subMenu.Path,
					Component: subMenu.Component,
					Icon:      subMenu.Icon,
					IsHidden:  subMenu.IsHidden == 1,
				})
			}
		} else {
			//一级菜单处理
			scUserMenuMessage = &pb.ScUserMenuMessage{
				Path:      menu.Path,
				Component: common.Component,
			}
			childrenList = append(childrenList, &pb.ScUserMenuMessage{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Icon:      menu.Icon,
				IsHidden:  menu.IsHidden == 1,
			})
		}
		scUserMenuMessage.IsHidden = menu.IsHidden == 1
		scUserMenuMessage.Children = childrenList
		scUserMenuList = append(scUserMenuList, scUserMenuMessage)
	}
	return scUserMenuList
}

func NewMenuServiceImpl() *MenuServiceImpl {
	return &MenuServiceImpl{}
}
