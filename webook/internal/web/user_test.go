package web

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/service"
	svcmocks "7day/webook/internal/service/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser_Signup(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) service.UserSVC
		reqBody  string
		wantCode int
		result   Result
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) service.UserSVC {
				usersvc := svcmocks.NewMockUserSVC(ctrl)
				usersvc.EXPECT().Signup(gomock.Any(), domain.User{
					Email:    "test@test.com",
					Password: "123",
				}).Return(nil)
				return usersvc
			},
			reqBody: `
{
	"email": "test@test.com",
	"password": "123",
	"confirm_password": "123"
}`,
			wantCode: 200,
			result: Result{
				Code: 400,
				Msg:  "注册成功",
				Data: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			server := gin.Default()
			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRouter(server)
			req, err := http.NewRequest(http.MethodPost,
				"/users/signup", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			// 这就是 HTTP 请求进去 GIN 框架的入口。
			// 当你这样调用的时候，GIN 就会处理这个请求
			// 响应写回到 resp 里
			server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var res Result
			json.NewDecoder(resp.Body).Decode(&res)
			assert.Equal(t, tc.result, res)
		})
	}
}
