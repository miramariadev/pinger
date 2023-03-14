package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.studionx.ru/id/pinger/app/pinger/presentation/controllers"
	"gitlab.studionx.ru/id/pinger/app/pinger/presentation/models"
)

func TestMonitoringController_HandlePing(t *testing.T) {
	type fields struct {
	}

	type args struct {
		r *http.Request
	}

	type want struct {
		code        int
		response    models.PingResponseModel
		contentType string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "positive test",
			fields: fields{},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/pinger/ping", nil)},
			want: want{
				code:        200,
				contentType: "application/json",
				response: models.PingResponseModel{
					Method: "ping",
					Result: "pong",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &controllers.Controller{}

			w := httptest.NewRecorder()
			w.Header().Set("Content-Type", tt.want.contentType)

			h := http.HandlerFunc(controller.HandlePing)

			h.ServeHTTP(w, tt.args.r)

			response := models.PingResponseModel{}
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, response, tt.want.response)
			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))
		})
	}
}
