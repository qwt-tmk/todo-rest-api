package user

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
	"github.com/qwt-tmk/todo-rest-api/domain/user"
	"go.uber.org/mock/gomock"
)

func TestUser_UpdateProfileUsecase_Run(t *testing.T) {
	tests := []struct {
		name string
		mockFn func(mr *user.MockUserRepository)
		input UpdateProfileUsecaseInputDTO
		want *UpdateProfileUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系：プロフィールを編集できる",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID: "1",
				Email: "updated@test.com",
				Name: "updatedUser",
			},
			want: &UpdateProfileUsecaseOutputDTO{
				ID: "1",
				Email: "updated@test.com",
				Name: "updatedUser",
			},
		},
		{
			name: "正常系：空のフィールドは更新されない",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID: "1",
				Name: "updatedUser",
			},
			want: &UpdateProfileUsecaseOutputDTO{
				ID: "1",
				Email: "user@test.com",
				Name: "updatedUser",
			},
			wantErr: false,
		},
		{
			name: "凖正常系：存在しないユーザーは編集できない",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.ErrNotFoundUser)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID: "0",
				Email: "noexistent@test.com",
				Name: "noexistent",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUserRepository := user.NewMockUserRepository(ctrl)
			tt.mockFn(mockUserRepository)
			UpdateProfileUsecase := NewUpdateProfileUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := UpdateProfileUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfileUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UpdateProfileUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
