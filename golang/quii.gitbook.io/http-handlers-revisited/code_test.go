package httphandlersrevisited

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type MockUserService struct {
	RegisterFunc    func(user User) (string, error)
	UsersRegistered []User
}

func (m *MockUserService) Register(user User) (string, error) {
	m.UsersRegistered = append(m.UsersRegistered, user)
	return m.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
	t.Run("can register valid users", func(t *testing.T) {
		user := User{Name: "CJ"}
		expectedInsertID := "whatever"

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return expectedInsertID, nil
			},
		}
		server := NewUserServer(service)

		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(t, user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusCreated)

		if res.Body.String() != expectedInsertID {
			t.Errorf("expected body of %q but got %q", res.Body.String(), expectedInsertID)
		}

		if len(service.UsersRegistered) != 1 {
			t.Fatalf("expected 1 user added but got %d", len(service.UsersRegistered))
		}

		if !reflect.DeepEqual(service.UsersRegistered[0], user) {
			t.Errorf("the user registered %+v was not what was expected %+v", service.UsersRegistered[0], user)
		}
	})
	t.Run("returns 400 bad request if body is not valid user JSON", func(t *testing.T) {
		server := NewUserServer(nil)

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("trouble will find me"))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusBadRequest)
	})
	t.Run("returns 500 internal server error if the service fails", func(t *testing.T) {
		user := User{Name: "CJ"}

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return "", errors.New("couldn't add user")
			},
		}
		server := NewUserServer(service)

		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(t, user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusInternalServerError)
	})
}

func userToJSON(t testing.TB, user User) *bytes.Buffer {
	t.Helper()
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(user)
	return &buf
}

func assertStatus(t testing.TB, res *httptest.ResponseRecorder, status int) {
	t.Helper()
	if res.Code != status {
		t.Errorf("response code is not correct, got %d, want %d", status, res.Code)
	}
}
