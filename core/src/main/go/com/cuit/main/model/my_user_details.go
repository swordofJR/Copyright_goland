package model

type MyUserDetails struct {
	User        User
	Permissions []Resource
}

func (m *MyUserDetails) GetAuthorities() []string {
	var authorities []string
	for _, permission := range m.Permissions {
		if len(permission.Value) > 0 {
			authorities = append(authorities, permission.Value)
		}
	}
	return authorities
}

func (m *MyUserDetails) GetPassword() string {
	return m.User.Password
}

func (m *MyUserDetails) GetUsername() string {
	return m.User.Username
}

func (m *MyUserDetails) IsAccountNonExpired() bool {
	return true
}

func (m *MyUserDetails) IsAccountNonLocked() bool {
	return true
}

func (m *MyUserDetails) IsCredentialsNonExpired() bool {
	return true
}

func (m *MyUserDetails) IsEnabled() bool {
	return true
}
