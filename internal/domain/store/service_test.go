package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/domain/address"
	"github.com/oprimogus/cardapiogo/internal/domain/owner"
	"github.com/oprimogus/cardapiogo/internal/domain/store"
	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/errors"
	mock_owner "github.com/oprimogus/cardapiogo/internal/infra/mocks/owner"
	mock_store "github.com/oprimogus/cardapiogo/internal/infra/mocks/store"
)

type ServiceSuite struct {
	suite.Suite
	controller      *gomock.Controller
	repository      *mock_store.MockRepository
	ownerRepository *mock_owner.MockRepository
	Service         *store.Service
	ownerService    *owner.Service
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TearDownTest(t *testing.T) {
	defer s.controller.Finish()
}

func (s *ServiceSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.repository = mock_store.NewMockRepository(s.controller)
	s.ownerRepository = mock_owner.NewMockRepository(s.controller)
	s.ownerService = owner.NewService(s.ownerRepository)
	s.Service = store.NewService(s.repository, *s.ownerService)
}

func (s *ServiceSuite) TestVerifyStoreExist() {
	type storesTestType struct {
		name      string
		store     store.StoreDetail
		err       error
		wantBool  bool
		wantError error
	}
	storesTest := []storesTestType{
		{"Store exists", store.StoreDetail{Name: "Test 1"}, nil, true, nil},
		{
			"Store doesnt't exist",
			store.StoreDetail{},
			fmt.Errorf(errors.NOT_FOUND_RECORD),
			false,
			nil,
		},
	}

	for _, v := range storesTest {
		s.repository.EXPECT().
			GetStoreByID(gomock.Any(), gomock.Any()).
			Return(v.store, v.err).
			Times(1)
		gotBool, err := store.StoreWithIDExist(context.Background(), s.Service, "")
		assert.Equal(s.T(), gotBool, v.wantBool, v.name)
		assert.Equal(s.T(), err, v.wantError, v.name)
	}
}

func (s *ServiceSuite) TestVerifyStoreExistWithCpfCnpj() {
	type storesTestType struct {
		name      string
		store     store.StoreDetail
		err       error
		wantBool  bool
		wantError error
	}
	storesTest := []storesTestType{
		{"Store exists", store.StoreDetail{Name: "Test 1"}, nil, true, nil},
		{
			"Store doesnt't exist",
			store.StoreDetail{},
			fmt.Errorf(errors.NOT_FOUND_RECORD),
			false,
			nil,
		},
	}

	for _, v := range storesTest {
		s.repository.EXPECT().
			GetStoreByCpfCnpj(gomock.Any(), "").
			Return(v.store, v.err).
			Times(1)
		gotBool, err := store.StoreWithCpfCnpjExist(context.Background(), s.Service, "")
		assert.Equal(s.T(), gotBool, v.wantBool, v.name)
		assert.Equal(s.T(), err, v.wantError, v.name)
	}
}

func (s *ServiceSuite) TestCreateStore() {
	type CreateStoresData struct {
		name                       string
		input                      store.CreateStoreParams
		mockGetStoreByCpfCnpjError error
		mockCreateStore            error
		want                       error
	}

	data := []CreateStoresData{
		{
			name: "Doesn't exist store yet",
			input: store.CreateStoreParams{
				Name:      "Test1",
				CpfCnpj:   "",
				Phone:     "",
				Address:   address.CreateAddressParams{},
				StoreType: types.StoreShopRestaurant,
			},
			mockGetStoreByCpfCnpjError: fmt.Errorf(errors.NOT_FOUND_RECORD),
			mockCreateStore:            nil,
			want:                       nil,
		},
		{
			name: "Exist store",
			input: store.CreateStoreParams{
				Name:      "Test2",
				CpfCnpj:   "",
				Phone:     "",
				Address:   address.CreateAddressParams{},
				StoreType: types.StoreShopRestaurant,
			},
			mockGetStoreByCpfCnpjError: nil,
			mockCreateStore:            nil,
			want:                       errors.ConflictError(store.ExistRegisteredStore),
		},
	}

	for _, v := range data {
		s.repository.EXPECT().
			GetStoreByCpfCnpj(gomock.Any(), "").
			Return(store.StoreDetail{}, v.mockGetStoreByCpfCnpjError)

		s.repository.EXPECT().
			CreateStore(gomock.Any(), gomock.Any()).
			Return(v.mockCreateStore).
			AnyTimes()

		err := s.Service.CreateStore(context.Background(), v.input)
		assert.Equal(s.T(), v.want, err, v.name)
	}
}

func (s *ServiceSuite) TestUpdateStore() {
	type UpdateStoresData struct {
		name                  string
		input                 store.UpdateStoreParams
		mockGetStoreByIDError error
		mockIsOwnerValue      bool
		mockIsOwnerError      error
		mockIsOwnerTimes      int
		mockUpdateStore       error
		mockUpdateStoreTimes  int
		want                  error
	}

	data := []UpdateStoresData{
		{
			name: "Doesn't exist store yet",
			input: store.UpdateStoreParams{
				ExternalID: "",
				Name:       "Test1",
				CpfCnpj:    "",
				Phone:      "",
				Address:    address.CreateAddressParams{},
				StoreType:  types.StoreShopRestaurant,
			},
			mockGetStoreByIDError: fmt.Errorf(errors.NOT_FOUND_RECORD),
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      0,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  0,
			want:                  errors.NotFound(store.NotRegisteredStore),
		},
		{
			name: "Store exist and is owner doing request",
			input: store.UpdateStoreParams{
				ExternalID: "",
				Name:       "Test2",
				CpfCnpj:    "",
				Phone:      "",
				Address:    address.CreateAddressParams{},
				StoreType:  types.StoreShopRestaurant,
			},
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      true,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  1,
			want:                  nil,
		},
		{
			name: "Store exist and not is owner doing request",
			input: store.UpdateStoreParams{
				ExternalID: "",
				Name:       "Test3",
				CpfCnpj:    "",
				Phone:      "",
				Address:    address.CreateAddressParams{},
				StoreType:  types.StoreShopRestaurant,
			},
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  0,
			want:                  errors.Forbidden(store.IsNotOwner),
		},
	}

	for _, v := range data {
		s.repository.EXPECT().
			GetStoreByID(gomock.Any(), "").
			Return(store.StoreDetail{}, v.mockGetStoreByIDError)

		s.ownerRepository.EXPECT().
			IsOwner(gomock.Any(), "", gomock.Any()).
			Return(v.mockIsOwnerValue, v.mockIsOwnerError).Times(v.mockIsOwnerTimes)

		s.repository.EXPECT().
			UpdateStore(gomock.Any(), gomock.Any()).
			Return(v.mockUpdateStore).Times(v.mockUpdateStoreTimes)

		err := s.Service.UpdateStore(context.Background(), v.input, "")
		assert.Equal(s.T(), v.want, err, v.name)
	}
}

func (s *ServiceSuite) TestUpdateStoreCpfCnpj() {
	type UpdateStoresData struct {
		name                  string
		input                 store.UpdateStoreCpfCnpjParams
		mockGetStoreByIDError error
		mockIsOwnerValue      bool
		mockIsOwnerError      error
		mockIsOwnerTimes      int
		mockUpdateStore       error
		mockUpdateStoreTimes  int
		want                  error
	}

	data := []UpdateStoresData{
		{
			name: "Doesn't exist store yet",
			input: store.UpdateStoreCpfCnpjParams{
				ExternalID: "",
				CpfCnpj:    "",
			},
			mockGetStoreByIDError: fmt.Errorf(errors.NOT_FOUND_RECORD),
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      0,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  0,
			want:                  errors.NotFound(store.NotRegisteredStore),
		},
		{
			name: "Store exist and is owner doing request",
			input: store.UpdateStoreCpfCnpjParams{
				ExternalID: "",
				CpfCnpj:    "",
			},
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      true,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  1,
			want:                  nil,
		},
		{
			name: "Store exist and not is owner doing request",
			input: store.UpdateStoreCpfCnpjParams{
				ExternalID: "",
				CpfCnpj:    "",
			},
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockUpdateStore:       nil,
			mockUpdateStoreTimes:  0,
			want:                  errors.Forbidden(store.IsNotOwner),
		},
	}

	for _, v := range data {
		s.repository.EXPECT().
			GetStoreByID(gomock.Any(), "").
			Return(store.StoreDetail{}, v.mockGetStoreByIDError)

		s.ownerRepository.EXPECT().
			IsOwner(gomock.Any(), "", gomock.Any()).
			Return(v.mockIsOwnerValue, v.mockIsOwnerError).Times(v.mockIsOwnerTimes)

		s.repository.EXPECT().
			UpdateStoreCpfCnpj(gomock.Any(), gomock.Any()).
			Return(v.mockUpdateStore).Times(v.mockUpdateStoreTimes)

		err := s.Service.UpdateStoreCpfCnpj(context.Background(), v.input, "")
		assert.Equal(s.T(), v.want, err, v.name)
	}
}

func (s *ServiceSuite) TestDeleteStore() {
	type DeleteStoresData struct {
		name                  string
		ID                    string
		userID                string
		mockGetStoreByIDError error
		mockIsOwnerValue      bool
		mockIsOwnerError      error
		mockIsOwnerTimes      int
		mockDeleteStore       error
		mockDeleteStoreTimes  int
		want                  error
	}

	data := []DeleteStoresData{
		{
			name:                  "Doesn't exist store yet",
			ID:                    "",
			userID:                "",
			mockGetStoreByIDError: fmt.Errorf(errors.NOT_FOUND_RECORD),
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      0,
			mockDeleteStore:       nil,
			mockDeleteStoreTimes:  0,
			want:                  errors.NotFound(store.NotRegisteredStore),
		},
		{
			name:                  "Store exist and is owner doing request",
			ID:                    "",
			userID:                "",
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      true,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockDeleteStore:       nil,
			mockDeleteStoreTimes:  1,
			want:                  nil,
		},
		{
			name:                  "Store exist and not is owner doing request",
			ID:                    "",
			userID:                "",
			mockGetStoreByIDError: nil,
			mockIsOwnerValue:      false,
			mockIsOwnerError:      nil,
			mockIsOwnerTimes:      1,
			mockDeleteStore:       nil,
			mockDeleteStoreTimes:  0,
			want:                  errors.Forbidden(store.IsNotOwner),
		},
	}

	for _, v := range data {
		s.repository.EXPECT().
			GetStoreByID(gomock.Any(), "").
			Return(store.StoreDetail{}, v.mockGetStoreByIDError)

		s.ownerRepository.EXPECT().
			IsOwner(gomock.Any(), "", gomock.Any()).
			Return(v.mockIsOwnerValue, v.mockIsOwnerError).Times(v.mockIsOwnerTimes)

		s.repository.EXPECT().
			DeleteStore(gomock.Any(), gomock.Any()).
			Return(v.mockDeleteStore).Times(v.mockDeleteStoreTimes)

		err := s.Service.DeleteStore(context.Background(), "", "")
		assert.Equal(s.T(), v.want, err, v.name)
	}
}
