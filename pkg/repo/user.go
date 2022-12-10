package repo

//
//import "context"
//
//package repo
//
//import (
//"context"
//"iot/pkg/model"
//"iot/pkg/utils"
//)
//
//func (r *RepoPG) CreateUser(ctx context.Context, ob *model.User) error {
//	tx, cancel := r.DBWithTimeout(ctx)
//	defer cancel()
//	return tx.Create(ob).Error
//}
//
//func (r *RepoPG) UpdateUser(ctx context.Context, ob *model.User) error {
//	tx, cancel := r.DBWithTimeout(ctx)
//	defer cancel()
//	return tx.Where("id = ?", ob.ID).Save(ob).Error
//}
//
//func (r *RepoPG) DeleteUser(ctx context.Context, id string) error {
//	tx, cancel := r.DBWithTimeout(ctx)
//	defer cancel()
//	return tx.Where("id = ?", id).Delete(&model.User{}).Error
//}
//
//func (r *RepoPG) GetOneUser(ctx context.Context, id string) (*model.User, error) {
//	tx, cancel := r.DBWithTimeout(ctx)
//	defer cancel()
//
//	rs := model.User{}
//	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
//		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
//	}
//
//	return nil, nil
//}
//
