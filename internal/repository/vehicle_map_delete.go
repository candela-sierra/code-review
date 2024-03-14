package repository

func (r *VehicleMap) Delete(id int) (err error) {
	_, err = r.FindById(id)
	if err != nil {
		return
	}
	delete(r.db, id)
	return
}
