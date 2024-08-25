package store

type StoreModule struct {
	repository Repository
	Create useCaseCreate
	Update useCaseUpdate
	AddBusinessHour useCaseAddBusinessHour
	DeleteBusinessHour useCaseDeleteBusinessHour
	GetByID useCaseGetByID
	GetByFilter useCaseGetByFilter
}

func NewStoreModule(storeRepository Repository) StoreModule {
	return StoreModule{
		repository: storeRepository,
		Create: newUseCaseCreate(storeRepository),
		Update: newUseCaseUpdate(storeRepository),
		AddBusinessHour: newUseCaseAddBusinessHour(storeRepository),
		DeleteBusinessHour: newUseCaseDeleteBusinessHour(storeRepository),
		GetByID: newUseCaseGetByID(storeRepository),
		GetByFilter: newUseCaseGetByFilter(storeRepository),
	}
}